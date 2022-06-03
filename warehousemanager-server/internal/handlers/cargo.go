package handlers

import (
	"context"
	"encoding/json"
	"github.com/martinkaburu/warehouse-manager/pkg/models"
	"github.com/martinkaburu/warehouse-manager/pkg/utils"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

const PageLimit int = 10

func CargoManifestHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		utils.SetResponseContentType(w)

		pageID := r.Context().Value("page_id")
		limit := r.Context().Value("limit")

		var cargo models.CargoList

		db.Preload("Orders").Where("id >= ?", pageID).Order("id").Limit(limit.(int) + 1).Find(&cargo.Cargo)

		if len(cargo.Cargo) == limit.(int)+1 {
			cargo.NextPageID = int(cargo.Cargo[len(cargo.Cargo)-1].ID)
			cargo.Cargo = cargo.Cargo[:limit.(int)] // this shortens the slice by 1
		}

		json.NewEncoder(w).Encode(cargo)

		return
	}
}

func Pagination(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		PageID := r.URL.Query().Get("page_id")
		limit := r.URL.Query().Get("limit")
		intPageID := 0
		intLimit := PageLimit
		var err error
		if PageID != "" {
			intPageID, err = strconv.Atoi(PageID)
			if err != nil {
				panic(err)
				return
			}
		}
		ctx := context.WithValue(r.Context(), "page_id", intPageID)
		if limit != "" {
			intLimit, err = strconv.Atoi(limit)
			if err != nil {
				panic(err)
				return
			}
		}
		ctx = context.WithValue(ctx, "limit", intLimit)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
