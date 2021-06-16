package database

import (
	"context"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestUserRepo_Create(t *testing.T) {
	userRepo := getUserRepo()
	user := getUser()
	userRepo.Create(user)
	testUser := getUser()
	if user.GetTelegramId() != testUser.GetTelegramId() {
		t.Errorf("Wrong TelegramId. Expected: %d, Current: %d", testUser.GetTelegramId(), user.GetTelegramId())
	}
	if user.GetIsPaid() != testUser.GetIsPaid() {
		t.Errorf("Wrong IsPaid Status. Expected: %t, Current: %t", testUser.GetIsPaid(), user.GetIsPaid())

	}
	if user.GetIsActive() != testUser.GetIsActive() {
		t.Errorf("Wrong IsActive Status. Expected: %t, Current: %t", testUser.GetIsActive(), user.GetIsActive())

	}
}

func TestUserRepo_MarkAsPaid(t *testing.T) {
	userRepo := getUserRepo()
	user := getUser()
	userRepo.Create(user)
	userRepo.MarkAsPaid(user, 4)

	testUser := userRepo.FindById(user.GetId())
	if testUser.GetIsPaid() != true {
		t.Errorf("Wrong IsPaid Status. Expected: true, Current: %t", testUser.GetIsPaid())
	}

	assertCompareDates(
		testUser.GetPaidUntil(),
		time.Now().AddDate(0, 1, 0),
		t,
	)
}

func TestUserRepo_MarkAsPaid_PaidProlongation(t *testing.T) {
	userRepo := getUserRepo()
	user := getUser()
	userLastPaidTime := time.Now().AddDate(0, 0, 10)
	user.PaidUntil = userLastPaidTime
	user.IsPaid = true
	userRepo.Create(user)
	userRepo.MarkAsPaid(user, 4)

	testUser := userRepo.FindById(user.GetId())
	assertCompareDates(
		testUser.GetPaidUntil(),
		userLastPaidTime.AddDate(0, 1, 0),
		t,
	)
}

func TestUserRepo_MarkAsPaid_PaidProlongation_OldPaidUntil(t *testing.T) {
	userRepo := getUserRepo()
	user := getUser()
	userLastPaidTime := time.Now().AddDate(0, 0, -10)
	user.PaidUntil = userLastPaidTime
	user.IsPaid = true
	userRepo.Create(user)
	userRepo.MarkAsPaid(user, 4)

	testUser := userRepo.FindById(user.GetId())
	assertCompareDates(
		testUser.GetPaidUntil(),
		time.Now().AddDate(0, 1, 0),
		t,
	)
}

func TestUserRepo_MarkAsPaid_PaidProlongation_IsNotPaid(t *testing.T) {
	userRepo := getUserRepo()
	user := getUser()
	userLastPaidTime := time.Now().AddDate(0, 0, 10)
	user.PaidUntil = userLastPaidTime
	userRepo.Create(user)
	userRepo.MarkAsPaid(user, 4)

	testUser := userRepo.FindById(user.GetId())
	assertCompareDates(
		testUser.GetPaidUntil(),
		time.Now().AddDate(0, 1, 0),
		t,
	)
}

func TestUserRepo_MarkAsPaid_PaidProlongation_IsNotPaid_OldPaidUntil(t *testing.T) {
	userRepo := getUserRepo()
	user := getUser()
	userLastPaidTime := time.Now().AddDate(0, 0, -10)
	user.PaidUntil = userLastPaidTime
	user.IsPaid = false
	userRepo.Create(user)
	userRepo.MarkAsPaid(user, 4)

	testUser := userRepo.FindById(user.GetId())
	assertCompareDates(
		testUser.GetPaidUntil(),
		time.Now().AddDate(0, 1, 0),
		t,
	)
}

func assertCompareDates(current time.Time, expected time.Time, t *testing.T) {
	currentYear, currentMonth, currentDay := current.Date()
	expectedYear, expectedMonth, expectedDay := expected.Date()
	if currentYear != expectedYear {
		t.Errorf("Wrong PaidUntil Year. Expected: %d, Current: %d", expectedYear, currentYear)
	}
	if currentMonth != expectedMonth {
		t.Errorf("Wrong PaidUntil Month. Expected: %s, Current: %s", expectedMonth, currentMonth)
	}
	if currentDay != expectedDay {
		t.Errorf("Wrong PaidUntil Day. Expected: %d, Current: %d", expectedDay, currentDay)
	}
}

func getUserRepo() *UserRepo {
	ctx := Boot(context.Background())
	db := ctx.Value("db").(*gorm.DB)
	return NewUserRepo(db)
}

func getUser() *User {
	return &User{
		TelegramId: 1251513,
		IsPaid:     false,
		IsActive:   true,
	}
}
