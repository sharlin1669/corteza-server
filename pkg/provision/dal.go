package provision

import (
	"context"

	"github.com/cortezaproject/corteza-server/pkg/auth"
	"github.com/cortezaproject/corteza-server/pkg/dal"
	"github.com/cortezaproject/corteza-server/pkg/id"
	"github.com/cortezaproject/corteza-server/store"
	"github.com/cortezaproject/corteza-server/system/types"
)

// Injects primary connection
func defaultDalConnection(ctx context.Context, s store.DalConnections) error {
	conn, err := store.LookupDalConnectionByHandle(ctx, s, types.DalPrimaryConnectionHandle)
	if err != nil && err != store.ErrNotFound {
		return err
	}

	// Already exists
	if conn != nil {
		return nil
	}

	conn = &types.DalConnection{
		// Using id.Next since we dropped "special" ids a while ago.
		// If needed, use the handle
		ID:     id.Next(),
		Handle: types.DalPrimaryConnectionHandle,
		Type:   types.DalPrimaryConnectionResourceType,

		Meta: types.ConnectionMeta{
			Name: "Primary Database",
		},

		Config: types.ConnectionConfig{
			DAL: &types.ConnectionConfigDAL{
				ModelIdent: "compose_record",
				Operations: dal.FullOperations(),
			},
		},

		CreatedAt: *now(),
		CreatedBy: auth.ServiceUser().ID,
	}

	return store.CreateDalConnection(ctx, s, conn)
}
