package ShortID

import (
	"log"
	"sync"
	"testing"
	"time"
)

func TestRuneDataShuffle(t *testing.T) {
	list := make([]string, 0)

	wg := sync.WaitGroup{}

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			runeData, _ := createCharacterRuneData(DefaultCharacters, time.Now().UnixNano())

			runeData.shuffle()
			str := string(runeData.GetCharacters())
			list = append(list, str)
			//	log.Println(str)
		}()
	}

	wg.Wait()

	for _, left := range list {
		count := 0
		for _, right := range list {
			if left == right {
				count++

				if count >= 2 {
					t.Fatal("Same characters")
				}
			}
		}
	}

	for _, characters := range list {
		log.Println(characters)
	}

	//for _, characters := range list {
	//	log.Println(characters)
	//}
}

func TestBit(t *testing.T) {
	const SEQUENCE_BITS int64 = 41
	const NODEID_BITS int64 = 41

	//timeNano := time.Now().UnixNano()
	var timeNano int64 = 1666395678157808600

	var nodeID int64 = 20000000
	var sequence int64 = 20000000

	log.Println("=>", timeNano)

	timeNano |= (sequence << SEQUENCE_BITS)
	timeNano |= (nodeID << NODEID_BITS)

	log.Println("=>", timeNano)

	log.Println("(sequence) =>", (sequence << SEQUENCE_BITS))
	log.Println("(nodeID) =>", (nodeID << NODEID_BITS))
	//1666395678157808600 // normal
	//1666395678157 812697 // 1-1
	//1666395 953035 715545
	//1666395 678157 808601
	//1666395 678157 808601
	//1666395 678157 808601
	//1666395 678159 905753
	//1666402000349668351

	//1666395678157808600
	//1667596344855340027
	//1666537240279884795
	//1666395678828897275
	//1666395678159909883
	//1666395678158338043
}
