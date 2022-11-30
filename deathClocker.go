package main
import (
	"math/rand"
	"time"
	"fmt"
	"net/http"
	"strings"
	"strconv"
)


func genDeathdate(x int) time.Time {
	rand.Seed(time.Now().UnixNano())
	now:=time.Now()
	yearRand:=rand.Intn(50)
	monthRand:=rand.Intn(12)
	dayRand:=rand.Intn(30)
	death:=now.AddDate(yearRand+x,monthRand,dayRand)
	return death
	
}

func deathClocker (w http.ResponseWriter, r *http.Request) {
age,_ := strconv.Atoi (strings.TrimPrefix(r.URL.Path, "/age/"))
	deathDate:= genDeathdate(age) 
	z:=deathDate.Format("02-Jan-2006")
	w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Fprintf(w, "Поздравляю, ты умрешь только %v!", z)
}