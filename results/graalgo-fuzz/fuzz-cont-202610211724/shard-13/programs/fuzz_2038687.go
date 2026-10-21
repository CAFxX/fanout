package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 bool }
func genFn0() int8 {
    var v4 uint16 = 255
    v4 = 2
    v5 := GenS1{F0: -233, F1: 1, F2: (!((v4 | 255) != 255))}
    v6 := GenS1{F0: 42, F1: 4294967295, F2: true}
    _ = v4
    _ = v5
    _ = v6
    return -128
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 65535
    v3 := make(map[int]int)
    var v7 int32 = 2147483647
    v2 -= v2
    switch 1 {
    case 32767:
        var v8 uint64 = 42
        _ = v8
    case 233:
        var v9 float64 = 0.0
        _ = v9
    }
    fmt.Println(v2, v7)
    v7 = (^(v7))
    v7 -= v7
    var v10 int16 = -32768
    v7 = (v7 % ((v7) | 1))
    v2 -= v2
    for fi11 := 0; fi11 < 12; fi11++ {
        v2 = v2
        _ = fi11
    }
    _ = len(v1)
    fmt.Println(v2)
    v12 := GenS1{F0: v7, F1: 42, F2: false}
    for fi13 := 0; fi13 < 1; fi13++ {
        var v14 uint = 438
        _ = fi13
        _ = v14
    }
    v15 := GenS1{F0: (v7 << ((65535) & 63)), F1: 540, F2: true}
    v16 := GenS1{F0: v7, F1: 178, F2: ("foo bar" <= "\n\t")}
    for fi17 := 0; fi17 < 7; fi17++ {
        switch (fi17 & (fi17 | 689)) {
        case 503:
            v10 += v10
        case 329:
            var v18 int64 = 1
            _ = v18
        default:
            var v19 uint16 = v2
            _ = v19
        }
        switch "foo bar" {
        case "":
            v10 = (^(v10))
        }
        _ = fi17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v7
}
