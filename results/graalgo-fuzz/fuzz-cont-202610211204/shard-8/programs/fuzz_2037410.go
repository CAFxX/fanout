package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 uint }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint8, p1 string) int16 {
    v7 := GenS1{F0: (true || (1000 <= -1)), F1: 1, F2: 880}
    v8 := GenS1{F0: true, F1: 903, F2: 255}
    var v9 uint16 = 42
    for fi10 := 0; fi10 < 7; fi10++ {
        v11 := GenS1{F0: true, F1: 2147483647, F2: 1}
        _ = fi10
        _ = v11
    }
    v12 := []int{-9223372036854775808, 127, -354, 0}
    _ = p0
    _ = p1
    _ = v12
    _ = v7
    _ = v8
    _ = v9
    return 1000
}

func main() {
    var v1 uint8 = 79
    var v2 uint8 = (v1 >> ((18446744073709551615) & 63))
    var v3 uint64 = 4294967295
    v4 := make(map[string]string)
    v5 := []uint32{65535, 902}
    var v6 float32 = 0.1
    switch "hello" {
    case "":
        v5 = append(v5, 175)
    case "a":
        v6 -= v6
    }
    v13 := GenS2{F0: 0}
    v14 := v4["\n\t"]
    v14 = "\n\t"
    for fi15 := 0; fi15 < 20; fi15++ {
        v2 += 128
        _ = fi15
    }
    fmt.Println(v3, v2, v2)
    for fi16 := 0; fi16 < 20; fi16++ {
        if false {
            switch (fi16 / (((fi16 ^ 770)) | 1)) {
            case 0:
                v3 -= (^(v3))
            default:
                switch fi16 {
                case 128:
                    v2 -= v1
                case 915:
                    var v17 float64 = float64(v6)
                    _ = v17
                case 0:
                    if (fi16 >= fi16) {
                        if len(v5) > 0 {
                            v5[0] = 4294967295
                        }
                        var v18 string = (v14 + "a")
                        _ = v18
                    }
                }
            }
        }
        var v19 uint = 255
        _ = fi16
        _ = v19
    }
    var v20 int64 = -9223372036854775808
    v3 = v3
    if (false && false) {
        for fi21 := 0; fi21 < 1; fi21++ {
            fmt.Println(v6, v1)
            _ = len(v4)
            v22 := GenS1{F0: true, F1: -307, F2: 256}
            _ = fi21
            _ = v22
        }
        _ = len(v4)
    }
    v3 += (v3 << ((51) & 63))
    var v23 int32 = int32(v1)
    if true {
        v6 = 0.1
    } else {
        var v24 uint32 = 326
        v25 := GenS2{F0: 255}
        var v26 int = 4294967295
        _ = v24
        _ = v25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
