//Count the number of offline services
function countOffline(services){
	var i = 0;
	for(s in services){
		for(t in services[s]){
			console.log(t);
			if(services[s][t].Status != true){
				i++;
			}
		}
	}
	return i;
}
//Count the number of online services
function countOnline(services){
	var i = 0;
	for(s in services){
		for(t in services[s]){
			if(services[s][t].Status == true){
				i++;
			}
		}
	}
	return i;
}
function itDashboardController($scope,$http){
	$scope.onlineCount = 0;
	$scope.offlineCount = 0;
	$scope.services=[];
	$http.get("http://localhost:8080/services")
	.success(function(response){$scope.services = response;
		$scope.onlineCount = countOnline($scope.services); 
		$scope.offlineCount = countOffline($scope.services);});
}

var app = angular.module('itDashboard',[]);

app.controller('itDashboardController', itDashboardController);