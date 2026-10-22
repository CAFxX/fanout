package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 uint64 }
func genFn0(p0 uint32) uint64 {
    v6 := GenS1{F0: 464, F1: 1, F2: 891}
    if false {
        v7 := []int{32768}
        for fi8 := 0; fi8 < 13; fi8++ {
            var v9 int16 = int16(p0)
            if ((3.14159 * (1.5 - (3.14159 * 1e10))) <= (((1.5 + 0.0) / float64(1.5)) / 0.1)) {
                var v10 uint8 = 115
                _ = v10
            }
            _ = fi8
            _ = v9
        }
        var v11 uint64 = 92
        _ = v11
        _ = v7
    } else {
        p0 = 2147483648
    }
    var v12 uint8 = 0
    _ = p0
    _ = v12
    _ = v6
    return 42
}
func genFn1() uint8 {
    var v13 int8 = 127
    v14 := []int{}
    v13 = 109
    var v15 int32 = -1
    if (-80 > 0) {
        switch -1 {
        case 127:
            var v16 int32 = (-(v15))
            _ = v16
        default:
            v15 -= -777
        }
        v13 = 0
    }
    _ = v13
    _ = v14
    _ = v15
    return 0
}

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[uint32]string)
    var v3 uint32 = 0
    var v4 int = 128
    v5 := make(map[uint32]int)
    switch "" {
    case "hello":
        v17 := GenS1{F0: 0, F1: 0, F2: v1}
        _ = v17
    default:
        if (!false) {
            if (1000 < 1000) {
                fmt.Println(v3)
                var v18 uint = 255
                v3 += (^(v3))
                _ = v18
            }
        }
    }
    v19 := []int{-781, v4, v4, (v4 ^ (v4 * -34))}
    fmt.Println(v1, v3)
    for fi20 := 0; fi20 < 5; fi20++ {
        v3 += v3
        _ = fi20
    }
    v3 -= (v3 ^ (v3 + (v3 | (^(v3)))))
    if ("" == "hello") {
        if len(v19) > 0 {
            v19[0] = v4
        }
        for fi21 := 0; fi21 < 4; fi21++ {
            v1 = v1
            fmt.Println(v4, v4, v3)
            _ = fi21
        }
    } else {
        switch "" {
        case "hello":
            v3 = v3
        default:
            v22 := v2[677]
            _ = v22
        }
    }
    var v23 uint32 = 2
    v23 += v3
    if (2 <= 255) {
        if true {
            fmt.Println(v23, v4, v23)
            var v24 int8 = -128
            _ = v24
        }
        delete(v5, 2147483647)
    } else {
        var v25 int32 = 2147483647
        _ = v25
    }
    v1 = (v1 >> ((947) & 63))
    var v26 float32 = float32(((float64(float32((3.14159 * 1e10))) + 3.14159) * float64(((1.5 / 1e10) * 1.5))))
    fmt.Println(v4)
    v4 = 5
    var v27 int = (-(v4))
    switch ("" + ("hello" + "foo bar")) {
    case "a":
        var v28 float32 = -2.25
        _ = v28
    case "foo bar":
        switch (v27 ^ (v4 * (v4 ^ (v4 & v27)))) {
        case -70:
            var v29 string = ("" + "a")
            _ = v29
        default:
            v4 = (v27 / ((-680) | 1))
        }
    }
    v3 = 638
    v3 += 0
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
