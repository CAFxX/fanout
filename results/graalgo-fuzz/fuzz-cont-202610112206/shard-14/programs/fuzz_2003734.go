package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
func genFn0(p0 int16) int64 {
    if false {
        for fi7 := 0; fi7 < 16; fi7++ {
            var v8 int64 = -1
            _ = fi7
            _ = v8
        }
        for fi9 := 0; fi9 < 1; fi9++ {
            p0 = (-(p0))
            _ = fi9
        }
    }
    v10 := []int{366, 4294967295}
    _ = len(v10)
    switch 42 {
    case 128:
        _ = len(v10)
    case -599:
        p0 = (p0 % ((p0) | 1))
    }
    _ = p0
    _ = v10
    return 0
}

func main() {
    var v1 uint64 = 2
    var v2 uint = 4294967295
    var v3 int32 = -1
    v4 := []string{"hello", "", "hello", "a"}
    var v5 float32 = 1.5
    v6 := make(map[int]string)
    v3 = v3
    if false {
        v11 := GenS1{F0: v3, F1: (^(v1))}
        _ = v11
    } else {
        fmt.Println(v5)
        v5 = ((v5 * ((float32(float64((1.5 * v5))) - v5) - v5)) + v5)
    }
    fmt.Println(v3, v3, v2)
    v2 -= (^(v2))
    for fi12 := 0; fi12 < 12; fi12++ {
        _ = len(v4)
        var v13 float32 = (float32((3.14159 / 1.5)) / v5)
        v5 = v5
        _ = fi12
        _ = v13
    }
    var v14 float32 = (0.0 - 0.1)
    v15 := GenS1{F0: (v3 ^ (-(v3))), F1: (v1 + (v1 >> ((v2) & 63)))}
    v14 += (v14 + (-2.25 + v14))
    fmt.Println(v2)
    switch "" {
    case "foo bar":
        if len(v4) > 0 {
            v4[0] = "hello"
        }
    case "\n\t":
        var v16 uint16 = 0
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
