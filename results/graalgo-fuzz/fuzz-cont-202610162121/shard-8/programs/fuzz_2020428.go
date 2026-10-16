package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64; F1 uint16 }
func genFn0(p0 int32, p1 uint, p2 uint32) uint32 {
    if (0.0 != (float32(0.1) - 0.0)) {
        for fi7 := 0; fi7 < 13; fi7++ {
            var v8 int16 = 1000
            v9 := []int{-9223372036854775808, -425, 9223372036854775807, fi7}
            _ = fi7
            _ = v8
            _ = v9
        }
        v10 := []int{0, 0, -1, 9223372036854775807}
        switch "hello" {
        case "":
            p2 = uint32(p1)
        default:
            v11 := GenS1{F0: 1, F1: 127}
            _ = v11
        }
        _ = v10
    }
    var v12 uint = 4294967295
    for fi13 := 0; fi13 < 8; fi13++ {
        if ((974 < -547) || (1000 < 0)) {
            var v14 string = "\n\t"
            v15 := []int{}
            switch fi13 {
            case 1:
                var v16 int16 = -32768
                _ = v16
            case 128:
                p1 += p1
            default:
                p2 = p2
            }
            _ = v14
            _ = v15
        }
        var v17 uint32 = (p2 - (p2 >> ((v12) & 63)))
        _ = fi13
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return (p2 % ((65535) | 1))
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]int)
    v3 := make(map[uint32]int)
    var v4 int8 = -1
    var v5 float64 = 1.5
    v6 := make(map[uint32]uint32)
    delete(v6, 2)
    delete(v2, -804)
    var v18 float32 = 0.1
    v19 := GenS2{F0: "a"}
    var v20 int = -1
    for fi21 := 0; fi21 < 4; fi21++ {
        var v22 float64 = 1e10
        v23 := []int{}
        _ = fi21
        _ = v22
        _ = v23
    }
    v24 := GenS3{F0: 0, F1: 401}
    var v25 bool = false
    v26 := []int{1, v20}
    v27 := v6[963]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
