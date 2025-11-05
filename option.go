package config

import "time"

type Option struct {
	Filename    OptionString
	FileType    OptionString
	Path        OptionString
	Env         OptionString
	DebounceDur OptionTimeDuration
}

// NewOption 创建默认配置
func NewOption() *Option {
	opt := &Option{}
	opt.setDefaultValue()
	return opt
}

// 初始化默认值
func (s *Option) setDefaultValue() *Option {
	s.Filename.Set(OptionFilename, false)
	s.FileType.Set(OptionFileType, false)
	s.Path.Set(OptionPath, false)
	s.Env.Set(OptionEnv, false)
	s.DebounceDur.Set(OptionTimeDuration(OptionDebounceDur), false)
	return s
}

type OptionString string
type OptionTimeDuration time.Duration

func (o *OptionString) Set(newStr OptionString, reset ...bool) {
	if len(reset) == 0 {
		reset = []bool{true}
	}
	if *o != "" && !reset[0] {
		return
	}
	*o = newStr
}

func (o *OptionString) ToValue() string {
	return string(*o)
}

func (o *OptionTimeDuration) Set(newDate OptionTimeDuration, reset ...bool) {
	if len(reset) == 0 {
		reset = []bool{true}
	}
	if *o != 0 && !reset[0] {
		return
	}
	*o = newDate
}

func (o *OptionTimeDuration) ToValue() time.Duration {
	return time.Duration(*o)
}
