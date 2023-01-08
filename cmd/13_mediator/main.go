package mediator

/*
クライアント・コード
*/
func main() {
	stationManager := newStationManager()

	passengerTrain := &PassengerTrain{
		mediator: stationManager,
	}
	freightTrain := &FreightTrain{
		mediator: stationManager,
	}

	passengerTrain.arrive()
	freightTrain.arrive()
	passengerTrain.depart()
}

/*
実行結果
PassengerTrain: Arrived
FreightTrain: Arrival blocked, waiting
PassengerTrain: Leaving
FreightTrain: Arrival permitted
FreightTrain: Arrived
*/
