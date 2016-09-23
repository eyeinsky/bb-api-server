angular.module('starter.controllers', [])

.controller('Channels', function($scope, $stateParams, $http) {

  var baseUrl = 'http://bb:8080'

  $scope.channels = [
      {name: 'P1', suffix: 'p1'}
    , {name: 'Kuku', suffix: 'kuku'}
    , {name: 'Raadio 2', suffix: 'r2'}
  ]
  $scope.selectChannel = function (c) {
    console.log('selectChannel', c)
    $http({
      method: 'GET',
      url: baseUrl + '/' + c.suffix
    }).then(function (response) {
      console.log('success', response)
    }, function (response) {
      console.log('error', response)
    })
  }
  console.log('cl')
})
