package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 int }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 string; F1 uint; F2 uint64 }

func main() {
    v1 := []int{2147483647, 896, 32767}
    var v2 uint8 = 255
    var v3 uint64 = 847
    v4 := []uint32{2147483647, 450, 4294967295, 401}
    v5 := make(map[int]int)
    var v6 int16 = 0
    switch ("foo bar" + "a") {
    case "\n\t":
        for fi7 := 0; fi7 < 12; fi7++ {
            v6 = v6
            _ = fi7
        }
    case "":
        delete(v5, 128)
    default:
        switch ("" + "a") {
        case "foo bar":
            var v8 uint32 = 1
            _ = v8
        default:
            v9 := GenS1{F0: 1, F1: ((1.5 * (-2.25 + (1e10 / float64(-2.25)))) < (-2.25 - 0.0)), F2: 461}
            _ = v9
        }
    }
    v3 = (v3 % ((v3) | 1))
    v10 := GenS2{F0: uint16(v3), F1: 624}
    v11 := v5[624]
    v12 := GenS3{F0: "", F1: 0, F2: 18446744073709551615}
    v2 = v2
    switch v11 {
    case -2147483648:
        if (!(0 == 82)) {
            switch 652 {
            case 9223372036854775807:
                for fi13 := 0; fi13 < 3; fi13++ {
                    for fi14 := 0; fi14 < 4; fi14++ {
                        var v15 int = (fi13 / ((v11) | 1))
                        var v16 int = v15
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                    var v17 uint64 = (v3 ^ (v3 / (((^(v3))) | 1)))
                    _ = fi13
                    _ = v17
                }
            default:
                v6 = (^(v6))
            }
            v6 += 1000
            v18 := GenS1{F0: 1, F1: false, F2: v11}
            _ = v18
        } else {
            for fi19 := 0; fi19 < 1; fi19++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                if (121 < uint(v2)) {
                    fmt.Println(v11, v2, fi19)
                    var v20 int64 = 42
                    _ = v20
                }
                _ = fi19
            }
        }
    case 2147483647:
        if true {
            v2 = (v2 * v2)
            var v21 float32 = (-2.25 / float32(((1.5 - 3.14159) + 0.1)))
            _ = v21
        } else {
            if len(v1) > 0 {
                v1[0] = (v11 % ((0) | 1))
            }
            v22 := GenS2{F0: 255, F1: 0}
            v23 := GenS2{F0: 2, F1: 65535}
            _ = v22
            _ = v23
        }
    case 876:
        for fi24 := 0; fi24 < 8; fi24++ {
            var v25 int32 = 0
            delete(v5, v11)
            _ = len(v5)
            _ = fi24
            _ = v25
        }
    }
    fmt.Println(v2, v6)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
