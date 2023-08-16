## telegram bot allows you get ecology data in Ukraine

###

<div align="center">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="150" alt="go logo"  />
  <img width="25" />
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/docker/docker-original.svg" height="150" alt="docker logo"  />
  <img width="25" />
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/mongodb/mongodb-original.svg" height="150" alt="mongodb logo"  />
</div>

###

##
- ecoman is a telegram bot made with golang
- you are able to choose city and station you interested in

## API in use:
- https://www.saveecobot.com
##

## data you get example
```md
🏙️ City: Kyiv
🏠 Station: Street Henerala Zhmachenka, 4
🧭 Latitude: 50.458479
🧭 Longitude: 30.6079481
🕒 Timezone: +0300
💎 Pollutant: Humidity
   + Units: %
   + Time: 2020-05-25T23:00:22.000000Z
   + Value: 100
   + Average: 2 minutes
💎 Pollutant: PM10
   + Units: ug/m3
   + Time: 2020-05-25T23:00:22.000000Z
   + Value: 8.47
   + Average: 2 minutes
💎 Pollutant: PM2.5
   + Units: ug/m3
   + Time: 2020-05-25T23:00:22.000000Z
   + Value: 5
   + Average: 2 minutes
💎 Pollutant: Pressure
   + Units: hPa
   + Time: 2020-05-21T18:32:09.000000Z
   + Value: 998.38
   + Average: 2 minutes
💎 Pollutant: Temperature
   + Units: Celcius
   + Time: 2020-05-25T23:00:22.000000Z
   + Value: 10.94
   + Average: 2 minutes
💎 Pollutant: Air Quality Index
   + Units: aqi
   + Time: 2020-05-26 00:00:00
   + Value: 23
   + Average: 1 hour
```
- after getting data user is asked:
```
would you like to receive advice on what is best to do on this day (based on data you got)?
```
- if answer "yes" -> user will receive advice, example:

```md
Based on the data provided for Kyiv on this day, here are some advice:

1. Humidity is at 100%.
It might be a good idea to stay indoors or find activities
that can be done inside to avoid discomfortcaused by the high humidity.

2. PM10 levels are at 8.47 ug/m3, which is considered low.
However, it is still advisable to reduce exposure to outdoor air pollution.
If possible, avoid spending extended periods of time in heavily polluted areas.

3. PM2.5 levels are at 5 ug/m3, also considered low. Nonetheless,
take precautions if you have any respiratory conditions.
Consider wearing a mask if necessary.

4. The pressure is at 998.38 hPa,
indicating relatively stable weather conditions.
It would be a good day for outdoor activities or exploring the city.

5. The temperature is 10.94 degrees Celsius.
Dress accordingly to stay comfortable outdoors.
You might need a light jacket or sweater.

6. The Air Quality Index (AQI) is at 23, indicating good air quality.
Enjoy outdoor activities without any major concerns for air pollution.

Remember, these recommendations are based on the available data and general guidelines.
Use your judgment and take into account any personal health considerations or local regulations.
```

## installation

```shell
git clone https://github.com/kenjitheman/eco_tg_bot 
```
- to install all dependencies
```shell
go mod tidy
```

## usage

- create .env file with:

```.env
TELEGRAM_APITOKEN=YOUR_TELEGRAM_API_TOKEN
USERNAME=YOUR_MONGODB_USERNAME
PASSWORD=YOUR_MONGODB_PASSWORD
OPENAI_APITOKEN=YOUR_OPENAI_API_TOKEN
```

- run docker container:

```shell
docker build -t your_image_name .
docker run -d -p 8080:80 your_image_name
```

- or run:

```shell
cd cmd
go run main.go
```

### ! you need your own mongo database and collection configured to run this project !

## contributing

- pull requests are welcome, for major changes, please open an issue first
to discuss what you would like to change

## license

- [MIT](https://choosealicense.com/licenses/mit/)
