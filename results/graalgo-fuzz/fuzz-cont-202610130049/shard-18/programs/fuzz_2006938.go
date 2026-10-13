package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8; F1 int }
func genFn0(p0 int, p1 int16, p2 int8) uint8 {
    v7 := GenS1{F0: false}
    for fi8 := 0; fi8 < 15; fi8++ {
        p0 = (p0 << ((0) & 63))
        var v9 int16 = 0
        for fi10 := 0; fi10 < 20; fi10++ {
            var v11 string = "foo bar"
            _ = fi10
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}
func genFn1(p0 int16, p1 int16, p2 int8) uint8 {
    fmt.Println(p2, p0, p0)
    switch "hello" {
    case "hello":
        v12 := GenS1{F0: (!(1 >= 2))}
        _ = v12
    case "a":
        fmt.Println(p0, p2, p0)
    default:
        v13 := GenS1{F0: true}
        _ = v13
    }
    p0 = (p0 << ((846) & 63))
    _ = p0
    _ = p1
    _ = p2
    return 128
}

func main() {
    var v1 int = 1
    v2 := make(map[int]int)
    var v3 float64 = ((1e10 * -2.25) / (-2.25 + 1e10))
    v4 := make(map[string]string)
    var v5 uint8 = 1
    var v6 int = v1
    v1 -= -788
    for fi14 := 0; fi14 < 9; fi14++ {
        if (!(true && false)) {
            v15 := GenS2{F0: v5}
            var v16 uint64 = 1
            _ = v15
            _ = v16
        }
        switch "" {
        case "hello":
            var v17 float32 = (0.0 / 1e10)
            _ = v17
        }
        _ = fi14
    }
    if true {
        v4[""] = ""
        var v18 int8 = -128
        v18 = (v18 + v18)
        _ = v18
    } else {
        var v19 int64 = -9223372036854775808
        var v20 uint64 = 42
        _ = v19
        _ = v20
    }
    var v21 int = 2147483647
    var v22 int8 = -128
    delete(v4, (("foo bar" + ("hello" + "hello")) + ("" + (("hello" + "") + ""))))
    fmt.Println(v6)
    for fi23 := 0; fi23 < 7; fi23++ {
        var v24 bool = ((((v22 * v22) > (v22 | v22)) && (((v22 * 100) >= v22) && false)) || false)
        v25 := GenS1{F0: (v24 && (2 >= 42))}
        for fi26 := 0; fi26 < 1; fi26++ {
            v5 = 42
            switch -1 {
            case 363:
                v27 := GenS2{F0: 0}
                _ = v27
            case 127:
                var v28 float64 = v3
                _ = v28
            default:
                var v29 uint32 = 1
                fmt.Println(fi26)
                if true {
                    fmt.Println(v5, v5)
                }
                v3 += (((float64((1e10 + (float32((v3 * (0.0 - ((v3 + 1.5) + 0.1)))) - 0.0))) * (1.5 - v3)) - (-2.25 + (((v3 - v3) / float64((float32(v3) * ((float32(3.14159) - (0.1 * (0.0 * float32(v3)))) + 0.1)))) - v3))) + 1e10)
                switch v6 {
                case 1:
                    v30 := GenS3{F0: (v5 / ((v5) | 1)), F1: (v1 ^ 2147483647)}
                    _ = v30
                }
                _ = v29
            }
            _ = fi26
        }
        _ = fi23
        _ = v24
        _ = v25
    }
    var v31 string = ("foo bar" + "foo bar")
    v32 := GenS2{F0: (v5 / ((127) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
