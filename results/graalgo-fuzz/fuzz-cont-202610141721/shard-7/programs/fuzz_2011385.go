package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string; F1 int8; F2 int }
type GenS4 struct { F0 int64; F1 uint8 }
func genFn0(p0 int32, p1 bool, p2 int32) uint8 {
    if true {
        p2 = p0
    }
    v6 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 254
}

func main() {
    var v1 int = -549
    var v2 float32 = (float32((float64((-2.25 * 3.14159)) * 1.5)) + float32(0.0))
    v3 := []uint32{713}
    var v4 int64 = 9223372036854775807
    v5 := []uint32{42, 339, 920, 65535}
    var v7 int32 = -2147483648
    if len(v5) > 0 {
        v5[0] = 1
    }
    if true {
        for fi8 := 0; fi8 < 14; fi8++ {
            v3 = append(v3, 363)
            _ = fi8
        }
    }
    var v9 int32 = v7
    var v10 int = v1
    switch 1 {
    case -293:
        for fi11 := 0; fi11 < 12; fi11++ {
            if len(v5) > 0 {
                v5[0] = 2
            }
            var v12 string = (("\n\t" + ("hello" + "a")) + "a")
            _ = fi11
            _ = v12
        }
    default:
        v1 += 128
    }
    if false {
        v13 := GenS1{F0: -1}
        _ = v13
    } else {
        for fi14 := 0; fi14 < 10; fi14++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi14
        }
        v1 += 4294967295
    }
    if false {
        v1 -= (^(v1))
    } else {
        fmt.Println(v10, v4)
    }
    v4 = 1
    for fi15 := 0; fi15 < 11; fi15++ {
        var v16 bool = (((v1 ^ (v1 << ((255) & 63))) >= (fi15 % (((^(v1))) | 1))) || false)
        var v17 int32 = v7
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 bool = (true || false)
    var v19 bool = true
    v20 := GenS1{F0: -1}
    if ("\n\t" >= "a") {
        v21 := GenS2{F0: (!v19)}
        v22 := GenS3{F0: "hello", F1: int8(v4), F2: v10}
        _ = v21
        _ = v22
    } else {
        v23 := GenS1{F0: -1}
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        switch "hello" {
        case "foo bar":
            v4 = (v4 % (((^(v4))) | 1))
        default:
            fmt.Println(v2, v4, v18)
        }
        _ = v23
    }
    v24 := GenS4{F0: 9223372036854775807, F1: 128}
    v3 = append(v3, 65535)
    if (!v18) {
        v4 = (v4 % ((-1) | 1))
    } else {
        if v18 {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v25 int = v1
        _ = v25
    }
    fmt.Println(v7, v9, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v9
}
