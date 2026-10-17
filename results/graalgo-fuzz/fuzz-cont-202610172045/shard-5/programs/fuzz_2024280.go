package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 uint8; F1 int }
func genFn0(p0 uint, p1 uint16, p2 int16) int64 {
    if (!false) {
        v5 := GenS1{F0: 4294967295, F1: 127}
        v6 := GenS1{F0: 202, F1: -128}
        v7 := GenS1{F0: 801, F1: -128}
        _ = v5
        _ = v6
        _ = v7
    } else {
        v8 := []int{134, int(p2)}
        v9 := GenS1{F0: 433, F1: 127}
        v10 := GenS1{F0: 4294967295, F1: -17}
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 uint32 = 713
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 42
}

func main() {
    var v1 float64 = (((-2.25 * (3.14159 - 1e10)) * (0.1 - 1e10)) + float64(1e10))
    var v2 uint64 = 18446744073709551615
    var v3 uint32 = 1
    var v4 string = "foo bar"
    v12 := []int{4294967295, 32768}
    for fi13 := 0; fi13 < 12; fi13++ {
        switch (fi13 ^ -963) {
        case 1:
            _ = len(v12)
        }
        _ = fi13
    }
    v14 := GenS1{F0: v2, F1: 1}
    switch 32767 {
    case 1:
        v15 := GenS1{F0: (v2 / (((v2 * v2)) | 1)), F1: 42}
        _ = v15
    case -29:
        fmt.Println(v4, v3, v2)
    case 42:
        if false {
            var v16 float64 = (v1 - v1)
            _ = v16
        }
    }
    v3 -= (v3 & v3)
    var v17 uint32 = v3
    if (2147483647 != 0) {
        v2 = (v2 ^ 2)
        var v18 bool = false
        _ = v18
    }
    if ((v2 & (v2 / ((18446744073709551615) | 1))) > (^(v2))) {
        v19 := GenS2{F0: 0, F1: 127}
        v3 = 65535
        _ = v19
    }
    var v20 int64 = genFn0(0, 255, 42)
    var v21 string = ("hello" + v4)
    v12 = append(v12, 42)
    v21 = ((v21 + v21) + v21)
    var v22 int32 = 447
    _ = len(v12)
    v20 = v20
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v14
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
