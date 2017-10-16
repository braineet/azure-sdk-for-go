package job

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"errors"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

// CompileMode enumerates the values for compile mode.
type CompileMode string

const (
	// Full specifies the full state for compile mode.
	Full CompileMode = "Full"
	// Semantic specifies the semantic state for compile mode.
	Semantic CompileMode = "Semantic"
	// SingleBox specifies the single box state for compile mode.
	SingleBox CompileMode = "SingleBox"
)

// ResourceType enumerates the values for resource type.
type ResourceType string

const (
	// StatisticsResource specifies the statistics resource state for resource type.
	StatisticsResource ResourceType = "StatisticsResource"
	// VertexResource specifies the vertex resource state for resource type.
	VertexResource ResourceType = "VertexResource"
)

// Result enumerates the values for result.
type Result string

const (
	// Cancelled specifies the cancelled state for result.
	Cancelled Result = "Cancelled"
	// Failed specifies the failed state for result.
	Failed Result = "Failed"
	// None specifies the none state for result.
	None Result = "None"
	// Succeeded specifies the succeeded state for result.
	Succeeded Result = "Succeeded"
)

// SeverityTypes enumerates the values for severity types.
type SeverityTypes string

const (
	// Error specifies the error state for severity types.
	Error SeverityTypes = "Error"
	// Warning specifies the warning state for severity types.
	Warning SeverityTypes = "Warning"
)

// State enumerates the values for state.
type State string

const (
	// StateAccepted specifies the state accepted state for state.
	StateAccepted State = "Accepted"
	// StateCompiling specifies the state compiling state for state.
	StateCompiling State = "Compiling"
	// StateEnded specifies the state ended state for state.
	StateEnded State = "Ended"
	// StateNew specifies the state new state for state.
	StateNew State = "New"
	// StatePaused specifies the state paused state for state.
	StatePaused State = "Paused"
	// StateQueued specifies the state queued state for state.
	StateQueued State = "Queued"
	// StateRunning specifies the state running state for state.
	StateRunning State = "Running"
	// StateScheduling specifies the state scheduling state for state.
	StateScheduling State = "Scheduling"
	// StateStarting specifies the state starting state for state.
	StateStarting State = "Starting"
	// StateWaitingForCapacity specifies the state waiting for capacity state for state.
	StateWaitingForCapacity State = "WaitingForCapacity"
)

// Type enumerates the values for type.
type Type string

const (
	// Hive specifies the hive state for type.
	Hive Type = "Hive"
	// USQL specifies the usql state for type.
	USQL Type = "USql"
)

// DataPath is a Data Lake Analytics U-SQL job data path item.
type DataPath struct {
	autorest.Response `json:"-"`
	JobID             *uuid.UUID `json:"jobId,omitempty"`
	Command           *string    `json:"command,omitempty"`
	Paths             *[]string  `json:"paths,omitempty"`
}

// ErrorDetails is the Data Lake Analytics job error details.
type ErrorDetails struct {
	Description         *string       `json:"description,omitempty"`
	Details             *string       `json:"details,omitempty"`
	EndOffset           *int32        `json:"endOffset,omitempty"`
	ErrorID             *string       `json:"errorId,omitempty"`
	FilePath            *string       `json:"filePath,omitempty"`
	HelpLink            *string       `json:"helpLink,omitempty"`
	InternalDiagnostics *string       `json:"internalDiagnostics,omitempty"`
	LineNumber          *int32        `json:"lineNumber,omitempty"`
	Message             *string       `json:"message,omitempty"`
	Resolution          *string       `json:"resolution,omitempty"`
	InnerError          *InnerError   `json:"InnerError,omitempty"`
	Severity            SeverityTypes `json:"severity,omitempty"`
	Source              *string       `json:"source,omitempty"`
	StartOffset         *int32        `json:"startOffset,omitempty"`
}

