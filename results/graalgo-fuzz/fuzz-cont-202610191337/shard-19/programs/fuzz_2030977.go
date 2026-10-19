package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 bool }
func genFn0(p0 string, p1 uint16) uint16 {
    fmt.Println(p1, p1)
    fmt.Println(p1)
    p1 -= (p1 / ((p1) | 1))
    _ = p0
    _ = p1
    return 65535
}

func main() {
    v1 := []string{"hello"}
    v2 := make(map[string]string)
    var v3 uint16 = 940
    for fi4 := 0; fi4 < 17; fi4++ {
        var v5 uint16 = v3
        v6 := GenS1{F0: 2147483648, F1: 478}
        _ = fi4
        _ = v5
        _ = v6
    }
    if len(v1) > 0 {
        v1[0] = "a"
    }
    v3 -= v3
    var v7 uint64 = 1
    fmt.Println(v3, v3)
    fmt.Println(v7, v7)
    _ = len(v1)
    if false {
        v8 := v2["a"]
        for fi9 := 0; fi9 < 11; fi9++ {
            v10 := v2[v8]
            for fi11 := 0; fi11 < 18; fi11++ {
                var v12 float32 = 1.5
                if (!(!(0 >= 898))) {
                    v2["foo bar"] = v10
                    var v13 uint16 = 0
                    _ = v13
                } else {
                    v8 = v10
                }
                var v14 int = -779
                _ = fi11
                _ = v12
                _ = v14
            }
            v15 := GenS1{F0: 42, F1: 287}
            _ = fi9
            _ = v10
            _ = v15
        }
        v1 = append(v1, v8)
        _ = v8
    } else {
        var v16 float64 = 1e10
        _ = v16
    }
    v3 += (v3 >> ((255) & 63))
    var v17 string = ""
    switch v17 {
    case "a":
        v18 := GenS1{F0: 0, F1: 61}
        _ = v18
    default:
        v19 := GenS2{F0: (!(721 > 32768))}
        _ = v19
    }
    v17 = "\n\t"
    var v20 int8 = -128
    var v21 uint = 255
    v22 := GenS1{F0: 65535, F1: v21}
    _ = len(v2)
    _ = len(v2)
    v21 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v7
}
