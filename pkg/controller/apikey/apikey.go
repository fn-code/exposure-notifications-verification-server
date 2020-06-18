// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package apikey contains web controllers for listing and adding API Keys.
package apikey

import (
	"context"
	"net/http"

	"github.com/google/exposure-notifications-verification-server/pkg/config"
	"github.com/google/exposure-notifications-verification-server/pkg/controller"
	"github.com/google/exposure-notifications-verification-server/pkg/controller/flash"
	"github.com/google/exposure-notifications-verification-server/pkg/database"
	"github.com/google/exposure-notifications-verification-server/pkg/logging"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type apikeyListController struct {
	config *config.Config
	db     *database.Database
	logger *zap.SugaredLogger
}

func NewListController(ctx context.Context, config *config.Config, db *database.Database) controller.Controller {
	return &apikeyListController{config, db, logging.FromContext(ctx)}
}

func (lc *apikeyListController) Execute(c *gin.Context) {
	user := c.MustGet("user").(*database.User)
	flash := flash.FromContext(c)

	m := controller.NewTemplateMapFromSession(lc.config, c)
	m["user"] = user

	apps, err := lc.db.ListAuthorizedApps(true)
	if err != nil {
		flash.ErrorNow("Error loading API Keys: %v", err)
	}

	m["apps"] = apps
	c.HTML(http.StatusOK, "apikeys", m)
}