package xfsquota

import (
	"fmt"
	"strings"
)

type LimitCommandOption struct {
	// Equal to `-g` flag on commandline.
	// Set when specifying Id or Name as that of a Group.
	Group bool
	// Equal to `-u` flag on commandline.
	// Set when specifying Id or Name as that of a User.
	User bool
	// Equal to `-p` flag on commandline.
	// Set when specifying Id or Name as that of a Project.
	Project bool
	// Equal to `bsoft=N` argument on commandline.
	// Set quota block soft limits.
	Bsoft uint32
	// Equal to `bhard=N` argument on commandline.
	// Set quota block hard limits.
	Bhard uint32
	// Equal to `isoft=N` argument on commandline.
	// Set quota inode count soft limits.
	Isoft uint32
	// Equal to `ihard=N` argument on commandline.
	// Set quota inode count hard limits.
	Ihard uint32
	// Equal to `rtbsoft=N` argument on commandline.
	// Set quota realtime block soft limits.
	Rtbsoft uint32
	// Equal to `rtbhard=N` argument on commandline.
	// Set quota realtime block hard limits.
	Rtbhard uint32
	// ID to target
	Id []uint32
	// Name to target
	Name []string
}

func (o LimitCommandOption) SubCommandString() string {
	cmds := []string{}
	cmds = append(cmds, "limit")

	if o.Group {
		cmds = append(cmds, "-g")
	}

	if o.Project {
		cmds = append(cmds, "-p")
	}

	if o.User {
		cmds = append(cmds, "-u")
	}

	if o.Bsoft != 0 {
		cmds = append(cmds, fmt.Sprintf("bsoft=%d", o.Bsoft))
	}

	if o.Bhard != 0 {
		cmds = append(cmds, fmt.Sprintf("bhard=%d", o.Bhard))
	}

	if o.Isoft != 0 {
		cmds = append(cmds, fmt.Sprintf("isoft=%d", o.Isoft))
	}

	if o.Ihard != 0 {
		cmds = append(cmds, fmt.Sprintf("ihard=%d", o.Ihard))
	}

	if o.Rtbsoft != 0 {
		cmds = append(cmds, fmt.Sprintf("rtbsoft=%d", o.Rtbsoft))
	}

	if o.Rtbhard != 0 {
		cmds = append(cmds, fmt.Sprintf("rtbhard=%d", o.Rtbhard))
	}

	return strings.Join(cmds, " ")
}

func (c *Command) Limit(globalOpt GlobalOption, subopt LimitCommandOption) error {
	c.SubOpt = subopt
	c.GlobalOpt = globalOpt
	return c.Execute()
}

func (c *Command) LimitProject(bsoft, bhard, isoft, ihard, rtbsoft, rtbhard uint32) error {
	opt := LimitCommandOption{
    Project: true,

		Bsoft:   bsoft,
		Bhard:   bhard,
		Isoft:   isoft,
		Ihard:   ihard,
		Rtbsoft: rtbsoft,
		Rtbhard: rtbhard,
	}

	gopt := GlobalOption{
		EnableExpertMode: true,
	}

	return c.Limit(gopt, opt)
}