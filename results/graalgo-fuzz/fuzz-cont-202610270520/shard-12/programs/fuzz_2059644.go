package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int64; F1 int8 }
func genFn0(p0 bool) uint8 {
    v7 := []int{-9223372036854775808, 152, 409, -9223372036854775808}
    for fi8 := 0; fi8 < 19; fi8++ {
        for fi9 := 0; fi9 < 8; fi9++ {
            if p0 {
                v10 := GenS1{F0: 480, F1: 42}
                if p0 {
                    fmt.Println(fi9)
                    var v11 uint8 = 127
                    _ = v11
                }
                _ = v10
            } else {
                p0 = false
            }
            p0 = ((fi8 - fi8) <= -9223372036854775808)
            fmt.Println(fi8, p0)
            switch fi9 {
            case 701:
                p0 = p0
            }
            var v12 int64 = 0
            _ = fi9
            _ = v12
        }
        _ = fi8
    }
    var v13 float32 = 1e10
    for fi14 := 0; fi14 < 17; fi14++ {
        for fi15 := 0; fi15 < 1; fi15++ {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = fi15
        }
        _ = fi14
    }
    _ = len(v7)
    _ = p0
    _ = v13
    _ = v7
    return 1
}
func genFn1() uint {
    var v16 int32 = 42
    v17 := GenS1{F0: 1, F1: -2147483648}
    var v18 float32 = ((0.0 - float32(3.14159)) / 0.1)
    fmt.Println(v16, v16)
    _ = v16
    _ = v17
    _ = v18
    return 631
}

func main() {
    var v1 int16 = -433
    var v2 int16 = 739
    var v3 uint8 = 127
    v4 := []int{42, 32768, 127, 305}
    v5 := make(map[string]string)
    var v6 string = "foo bar"
    v5[v6] = ("hello" + (v6 + "hello"))
    v2 -= 42
    v6 = ("" + v6)
    v5[""] = "foo bar"
    fmt.Println(v3, v6, v3)
    v3 += v3
    v6 = ("a" + ("a" + v6))
    v1 = v2
    fmt.Println(v2)
    v3 -= v3
    var v19 uint32 = 42
    for fi20 := 0; fi20 < 17; fi20++ {
        if (false || true) {
            for fi21 := 0; fi21 < 14; fi21++ {
                fmt.Println(fi21, fi21, v2)
                _ = fi21
            }
            v22 := GenS2{F0: 9223372036854775807, F1: 106}
            var v23 uint8 = (v3 - v3)
            _ = v22
            _ = v23
        }
        _ = fi20
    }
    v24 := GenS2{F0: 670, F1: 1}
    var v25 uint = 1
    v2 += v1
    var v26 int32 = -380
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
