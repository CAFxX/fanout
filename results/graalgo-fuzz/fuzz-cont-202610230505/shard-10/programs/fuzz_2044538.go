package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint8; F1 int }
func genFn0() uint8 {
    if (0 >= 4294967295) {
        var v7 uint64 = 1
        var v8 uint8 = 255
        _ = v7
        _ = v8
    } else {
        var v9 uint32 = 2
        _ = v9
    }
    v10 := []int{-1}
    var v11 int16 = 1
    _ = v10
    _ = v11
    return 0
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]int)
    v3 := []uint32{}
    var v4 int8 = -128
    var v5 int8 = 1
    var v6 uint32 = 0
    if (2 <= 42) {
        var v12 int = 42
        if false {
            v13 := GenS1{F0: 1, F1: 100}
            v14 := GenS1{F0: -32768, F1: (v4 >> ((820) & 63))}
            switch "hello" {
            case "foo bar":
                var v15 uint32 = v6
                _ = v15
            case "\n\t":
                var v16 int = v12
                _ = v16
            case "a":
                for fi17 := 0; fi17 < 9; fi17++ {
                    var v18 bool = true
                    for fi19 := 0; fi19 < 9; fi19++ {
                        var v20 uint32 = (^(v6))
                        v21 := GenS2{F0: 127, F1: fi17}
                        var v22 int64 = 42
                        _ = fi19
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                    _ = fi17
                    _ = v18
                }
            default:
                fmt.Println(v4)
            }
            _ = v13
            _ = v14
        }
        _ = v12
    }
    var v23 float64 = 1.5
    var v24 int8 = (v5 << ((2) & 63))
    v25 := GenS2{F0: genFn0(), F1: -1}
    v6 += (v6 * v6)
    var v26 uint32 = (v6 ^ (v6 + (v6 >> ((255) & 63))))
    if true {
        for fi27 := 0; fi27 < 13; fi27++ {
            var v28 uint32 = v6
            _ = len(v3)
            _ = fi27
            _ = v28
        }
        if false {
            v3 = append(v3, (v26 >> ((0) & 63)))
            if (!(1 == 1)) {
                v4 -= v5
            } else {
                v24 = (-(v4))
            }
        } else {
            var v29 int32 = -2147483648
            v2[v6] = 128
            _ = v29
        }
        var v30 uint64 = 2
        _ = v30
    }
    var v31 int16 = 1000
    var v32 float32 = 0.1
    for fi33 := 0; fi33 < 8; fi33++ {
        if (v24 > (v24 * v5)) {
            v31 = 32767
        }
        _ = fi33
    }
    var v34 int8 = (-(v24))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
