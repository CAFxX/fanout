package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 int32; F1 uint32; F2 uint32 }
func genFn0(p0 string) int32 {
    p0 = p0
    var v5 uint16 = 2
    p0 = "\n\t"
    var v6 uint = 1
    _ = p0
    _ = v5
    _ = v6
    return 42
}
func genFn1() uint16 {
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := []int{(fi7 >> ((42) & 63)), 633, fi7}
        var v9 string = ("a" + "\n\t")
        v10 := GenS1{F0: (fi7 % ((fi7) | 1)), F1: 908}
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 uint32 = 2147483648
    v11 -= v11
    _ = v11
    return 255
}

func main() {
    var v1 int = -2147483648
    var v2 int = v1
    var v3 bool = true
    v4 := []uint32{1, 1, 4294967295, 739}
    if len(v4) > 0 {
        v4[0] = 444
    }
    switch (v1 % (((v1 + (v1 ^ (v1 / ((v2) | 1))))) | 1)) {
    case 437:
        var v12 int16 = 1
        _ = v12
    }
    var v13 int8 = -32
    var v14 uint = 4294967295
    _ = len(v4)
    var v15 uint16 = 42
    if ((v14 + v14) != 65535) {
        if v3 {
            for fi16 := 0; fi16 < 12; fi16++ {
                switch "" {
                case "a":
                    v3 = ((v13 + -128) > v13)
                default:
                    v14 = (v14 ^ 255)
                }
                _ = fi16
            }
        } else {
            var v17 uint32 = 2147483647
            _ = v17
        }
        switch ((("\n\t" + ("foo bar" + "\n\t")) + "hello") + (((("foo bar" + "a") + "foo bar") + ("" + ("hello" + "foo bar"))) + "\n\t")) {
        case "\n\t":
            for fi18 := 0; fi18 < 13; fi18++ {
                v14 -= (v14 / (((^(v14))) | 1))
                _ = fi18
            }
        case "hello":
            var v19 int = v1
            _ = v19
        case "foo bar":
            v20 := GenS1{F0: v1, F1: (^(v14))}
            _ = v20
        default:
            v15 += (v15 & 763)
        }
    }
    var v21 uint = (v14 | 758)
    v14 = 1
    if false {
        var v22 float64 = -2.25
        _ = v22
    }
    fmt.Println(v3, v21, v2)
    v23 := GenS2{F0: 0, F1: 2, F2: 2}
    switch "\n\t" {
    case "a":
        var v24 int = v1
        _ = v24
    case "hello":
        v1 = v2
    case "":
        fmt.Println(v13, v3)
    default:
        v25 := GenS1{F0: (v2 << ((654) & 63)), F1: v14}
        _ = v25
    }
    var v26 string = ""
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v23
    _ = v26
    _ = v3
    _ = v4
}
