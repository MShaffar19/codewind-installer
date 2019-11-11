/*******************************************************************************
 * Copyright (c) 2019 IBM Corporation and others.
 * All rights reserved. This program and the accompanying materials
 * are made available under the terms of the Eclipse Public License v2.0
 * which accompanies this distribution, and is available at
 * http://www.eclipse.org/legal/epl-v20.html
 *
 * Contributors:
 *     IBM Corporation - initial API and implementation
 *******************************************************************************/

package actions

import (
	"strings"

	"github.com/eclipse/codewind-installer/pkg/utils"
	logr "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

//RemoveCommand to remove all codewind and project images
func RemoveCommand(c *cli.Context) {
	tag := c.String("tag")
	imageArr := []string{
		"eclipse/codewind-pfe-amd64:" + tag,
		"eclipse/codewind-performance-amd64:" + tag,
		"cw-",
	}
	networkName := "codewind"

	images := utils.GetImageList()

	logr.Infoln("Removing Codewind docker images..")

	for _, image := range images {
		imageRepo := strings.Join(image.RepoDigests, " ")
		imageTags := strings.Join(image.RepoTags, " ")
		for _, key := range imageArr {
			if strings.HasPrefix(imageRepo, key) || strings.HasPrefix(imageTags, key) {
				if len(image.RepoTags) > 0 {
					logr.Infoln("Deleting Image ", image.RepoTags[0], "... ")
				} else {
					logr.Infoln("Deleting Image ", image.ID, "... ")
				}
				utils.RemoveImage(image.ID)
			}
		}
	}

	networks := utils.GetNetworkList()

	for _, network := range networks {
		if strings.Contains(network.Name, networkName) {
			logr.Infoln("Removing docker network: ", network.Name, "... ")
			utils.RemoveNetwork(network)
		}
	}
}
