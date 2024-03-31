// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package googlecloudstorageexporter // Package googlecloudstorageexporter import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/googlecloudstorageexporter"
import (
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

type Config struct {
	exporterhelper.TimeoutSettings `mapstructure:",squash"`
	exporterhelper.QueueSettings   `mapstructure:"sending_queue"`
	configretry.BackOffConfig      `mapstructure:"retry_on_failure"`
	ProjectID                      string `mapstructure:"project"`
	BucketName                     string `mapstructure:"bucket"`
}
