package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 int64 }
func genFn0(p0 int16, p1 uint8) string {
    for fi7 := 0; fi7 < 4; fi7++ {
        for fi8 := 0; fi8 < 16; fi8++ {
            v9 := GenS1{F0: 1}
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        p0 -= (-(p0))
        p0 -= p0
        p1 = p1
        _ = fi10
    }
    if false {
        v11 := GenS2{F0: p1}
        _ = v11
    }
    fmt.Println(p0, p1)
    switch int(p1) {
    case -702:
        v12 := GenS3{F0: 127, F1: 9223372036854775807}
        _ = v12
    case 51:
        var v13 uint16 = 32767
        _ = v13
    case -1:
        v14 := []int{}
        _ = v14
    }
    _ = p0
    _ = p1
    return ("\n\t" + "")
}

func main() {
    var v1 uint32 = 2147483647
    var v2 int32 = -2147483648
    var v3 string = ("" + "a")
    var v4 int64 = -9223372036854775808
    var v5 float32 = 1.5
    v6 := []string{v3}
    var v15 int = 128
    var v16 float64 = float64(-2.25)
    var v17 int32 = int32(v15)
    fmt.Println(v16)
    if (v3 != (v3 + v3)) {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        var v18 string = v3
        _ = v18
    } else {
        if len(v6) > 0 {
            v6[0] = v3
        }
        v19 := GenS3{F0: 0, F1: 9223372036854775807}
        fmt.Println(v1, v2)
        _ = v19
    }
    if true {
        switch "" {
        case "a":
            var v20 int = (v15 << ((1) & 63))
            _ = v20
        case "foo bar":
            v3 = v3
        }
    } else {
        v1 -= 1
        var v21 int = (v15 & v15)
        v2 += v2
        _ = v21
    }
    var v22 uint32 = (^(v1))
    fmt.Println(v1, v16)
    for fi23 := 0; fi23 < 12; fi23++ {
        if (-2.25 != 3.14159) {
            var v24 bool = (v3 < v3)
            _ = v24
        } else {
            var v25 uint8 = 0
            _ = v25
        }
        if false {
            if true {
                fmt.Println(v4, v17, v5)
            }
            var v26 uint = 0
            _ = v26
        }
        v6 = append(v6, (v3 + v3))
        _ = fi23
    }
    var v27 int32 = (v17 ^ v2)
    v15 -= -1
    var v28 float64 = 1e10
    fmt.Println(v1, v27)
    switch "\n\t" {
    case "hello":
        if ((0.1 + v28) >= ((v28 / ((v28 / v16) * (0.1 / 1.0))) * 0.0)) {
            var v29 float64 = float64(-2.25)
            _ = v29
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
