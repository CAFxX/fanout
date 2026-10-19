package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int8 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 int16 = -1
    var v2 float32 = (3.14159 - (0.1 / 3.14159))
    var v3 int64 = -44
    var v4 uint64 = 936
    var v5 int8 = 96
    var v6 float32 = (v2 - (-2.25 - v2))
    v5 = v5
    if false {
        v7 := GenS1{F0: 32767, F1: v1, F2: v5}
        fmt.Println(v3, v1, v2)
        v6 += v2
        _ = v7
    }
    fmt.Println(v4)
    if false {
        v8 := []int{32767, 32767, 128, -9223372036854775808}
        _ = v8
    }
    v9 := GenS1{F0: 65535, F1: (v1 - 32767), F2: v5}
    if true {
        v10 := GenS1{F0: 151, F1: v1, F2: (v5 >> ((0) & 63))}
        var v11 float64 = -2.25
        _ = v10
        _ = v11
    } else {
        v3 += 42
    }
    var v12 uint = 65535
    v13 := GenS2{F0: -20, F1: 255}
    fmt.Println(v1, v12, v1)
    v14 := GenS3{F0: 42}
    v12 += (v12 & v12)
    switch -454 {
    case 2147483647:
        v2 = (float32(0.0) + (v2 / v2))
    case 32767:
        v15 := []int{9223372036854775807, 0}
        _ = v15
    }
    v16 := []int{32768, 4294967295}
    if false {
        switch 9223372036854775807 {
        case 2147483647:
            if false {
                if len(v16) > 0 {
                    v16[0] = 42
                }
            } else {
                for fi17 := 0; fi17 < 8; fi17++ {
                    var v18 int16 = v1
                    _ = fi17
                    _ = v18
                }
            }
        case 128:
            v1 = 42
        case 166:
            if false {
                var v19 string = ""
                v1 += (v1 | v1)
                v5 -= 42
                _ = v19
            }
        default:
            var v20 uint16 = 338
            _ = v20
        }
        v5 += (v5 ^ v5)
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
    } else {
        v2 = 3.14159
        if true {
            var v21 string = ""
            _ = v21
        }
        v6 -= v6
    }
    v22 := GenS2{F0: v5, F1: 26}
    v4 = 2
    v6 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
