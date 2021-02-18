module github.com/IBM-Cloud/terraform-provider-ibm

go 1.12

require (
	github.com/Bowery/prompt v0.0.0-20190916142128-fa8279994f75 // indirect
	github.com/IBM-Cloud/bluemix-go v0.0.0-20210203095940-db28d5e07b55
	github.com/IBM-Cloud/power-go-client v1.0.55
	github.com/IBM/apigateway-go-sdk v0.0.0-20200414212859-416e5948678a
	github.com/IBM/go-sdk-core v1.1.0
	github.com/IBM/go-sdk-core/v3 v3.3.1
	github.com/IBM/go-sdk-core/v4 v4.10.0
	github.com/IBM/ibm-cos-sdk-go v1.3.1
	github.com/IBM/ibm-cos-sdk-go-config v1.0.1
	github.com/IBM/keyprotect-go-client v0.5.2
	github.com/IBM/networking-go-sdk v0.12.1
	github.com/IBM/platform-services-go-sdk v0.17.8
	github.com/IBM/vpc-go-sdk v0.4.0
	github.com/ScaleFT/sshkeys v0.0.0-20200327173127-6142f742bca5
	github.com/Shopify/sarama v1.27.2
	github.com/apache/incubator-openwhisk-client-go v0.0.0-20171128215515-ad814bc98c32 // indirect
	github.com/apache/openwhisk-client-go v0.0.0-20200201143223-a804fb82d105
	github.com/apparentlymart/go-cidr v1.1.0
	github.com/cloudfoundry/jibber_jabber v0.0.0-20151120183258-bcc4c8345a21 // indirect
	github.com/dchest/safefile v0.0.0-20151022103144-855e8d98f185 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/ghodss/yaml v1.0.0
	github.com/go-openapi/validate v0.20.1 // indirect
	github.com/go-test/deep v1.0.4 // indirect
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect
	github.com/google/uuid v1.1.1
	github.com/hashicorp/go-uuid v1.0.2
	github.com/hashicorp/go-version v1.2.1
	github.com/hashicorp/hil v0.0.0-20200423225030-a18a1cd20038 // indirect
	github.com/hashicorp/terraform v0.12.28 // indirect
	github.com/hashicorp/terraform-plugin-sdk v1.6.0
	github.com/hokaccha/go-prettyjson v0.0.0-20170213120834-e6b9231a2b1c // indirect
	github.com/kardianos/govendor v1.0.9 // indirect
	github.com/minsikl/netscaler-nitro-go v0.0.0-20170827154432-5b14ce3643e3
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/gox v1.0.1 // indirect
	github.com/renier/xmlrpc v0.0.0-20170708154548-ce4a1a486c03 // indirect
	github.com/satori/go.uuid v1.2.0
	github.com/softlayer/softlayer-go v0.0.0-20190814165317-b9062a914a22
	github.ibm.com/ibmcloud/namespace-go-sdk v0.0.0-00010101000000-000000000000
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad
	golang.org/x/tools v0.0.0-20210107193943-4ed967dd8eff // indirect
	gotest.tools v2.2.0+incompatible
)

replace github.com/softlayer/softlayer-go v0.0.0-20190814165317-b9062a914a22 => ./common/github.com/softlayer/softlayer-go

replace github.ibm.com/ibmcloud/namespace-go-sdk => ./common/github.ibm.com/ibmcloud/namespace-go-sdk
