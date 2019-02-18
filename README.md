# Welcome to Augeas!

A service tells you where and how will it take to park your car!



##  Install

You'll need at least Go-1.11 to build the project: [Here](https://github.com/golang/go/wiki/Ubuntu) is how.

## Build

Just type `make build`

## Tests

Just type `make test`

## Run the service

You'll need a CSV which describes all car parks in your Augeas
Here is an example of the CSV

```
poi_id;poi_name;poi_weight;poi_visible;poi_lat;poi_lon;poi_type_id;poi_address_number;poi_address_name
937854398;"Foch";0;1;48.872505;2.285156;amenity:parking;;""
 ```

Run `./augeas --help` for more information about the usage of the command. 

Once you have your CSV correctly formated, just type `./augeas -poi your_csv`
 
## Available APIs

The service runs on `:1337` by default 

`localhost:1337/v0/car_parks`: returns a list of all imported car parks

`localhost:1337/v0/car_parks/id`: returns a car park with id 

`localhost:1337/v0/park_duration?lon=2.300731&lat=48.87445`: Given the coordinate of your location, the api will return a  list of nearest car parks with distances and durations by walk from your location. 
Parameters of this end point: 
`lon` and `lot`: the coordinate of  your location

`n`: max number of car parks that Augeas should return, `n=5` by default 

`walking_speed`: your waking speed in `meter/second`, `walking_speed=1.11` by default

`max_park_duration`: the max time(in second) it'll take to park your car, all car parks whose `duration` are greater than this value will be filtered,  `walking_speed=1200`
   

 



