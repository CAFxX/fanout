package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 bool; F1 int64; F2 uint64 }
func genFn0(p0 uint32, p1 int8) uint32 {
    if false {
        v5 := []int{-2147483648, -2147483648, -513, 739}
        for fi6 := 0; fi6 < 19; fi6++ {
            p0 -= 1
            _ = len(v5)
            _ = fi6
        }
        _ = v5
    } else {
        var v7 uint8 = 0
        for fi8 := 0; fi8 < 18; fi8++ {
            v9 := []int{42, (fi8 / (((fi8 >> ((18446744073709551615) & 63))) | 1))}
            var v10 uint8 = 1
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = v7
    }
    switch ((("\n\t" + "") + "hello") + ("foo bar" + "a")) {
    case "":
        p0 += (p0 & (p0 & p0))
    case "a":
        var v11 int64 = 506
        _ = v11
    default:
        if true {
            switch 753 {
            case 128:
                var v12 int64 = 912
                _ = v12
            default:
                switch "\n\t" {
                case "hello":
                    switch 0 {
                    case -258:
                        p0 = (p0 + (p0 / ((p0) | 1)))
                    default:
                        switch "a" {
                        case "\n\t":
                            var v13 uint64 = 0
                            _ = v13
                        }
                    }
                case "\n\t":
                    switch -2147483648 {
                    case -493:
                        for fi14 := 0; fi14 < 5; fi14++ {
                            if true {
                                v15 := []int{(^(fi14)), (fi14 / (((fi14 & (^(fi14)))) | 1)), (fi14 + (fi14 ^ (fi14 >> ((255) & 63)))), int(p1)}
                                _ = v15
                            } else {
                                var v16 uint64 = 0
                                _ = v16
                            }
                            var v17 int32 = -2147483648
                            _ = fi14
                            _ = v17
                        }
                    case 127:
                        v18 := GenS1{F0: -1, F1: (p0 << ((0) & 63))}
                        _ = v18
                    case 4294967295:
                        v19 := []int{}
                        _ = v19
                    default:
                        var v20 bool = false
                        _ = v20
                    }
                default:
                    p0 = (^(p0))
                }
            }
        }
    }
    _ = p0
    _ = p1
    return p0
}

func main() {
    var v1 float32 = 1e10
    var v2 int16 = -1
    var v3 bool = true
    var v4 float64 = 1e10
    v4 = ((v4 * ((((1e10 + 0.0) + v4) * v4) + v4)) + v4)
    for fi21 := 0; fi21 < 5; fi21++ {
        v3 = v3
        var v22 int16 = 32767
        var v23 uint8 = 108
        _ = fi21
        _ = v22
        _ = v23
    }
    if v3 {
        v24 := []int{127, 2147483647, -1, -434}
        v25 := GenS1{F0: -1, F1: 453}
        v26 := GenS2{F0: false, F1: 42, F2: 42}
        _ = v24
        _ = v25
        _ = v26
    } else {
        v27 := GenS1{F0: 0, F1: 42}
        v28 := []int{-9223372036854775808, 42, 9223372036854775807}
        v3 = (-9223372036854775808 == -9223372036854775808)
        _ = v27
        _ = v28
    }
    v3 = true
    var v29 bool = (!("foo bar" > ("hello" + "\n\t")))
    for fi30 := 0; fi30 < 13; fi30++ {
        v2 += v2
        _ = fi30
    }
    v31 := []int{}
    if ((0.0 + v1) <= float32(0.1)) {
        if len(v31) > 0 {
            v31[0] = 2147483647
        }
        fmt.Println(v1)
    }
    for fi32 := 0; fi32 < 20; fi32++ {
        v3 = (899 <= 42)
        _ = fi32
    }
    var v33 uint16 = 65535
    v34 := GenS1{F0: 1, F1: 15}
    var v35 uint8 = 2
    var v36 int = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v31
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