// HiveJobProperties is
type HiveJobProperties struct {
	RuntimeVersion         *string                 `json:"runtimeVersion,omitempty"`
	Script                 *string                 `json:"script,omitempty"`
	Type                   Type                    `json:"type,omitempty"`
	StatementInfo          *[]HiveJobStatementInfo `json:"statementInfo,omitempty"`
	LogsLocation           *string                 `json:"logsLocation,omitempty"`
	WarehouseLocation      *string                 `json:"warehouseLocation,omitempty"`
	StatementCount         *int32                  `json:"statementCount,omitempty"`
	ExecutedStatementCount *int32                  `json:"executedStatementCount,omitempty"`
}

// MarshalJSON is the custom marshaler for HiveJobProperties.
func (hjp HiveJobProperties) MarshalJSON() ([]byte, error) {
	hjp.Type = Hive
	type Alias HiveJobProperties
	return json.Marshal(&struct {
		Alias
	}{
		Alias: (Alias)(hjp),
	})
}

// AsUSQLJobProperties is the Properties implementation for HiveJobProperties.
func (hjp HiveJobProperties) AsUSQLJobProperties() (*USQLJobProperties, bool) {
	return nil, false
}

// AsHiveJobProperties is the Properties implementation for HiveJobProperties.
func (hjp HiveJobProperties) AsHiveJobProperties() (*HiveJobProperties, bool) {
	return &hjp, true
}

// HiveJobStatementInfo is
type HiveJobStatementInfo struct {
	LogLocation           *string `json:"logLocation,omitempty"`
	ResultPreviewLocation *string `json:"resultPreviewLocation,omitempty"`
	ResultLocation        *string `json:"resultLocation,omitempty"`
	ErrorMessage          *string `json:"errorMessage,omitempty"`
}

// InfoListResult is list of jobInfo items.
type InfoListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Information `json:"value,omitempty"`
	NextLink          *string        `json:"nextLink,omitempty"`
	Count             *int64         `json:"count,omitempty"`
}

// InfoListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client InfoListResult) InfoListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// Information is the common Data Lake Analytics job information properties.
type Information struct {
	autorest.Response   `json:"-"`
	JobID               *uuid.UUID          `json:"jobId,omitempty"`
	Name                *string             `json:"name,omitempty"`
	Type                Type                `json:"type,omitempty"`
	Submitter           *string             `json:"submitter,omitempty"`
	ErrorMessage        *[]ErrorDetails     `json:"errorMessage,omitempty"`
	DegreeOfParallelism *int32              `json:"degreeOfParallelism,omitempty"`
	Priority            *int32              `json:"priority,omitempty"`
	SubmitTime          *date.Time          `json:"submitTime,omitempty"`
	StartTime           *date.Time          `json:"startTime,omitempty"`
	EndTime             *date.Time          `json:"endTime,omitempty"`
	State               State               `json:"state,omitempty"`
	Result              Result              `json:"result,omitempty"`
	StateAuditRecords   *[]StateAuditRecord `json:"stateAuditRecords,omitempty"`
	Properties          Properties          `json:"properties,omitempty"`
}

