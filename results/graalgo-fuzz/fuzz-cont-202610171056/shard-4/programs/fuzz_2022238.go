package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 int8 }
type GenS2 struct { F0 int64; F1 int; F2 uint }
func genFn0() uint64 {
    switch 9223372036854775807 {
    case 9223372036854775807:
        switch 1 {
        case -9223372036854775808:
            var v6 int16 = 1000
            _ = v6
        default:
            v7 := GenS1{F0: 1, F1: 763, F2: 0}
            _ = v7
        }
    case 128:
        for fi8 := 0; fi8 < 10; fi8++ {
            v9 := []int{-91, fi8}
            var v10 float64 = (((3.14159 * (1e10 - 0.1)) - 0.1) * (float64(((0.1 * float32(0.1)) - -2.25)) / 1.5))
            v10 = v10
            _ = fi8
            _ = v10
            _ = v9
        }
    default:
        if ((3.14159 / 1.0) == 0.1) {
            v11 := GenS1{F0: 42, F1: 57, F2: 0}
            v12 := GenS1{F0: 9223372036854775807, F1: 478, F2: 1}
            _ = v11
            _ = v12
        }
    }
    var v13 uint32 = 650
    _ = v13
    return 2
}

func main() {
    v1 := []int{42}
    v2 := []uint32{65535, 456, 65535, 2147483647}
    var v3 int = 394
    v4 := make(map[uint32]uint32)
    var v5 uint = 633
    switch ("hello" + "hello") {
    case "\n\t":
        var v14 uint8 = 0
        _ = v14
    case "":
        switch 128 {
        case -2147483648:
            v15 := GenS1{F0: 1, F1: 533, F2: -76}
            _ = v15
        }
    default:
        if (1 == 2) {
            var v16 int16 = 360
            _ = v16
        } else {
            var v17 float64 = -2.25
            var v18 float32 = (((1.5 + 1.5) + 0.0) * -2.25)
            _ = v17
            _ = v18
        }
    }
    switch -813 {
    case -689:
        for fi19 := 0; fi19 < 12; fi19++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi19
        }
    case 42:
        var v20 string = "foo bar"
        _ = v20
    case -606:
        var v21 uint32 = 0
        _ = v21
    default:
        v22 := GenS2{F0: 1, F1: v3, F2: v5}
        _ = v22
    }
    v5 = 817
    v3 = v3
    v3 += (v3 >> (((v5 & v5)) & 63))
    fmt.Println(v3)
    fmt.Println(v5)
    fmt.Println(v5, v5)
    v5 += 0
    for fi23 := 0; fi23 < 13; fi23++ {
        v3 = -1
        var v24 uint16 = 980
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 5; fi25++ {
        v26 := GenS1{F0: 0, F1: 2147483647, F2: 127}
        _ = fi25
        _ = v26
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
