package ShortID

import (
	"log"
	"math"
	"sync"
	"testing"
	"time"
)

func TestRuneDataShuffle(t *testing.T) {
	list := make([]string, 0)

	wg := sync.WaitGroup{}

	for i := 0; i < 55; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			runeData, _ := createCharacterRuneData(DefaultCharacters, GetSeed())

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

	var datacenterId int64 = 1024
	var seedGeneratorID int64 = 1024
	var sequence int64 = 4095

	log.Println(time.Now().UnixMilli())
	//1666464163170686300
	//2802536687000686300

	for {
		timestamp := 2802536687000686300 - CUSTOM_EPOCH

		log.Println("currentTimestamp =>", timestamp)

		timestamp <<= TIMESTAMP_LEFT_SHIFT

		log.Println("TIMESTAMP_LEFT_SHIFT =>", timestamp)

		timestamp |= datacenterId << DATACENTER_ID_SHIFT
		log.Println("DATACENTER_ID_SHIFT =>", timestamp)
		timestamp |= seedGeneratorID << SEED_GENERATOR_ID_SHIFT
		log.Println("SEED_GENERATOR_ID_SHIFT =>", timestamp)
		timestamp |= sequence
		log.Println("sequence =>", timestamp)

		time.Sleep(time.Second)

		sequence++
	}

	/*
		2022/10/22 21:36:45 currentTimestamp => 663454479600
		2022/10/22 21:36:45 TIMESTAMP_LEFT_SHIFT => 2782729777604198400
		2022/10/22 21:36:45 DATACENTER_ID_SHIFT => 2782729777604198400
		2022/10/22 21:36:45 SEED_GENERATOR_ID_SHIFT => 2782729777604198400
		2022/10/22 21:36:45 sequence => 2782729777604198400

		2022/10/22 21:36:46 currentTimestamp => 664455645900
		2022/10/22 21:36:46 TIMESTAMP_LEFT_SHIFT => 2786928973420953600
		2022/10/22 21:36:46 DATACENTER_ID_SHIFT => 2786928973420953600
		2022/10/22 21:36:46 SEED_GENERATOR_ID_SHIFT => 2786928973420953600
		2022/10/22 21:36:46 sequence => 2786928973420953601
	*/
}

func TestSequence(t *testing.T) {
	var sequence int64 = 0

	const SEQUENCE_BITS int64 = 12
	max_sequence_bits := math.Pow(2, float64(SEQUENCE_BITS)) - 1

	for {
		sequence = (sequence + 1) & int64(max_sequence_bits)
		log.Println(sequence)
		//time.Sleep(time.Second)

		if sequence == 0 {
			break
		}
	}
}

func TestRSG(t *testing.T) {
	for {
		log.Println(GetSeed())
		time.Sleep(time.Second)
	}
}

func TestSameRSG(t *testing.T) {
	list := make([]int64, 0)
	//sameList := make([]int64, 0)
	wg := sync.WaitGroup{}

	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			list = append(list, GetSeed())
		}()
	}

	wg.Wait()

	for _, left := range list {
		count := 0
		for _, right := range list {
			if left == right {
				count++
			}
		}

		if count >= 2 {
			log.Println("Same Seed:", left, " Count:", count)
			t.Fatal("Same seed")
			//sameList = append(sameList, left)
		}
	}

	for _, seed := range list {
		log.Println(seed)
	}
}

func TestShortID(t *testing.T) {
	shortID, err := CreateShortID("x12", 120)

	if err != nil {
		t.Fatal(err)
	}

	log.Println("ID:", shortID.Generate())
}
