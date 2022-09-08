package compress

import (
	"bytes"
	"repo-test-CICD-S3/pkg/compress"
	"testing"
)

func TestCompress(t *testing.T) {
	inputs := [][]byte{
		[]byte("05KNR_IDTick_Test                               4000E1301#0/T                                 9   000000000                0       0"),
		[]byte("09:00\t11:30\t12:30\t15:00\t\t\t+09:00\nDOP\tDOR\tDHP\tDHR\tDLP\tDLR\tDPP\tDPR\tDV\tQAP\tQAR\tQAS\tQBP\tQBR\tQBS\n1\t0900\t0\t3365\t\t3365\t\t3365\t\t3365\t\t32600\t\t\t\t\t\t\n2\t0901\t0\t3365\t\t3370\t\t3365\t\t3370\t\t1100"),
		[]byte("09:00\t11:30\t12:30\t15:00\t\t\t+09:00\nDOP\tDOR\tDHP\tDHR\tDLP\tDLR\tDPP\tDPR\tDV\tQAP\tQAR\tQAS\tQBP\tQBR\tQBS\n1\t0900\t0\t3365\t\t3365\t\t3365\t\t3365\t\t32600\t\t\t\t\t\t\n2\t0901\t0\t3365\t\t3370\t\t3365\t\t3370\t\t1100\t\t\t\t\t\t\n3\t0902\t0\t3370\t\t3370\t\t3370\t\t3370\t\t900\t\t\t\t\t\t\n4\t0903\t0\t3365\t\t3370\t\t3365\t\t3370\t\t2200\t\t\t\t\t\t\n5\t0904\t0\t3365\t\t3370\t\t3365\t\t3365\t\t4100\t\t\t\t\t\t\n6\t0905\t0\t3365\t\t3370\t\t3365\t\t3365\t\t2900\t\t\t\t\t\t\n7\t0907\t0\t3365\t\t3365\t\t3365\t\t3365\t\t200\t\t\t\t\t\t\n8\t0908\t0\t3370\t\t3370\t\t3370\t\t3370\t\t2400\t\t\t\t\t\t\n9\t0909\t0\t3365\t\t3370\t\t3365\t\t3365\t\t2000"),
		nil,
	}
	outputsExpect := []*bytes.Buffer{
		bytes.NewBuffer([]byte{120, 156, 50, 48, 245, 246, 11, 138, 247, 116, 9, 201, 76, 206, 142, 15, 73, 45, 46, 81, 192, 15, 76, 12, 12, 12, 92, 13, 141, 13, 12, 149, 13, 244, 67, 8, 168, 85, 80, 176, 84, 80, 80, 48, 128, 1, 116, 73, 152, 128, 1, 32, 0, 0, 255, 255, 109, 76, 21, 244}),
		bytes.NewBuffer([]byte{120, 156, 108, 141, 49, 170, 196, 48, 12, 68, 235, 231, 171, 252, 70, 178, 73, 62, 73, 103, 163, 34, 69, 32, 146, 2, 123, 255, 163, 44, 113, 181, 11, 219, 60, 134, 153, 7, 35, 219, 46, 130, 234, 222, 4, 173, 147, 203, 211, 192, 223, 156, 138, 93, 142, 93, 137, 29, 142, 29, 137, 157, 142, 157, 137, 185, 99, 158, 216, 139, 232, 78, 244, 36, 250, 77, 12, 39, 70, 18, 227, 46, 138, 108, 34, 8, 173, 173, 11, 191, 88, 215, 121, 5, 148, 250, 200, 250, 33, 255, 11, 95, 89, 85, 228, 29, 0, 0, 255, 255, 50, 193, 30, 147}),
		bytes.NewBuffer([]byte{120, 156, 124, 208, 65, 206, 131, 32, 16, 5, 224, 245, 243, 42, 255, 102, 24, 64, 197, 157, 102, 22, 46, 76, 28, 48, 249, 239, 127, 148, 102, 104, 130, 109, 211, 178, 121, 33, 240, 5, 30, 67, 105, 33, 130, 115, 139, 39, 56, 174, 25, 109, 7, 248, 171, 71, 131, 156, 10, 57, 11, 100, 87, 200, 94, 32, 135, 66, 142, 2, 81, 133, 104, 129, 252, 35, 175, 138, 188, 22, 228, 245, 66, 222, 20, 121, 43, 200, 219, 53, 56, 80, 34, 2, 193, 251, 49, 226, 91, 242, 88, 159, 2, 48, 176, 97, 247, 130, 39, 194, 219, 218, 185, 102, 189, 89, 174, 246, 169, 62, 51, 53, 26, 140, 250, 206, 181, 204, 205, 70, 179, 225, 135, 181, 12, 119, 133, 209, 108, 236, 88, 190, 59, 76, 102, 167, 206, 28, 238, 10, 179, 209, 185, 243, 51, 14, 205, 38, 179, 169, 87, 129, 136, 30, 1, 0, 0, 255, 255, 159, 228, 71, 38}),
		bytes.NewBuffer([]byte{}),
	}
	for idx := range inputs {
		output, err := compress.Compress(inputs[idx])
		if err != nil || !bytes.Equal(output.Bytes(), outputsExpect[idx].Bytes()) {
			t.Errorf("CASE %d: FAILED", idx)
		} else {
			t.Logf("CASE %d: SUCCESS", idx)
		}
	}
}