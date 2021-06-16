package payment

import (
	"testing"
)

func TestBuilderURL(t *testing.T) {
	in := buildRedirectURL("lel", "lel", 11, 500, "KEK", true)
	out := "https://auth.robokassa.ru/Merchant/Index.aspx?Desc=KEK&InvId=11&IsTest=1&MrchLogin=lel&OutSum=500&SignatureValue=e8c1c7bcacfa991b8612f2759804abd9"
	if in != out {
		t.Errorf("Wrong url. Expected: %s, Current: %s ", in, out)
		return
	}
}

func TestRobokassa_VerifyResult(t *testing.T) {
	vr1 := verifyResult("password", 666, "30.000000", "096684ec7388dcfe48ba1a3b5d2c6565")
	vr2 := verifyResult("password", 666, "32.000000", "096684ec7388dcfe48ba1a3b5d2c6565")

	if vr1 != true {
		t.Errorf("Wrong result. Expected: %t, Current: %t ", vr1, true)
		return
	}
	if vr2 != false {
		t.Errorf("Wrong result. Expected: %t, Current: %t ", vr2, false)
		return
	}

}

func TestRobokassa_VerifyRequest(t *testing.T) {
	var result = &RobokassaResult{"30", "666", "6214d840484820c470adfa20008d4507"}
	vrR1 := verifyRequest("password", result)
	vrR2 := verifyRequest("test", result)
	if vrR1 != true {
		t.Errorf("Wrong result. Expected: %t, Current: %t ", vrR1, true)
		return
	}
	if vrR2 != false {
		t.Errorf("Wrong result. Expected: %t, Current: %t ", vrR2, false)
		return
	}
}

func TestRobokassa_NewRobokassa(t *testing.T) {
	c := NewRobokassa("login", "pwd1", "password")
	in := c.URL(110, 2000, "description", true)
	out := "https://auth.robokassa.ru/Merchant/Index.aspx?Desc=description&InvId=110&IsTest=1&MrchLogin=login&OutSum=2000&SignatureValue=1364f38f54e76a0affe62974bfdbde85"
	if in != out {
		t.Errorf("Wrong url in TestRobokassa_NewRobokassa. Expected: %s, Current: %s ", in, out)
		return
	}
}

func TestRobokassa_CheckResult(t *testing.T) {
	c := NewRobokassa("login", "pwd1", "password")
	var result = &RobokassaResult{"1200", "666", "3a3869287aaa475dda04d93280705839"}
	checkR := c.CheckResult(result)

	if checkR != true {
		t.Errorf("Wrong test in TestRobokassa_CheckResult. Expected: %t, Current: %t ", checkR, true)
	}
}

func TestRobokassa_CheckSuccess(t *testing.T) {
	c := NewRobokassa("login", "pwd1", "password")
	var result = &RobokassaResult{"1200", "666", "1bc5c075a999e194e6d46ba351e4c11e"}
	checkS := c.CheckSuccess(result)

	if checkS != true {
		t.Errorf("Wrong test in TestRobokassa_CheckSuccess. Expected: %t, Current: %t ", checkS, true)
	}
}

func TestRobokassa_CheckSuccess_NotNumber(t *testing.T) {
	c := NewRobokassa("login", "pwd1", "password")
	var result = &RobokassaResult{"1200", "8666", "1bc5c075a999e194e6d46ba351e4c11e"}
	checkS := c.CheckSuccess(result)

	if checkS != false {
		t.Errorf("Wrong test in TestRobokassa_CheckSuccess. Expected: %t, Current: %t ", checkS, false)
	}

}

func TestRobokassa_CheckSuccess_OutSum(t *testing.T) {
	c := NewRobokassa("login", "pwd1", "password")
	var result = &RobokassaResult{"12900", "666", "SignatureValue"}

	checkOS := c.CheckSuccess(result)

	if checkOS != false {
		t.Errorf("Wrong test in TestRobokassa_CheckSuccess. Expected: %t, Current: %t ", checkOS, false)
	}
}
