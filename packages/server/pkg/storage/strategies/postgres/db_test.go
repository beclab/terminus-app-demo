package postgres

import "testing"

func TestNewDB(t *testing.T) {
	op := New()
	var res []string
	err := op.db.Raw("select 'aaa'").Scan(&res).Error
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	t.Log(res)
}
