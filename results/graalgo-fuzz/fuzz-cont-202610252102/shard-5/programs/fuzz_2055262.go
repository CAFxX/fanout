package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 uint; F2 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint; F1 int16; F2 uint8 }
func genFn0(p0 uint8) uint64 {
    var v4 float64 = ((float64((((0.1 * ((3.14159 * -2.25) - (1.5 - 0.0))) + 0.1) - 3.14159)) / (0.1 / 1.5)) + 1e10)
    if false {
        v5 := GenS1{F0: 2}
        fmt.Println(v4, p0)
        v6 := GenS1{F0: 841}
        _ = v5
        _ = v6
    } else {
        p0 = (^(p0))
        p0 -= (p0 << ((401) & 63))
    }
    _ = p0
    _ = v4
    return 0
}

func main() {
    var v1 string = ("hello" + ("\n\t" + "hello"))
    var v2 float32 = 0.0
    var v3 int16 = 372
    if true {
        var v7 uint8 = 194
        v8 := GenS2{F0: 1, F1: 1, F2: -595}
        var v9 uint16 = 352
        _ = v7
        _ = v8
        _ = v9
    }
    v1 = v1
    v10 := []int{-282, 127, -1}
    v11 := GenS3{F0: -1}
    if (true || (!false)) {
        fmt.Println(v3)
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        if len(v10) > 0 {
            v10[0] = 651
        }
    }
    switch "a" {
    case "\n\t":
        var v12 uint64 = 982
        _ = v12
    case "a":
        for fi13 := 0; fi13 < 11; fi13++ {
            fmt.Println(v2)
            var v14 uint32 = 4294967295
            v1 = v1
            _ = fi13
            _ = v14
        }
    }
    v15 := GenS4{F0: 18446744073709551615, F1: (v3 | (v3 % ((1) | 1))), F2: 0}
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    switch -9223372036854775808 {
    case 509:
        for fi16 := 0; fi16 < 1; fi16++ {
            for fi17 := 0; fi17 < 3; fi17++ {
                var v18 float64 = 0.1
                _ = fi17
                _ = v18
            }
            _ = fi16
        }
    case 0:
        v3 = (v3 << ((65535) & 63))
    default:
        var v19 uint32 = 2
        _ = v19
    }
    if true {
        var v20 int16 = v3
        _ = v20
    } else {
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        fmt.Println(v2, v3)
    }
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    for fi21 := 0; fi21 < 4; fi21++ {
        v1 = v1
        var v22 float32 = v2
        _ = fi21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v3
}
