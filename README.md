# Huawei Health TCX

Golang app that generates TCX files from Huawei HiTrack dump. Hitrack data is what Huawei wearables generates after an activiry ([Huawei Band 3 PRO](https://consumer.huawei.com/en/wearables/band3-pro/) is an example).

The app gets the HiTrack data reading it from the Sqlite3 database you can obtain from the Huawei Health app backup.
The outputted `.TCX` files will contain timestamped GPS, altitude, heart-rate, and cadence data where available.

This app gets inspiration from [Huawei TCX Converter](https://github.com/aricooperdavis/Huawei-TCX-Converter) which should be used if your backup has a different format (see below).

## How to get the Huawei Health db

- Open the Huawei Health app and open the exercise that you want to convert to view it's trajectory. This ensures that its HiTrack file is generated.
- Download the [Huawei Backup App](https://play.google.com/store/apps/details?id=com.huawei.KoBackup&hl=en_GB) onto your phone.
- Start a new **unencrypted** backup of the Huawei Health app data to your external storage (SD Card)
- Connect the phone to the pc using a USB cable. When prompted on the phone, authorize the pc to access data.
- Navigate to `/HuaweiBackup/backupFiles/<backup folder>/` and copy `com.huawei.health.db` to your computer. If you can't find the `.db` file but you find the `com.huawei.health.tar` file, than you should use [Huawei TCX Converter](https://github.com/aricooperdavis/Huawei-TCX-Converter).

## How to run the app

```
git clone git@github.com:tommyblue/go-huawei-health-tcx.git
./scripts/setup
./dist/ghht <path to db file>
```

In the folder you're running the app from, you'll find a TCX file for each activity.

## How to compile

* [go-sqlite3](https://github.com/mattn/go-sqlite3)

```
go get <dependencies>
go get <repo>
go build
go install
```

## Contributing

TODO
