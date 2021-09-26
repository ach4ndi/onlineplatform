first install golang library on https://golang.org/. you can choose what your platform

 ## Running Local
You should install some module first before running the application :

``` bash
go get github.com/badoux/checkmail
go get github.com/jinzhu/gorm
go get golang.org/x/crypto/bcrypt
go get github.com/dgrijalva/jwt-go
go get github.com/gorilla/mux
go get github.com/jinzhu/gorm/dialects/mysql
go get github.com/jinzhu/gorm/dialects/sqlite
go get github.com/joho/godotenv
go get gopkg.in/go-playground/assert.v1
go get github.com/cloudinary/cloudinary-go
go get github.com/cbrake/golang.org-x-image
```
After already install all module, then running go run main.go

before run `go run main.go`
```
# cloudinary
CLOUDINARY_CLOUDNAME={get from cloudinary}
CLOUDINARY_APIKEY={get from cloudinary}
CLOUDINARY_APISecret={get from cloudinary}

# local image
IMG_DIR=images
IMG_LIMIT=10

# select sqlite3
#DB_DRIVER=sqlite3
#API_SECRET=
#DB_NAME=andionlinecouse.db

# Mysql
DB_HOST=127.0.0.1
DB_DRIVER=mysql 
API_SECRET= 91749ywifhakn
DB_USER=andi
DB_PASSWORD=12345abcderf
DB_NAME=andi_onlinecouse
DB_PORT=3306 #Default mysql port

#Server
WEB_PORT=:8080
LIMITLV=1
SEED_LOAD=0
```
 notes:
 

 - if you want regenerate sampledata from seed data, just set SEED_LOAD to 1.
 - LIMITLV is confugiration number to check the user is admin or not, make sure add data user status with level_num following LIMITLV .
	 ```
	{
		"level_name": "Admin",
		"level_num":1
	}
	```
- if you want using cloudinary, register to get api key.
- 