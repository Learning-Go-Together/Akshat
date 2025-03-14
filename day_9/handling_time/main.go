package main

import (
	"fmt"
	"time"
)

/*
Basic full date  "Mon Jan 2 15:04:05 MST 2006" gives "Wed Feb 25 11:06:39 PST 2015"
Basic short date "2006/01/02" gives "2015/02/25"
AM/PM            "3PM==3pm==15h" gives "11AM==11am==11h"
No fraction      "Mon Jan _2 15:04:05 MST 2006" gives "Wed Feb 25 11:06:39 PST 2015"
0s for fraction  "15:04:05.00000" gives "11:06:39.12340"
9s for fraction  "15:04:05.99999999" gives "11:06:39.1234"
*/
func main() {
	present_time := time.Now()
	fmt.Printf("Present time (default format): %v\n", present_time)
	fmt.Printf("Present time (dd-mm-yyyy format): %v\n", present_time.Format("01-02-2006"))
	fmt.Printf("Present time (dd-mm-yy format): %v\n", present_time.Format("01-02-06"))
	fmt.Printf("Present time (date,day month year format): %v\n", present_time.Format("02, Monday January 2006"))

	//we can also create any date

	createdDate := time.Date(2003, time.November, 7, 00, 0, 0, 0, time.UTC).Format("02, Monday January 2006")
	fmt.Println("Created Date:", createdDate)
}
