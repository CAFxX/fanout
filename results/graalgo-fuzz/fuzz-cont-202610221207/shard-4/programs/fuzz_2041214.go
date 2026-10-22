package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int; F1 uint }
type GenS4 struct { F0 uint8; F1 uint; F2 int8 }
func genFn0() int8 {
    for fi7 := 0; fi7 < 14; fi7++ {
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 int32 = -319
            if ((469 <= 65535) || ((false || (1 > -1)) || false)) {
                v10 := GenS1{F0: 65535, F1: 255}
                _ = v10
            }
            v11 := []int{fi7, 984}
            if false {
                v12 := GenS1{F0: 42, F1: 1}
                v9 += 1
                _ = v12
            } else {
                v9 = -683
                v9 -= (v9 * (v9 % ((v9) | 1)))
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
            }
            v9 -= 743
            _ = fi8
            _ = v11
            _ = v9
        }
        _ = fi7
    }
    switch "foo bar" {
    case "":
        v13 := []int{}
        _ = v13
    case "\n\t":
        var v14 uint16 = 42
        _ = v14
    case "hello":
        v15 := GenS2{F0: 538}
        _ = v15
    }
    var v16 uint = 18446744073709551615
    v17 := []int{127, 2147483647}
    _ = v16
    _ = v17
    return -57
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 1
    var v3 float32 = 1e10
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]string)
    var v6 int8 = -128
    v18 := []int{2147483647, 32768, 87}
    if false {
        v19 := GenS3{F0: 32768, F1: 0}
        _ = v19
    } else {
        switch "a" {
        case "foo bar":
            var v20 bool = true
            _ = v20
        case "\n\t":
            _ = len(v18)
        default:
            delete(v4, 9223372036854775807)
        }
        v21 := GenS1{F0: 0, F1: 127}
        _ = v21
    }
    var v22 uint64 = 985
    fmt.Println(v2, v3, v22)
    v6 = v6
    v18 = append(v18, 42)
    var v23 int = -887
    var v24 uint = 42
    for fi25 := 0; fi25 < 17; fi25++ {
        switch fi25 {
        case 128:
            switch "a" {
            case "":
                v26 := GenS4{F0: uint8(v22), F1: (v24 * v24), F2: (^(v6))}
                _ = v26
            case "a":
                var v27 int64 = 0
                _ = v27
            case "hello":
                v3 = 1e10
            }
        }
        _ = fi25
    }
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    v2 = (^(v2))
    var v28 int32 = 1
    var v29 uint16 = (v2 / ((1) | 1))
    v6 -= (v6 | 0)
    var v30 int = v23
    var v31 uint8 = 124
    v18 = append(v18, (v30 | (v30 * 32768)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
