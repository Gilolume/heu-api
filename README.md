# heu-api
api rest fo rheu project

Create `heu/api` image :
```sh
docker build -t heu/api .
```

Create `config.json` file with your key for use google speech api :

Example : 
```
{
    "google_speech_key": "xxxxx"
}
```

Start heu-api :
```sh
docker-compose up -d
```
