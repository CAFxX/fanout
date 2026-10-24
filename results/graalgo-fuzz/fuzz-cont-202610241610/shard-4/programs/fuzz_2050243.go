package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 int }
func genFn0(p0 bool) int {
    fmt.Println(p0)
    for fi5 := 0; fi5 < 18; fi5++ {
        v6 := []int{(fi5 & -2147483648), (^(fi5))}
        for fi7 := 0; fi7 < 11; fi7++ {
            var v8 int64 = -9223372036854775808
            _ = len(v6)
            var v9 uint32 = 65535
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi5
        _ = v6
    }
    _ = p0
    return -625
}

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{}
    var v3 uint32 = 2
    var v4 int16 = 42
    var v10 uint = 383
    if ("foo bar" != "hello") {
        var v11 int32 = -423
        _ = v11
    }
    var v12 uint16 = 800
    v10 -= 255
    var v13 int16 = v4
    v1["foo bar"] = (((("a" + (((("hello" + "hello") + "hello") + "a") + "")) + "a") + (("a" + "hello") + "\n\t")) + "")
    v14 := GenS1{F0: 0, F1: 255}
    v12 = (v12 & v12)
    v10 += (v10 >> (((v10 >> ((v10) & 63))) & 63))
    switch -792 {
    case 81:
        switch int(v13) {
        case 128:
            _ = len(v2)
        case 1:
            for fi15 := 0; fi15 < 14; fi15++ {
                v13 -= (^(v4))
                _ = fi15
            }
        case 0:
            for fi16 := 0; fi16 < 4; fi16++ {
                var v17 bool = false
                _ = fi16
                _ = v17
            }
        default:
            fmt.Println(v12, v10, v13)
        }
    case -768:
        for fi18 := 0; fi18 < 12; fi18++ {
            if len(v2) > 0 {
                v2[0] = uint32(v10)
            }
            if (!true) {
                v4 += v4
                var v19 float32 = (0.0 * 1.5)
                if (false && false) {
                    fmt.Println(v19, v13)
                    var v20 uint64 = 919
                    for fi21 := 0; fi21 < 19; fi21++ {
                        var v22 uint32 = (v3 % ((v3) | 1))
                        _ = fi21
                        _ = v22
                    }
                    _ = v20
                }
                _ = v19
            } else {
                for fi23 := 0; fi23 < 11; fi23++ {
                    v24 := GenS2{F0: genFn0(true)}
                    var v25 uint8 = 42
                    if false {
                        v3 += (v3 << ((v10) & 63))
                        fmt.Println(v4)
                    }
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            }
            _ = fi18
        }
    case 3:
        v1["\n\t"] = "foo bar"
    }
    v26 := GenS1{F0: -958, F1: uint8(v3)}
    delete(v1, ("\n\t" + "foo bar"))
    var v27 uint8 = 255
    v3 = 2147483647
    _ = len(v2)
    fmt.Println(v4, v3, v3)
    v27 += (^(v27))
    if (false && (!(18446744073709551615 == 537))) {
        var v28 uint32 = (v3 * uint32(v27))
        var v29 float32 = (float32(float64(3.14159)) - 3.14159)
        v1["\n\t"] = ("" + ("" + "hello"))
        _ = v28
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
