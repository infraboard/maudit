package client_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/infraboard/maudit/apps/book"
	"github.com/infraboard/maudit/client"
)

func TestBookQuery(t *testing.T) {
	should := assert.New(t)

	c, err := client.NewClient(client.NewDefaultConfig())
	should.NoError(err)

	resp, err := c.Book().QueryBook(
		context.Background(),
		book.NewQueryBookRequest(),
	)
	should.NoError(err)
	fmt.Println(resp.Items)
}
