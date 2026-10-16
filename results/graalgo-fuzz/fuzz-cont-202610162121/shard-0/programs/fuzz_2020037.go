package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 uint16; F2 uint16 }
func genFn0(p0 uint32) uint32 {
    v7 := []int{152, 42}
    var v8 uint = 1
    for fi9 := 0; fi9 < 7; fi9++ {
        var v10 uint8 = 1
        _ = len(v7)
        _ = fi9
        _ = v10
    }
    var v11 uint8 = 42
    _ = p0
    _ = v11
    _ = v7
    _ = v8
    return 94
}

func main() {
    var v1 float64 = 1e10
    v2 := make(map[uint32]uint32)
    var v3 int16 = -32768
    v4 := make(map[string]string)
    var v5 uint32 = 556
    v6 := []uint32{v5, 2147483648}
    fmt.Println(v5, v1)
    var v12 uint16 = 0
    var v13 float64 = v1
    v1 = v13
    for fi14 := 0; fi14 < 15; fi14++ {
        if false {
            fmt.Println(v3, fi14)
            if len(v6) > 0 {
                v6[0] = 2
            }
        } else {
            fmt.Println(v5, fi14)
            var v15 uint8 = 0
            var v16 uint32 = 0
            _ = v15
            _ = v16
        }
        _ = len(v4)
        _ = fi14
    }
    var v17 uint32 = genFn0((^(v5)))
    v18 := GenS1{F0: true}
    v5 -= (v5 ^ (^(v17)))
    for fi19 := 0; fi19 < 18; fi19++ {
        v20 := GenS1{F0: (1 < 1)}
        for fi21 := 0; fi21 < 19; fi21++ {
            v1 = v13
            _ = fi21
        }
        _ = fi19
        _ = v20
    }
    v22 := GenS2{F0: v3, F1: (v12 + uint16(v5)), F2: v12}
    var v23 uint64 = 42
    fmt.Println(v5)
    switch 32768 {
    case -368:
        delete(v4, ("foo bar" + "\n\t"))
    case 33:
        for fi24 := 0; fi24 < 1; fi24++ {
            if false {
                var v25 float32 = 1e10
                _ = v25
            } else {
                fmt.Println(v3, v13)
            }
            _ = fi24
        }
    case 32767:
        if ((0.0 < v13) && (!(!true))) {
            var v26 int = -476
            _ = v26
        }
    default:
        v3 -= (v3 + v3)
    }
    for fi27 := 0; fi27 < 12; fi27++ {
        v5 += v17
        fmt.Println(v5)
        _ = fi27
    }
    var v28 uint16 = v12
    var v29 int = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
