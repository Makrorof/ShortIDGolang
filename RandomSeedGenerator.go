package ShortID

import (
	"fmt"
	"sync"
	"time"
)

//////////////////////////////////////////////////////////
//https://github.com/twitter-archive/snowflake To GOLANG//
//////////////////////////////////////////////////////////

///
//Bits
///

//Max thread = 31
//Max datacenter = 31
//Max sequence = 4095
//MAX YEAR => 2058

const p_SEED_GENERATOR_ID_BITS int64 = 5 //Thread ID bits
const p_DATACENTER_ID_BITS int64 = 5
const p_SEQUENCE_BITS int64 = 12

const p_MAX_SEED_GENERATOR_ID int64 = -1 ^ (-1 << p_SEED_GENERATOR_ID_BITS)
const p_MAX_DATACENTER_ID int64 = -1 ^ (-1 << p_DATACENTER_ID_BITS)
const p_MAX_SEQUENCE int64 = -1 ^ (-1 << p_SEQUENCE_BITS)

const p_SEED_GENERATOR_ID_SHIFT = p_SEQUENCE_BITS
const p_DATACENTER_ID_SHIFT = p_SEQUENCE_BITS + p_SEED_GENERATOR_ID_BITS
const p_TIMESTAMP_LEFT_SHIFT = p_SEQUENCE_BITS + p_SEED_GENERATOR_ID_BITS + p_DATACENTER_ID_BITS

//GMT: Saturday, 22 October 2022 19:08:19
const CUSTOM_EPOCH int64 = 1666465699000

type randomSeedGenerator struct {
	generatorID  int64 //Thread ID
	datacenterID int64 //Datacenter ID

	sequence int64
	locker   sync.Mutex

	lastTimestamp int64
}

func MustCreateSeedGenerator(generatorID int64, datacenterID int64) *randomSeedGenerator {
	gen, err := CreateSeedGenerator(generatorID, datacenterID)

	if err != nil {
		panic(err)
	}

	return gen
}

func CreateSeedGenerator(generatorID int64, datacenterID int64) (*randomSeedGenerator, error) {
	if generatorID > p_MAX_SEED_GENERATOR_ID || generatorID < 0 {
		return nil, fmt.Errorf("generatorID can't be greater than %d or less than 0", p_MAX_SEED_GENERATOR_ID)
	}

	if datacenterID > p_MAX_DATACENTER_ID || datacenterID < 0 {
		return nil, fmt.Errorf("datacenterID can't be greater than %d or less than 0", p_MAX_DATACENTER_ID)
	}

	gen := &randomSeedGenerator{
		generatorID:   generatorID,
		datacenterID:  datacenterID,
		sequence:      0,
		lastTimestamp: 0,
	}

	gen.lastTimestamp = gen.getTime()

	return gen, nil
}

//Return random seed
func (rnd *randomSeedGenerator) GetSeed() int64 {
	return rnd.getRandomSeed()
}

func (rnd *randomSeedGenerator) getRandomSeed() int64 {
	rnd.locker.Lock()
	defer rnd.locker.Unlock()

	currentTimestamp := rnd.getTime()

	if rnd.lastTimestamp == currentTimestamp {
		rnd.sequence = (rnd.sequence + 1) & p_MAX_SEQUENCE

		if rnd.sequence == 0 {
			currentTimestamp = rnd.waitNextTime()
		}
	} else {
		rnd.sequence = 0
	}

	rnd.lastTimestamp = currentTimestamp

	currentTimestamp <<= p_TIMESTAMP_LEFT_SHIFT

	currentTimestamp |= (rnd.datacenterID << p_DATACENTER_ID_SHIFT)
	currentTimestamp |= (rnd.generatorID << p_SEED_GENERATOR_ID_SHIFT)
	currentTimestamp |= rnd.sequence

	return currentTimestamp
}

func (rnd *randomSeedGenerator) waitNextTime() int64 {
	timeNano := time.Now().UnixNano()

	for timeNano <= rnd.lastTimestamp {
		timeNano = rnd.getTime()
	}

	return timeNano
}

func (rnd *randomSeedGenerator) getTime() int64 {
	return time.Now().UnixMilli() - CUSTOM_EPOCH
}
