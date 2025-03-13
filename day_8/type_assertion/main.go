package main

func getExpenseReport(e expense) (string, float64) {
	//
	// switch t := e.(type) {
	// case email:
	// 	return t.toAddress, t.cost()
	// case sms:
	// 	return t.toPhoneNumber, t.cost()
	// default:
	// 	return "", t.cost()
	// }

	var t1, ok1 = e.(email)
	if ok1 {
		return t1.toAddress, t1.cost()
	}
	var t2, ok2 = e.(sms)
	if ok2 {
		return t2.toPhoneNumber, t2.cost()
	}
	var t3, ok3 = e.(invalid)
	if ok3 {
		return "", t3.cost()
	}
	return "", 0.0

}

// don't touch below this line

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}
