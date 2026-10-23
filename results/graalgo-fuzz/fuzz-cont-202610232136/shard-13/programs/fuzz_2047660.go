package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint32 }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 bool; F1 bool }
type GenS5 struct { F0 uint16; F1 uint64; F2 uint16 }

func main() {
    var v1 uint16 = 42
    v2 := []string{"a", ""}
    var v3 uint64 = 809
    if true {
        v4 := GenS1{F0: 42, F1: -805, F2: 0}
        _ = v4
    }
    switch "a" {
    case "a":
        fmt.Println(v3, v1)
    case "hello":
        if false {
            for fi5 := 0; fi5 < 1; fi5++ {
                v3 = v3
                _ = fi5
            }
            for fi6 := 0; fi6 < 15; fi6++ {
                var v7 int = -2147483648
                if (!false) {
                    v1 -= 1
                    if len(v2) > 0 {
                        v2[0] = "hello"
                    }
                }
                switch 587 {
                case 4294967295:
                    for fi8 := 0; fi8 < 2; fi8++ {
                        v3 -= v3
                        fmt.Println(v1, fi6, fi8)
                        _ = fi8
                    }
                case 127:
                    switch "" {
                    case "":
                        v9 := GenS2{F0: 32767, F1: (^(v1))}
                        _ = v9
                    }
                default:
                    var v10 string = "\n\t"
                    _ = v10
                }
                _ = fi6
                _ = v7
            }
        }
    case "foo bar":
        v1 -= v1
    default:
        var v11 uint16 = (v1 | (^(v1)))
        _ = v11
    }
    switch ("" + "a") {
    case "a":
        if (42 != 42) {
            v12 := GenS3{F0: 42, F1: 128}
            v1 = (v1 - 535)
            v3 += v3
            _ = v12
        } else {
            var v13 uint64 = 0
            var v14 int16 = 32767
            v15 := GenS3{F0: -396, F1: 216}
            _ = v13
            _ = v14
            _ = v15
        }
    case "hello":
        v3 = (v3 >> ((55) & 63))
    }
    var v16 int32 = -1
    switch 0 {
    case -2147483648:
        v17 := GenS2{F0: -1, F1: (v1 << ((4294967295) & 63))}
        _ = v17
    case 458:
        v18 := GenS4{F0: true, F1: (((0.1 - (0.0 / 3.14159)) != 1e10) || (!false))}
        _ = v18
    case 540:
        fmt.Println(v16, v3)
    default:
        var v19 float32 = (3.14159 / 1.0)
        _ = v19
    }
    var v20 int32 = (v16 & (v16 - (v16 << ((939) & 63))))
    if true {
        v3 = 305
        v3 -= 273
        var v21 int16 = -32768
        _ = v21
    } else {
        if ((3.14159 * (float64(3.14159) - 3.14159)) <= 0.0) {
            var v22 bool = false
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            fmt.Println(v1, v20, v20)
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 8; fi23++ {
        v3 = 191
        v3 = (^(v3))
        _ = fi23
    }
    fmt.Println(v20, v3)
    v24 := GenS5{F0: v1, F1: (^(v3)), F2: v1}
    var v25 int64 = 9223372036854775807
    switch -9223372036854775808 {
    case -1:
        v3 -= v3
    }
    v26 := GenS3{F0: v16, F1: 128}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
