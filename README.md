# ecoman

### Telegram bot allows you get ecology data in Ukraine

<div align="center">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="150" alt="go logo"  />
  <img width="25" />
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/docker/docker-original.svg" height="150" alt="docker logo"  />
  <img width="25" />
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/mongodb/mongodb-original.svg" height="150" alt="mongodb logo"  />
</div>

## Project structure

```go
.
├── api
│   ├── api.go
│   └── api_test.go
├── bot
│   ├── bot.go
│   ├── core.go
│   ├── keyboards.go
│   └── vars.go
├── db
│   ├── db.go
│   └── db_test.go
├── Dockerfile
├── go.mod
├── go.sum
├── LICENSE
├── main.go
├── openai
│   ├── openai.go
│   └── openai_test.go
├── README.md
└── vars
    └── vars.go
```

## Installation

```sh
git clone https://github.com/kenjitheman/ecoman
```

## Usage

- Create .env file with:

```.env
TELEGRAM_APITOKEN=YOUR_TELEGRAM_API_TOKEN
OPENAI_APITOKEN=YOUR_OPENAI_API_TOKEN
USERNAME=YOUR_MONGODB_USERNAME
PASSWORD=YOUR_MONGODB_PASSWORD
MONGO_URI=YOUR_MONGO_URI
```

- Verify your dependencies:

```sh
go mod tidy
```

- To run:

```sh
go run main.go
```

- Run it using docker
- You need to paste your .env variables

```dockerfile
ENV TELEGRAM_APITOKEN=YOUR_TELEGRAM_API_TOKEN
ENV OPENAI_APITOKEN=YOUR_OPENAI_API_TOKEN
ENV USERNAME=YOUR_MONGODB_USERNAME
ENV PASSWORD=YOUR_MONGODB_PASSWORD
ENV MONGO_URI=YOUR_MONGO_URI
```

```sh
docker build -t your_image_name .
docker run -d -p 8080:80 your_image_name
```

## Data you get example:

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

- After getting data user is asked:

```
Would you like to receive advice on what is best to do on this day?
```

- If your answer is yes - user will receive an advice, for example:

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

## Contributing

- Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

- Please make sure to update tests as appropriate.

## License

- [MIT](https://choosealicense.com/licenses/mit/)
