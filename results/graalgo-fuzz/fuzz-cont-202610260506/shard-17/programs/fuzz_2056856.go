package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int16 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint) int32 {
    var v6 uint = p0
    v7 := GenS1{F0: 18446744073709551615, F1: (p0 << (((p0 | (p0 ^ (^(v6))))) & 63)), F2: 32767}
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 int32 = 2147483647
        fmt.Println(v9, fi8)
        v10 := []int{}
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 int64 = -1
    fmt.Println(p0, v6, v11)
    _ = p0
    _ = v11
    _ = v6
    _ = v7
    return 0
}

func main() {
    v1 := []string{"foo bar", ((("a" + "foo bar") + (((("hello" + (("foo bar" + "a") + "foo bar")) + "hello") + "") + ("hello" + ("" + "")))) + "")}
    var v2 float64 = (float64(-2.25) + (0.1 - (0.1 * 0.0)))
    var v3 int16 = 431
    var v4 string = "foo bar"
    var v5 uint32 = 2147483648
    _ = len(v1)
    var v12 bool = ("a" == v4)
    v4 = v4
    v2 += (0.0 / v2)
    var v13 int32 = 0
    v2 = 1e10
    var v14 uint = 2
    var v15 int8 = 0
    v15 += v15
    if len(v1) > 0 {
        v1[0] = ((v4 + ((v4 + "a") + (("hello" + ("hello" + "a")) + "a"))) + ((v4 + ("a" + "hello")) + "hello"))
    }
    switch 463 {
    case 42:
        v16 := GenS2{F0: 2}
        _ = v16
    case -916:
        v13 -= (-(v13))
    case 127:
        v17 := GenS1{F0: 0, F1: v14, F2: v3}
        _ = v17
    default:
        if len(v1) > 0 {
            v1[0] = v4
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
