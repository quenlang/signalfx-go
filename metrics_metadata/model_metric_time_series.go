/*
 * Metrics Metadata API
 *
 * API for creating, retrieving, updating, and deleting metric names and MTS metadata.<br> **NOTE:*() Although you can't set custom properties or tags for a metric, you *can* retrieve them for metrics and metric time series (**MTS**).
 *
 * API version: 3.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package metrics_metadata

// The metadata for a single retrieved metric timeseries
type MetricTimeSeries struct {
	// Name of the MTS. Metric names are UTF-8 strings with a maximum length of 256 characters (1024 bytes).
	Metric string `json:"metric,omitempty"`
	// Metric type of the MTS for this metadata. The possible values are \"GAUGE\", \"COUNTER\", and \"CUMULATIVE_COUNTER\".
	Type string `json:"type,omitempty"`
	// Dimension metadata for the MTS, in the form of a JSON object (dictionary). Each property is a dimension name and dimension value. The section [Dimension Criteria](https://developers.signalfx.com/metrics/metrics_metadata_overview.html#_dimension_criteria) lists the requirements for dimensions.
	Dimensions map[string]interface{} `json:"dimensions,omitempty"`
	// Custom property metadata for the MTS, in the form of a JSON object (dictionary). Each property is a custom property name and value. The section [Custom Properties Criteria](https://developers.signalfx.com/metrics/metrics_metadata_overview.html#_custom_properties_criteria) lists the requirements for custom properties. Custom property metadata for the MTS, in the form of a JSON object (dictionary). Each property is a custom property name and value. Custom property names and values have these criteria: <br> **Name:** <br>   * UTF-8 string, maximum length of 128 characters (512 bytes)   * Must start with an uppercase or lowercase letter. The rest of     the name can contain letters, numbers, underscores (`_`) and     hyphens (`-`).   * Must not start with the underscore character (`_`)  <br> **Value:**   * String: Maximum length 256 UTF-8 characters (1024 bytes)   * Integer or float: Maximum length 8192 bits (1024 bytes)
	CustomProperties map[string]string `json:"customProperties,omitempty"`
	// Tag metadata for the MTS, in the form of a JSON array (a list) with one element for each tag. <br> Each tag is a UTF-8 string, starting with an uppercase or lowercase alphabetic character. The maximum length is expressed in characters; if a string consists solely of single-byte UTF-8 entities, 1024 characters are available.
	Tags []string `json:"tags,omitempty"`
	// The time that the MTS was created, in Unix time UTC-relative
	Created int64 `json:"created,omitempty"`
	// SignalFx ID of the user who created the MTS. If the value is \"AAAAAAAAAAA\", SignalFx created the MTS.
	Creator string `json:"creator,omitempty"`
	// The time that the MTS was last updated, in Unix time UTC-relative
	LastUpdated int64 `json:"lastUpdated,omitempty"`
	// SignalFx ID of the user who last updated the MTS. If the value is \"AAAAAAAAAAA\", SignalFx last updated the metric.
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
}