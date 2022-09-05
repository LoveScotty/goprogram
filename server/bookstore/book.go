package bookstore

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/LoveScotty/goprogram/internal/store/dto/bookstore"
)

func (srv *Server) bookRouter(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/create", srv.createBookHandler)
	routerGroup.PUT("/update", srv.updateBookHandler)

	routerGroup.GET("/get/:id", srv.getBookHandler)
	routerGroup.GET("/get/all", srv.getAllBooksHandler)
	routerGroup.DELETE("/del/:id", srv.delBookHandler)
}

func (srv *Server) createBookHandler(c *gin.Context) {
	var book bookstore.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := srv.s.Book().Add(c, &book); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func (srv *Server) updateBookHandler(c *gin.Context) {
	var book bookstore.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err := srv.s.Book().Update(c, &book); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func (srv *Server) getBookHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	book, err := srv.s.Book().Get(c, id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, book)
}

func (srv *Server) getAllBooksHandler(c *gin.Context) {
	bookList, err := srv.s.Book().All(c)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, bookList)
}

func (srv *Server) delBookHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = srv.s.Book().Delete(c, id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}
