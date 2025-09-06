package sseclient

import (
	"io"
	"net/http"
	"testing"
	"time"
)

var testServerStarted = false
var testClient *client

func createClient() *client {
	if testClient != nil {
		return testClient
	}

	testClient := NewClient("http://127.0.0.1:8080/stream")

	return testClient
}

func startLocalServer(t *testing.T) {
	if testServerStarted {
		return
	}

	errCh := make(chan error, 1)

	go func() {
		http.HandleFunc("/stream", func(w http.ResponseWriter, _ *http.Request) {
			io.WriteString(w, ":ok\r\n")
			io.WriteString(w, "\r\n")

			io.WriteString(w, "event: message\r\n")
			io.WriteString(w, `id: [{"topic":"eqiad.mediawiki.page-create","partition":0,"timestamp":1757129333323},{"topic":"codfw.mediawiki.page-create","partition":0,"offset":-1}]`+"\r\n")
			io.WriteString(w, `data: {"$schema":"/mediawiki/revision/create/2.0.0","meta":{"uri":"https://commons.wikimedia.org/wiki/File:Keith%27s_Magazine_on_Home_Building,_1915;_Keith%27s_Home-Builder_-_DPLA_-_baf2c6b254fbda2f3d9d0b3059945d19_(page_617).jpg","request_id":"61ac3485-b77b-4860-adc3-664d75ebb68e","id":"b6a5fea2-91be-48db-b78c-6abbb72bdab4","domain":"commons.wikimedia.org","stream":"mediawiki.page-create","dt":"2025-09-06T03:28:53.322Z","topic":"eqiad.mediawiki.page-create","partition":0,"offset":329279418},"database":"commonswiki","page_id":174482800,"page_title":"File:Keith's_Magazine_on_Home_Building,_1915;_Keith's_Home-Builder_-_DPLA_-_baf2c6b254fbda2f3d9d0b3059945d19_(page_617).jpg","page_namespace":6,"rev_id":1081516708,"rev_timestamp":"2025-09-06T03:28:51Z","rev_sha1":"jww5bk9r8was0okyhe6ig9px7w98kby","rev_minor_edit":false,"rev_len":1172,"rev_content_model":"wikitext","rev_content_format":"text/x-wiki","performer":{"user_text":"DPLA bot","user_groups":["bot","filemover","*","user","autoconfirmed"],"user_is_bot":true,"user_id":8609812,"user_registration_dt":"2019-12-18T13:57:27Z","user_edit_count":26531708},"page_is_redirect":false,"comment":"Uploading DPLA ID \"[[dpla:baf2c6b254fbda2f3d9d0b3059945d19|baf2c6b254fbda2f3d9d0b3059945d19]]\".","parsedcomment":"Uploading DPLA ID &quot;<a href=\"https://dp.la/item/baf2c6b254fbda2f3d9d0b3059945d19\" class=\"extiw\" title=\"dpla:baf2c6b254fbda2f3d9d0b3059945d19\">baf2c6b254fbda2f3d9d0b3059945d19</a>&quot;.","dt":"2025-09-06T03:28:51Z","rev_slots":{"main":{"rev_slot_content_model":"wikitext","rev_slot_sha1":"jww5bk9r8was0okyhe6ig9px7w98kby","rev_slot_size":1172,"rev_slot_origin_rev_id":1081516708}}}`+"\r\n")
			io.WriteString(w, "\r\n")

			io.WriteString(w, "event: message\r\n")
			io.WriteString(w, `id: [{"topic":"eqiad.mediawiki.page-create","partition":0,"timestamp":1757129333851},{"topic":"codfw.mediawiki.page-create","partition":0,"offset":-1}]`+"\r\n")
			io.WriteString(w, `data: {"$schema":"/mediawiki/revision/create/2.0.0","meta":{"uri":"https://commons.wikimedia.org/wiki/File:Aufn._11%2B12%2B14_-_LABW_-_Staatsarchiv_Ludwigsburg_PL_734_FM1_131.jpg","request_id":"a423e402-ae63-4264-b5f7-3ab26d729b72","id":"11695ed6-0c8d-4491-ac06-0d7e328fe6bd","domain":"commons.wikimedia.org","stream":"mediawiki.page-create","dt":"2025-09-06T03:28:53.850Z","topic":"eqiad.mediawiki.page-create","partition":0,"offset":329279419},"database":"commonswiki","page_id":174482801,"page_title":"File:Aufn._11+12+14_-_LABW_-_Staatsarchiv_Ludwigsburg_PL_734_FM1_131.jpg","page_namespace":6,"rev_id":1081516711,"rev_timestamp":"2025-09-06T03:28:52Z","rev_sha1":"fx9t9e9ecnxo0gsphfmep718ybx6cbi","rev_minor_edit":false,"rev_len":1470,"rev_content_model":"wikitext","rev_content_format":"text/x-wiki","performer":{"user_text":"CuratorBot","user_groups":["bot","ipblock-exempt","*","user","autoconfirmed"],"user_is_bot":true,"user_id":13158996,"user_registration_dt":"2024-07-31T10:17:37Z","user_edit_count":2027243},"page_is_redirect":false,"comment":"upload Landesarchiv Baden-Württemberg archives Staatsarchiv Ludwigsburg Findbuch PL 734 ([[:toollabs:editgroups-commons/b/OR/3673fa23ba1|details]])","parsedcomment":"upload Landesarchiv Baden-Württemberg archives Staatsarchiv Ludwigsburg Findbuch PL 734 (<a href=\"https://iw.toolforge.org/editgroups-commons/b/OR/3673fa23ba1\" class=\"extiw\" title=\"toollabs:editgroups-commons/b/OR/3673fa23ba1\">details</a>)","dt":"2025-09-06T03:28:52Z","rev_slots":{"main":{"rev_slot_content_model":"wikitext","rev_slot_sha1":"fx9t9e9ecnxo0gsphfmep718ybx6cbi","rev_slot_size":1470,"rev_slot_origin_rev_id":1081516711}}}`+"\r\n")
			io.WriteString(w, "\r\n")

			io.WriteString(w, "event: message\r\n")
			io.WriteString(w, `id: [{"topic":"eqiad.mediawiki.page-create","partition":0,"timestamp":1757129336258},{"topic":"codfw.mediawiki.page-create","partition":0,"offset":-1}]`+"\r\n")
			io.WriteString(w, `data: {"$schema":"/mediawiki/revision/create/2.0.0","meta":{"uri":"https://commons.wikimedia.org/wiki/File:Keith%27s_Magazine_on_Home_Building,_1915;_Keith%27s_Home-Builder_-_DPLA_-_baf2c6b254fbda2f3d9d0b3059945d19_(page_618).jpg","request_id":"da0d3310-a979-4f69-ad9c-8d8fab7d58f3","id":"bc0ee844-8e5e-4f1c-a8fa-fe46997ba56c","domain":"commons.wikimedia.org","stream":"mediawiki.page-create","dt":"2025-09-06T03:28:56.257Z","topic":"eqiad.mediawiki.page-create","partition":0,"offset":329279420},"database":"commonswiki","page_id":174482802,"page_title":"File:Keith's_Magazine_on_Home_Building,_1915;_Keith's_Home-Builder_-_DPLA_-_baf2c6b254fbda2f3d9d0b3059945d19_(page_618).jpg","page_namespace":6,"rev_id":1081516723,"rev_timestamp":"2025-09-06T03:28:53Z","rev_sha1":"jww5bk9r8was0okyhe6ig9px7w98kby","rev_minor_edit":false,"rev_len":1172,"rev_content_model":"wikitext","rev_content_format":"text/x-wiki","performer":{"user_text":"DPLA bot","user_groups":["bot","filemover","*","user","autoconfirmed"],"user_is_bot":true,"user_id":8609812,"user_registration_dt":"2019-12-18T13:57:27Z","user_edit_count":26531709},"page_is_redirect":false,"comment":"Uploading DPLA ID \"[[dpla:baf2c6b254fbda2f3d9d0b3059945d19|baf2c6b254fbda2f3d9d0b3059945d19]]\".","parsedcomment":"Uploading DPLA ID &quot;<a href=\"https://dp.la/item/baf2c6b254fbda2f3d9d0b3059945d19\" class=\"extiw\" title=\"dpla:baf2c6b254fbda2f3d9d0b3059945d19\">baf2c6b254fbda2f3d9d0b3059945d19</a>&quot;.","dt":"2025-09-06T03:28:53Z","rev_slots":{"main":{"rev_slot_content_model":"wikitext","rev_slot_sha1":"jww5bk9r8was0okyhe6ig9px7w98kby","rev_slot_size":1172,"rev_slot_origin_rev_id":1081516723}}}`+"\r\n")
			io.WriteString(w, "\r\n")

			io.WriteString(w, "event: message\r\n")
			io.WriteString(w, `id: [{"topic":"eqiad.mediawiki.page-create","partition":0,"timestamp":1757129337590},{"topic":"codfw.mediawiki.page-create","partition":0,"offset":-1}]`+"\r\n")
			io.WriteString(w, `data: {"$schema":"/mediawiki/revision/create/2.0.0","meta":{"uri":"https://commons.wikimedia.org/wiki/Commons:Deletion_requests/File:HipstamaticPhoto-628876182.689958.jpg","request_id":"ded1a29b-517c-43ee-b2c5-d51c4da6f6cc","id":"54044d3f-325a-465b-82d0-43202633f7e4","domain":"commons.wikimedia.org","stream":"mediawiki.page-create","dt":"2025-09-06T03:28:57.589Z","topic":"eqiad.mediawiki.page-create","partition":0,"offset":329279421},"database":"commonswiki","page_id":174482803,"page_title":"Commons:Deletion_requests/File:HipstamaticPhoto-628876182.689958.jpg","page_namespace":4,"rev_id":1081516735,"rev_timestamp":"2025-09-06T03:28:57Z","rev_sha1":"dpshkiafueenhu9zmmpg03q5c5urtq3","rev_minor_edit":false,"rev_len":204,"rev_content_model":"wikitext","rev_content_format":"text/x-wiki","performer":{"user_text":"191.126.129.224","user_groups":["*"],"user_is_bot":false},"page_is_redirect":false,"comment":"Starting deletion request","parsedcomment":"Starting deletion request","dt":"2025-09-06T03:28:57Z","rev_slots":{"main":{"rev_slot_content_model":"wikitext","rev_slot_sha1":"dpshkiafueenhu9zmmpg03q5c5urtq3","rev_slot_size":204,"rev_slot_origin_rev_id":1081516735}}}`+"\r\n")
			io.WriteString(w, "\r\n")
		})

		errCh <- http.ListenAndServe("127.0.0.1:8080", nil)
	}()

	select {
	case err := <-errCh:
		testServerStarted = false
		t.Fatal("startLocalServer: Failed starting the local server stream: ", err)
	case <-time.After(1 * time.Second):
		testServerStarted = true
		t.Log("startLocalServer: Started the local server stream")
	}
}

