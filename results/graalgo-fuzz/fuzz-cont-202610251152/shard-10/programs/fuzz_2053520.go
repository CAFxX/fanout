package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 uint8; F1 int8; F2 int8 }
func genFn0() uint16 {
    var v7 string = "\n\t"
    for fi8 := 0; fi8 < 5; fi8++ {
        v7 = v7
        _ = fi8
    }
    _ = v7
    return 2
}
func genFn1(p0 int32, p1 bool, p2 uint) bool {
    v9 := []int{}
    var v10 float32 = (1.5 / 1e10)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return p1
}

func main() {
    var v1 float32 = -2.25
    var v2 float32 = v1
    var v3 int8 = -50
    v4 := make(map[uint32]bool)
    v5 := make(map[uint32]bool)
    v6 := make(map[int]bool)
    v11 := GenS1{F0: 4294967295, F1: 42}
    v12 := []int{9223372036854775807, -1, -2147483648}
    switch -9223372036854775808 {
    case 891:
        for fi13 := 0; fi13 < 19; fi13++ {
            var v14 int8 = v3
            _ = fi13
            _ = v14
        }
    case 32768:
        v15 := GenS1{F0: 0, F1: 1}
        _ = v15
    }
    fmt.Println(v3, v2)
    for fi16 := 0; fi16 < 5; fi16++ {
        _ = len(v6)
        for fi17 := 0; fi17 < 10; fi17++ {
            var v18 float64 = 1.5
            v19 := GenS2{F0: 127, F1: (v3 ^ (v3 % (((v3 | v3)) | 1))), F2: (v3 & -84)}
            delete(v6, 1)
            _ = fi17
            _ = v18
            _ = v19
        }
        v20 := GenS1{F0: 2147483647, F1: -2147483648}
        _ = fi16
        _ = v20
    }
    v21 := GenS2{F0: 2, F1: 100, F2: v3}
    switch "\n\t" {
    case "":
        var v22 uint32 = 922
        _ = v22
    case "foo bar":
        var v23 int64 = -1
        _ = v23
    }
    v3 += (v3 % (((^(v3))) | 1))
    var v24 int32 = 42
    var v25 uint = 4294967295
    var v26 uint8 = 136
    fmt.Println(v2, v26)
    var v27 uint8 = v26
    v6[-487] = (!((!true) || ((true || false) && true)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
