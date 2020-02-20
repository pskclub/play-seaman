module golang-guideline

go 1.13

replace (
	github.com/Sirupsen/logrus v1.0.5 => github.com/sirupsen/logrus v1.0.5
	github.com/Sirupsen/logrus v1.3.0 => github.com/Sirupsen/logrus v1.0.6
	github.com/Sirupsen/logrus v1.4.0 => github.com/sirupsen/logrus v1.0.6
	github.com/smartystreets/goconvey v0.0.0-20190611202728-68dc04aab96a => github.com/smartystreets/goconvey v0.0.0-20181108003508-044398e4856c
)

require (
	bitbucket.org/3dsinteractive/pam4 v0.0.0-20200116042303-a8873762fc23
	bitbucket.org/3dsinteractive/seaman v0.0.0-20200217084055-8e7944fb66df
)
