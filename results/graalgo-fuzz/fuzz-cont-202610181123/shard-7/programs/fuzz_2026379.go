package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int32 }
type GenS2 struct { F0 uint16; F1 int32 }

func main() {
    var v1 int32 = 1
    v2 := make(map[uint32]bool)
    v3 := []int{-57, -498, 1, 128}
    v4 := []uint32{0, 445}
    switch 128 {
    case -172:
        if true {
            v1 = (v1 % ((v1) | 1))
            for fi5 := 0; fi5 < 1; fi5++ {
                fmt.Println(fi5, v1)
                _ = fi5
            }
        }
    case 2147483647:
        var v6 uint = 65535
        _ = v6
    }
    v1 = v1
    v7 := GenS1{F0: 42, F1: 100, F2: -2147483648}
    fmt.Println(v1)
    fmt.Println(v1)
    for fi8 := 0; fi8 < 19; fi8++ {
        delete(v2, 2147483647)
        if false {
            if false {
                var v9 float32 = (1e10 + 0.1)
                v10 := GenS1{F0: 42, F1: 42, F2: v1}
                _ = v10
                _ = v9
            }
        } else {
            if false {
                fmt.Println(v1, v1)
                var v11 int8 = 100
                switch "hello" {
                case "\n\t":
                    var v12 int32 = v1
                    _ = v12
                case "a":
                    v1 -= (v1 & 2147483647)
                case "hello":
                    var v13 uint64 = 1
                    _ = v13
                default:
                    var v14 bool = true
                    _ = v14
                }
                _ = v11
            } else {
                var v15 int64 = 9223372036854775807
                v15 += -124
                _ = v15
            }
            switch fi8 {
            case 0:
                for fi16 := 0; fi16 < 18; fi16++ {
                    v17 := GenS1{F0: 0, F1: 128, F2: v1}
                    _ = fi16
                    _ = v17
                }
            default:
                if true {
                    _ = len(v3)
                    var v18 int32 = (v1 & 0)
                    _ = v18
                }
            }
        }
        v1 += -2147483648
        _ = fi8
    }
    var v19 int32 = (v1 >> ((42) & 63))
    var v20 bool = true
    v21 := GenS2{F0: 42, F1: v19}
    var v22 float32 = 0.0
    var v23 uint32 = 234
    v24 := GenS2{F0: 369, F1: 0}
    var v25 uint64 = 0
    v26 := GenS1{F0: 1, F1: 2, F2: v1}
    fmt.Println(v25)
    var v27 int = 1
    var v28 uint16 = 0
    var v29 int16 = int16(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v7
}
