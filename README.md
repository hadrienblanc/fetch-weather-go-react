# go-reactjs-weather

Welcome,

This project is :
- a Go server
- a react front-end in the front folder


## The GO Server

The go server fetched on the openweather API the current pollution and the average of :
- the weather conditions of now
- the weather conditions in 3 hours
- the weather conditions in 6 hours

The Go server converts the weather condition to a RVB html color code. (Example : #f49A66)

## The React.js Front


## Launch the app

### Back-end

To launch the Go server :
```
go run app.go
```

To test :
```
curl localhost:8080/
```
```
curl localhost:8080/Paris
```

### Front-end

To launch the React front-end (with npm):
```
cd front
cd my-app
npm install
npm start
```

On a Browser :  [http://localhost:3000](http://localhost:3000/)
