package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int8 }
func genFn0(p0 int64) int32 {
    var v7 int16 = 0
    var v8 int8 = -107
    v9 := GenS1{F0: -2147483648, F1: p0, F2: 7}
    for fi10 := 0; fi10 < 1; fi10++ {
        var v11 uint8 = 128
        _ = fi10
        _ = v11
    }
    switch -2147483648 {
    case 127:
        if ((v8 ^ v8) == v8) {
            v8 = (-(v8))
        } else {
            v12 := []int{127, 747, 32767, 127}
            v8 = (v8 >> ((544) & 63))
            _ = v12
        }
    case 96:
        v8 = 42
    }
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return -2147483648
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 0
    var v3 uint = 742
    v4 := []int{-1}
    var v5 uint8 = (v2 & (^(v2)))
    v6 := make(map[int]bool)
    v5 = (v2 / (((v2 & v2)) | 1))
    fmt.Println(v5, v2)
    v13 := v1[2147483648]
    var v14 uint16 = 1
    for fi15 := 0; fi15 < 13; fi15++ {
        v5 += (v2 << (((v3 + (v3 * v3))) & 63))
        _ = fi15
    }
    _ = len(v4)
    var v16 string = ("hello" + "\n\t")
    for fi17 := 0; fi17 < 10; fi17++ {
        v18 := GenS1{F0: -262, F1: 174, F2: -1}
        if (v3 == 2) {
            v19 := GenS1{F0: 42, F1: 1, F2: 100}
            v20 := GenS1{F0: -1, F1: -1, F2: 42}
            _ = v19
            _ = v20
        }
        v16 = v16
        _ = fi17
        _ = v18
    }
    if (((v14 & (^(v14))) == (v14 * (v14 + 1))) || v13) {
        var v21 string = "hello"
        _ = v21
    } else {
        var v22 int8 = -1
        _ = v22
    }
    fmt.Println(v3, v3, v16)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
