package thousandeyes

import (
	"encoding/json"
	"fmt"
)

// RTP Stream, labeled "voice"

// RTPStream - RTPStream trace test
type RTPStream struct {
	// Common test fields
	AlertsEnabled      *bool                `json:"alertsEnabled,omitempty" te:"int-bool"`
	AlertRules         *[]AlertRule         `json:"alertRules,omitempty"`
	APILinks           *[]APILink           `json:"apiLinks,omitempty"`
	CreatedBy          *string              `json:"createdBy,omitempty"`
	CreatedDate        *string              `json:"createdDate,omitempty"`
	Description        *string              `json:"description,omitempty"`
	Enabled            *bool                `json:"enabled,omitempty" te:"int-bool"`
	Groups             *[]GroupLabel        `json:"groups,omitempty"`
	ModifiedBy         *string              `json:"modifiedBy,omitempty"`
	ModifiedDate       *string              `json:"modifiedDate,omitempty"`
	SavedEvent         *bool                `json:"savedEvent,omitempty" te:"int-bool"`
	SharedWithAccounts *[]SharedWithAccount `json:"sharedWithAccounts,omitempty"`
	TestID             *int64               `json:"testId,omitempty"`
	TestName           *string              `json:"testName,omitempty"`
	Type               *string              `json:"type,omitempty"`
	LiveShare          *bool                `json:"liveShare,omitempty" te:"int-bool"`

	// Fields unique to this test
	Agents          *[]Agent      `json:"agents,omitempty"`
	BGPMeasurements *bool         `json:"bgpMeasurements,omitempty" te:"int-bool"`
	BGPMonitors     *[]BGPMonitor `json:"bgpMonitors,omitempty"`
	Codec           *string       `json:"codec,omitempty"`
	CodecID         *int64        `json:"codecId,omitempty"`
	DSCP            *string       `json:"dscp,omitempty"`
	DSCPID          *int64        `json:"dscpId,omitempty"`
	Duration        *int          `json:"duration,omitempty"`
	Interval        *int          `json:"interval,omitempty"`
	JitterBuffer    *int          `json:"jitterBuffer,omitempty"`
	MTUMeasurements *bool         `json:"mtuMeasurements,omitempty" te:"int-bool"`
	NumPathTraces   *int          `json:"numPathTraces,omitempty"`
	TargetAgentID   *int64        `json:"targetAgentId,omitempty"`
	UsePublicBGP    *bool         `json:"usePublicBgp,omitempty" te:"int-bool"`
	// server field is present in response, but we should not track it.
	//Server          *string       `json:"server,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface. It ensures
// that ThousandEyes int fields that only use the values 0 or 1 are
// treated as booleans.
func (t RTPStream) MarshalJSON() ([]byte, error) {
	type aliasTest RTPStream

	data, err := json.Marshal((aliasTest)(t))
	if err != nil {
		return nil, err
	}

	return jsonBoolToInt(&t, data)
}

// UnmarshalJSON implements the json.Unmarshaler interface. It ensures
// that ThousandEyes int fields that only use the values 0 or 1 are
// treated as booleans.
func (t *RTPStream) UnmarshalJSON(data []byte) error {
	type aliasTest RTPStream
	test := (*aliasTest)(t)

	data, err := jsonIntToBool(t, data)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &test)
}

// AddAgent - Add agent to voice call  test
func (t *RTPStream) AddAgent(id int64) {
	agent := Agent{AgentID: Int64(id)}
	*t.Agents = append(*t.Agents, agent)
}

// GetRTPStream - get voice call test
func (c *Client) GetRTPStream(id int64) (*RTPStream, error) {
	resp, err := c.get(fmt.Sprintf("/tests/%d", id))
	if err != nil {
		return &RTPStream{}, err
	}
	var target map[string][]RTPStream
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

//CreateRTPStream - Create voice call test
func (c Client) CreateRTPStream(t RTPStream) (*RTPStream, error) {
	resp, err := c.post("/tests/voice/new", t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 201 {
		return &t, fmt.Errorf("failed to create voice test, response code %d", resp.StatusCode)
	}
	var target map[string][]RTPStream
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}

//DeleteRTPStream - delete voice call test
func (c *Client) DeleteRTPStream(id int64) error {
	resp, err := c.post(fmt.Sprintf("/tests/voice/%d/delete", id), nil, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to delete voice test, response code %d", resp.StatusCode)
	}
	return nil
}

//UpdateRTPStream - update voice call test
func (c *Client) UpdateRTPStream(id int64, t RTPStream) (*RTPStream, error) {
	resp, err := c.post(fmt.Sprintf("/tests/voice/%d/update", id), t, nil)
	if err != nil {
		return &t, err
	}
	if resp.StatusCode != 200 {
		return &t, fmt.Errorf("failed to update test, response code %d", resp.StatusCode)
	}
	var target map[string][]RTPStream
	if dErr := c.decodeJSON(resp, &target); dErr != nil {
		return nil, fmt.Errorf("Could not decode JSON response: %v", dErr)
	}
	return &target["test"][0], nil
}
