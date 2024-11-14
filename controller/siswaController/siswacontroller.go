package siswacontroller

import (
	"net/http"
	"strconv"

	"github.com/Red651/belajar-web-api/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var siswa []database.Siswa

	database.DB.Find(&siswa)

	c.JSON(http.StatusOK, gin.H{
		"siswa": siswa,
	})

}

func ShowById(c *gin.Context) {
	var siswa database.Siswa

	id := c.Param("id")
	if err := database.DB.First(&siswa, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"massage": "data tidak ditemukan"})
			return

		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"massage": err.Error()})
			return
		}

	}
	c.JSON(http.StatusOK, gin.H{"siswa": siswa})

}

func ShowByNama(c *gin.Context) {
	var siswa []database.Siswa

	nama := c.Param("nama")
	if err := database.DB.Where("nama = ?", nama).Find(&siswa).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"massage": "data tidak ditemukan"})
			return

		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"massage": err.Error()})
			return
		}

	}
	c.JSON(http.StatusOK, gin.H{"siswa": siswa})

}

func Insert(c *gin.Context) {
	var siswa database.Siswa

	if err := c.ShouldBindJSON(&siswa); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&siswa).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil disimpan", "data": siswa})
}

func Update(c *gin.Context) {
	var siswa database.Siswa

	id := c.Param("id")

	if err := c.ShouldBindJSON(&siswa); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result := database.DB.Model(&database.Siswa{}).Where("id = ?", id).Updates(siswa)

	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal mengupdate data: " + result.Error.Error(),
		})
		return
	}

	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Data tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diubah"})
}

func Delete(c *gin.Context) {
	var siswa database.Siswa

	id := c.Param("id")

	if idInt, err := strconv.ParseInt(id, 10, 64); err == nil {
		if database.DB.Delete(&siswa, idInt).RowsAffected == 0 {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Tidak dapat menemukan data"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "ID tidak valid"})
	}
}
