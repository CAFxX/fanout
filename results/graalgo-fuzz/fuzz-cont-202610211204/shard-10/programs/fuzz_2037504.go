package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 uint64; F1 int64 }
func genFn0() int {
    var v4 int32 = 42
    v5 := []int{851, 32767, 4294967295, 491}
    _ = v4
    _ = v5
    return 128
}
func genFn1() int64 {
    var v6 int16 = 42
    v7 := GenS1{F0: 220}
    _ = v6
    _ = v7
    return -642
}

func main() {
    var v1 uint32 = 2
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]int)
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 uint64 = 4294967295
        _ = fi8
        _ = v9
    }
    var v10 int64 = genFn1()
    var v11 int32 = 632
    fmt.Println(v1, v11, v1)
    if false {
        delete(v2, 658)
        v12 := GenS2{F0: 0, F1: ((1.5 * 0.0) <= ((3.14159 * float32(((((-2.25 + ((1e10 / 0.1) * 1.5)) / -2.25) + -2.25) + (1.5 * 1e10)))) - float32((-2.25 + -2.25))))}
        _ = v12
    }
    var v13 float64 = ((-2.25 * float64(1e10)) / 3.14159)
    for fi14 := 0; fi14 < 11; fi14++ {
        delete(v2, 1)
        var v15 uint = 778
        _ = fi14
        _ = v15
    }
    var v16 int8 = 127
    for fi17 := 0; fi17 < 17; fi17++ {
        v18 := GenS3{F0: v11}
        _ = fi17
        _ = v18
    }
    var v19 float32 = (0.0 - -2.25)
    for fi20 := 0; fi20 < 1; fi20++ {
        var v21 uint16 = 1
        fmt.Println(v16, v1, v1)
        _ = fi20
        _ = v21
    }
    var v22 float64 = float64(0.1)
    v13 += 3.14159
    v23 := GenS4{F0: "foo bar", F1: 128}
    v24 := GenS3{F0: v11}
    for fi25 := 0; fi25 < 20; fi25++ {
        for fi26 := 0; fi26 < 5; fi26++ {
            v19 = ((1.5 / v19) + v19)
            delete(v3, 563)
            _ = fi26
        }
        _ = fi25
    }
    switch 32768 {
    case -308:
        if false {
            if true {
                v27 := GenS5{F0: 250, F1: (^(v10))}
                _ = v27
            } else {
                v2[42] = -2147483648
                v28 := GenS4{F0: (("" + "\n\t") + ""), F1: 1}
                var v29 uint16 = 32767
                _ = v28
                _ = v29
            }
        }
    case 42:
        v30 := []int{165, 1, 611}
        _ = v30
    case 350:
        var v31 int64 = (v10 >> ((255) & 63))
        _ = v31
    }
    var v32 uint64 = 937
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v16
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v32
}
