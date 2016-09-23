angular.module('starter.controllers', [])

.controller('Channels', function($scope, $stateParams, $http) {

  var baseUrl = 'http://bb:8080'

  $scope.nowPlaying = 'off'
  $scope.channels = [
      {name: 'P1', suffix: 'p1'}
    , {name: 'Kuku', suffix: 'kuku'}
    , {name: 'Raadio 2', suffix: 'r2'}
    , {name: 'Bandit Rock', suffix: 'brock'}
  ]

  var setPlaying = function (name) {
    console.log('setPlaying', name)
    $scope.channels.map(function (c) {
      console.log('c is', c)
      if (c.suffix === name) {
        c.playing = true
      } else {
        c.playing = false
      }
    })
  }

  $scope.selectChannel = function (c) {
    console.log('selectChannel', c)
    $http.get(baseUrl + '/' + c.suffix).then(function (response) {
      setPlaying(response.data)
      console.log('success', response)
    }).catch(function (err) {
      setPlaying(undefined)
      console.log('error', err)
    })
  }

  $scope.getNowPlaying = function () {
    $http.get(baseUrl + '/playing').then(function (response) {
      console.log('success', response)
      setPlaying(response.data)
    }).catch(function (err) {
      console.log('error', err)
    })
  }

  $scope.getNowPlaying()
})
