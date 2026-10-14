package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 int16 }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint32 }
func genFn0() int {
    v4 := []int{-9223372036854775808, -670}
    var v5 bool = (true && true)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = len(v4)
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 bool, p1 uint64, p2 int) uint32 {
    v6 := GenS1{F0: 941, F1: 1, F2: 42}
    v7 := []int{(p2 * (p2 / (((^(p2))) | 1))), genFn0(), p2}
    v7 = append(v7, p2)
    if len(v7) > 0 {
        v7[0] = p2
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 403
}

func main() {
    var v1 uint16 = 255
    v2 := make(map[uint32]string)
    var v3 int64 = -422
    var v8 uint8 = 1
    var v9 uint64 = 42
    v10 := GenS1{F0: 65535, F1: (v3 * (v3 % (((v3 ^ -112)) | 1))), F2: -1}
    if ("" < "\n\t") {
        _ = len(v2)
        var v11 int32 = 242
        fmt.Println(v3, v8)
        _ = v11
    } else {
        var v12 int32 = -966
        _ = v12
    }
    var v13 uint64 = (v9 + v9)
    v14 := []int{}
    for fi15 := 0; fi15 < 17; fi15++ {
        for fi16 := 0; fi16 < 18; fi16++ {
            if false {
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
            }
            _ = fi16
        }
        _ = fi15
    }
    for fi17 := 0; fi17 < 6; fi17++ {
        v18 := GenS2{F0: v1, F1: v8, F2: 934}
        _ = fi17
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
