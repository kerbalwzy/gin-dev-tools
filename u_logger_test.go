package kerbalwzygo

import "testing"

func TestXLogger(t *testing.T) {
	logger = GetLogger()
	logger.SetLevel(Debug)
	logger.Debug("Test Debug")
	logger.SetLevel(Info)
	logger.Debug("Test Debug 2")
	logger.Info("Test Info")
	t.Logf("%p", logger)
	logger = GetLogger()
	t.Logf("%p", logger)
}
