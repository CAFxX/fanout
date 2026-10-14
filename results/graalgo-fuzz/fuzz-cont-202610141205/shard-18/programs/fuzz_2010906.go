package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 uint8; F2 int }
type GenS3 struct { F0 bool }
func genFn0(p0 int64) int {
    fmt.Println(p0)
    v6 := []int{}
    if len(v6) > 0 {
        v6[0] = 127
    }
    var v7 bool = false
    if (v7 || (v7 && true)) {
        for fi8 := 0; fi8 < 9; fi8++ {
            fmt.Println(fi8, v7)
            _ = len(v6)
            _ = fi8
        }
    }
    _ = p0
    _ = v6
    _ = v7
    return -2147483648
}
func genFn1(p0 int64) int64 {
    fmt.Println(p0)
    v9 := []int{42}
    if false {
        v9 = append(v9, 1)
        v10 := GenS1{F0: 65535}
        switch ("a" + "a") {
        case "hello":
            fmt.Println(p0)
        }
        _ = v10
    } else {
        if (65535 < 597) {
            p0 += 511
            var v11 uint32 = 65535
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            _ = v11
        }
        switch 1 {
        case 32768:
            v12 := GenS1{F0: 442}
            _ = v12
        default:
            fmt.Println(p0)
        }
        var v13 int64 = (p0 % ((p0) | 1))
        _ = v13
    }
    var v14 int32 = -2147483648
    _ = p0
    _ = v14
    _ = v9
    return 9223372036854775807
}

func main() {
    var v1 int8 = 42
    var v2 int32 = 826
    v3 := []string{}
    var v4 float64 = ((-2.25 - (1e10 + (0.0 / 1e10))) - (-2.25 - 1e10))
    var v5 uint32 = 4294967295
    v4 -= (1.5 * (v4 + v4))
    v15 := GenS2{F0: "foo bar", F1: 221, F2: 0}
    v5 = 65535
    v4 = float64((float32(0.1) / float32(v4)))
    v16 := GenS3{F0: (0.0 > 0.0)}
    v1 += (v1 & -114)
    v3 = append(v3, ("\n\t" + (("hello" + "hello") + ("a" + "hello"))))
    var v17 int = 32767
    v17 -= genFn0(42)
    for fi18 := 0; fi18 < 2; fi18++ {
        fmt.Println(v1, v17, v2)
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
