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
	"strconv"
	"testing"
)

func TestProducer(t *testing.T) {
	producer := NewProducer([]string{"10.77.109.120:9092"})
	for i := 1; i < 10; i++ {
		err := producer.Send("test-topic", []byte("hello"+strconv.Itoa(i)))
		if err != nil {
			t.Fatal(err)
		}
	}
}
