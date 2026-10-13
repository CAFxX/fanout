package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 int; F2 int32 }
type GenS4 struct { F0 int8; F1 bool; F2 uint32 }
type GenS5 struct { F0 int32; F1 uint8 }
func genFn0(p0 uint64) uint64 {
    if true {
        v7 := GenS1{F0: 1}
        _ = v7
    }
    p0 = (p0 % (((p0 - p0)) | 1))
    var v8 uint = 255
    v9 := GenS2{F0: 2147483647}
    _ = p0
    _ = v8
    _ = v9
    return 2
}

func main() {
    var v1 uint8 = 128
    v2 := make(map[uint32]bool)
    var v3 uint32 = 2
    v4 := make(map[int]bool)
    v5 := []int{32768, 719}
    v6 := make(map[int]string)
    v3 -= 1
    var v10 int64 = 42
    v11 := GenS1{F0: (v3 + v3)}
    v12 := GenS3{F0: -1, F1: 0, F2: -100}
    var v13 int8 = 42
    fmt.Println(v1, v13)
    v3 = 686
    if true {
        var v14 uint64 = 2
        var v15 int32 = 42
        if false {
            v15 = -1
        } else {
            v16 := GenS4{F0: v13, F1: (true || (1000 > -43)), F2: 262}
            switch 176 {
            case 992:
                v17 := GenS4{F0: (v13 << ((1) & 63)), F1: (!true), F2: (v3 % (((v3 & 166)) | 1))}
                _ = v17
            }
            _ = v16
        }
        _ = v14
        _ = v15
    } else {
        var v18 uint = 452
        if false {
            fmt.Println(v3, v13, v10)
        }
        v3 = (v3 + (^(v3)))
        _ = v18
    }
    if (!false) {
        switch int(v1) {
        case 127:
            v19 := GenS5{F0: 13, F1: v1}
            _ = v19
        default:
            if true {
                v20 := GenS3{F0: -870, F1: 42, F2: -565}
                _ = v20
            } else {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v13 = v13
            }
        }
        var v21 int32 = 255
        _ = v21
    }
    v1 += 32
    v22 := GenS5{F0: -1, F1: 0}
    v5 = append(v5, 681)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
