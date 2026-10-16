package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 uint8 }
func genFn0() int8 {
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 uint = 4294967295
        for fi8 := 0; fi8 < 14; fi8++ {
            var v9 int64 = 783
            _ = fi8
            _ = v9
        }
        fmt.Println(v7, fi6)
        _ = fi6
        _ = v7
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        for fi11 := 0; fi11 < 11; fi11++ {
            var v12 int16 = 0
            v12 += v12
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    return 5
}

func main() {
    v1 := make(map[uint32]string)
    var v2 bool = false
    var v3 int16 = -271
    var v4 uint32 = 2147483648
    var v5 string = "a"
    _ = len(v1)
    v13 := []int{128, -2147483648, 701, -779}
    v14 := GenS1{F0: 42}
    v1[v4] = v5
    v15 := GenS1{F0: v3}
    var v16 bool = (195 >= 255)
    for fi17 := 0; fi17 < 8; fi17++ {
        if true {
            switch "\n\t" {
            case "a":
                v13 = append(v13, fi17)
            case "\n\t":
                if v2 {
                    var v18 uint64 = 4294967295
                    v3 += v3
                    v4 = 2147483647
                    _ = v18
                }
            default:
                var v19 uint32 = v4
                _ = v19
            }
        } else {
            v5 = v5
        }
        _ = fi17
    }
    var v20 uint8 = 1
    var v21 int32 = 2147483647
    if v2 {
        v22 := GenS1{F0: 1}
        _ = v22
    }
    v13 = append(v13, 580)
    v21 = 2147483647
    switch "a" {
    case "\n\t":
        v2 = (v21 != v21)
    case "":
        v4 = (v4 % (((v4 % ((v4) | 1))) | 1))
    default:
        v5 = ""
    }
    v23 := GenS2{F0: (v20 <= 128), F1: 0}
    var v24 float32 = (-2.25 + (float32(0.0) - 0.1))
    v24 -= v24
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
