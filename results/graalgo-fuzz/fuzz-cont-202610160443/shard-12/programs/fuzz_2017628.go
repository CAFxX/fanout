package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 uint32; F1 uint32; F2 int8 }
func genFn0(p0 bool, p1 bool) uint64 {
    v5 := []int{128, 1, 82}
    var v6 int64 = -9223372036854775808
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 853
}

func main() {
    var v1 string = "a"
    var v2 int = -754
    v3 := []uint32{728, 42}
    var v4 uint = 956
    v1 = v1
    var v7 uint16 = 1
    v8 := GenS1{F0: 9223372036854775807}
    var v9 float64 = (float64(float32(-2.25)) + 3.14159)
    fmt.Println(v9, v1)
    fmt.Println(v1, v4, v4)
    var v10 uint64 = 1
    switch v1 {
    case "a":
        var v11 uint16 = (v7 << ((v4) & 63))
        _ = v11
    case "foo bar":
        fmt.Println(v2, v7, v1)
    case "\n\t":
        var v12 uint8 = 42
        _ = v12
    }
    if (true || true) {
        fmt.Println(v9)
    } else {
        if false {
            for fi13 := 0; fi13 < 4; fi13++ {
                var v14 int64 = 1
                switch (v1 + "a") {
                case "":
                    for fi15 := 0; fi15 < 2; fi15++ {
                        var v16 uint16 = v7
                        var v17 float64 = -2.25
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                case "a":
                    fmt.Println(v10)
                }
                _ = fi13
                _ = v14
            }
            var v18 float32 = 0.1
            _ = v18
        } else {
            if false {
                v19 := GenS2{F0: (v7 >> (((v4 >> ((v4) & 63))) & 63)), F1: 9223372036854775807}
                var v20 uint32 = 2147483648
                v21 := GenS2{F0: (v7 + v7), F1: 9223372036854775807}
                _ = v19
                _ = v20
                _ = v21
            } else {
                if true {
                    v9 = ((float64(0.0) / (v9 / (v9 - ((-2.25 - 1e10) / (0.1 + (1.5 - (0.1 / 1.0))))))) / 3.14159)
                } else {
                    v2 -= (-(v2))
                    v3 = append(v3, 308)
                }
                fmt.Println(v2, v2, v9)
            }
        }
    }
    v10 += (v10 >> ((161) & 63))
    v2 = (v2 & (v2 * (v2 % ((v2) | 1))))
    var v22 uint16 = (v7 >> ((v4) & 63))
    fmt.Println(v9, v2)
    v4 += (v4 + (^(v4)))
    if false {
        v2 = v2
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi23 := 0; fi23 < 3; fi23++ {
            switch fi23 {
            case 587:
                if len(v3) > 0 {
                    v3[0] = 398
                }
            case 4294967295:
                v24 := GenS3{F0: uint32(v10), F1: 488, F2: 1}
                _ = v24
            case -322:
                v9 = (v9 - 1.5)
            }
            var v25 uint16 = v22
            fmt.Println(v25, v2, v1)
            _ = fi23
            _ = v25
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
