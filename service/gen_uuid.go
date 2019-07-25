package service

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/chilts/sid"
	guuid "github.com/google/uuid"
	"github.com/kjk/betterguid"
	"github.com/lithammer/shortuuid"
	"github.com/oklog/ulid"
	"github.com/rs/xid"
	"github.com/satori/go.uuid"
	"github.com/segmentio/ksuid"
	"github.com/sony/sonyflake"
	"github.com/speps/go-hashids"
)

/*
	Link tham khảo: https://blog.kowalczyk.info/article/JyRZ/generating-good-unique-ids-in-go.html
 */


func genShortUUID() {
	id := shortuuid.New()
	fmt.Printf("github.com/lithammer/shortuuid: %s\n", id)
}

func genUUID() {
	id := guuid.New()
	fmt.Printf("github.com/google/uuid:         %s\n", id.String())
}

func genXid() {
	id := xid.New()
	fmt.Printf("github.com/rs/xid:              %s\n", id.String())
}

func genKsuid() {
	id := ksuid.New()
	fmt.Printf("github.com/segmentio/ksuid:     %s\n", id.String())
}

func genBetterGUID() {
	id := betterguid.New()
	fmt.Printf("github.com/kjk/betterguid:      %s\n", id)
}

func genUlid() {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	fmt.Printf("github.com/oklog/ulid:          %s\n", id.String())
}

func genSonyflake() {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		log.Fatalf("flake.NextID() failed with %s\n", err)
	}
	// Note: this is base16, could shorten by encoding as base62 string
	fmt.Printf("github.com/sony/sonyflake:      %x\n", id)
}

func genSid() {
	id := sid.Id()
	fmt.Printf("github.com/chilts/sid:          %s\n", id)
}

func genUUIDv4() {
	id := uuid.NewV4()

	fmt.Printf("github.com/satori/go.uuid:      %s\n", id)
}

func HashIds(){

	startTime := time.Now().UnixNano()

	hd := hashids.NewData()
	hd.Salt = "this is my salt"
	hd.MinLength = 30
	h, _ := hashids.NewWithData(hd)

	number := 1563878884

	e, _ := h.Encode([]int{number})
	fmt.Println("Chuỗi hash", e)

	d, _ := h.DecodeWithError(e)

	fmt.Println("Chuỗi decode", d)

	endTime := time.Now().UnixNano()

	fmt.Println("endTime...", endTime)

	fmt.Println("tốn hết ", endTime-startTime)
}
