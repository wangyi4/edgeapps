#!/bin/bash
#########################################################
# Copyright 2019 Intel Corporation. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#########################################################

while :
do
  taskset -c 1 ffmpeg -re -async 1 -vsync -1 -f mjpeg -r 30 -i vidfifo.mjpeg \
    -vcodec mjpeg -b:v 50M -s 1280x720 -c:v copy \
    -f rtp rtp://analytics.community.appliance.mec:5001 > \
    /dev/null 2>&1 < /dev/null
done
