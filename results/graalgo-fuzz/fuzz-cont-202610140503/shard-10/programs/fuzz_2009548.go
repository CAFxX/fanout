package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int32 }
type GenS2 struct { F0 uint64 }
func genFn0() int16 {
    v4 := GenS1{F0: 1, F1: "", F2: 468}
    var v5 int16 = -974
    var v6 int8 = -78
    v5 += (v5 >> ((930) & 63))
    var v7 int64 = 42
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return v5
}

func main() {
    var v1 uint32 = 464
    var v2 float64 = (0.0 * float64((3.14159 / 1.0)))
    v3 := []string{"", "hello"}
    var v8 int32 = -1
    v8 = 1
    var v9 int8 = 1
    fmt.Println(v9, v2)
    if true {
        v9 = v9
        v10 := GenS1{F0: 0, F1: "foo bar", F2: (v8 * (v8 - (v8 >> ((18446744073709551615) & 63))))}
        v2 += ((1e10 * v2) + (v2 - 0.1))
        _ = v10
    }
    var v11 int64 = -188
    var v12 uint = 365
    var v13 int = 128
    if (false && true) {
        fmt.Println(v8, v2)
    }
    v13 = (v13 >> (((v12 / (((v12 + v12)) | 1))) & 63))
    if ((v9 ^ (v9 % (((v9 % ((-82) | 1))) | 1))) >= v9) {
        if (-138 < genFn0()) {
            v3 = append(v3, "")
        } else {
            var v14 uint = v12
            for fi15 := 0; fi15 < 16; fi15++ {
                var v16 uint8 = 184
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi15
                _ = v16
            }
            _ = v14
        }
        var v17 int8 = (^(v9))
        v2 -= (float64(1e10) / 1.5)
        _ = v17
    }
    v13 -= -535
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    v18 := GenS2{F0: 18446744073709551615}
    v19 := GenS2{F0: 679}
    v8 += (v8 << (((v12 + (v12 / ((255) | 1)))) & 63))
    v12 += v12
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
