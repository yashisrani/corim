package cmd

import "github.com/veraison/corim/comid"

var (
	// note: embedded CoSWIDs are not validated {0: h'5C57E8F446CD421B91C908CF93E13CFC', 1: [505(h'deadbeef')]}
	testCorimValid = comid.MustHexDecode(nil, "a200505c57e8f446cd421b91c908cf93e13cfc0181d901f944deadbeef")
	// {0: h'5C57E8F446CD421B91C908CF93E13CFC'}
	testCorimInvalid = comid.MustHexDecode(nil, "a100505c57e8f446cd421b91c908cf93e13cfc")
	testMetaInvalid  = []byte("{}")
	testMetaValid    = []byte(`{
		"signer": {
			"name": "ACME Ltd signing key",
			"uri": "https://acme.example"
		},
		"validity": {
			"not-before": "2021-12-31T00:00:00Z",
			"not-after": "2025-12-31T00:00:00Z"
		}
	}`)
	testECKey = []byte(`{
		"kty": "EC",
		"crv": "P-256",
		"x": "MKBCTNIcKUSDii11ySs3526iDZ8AiTo7Tu6KPAqv7D4",
		"y": "4Etl6SRW2YiLUrN5vfvVHuhp7x8PxltmWWlbbM4IFyM",
		"d": "870MB6gfuTJ4HtUnUvYMyJpr5eUZNP4Bk43bVdj3eAE",
		"use": "enc",
		"kid": "1"
	  }`)

	testSignedCorimValid = comid.MustHexDecode(nil, `
d284585da3012603746170706c69636174696f6e2f72696d2b63626f7208
5841a200a2007441434d45204c7464207369676e696e67206b657901d820
7468747470733a2f2f61636d652e6578616d706c6501a200c11a61ce4800
01c11a69546780a059061aa600505c57e8f446cd421b91c908cf93e13cfc
0183590216d901faa40065656e2d474201a10050366d0a0a598845ed8488
2f2a544f62420281a3006941434d45204c74642e01d8207468747470733a
2f2f61636d652e6578616d706c65028300010204a1028282a200a300d902
27582061636d652d696d706c656d656e746174696f6e2d69642d30303030
3030303031016441434d45026a526f616452756e6e657201d90226582101
ceebae7b8927a3227e5303cf5e0f1f7b34bb542ad7250ac03fbcde36ec2f
150881a100787c4d466b77457759484b6f5a497a6a3043415159494b6f5a
497a6a30444151634451674145466e3074616f41775233506d724b6b594c
74417344396f30354b534d366d6267664e436770754c306736567054486b
5a6c3733776b354244786f56376e2b4f656565306949716b5733484d5a54
334554696e694a64673d3d82a200a300d90227582061636d652d696d706c
656d656e746174696f6e2d69642d303030303030303031016441434d4502
6a526f616452756e6e657201d902265821014ca3e4f50bf248c39787020d
68ffd05c88767751bf2645ca923f57a98becd29681a100787c4d466b7745
7759484b6f5a497a6a3043415159494b6f5a497a6a304441516344516741
453656777165376879334f385970612b425545544c556a424e5533724558
565579743958485237484a574c473758544b51643969316b565258654250
444c466e66597275312f657578526e4a4d374839556f46444c64413d3d59
01a3d901faa40065656e2d474201a1005043bbe37f2e614b33aed353cff1
428b160281a3006941434d45204c74642e01d8207468747470733a2f2f61
636d652e6578616d706c65028300010204a1008182a100a300d902275820
61636d652d696d706c656d656e746174696f6e2d69642d30303030303030
3031016441434d45026a526f616452756e6e657283a200d90258a3016242
4c0465322e312e30055820acbb11c7e4da217205523ce4ce1a245ae1a239
ae3c6bfd9e7871f7e5d8bae86b01a102818201582087428fc522803d3106
5e7bce3cf03fe475096631e5e07bbd7a0fde60c4cf25c7a200d90258a301
6450526f540465312e332e35055820acbb11c7e4da217205523ce4ce1a24
5ae1a239ae3c6bfd9e7871f7e5d8bae86b01a10281820158200263829989
b6fd954f72baaf2fc64bc2e2f01d692d4de72986ea808f6e99813fa200d9
0258a3016441526f540465302e312e34055820acbb11c7e4da217205523c
e4ce1a245ae1a239ae3c6bfd9e7871f7e5d8bae86b01a1028182015820a3
a5e715f0cc574a73c3f9bebb6bc24f32ffd5b67b387244c2c909da779a14
7859017cd901f9a8007820636f6d2e61636d652e727264323031332d6365
2d7370312d76342d312d352d300c0001783041434d4520526f616472756e
6e6572204465746563746f72203230313320436f796f7465204564697469
6f6e205350310d65342e312e3505a5182b65747269616c182d6432303133
182f66636f796f7465183473526f616472756e6e6572204465746563746f
721836637370310282a3181f745468652041434d4520436f72706f726174
696f6e18206861636d652e636f6d1821820102a3181f75436f796f746520
53657276696365732c20496e632e18206c6d79636f796f74652e636f6d18
210404a21826781c7777772e676e752e6f72672f6c6963656e7365732f67
706c2e7478741828676c6963656e736506a110a318186a72726465746563
746f7218196d2570726f6772616d6461746125181aa111a318186e727264
65746563746f722e657865141a000820e80782015820a314fc2dc663ae7a
6b6bc6787594057396e6b3f569cd50fd5ddb4d1bbafd2b6a0281a200d820
784068747470733a2f2f706172656e742e6578616d706c652f72696d732f
63636233616138352d363162342d343066312d383438652d303261643665
3861323534620182015820e45b72f5c0c0b572db4d8d3ab7e97f368ff74e
62347a824decb67a84e5224d750382482b06010401a02064781c68747470
3a2f2f61726d2e636f6d2f696f742f70726f66696c652f3104a200c11a61
ce480001c11a695467800581a3006941434d45204c74642e01d8206c6163
6d652e6578616d706c6502810158400d35302dda8e1249e295154255370f
ade855043ae59c7d460b49415fddec40d70e22f21f98ac18bc69a18d1757
30f6d7c6c28fa0819aeded8653f69d71d60172
	`)
	testSignedCorimInvalid = comid.MustHexDecode(nil, `
d284585da3012603746170706c69636174696f6e2f72696d2b63626f7208
5841a200a2007441434d45204c7464207369676e696e67206b657901d820
7468747470733a2f2f61636d652e6578616d706c6501a200c11a61ce4800
01c11a69546780a041a044deadbeef
	`)
)
