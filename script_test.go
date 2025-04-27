package bitcoin

import "testing"

func TestGetScriptFromAddress(t *testing.T) {
	t.Parallel()

	tests := []struct {
		inputAddress   string
		inputNetwork   NetworkType
		expectedScript string
		expectedError  bool
	}{
		{inputAddress: "bc1qr8063yn4gk44elj8sy6zk59y32v5t9jwjyv080", inputNetwork: Mainnet, expectedScript: "001419dfa8927545ab5cfe4781342b50a48a9945964e", expectedError: false},
		{inputAddress: "bc1q7cyrfmck2ffu2ud3rn5l5a8yv6f0chkp0zpemf", inputNetwork: Mainnet, expectedScript: "0014f60834ef165253c571b11ce9fa74e46692fc5ec1", expectedError: false},
		{inputAddress: "1LR8Cotk1LDxsVRtqUs16QJ3ExPGwDfeMz", inputNetwork: Mainnet, expectedScript: "76a914d4fa62e0243e52eeddd60812c9cd421bd337356588ac", expectedError: false},
		{inputAddress: "", inputNetwork: Mainnet, expectedScript: "", expectedError: true},
	}

	for _, test := range tests {
		script, err := GetScriptFromAddress(test.inputAddress, test.inputNetwork)
		if err != nil && !test.expectedError {
			t.Fatalf("%s Failed: [%v] inputted and error not expected but got: %s", t.Name(), test.inputAddress, err.Error())
		} else if err == nil && test.expectedError {
			t.Fatalf("%s Failed: [%v] inputted and error expected but got nil", t.Name(), test.inputAddress)
		} else if test.expectedScript != script {
			t.Fatalf("%s Failed: [%v] inputted [%s] expected but failed comparison of scripts, got: %s", t.Name(), test.inputAddress, test.expectedScript, script)
		}
	}
}
