/*
Copyright 2009-2016 Weibo, Inc.

All files licensed under the Apache License, Version 2.0 (the "License");
you may not use these files except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package service

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/weibocom/wqs/config"
)

func TestQueueService(t *testing.T) {
	queueService := NewQueueService(config.NewConfig())

	queueService.DeleteQueue("test_queue")

	groupinfo1 := queueService.LookupGroup("test_group2")
	result3, _ := json.Marshal(groupinfo1)
	fmt.Println(string(result3))
}
