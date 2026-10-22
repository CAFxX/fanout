package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 int64; F1 bool; F2 uint16 }
type GenS3 struct { F0 uint8 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]bool)
    var v3 int16 = -32768
    var v4 int64 = 366
    var v5 int8 = 127
    v6 := []int{-590, -1, 127, 42}
    if false {
        v3 -= (^(v3))
    } else {
        switch 128 {
        case -2147483648:
            if (true || (1 < -739)) {
                v7 := GenS1{F0: 18446744073709551615, F1: 1}
                _ = v7
            }
        case -1:
            v3 -= (v3 + (v3 << ((39) & 63)))
        }
        var v8 string = "a"
        _ = v8
    }
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 bool = true
        delete(v1, 434)
        _ = fi9
        _ = v10
    }
    var v11 uint16 = 42
    for fi12 := 0; fi12 < 14; fi12++ {
        for fi13 := 0; fi13 < 9; fi13++ {
            v14 := GenS1{F0: 42, F1: 42}
            if ((^(v11)) == (v11 ^ v11)) {
                fmt.Println(v11, v4, v11)
            }
            _ = fi13
            _ = v14
        }
        var v15 bool = true
        for fi16 := 0; fi16 < 15; fi16++ {
            v1[2] = 339
            v6 = append(v6, (fi16 + fi12))
            _ = fi16
        }
        _ = fi12
        _ = v15
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        v5 += (v5 % (((v5 / (((v5 + (v5 & (v5 / ((v5) | 1))))) | 1))) | 1))
        var v18 int = (^(fi17))
        _ = fi17
        _ = v18
    }
    var v19 int32 = 2147483647
    var v20 int8 = v5
    var v21 float64 = (-2.25 - 1e10)
    v22 := GenS2{F0: (v4 & 604), F1: true, F2: v11}
    v5 += (v20 >> ((65535) & 63))
    switch 128 {
    case -1:
        if true {
            v11 += (v11 ^ 255)
        }
    case 42:
        v23 := GenS3{F0: 1}
        _ = v23
    case 555:
        fmt.Println(v20, v20, v19)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
