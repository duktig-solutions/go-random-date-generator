package randomDataTime

import (
	"crypto/rand"
	"math/big"
	"time"
	"errors"
)

// GenerateDate - Random date between given dates
// Example:
//   date, err := GenerateDate("2022-01-22", "2022-06-25")
func GenerateDate(startDateStr, endDateStr string) (string, error) {

	startTime, err := time.Parse("2006-01-02", startDateStr)

	if err != nil {
		return "", err
	}

	endTime, err := time.Parse("2006-01-02", endDateStr)

	if err != nil {
		return "", err
	}

	randomTime, err := genRandNum(startTime.Unix(), endTime.Unix())

	if err != nil {
		return "", err
	}

	return time.Unix(randomTime, 0).Format("2006-01-02"), nil

}

// GenerateTime - Random Time between given times
// Example:
//   _time, err := GenerateTime(10, 18)
func GenerateTime(startTime, endTime int) (string, error) {

	if startTime < 0 || startTime > 23 {
		return "", errors.New("invalid startTime. must be 0 - 23")
	}

	if endTime < 1 || endTime > 23 {
		return "", errors.New("invalid endTime. must be 1 - 23")
	}

	timeFrom := time.Date(1970, 01, 01, startTime, 00, 00, 00, time.UTC)
	timeTo := time.Date(1970, 01, 01, endTime, 00, 00, 00, time.UTC)
	randomTime, err := genRandNum(timeFrom.Unix(), timeTo.Unix())

	if err != nil {
		return "", err
	}

	return time.Unix(randomTime, 0).UTC().Format("15:04"), nil

}

// GenerateDateTime - Random Date and Time between given values
// Example:
//   _Datetime, err := GenerateDateTime("2022-01-01 00:00:00", "2022-08-21 17:08:26")
func GenerateDateTime(startDateTime, endDateTime string) (string, error) {

	startTime, err := time.Parse("2006-01-02 15:04:05", startDateTime)

	if err != nil {
		return "", err
	}

	endTime, err := time.Parse("2006-01-02 15:04:05", endDateTime)

	if err != nil {
		return "", err
	}

	randomTime, err := genRandNum(startTime.Unix(), endTime.Unix())

	if err != nil {
		return "", err
	}

	return time.Unix(randomTime, 0).Format("2006-01-02 15:04:05"), nil

}

// GenerateDOB - Random DOB (Date of birt) between given values
// Example:
//   _dob, err := GenerateDOB(18, 65)
func GenerateDOB(ageFrom, ageTo int64) (string, error) {

	ageTimeFrom := time.Now().Add(-time.Hour * time.Duration(ageFrom*365*24))
	ageTimeTo := time.Now().Add(-time.Hour * time.Duration(ageTo*365*24))
	randomAgeUnix, err := genRandNum(ageTimeTo.Unix(), ageTimeFrom.Unix())

	if err != nil {
		return "", err
	}

	return time.Unix(randomAgeUnix, 0).Format("2006-01-02"), nil

}

func genRandNum(min, max int64) (int64, error) {

	// calculate the max value
	bg := big.NewInt(max - min)

	// get big.Int between 0 and bg
	// in this case 0 to 20
	n, err := rand.Int(rand.Reader, bg)
	if err != nil {
		return 0, err
	}

	// add n to min to support the passed in range
	return n.Int64() + min, nil
}

/*
func main() {

	for i := 0; i <= 100; i++ {
		date, err := GenerateDate("2010-08-01", "2022-08-01")

		if err != nil {

		}

		fmt.Println(date)
	}

	for i := 0; i <= 100; i++ {
		dob, err := GenerateDOB(18, 80)
		if err != nil {

		}

		fmt.Println(dob)
	}

	for i := 0; i <= 100; i++ {
		tm, err := GenerateTime(1, 8)
		if err != nil {

		}

		fmt.Println(tm)
	}

	for i := 0; i <= 100; i++ {
		tm, err := GenerateDateTime("2022-01-01 00:00:00", "2022-08-21 17:08:26")
		if err != nil {

		}

		fmt.Println(tm)
	}

}
*/
