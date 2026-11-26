package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 uint; F2 uint }
func genFn0(p0 bool, p1 uint16) int16 {
    var v6 uint = 379
    p0 = p0
    if (false && (!p0)) {
        var v7 uint16 = (p1 / (((^(p1))) | 1))
        v6 = v6
        _ = v7
    }
    v8 := []int{32768}
    switch -9223372036854775808 {
    case -9223372036854775808:
        p0 = p0
    case 608:
        fmt.Println(p1, p0, v6)
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v8
    return 1000
}

func main() {
    var v1 uint16 = 457
    v2 := []int{-938, 128, 236, 535}
    v3 := []string{((((("" + "a") + "a") + "\n\t") + ("hello" + ("hello" + "foo bar"))) + "foo bar"), "\n\t", "hello", ("foo bar" + "foo bar")}
    var v4 uint32 = 892
    var v5 int32 = -2147483648
    v2 = append(v2, -269)
    for fi9 := 0; fi9 < 19; fi9++ {
        v5 -= v5
        var v10 uint = 65535
        fmt.Println(v1, v1, v5)
        _ = fi9
        _ = v10
    }
    v5 = v5
    for fi11 := 0; fi11 < 1; fi11++ {
        switch "hello" {
        case "foo bar":
            _ = len(v2)
        }
        v4 += v4
        v4 = (v4 * (v4 | (v4 * (v4 << ((4294967295) & 63)))))
        _ = fi11
    }
    v12 := GenS1{F0: (v1 % ((713) | 1))}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v4 += v4
    fmt.Println(v1)
    var v13 uint8 = 1
    var v14 uint = 197
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2 = append(v2, 42)
    fmt.Println(v5, v4)
    var v15 uint16 = 493
    v5 = 0
    v16 := GenS1{F0: (v15 << (((v14 + v14)) & 63))}
    v17 := GenS2{F0: (!false), F1: (^(v14)), F2: v14}
    switch "\n\t" {
    case "\n\t":
        var v18 uint8 = v13
        _ = v18
    case "":
        fmt.Println(v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
