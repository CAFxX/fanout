package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 uint32 }
type GenS2 struct { F0 string }

func main() {
    var v1 uint32 = 2147483648
    var v2 int = 251
    v3 := []int{(^(v2)), v2, v2, v2}
    v4 := make(map[string]uint32)
    var v5 uint64 = 358
    var v6 float32 = 1.5
    for fi7 := 0; fi7 < 8; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            var v9 int = (^(fi7))
            var v10 uint64 = (v5 - 42)
            if true {
                v11 := GenS1{F0: (!(true && ((v1 == 708) || (-2147483648 < -2147483648)))), F1: 0, F2: (v1 / (((v1 % (((v1 | 2147483648)) | 1))) | 1))}
                _ = v11
            }
            _ = fi8
            _ = v10
            _ = v9
        }
        var v12 int32 = 0
        var v13 string = ""
        _ = fi7
        _ = v12
        _ = v13
    }
    if true {
        var v14 string = ((("\n\t" + "foo bar") + "") + (("" + ("a" + "")) + "foo bar"))
        for fi15 := 0; fi15 < 14; fi15++ {
            v2 += fi15
            _ = fi15
        }
        _ = v14
    }
    var v16 float32 = float32((3.14159 * 0.1))
    if true {
        v17 := GenS2{F0: "hello"}
        _ = v17
    } else {
        switch v2 {
        case 128:
            v18 := GenS2{F0: "foo bar"}
            _ = v18
        case -9223372036854775808:
            var v19 int32 = 1
            _ = v19
        case -241:
            for fi20 := 0; fi20 < 1; fi20++ {
                for fi21 := 0; fi21 < 5; fi21++ {
                    var v22 int16 = 764
                    _ = fi21
                    _ = v22
                }
                _ = fi20
            }
        }
    }
    v1 -= v1
    v6 += float32(float64(v6))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
