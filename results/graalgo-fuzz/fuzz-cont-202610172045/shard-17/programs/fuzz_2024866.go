package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 bool }
func genFn0(p0 bool, p1 string, p2 int8) uint {
    v4 := []int{987}
    v5 := GenS1{F0: -32768, F1: (0 < 65535), F2: -1}
    v6 := GenS1{F0: -1, F1: false, F2: 2147483647}
    var v7 uint32 = 2147483647
    p2 = 53
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 uint64) uint64 {
    v8 := GenS1{F0: 1000, F1: true, F2: 128}
    var v9 uint32 = 646
    _ = p0
    _ = v8
    _ = v9
    return p0
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 327
    var v3 uint = 276
    v10 := GenS2{F0: 255}
    if ((3.14159 / 3.14159) != 1e10) {
        v11 := GenS1{F0: 42, F1: false, F2: 1}
        var v12 float32 = 3.14159
        v12 += ((1e10 + 1.5) - v12)
        _ = v11
        _ = v12
    }
    v13 := GenS2{F0: 255}
    v3 += (v3 << ((genFn0((1000 != 1), "hello", int8(v2))) & 63))
    v2 = v2
    fmt.Println(v3)
    var v14 uint64 = 229
    v15 := []int{}
    for fi16 := 0; fi16 < 4; fi16++ {
        fmt.Println(fi16, v3)
        if true {
            for fi17 := 0; fi17 < 4; fi17++ {
                if (!(1 <= 1000)) {
                    if (2147483648 >= 0) {
                        v2 -= (v2 % ((v2) | 1))
                        v18 := GenS1{F0: 1000, F1: false, F2: int(v14)}
                        _ = v18
                    } else {
                        v14 = genFn1(v14)
                    }
                    switch "a" {
                    case "\n\t":
                        var v19 int32 = -266
                        _ = v19
                    case "":
                        v3 += (v3 % (((v3 << (((v3 - v3)) & 63))) | 1))
                    case "hello":
                        v20 := GenS3{F0: genFn0((false || (42 == 2147483647)), "hello", 100), F1: true}
                        _ = v20
                    default:
                        _ = len(v15)
                        var v21 bool = false
                        _ = v21
                    }
                    var v22 uint16 = 687
                    _ = v22
                } else {
                    var v23 int64 = 0
                    v2 -= (v2 + (^(v2)))
                    _ = v23
                }
                var v24 float32 = 0.1
                fmt.Println(fi17, v3)
                _ = fi17
                _ = v24
            }
        }
        _ = fi16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
