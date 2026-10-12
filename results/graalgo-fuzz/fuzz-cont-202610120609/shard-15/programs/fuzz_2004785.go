package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 int16; F1 int32; F2 int8 }
func genFn0(p0 int64, p1 string) uint16 {
    for fi5 := 0; fi5 < 10; fi5++ {
        var v6 uint64 = 4294967295
        for fi7 := 0; fi7 < 12; fi7++ {
            v8 := GenS1{F0: 2147483647, F1: v6}
            v6 -= 407
            fmt.Println(p1, fi5)
            var v9 int32 = -767
            _ = fi7
            _ = v8
            _ = v9
        }
        v10 := []int{fi5, 1}
        _ = fi5
        _ = v10
        _ = v6
    }
    var v11 float64 = 3.14159
    v12 := GenS2{F0: 673, F1: -2147483648, F2: 125}
    v11 = (v11 / ((v11 + 1.5) * (1e10 / v11)))
    for fi13 := 0; fi13 < 6; fi13++ {
        var v14 int64 = (p0 << ((4294967295) & 63))
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 0
}
func genFn1(p0 bool) string {
    v15 := GenS1{F0: 480, F1: 0}
    p0 = true
    v16 := []int{}
    p0 = p0
    _ = p0
    _ = v15
    _ = v16
    return "foo bar"
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]string)
    v3 := []int{}
    v4 := make(map[int]int)
    var v17 uint = 748
    if (!(false && (!false))) {
        var v18 uint16 = genFn0(0, "foo bar")
        var v19 float32 = 0.0
        switch "a" {
        case "":
            var v20 int16 = 1
            _ = v20
        case "foo bar":
            fmt.Println(v17)
        default:
            v19 += v19
        }
        _ = v18
        _ = v19
    }
    v21 := GenS2{F0: 1000, F1: -2147483648, F2: -1}
    var v22 float32 = 3.14159
    var v23 int32 = 0
    var v24 int32 = (v23 / (((v23 / ((42) | 1))) | 1))
    _ = len(v2)
    _ = len(v3)
    fmt.Println(v17, v24, v22)
    var v25 uint16 = 0
    var v26 int16 = 32767
    v27 := GenS1{F0: (v24 % ((v23) | 1)), F1: 2}
    v24 = -674
    var v28 uint = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
