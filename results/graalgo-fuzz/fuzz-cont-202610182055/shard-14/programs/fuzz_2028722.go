package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 string; F1 uint8; F2 uint }
func genFn0(p0 int16, p1 string, p2 uint16) int8 {
    switch "" {
    case "foo bar":
        p0 = (p0 % ((p0) | 1))
    case "hello":
        v4 := GenS1{F0: 42, F1: 2147483648}
        _ = v4
    }
    var v5 bool = ((p2 & (p2 / (((p2 * (p2 % ((p2) | 1)))) | 1))) == p2)
    v6 := []int{-2147483648, 689}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -77
}

func main() {
    var v1 string = ("\n\t" + ("foo bar" + ("a" + ("a" + "a"))))
    v2 := []string{(v1 + v1), (v1 + ("\n\t" + "hello")), ("foo bar" + ((("" + v1) + "\n\t") + v1))}
    var v3 int64 = 610
    var v7 float64 = 0.0
    var v8 float64 = ((3.14159 * 0.1) * (v7 * v7))
    v9 := GenS1{F0: 32767, F1: 421}
    if false {
        v8 += v7
        if len(v2) > 0 {
            v2[0] = v1
        }
    }
    if (42 < 65535) {
        fmt.Println(v7)
        v7 -= v7
    }
    switch -1 {
    case 9223372036854775807:
        for fi10 := 0; fi10 < 9; fi10++ {
            v11 := GenS2{F0: v1, F1: 255, F2: 0}
            v2 = append(v2, "a")
            _ = fi10
            _ = v11
        }
    case -462:
        var v12 uint16 = 1
        _ = v12
    }
    v1 = v1
    for fi13 := 0; fi13 < 8; fi13++ {
        v1 = v1
        v14 := GenS2{F0: "", F1: 49, F2: 65535}
        var v15 int16 = 1000
        _ = fi13
        _ = v14
        _ = v15
    }
    v1 = ("a" + v1)
    if ("foo bar" != ("a" + v1)) {
        switch 127 {
        case 824:
            var v16 int = 4294967295
            _ = v16
        case 657:
            var v17 int8 = 1
            _ = v17
        default:
            for fi18 := 0; fi18 < 5; fi18++ {
                v3 += -376
                v19 := GenS1{F0: (fi18 ^ fi18), F1: 65535}
                _ = fi18
                _ = v19
            }
        }
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
