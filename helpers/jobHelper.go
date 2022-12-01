package helpers

import (
	"fmt"
    "math/rand"
    "strconv"
	"gopkg.in/robfig/cron.v2"
    "github.com/go-resty/resty/v2"
)

var limitStart 	= 1
var limitEnd 	= 20
var jobSeconds 	= 15
type IntRange struct {
	min, max int
}
type IntWeather struct{
	water int `json:"water"` 
	wind int `json:"wind"`
}

func RunCronJobs_Weather() {
	s := cron.New()
	client := resty.New()

	r := rand.New(rand.NewSource(55))
	ir := IntRange{limitStart, limitEnd}
	var vals IntWeather
	var v_err any 
	var v_resp any
	var v_body map[string]interface{}
	fmt.Println("Helo, there. \nThe value will update every "+strconv.Itoa(jobSeconds)+" seconds. \nStart from now..\n\n")

	s.AddFunc("@every "+strconv.Itoa(jobSeconds)+"s", func() {
		vals = IntWeather{ir.NextRandom(r), ir.NextRandom(r)}
		v_body = map[string]interface{}{
			"water": vals.water,
			"wind": vals.wind,
		}
		_,err := client.R().
			SetBody(v_body).
			SetResult(&v_resp).
			SetError(&v_err).
			Post("http://localhost:8080/weather")
		if err != nil {
			fmt.Println("________FAILED_to_INSERT___________________")
		}else{
			fmt.Println("________SUCCESS_to_INSERT;_NOW_EXAMINE:____")
			fmt.Printf("\n%+v", vals)
			fmt.Println(vals.ExamineData())
		}
	})
	s.Start()
}

func (ir *IntRange) NextRandom(r *rand.Rand) int {
	return r.Intn(ir.max-ir.min+1) + ir.min
}

func (vals *IntWeather) ExamineData() string {
	str := "\nstatus water: "
	switch true {
		case vals.water <= 5:
			str += "AMAN"
		break;
		case vals.water > 5 && vals.water < 9:
			str += "SIAGA"
		break;
		default:
			str += "BAHAYA"
		break;

	}
	str += "\nstatus wind: "
	switch true {
		case vals.wind <= 6:
			str += "AMAN"
		break;
		case vals.wind > 6 && vals.wind < 16:
			str += "SIAGA"
		break;
		default:
			str += "BAHAYA"
		break;

	}
	return str
}

   