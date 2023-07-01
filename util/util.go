package util

func GetRsrq(rsrq int) int {
	rsrq = rsrq * -1
	switch {
	case rsrq <= 10:
		return 100
	case rsrq <= 9 || rsrq <= 14:
		return 75 * rsrq / 14
	case rsrq <= 16 || rsrq <= 19:
		return 50 * rsrq / 19
	case rsrq >= 20:
		return 0
	}
	return 0
}

func GetRsp(rsrp int) int {
	rsrp = rsrp * -1
	switch {
	case rsrp <= 80:
		return 100
	case rsrp <= 81 || rsrp <= 90:
		return 75 * rsrp / 90
	case rsrp <= 91 || rsrp <= 99:
		return 50 * rsrp / 99
	case rsrp >= 100:
		return 0
	}
	return 0
}

func GetSirn(sinr int) int {
	switch {
	case sinr > 20:
		return 100
	case sinr > 13 || sinr >= 20:
		return 75 / 20 * sinr
	case sinr > 0 || sinr >= 13:
		return 50 / 13 * sinr
	case sinr <= 0:
		return 0
	}
	return 0
}

func GetRssiPercentage(rssi int) int {
	rssi = rssi * -1
	switch {
	case rssi <= 65:
		return 100
	case rssi <= 64 || rssi <= 74:
		return 80 * rssi / 74
	case rssi <= 75 || rssi <= 85:
		return 60 * rssi / 85
	case rssi <= 86 || rssi <= 94:
		return 40 * rssi / 94
	case rssi >= 95:
		return 0
	}
	return 0
}

func ByteToMB(b float64) int64 {
	return int64(b * 8 / 1000000)
}

func GetRSSI(rssi int) int {
	return -128 + rssi
}
