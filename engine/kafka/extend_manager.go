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

package kafka

import (
	log "github.com/Sirupsen/logrus"

	"encoding/json"
	"strings"

	"github.com/weibocom/wqs/config"
	"github.com/weibocom/wqs/engine/zookeeper"
	"github.com/weibocom/wqs/model"
)

const (
	extengPath      = "/wqs"
	groupConfigPath = extengPath + "/groupconfig"
	queuePath       = extengPath + "/queue"

	emptyString = ""
)

type ExtendManager struct {
	config   *config.Config
	zkClient *zookeeper.ZkClient
}

func NewExtendManager(config *config.Config) *ExtendManager {
	extendManager := ExtendManager{}
	extendManager.config = config
	extendManager.zkClient = zookeeper.NewZkClient(strings.Split(config.ZookeeperAddr, ","))
	return &extendManager
}

//========extend配置相关函数========//

func (this *ExtendManager) AddGroupConfig(group string, queue string, write bool, read bool, url string, ips []string) bool {
	path := buildConfigPath(group, queue)
	groupConfig := model.GroupConfig{Group: group, Queue: queue, Write: write, Read: read, Url: url, Ips: ips}
	data := groupConfig.ToJson()
	log.Infof("add group config, zk path:%s, data:%s", path, data)
	return this.zkClient.CreateRec(path, data)
}

func (this *ExtendManager) DeleteGroupConfig(group string, queue string) bool {
	path := buildConfigPath(group, queue)
	log.Infof("delete group config, zk path:%s", path)
	return this.zkClient.DeleteRec(path)
}

func (this *ExtendManager) UpdateGroupConfig(group string, queue string, write bool, read bool, url string, ips []string) bool {
	path := buildConfigPath(group, queue)
	groupConfig := model.GroupConfig{Group: group, Queue: queue, Write: write, Read: read, Url: url, Ips: ips}
	data := groupConfig.ToJson()
	log.Infof("update group config, zk path:%s, data:%s", path, data)
	return this.zkClient.Set(path, data)
}

func (this *ExtendManager) GetGroupConfig(group string, queue string) *model.GroupConfig {
	path := buildConfigPath(group, queue)
	data, _ := this.zkClient.Get(path)
	if len(data) == 0 {
		log.Infof("get group config, zk path:%s, data:null", path)
		return nil
	} else {
		groupConfig := model.GroupConfig{}
		json.Unmarshal([]byte(data), &groupConfig)
		log.Infof("get group config, zk path:%s, data:%s", path, groupConfig.ToJson())
		return &groupConfig
	}
}

func (this *ExtendManager) GetAllGroupConfig() map[string]*model.GroupConfig {
	keys, _ := this.zkClient.Children(groupConfigPath)
	allGroupConfig := make(map[string]*model.GroupConfig)
	for _, key := range keys {
		data, _ := this.zkClient.Get(groupConfigPath + "/" + key)
		groupConfig := model.GroupConfig{}
		json.Unmarshal([]byte(data), &groupConfig)
		allGroupConfig[key] = &groupConfig
	}
	return allGroupConfig
}

func (this *ExtendManager) GetGroupMap() map[string][]string {
	groupmap := make(map[string][]string)
	keys, _ := this.zkClient.Children(groupConfigPath)
	for _, k := range keys {
		group := strings.Split(k, ".")[0]
		queues, ok := groupmap[group]
		if ok {
			queues = append(queues, strings.Split(k, ".")[1])
			groupmap[group] = queues
		} else {
			tempqueues := make([]string, 0)
			tempqueues = append(tempqueues, strings.Split(k, ".")[1])
			groupmap[group] = tempqueues
		}
	}
	return groupmap
}

func (this *ExtendManager) GetQueueMap() map[string][]string {
	queuemap := make(map[string][]string)
	keys, _ := this.zkClient.Children(groupConfigPath)
	queues, _ := this.zkClient.Children(queuePath)
	for _, k := range keys {
		queue := strings.Split(k, ".")[1]
		groups, ok := queuemap[queue]
		if ok {
			groups = append(groups, strings.Split(k, ".")[0])
			queuemap[queue] = groups
		} else {
			tempgroups := make([]string, 0)
			tempgroups = append(tempgroups, strings.Split(k, ".")[0])
			queuemap[queue] = tempgroups
		}
	}
	for _, queue := range queues {
		_, ok := queuemap[queue]
		if ok {
			continue
		} else {
			queuemap[queue] = make([]string, 0)
		}
	}
	return queuemap
}

func (this *ExtendManager) AddQueue(queue string) bool {
	path := buildQueuePath(queue)
	data := ""
	log.Infof("add queue, zk path:%s, data:%s", path, data)
	return this.zkClient.CreateRec(path, data)
}

func (this *ExtendManager) DelQueue(queue string) bool {
	path := buildQueuePath(queue)
	log.Infof("del queue, zk path:%s", path)
	return this.zkClient.DeleteRec(path)
}

func (this *ExtendManager) GetQueues() []string {
	queues, _ := this.zkClient.Children(queuePath)
	return queues
}

func buildConfigPath(group string, queue string) string {
	return groupConfigPath + "/" + group + "." + queue
}

func buildQueuePath(queue string) string {
	return queuePath + "/" + queue
}
