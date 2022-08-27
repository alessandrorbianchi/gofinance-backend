package db

import (
	"context"
	"testing"

	"github.com/alessandrorbianchi/gofinance-backend/util"
	"github.com/stretchr/testify/require"
)

func createRandomCategory(t *testing.T) Category {
	user := createRandomUser(t)
	arg := CreateCategoryParams{
		CoUserID:      user.ID,
		StTitle:       util.RandomString(12),
		StType:        "debit",
		StDescription: util.RandomString(20),
	}

	category, err := testQueries.CreateCategory(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, arg.CoUserID, category.CoUserID)
	require.Equal(t, arg.StTitle, category.StTitle)
	require.Equal(t, arg.StType, category.StType)
	require.Equal(t, arg.StDescription, category.StDescription)
	require.NotEmpty(t, category.DtCreatedAt)

	return category
}

func TestCreateCategory(t *testing.T) {
	createRandomCategory(t)
}

func TestGetCategory(t *testing.T) {
	category1 := createRandomCategory(t)
	category2, err := testQueries.GetCategory(context.Background(), category1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, category1.CoUserID, category2.CoUserID)
	require.Equal(t, category1.StTitle, category2.StTitle)
	require.Equal(t, category1.StType, category2.StType)
	require.Equal(t, category1.StDescription, category2.StDescription)
	require.NotEmpty(t, category2.DtCreatedAt)
}

func TestDeleteCategory(t *testing.T) {
	category1 := createRandomCategory(t)
	err := testQueries.DeleteCategories(context.Background(), category1.ID)
	require.NoError(t, err)
}

func TestUpdateCategory(t *testing.T) {
	category1 := createRandomCategory(t)

	arg := UpdateCategoriesParams{
		ID:            category1.ID,
		StTitle:       util.RandomString(12),
		StDescription: util.RandomString(20),
	}

	category2, err := testQueries.UpdateCategories(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, category2)

	require.Equal(t, category1.ID, category2.ID)
	require.Equal(t, arg.StTitle, category2.StTitle)
	require.Equal(t, arg.StDescription, category2.StDescription)
	require.NotEmpty(t, category2.DtCreatedAt)
}

func TestListCategories(t *testing.T) {
	lastCategory := createRandomCategory(t)

	arg := GetCategoriesParams{
		CoUserID:      lastCategory.CoUserID,
		StType:        lastCategory.StType,
		StTitle:       lastCategory.StTitle,
		StDescription: lastCategory.StDescription,
	}

	categorys, err := testQueries.GetCategories(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, categorys)

	for _, category := range categorys {
		require.Equal(t, lastCategory.ID, category.ID)
		require.Equal(t, lastCategory.CoUserID, category.CoUserID)
		require.Equal(t, lastCategory.StType, category.StType)
		require.Equal(t, lastCategory.StTitle, category.StTitle)
		require.Equal(t, lastCategory.StDescription, category.StDescription)
		require.NotEmpty(t, lastCategory.DtCreatedAt)
	}
}
