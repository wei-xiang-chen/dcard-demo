package setting

type RedisSetting struct {
	Addr     string
	Password string
	DB       int
}

var (
	RsSetting *RedisSetting
)

func (s *Setting) ReadRedisSetting() error {
	err := s.vp.UnmarshalKey("Redis", &RsSetting)
	if err != nil {
		return err
	}

	return nil
}
