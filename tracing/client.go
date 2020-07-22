

/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/SkyAPM/go2sky"
	httpPlugin "github.com/SkyAPM/go2sky/plugins/http"
	"github.com/SkyAPM/go2sky/reporter"
)

func main() {
	r, err := reporter.NewGRPCReporter("localhost:32699")
	if err != nil {
		log.Fatalf("new reporter error %v \n", err)
		return
	}

	tracer, err := go2sky.NewTracer("client", go2sky.WithReporter(r))
	if err != nil {
		log.Fatalf("create tracer error %v \n", err)
	}
	tracer.WaitUntilRegister()

	client, err := httpPlugin.NewClient(tracer)
	if err != nil {
		log.Fatalf("create client error %v \n", err)
	}

	// call server
	request, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:30575/productpage"), nil)
	if err != nil {
		log.Fatalf("unable to create http request: %+v\n", err)
	}
	for {
		log.Println("requesting....")
		res, err := client.Do(request)
		if err != nil {
			log.Fatalf("unable to do http request: %+v\n", err)
		}
		log.Printf("%+v", res)
		_ = res.Body.Close()
		time.Sleep(time.Second)
	}
}

