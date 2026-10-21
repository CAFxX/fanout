package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8; F1 int8; F2 uint }
type GenS6 struct { F0 bool }

func main() {
    var v1 uint8 = 1
    var v2 int = -78
    v3 := []int{855}
    var v4 uint32 = 831
    fmt.Println(v1)
    var v5 int16 = 1000
    _ = len(v3)
    if (1 > 1) {
        var v6 int8 = 86
        if len(v3) > 0 {
            v3[0] = 0
        }
        _ = v6
    } else {
        if true {
            if false {
                v3 = append(v3, (v2 | (v2 >> ((2) & 63))))
            } else {
                if len(v3) > 0 {
                    v3[0] = (v2 & v2)
                }
                for fi7 := 0; fi7 < 6; fi7++ {
                    v8 := GenS1{F0: 908, F1: (v5 >> ((42) & 63))}
                    _ = fi7
                    _ = v8
                }
                var v9 int = v2
                _ = v9
            }
        }
    }
    var v10 uint = 65535
    for fi11 := 0; fi11 < 11; fi11++ {
        v4 -= v4
        v12 := GenS1{F0: 0, F1: (v5 / (((-(v5))) | 1))}
        _ = fi11
        _ = v12
    }
    v13 := GenS1{F0: v4, F1: v5}
    for fi14 := 0; fi14 < 15; fi14++ {
        fmt.Println(v10, v10, fi14)
        v15 := GenS2{F0: 28, F1: v2}
        _ = fi14
        _ = v15
    }
    switch (v2 * v2) {
    case 4294967295:
        switch v2 {
        case 0:
            var v16 uint8 = (v1 & v1)
            _ = v16
        case -9223372036854775808:
            if (!true) {
                v10 = v10
                for fi17 := 0; fi17 < 14; fi17++ {
                    v18 := GenS2{F0: 127, F1: (v2 % ((fi17) | 1))}
                    fmt.Println(v10)
                    _ = fi17
                    _ = v18
                }
            } else {
                switch 2147483647 {
                case -9223372036854775808:
                    v5 -= v5
                case 9223372036854775807:
                    var v19 uint16 = 65535
                    _ = v19
                default:
                    v20 := GenS3{F0: v4}
                    _ = v20
                }
                if (0.0 <= 3.14159) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                fmt.Println(v10, v1, v4)
            }
        case 2147483647:
            v10 -= 0
        default:
            v4 -= (v4 ^ 0)
        }
    case 32768:
        v21 := GenS4{F0: v10}
        _ = v21
    case 42:
        for fi22 := 0; fi22 < 10; fi22++ {
            v2 = (v2 & 1)
            _ = fi22
        }
    default:
        var v23 int = (v2 + -376)
        _ = v23
    }
    switch v2 {
    case 128:
        v24 := GenS1{F0: v4, F1: v5}
        _ = v24
    }
    v25 := GenS5{F0: v1, F1: 0, F2: v10}
    var v26 bool = false
    v5 -= v5
    if false {
        fmt.Println(v10)
    } else {
        v27 := GenS6{F0: v26}
        v28 := GenS6{F0: (-1 < -428)}
        _ = v27
        _ = v28
    }
    for fi29 := 0; fi29 < 5; fi29++ {
        fmt.Println(v10)
        var v30 float32 = (1e10 / 3.14159)
        _ = fi29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
