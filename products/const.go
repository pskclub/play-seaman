package products

const (
	DefaultSomething int = 10
)

// ShipmentType is e-commerce shipment type
type ShipmentType string

// Shipment types
const (
	ShipmentTypeOnePackagePerOrder ShipmentType = "ONE_PACKAGE_PER_ORDER"
	ShipmentTypeOnePackagePerSKU   ShipmentType = "ONE_PACKAGE_PER_SKU"
)

// Kafka topic
const (
	// TopicOrderCreated using when create order success
	TopicOrderCreated   = "order_created"
	TopicPaymentSuccess = "order_paid_success"
)

// ConsumerGroupShipment is shipment consumer group
const (
	ConsumerGroupShipment    = "shipment"
	ConsumerGroupTransaction = "transaction"
)
