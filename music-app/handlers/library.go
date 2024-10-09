package handlers

import (
	"context" 
	"music-app/models"
	"net/http"
	"music-app/utils"
	"github.com/gin-gonic/gin"
)

func AddSong(c *gin.Context) {
	db := utils.GetDB()
	var newSong models.Song
	if err := c.ShouldBindJSON(&newSong); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Недопустимый запрос"})
		return
	}

	// Логика для добавления песни в базу данных
	_, err := db.Exec(context.Background(), "INSERT INTO songs (title, artist) VALUES ($1, $2)", newSong.Title, newSong.Artist)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при добавлении песни"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Песня добавлена", "song": newSong})
}

func GetSongText(c *gin.Context) {
	id := c.Param("id")
	db := utils.GetDB()

	var lyrics string
	err := db.QueryRow(context.Background(), "SELECT lyrics FROM songs WHERE id = $1", id).Scan(&lyrics)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Песня не найдена"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"lyrics": lyrics})
}

func UpdateSong(c *gin.Context) {
	db := utils.GetDB()
	id := c.Param("id")
	var updatedSong models.Song
	if err := c.ShouldBindJSON(&updatedSong); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Недопустимый запрос"})
		return
	}

	// Логика для обновления песни в базе данных
	_, err := db.Exec(context.Background(), "UPDATE songs SET title = $1, artist = $2 WHERE id = $3", updatedSong.Title, updatedSong.Artist, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении песни"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Песня обновлена"})
}

func DeleteSong(c *gin.Context) {
	db := utils.GetDB()
	id := c.Param("id")

	// Логика для удаления песни из базы данных
	_, err := db.Exec(context.Background(), "DELETE FROM songs WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении песни"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Песня удалена"})
}

func GetLibrary(c *gin.Context) {
	db := utils.GetDB()
	page, limit := utils.GetPaginationParams(c)

	var songs []models.Song
	rows, err := db.Query(context.Background(), "SELECT * FROM songs ORDER BY id LIMIT $1 OFFSET $2", limit, (page-1)*limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить песни"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var song models.Song
		err := rows.Scan(&song.ID, &song.Title, &song.Artist) // Замените на свои поля
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сканировании данных"})
			return
		}
		songs = append(songs, song)
	}

	c.JSON(http.StatusOK, songs)
}
