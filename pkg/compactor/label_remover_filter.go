// SPDX-License-Identifier: AGPL-3.0-only
// Provenance-includes-location: https://github.com/cortexproject/cortex/blob/master/pkg/compactor/label_remover_filter.go
// Provenance-includes-license: Apache-2.0
// Provenance-includes-copyright: The Cortex Authors.

package compactor

import (
	"context"

	"github.com/oklog/ulid"
	"github.com/thanos-io/thanos/pkg/block/metadata"
	"github.com/thanos-io/thanos/pkg/extprom"
)

type LabelRemoverFilter struct {
	labels []string
}

// NewLabelRemoverFilter creates a LabelRemoverFilter.
func NewLabelRemoverFilter(labels []string) *LabelRemoverFilter {
	return &LabelRemoverFilter{labels: labels}
}

// Filter modifies external labels of existing blocks, removing given labels from the metadata of blocks that have it.
func (f *LabelRemoverFilter) Filter(_ context.Context, metas map[ulid.ULID]*metadata.Meta, _ *extprom.TxGaugeVec, _ *extprom.TxGaugeVec) error {
	for _, meta := range metas {
		for _, l := range f.labels {
			delete(meta.Thanos.Labels, l)
		}
	}

	return nil
}
