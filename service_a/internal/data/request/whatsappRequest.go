package request

type ContactReceive struct {
	CodeEnterprise string            `json:"code_enterprise"`
	To             string            `json:"to"`
	DriverName     string            `json:"driver_name"`
	Photo          string            `json:"photo,omitempty"`
	Services       []ServicesReceive `json:"services,omitempty"`
	Text           Text              `json:"text,omitempty"`
}

type ServicesReceive struct {
	ID                string `json:"id"`
	StartDate         string `json:"start_date"`
	StartTime         string `json:"start_time"`
	ServiceType       string `json:"service_type"`
	Description       string `json:"description"`
	DriverInformation string `json:"driver_information"`
}

type Language struct {
	Code string `json:"code"`
}
type Template struct {
	Name       string       `json:"name"`
	Language   Language     `json:"language"`
	Components []Components `json:"components"`
}
type Parameters struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
type Components struct {
	Type       string       `json:"type"`
	Parameters []Parameters `json:"parameters"`
}
type WhatsappRequest struct {
	MessagingProduct string   `json:"messaging_product"`
	RecipientType    string   `json:"recipient_type"`
	To               string   `json:"to" validate:"required"`
	Type             string   `json:"type"`
	Template         Template `json:"template"`
}
type WhatsappTextRequest struct {
	MessagingProduct string `json:"messaging_product"`
	RecipientType    string `json:"recipient_type"`
	To               string `json:"to" validate:"required"`
	Type             string `json:"type"`
	Text             Text   `json:"text"`
}

type Text struct {
	PreviewUrl bool   `json:"preview_url"`
	Body       string `json:"body"`
}

type ReceiveMessage struct {
	Field string `json:"field"`
	Value struct {
		MessagingProduct string `json:"messaging_product"`
		Metadata         struct {
			DisplayPhoneNumber string `json:"display_phone_number"`
			PhoneNumberID      string `json:"phone_number_id"`
		} `json:"metadata"`
		Contacts []struct {
			Profile struct {
				Name string `json:"name"`
			} `json:"profile"`
			WaID string `json:"wa_id"`
		} `json:"contacts"`
		Messages []struct {
			From      string `json:"from"`
			ID        string `json:"id"`
			Timestamp string `json:"timestamp"`
			Type      string `json:"type"`
			Text      struct {
				Body string `json:"body"`
			} `json:"text"`
		} `json:"messages"`
	} `json:"value"`
}

type WhatsappWebhook struct {
	Object string `json:"object"`
	Entry  []struct {
		ID      string `json:"id"`
		Changes []struct {
			Field string `json:"field"`
			Value struct {
				MessagingProduct string `json:"messaging_product"`
				Metadata         struct {
					DisplayPhoneNumber string `json:"display_phone_number"`
					PhoneNumberID      string `json:"phone_number_id"`
				} `json:"metadata"`
				Contacts []struct {
					Profile struct {
						Name string `json:"name"`
					} `json:"profile"`
					WaID string `json:"wa_id"`
				} `json:"contacts"`
				Statuses []struct {
					ID          string `json:"id"`
					Status      string `json:"status"`
					Timestamp   string `json:"timestamp"`
					RecipientID string `json:"recipient_id"`
					Errors      []struct {
						Code      int    `json:"code"`
						Title     string `json:"title"`
						Message   string `json:"message"`
						ErrorData struct {
							Details string `json:"details"`
						}
						Href string `json:"href"`
					} `json:"errors"`
					Conversation struct {
						ID                  string `json:"id"`
						ExpirationTimestamp string `json:"expiration_timestamp"`
						Origin              struct {
							Type string `json:"type"`
						} `json:"origin"`
					} `json:"conversation"`
					Pricing struct {
						Billable     string `json:"billable"`
						PricingModel string `json:"pricing_model"`
						Category     string `json:"category"`
					} `json:"pricing"`
				} `json:"statuses"`
				Messages []struct {
					From      string `json:"from"`
					ID        string `json:"id"`
					Timestamp string `json:"timestamp"`
					Type      string `json:"type"`
					Reaction  struct {
						MessageId string `json:"message_id"`
						Emoji     string `json:"emoji"`
					} `json:"reaction"`
					Text struct {
						Body string `json:"body"`
					} `json:"text"`
				} `json:"messages"`
			} `json:"value"`
		} `json:"changes"`
	} `json:"entry"`
}