// UnmarshalJSON is the custom unmarshaler for Information struct.
func (i *Information) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	var v *json.RawMessage

	v = m["jobId"]
	if v != nil {
		var jobID uuid.UUID
		err = json.Unmarshal(*m["jobId"], &jobID)
		if err != nil {
			return err
		}
		i.JobID = &jobID
	}

	v = m["name"]
	if v != nil {
		var name string
		err = json.Unmarshal(*m["name"], &name)
		if err != nil {
			return err
		}
		i.Name = &name
	}

	v = m["type"]
	if v != nil {
		var typeVar Type
		err = json.Unmarshal(*m["type"], &typeVar)
		if err != nil {
			return err
		}
		i.Type = typeVar
	}

	v = m["submitter"]
	if v != nil {
		var submitter string
		err = json.Unmarshal(*m["submitter"], &submitter)
		if err != nil {
			return err
		}
		i.Submitter = &submitter
	}

	v = m["errorMessage"]
	if v != nil {
		var errorMessage []ErrorDetails
		err = json.Unmarshal(*m["errorMessage"], &errorMessage)
		if err != nil {
			return err
		}
		i.ErrorMessage = &errorMessage
	}

	v = m["degreeOfParallelism"]
	if v != nil {
		var degreeOfParallelism int32
		err = json.Unmarshal(*m["degreeOfParallelism"], &degreeOfParallelism)
		if err != nil {
			return err
		}
		i.DegreeOfParallelism = &degreeOfParallelism
	}

	v = m["priority"]
	if v != nil {
		var priority int32
		err = json.Unmarshal(*m["priority"], &priority)
		if err != nil {
			return err
		}
		i.Priority = &priority
	}

	v = m["submitTime"]
	if v != nil {
		var submitTime date.Time
		err = json.Unmarshal(*m["submitTime"], &submitTime)
		if err != nil {
			return err
		}
		i.SubmitTime = &submitTime
	}

	v = m["startTime"]
	if v != nil {
		var startTime date.Time
		err = json.Unmarshal(*m["startTime"], &startTime)
		if err != nil {
			return err
		}
		i.StartTime = &startTime
	}

	v = m["endTime"]
	if v != nil {
		var endTime date.Time
		err = json.Unmarshal(*m["endTime"], &endTime)
		if err != nil {
			return err
		}
		i.EndTime = &endTime
	}

	v = m["state"]
	if v != nil {
		var state State
		err = json.Unmarshal(*m["state"], &state)
		if err != nil {
			return err
		}
		i.State = state
	}

	v = m["result"]
	if v != nil {
		var resultVar Result
		err = json.Unmarshal(*m["result"], &resultVar)
		if err != nil {
			return err
		}
		i.Result = resultVar
	}

	v = m["stateAuditRecords"]
	if v != nil {
		var stateAuditRecords []StateAuditRecord
		err = json.Unmarshal(*m["stateAuditRecords"], &stateAuditRecords)
		if err != nil {
			return err
		}
		i.StateAuditRecords = &stateAuditRecords
	}

	v = m["properties"]
	if v != nil {
		properties, err := unmarshalProperties(*m["properties"])
		if err != nil {
			return err
		}
		i.Properties = properties
	}

	return nil
}

// InnerError is the Data Lake Analytics job error details.
type InnerError struct {
	DiagnosticCode      *int32        `json:"diagnosticCode,omitempty"`
	Severity            SeverityTypes `json:"severity,omitempty"`
	Details             *string       `json:"details,omitempty"`
	Component           *string       `json:"component,omitempty"`
	ErrorID             *string       `json:"errorId,omitempty"`
	HelpLink            *string       `json:"helpLink,omitempty"`
	InternalDiagnostics *string       `json:"internalDiagnostics,omitempty"`
	Message             *string       `json:"message,omitempty"`
	Resolution          *string       `json:"resolution,omitempty"`
	Source              *string       `json:"source,omitempty"`
	Description         *string       `json:"description,omitempty"`
}

// Properties is the common Data Lake Analytics job properties.
type Properties interface {
	AsUSQLJobProperties() (*USQLJobProperties, bool)
	AsHiveJobProperties() (*HiveJobProperties, bool)
}

func unmarshalProperties(body []byte) (Properties, error) {
	var m map[string]interface{}
	err := json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}

	switch m["type"] {
	case string(USQL):
		var usjp USQLJobProperties
		err := json.Unmarshal(body, &usjp)
		return usjp, err
	case string(Hive):
		var hjp HiveJobProperties
		err := json.Unmarshal(body, &hjp)
		return hjp, err
	default:
		return nil, errors.New("Unsupported type")
	}
}
func unmarshalPropertiesArray(body []byte) ([]Properties, error) {
	var rawMessages []*json.RawMessage
	err := json.Unmarshal(body, &rawMessages)
	if err != nil {
		return nil, err
	}

	pArray := make([]Properties, len(rawMessages))

	for index, rawMessage := range rawMessages {
		p, err := unmarshalProperties(*rawMessage)
		if err != nil {
			return nil, err
		}
		pArray[index] = p
	}
	return pArray, nil
}

// Resource is the Data Lake Analytics U-SQL job resources.
type Resource struct {
	Name         *string      `json:"name,omitempty"`
	ResourcePath *string      `json:"resourcePath,omitempty"`
	Type         ResourceType `json:"type,omitempty"`
}

