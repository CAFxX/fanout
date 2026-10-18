package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 int64 }
type GenS2 struct { F0 uint; F1 int32; F2 uint }
type GenS3 struct { F0 int64 }
func genFn0(p0 int32, p1 uint64) uint16 {
    if false {
        for fi5 := 0; fi5 < 19; fi5++ {
            v6 := []int{}
            _ = fi5
            _ = v6
        }
    } else {
        v7 := GenS1{F0: 128, F1: p0, F2: 9223372036854775807}
        var v8 string = "hello"
        var v9 float64 = 0.1
        _ = v7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 140, F1: 1, F2: 0}
    _ = p0
    _ = p1
    _ = v10
    return 499
}
func genFn1(p0 uint64) string {
    p0 += (^(p0))
    if false {
        v11 := []int{2147483647, 109}
        if (0.1 < 1e10) {
            var v12 int = -9223372036854775808
            var v13 float32 = (1.5 - 0.0)
            _ = v12
            _ = v13
        } else {
            var v14 uint16 = 581
            var v15 int32 = 2147483647
            v14 = genFn0(-1, 176)
            _ = v14
            _ = v15
        }
        var v16 uint16 = 892
        _ = v11
        _ = v16
    } else {
        v17 := GenS1{F0: 1, F1: 2147483647, F2: 946}
        _ = v17
    }
    _ = p0
    return ("" + "a")
}

func main() {
    v1 := []string{"hello", "", "\n\t", "hello"}
    var v2 int32 = 42
    v3 := make(map[string]bool)
    v4 := []int{19, 128}
    fmt.Println(v2)
    var v18 int16 = -1
    if len(v4) > 0 {
        v4[0] = 4294967295
    }
    v19 := v3["foo bar"]
    var v20 int8 = -105
    v21 := GenS1{F0: 143, F1: (-(v2)), F2: 1}
    v20 -= (v20 / ((v20) | 1))
    var v22 int16 = v18
    var v23 float64 = ((-2.25 / 1.0) - float64(0.1))
    var v24 uint32 = 2147483648
    v25 := GenS1{F0: 0, F1: 2147483647, F2: 1}
    v26 := GenS2{F0: 42, F1: -538, F2: 94}
    switch "\n\t" {
    case "":
        _ = len(v4)
    case "hello":
        var v27 int32 = v2
        _ = v27
    default:
        v24 -= (v24 / (((v24 + 2147483648)) | 1))
    }
    for fi28 := 0; fi28 < 17; fi28++ {
        var v29 int32 = (v2 % ((-2147483648) | 1))
        var v30 int = (fi28 / ((-877) | 1))
        _ = fi28
        _ = v29
        _ = v30
    }
    v31 := GenS3{F0: -9223372036854775808}
    var v32 float32 = 0.0
    fmt.Println(v2)
    var v33 float32 = 0.1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
