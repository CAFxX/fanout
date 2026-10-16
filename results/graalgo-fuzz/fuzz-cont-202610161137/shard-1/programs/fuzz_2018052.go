package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 uint16; F1 int; F2 int }
func genFn0() string {
    var v5 int64 = 9223372036854775807
    fmt.Println(v5)
    for fi6 := 0; fi6 < 9; fi6++ {
        fmt.Println(fi6, fi6)
        v7 := []int{fi6}
        _ = fi6
        _ = v7
    }
    v5 += v5
    _ = v5
    return "a"
}

func main() {
    v1 := make(map[int]string)
    var v2 int64 = 9223372036854775807
    var v3 int16 = 0
    var v4 uint = 547
    for fi8 := 0; fi8 < 8; fi8++ {
        v9 := GenS1{F0: 2147483648, F1: int32(v4)}
        for fi10 := 0; fi10 < 1; fi10++ {
            v3 = v3
            _ = fi10
        }
        _ = fi8
        _ = v9
    }
    delete(v1, 127)
    var v11 uint64 = 460
    switch 32768 {
    case -9223372036854775808:
        for fi12 := 0; fi12 < 4; fi12++ {
            v3 -= (v3 + (v3 >> ((v4) & 63)))
            var v13 uint16 = 32767
            _ = fi12
            _ = v13
        }
    case 2147483647:
        var v14 int = -87
        _ = v14
    case -173:
        v15 := GenS1{F0: 42, F1: 42}
        _ = v15
    default:
        switch -1 {
        case -2147483648:
            v16 := []int{2147483647, 42, 2147483647, -1}
            _ = v16
        }
    }
    switch 42 {
    case 1:
        var v17 bool = (0 == 42)
        _ = v17
    case 32768:
        v18 := []int{}
        _ = v18
    default:
        for fi19 := 0; fi19 < 9; fi19++ {
            switch 32768 {
            case 2147483647:
                fmt.Println(v2)
            default:
                v20 := GenS1{F0: 2147483648, F1: -2147483648}
                _ = v20
            }
            v21 := GenS2{F0: uint16(v11), F1: fi19, F2: (fi19 & fi19)}
            _ = fi19
            _ = v21
        }
    }
    if (!(2147483647 != 0)) {
        var v22 float32 = 1e10
        if ((false && (128 >= 2)) || false) {
            var v23 uint8 = 255
            var v24 float64 = 0.1
            v22 = float32((v24 * v24))
            _ = v23
            _ = v24
        }
        var v25 int8 = 42
        _ = v22
        _ = v25
    }
    v26 := GenS2{F0: 2, F1: -354, F2: -2147483648}
    var v27 string = ("\n\t" + "hello")
    var v28 uint64 = v11
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
