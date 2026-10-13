package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int16; F1 uint32 }
func genFn0() uint16 {
    var v7 string = "hello"
    if false {
        fmt.Println(v7)
        fmt.Println(v7)
        v8 := GenS1{F0: 560, F1: v7}
        _ = v8
    } else {
        var v9 int = 32768
        fmt.Println(v9, v9)
        v9 += 0
        _ = v9
    }
    _ = v7
    return 65535
}
func genFn1(p0 uint32, p1 int, p2 int64) int64 {
    v10 := []int{(p1 % ((p1) | 1))}
    for fi11 := 0; fi11 < 13; fi11++ {
        switch "hello" {
        case "foo bar":
            p1 += 32767
        case "\n\t":
            switch "a" {
            case "\n\t":
                _ = len(v10)
            case "foo bar":
                p2 -= 1
            case "a":
                var v12 uint8 = 239
                _ = v12
            }
        case "":
            var v13 string = "hello"
            _ = v13
        }
        _ = fi11
    }
    fmt.Println(p1)
    _ = len(v10)
    if (-32768 >= -32768) {
        var v14 float64 = ((0.0 + (1.5 * 1.5)) + (1e10 * (3.14159 * 3.14159)))
        for fi15 := 0; fi15 < 2; fi15++ {
            v16 := GenS2{F0: 1000, F1: 827}
            _ = fi15
            _ = v16
        }
        p2 += p2
        _ = v14
    } else {
        for fi17 := 0; fi17 < 6; fi17++ {
            _ = len(v10)
            _ = fi17
        }
        if len(v10) > 0 {
            v10[0] = p1
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 0
}

func main() {
    var v1 int = 39
    var v2 int = 757
    v3 := make(map[string]bool)
    var v4 int32 = 138
    var v5 int16 = -946
    v6 := []int{0, v1, (v2 ^ v2), (v2 / ((v1) | 1))}
    v4 += v4
    var v18 int64 = 423
    v5 += v5
    fmt.Println(v4, v2, v4)
    v6 = append(v6, (v2 * (-(v2))))
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v19 := GenS2{F0: -1, F1: 499}
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
