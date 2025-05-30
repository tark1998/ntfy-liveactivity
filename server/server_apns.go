package server

import (
	"log"
	"fmt"
	//"time"

	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/token"
)

type apnsClient struct {
	client *apns2.Client
}
func newapnsClient() *apnsClient {
	authKey, err := token.AuthKeyFromFile("/etc/ntfy/AuthKey_XXX.p8")
	if err != nil {
		log.Fatal("token error:", err)
	}
	token := &token.Token{
		AuthKey: authKey,
		// KeyID from developer account (Certificates, Identifiers & Profiles -> Keys)
		KeyID:   "ABC123DEFG",
		// TeamID from developer account (View Account -> Membership)
		TeamID:  "DEF123GHIJ",
	}
	client := apns2.NewTokenClient(token)
	return &apnsClient{
		client: client,
	}
}
func (c *apnsClient) Send(v *visitor, m *message, P2STokens P2SToken, P2UTokens P2UToken) error {
	notification := &apns2.Notification{}
	notification.Topic = "com.tark1998.trp-activity.push-type.liveactivity"
	notification.DeviceToken = P2UTokens.Token
	notification.Priority = 10
	notification.PushType = "liveactivity"
	switch m.Activity {
	case 1:
		notification.DeviceToken = P2STokens.Token 
		notification.Payload = []byte(m.Message)/*fmt.Sprintf(`{
			"aps": {
				"timestamp": %d,
				"event": "start",
				"content-state": {
					"got": %s
				},
				"attributes-type": "MywidgetAttributes",
				"attributes": {
					"remain": %s,
					"name": "Apple"
				},
				"alert": {
					"title": "Hello",
					"body": "World",
					"sound": "chime.aiff"
				}
			}
		}`,time.Now().Unix(), m.Message, m.Title))*/

	case 2:
		notification.Payload = []byte(m.Message)/*fmt.Sprintf(`{
			"aps": {
				"timestamp": %d,
				"event": "update",
				"content-state": {
					"got": %s
				}
			}
		}`,time.Now().Unix(), m.Message))*/

	case 3:
		notification.Payload = []byte(m.Message)/*fmt.Sprintf(`{
			"aps": {
				"timestamp": %d,
				"event": "end",
				"content-state": {
					"got": %s
				},
				"dismissal-date": %d,
				"alert": {
					"title": "Hello",
					"body": "Update World",
					"sound": "chime.aiff"
				}
			}
		}`,time.Now().Unix(), m.Message, time.Now().Unix()+10))*/
	}

	res, err := c.client.Push(notification)

	if err != nil {
		log.Fatal("Error:", err)
		return err
	}

	ev := logvm(v, m).Tag(tagapns)
	if ev.IsTrace() {
		//ev.Field("apns_message", util.MaybeMarshalJSON(fbm)).Trace("Apns message")
		fmt.Printf("%v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)
	}
	return err
}
