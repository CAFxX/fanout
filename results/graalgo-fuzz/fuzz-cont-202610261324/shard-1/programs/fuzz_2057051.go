package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int16 }
func genFn0(p0 int16, p1 int8, p2 uint) uint16 {
    v5 := GenS1{F0: (^(p2))}
    var v6 uint8 = 128
    var v7 string = ("hello" + "\n\t")
    if true {
        switch v7 {
        case "a":
            var v8 float64 = -2.25
            _ = v8
        case "hello":
            v6 = 2
        case "\n\t":
            v9 := GenS2{F0: 65535, F1: p0}
            _ = v9
        default:
            v10 := []int{32767}
            _ = v10
        }
    } else {
        switch ((v7 + "a") + v7) {
        case "":
            for fi11 := 0; fi11 < 8; fi11++ {
                var v12 int8 = p1
                var v13 int32 = 2147483647
                _ = fi11
                _ = v12
                _ = v13
            }
        case "foo bar":
            v14 := []int{101, -1, 675}
            _ = v14
        case "a":
            v15 := []int{32768, 765}
            _ = v15
        default:
            v16 := []int{}
            _ = v16
        }
    }
    var v17 int = int(v6)
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v5
    _ = v6
    _ = v7
    return 43
}
func genFn1(p0 uint8, p1 int32, p2 uint16) int16 {
    switch 1 {
    case 4294967295:
        switch "hello" {
        case "\n\t":
            var v18 int8 = 1
            _ = v18
        default:
            fmt.Println(p1)
        }
    case 633:
        var v19 int64 = 0
        _ = v19
    case 0:
        switch ("foo bar" + ("foo bar" + "a")) {
        case "foo bar":
            var v20 float64 = (((1.5 - 0.0) + (float64(float32(0.0)) / 1.5)) + 1e10)
            _ = v20
        }
    }
    p1 = (p1 % ((p1) | 1))
    var v21 int8 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    return -1
}

func main() {
    var v1 uint32 = 342
    var v2 uint32 = (v1 & 1)
    var v3 int32 = 1
    v4 := []int{}
    var v22 uint8 = 128
    v4 = append(v4, 42)
    for fi23 := 0; fi23 < 20; fi23++ {
        var v24 int16 = 1000
        v2 += v2
        if true {
            for fi25 := 0; fi25 < 15; fi25++ {
                v26 := GenS2{F0: 211, F1: genFn1(2, (v3 | 0), 32767)}
                var v27 string = ("hello" + "hello")
                _ = fi25
                _ = v26
                _ = v27
            }
        } else {
            var v28 string = "foo bar"
            _ = v28
        }
        _ = fi23
        _ = v24
    }
    if (false || true) {
        v4 = append(v4, 4294967295)
        var v29 uint16 = genFn0(-1, 127, 344)
        _ = v29
    } else {
        if (v2 == v2) {
            var v30 int64 = 705
            _ = v30
        }
    }
    v22 = (^(v22))
    _ = len(v4)
    v31 := GenS2{F0: 549, F1: -1}
    for fi32 := 0; fi32 < 3; fi32++ {
        v4 = append(v4, 127)
        _ = fi32
    }
    switch -380 {
    case -2147483648:
        fmt.Println(v3, v2)
    case -9223372036854775808:
        var v33 uint32 = v1
        _ = v33
    case 4294967295:
        v4 = append(v4, 127)
    default:
        var v34 int = 1
        _ = v34
    }
    v35 := GenS2{F0: 255, F1: 764}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v31
    _ = v35
    _ = v4
}
