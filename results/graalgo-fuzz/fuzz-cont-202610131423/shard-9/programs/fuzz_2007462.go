package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int64 }
type GenS2 struct { F0 bool; F1 string }
func genFn0(p0 int64, p1 bool) uint8 {
    v5 := []int{}
    p0 = p0
    switch 32768 {
    case 9223372036854775807:
        var v6 float64 = 0.1
        _ = v6
    case -2147483648:
        p0 = 879
    default:
        v7 := GenS1{F0: 42, F1: 127, F2: (p0 % ((595) | 1))}
        _ = v7
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 int16 = -32768
        _ = fi8
        _ = v9
    }
    if false {
        var v10 uint = 4294967295
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v5
    return 255
}

func main() {
    var v1 int8 = 0
    var v2 string = ("a" + "")
    var v3 int16 = 32767
    var v4 string = ("hello" + ((("foo bar" + ("\n\t" + v2)) + "foo bar") + (v2 + v2)))
    v11 := []int{-178}
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 bool = (204 > 32)
        switch 0 {
        case 9223372036854775807:
            v14 := GenS1{F0: 4294967295, F1: 135, F2: 1}
            _ = v14
        }
        var v15 float64 = 1.5
        _ = fi12
        _ = v13
        _ = v15
    }
    fmt.Println(v2, v1)
    v3 -= 1000
    v16 := GenS1{F0: 524, F1: 127, F2: 9223372036854775807}
    switch "\n\t" {
    case "":
        v3 -= 1000
    case "\n\t":
        var v17 int16 = (v3 * v3)
        _ = v17
    default:
        fmt.Println(v4)
    }
    v3 += v3
    var v18 int8 = (v1 ^ 42)
    switch 32768 {
    case 32768:
        if ((-2.25 + (1.5 / (3.14159 * float32(3.14159)))) == (float32(float64(1.5)) * float32(0.1))) {
            v19 := GenS2{F0: true, F1: "foo bar"}
            _ = v19
        }
    case 9223372036854775807:
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
    }
    var v20 uint = 669
    v20 -= 938
    var v21 uint = 168
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v16
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
