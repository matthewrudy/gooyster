GoOyster
========

A simple implementation of the OysterCard

## Run

```
$ make

OysterCard simulation!!!!!

Action: Fresh new card
Balance: £0.00

Action: Topped up £30
Balance: £30.00

Action: Entered tube at Holborn
Balance: £26.80

Action: Exited tube at Earl's Court
Balance: £27.50

Action: Got on Bus at Earl's Court
Balance: £25.70

Action: Got off Bus at Chelsea
Balance: £25.70

Action: Entered tube at Earl's Court
Balance: £22.50

Action: Exited tube at Hammersmith
Balance: £23.70
```

## Test

```
$ make test

go test -v ./...
=== RUN   TestCard_Balance
--- PASS: TestCard_Balance (0.00s)
=== RUN   TestCard_Enter_and_Exit
--- PASS: TestCard_Enter_and_Exit (0.00s)
=== RUN   TestJourney_Fare_BusIsAlways_180
--- PASS: TestJourney_Fare_BusIsAlways_180 (0.00s)
=== RUN   TestJourney_Fare_StartsAtMaximum
--- PASS: TestJourney_Fare_StartsAtMaximum (0.00s)
=== RUN   TestJourney_Fare_Zone1
--- PASS: TestJourney_Fare_Zone1 (0.00s)
=== RUN   TestJourney_Fare_Zone2
--- PASS: TestJourney_Fare_Zone2 (0.00s)
=== RUN   TestJourney_Fare_Zone3
--- PASS: TestJourney_Fare_Zone3 (0.00s)
=== RUN   TestJourney_Fare_Zone1To2
--- PASS: TestJourney_Fare_Zone1To2 (0.00s)
=== RUN   TestJourney_Fare_Zone2To3
--- PASS: TestJourney_Fare_Zone2To3 (0.00s)
=== RUN   TestJourney_Fare_Zone1To3
--- PASS: TestJourney_Fare_Zone1To3 (0.00s)
PASS
ok  	github.com/matthewrudy/gooyster	(cached)
```