package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 int16; F2 int8 }
func genFn0(p0 uint64) uint {
    if false {
        v5 := []int{-108, 32768, 9223372036854775807, 32768}
        _ = v5
    } else {
        switch 498 {
        case -2147483648:
            switch 2147483647 {
            case -1:
                var v6 uint32 = 2147483648
                _ = v6
            case 9223372036854775807:
                v7 := []int{32767}
                _ = v7
            case -506:
                p0 -= (p0 * 58)
            }
        case 32767:
            switch 776 {
            case 127:
                v8 := GenS1{F0: 1}
                _ = v8
            case 32767:
                if (true && false) {
                    var v9 uint16 = 255
                    var v10 int16 = -1
                    var v11 int64 = 9223372036854775807
                    _ = v10
                    _ = v11
                    _ = v9
                }
            case -1:
                for fi12 := 0; fi12 < 9; fi12++ {
                    var v13 int32 = 42
                    switch "a" {
                    case "hello":
                        v13 = (v13 >> ((42) & 63))
                    }
                    var v14 int32 = v13
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            }
        default:
            var v15 int32 = -796
            _ = v15
        }
        v16 := []int{}
        p0 += 2
        _ = v16
    }
    switch -102 {
    case -745:
        v17 := []int{35, -541, 4294967295, 2147483647}
        _ = v17
    case 0:
        v18 := GenS1{F0: 308}
        _ = v18
    case 1:
        v19 := GenS1{F0: 541}
        _ = v19
    }
    if true {
        p0 = 142
    }
    _ = p0
    return 0
}

func main() {
    var v1 int = 0
    var v2 int32 = -308
    v3 := make(map[int]bool)
    var v4 int32 = v2
    v20 := GenS2{F0: false}
    var v21 uint32 = 42
    v22 := GenS3{F0: 4294967295, F1: 1, F2: 1}
    var v23 uint32 = (^(v21))
    var v24 string = "\n\t"
    for fi25 := 0; fi25 < 9; fi25++ {
        var v26 int16 = 1000
        switch 32768 {
        case -9223372036854775808:
            v27 := []int{}
            _ = v27
        case -222:
            var v28 uint16 = 239
            _ = v28
        }
        delete(v3, fi25)
        _ = fi25
        _ = v26
    }
    for fi29 := 0; fi29 < 14; fi29++ {
        if (v4 != (v2 / (((v2 >> ((0) & 63))) | 1))) {
            v21 -= v21
        }
        _ = fi29
    }
    v3[v1] = ((v24 < v24) || true)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
