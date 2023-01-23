package api

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"qkdc-service/data"
	"qkdc-service/utils"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var hashAlgorithmId = []byte{0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x11}

func ListenAndServe(manager *data.KeyManager, APIPort int, logRequests bool) {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		for {
			msgType, body, err := conn.ReadMessage() // msgType https://www.rfc-editor.org/rfc/rfc6455.html#section-11.8
			if err != nil {
				return
			}
			seq, err := DecodeDERSequence(body)
			if err != nil {
				log.Println(err)
			}
			stateId := seq[0].AsInt()
			switch stateId {
			case 0x01: // reserveKeyAndGetHalf
				if logRequests {
					log.Println("reserveKeyAndGetKeyHalf request: ", seq.ToString())
				}
				if len(seq) != 3 {
					log.Println("sequence of length 3 was expected")
					continue
				}
				_, callId := seq[1].AsInt(), seq[2].AsInt()
				keyId, thisHalf, otherHash, err := manager.ReserveKeyAndGetHalf()
				errCode := 0
				if err != nil {
					errCode = 1
				}
				res := DERSequence{}
				res = append(res, CreateIntSeqElement(0xff)) // reserveKeyAndGetHalf result
				res = append(res, CreateIntSeqElement(callId))
				res = append(res, CreateIntSeqElement(errCode))
				res = append(res, CreateArrSeqElement(keyId))
				res = append(res, CreateArrSeqElement(thisHalf))
				res = append(res, CreateArrSeqElement(otherHash))
				res = append(res, CreateObjSeqElement(hashAlgorithmId))
				if logRequests {
					log.Println("reserveKeyAndGetKeyHalf response: ", utils.BytesToHexOctets(res.ToByteArray()))
				}
				err = conn.WriteMessage(msgType, res.ToByteArray())
			case 0x02: // getKeyHalf
				if logRequests {
					log.Println("getKeyHalf request: ", seq.ToString())
				}
				if len(seq) != 4 {
					log.Println("sequence of length 4 was expected")
					continue
				}
				_, keyId, callId := seq[1].AsInt(), seq[2].AsBytes(), seq[3].AsInt()
				thisHalf, otherHash, err := manager.GetKeyThisHalfOtherHash(keyId)
				errCode := 0
				if err != nil {
					log.Println(err)
					errCode = 1
				}
				res := DERSequence{}
				res = append(res, CreateIntSeqElement(0xfe)) // getKeyHalf result
				res = append(res, CreateIntSeqElement(callId))
				res = append(res, CreateIntSeqElement(errCode))
				res = append(res, CreateArrSeqElement(thisHalf))
				res = append(res, CreateArrSeqElement(otherHash))
				res = append(res, CreateObjSeqElement(hashAlgorithmId))
				if logRequests {
					log.Println("getKeyHalf response: ", utils.BytesToHexOctets(res.ToByteArray()))
				}
				err = conn.WriteMessage(msgType, res.ToByteArray())
			case 0x03: // getState
				if logRequests {
					log.Println("getState request: ", seq.ToString())
				}
				if len(seq) != 0 {
					log.Println("sequence of length 0 was expected")
					continue
				}
				res := DERSequence{}
				// i need to get state from manager
				// i also need to get first keys of each parity
				err = conn.WriteMessage(msgType, res.ToByteArray())
			case 0x04: // setState
				if logRequests {
					log.Println("setState request: ", seq.ToString())
				}
				if len(seq) != 0 {
					log.Println("sequence of length 0 was expected")
					continue
				}
				res := DERSequence{}
				err = conn.WriteMessage(msgType, res.ToByteArray())
			}

			if err != nil {
				log.Println(err)
				return
			}
		}
	})

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", APIPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("server listening at %v", lis.Addr())
	log.Panic(http.Serve(lis, http.DefaultServeMux))
}
