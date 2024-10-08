package valueObject

import "testing"

func TestDatabaseType(t *testing.T) {
	t.Run("ValidDatabaseType", func(t *testing.T) {
		validDbTypes := []interface{}{
			"mariadb", "mysql", "percona", "postgresql", "postgres",
		}

		for _, dbType := range validDbTypes {
			_, err := NewDatabaseType(dbType)
			if err != nil {
				t.Errorf("Expected no error for '%v', got '%s'", dbType, err.Error())
			}
		}
	})

	t.Run("InvalidDatabaseType", func(t *testing.T) {
		invalidDbTypes := []interface{}{
			"cassandra", "sql-server", "cosmosdb",
		}

		for _, dbType := range invalidDbTypes {
			_, err := NewDatabaseType(dbType)
			if err == nil {
				t.Errorf("Expected error for '%v', got nil", dbType)
			}
		}
	})
}
