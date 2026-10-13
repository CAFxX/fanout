package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int32) string {
    p0 -= (^(p0))
    v4 := GenS1{F0: "a", F1: 127}
    v5 := GenS1{F0: "a", F1: 2}
    for fi6 := 0; fi6 < 1; fi6++ {
        p0 = p0
        p0 = 0
        _ = fi6
    }
    for fi7 := 0; fi7 < 7; fi7++ {
        v8 := []int{(fi7 * (fi7 >> ((4294967295) & 63))), -238, -1}
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = v4
    _ = v5
    return "a"
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint = 635
    v3 := []string{"a", "a", "hello"}
    var v9 int = 32767
    var v10 int8 = 0
    var v11 int16 = 42
    for fi12 := 0; fi12 < 5; fi12++ {
        if len(v3) > 0 {
            v3[0] = ("a" + "")
        }
        v13 := GenS2{F0: 138}
        var v14 int32 = -836
        _ = fi12
        _ = v13
        _ = v14
    }
    v11 += (v11 | 292)
    for fi15 := 0; fi15 < 12; fi15++ {
        fmt.Println(v2, v2, fi15)
        fmt.Println(v2)
        _ = fi15
    }
    if (float64(-2.25) <= 0.0) {
        var v16 bool = (!true)
        _ = v16
    } else {
        fmt.Println(v10, v9, v2)
    }
    var v17 uint32 = 65535
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v9
}
