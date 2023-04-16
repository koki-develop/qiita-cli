package flags

import "github.com/spf13/cobra"

type IFlag interface {
	AddTo(cmds ...*cobra.Command)
}

type Flags []IFlag

func (fs Flags) AddTo(cmds ...*cobra.Command) {
	for _, f := range fs {
		for _, cmd := range cmds {
			f.AddTo(cmd)
		}
	}
}

var (
	_ IFlag = (*String)(nil)
	_ IFlag = (*Int)(nil)
	_ IFlag = (*StringSlice)(nil)
)

type Flag struct {
	Name        string
	ShortName   string
	Description string
}

func (f *Flag) Changed(cmd *cobra.Command) bool {
	return cmd.Flag(f.Name).Changed
}

type String struct {
	*Flag
	Default string
	Value   string
}

func (f *String) AddTo(cmds ...*cobra.Command) {
	for _, cmd := range cmds {
		if f.ShortName != "" {
			cmd.Flags().StringVarP(&f.Value, f.Name, f.ShortName, f.Default, f.Description)
		} else {
			cmd.Flags().StringVar(&f.Value, f.Name, f.Default, f.Description)
		}
	}
}

func (f *String) Get(cmd *cobra.Command, nonnull bool) *string {
	if nonnull || f.Changed(cmd) {
		return &f.Value
	}
	return nil
}

type Int struct {
	*Flag
	Default int
	Value   int
}

func (f *Int) AddTo(cmds ...*cobra.Command) {
	for _, cmd := range cmds {
		if f.ShortName != "" {
			cmd.Flags().IntVarP(&f.Value, f.Name, f.ShortName, f.Default, f.Description)
		} else {
			cmd.Flags().IntVar(&f.Value, f.Name, f.Default, f.Description)
		}
	}
}

func (f *Int) Get(cmd *cobra.Command, nonnull bool) *int {
	if nonnull || f.Changed(cmd) {
		return &f.Value
	}
	return nil
}

type StringSlice struct {
	*Flag
	Default []string
	Value   []string
}

func (f *StringSlice) AddTo(cmds ...*cobra.Command) {
	for _, cmd := range cmds {
		if f.ShortName != "" {
			cmd.Flags().StringSliceVarP(&f.Value, f.Name, f.ShortName, f.Default, f.Description)
		} else {
			cmd.Flags().StringSliceVar(&f.Value, f.Name, f.Default, f.Description)
		}
	}
}

func (f *StringSlice) Get(cmd *cobra.Command, nonnull bool) *[]string {
	if nonnull || f.Changed(cmd) {
		return &f.Value
	}
	return nil
}
