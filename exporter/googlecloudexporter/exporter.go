package googlecloudexporter

import (
	"go.opentelemetry.io/collector/exporter"
	"go.uber.org/zap"
)

type gcsExporter struct {
	config *Config
	logger *zap.Logger
	//TODO(swiftdiaries): Create datawriter, marshaler similar to S3Exporter
}

func newGCSExporter(config *Config, params exporter.CreateSettings) *gcsExporter {
	gcsExporter := &gcsExporter{
		config: config,
		logger: params.Logger,
	}
	return gcsExporter
}
