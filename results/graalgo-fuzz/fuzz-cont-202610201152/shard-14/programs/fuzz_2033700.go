package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint }
func genFn0(p0 uint16) uint {
    for fi6 := 0; fi6 < 8; fi6++ {
        if true {
            v7 := []int{}
            p0 = p0
            _ = v7
        }
        _ = fi6
    }
    p0 = p0
    _ = p0
    return 65535
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = -1
    var v3 float32 = 0.1
    var v4 float64 = 0.0
    var v5 int8 = 47
    v5 = (v5 * v5)
    for fi8 := 0; fi8 < 19; fi8++ {
        if true {
            switch "hello" {
            case "a":
                for fi9 := 0; fi9 < 13; fi9++ {
                    var v10 int = (fi8 << ((255) & 63))
                    _ = fi9
                    _ = v10
                }
            case "":
                var v11 int8 = v5
                _ = v11
            case "foo bar":
                for fi12 := 0; fi12 < 20; fi12++ {
                    v2 -= 70
                    var v13 bool = (true || false)
                    var v14 uint = 334
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            default:
                v5 = (v5 % ((42) | 1))
            }
            if true {
                if false {
                    v15 := GenS1{F0: 4294967295, F1: 560, F2: genFn0(42)}
                    v16 := GenS1{F0: 42, F1: 0, F2: 704}
                    if true {
                        v2 += v2
                    }
                    _ = v15
                    _ = v16
                }
                var v17 int16 = -1
                fmt.Println(v3)
                _ = v17
            }
        }
        v2 -= v2
        var v18 float32 = float32((float64(v3) / v4))
        _ = fi8
        _ = v18
    }
    v19 := []int{}
    _ = len(v19)
    var v20 float32 = v3
    if false {
        v20 -= v20
        switch 127 {
        case 128:
            v5 = (v5 >> ((719) & 63))
        case -1:
            v1[2147483647] = "hello"
        }
    }
    for fi21 := 0; fi21 < 17; fi21++ {
        v4 -= (v4 / v4)
        _ = fi21
    }
    var v22 int16 = 625
    v2 = v2
    var v23 int32 = (^(v2))
    v24 := v1[750]
    for fi25 := 0; fi25 < 1; fi25++ {
        v2 = v23
        v3 = v3
        var v26 uint16 = 255
        _ = fi25
        _ = v26
    }
    v27 := GenS1{F0: 42, F1: 1000, F2: 18446744073709551615}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
