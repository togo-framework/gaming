// Package gaming is the Gaming content pack for togo: GamingActivity resource: platform, rank, hours, gradient theming and looking-for-team, with slug detail pages.
//
// A content pack bundles resource definitions (model + sqlc + Atlas + REST +
// GraphQL + a page) that a togo app pulls in. Generate the resource with
// `togo make:resource` in the host app; this plugin wires the provider hook.
package gaming

import "github.com/togo-framework/togo"

func init() {
	togo.RegisterProviderFunc("gaming", togo.PriorityService, func(k *togo.Kernel) error {
		return nil
	})
}
