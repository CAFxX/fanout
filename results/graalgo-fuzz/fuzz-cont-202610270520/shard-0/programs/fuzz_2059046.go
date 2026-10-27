package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int8 }
func genFn0(p0 uint) uint16 {
    if true {
        if (false || ("a" == (("\n\t" + "hello") + ("a" + ("hello" + "\n\t"))))) {
            var v4 int16 = -1
            p0 = (p0 | p0)
            v4 += (v4 + 594)
            _ = v4
        }
        v5 := []int{1, -613, 32768, 11}
        if len(v5) > 0 {
            v5[0] = -1
        }
        _ = v5
    }
    v6 := GenS1{F0: 9223372036854775807, F1: -27, F2: 127}
    if true {
        if true {
            v7 := GenS1{F0: 9223372036854775807, F1: 42, F2: -1}
            p0 += 4294967295
            _ = v7
        }
    }
    fmt.Println(p0)
    _ = p0
    _ = v6
    return 732
}

func main() {
    var v1 int32 = -748
    var v2 int = -965
    var v3 float64 = 1e10
    v2 = (v2 << ((65535) & 63))
    v2 = (-(v2))
    var v8 string = ("" + "hello")
    var v9 int = (v2 / (((v2 / ((v2) | 1))) | 1))
    v10 := []int{v9, v2, 499, v2}
    var v11 int16 = -32768
    if (v11 == (-(v11))) {
        v9 -= (v2 - (v2 >> ((1) & 63)))
        for fi12 := 0; fi12 < 18; fi12++ {
            fmt.Println(v11, v9, v11)
            _ = fi12
        }
        var v13 int8 = 1
        _ = v13
    }
    v11 = (v11 >> ((983) & 63))
    if false {
        var v14 uint = 939
        var v15 string = ((("" + "\n\t") + v8) + ("foo bar" + ("hello" + v8)))
        _ = v14
        _ = v15
    } else {
        var v16 uint = 992
        _ = v16
    }
    v17 := GenS1{F0: -1, F1: (-(v1)), F2: -1}
    fmt.Println(v8)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
