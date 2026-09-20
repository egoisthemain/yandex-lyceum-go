package main

import (
	"bufio"
	"errors"
	"strings"
	"time"
)

type Ticket struct {
	Ticket string
	User   string
	Status string
	Date   time.Time
}

const (
	WrongTicketStringError = "wrong ticket string"
	InvalidStatus          = "wrong status"
	InvalidDate            = "wrong date"
	InvalideName           = "wrong name"
)

func parseTicketString(s string) (Ticket, error) {
	validStatus := []string{"Готово", "В работе", "Не будет сделано"}

	ticket := Ticket{}
	parts := strings.Split(s, "_")
	if len(parts) != 4 {
		return ticket, errors.New(WrongTicketStringError)
	}

	ticketNum, user, status, date := parts[0], parts[1], parts[2], parts[3]

	if ticketNum[:6] != "TICKET" && ticketNum[6] != '-' {
		return ticket, errors.New(WrongTicketStringError)
	}

	for i, valStatus := range validStatus {
		if status == valStatus {
			break
		}
		if i == len(validStatus)-1 {
			return ticket, errors.New(InvalidStatus)
		}
	}

	timeStamp, err := time.Parse("2006-01-02", date)
	if err != nil {
		return ticket, errors.New(InvalidDate)
	}

	ticket.Ticket = ticketNum
	ticket.User = user
	ticket.Status = status
	ticket.Date = timeStamp

	return ticket, nil
}

func GetTasks(text string, user *string, status *string) []Ticket {
	targetTickets := []Ticket{}

	scanner := bufio.NewScanner(strings.NewReader(text))

	for scanner.Scan() {

		currString := scanner.Text()
		currString = strings.TrimSpace(currString)
		if currString == "" || len(currString) < 1 {
			continue
		}

		newTicket, err := parseTicketString(currString)
		if err != nil {
			continue
		}

		if user == nil && status == nil {
			targetTickets = append(targetTickets, newTicket)
		} else if user != nil && status != nil {
			if newTicket.User == *user && newTicket.Status == *status {
				targetTickets = append(targetTickets, newTicket)
			}
		} else if user != nil && status == nil {
			if newTicket.User == *user {
				targetTickets = append(targetTickets, newTicket)
			}
		} else if user == nil && status != nil {
			if newTicket.Status == *status {
				targetTickets = append(targetTickets, newTicket)
			}
		}
	}

	return targetTickets
}
