package app

import (
	"awesomeProject1/internal/app/model"
	"awesomeProject1/swagger/comics"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (a *Application) StartServer() {
	log.Println("server start up")

	r := gin.Default()

	r.GET("/comics", a.GetComics)

	r.GET("/comics/:uuid", a.GetComicsPrice)

	r.POST("/comics", a.CreateComics)

	r.PUT("/comics/:uuid", a.ChangePrice)

	r.DELETE("/comics/:uuid", a.DeleteComics)

	_ = r.Run()

	log.Println("server down")
}

// GetComics 		godoc
// @Summary      	Get all records
// @Description  	Get a list of all comics
// @Tags         	Info
// @Produce      	json
// @Success      	200 {object} model.Comics
// @Failure 		500 {object} comics.ComicsError
// @Router       	/comics [get]
func (a *Application) GetComics(gCtx *gin.Context) {
	resp, err := a.repo.GetComics()
	if err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&comics.ComicsError{
				Description: "Can't get a list of comics codes",
				Error:       comics.Err500,
				Type:        comics.TypeInternalReq,
			})
		return
	}

	gCtx.JSON(http.StatusOK, resp)
}

// GetComicsPrice  	godoc
// @Summary      	Get price for a comics
// @Description  	Get the price using the comics uuid
// @Tags         	Info
// @Produce      	json
// @Param 			UUID query string true "UUID комикс" format(uuid)
// @Success      	200 {object} comics.ComicsPrice
// @Failure 	 	400 {object} comics.ComicsError
// @Failure 	 	404 {object} comics.ComicsError
// @Failure 	 	500 {object} comics.ComicsError
// @Router       	/comics/:uuid [get]
func (a *Application) GetComicsPrice(gCtx *gin.Context) {
	UUID, err := uuid.Parse(gCtx.Param("uuid"))
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&comics.ComicsError{
				Description: "Invalid UUID format",
				Error:       comics.Err400,
				Type:        comics.TypeClientReq,
			})
		return
	}

	var comics1 model.Comics
	code, err := a.repo.GetComicsPrice(UUID, &comics1)
	if err != nil {
		if code == 404 {
			gCtx.JSON(
				http.StatusNotFound,
				&comics.ComicsError{
					Description: "UUID Not Found",
					Error:       comics.Err404,
					Type:        comics.TypeClientReq,
				})
			return
		} else {
			gCtx.JSON(
				http.StatusInternalServerError,
				&comics.ComicsError{
					Description: "Get comics price failed",
					Error:       comics.Err500,
					Type:        comics.TypeInternalReq,
				})
			return
		}
	}

	gCtx.JSON(
		http.StatusOK,
		&comics.ComicsPrice{
			Price: comics1.Price,
		})
}

// CreateComics		godoc
// @Summary     	Add a new comics
// @Description		Adding a new comics to database
// @Tags			Add
// @Produce      	json
// @Param 			Name query string true "Название"
// @Param 			Price query uint64 true "Цена"
// @Param 			Year query uint64 true "Год выпуска"
// @Success 		201 {object} comics.ComicsCreated
// @Failure 		400 {object} comics.ComicsError
// @Failure 		500 {object} comics.ComicsError
// @Router  		/comics [post]
func (a *Application) CreateComics(gCtx *gin.Context) {
	comics1 := model.Comics{}
	err := gCtx.BindJSON(&comics1)
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&comics.ComicsError{
				Description: "Invalid parameters",
				Error:       comics.Err400,
				Type:        comics.TypeClientReq,
			})
		return
	}

	err = a.repo.AddComics(comics1)
	if err != nil {
		gCtx.JSON(
			http.StatusInternalServerError,
			&comics.ComicsError{
				Description: "Create failed",
				Error:       comics.Err500,
				Type:        comics.TypeInternalReq,
			})
		return
	}

	gCtx.JSON(
		http.StatusCreated,
		&comics.ComicsCreated{
			Success: true,
		})
}

// ChangePrice		godoc
// @Summary      	Change comics price
// @Description  	Change the comics price using its uuid
// @Tags         	Change
// @Produce      	json
// @Param 			UUID query string true "UUID комикс" format(uuid)
// @Param 			Price query uint64 true "Новая цена"
// @Success      	200 {object} comics.ComicsChanged
// @Failure 		400 {object} comics.ComicsError
// @Failure 		404 {object} comics.ComicsError
// @Failure 	 	500 {object} comics.ComicsError
// @Router       	/comics/:uuid [put]
func (a *Application) ChangePrice(gCtx *gin.Context) {
	UUID, err := uuid.Parse(gCtx.Param("uuid"))
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&comics.ComicsError{
				Description: "Invalid UUID format",
				Error:       comics.Err400,
				Type:        comics.TypeClientReq,
			})
		return
	}

	comics1 := model.Comics{}
	err = gCtx.BindJSON(&comics1)
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&comics.ComicsError{
				Description: "The price is negative or not int",
				Error:       comics.Err400,
				Type:        comics.TypeClientReq,
			})
		return
	}

	code, err := a.repo.ChangePrice(UUID, comics1.Price)
	if err != nil {
		if code == 404 {
			gCtx.JSON(
				http.StatusNotFound,
				&comics.ComicsError{
					Description: "UUID Not Found",
					Error:       comics.Err404,
					Type:        comics.TypeClientReq,
				})
			return
		} else {
			gCtx.JSON(
				http.StatusInternalServerError,
				&comics.ComicsError{
					Description: "Change failed",
					Error:       comics.Err500,
					Type:        comics.TypeInternalReq,
				})
			return
		}
	}

	gCtx.JSON(
		http.StatusOK,
		&comics.ComicsChanged{
			Success: true,
		})
}

// DeleteComics		godoc
// @Summary     	Delete a comics
// @Description 	Delete a comics using its uuid
// @Tags         	Delete
// @Produce      	json
// @Param 			UUID query string true "UUID комикс" format(uuid)
// @Success      	200 {object} comics.ComicsDeleted
// @Failure 		400 {object} comics.ComicsError
// @Failure 		404 {object} comics.ComicsError
// @Failure 	 	500 {object} comics.ComicsError
// @Router       	/comics/:uuid [delete]
func (a *Application) DeleteComics(gCtx *gin.Context) {
	UUID, err := uuid.Parse(gCtx.Param("uuid"))
	if err != nil {
		gCtx.JSON(
			http.StatusBadRequest,
			&comics.ComicsError{
				Description: "Invalid UUID format",
				Error:       comics.Err400,
				Type:        comics.TypeClientReq,
			})
		return
	}

	code, err := a.repo.DeleteComics(UUID)
	if err != nil {
		if code == 404 {
			gCtx.JSON(
				http.StatusNotFound,
				&comics.ComicsError{
					Description: "UUID Not Found",
					Error:       comics.Err404,
					Type:        comics.TypeClientReq,
				})
			return
		} else {
			gCtx.JSON(
				http.StatusInternalServerError,
				&comics.ComicsError{
					Description: "Delete failed",
					Error:       comics.Err500,
					Type:        comics.TypeInternalReq,
				})
			return
		}
	}

	gCtx.JSON(
		http.StatusOK,
		&comics.ComicsDeleted{
			Success: true,
		})
}
