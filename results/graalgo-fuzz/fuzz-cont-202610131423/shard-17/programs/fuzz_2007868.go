package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 int64; F2 bool }
type GenS3 struct { F0 uint16; F1 uint16; F2 int }
func genFn0(p0 bool, p1 int32) int16 {
    var v7 uint32 = 0
    p0 = (85 == 32767)
    _ = p0
    _ = p1
    _ = v7
    return -890
}

func main() {
    var v1 uint32 = 2147483648
    v2 := []uint32{}
    var v3 uint32 = (v1 << ((114) & 63))
    var v4 float64 = (0.0 / (0.1 * -2.25))
    var v5 uint32 = (v1 + v1)
    v6 := []string{("" + "a"), "hello"}
    if true {
        v6 = append(v6, "")
        v6 = append(v6, (("a" + "\n\t") + "a"))
        fmt.Println(v3)
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v2 = append(v2, 2147483648)
    var v8 uint64 = 0
    v3 = 783
    var v9 float64 = -2.25
    if true {
        switch -1 {
        case -375:
            fmt.Println(v8, v3)
        case 9223372036854775807:
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        case -531:
            switch 32767 {
            case -772:
                if true {
                    switch ("" + "\n\t") {
                    case "a":
                        fmt.Println(v4)
                    case "\n\t":
                        var v10 bool = true
                        _ = v10
                    }
                    fmt.Println(v1, v5)
                } else {
                    var v11 float32 = (1.5 * 1e10)
                    _ = v11
                }
            case -568:
                fmt.Println(v4, v3, v1)
            default:
                v12 := GenS1{F0: 0}
                _ = v12
            }
        }
        for fi13 := 0; fi13 < 2; fi13++ {
            var v14 int32 = 2147483647
            _ = fi13
            _ = v14
        }
        if true {
            if ((false || true) || (v8 < v8)) {
                v15 := GenS2{F0: 199, F1: 9223372036854775807, F2: true}
                _ = v15
            }
            if false {
                _ = len(v2)
                for fi16 := 0; fi16 < 15; fi16++ {
                    for fi17 := 0; fi17 < 6; fi17++ {
                        v1 = 2147483647
                        v1 += 1
                        _ = fi17
                    }
                    _ = fi16
                }
                var v18 int16 = -5
                _ = v18
            }
        } else {
            var v19 uint64 = (v8 & v8)
            v4 = ((v9 / v4) + (-2.25 * v9))
            v2 = append(v2, v5)
            _ = v19
        }
    }
    v4 -= float64(float32(float64(1e10)))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if ((0.1 > (((-2.25 / 0.1) - (((((1.5 * 3.14159) / 0.1) - 0.1) - 1e10) / (0.1 + ((0.1 - 3.14159) + -2.25)))) / 1.5)) && true) {
        v2 = append(v2, (^(v3)))
        fmt.Println(v3)
        switch "\n\t" {
        case "foo bar":
            switch ((("hello" + "a") + "a") + "") {
            case "foo bar":
                fmt.Println(v1)
            }
        case "a":
            v20 := GenS2{F0: 0, F1: -374, F2: true}
            _ = v20
        }
    } else {
        v4 -= (float64((1e10 / -2.25)) / v9)
        v21 := GenS3{F0: 0, F1: uint16(v5), F2: 238}
        _ = v21
    }
    _ = len(v6)
    var v22 int = 2147483647
    for fi23 := 0; fi23 < 17; fi23++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi23
    }
    for fi24 := 0; fi24 < 4; fi24++ {
        v5 = 42
        v3 = v3
        _ = fi24
    }
    var v25 uint32 = (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
