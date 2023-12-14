package connect

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	Conf "BACKEND/configs"
)
type User struct {
    ID   int64  `json:"id"`
    Name  string `json:"user_name"`
	FullName  string `json:"user_fullname"`
    Email string `json:"user_email"`
	Password string `json:"user_password"`
}

func TestConnect(c *gin.Context) {
	var db, err = Conf.Connectdb()
    if err!= nil {
        c.JSON(http.StatusNotFound, gin.H{"result":"Missing Connection"})
		log.Println("Missing Connection")
		return
    }
    defer db.Close()
    c.JSON(http.StatusOK, gin.H{
        "result": "Success Connection",
    })
}

func GetDataByUserName(c *gin.Context){
	var db,err = Conf.Connectdb()
	if err!= nil {
		c.JSON(http.StatusNotFound, gin.H{"result":"Missing Connection"})
		log.Println("Missing Connection")
		return
    }
	defer db.Close()
	if c.Query("values") =="" || c.Query("id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required parameter"})
		return
	}
	keyValues := c.Query("values")
	keyIds := c.Query("id")
	var users []User
	rows, err := db.Query("SELECT * FROM tbcluster where id = ? and user_name = ?", keyIds, keyValues)
	if err!= nil {
        c.JSON(http.StatusNotFound, gin.H{"result":"404 Page Not Found"})
        log.Println("404 Page Not Found")
        return
    }
	for rows.Next(){
		var user User
		err = rows.Scan(&user.ID, &user.Name, &user.FullName, &user.Email, &user.Password)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"result":"Not Found"})
			log.Println("Not Found")
			return
		}
		users = append(users, user)
	}	
	response := gin.H{
		"user": users,
		"values": keyValues,
        "id": keyIds,
	}

	c.JSON(http.StatusOK, response)
}