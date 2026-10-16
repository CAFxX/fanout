package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 uint32; F1 int; F2 int32 }
func genFn0(p0 int, p1 int8, p2 int64) int {
    if (1e10 <= (0.0 - 1.5)) {
        var v7 float64 = (1.5 - 1.5)
        v8 := []int{}
        if true {
            var v9 bool = false
            var v10 uint64 = 215
            var v11 uint8 = 2
            _ = v10
            _ = v11
            _ = v9
        }
        _ = v7
        _ = v8
    }
    v12 := GenS1{F0: -2147483648, F1: 0}
    p0 += p0
    p0 -= (p0 << ((42) & 63))
    p0 = (p0 + p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return int(p2)
}

func main() {
    var v1 uint32 = 2
    v2 := make(map[string]int)
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]string)
    var v5 int8 = -104
    v6 := []string{("" + "a")}
    for fi13 := 0; fi13 < 7; fi13++ {
        v1 -= 1
        var v14 uint32 = (v1 % (((v1 & v1)) | 1))
        v1 += 1
        _ = fi13
        _ = v14
    }
    if true {
        v6 = append(v6, ("" + ("" + "")))
        switch "" {
        case "":
            v15 := GenS1{F0: 0, F1: (v1 - (^(v1)))}
            _ = v15
        default:
            v5 -= v5
        }
        v16 := GenS2{F0: v1, F1: 961, F2: 2147483647}
        _ = v16
    } else {
        var v17 int32 = -2147483648
        v5 = 127
        _ = v17
    }
    fmt.Println(v5, v5)
    var v18 float64 = -2.25
    v2["\n\t"] = 128
    var v19 int16 = 1
    v19 -= v19
    v19 -= v19
    v5 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
