package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 int8; F1 int16; F2 int }
type GenS4 struct { F0 int64; F1 string; F2 uint32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int16 = 17
    var v3 uint32 = 68
    var v4 uint16 = 255
    v5 := GenS1{F0: ((951 != 128) || (!(0.1 != (-2.25 * -2.25)))), F1: 722}
    v2 -= (v2 ^ 1)
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 int16 = v2
        switch fi6 {
        case 32768:
            v7 = 32767
        }
        fmt.Println(v2, v4, v2)
        _ = fi6
        _ = v7
    }
    v3 = (v3 * 4294967295)
    v8 := GenS1{F0: (255 == 128), F1: 4294967295}
    if ((0.0 - 3.14159) <= 1.5) {
        var v9 uint16 = v4
        v4 += v9
        _ = v9
    }
    v10 := GenS2{F0: 42, F1: 0}
    v11 := GenS3{F0: -128, F1: (v2 * v2), F2: 0}
    fmt.Println(v3, v2, v2)
    if false {
        if false {
            var v12 int16 = 1
            for fi13 := 0; fi13 < 12; fi13++ {
                var v14 float32 = 1.5
                v15 := GenS4{F0: -529, F1: "\n\t", F2: (^(v3))}
                for fi16 := 0; fi16 < 1; fi16++ {
                    v2 = 1000
                    _ = fi16
                }
                _ = fi13
                _ = v14
                _ = v15
            }
            v17 := v1["\n\t"]
            _ = v12
            _ = v17
        }
    } else {
        var v18 uint32 = v3
        switch -2147483648 {
        case 32768:
            v2 += v2
        case 4294967295:
            var v19 uint32 = 0
            _ = v19
        default:
            if (0.0 > -2.25) {
                var v20 int8 = -1
                v1["hello"] = v18
                for fi21 := 0; fi21 < 1; fi21++ {
                    v22 := GenS3{F0: v20, F1: 1, F2: fi21}
                    v23 := GenS2{F0: 0, F1: 127}
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                _ = v20
            }
        }
        _ = v18
    }
    fmt.Println(v2)
    v3 -= 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
