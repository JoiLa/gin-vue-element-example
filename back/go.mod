module Api

go 1.13

require (
	github.com/360EntSecGroup-Skylar/excelize v1.4.1 // indirect
	github.com/EDDYCJY/go-gin-example v0.0.0-20200505102242-63963976dee0 // indirect
	github.com/anthonynsimon/bild v0.13.0 // indirect
	github.com/astaxie/beego v1.12.2
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fogleman/gg v1.3.0
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.61.0
	github.com/go-openapi/spec v0.19.9 // indirect
	github.com/go-openapi/swag v0.19.9 // indirect
	github.com/go-playground/validator/v10 v10.3.0 // indirect
	github.com/go-redis/redis v6.14.2+incompatible
	github.com/gobuffalo/packr v1.30.1
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/manucorporat/try v0.0.0-20170609134256-2a0c6b941d52
	github.com/pkg/errors v0.9.1 // indirect
	github.com/robfig/cron v1.2.0 // indirect
	github.com/robfig/cron/v3 v3.0.0 // indirect
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/swaggo/swag v1.6.7 // indirect
	github.com/tealeg/xlsx v1.0.5 // indirect
	github.com/ugorji/go v1.1.8 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/image v0.0.0-20200801110659-972c09e46d76 // indirect
	golang.org/x/net v0.0.0-20200904194848-62affa334b73 // indirect
	golang.org/x/sys v0.0.0-20200918174421-af09f7315aff // indirect
	golang.org/x/tools v0.0.0-20200918232735-d647fc253266 // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/genproto v0.0.0-20200918140846-d0d605568037 // indirect
	google.golang.org/grpc v1.32.0 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

replace (
	./conf => ./vendor/github.com/EDDYCJY/go-gin-example/pkg/conf
	./middleware => ./vendor/github.com/EDDYCJY/go-gin-example/middleware
	./models => ./vendor/github.com/EDDYCJY/go-gin-example/models
	./pkg/e => ./vendor/github.com/EDDYCJY/go-gin-example/pkg/e
	./pkg/setting => ./vendor/github.com/EDDYCJY/go-gin-example/pkg/setting
	./pkg/util => ./vendor/github.com/EDDYCJY/go-gin-example/pkg/util
	./routers => ./vendor/github.com/EDDYCJY/go-gin-example/routers
)
