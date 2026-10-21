package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint16 }
type GenS2 struct { F0 uint8; F1 int8; F2 int8 }
func genFn0(p0 int8) string {
    fmt.Println(p0)
    var v4 int16 = 214
    v5 := []int{-9223372036854775808, -2147483648, 994}
    _ = p0
    _ = v4
    _ = v5
    return (((("a" + "") + "") + "") + "hello")
}
func genFn1(p0 int, p1 uint, p2 int) string {
    var v6 uint64 = 0
    for fi7 := 0; fi7 < 20; fi7++ {
        if false {
            switch p0 {
            case 1:
                if true {
                    var v8 float32 = ((-2.25 * 0.0) * 0.1)
                    p0 += p0
                    _ = v8
                }
            default:
                var v9 uint32 = 2147483648
                _ = v9
            }
            var v10 uint8 = 127
            v11 := []int{p0, -1, int(v10)}
            _ = v10
            _ = v11
        } else {
            v12 := GenS1{F0: 42, F1: -1, F2: 65535}
            var v13 int = fi7
            _ = v12
            _ = v13
        }
        v14 := []int{p2}
        v6 -= (^(v6))
        _ = fi7
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return ("foo bar" + "foo bar")
}

func main() {
    v1 := []int{32767, 9223372036854775807}
    v2 := make(map[uint32]int)
    v3 := []int{-565, 127}
    for fi15 := 0; fi15 < 4; fi15++ {
        var v16 uint16 = 32767
        _ = fi15
        _ = v16
    }
    _ = len(v2)
    v17 := v2[2147483648]
    v17 = (v17 * (v17 & 270))
    v17 -= (^(v17))
    v17 += (^(v17))
    _ = len(v2)
    switch v17 {
    case 42:
        var v18 int = -60
        _ = v18
    }
    v19 := GenS2{F0: 156, F1: 127, F2: 104}
    var v20 uint16 = 65535
    if len(v1) > 0 {
        v1[0] = (v17 & v17)
    }
    var v21 int64 = 1
    if true {
        switch "a" {
        case "a":
            v22 := GenS1{F0: 267, F1: -2147483648, F2: (^(v20))}
            _ = v22
        case "":
            v23 := GenS2{F0: 1, F1: 42, F2: -128}
            _ = v23
        case "foo bar":
            fmt.Println(v20, v21)
        }
        var v24 int32 = 0
        v20 = (v20 & v20)
        _ = v24
    }
    for fi25 := 0; fi25 < 19; fi25++ {
        var v26 string = "a"
        _ = fi25
        _ = v26
    }
    fmt.Println(v21, v17, v20)
    var v27 string = ("hello" + "")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v3
}
