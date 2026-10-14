package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 int8; F1 string; F2 uint64 }

func main() {
    var v1 uint32 = 42
    v2 := []string{}
    v3 := make(map[int]string)
    var v4 uint8 = 0
    var v5 uint64 = 2
    var v6 int64 = 1
    v6 = (v6 | 42)
    var v7 uint32 = (v1 ^ (v1 / ((2) | 1)))
    var v8 uint = 255
    var v9 float32 = 1e10
    var v10 int8 = 0
    v7 = v7
    fmt.Println(v6, v9)
    if (((v6 ^ (v6 - (-(v6)))) > (^(v6))) && true) {
        var v11 float32 = (((3.14159 - 1e10) * v9) - v9)
        _ = v11
    } else {
        var v12 float32 = ((float32((((((((float64(v9) / 1e10) + (((-2.25 + 1e10) + 3.14159) * 3.14159)) - 1e10) * ((float64(3.14159) + ((float64(v9) + 0.1) - (3.14159 - -2.25))) * float64(v9))) * 1.5) + float64(0.0)) + 0.1)) / 0.1) * v9)
        if false {
            fmt.Println(v4, v5)
            for fi13 := 0; fi13 < 4; fi13++ {
                v12 -= (((v12 * (((float32(((1e10 / float64(1e10)) * (1.5 + float64(1e10)))) * (v9 * ((0.1 + (v9 * (v9 + float32((3.14159 + float64(0.1)))))) / v12))) * v12) / (v9 + -2.25))) + (v9 * ((v9 + 1.5) + v9))) - 0.0)
                v14 := GenS1{F0: (fi13 ^ (^(fi13))), F1: (v7 & (v7 ^ (v1 ^ (v1 >> ((uint(v5)) & 63)))))}
                v8 -= v8
                _ = fi13
                _ = v14
            }
        }
        _ = v12
    }
    v2 = append(v2, "a")
    for fi15 := 0; fi15 < 20; fi15++ {
        for fi16 := 0; fi16 < 6; fi16++ {
            if (false && true) {
                switch "" {
                case "foo bar":
                    v3[(fi15 | 128)] = ("hello" + "foo bar")
                }
                var v17 string = ""
                var v18 uint64 = 4294967295
                _ = v17
                _ = v18
            } else {
                if ((-(v10)) > -128) {
                    v4 = (v4 ^ (v4 + (^(v4))))
                    v19 := GenS2{F0: (v10 ^ v10), F1: "a", F2: v5}
                    var v20 int64 = (v6 | (v6 << ((v8) & 63)))
                    _ = v19
                    _ = v20
                }
                var v21 float64 = 1.5
                _ = v21
            }
            if false {
                switch (fi15 | (fi15 & fi15)) {
                case 1:
                    var v22 uint8 = (v4 % (((v4 + v4)) | 1))
                    _ = v22
                case 4294967295:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                default:
                    var v23 uint = v8
                    _ = v23
                }
                v24 := v3[(^(fi16))]
                _ = v24
            }
            var v25 bool = true
            _ = fi16
            _ = v25
        }
        v4 += v4
        fmt.Println(v9)
        _ = fi15
    }
    fmt.Println(v7)
    v7 += (v7 * v1)
    for fi26 := 0; fi26 < 16; fi26++ {
        delete(v3, -2147483648)
        fmt.Println(v9, v7, v5)
        _ = fi26
    }
    v9 -= float32((1e10 / 3.14159))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
