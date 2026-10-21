package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 uint32; F1 int; F2 int16 }
type GenS3 struct { F0 uint16; F1 int16; F2 int64 }
type GenS4 struct { F0 uint16; F1 uint32; F2 bool }
type GenS5 struct { F0 string }
type GenS6 struct { F0 bool; F1 uint16 }
type GenS7 struct { F0 int64; F1 uint }

func main() {
    var v1 int8 = 100
    v2 := make(map[uint32]int)
    var v3 bool = true
    v4 := make(map[int]uint32)
    fmt.Println(v1, v3)
    for fi5 := 0; fi5 < 19; fi5++ {
        switch fi5 {
        case 4294967295:
            v6 := GenS1{F0: -1, F1: 1}
            _ = v6
        case -2147483648:
            var v7 int64 = -324
            _ = v7
        default:
            if true {
                v1 += v1
                v1 -= -128
                v1 -= v1
            } else {
                v8 := GenS2{F0: 4294967295, F1: (fi5 / ((127) | 1)), F2: 1000}
                if (1.5 > 3.14159) {
                    var v9 uint8 = 128
                    switch (fi5 + (fi5 + (fi5 % ((fi5) | 1)))) {
                    case -598:
                        v10 := []int{}
                        _ = v10
                    default:
                        v11 := GenS2{F0: 640, F1: (fi5 / ((fi5) | 1)), F2: 968}
                        _ = v11
                    }
                    if ((0 < 919) && (!v3)) {
                        var v12 int8 = 100
                        _ = v12
                    }
                    _ = v9
                }
                v13 := GenS3{F0: 734, F1: -1, F2: 275}
                _ = v13
                _ = v8
            }
        }
        _ = fi5
    }
    var v14 string = ("hello" + (("a" + "a") + ""))
    var v15 uint32 = 4294967295
    v16 := GenS4{F0: 0, F1: (v15 % (((v15 >> ((634) & 63))) | 1)), F2: (float32((float64((float32(0.1) + -2.25)) * (0.1 * float64((1e10 / 1.5))))) > float32(1e10))}
    if (v15 <= 292) {
        if ((true || v3) && ((v1 ^ (v1 * v1)) > (v1 ^ (v1 | v1)))) {
            for fi17 := 0; fi17 < 19; fi17++ {
                if true {
                    v1 = 90
                    v18 := GenS1{F0: 42, F1: 18446744073709551615}
                    v4[(fi17 - fi17)] = (^(v15))
                    _ = v18
                } else {
                    for fi19 := 0; fi19 < 16; fi19++ {
                        var v20 uint64 = 898
                        _ = fi19
                        _ = v20
                    }
                    v21 := GenS5{F0: "\n\t"}
                    v22 := []int{(^(fi17)), -155, int(v15), fi17}
                    _ = v21
                    _ = v22
                }
                _ = fi17
            }
        }
        v23 := GenS6{F0: v3, F1: 42}
        _ = v23
    }
    var v24 int16 = -125
    v1 = (v1 ^ v1)
    v25 := []int{4294967295, 2147483647, 2147483647, 339}
    var v26 int16 = -1
    var v27 bool = v3
    switch v14 {
    case "foo bar":
        var v28 uint16 = 2
        _ = v28
    case "hello":
        for fi29 := 0; fi29 < 8; fi29++ {
            v30 := GenS7{F0: -560, F1: 842}
            if len(v25) >= 2 {
                v25 = v25[1:len(v25)]
            }
            _ = fi29
            _ = v30
        }
    }
    var v31 int64 = -773
    var v32 uint64 = 28
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v4
}
