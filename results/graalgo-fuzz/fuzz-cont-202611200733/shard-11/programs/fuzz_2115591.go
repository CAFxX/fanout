package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 string }
type GenS2 struct { F0 int }

func main() {
    v1 := []uint32{0, 2}
    v2 := make(map[int]int)
    var v3 int32 = 2147483647
    var v4 uint = 65535
    v5 := []uint32{2147483648}
    v6 := make(map[string]bool)
    var v7 int = 127
    fmt.Println(v3)
    var v8 uint64 = 42
    if (false || (!false)) {
        v9 := GenS1{F0: false, F1: true, F2: "a"}
        fmt.Println(v3)
        _ = v9
    }
    if (1e10 == (float64((3.14159 * 1.5)) + 1.5)) {
        v4 = v4
    } else {
        fmt.Println(v3, v7, v3)
        var v10 float64 = 1e10
        _ = v10
    }
    var v11 int32 = v3
    v5 = append(v5, 831)
    var v12 float64 = 1.5
    switch 1 {
    case 4294967295:
        for fi13 := 0; fi13 < 20; fi13++ {
            var v14 int16 = 32767
            _ = fi13
            _ = v14
        }
    case 127:
        v8 = 1
    case 32768:
        var v15 int16 = -32768
        _ = v15
    default:
        for fi16 := 0; fi16 < 8; fi16++ {
            for fi17 := 0; fi17 < 5; fi17++ {
                v18 := GenS1{F0: false, F1: false, F2: "a"}
                v3 = 1
                v12 = 1e10
                _ = fi17
                _ = v18
            }
            v19 := GenS1{F0: (true && true), F1: (2 < 65535), F2: ("\n\t" + ("hello" + "hello"))}
            _ = fi16
            _ = v19
        }
    }
    v3 -= (^(v3))
    v20 := v2[v7]
    v4 += v4
    v3 += 0
    v3 = (v11 / ((-1) | 1))
    v11 += v11
    var v21 int16 = 1000
    for fi22 := 0; fi22 < 5; fi22++ {
        v23 := GenS2{F0: -838}
        _ = fi22
        _ = v23
    }
    var v24 bool = ((v7 * 0) < v20)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
