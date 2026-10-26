package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 int8 }
func genFn0() uint {
    var v5 int32 = 361
    v5 = (v5 * 0)
    var v6 uint32 = 65535
    _ = v5
    _ = v6
    return 18446744073709551615
}

func main() {
    v1 := []string{}
    var v2 float64 = 1.5
    var v3 float64 = -2.25
    var v4 int32 = 828
    for fi7 := 0; fi7 < 7; fi7++ {
        v8 := GenS1{F0: 23, F1: v4}
        _ = fi7
        _ = v8
    }
    v3 += float64((0.1 + 1e10))
    v9 := GenS1{F0: -345, F1: (v4 % ((0) | 1))}
    if false {
        v2 += 1e10
    }
    if true {
        v2 = ((0.0 - (((v2 * (float64(0.1) - 3.14159)) + v3) + (((float64(float32(-2.25)) / 1.5) * (v3 / float64((3.14159 / -2.25)))) * ((1e10 - v2) - float64(1.5))))) * v3)
    } else {
        v4 = 1
        fmt.Println(v2, v2, v2)
    }
    var v10 int = 877
    v11 := GenS2{F0: ""}
    if false {
        var v12 int32 = -985
        _ = v12
    }
    var v13 float64 = -2.25
    var v14 bool = true
    for fi15 := 0; fi15 < 14; fi15++ {
        for fi16 := 0; fi16 < 9; fi16++ {
            if (!false) {
                if (0.1 < (3.14159 / float64(1.5))) {
                    v1 = append(v1, "\n\t")
                    fmt.Println(v14, v2)
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                if (!false) {
                    _ = len(v1)
                    v3 += v13
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            }
            v10 += v10
            switch 32767 {
            case 9223372036854775807:
                var v17 int32 = (v4 | -2147483648)
                _ = v17
            default:
                v13 -= v13
            }
            _ = fi16
        }
        var v18 int64 = -973
        _ = fi15
        _ = v18
    }
    v19 := GenS3{F0: 0, F1: 127}
    var v20 int16 = -1
    var v21 int8 = 0
    v1 = append(v1, "a")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v9
}
