package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 int8 }
func genFn0(p0 uint16, p1 uint64) uint16 {
    v4 := []int{120}
    p1 = p1
    v4 = append(v4, 0)
    _ = p0
    _ = p1
    _ = v4
    return p0
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{0, -1, 32767, 42}
    var v3 int8 = 127
    v2 = append(v2, -250)
    fmt.Println(v3)
    v5 := GenS1{F0: 357, F1: -1, F2: v3}
    var v6 float64 = (0.0 / ((3.14159 - (-2.25 / 1.5)) + float64((((1e10 - (float32(3.14159) + 3.14159)) + 0.0) / 1.0))))
    if false {
        v7 := GenS1{F0: 2147483647, F1: 4294967295, F2: v3}
        for fi8 := 0; fi8 < 9; fi8++ {
            fmt.Println(v3, fi8)
            if (!false) {
                var v9 uint = 574
                v3 -= v3
                _ = v9
            } else {
                switch "a" {
                case "hello":
                    v6 += 3.14159
                }
            }
            delete(v1, 65535)
            _ = fi8
        }
        _ = v7
    } else {
        v3 += (v3 & (v3 >> ((719) & 63)))
        v6 = (v6 + v6)
        fmt.Println(v6, v6)
    }
    v2 = append(v2, 4294967295)
    v6 = v6
    switch "" {
    case "foo bar":
        switch "" {
        case "foo bar":
            fmt.Println(v6)
        }
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        v6 += float64(-2.25)
        fmt.Println(fi10)
        var v11 float64 = 0.0
        _ = fi10
        _ = v11
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
}
