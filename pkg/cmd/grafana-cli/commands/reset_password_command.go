package commands

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/rkrikbaev/grafana/pkg/bus"
	"github.com/rkrikbaev/grafana/pkg/cmd/grafana-cli/logger"
	"github.com/rkrikbaev/grafana/pkg/models"
	"github.com/rkrikbaev/grafana/pkg/util"
)

const AdminUserId = 1

func resetPasswordCommand(c CommandLine) error {
	newPassword := c.Args().First()

	password := models.Password(newPassword)
	if password.IsWeak() {
		return fmt.Errorf("New password is too short")
	}

	userQuery := models.GetUserByIdQuery{Id: AdminUserId}

	if err := bus.Dispatch(&userQuery); err != nil {
		return fmt.Errorf("Could not read user from database. Error: %v", err)
	}

	passwordHashed := util.EncodePassword(newPassword, userQuery.Result.Salt)

	cmd := models.ChangeUserPasswordCommand{
		UserId:      AdminUserId,
		NewPassword: passwordHashed,
	}

	if err := bus.Dispatch(&cmd); err != nil {
		return fmt.Errorf("Failed to update user password")
	}

	logger.Infof("\n")
	logger.Infof("Admin password changed successfully %s", color.GreenString("✔"))

	return nil
}
