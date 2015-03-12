(function() {
	var app = angular.module('itDashboard', []);

	app.controller('itDashboardController', function($scope, itService) {
		$scope.services = [];

		loadRemoteData();

		function loadRemoteData() {
			itService.getServices().then(function(services){
				$scope.services = services
			})
		}
	})

	app.factory('itService', function($http){
		return {
			getServices: function(){
				return $http.jsonp('http://localhost:8080/services')
				.then(function(response){ 
					console.log(response);
					return response.data; 
				});
			}
		}
	})
})