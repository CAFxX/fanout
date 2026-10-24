package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 string }
func genFn0(p0 int8, p1 uint16, p2 int64) string {
    var v6 int64 = p2
    v7 := []int{32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return ("foo bar" + "foo bar")
}

func main() {
    v1 := []string{"\n\t"}
    var v2 int8 = -120
    var v3 bool = (2147483648 > 78)
    v4 := []uint32{2147483648, 2}
    v5 := make(map[uint32]string)
    v8 := GenS1{F0: "hello"}
    if (960 <= 1) {
        if (0.0 < (1e10 * (1e10 + (1e10 - 1.5)))) {
            for fi9 := 0; fi9 < 4; fi9++ {
                switch (fi9 << ((4294967295) & 63)) {
                case 456:
                    if (!v3) {
                        var v10 uint64 = 1
                        _ = v10
                    }
                case -9223372036854775808:
                    _ = len(v5)
                case 32767:
                    switch "a" {
                    case "hello":
                        if false {
                            var v11 int32 = -1
                            _ = v11
                        } else {
                            v2 = (v2 >> ((78) & 63))
                        }
                    }
                default:
                    if ((1e10 - 1e10) == 0.0) {
                        _ = len(v1)
                        _ = len(v5)
                    } else {
                        v2 = -83
                    }
                }
                if (42 > 42) {
                    fmt.Println(v3, v2, v2)
                    v2 = (v2 % (((v2 & v2)) | 1))
                }
                _ = fi9
            }
        }
        for fi12 := 0; fi12 < 4; fi12++ {
            for fi13 := 0; fi13 < 5; fi13++ {
                var v14 int16 = 1
                v14 = v14
                v15 := GenS2{F0: -1, F1: "\n\t"}
                _ = fi13
                _ = v14
                _ = v15
            }
            v4 = append(v4, 427)
            delete(v5, 810)
            _ = fi12
        }
    }
    var v16 int64 = -1
    var v17 float32 = 3.14159
    switch 4294967295 {
    case -2147483648:
        v3 = (0 != -1)
    }
    var v18 string = "foo bar"
    v16 = v16
    v16 -= v16
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
