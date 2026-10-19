package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 string; F1 uint64; F2 uint16 }
func genFn0(p0 int) int8 {
    v4 := GenS1{F0: (0 > 990), F1: 65535, F2: 415}
    for fi5 := 0; fi5 < 2; fi5++ {
        var v6 uint8 = 42
        _ = fi5
        _ = v6
    }
    var v7 int8 = 76
    for fi8 := 0; fi8 < 3; fi8++ {
        v7 -= (-(v7))
        v9 := GenS1{F0: (!(2 == 87)), F1: 729, F2: 1}
        v10 := []int{p0, 593}
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = v4
    _ = v7
    return (v7 & v7)
}

func main() {
    var v1 bool = (-32768 <= 42)
    var v2 int = 323
    v3 := make(map[int]bool)
    var v11 uint64 = 42
    fmt.Println(v11, v2, v2)
    var v12 uint8 = 1
    var v13 uint64 = (v11 & v11)
    var v14 string = "foo bar"
    if v1 {
        switch 992 {
        case 42:
            v14 = (v14 + v14)
        }
        v15 := GenS1{F0: false, F1: 744, F2: -107}
        v16 := GenS1{F0: (!(genFn0((v2 & (^(v2)))) <= 127)), F1: 653, F2: -32768}
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        v18 := []int{}
        _ = fi17
        _ = v18
    }
    var v19 uint32 = 0
    v20 := GenS2{F0: (v2 % ((v2) | 1))}
    v21 := GenS1{F0: true, F1: 4294967295, F2: 1000}
    var v22 int16 = -32768
    if (32767 < (-(v22))) {
        v22 = v22
    } else {
        var v23 string = v14
        v24 := []int{v2}
        v25 := GenS3{F0: int(v12), F1: 442}
        _ = v23
        _ = v24
        _ = v25
    }
    switch v2 {
    case -125:
        if (42 != 631) {
            var v26 int32 = 42
            _ = v26
        } else {
            var v27 int32 = 2147483647
            fmt.Println(v27)
            v19 -= v19
            _ = v27
        }
    default:
        v28 := GenS4{F0: ((v14 + "hello") + ((("" + v14) + v14) + ("\n\t" + "hello"))), F1: 42, F2: 0}
        _ = v28
    }
    v2 = v2
    v29 := []int{(v2 + v2)}
    v14 = v14
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v29
    _ = v3
}
