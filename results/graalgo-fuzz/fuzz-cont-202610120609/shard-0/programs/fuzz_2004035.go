package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 uint64; F2 int64 }
type GenS4 struct { F0 int16; F1 int32; F2 int8 }
func genFn0(p0 int8, p1 int32) int16 {
    var v6 uint16 = 32767
    if (797 <= -93) {
        for fi7 := 0; fi7 < 17; fi7++ {
            switch int(p0) {
            case 1:
                p1 = (p1 | (p1 - p1))
            default:
                p1 = p1
            }
            _ = fi7
        }
        v8 := []int{-330, -943}
        _ = v8
    }
    v9 := []int{}
    v10 := GenS1{F0: (^(p0)), F1: "foo bar"}
    var v11 float64 = 1.5
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v6
    _ = v9
    return 1
}

func main() {
    var v1 uint32 = 0
    v2 := make(map[string]uint32)
    var v3 float32 = (-2.25 - 0.1)
    var v4 uint = 638
    var v5 uint = (v4 - v4)
    var v12 float64 = (1.5 - -2.25)
    for fi13 := 0; fi13 < 5; fi13++ {
        var v14 float32 = 3.14159
        var v15 float32 = float32(v12)
        _ = fi13
        _ = v14
        _ = v15
    }
    v4 += 18446744073709551615
    var v16 uint8 = 255
    v17 := GenS2{F0: -1}
    v18 := GenS3{F0: (!false), F1: 245, F2: 1}
    var v19 int8 = 100
    v20 := []int{42, 2147483647, -1, 128}
    switch "a" {
    case "a":
        var v21 int16 = 389
        _ = v21
    case "":
        var v22 int16 = int16(v4)
        _ = v22
    case "foo bar":
        var v23 float32 = v3
        _ = v23
    }
    v24 := GenS4{F0: -32768, F1: -1, F2: (-(v19))}
    var v25 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
