package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
func genFn0(p0 int16, p1 int32) int32 {
    v6 := GenS1{F0: "", F1: 0}
    v7 := []int{-1, 52, 32768}
    p0 -= p0
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{"a", ("foo bar" + "\n\t"), "a", ("\n\t" + "\n\t")}
    v3 := make(map[int]int)
    var v4 int64 = 577
    v5 := make(map[int]string)
    if true {
        var v8 bool = (!(!true))
        _ = v8
    } else {
        fmt.Println(v4)
    }
    for fi9 := 0; fi9 < 9; fi9++ {
        v10 := GenS1{F0: (((("" + ("\n\t" + "")) + "\n\t") + "") + ""), F1: (v4 & 0)}
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 float32 = (0.0 * (3.14159 * 1.5))
        _ = fi11
        _ = v12
    }
    var v13 uint32 = 4294967295
    var v14 uint32 = 0
    v15 := v3[-9223372036854775808]
    if ("hello" == "a") {
        var v16 int = (-(v15))
        _ = v16
    } else {
        v14 += v13
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        var v18 uint8 = 104
        _ = fi17
        _ = v18
    }
    fmt.Println(v15)
    v14 -= (v13 - v14)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
