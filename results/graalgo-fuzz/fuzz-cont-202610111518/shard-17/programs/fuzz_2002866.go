package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 uint16; F2 uint64 }
func genFn0() int64 {
    v6 := []int{-26, 196, 42}
    var v7 uint32 = 4294967295
    v7 = v7
    for fi8 := 0; fi8 < 17; fi8++ {
        v6 = append(v6, (^(fi8)))
        var v9 float64 = 0.0
        _ = fi8
        _ = v9
    }
    var v10 uint32 = v7
    _ = v10
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 int16, p1 int32) string {
    p1 = p1
    p1 = 661
    v11 := GenS1{F0: (("a" + "\n\t") + "")}
    _ = p0
    _ = p1
    _ = v11
    return "\n\t"
}

func main() {
    var v1 int32 = -1
    var v2 uint = 0
    var v3 bool = (true || true)
    v4 := make(map[int]uint32)
    var v5 uint32 = 0
    if false {
        fmt.Println(v3, v2, v5)
        for fi12 := 0; fi12 < 7; fi12++ {
            v13 := GenS2{F0: -722, F1: 50, F2: 2}
            v3 = v3
            _ = fi12
            _ = v13
        }
    } else {
        var v14 uint16 = 32767
        var v15 int8 = -128
        v3 = v3
        _ = v14
        _ = v15
    }
    v3 = (-2147483648 > 2147483647)
    v1 = (-(v1))
    var v16 float64 = 0.0
    switch "foo bar" {
    case "":
        v5 = v5
    }
    if false {
        v5 = v5
    }
    for fi17 := 0; fi17 < 13; fi17++ {
        v18 := v4[(fi17 + fi17)]
        _ = fi17
        _ = v18
    }
    if v3 {
        if (v3 || (!false)) {
            delete(v4, 127)
            v19 := GenS1{F0: "foo bar"}
            _ = v19
        } else {
            var v20 float32 = 3.14159
            v21 := v4[32767]
            v1 = (v1 / ((v1) | 1))
            _ = v20
            _ = v21
        }
        var v22 uint32 = (v5 - (v5 / (((v5 / ((v5) | 1))) | 1)))
        var v23 bool = ("a" >= ("foo bar" + "\n\t"))
        _ = v22
        _ = v23
    }
    v2 -= v2
    v16 -= v16
    if v3 {
        v24 := GenS2{F0: 322, F1: 255, F2: 927}
        for fi25 := 0; fi25 < 14; fi25++ {
            v26 := GenS2{F0: 9223372036854775807, F1: 876, F2: 759}
            _ = fi25
            _ = v26
        }
        _ = v24
    } else {
        for fi27 := 0; fi27 < 16; fi27++ {
            var v28 uint64 = 59
            _ = fi27
            _ = v28
        }
        var v29 uint32 = (^(v5))
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
