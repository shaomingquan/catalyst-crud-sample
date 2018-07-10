package main

import "github.com/shaomingquan/webcore-sample/imports"

func init() {
	go func() {
		<- prepare // wait for prepare
		app.Init()

		
		imports.Start_(&app)
		
		imports.Start_api(&app)
		
		imports.Start_api_article(&app)
		
		imports.Start_task(&app)
		// ###
	
	
	
	

		loaded <- 1 // i am loaded
	}()

}

// auto generate by _, dont modify
