package main

import "net/http"
import "gorm.io/gorm"


type AetherServer struct {
   
   DB *gorm.DB

}


(s AetherServer) DeleteConnectivityServices(w http.ResponseWriter, r *http.Request, target Target){


}
        // GET /connectivity-services Container
        // (GET /aether/v2.0.x/{target}/connectivity-services)
(s AetherServer) GetConnectivityServices(w http.ResponseWriter, r *http.Request, target Target){
 
          
        
         tx := s.DB.Create(t)
        return tx.Error

}
        // POST /connectivity-services
        // (POST /aether/v2.0.x/{target}/connectivity-services)
(s AetherServer) PostConnectivityServices(w http.ResponseWriter, r *http.Request, target Target){


}


