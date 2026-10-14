package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int }
func genFn0() int16 {
    v7 := GenS1{F0: true}
    if true {
        v8 := GenS1{F0: (255 == 18446744073709551615)}
        _ = v8
    }
    _ = v7
    return -1
}

func main() {
    v1 := make(map[string]bool)
    var v2 int8 = 100
    v3 := make(map[uint32]bool)
    v4 := []uint32{0, 4294967295}
    var v5 uint = 42
    v6 := make(map[uint32]uint32)
    delete(v6, 666)
    var v9 uint64 = 537
    for fi10 := 0; fi10 < 1; fi10++ {
        _ = len(v3)
        var v11 float64 = (3.14159 * 0.1)
        v5 = 923
        _ = fi10
        _ = v11
    }
    switch 290 {
    case 2147483647:
        v5 -= 2
    case -151:
        var v12 int = 32768
        _ = v12
    case -2147483648:
        var v13 float32 = float32(0.1)
        _ = v13
    default:
        for fi14 := 0; fi14 < 13; fi14++ {
            for fi15 := 0; fi15 < 18; fi15++ {
                if false {
                    var v16 float64 = 0.0
                    v16 = float64(0.1)
                    _ = v16
                }
                v9 -= 42
                v9 -= (v9 >> ((v5) & 63))
                _ = fi15
            }
            var v17 int8 = v2
            var v18 int64 = 0
            _ = fi14
            _ = v17
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 2; fi19++ {
        var v20 string = (("hello" + "\n\t") + ("\n\t" + "\n\t"))
        _ = fi19
        _ = v20
    }
    var v21 int16 = 0
    if false {
        if false {
            delete(v6, 65535)
        }
        v21 = v21
    } else {
        for fi22 := 0; fi22 < 11; fi22++ {
            for fi23 := 0; fi23 < 11; fi23++ {
                var v24 uint = (v5 * v5)
                _ = fi23
                _ = v24
            }
            _ = fi22
        }
    }
    if (249 != 70) {
        v21 -= (v21 / (((v21 | (v21 << ((255) & 63)))) | 1))
        v25 := GenS2{F0: 351}
        _ = v25
    }
    for fi26 := 0; fi26 < 16; fi26++ {
        var v27 float32 = ((0.1 / 1.0) * (((1e10 + -2.25) / 0.1) - float32(0.0)))
        var v28 int8 = v2
        _ = fi26
        _ = v27
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
