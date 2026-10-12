package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 string; F2 int8 }
type GenS3 struct { F0 int64 }
func genFn0() uint8 {
    v4 := GenS1{F0: 0}
    switch 127 {
    case -9223372036854775808:
        var v5 uint = 492
        _ = v5
    case 437:
        if false {
            var v6 float32 = ((0.1 / 3.14159) + (0.0 - 1.5))
            _ = v6
        }
    case 128:
        v7 := GenS2{F0: 0, F1: ((((((("a" + "a") + "\n\t") + "foo bar") + "a") + "\n\t") + ("a" + "foo bar")) + ("hello" + "foo bar")), F2: 127}
        _ = v7
    default:
        var v8 bool = true
        _ = v8
    }
    var v9 uint64 = 18446744073709551615
    _ = v4
    _ = v9
    return 128
}
func genFn1(p0 int64) int {
    fmt.Println(p0)
    v10 := GenS2{F0: 0, F1: "", F2: 96}
    if (("" + ("" + ("\n\t" + "a"))) <= ("a" + "hello")) {
        v11 := GenS2{F0: 42, F1: "hello", F2: -128}
        _ = v11
    }
    var v12 uint32 = 154
    switch -1 {
    case -722:
        p0 = p0
    case -1:
        for fi13 := 0; fi13 < 18; fi13++ {
            v14 := []int{fi13}
            _ = fi13
            _ = v14
        }
    default:
        fmt.Println(p0, p0)
    }
    _ = p0
    _ = v10
    _ = v12
    return 827
}

func main() {
    var v1 float64 = (1.5 + 3.14159)
    var v2 float32 = (0.1 * 3.14159)
    var v3 int16 = 0
    var v15 int = 4294967295
    v2 -= (((v2 - -2.25) / v2) + ((float32(v1) - ((v2 * (3.14159 * v2)) - v2)) / (1e10 / (v2 / v2))))
    fmt.Println(v3, v3)
    v15 += (^(v15))
    fmt.Println(v3)
    switch 42 {
    case 32768:
        for fi16 := 0; fi16 < 14; fi16++ {
            fmt.Println(v1, fi16)
            switch "\n\t" {
            case "\n\t":
                var v17 int = fi16
                _ = v17
            case "hello":
                v18 := GenS3{F0: 42}
                _ = v18
            case "a":
                fmt.Println(v2, v3)
            }
            _ = fi16
        }
    default:
        v1 -= v1
    }
    var v19 uint64 = 2
    if false {
        v19 -= (v19 - 42)
        for fi20 := 0; fi20 < 11; fi20++ {
            v21 := []int{9223372036854775807, 2147483647, v15}
            v3 = -1
            _ = fi20
            _ = v21
        }
    }
    switch (-(v15)) {
    case 1:
        var v22 int8 = 1
        _ = v22
    case 42:
        v1 = 3.14159
    case 32768:
        v23 := []int{(v15 << ((353) & 63)), genFn1(963)}
        _ = v23
    default:
        switch (^(v15)) {
        case -215:
            fmt.Println(v3, v2, v15)
        case 128:
            var v24 int = (v15 / (((v15 / ((v15) | 1))) | 1))
            _ = v24
        }
    }
    v25 := GenS3{F0: 982}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v25
    _ = v3
}