func TestClientEvents(t *testing.T) {
	startLocalServer(t)

	client := createClient()
	events, err := client.Events()
	if err != nil {
		t.Fatal("TestClientEvents: Failed listening to the local server stream: ", err)
		return
	}

	for _, event := range events {
		t.Log("TestClientEvents: ", event, client.status)
	}
}

func TestClientListen(t *testing.T) {
	startLocalServer(t)

	client := createClient()
	client.WithOnEvent(func(event *Event) {
		t.Log("TestClientListen: \r\n\r\n", "Event Name: ", event.Name, "\r\nEvent ID: ", event.Id, "\r\nEvent Data: ", event.Data, "\r\nClient Status: ", client.status)
	})

	err := client.Listen(false)
	if err != nil {
		t.Fatal("TestClientListen: Failed listening to the local server stream: ", err)
	}
}

func TestClientListenAsync(t *testing.T) {
	startLocalServer(t)

	client := createClient()
	client.WithOnEvent(func(event *Event) {
		t.Log("TestClientListenAsync: ", event, client.status)
	})

	err := client.Listen(true)
	if err != nil {
		t.Fatal("TestClientListenAsync: Failed listening to the local server stream: ", err)
	}

	// Pretend to do work.
	t.Log("TestClientListenAsync: Sleeping...")
	time.Sleep(5 * time.Second)
}

func TestRealStream(t *testing.T) {
	client := NewClient("https://stream.wikimedia.org/v2/stream/mediawiki.page-create")
	client.WithOnEvent(func(event *Event) {
		t.Log("TestRealStream: ", event, client.status)
	})

	err := client.Listen(false)
	if err != nil {
		t.Fatal("TestRealStream: Failed listening to the local server stream: ", err)
	}
}
