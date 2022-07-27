package app

import apicontent "github.com/henrikroschmann/F-api-ker/controller/api-content"

func mapUrls() {
	router.GET("/content", apicontent.Get)                   //Get all jsons
	router.GET("/content/:content_id", apicontent.GetById)   // Get specific jsons
	router.POST("/content", apicontent.Save)                 // Create jsons
	router.DELETE("/content/:content_id", apicontent.Delete) // Delete jsons
}
