package usercmd

import (
	"context"

	"github.com/fatih/color"
	"github.com/pkg/errors"
)

// ChartDefaultSet sets the local app chart default
func (c *EpinioClient) ChartDefaultSet(ctx context.Context, chartName string) error {
	log := c.Log.WithName("ChartDefaultSet")
	log.Info("start")
	defer log.Info("return")

	// Validate chosen app chart to exist
	if chartName != "" {
		_, err := c.API.ChartShow(chartName)
		if err != nil {
			return err
		}
	}

	// Save to settings
	c.Settings.AppChart = chartName
	err := c.Settings.Save()
	if err != nil {
		return errors.Wrap(err, "failed to save settings")
	}

	// And report
	if chartName == "" {
		c.ui.Note().
			Msg("Unset Default Application Chart")
	} else {
		c.ui.Note().
			WithStringValue("Name", c.Settings.AppChart).
			Msg("New Default Application Chart")
	}

	return nil
}

// ChartDefaultShow displays the current local app chart default
func (c *EpinioClient) ChartDefaultShow(ctx context.Context) error {
	log := c.Log.WithName("ChartDefaultShow")
	log.Info("start")
	defer log.Info("return")

	chart := c.Settings.AppChart
	if chart == "" {
		chart = color.CyanString("not set, system default applies")
	}

	c.ui.Note().
		WithStringValue("Name", chart).
		Msg("Default Application Chart")

	return nil
}

// ChartList displays a table of all known application charts.
func (c *EpinioClient) ChartList(ctx context.Context) error {
	log := c.Log.WithName("ChartList")
	log.Info("start")
	defer log.Info("return")

	c.ui.Note().
		Msg("Show Application Charts")

	charts, err := c.API.ChartList()
	if err != nil {
		return err
	}

	msg := c.ui.Success().WithTable("Default", "Name", "Created", "Description")

	for _, chart := range charts {
		mark := ""
		name := chart.Meta.Name
		created := chart.Meta.CreatedAt.String()
		short := chart.ShortDescription
		if chart.Meta.Name == c.Settings.AppChart {
			mark = color.BlueString("*")
			name = color.BlueString(name)
			created = color.BlueString(created)
			short = color.BlueString(short)
		}
		msg = msg.WithTableRow(mark, name, created, short)
	}

	msg.Msg("Ok")
	return nil
}

// ChartShow shows the value of the specified environment variable in
// the named application.
func (c *EpinioClient) ChartShow(ctx context.Context, name string) error {
	log := c.Log.WithName("ChartShow")
	log.Info("start")
	defer log.Info("return")

	c.ui.Note().
		WithStringValue("Name", name).
		Msg("Show application chart details")

	chart, err := c.API.ChartShow(name)
	if err != nil {
		return err
	}

	c.ui.Success().WithTable("Key", "Value").
		WithTableRow("Name", chart.Meta.Name).
		WithTableRow("Created", chart.Meta.CreatedAt.String()).
		WithTableRow("Short", chart.ShortDescription).
		WithTableRow("Description", chart.Description).
		WithTableRow("Helm Repository", chart.HelmRepo).
		WithTableRow("Helm Chart", chart.HelmChart).
		Msg("Details:")

	return nil
}

// ChartMatching retrieves all application charts in the cluster, for the given prefix
func (c *EpinioClient) ChartMatching(prefix string) []string {
	log := c.Log.WithName("ChartMatching")
	log.Info("start")
	defer log.Info("return")

	resp, err := c.API.ChartMatch(prefix)
	if err != nil {
		log.Error(err, "calling chart match endpoint")
		return []string{}
	}

	return resp.Names
}
