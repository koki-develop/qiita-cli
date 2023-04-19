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
	Required    bool
}

func (f *Flag) Changed(cmd *cobra.Command) bool {
	return cmd.Flag(f.Name).Changed
}

type String struct {
	*Flag
	Default string
	value   string
}

func (f *String) AddTo(cmds ...*cobra.Command) {
	for _, cmd := range cmds {
		if f.ShortName != "" {
			cmd.Flags().StringVarP(&f.value, f.Name, f.ShortName, f.Default, f.Description)
		} else {
			cmd.Flags().StringVar(&f.value, f.Name, f.Default, f.Description)
		}
		if f.Required {
			if err := cmd.MarkFlagRequired(f.Name); err != nil {
				panic(err)
			}
		}
	}
}

func (f *String) Get(cmd *cobra.Command, nonnull bool) *string {
	if nonnull || f.Changed(cmd) {
		return &f.value
	}
	return nil
}

type Int struct {
	*Flag
	Default int
	value   int
}

func (f *Int) AddTo(cmds ...*cobra.Command) {
	for _, cmd := range cmds {
		if f.ShortName != "" {
			cmd.Flags().IntVarP(&f.value, f.Name, f.ShortName, f.Default, f.Description)
		} else {
			cmd.Flags().IntVar(&f.value, f.Name, f.Default, f.Description)
		}
		if f.Required {
			if err := cmd.MarkFlagRequired(f.Name); err != nil {
				panic(err)
			}
		}
	}
}

func (f *Int) Get(cmd *cobra.Command, nonnull bool) *int {
	if nonnull || f.Changed(cmd) {
		return &f.value
	}
	return nil
}

type Bool struct {
	*Flag
	Default bool
	value   bool
}

func (f *Bool) AddTo(cmds ...*cobra.Command) {
	for _, cmd := range cmds {
		if f.ShortName != "" {
			cmd.Flags().BoolVarP(&f.value, f.Name, f.ShortName, f.Default, f.Description)
		} else {
			cmd.Flags().BoolVar(&f.value, f.Name, f.Default, f.Description)
		}
		if f.Required {
			if err := cmd.MarkFlagRequired(f.Name); err != nil {
				panic(err)
			}
		}
	}
}

func (f *Bool) Get(cmd *cobra.Command, nonnull bool) *bool {
	if nonnull || f.Changed(cmd) {
		return &f.value
	}
	return nil
}

type StringSlice struct {
	*Flag
	Default []string
	value   []string
}

func (f *StringSlice) AddTo(cmds ...*cobra.Command) {
	for _, cmd := range cmds {
		if f.ShortName != "" {
			cmd.Flags().StringSliceVarP(&f.value, f.Name, f.ShortName, f.Default, f.Description)
		} else {
			cmd.Flags().StringSliceVar(&f.value, f.Name, f.Default, f.Description)
		}
		if f.Required {
			if err := cmd.MarkFlagRequired(f.Name); err != nil {
				panic(err)
			}
		}
	}
}

func (f *StringSlice) Get(cmd *cobra.Command, nonnull bool) *[]string {
	if nonnull || f.Changed(cmd) {
		return &f.value
	}
	return nil
}
