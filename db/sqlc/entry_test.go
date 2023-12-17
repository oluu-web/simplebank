package db

import (
	"simplebank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateEntry(t *testing.T) {
	arg := CreateEntryParams{
		AccountID: util.RandomInt(2, 100),
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(ctx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.AccountID)
	require.NotZero(t, entry.Amount)
}

func createRandomEntry(t *testing.T) Entry {
	arg := CreateEntryParams{
		AccountID: util.RandomInt(2, 100),
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(ctx, arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.AccountID)
	require.NotZero(t, entry.Amount)

	return entry
}

func TestGetEntry(t *testing.T) {
	entry1 := createRandomEntry(t)
	entry2, err := testQueries.GetEntry(ctx, entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt.Time, entry2.CreatedAt.Time, time.Second)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntry(t)
	}

	arg := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}

	entries, err := testQueries.ListEntries(ctx, arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
