package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 uint; F1 int32; F2 int16 }
type GenS4 struct { F0 int32 }

func main() {
    var v1 uint64 = 4294967295
    var v2 int16 = -251
    var v3 int64 = 9223372036854775807
    fmt.Println(v3, v2, v3)
    var v4 uint64 = v1
    v5 := []int{}
    v1 = v4
    v6 := GenS1{F0: false, F1: 65535}
    switch -1 {
    case 872:
        v7 := GenS1{F0: false, F1: 65535}
        _ = v7
    case 2147483647:
        switch int(v3) {
        case 4294967295:
            v5 = append(v5, 216)
        case 805:
            var v8 string = ("hello" + "foo bar")
            _ = v8
        default:
            v9 := GenS2{F0: 889, F1: v3}
            _ = v9
        }
    case 764:
        v10 := GenS1{F0: (!(true && false)), F1: 65535}
        _ = v10
    default:
        var v11 uint8 = 42
        _ = v11
    }
    v4 -= 636
    var v12 int8 = 100
    v13 := GenS3{F0: 175, F1: 0, F2: int16(v4)}
    v5 = append(v5, int(v12))
    for fi14 := 0; fi14 < 17; fi14++ {
        var v15 int = int(v12)
        v16 := GenS4{F0: -190}
        _ = fi14
        _ = v15
        _ = v16
    }
    v4 += (v4 ^ 42)
    v12 -= -128
    var v17 float64 = ((0.1 / 1e10) * 1e10)
    v2 -= v2
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
