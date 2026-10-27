package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 bool; F1 uint64; F2 uint64 }
func genFn0(p0 int16, p1 uint32, p2 int) uint32 {
    p1 -= (^(p1))
    p1 = p1
    var v6 int64 = -9223372036854775808
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 544
}
func genFn1(p0 int, p1 uint8) string {
    var v7 string = ("foo bar" + "foo bar")
    v8 := []int{p0, p0, -2147483648}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return (v7 + (v7 + v7))
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]bool)
    var v3 uint8 = 59
    var v4 int16 = -428
    v5 := []int{-2147483648, -1, 682, 0}
    var v9 uint = 0
    v10 := GenS1{F0: -1, F1: 42}
    for fi11 := 0; fi11 < 18; fi11++ {
        delete(v2, "hello")
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi11
    }
    var v12 float32 = 0.1
    v9 -= v9
    var v13 uint = v9
    for fi14 := 0; fi14 < 13; fi14++ {
        delete(v1, "foo bar")
        if true {
            var v15 uint8 = v3
            _ = v15
        }
        v9 -= v9
        _ = fi14
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if false {
        v16 := GenS1{F0: -128, F1: (v3 | (v3 & (v3 + v3)))}
        v9 += v13
        _ = v16
    } else {
        switch ("" + "hello") {
        case "hello":
            v2["hello"] = true
        case "a":
            v17 := GenS2{F0: (0 == -76), F1: 206, F2: 195}
            _ = v17
        case "\n\t":
            v4 += (v4 / ((v4) | 1))
        default:
            v3 += (^(v3))
        }
        switch -1 {
        case 127:
            if false {
                for fi18 := 0; fi18 < 7; fi18++ {
                    var v19 float32 = -2.25
                    fmt.Println(v13)
                    if true {
                        fmt.Println(v13)
                    } else {
                        var v20 int8 = 127
                        fmt.Println(v12)
                        _ = v20
                    }
                    _ = fi18
                    _ = v19
                }
                v21 := GenS2{F0: ((v3 ^ (^(v3))) >= v3), F1: 42, F2: 949}
                _ = v21
            }
        case 752:
            fmt.Println(v9, v9)
        case 32768:
            var v22 string = ("foo bar" + ("hello" + "hello"))
            _ = v22
        default:
            var v23 int64 = -1
            _ = v23
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
