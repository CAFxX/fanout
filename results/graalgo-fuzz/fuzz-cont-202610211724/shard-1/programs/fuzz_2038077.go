package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int8 }
type GenS2 struct { F0 int8; F1 uint16; F2 uint8 }
type GenS3 struct { F0 uint64; F1 int32 }
func genFn0(p0 int8, p1 int, p2 int8) string {
    v4 := []int{}
    for fi5 := 0; fi5 < 6; fi5++ {
        for fi6 := 0; fi6 < 14; fi6++ {
            var v7 int16 = 42
            fmt.Println(p2, p1)
            v8 := GenS1{F0: 42, F1: (p0 / ((p2) | 1)), F2: p0}
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = fi5
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return "foo bar"
}

func main() {
    v1 := make(map[int]int)
    var v2 int16 = 1
    var v3 uint = 2
    v3 -= (v3 >> ((v3) & 63))
    switch 1 {
    case -18:
        if ((!false) || false) {
            switch -491 {
            case 120:
                fmt.Println(v3)
            case -2147483648:
                for fi9 := 0; fi9 < 2; fi9++ {
                    for fi10 := 0; fi10 < 6; fi10++ {
                        v3 += v3
                        _ = fi10
                    }
                    v2 = (v2 << ((v3) & 63))
                    _ = fi9
                }
            }
        } else {
            fmt.Println(v3, v2)
            if false {
                fmt.Println(v2)
                var v11 uint = (v3 - 1)
                v12 := []int{0, -530}
                _ = v11
                _ = v12
            }
        }
    default:
        for fi13 := 0; fi13 < 15; fi13++ {
            var v14 uint = v3
            _ = fi13
            _ = v14
        }
    }
    delete(v1, 956)
    fmt.Println(v2)
    var v15 bool = true
    for fi16 := 0; fi16 < 15; fi16++ {
        for fi17 := 0; fi17 < 6; fi17++ {
            var v18 float32 = ((0.0 * (0.1 - ((float32((-2.25 * 1.5)) / 0.1) - -2.25))) - (1.5 * 0.0))
            _ = fi17
            _ = v18
        }
        if ((v3 > 255) && v15) {
            v19 := GenS2{F0: 24, F1: 255, F2: 204}
            _ = v19
        } else {
            switch "foo bar" {
            case "foo bar":
                switch fi16 {
                case 516:
                    if (0 >= 1) {
                        v20 := GenS1{F0: 2, F1: -1, F2: 127}
                        _ = v20
                    } else {
                        v2 -= (v2 & -32768)
                        v21 := GenS2{F0: -1, F1: 2, F2: 60}
                        _ = v21
                    }
                case 606:
                    for fi22 := 0; fi22 < 1; fi22++ {
                        var v23 int64 = 924
                        _ = fi22
                        _ = v23
                    }
                default:
                    switch fi16 {
                    case 32768:
                        var v24 int32 = 256
                        _ = v24
                    case 462:
                        v3 -= 18446744073709551615
                    default:
                        var v25 int8 = -59
                        _ = v25
                    }
                }
            case "\n\t":
                if true {
                    v26 := []int{-361, (fi16 << ((v3) & 63))}
                    _ = v26
                }
            default:
                if (748 >= 816) {
                    var v27 int64 = -1
                    _ = v27
                }
            }
        }
        _ = fi16
    }
    v28 := GenS3{F0: 2, F1: 915}
    switch "a" {
    case "a":
        v1[128] = 1
    }
    if (true || (!v15)) {
        v2 = (v2 & (v2 + (v2 / (((v2 >> (((v3 >> ((v3) & 63))) & 63))) | 1))))
        delete(v1, 754)
    } else {
        v2 -= 355
        var v29 int = -2147483648
        _ = v29
    }
    v1[4294967295] = -1
    var v30 bool = (-583 >= -2147483648)
    if v15 {
        for fi31 := 0; fi31 < 4; fi31++ {
            v32 := []int{(-(fi31))}
            _ = fi31
            _ = v32
        }
    }
    fmt.Println(v3, v30, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v15
    _ = v2
    _ = v28
    _ = v3
    _ = v30
}
