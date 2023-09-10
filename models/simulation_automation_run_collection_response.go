package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SimulationAutomationRunCollectionResponse 
type SimulationAutomationRunCollectionResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewSimulationAutomationRunCollectionResponse instantiates a new simulationAutomationRunCollectionResponse and sets the default values.
func NewSimulationAutomationRunCollectionResponse()(*SimulationAutomationRunCollectionResponse) {
    m := &SimulationAutomationRunCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateSimulationAutomationRunCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSimulationAutomationRunCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSimulationAutomationRunCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SimulationAutomationRunCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSimulationAutomationRunFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SimulationAutomationRunable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SimulationAutomationRunable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *SimulationAutomationRunCollectionResponse) GetValue()([]SimulationAutomationRunable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SimulationAutomationRunable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SimulationAutomationRunCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *SimulationAutomationRunCollectionResponse) SetValue(value []SimulationAutomationRunable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
// SimulationAutomationRunCollectionResponseable 
type SimulationAutomationRunCollectionResponseable interface {
    BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]SimulationAutomationRunable)
    SetValue(value []SimulationAutomationRunable)()
}