// StateAuditRecord is the Data Lake Analytics U-SQL job state audit records for tracking the lifecycle of a job.
type StateAuditRecord struct {
	NewState        *string    `json:"newState,omitempty"`
	TimeStamp       *date.Time `json:"timeStamp,omitempty"`
	RequestedByUser *string    `json:"requestedByUser,omitempty"`
	Details         *string    `json:"details,omitempty"`
}

// Statistics is the Data Lake Analytics U-SQL job execution statistics.
type Statistics struct {
	autorest.Response `json:"-"`
	LastUpdateTimeUtc *date.Time               `json:"lastUpdateTimeUtc,omitempty"`
	Stages            *[]StatisticsVertexStage `json:"stages,omitempty"`
}

// StatisticsVertexStage is the Data Lake Analytics U-SQL job statistics vertex stage information.
type StatisticsVertexStage struct {
	DataRead              *int64  `json:"dataRead,omitempty"`
	DataReadCrossPod      *int64  `json:"dataReadCrossPod,omitempty"`
	DataReadIntraPod      *int64  `json:"dataReadIntraPod,omitempty"`
	DataToRead            *int64  `json:"dataToRead,omitempty"`
	DataWritten           *int64  `json:"dataWritten,omitempty"`
	DuplicateDiscardCount *int32  `json:"duplicateDiscardCount,omitempty"`
	FailedCount           *int32  `json:"failedCount,omitempty"`
	MaxVertexDataRead     *int64  `json:"maxVertexDataRead,omitempty"`
	MinVertexDataRead     *int64  `json:"minVertexDataRead,omitempty"`
	ReadFailureCount      *int32  `json:"readFailureCount,omitempty"`
	RevocationCount       *int32  `json:"revocationCount,omitempty"`
	RunningCount          *int32  `json:"runningCount,omitempty"`
	ScheduledCount        *int32  `json:"scheduledCount,omitempty"`
	StageName             *string `json:"stageName,omitempty"`
	SucceededCount        *int32  `json:"succeededCount,omitempty"`
	TempDataWritten       *int64  `json:"tempDataWritten,omitempty"`
	TotalCount            *int32  `json:"totalCount,omitempty"`
	TotalFailedTime       *string `json:"totalFailedTime,omitempty"`
	TotalProgress         *int32  `json:"totalProgress,omitempty"`
	TotalSucceededTime    *string `json:"totalSucceededTime,omitempty"`
}

// USQLJobProperties is
type USQLJobProperties struct {
	RuntimeVersion           *string     `json:"runtimeVersion,omitempty"`
	Script                   *string     `json:"script,omitempty"`
	Type                     Type        `json:"type,omitempty"`
	Resources                *[]Resource `json:"resources,omitempty"`
	Statistics               *Statistics `json:"statistics,omitempty"`
	DebugData                *DataPath   `json:"debugData,omitempty"`
	AlgebraFilePath          *string     `json:"algebraFilePath,omitempty"`
	TotalCompilationTime     *string     `json:"totalCompilationTime,omitempty"`
	TotalPauseTime           *string     `json:"totalPauseTime,omitempty"`
	TotalQueuedTime          *string     `json:"totalQueuedTime,omitempty"`
	TotalRunningTime         *string     `json:"totalRunningTime,omitempty"`
	RootProcessNodeID        *string     `json:"rootProcessNodeId,omitempty"`
	YarnApplicationID        *string     `json:"yarnApplicationId,omitempty"`
	YarnApplicationTimeStamp *int64      `json:"yarnApplicationTimeStamp,omitempty"`
	CompileMode              CompileMode `json:"compileMode,omitempty"`
}

// MarshalJSON is the custom marshaler for USQLJobProperties.
func (usjp USQLJobProperties) MarshalJSON() ([]byte, error) {
	usjp.Type = USQL
	type Alias USQLJobProperties
	return json.Marshal(&struct {
		Alias
	}{
		Alias: (Alias)(usjp),
	})
}

// AsUSQLJobProperties is the Properties implementation for USQLJobProperties.
func (usjp USQLJobProperties) AsUSQLJobProperties() (*USQLJobProperties, bool) {
	return &usjp, true
}

// AsHiveJobProperties is the Properties implementation for USQLJobProperties.
func (usjp USQLJobProperties) AsHiveJobProperties() (*HiveJobProperties, bool) {
	return nil, false
}
