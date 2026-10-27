package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 int; F2 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32; F1 uint; F2 bool }
func genFn0(p0 int) int64 {
    v5 := []int{(p0 - 365), (p0 << ((1) & 63))}
    p0 += (p0 * (p0 - (p0 << ((2) & 63))))
    _ = len(v5)
    if len(v5) > 0 {
        v5[0] = (p0 - p0)
    }
    if true {
        if true {
            fmt.Println(p0)
        }
        p0 += p0
    } else {
        var v6 int64 = 9223372036854775807
        fmt.Println(v6)
        switch "\n\t" {
        case "foo bar":
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        case "\n\t":
            v6 = v6
        case "":
            fmt.Println(v6)
        default:
            p0 += (p0 % ((p0) | 1))
        }
        _ = v6
    }
    _ = p0
    _ = v5
    return -340
}
func genFn1(p0 int, p1 uint, p2 uint16) uint8 {
    var v7 int = -1
    for fi8 := 0; fi8 < 1; fi8++ {
        p0 -= -992
        v7 = 0
        _ = fi8
    }
    v9 := GenS1{F0: 42}
    v10 := []int{42}
    p0 = (v7 * -391)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v9
    return 0
}

func main() {
    var v1 uint16 = 516
    var v2 int = 1
    var v3 bool = ((((0.1 * (float64(float32(1e10)) - (-2.25 - ((0.1 - 0.0) - ((1e10 - 0.0) + (1.5 * float64(0.1))))))) + float64(1e10)) - 1e10) >= -2.25)
    var v4 float64 = (1e10 + (-2.25 / (-2.25 * 1e10)))
    for fi11 := 0; fi11 < 16; fi11++ {
        switch ("a" + ((("" + "\n\t") + ("hello" + (("a" + ("hello" + "")) + ""))) + ((("hello" + "") + ("hello" + (("foo bar" + "") + "\n\t"))) + (("\n\t" + "") + ("a" + "a"))))) {
        case "\n\t":
            v1 = 42
        case "":
            v12 := GenS1{F0: 2147483647}
            _ = v12
        default:
            v1 -= 2
        }
        v3 = true
        v13 := GenS2{F0: 2, F1: (^(v2)), F2: 1}
        _ = fi11
        _ = v13
    }
    v4 += 1.5
    v3 = (v3 || true)
    var v14 uint64 = 18446744073709551615
    v15 := GenS3{F0: 0}
    var v16 float32 = (3.14159 / 3.14159)
    v17 := []int{626, (v2 >> ((18446744073709551615) & 63)), (v2 * v2), (^(v2))}
    v18 := GenS4{F0: 65535, F1: 4294967295, F2: ((2 < 65535) || v3)}
    fmt.Println(v14)
    v19 := GenS1{F0: 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
