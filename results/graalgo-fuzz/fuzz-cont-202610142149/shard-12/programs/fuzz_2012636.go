package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 uint8; F2 string }
type GenS4 struct { F0 int16; F1 bool; F2 uint }
type GenS5 struct { F0 int16; F1 int64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint8; F1 int32 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint16; F1 uint; F2 bool }

func main() {
    var v1 string = "foo bar"
    var v2 int8 = 1
    v3 := []uint32{0, 2147483648, 851}
    var v4 float64 = (0.0 / (-2.25 / float64((0.1 - -2.25))))
    v5 := make(map[uint32]uint32)
    var v6 int32 = -2147483648
    v5[2147483648] = 362
    v7 := v5[816]
    switch v1 {
    case "a":
        v8 := v5[v7]
        _ = v8
    case "hello":
        var v9 float32 = ((1.5 * -2.25) / 1.5)
        _ = v9
    default:
        switch v1 {
        case "a":
            for fi10 := 0; fi10 < 7; fi10++ {
                v11 := GenS1{F0: v1, F1: 604, F2: 42}
                v12 := GenS2{F0: -2147483648}
                _ = fi10
                _ = v11
                _ = v12
            }
        }
    }
    var v13 int8 = 100
    if false {
        v3 = append(v3, (v7 & (v7 * (v7 / ((v7) | 1)))))
    } else {
        for fi14 := 0; fi14 < 1; fi14++ {
            v2 += -82
            _ = fi14
        }
        if false {
            for fi15 := 0; fi15 < 20; fi15++ {
                for fi16 := 0; fi16 < 17; fi16++ {
                    _ = len(v5)
                    if true {
                        v13 -= (v13 % ((v2) | 1))
                        var v17 string = v1
                        _ = v17
                    } else {
                        v18 := GenS3{F0: ((v1 + "a") + "foo bar"), F1: 137, F2: v1}
                        fmt.Println(v6, fi15)
                        _ = v18
                    }
                    _ = fi16
                }
                var v19 int32 = v6
                _ = fi15
                _ = v19
            }
            v20 := GenS4{F0: 32767, F1: true, F2: 3}
            v21 := GenS5{F0: 32767, F1: 186}
            _ = v20
            _ = v21
        }
        _ = len(v3)
    }
    v4 += -2.25
    if false {
        for fi22 := 0; fi22 < 3; fi22++ {
            for fi23 := 0; fi23 < 16; fi23++ {
                v24 := GenS6{F0: -18}
                _ = fi23
                _ = v24
            }
            _ = fi22
        }
    }
    var v25 string = v1
    v26 := GenS7{F0: 128, F1: -586}
    for fi27 := 0; fi27 < 5; fi27++ {
        switch fi27 {
        case 702:
            var v28 int16 = -420
            _ = v28
        }
        var v29 uint64 = 853
        _ = fi27
        _ = v29
    }
    var v30 int16 = 1000
    for fi31 := 0; fi31 < 16; fi31++ {
        var v32 float32 = (((0.0 / 1e10) + 1.5) / -2.25)
        v33 := GenS8{F0: (v2 >> ((2) & 63))}
        _ = fi31
        _ = v32
        _ = v33
    }
    v30 -= (v30 - v30)
    v34 := GenS9{F0: 65535, F1: 1, F2: true}
    for fi35 := 0; fi35 < 6; fi35++ {
        v36 := GenS6{F0: v13}
        _ = fi35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
