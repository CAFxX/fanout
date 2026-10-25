package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
func genFn0(p0 int32) int {
    p0 -= p0
    var v6 int32 = p0
    var v7 int = 377
    _ = p0
    _ = v6
    _ = v7
    return (v7 << ((457) & 63))
}

func main() {
    var v1 float64 = ((0.0 / ((1.5 + (3.14159 + 0.0)) - -2.25)) + 0.1)
    var v2 bool = false
    var v3 int64 = 9223372036854775807
    v4 := make(map[int]string)
    var v5 string = "a"
    v5 = (("" + "a") + v5)
    if v2 {
        var v8 bool = (v2 || (float32((v1 - (3.14159 / v1))) > 0.0))
        _ = v8
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        if true {
            _ = len(v4)
            var v10 float32 = -2.25
            _ = v10
        }
        _ = fi9
    }
    if (2147483647 == 1) {
        v3 += 0
        for fi11 := 0; fi11 < 13; fi11++ {
            v12 := v4[(fi11 - fi11)]
            var v13 uint8 = 1
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    fmt.Println(v5, v2)
    var v14 int8 = 42
    var v15 int32 = 202
    if (((v14 % (((v14 ^ 1)) | 1)) == -90) || v2) {
        for fi16 := 0; fi16 < 1; fi16++ {
            var v17 float32 = -2.25
            for fi18 := 0; fi18 < 17; fi18++ {
                fmt.Println(v14)
                _ = fi18
            }
            _ = fi16
            _ = v17
        }
        for fi19 := 0; fi19 < 8; fi19++ {
            v3 = (v3 - 9223372036854775807)
            v4[(fi19 ^ -1)] = ""
            fmt.Println(v2)
            _ = fi19
        }
        fmt.Println(v2)
    } else {
        v2 = true
        v5 = v5
    }
    v5 = v5
    v20 := GenS1{F0: 1, F1: (v15 - (v15 >> ((18446744073709551615) & 63)))}
    v21 := GenS1{F0: 127, F1: -2147483648}
    fmt.Println(v1, v14, v2)
    v15 += (v15 | v15)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
