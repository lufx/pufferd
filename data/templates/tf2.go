/*
 Copyright 2016 Padduck, LLC

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

 	http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package templates

const TF2 = `{
  "pufferd": {
    "type": "srcds",
    "install": {
      "commands": [
        {
          "files": "https://steamcdn-a.akamaihd.net/client/installer/steamcmd_linux.tar.gz",
          "type": "download"
        },
        {
          "commands": [
            "mkdir steamcmd",
            "tar --no-same-owner -xzvf steamcmd_linux.tar.gz -C steamcmd",
            "steamcmd/steamcmd.sh +login anonymous +force_install_dir ${rootdir} +app_update 232250 +quit",
            "mkdir -p .steam/sdk32",
            "cp steamcmd/linux32/steamclient.so .steam/sdk32/steamclient.so"
          ],
          "type": "command"
        }
      ]
    },
    "run": {
      "stop": "stop",
      "pre": [],
      "post": [],
      "arguments": [
      	"--game tf",
      	"-console",
        "+map ctf_2fort",
      	"-norestart"
      ],
      "program": "./srcds_run"
    },
    "data": {
      "map": {
      	"value": "ctf_2fort",
      	"required": true,
      	"desc": "TF2 Map",
      	"display": "Team Fortess 2 Map to load",
      	"internal": false
      }
    }
  }
}`
