package mqClient
import(
	// "time"
)

func MessageProducer(msg Message){
	// message := Message{userId,"message created"}
	// time.Sleep(time.Second*4)
	SendHelper(msg)
}