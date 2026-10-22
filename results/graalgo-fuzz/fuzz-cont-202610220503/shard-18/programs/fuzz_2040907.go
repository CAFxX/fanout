package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32 }
func genFn0(p0 int32, p1 uint16, p2 string) uint8 {
    p0 -= 2147483647
    var v5 string = ((("foo bar" + "a") + "hello") + p2)
    switch v5 {
    case "\n\t":
        var v6 uint64 = 0
        _ = v6
    case "a":
        p1 = 65535
    case "foo bar":
        if (false && (!false)) {
            fmt.Println(p0, p0, p1)
        }
    default:
        fmt.Println(p0, p1)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 255
}
func genFn1(p0 uint64, p1 bool) int {
    v7 := GenS1{F0: 1000}
    for fi8 := 0; fi8 < 20; fi8++ {
        fmt.Println(p1, fi8)
        p0 += (p0 % (((^(p0))) | 1))
        for fi9 := 0; fi9 < 16; fi9++ {
            switch "\n\t" {
            case "foo bar":
                var v10 uint = 4294967295
                _ = v10
            }
            p1 = false
            fmt.Println(fi9, p0)
            v11 := GenS1{F0: -1}
            _ = fi9
            _ = v11
        }
        _ = fi8
    }
    p1 = (-2.25 <= ((0.0 * 1e10) - 1.5))
    _ = p0
    _ = p1
    _ = v7
    return -105
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{(("" + "") + ""), "a", "foo bar"}
    v3 := []int{127, -2147483648, -990, 32767}
    v4 := make(map[uint32]uint32)
    for fi12 := 0; fi12 < 16; fi12++ {
        v13 := GenS2{F0: -2147483648}
        var v14 uint64 = 42
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 uint = 42
    for fi16 := 0; fi16 < 13; fi16++ {
        v15 = (v15 * (v15 / ((v15) | 1)))
        _ = fi16
    }
    v15 -= (v15 % (((v15 - (v15 << ((v15) & 63)))) | 1))
    var v17 int = 2147483647
    v18 := GenS1{F0: 32767}
    v15 = v15
    v19 := GenS2{F0: -1}
    delete(v1, "a")
    for fi20 := 0; fi20 < 10; fi20++ {
        var v21 string = "foo bar"
        _ = fi20
        _ = v21
    }
    v15 = v15
    v22 := v4[2147483647]
    var v23 int8 = -128
    _ = len(v3)
    var v24 float64 = float64((-2.25 * -2.25))
    if true {
        switch (v17 >> (((^(v15))) & 63)) {
        case 619:
            switch "hello" {
            case "":
                v23 += 100
            case "a":
                var v25 int16 = 0
                _ = v25
            case "\n\t":
                v22 -= (v22 + (v22 % ((2147483647) | 1)))
            default:
                if (genFn0(-2147483648, 389, "") < 2) {
                    var v26 int64 = -694
                    _ = v26
                }
            }
        case 2147483647:
            v22 -= v22
        }
        v1[("a" + "a")] = v22
        fmt.Println(v22, v23, v15)
    }
    if true {
        v4[(v22 + v22)] = (^(v22))
        switch "\n\t" {
        case "hello":
            var v27 float64 = v24
            _ = v27
        default:
            if true {
                v28 := GenS1{F0: -170}
                v3 = append(v3, 127)
                v29 := GenS2{F0: 1}
                _ = v28
                _ = v29
            }
        }
        var v30 float32 = (1e10 - float32((v24 - ((0.0 - float64((((float32(-2.25) + 1.5) + float32(float64(1e10))) / 3.14159))) * v24))))
        _ = v30
    }
    fmt.Println(v23, v15)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
