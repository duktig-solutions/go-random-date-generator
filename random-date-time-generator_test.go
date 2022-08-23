package randomDataTime

import "testing"
import "fmt"

func TestGenerateDate(t *testing.T) {

	t.Parallel()

	for i := 1; i <= 100; i++ {
		result, err := GenerateDate("2020-01-01", "2022-01-01")
		if err != nil {
			t.Errorf("Method %s returned error - %s", "GenerateDate", err.Error())
		} else {
			fmt.Println(fmt.Sprintf("Generated Date: %s", result))
		}
	}

}

func TestGenerateTime(t *testing.T) {

	t.Parallel()

	for i := 1; i <= 100; i++ {
		result, err := GenerateTime(13, 19)

		if err != nil {
			t.Errorf("Method %s returned error - %s", "GenerateTime", err.Error())
		} else {
			fmt.Println(fmt.Sprintf("Generated Time: %s", result))
		}
	}

}

func TestGenerateDateTime(t *testing.T) {

	t.Parallel()

	for i := 1; i <= 100; i++ {
		result, err := GenerateDateTime("2020-01-01 10:56:54", "2022-01-01 23:18:40")

		if err != nil {
			t.Errorf("Method %s returned error - %s", "GenerateDateTime", err.Error())
		} else {
			fmt.Println(fmt.Sprintf("Generated Date/Time: %s", result))
		}
	}

}

func TestGenerateDOB(t *testing.T) {

	t.Parallel()

	for i := 1; i <= 100; i++ {
		result, err := GenerateDOB(16, 53)

		if err != nil {
			t.Errorf("Method %s returned error - %s", "GenerateDOB", err.Error())
		} else {
			fmt.Println(fmt.Sprintf("Genrated DOB: %s", result))
		}
	}
}