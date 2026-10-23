package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 int; F2 string }
func genFn0(p0 uint8) int8 {
    var v4 int8 = 0
    var v5 int16 = 1
    p0 = (^(p0))
    var v6 uint32 = 4294967295
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 127
}
func genFn1(p0 uint64) uint8 {
    if (-128 < -1) {
        var v7 int64 = 752
        if (((false || (!false)) || false) && false) {
            if (!false) {
                for fi8 := 0; fi8 < 9; fi8++ {
                    v9 := GenS1{F0: 58}
                    var v10 int32 = 1
                    v11 := GenS1{F0: 255}
                    _ = fi8
                    _ = v10
                    _ = v11
                    _ = v9
                }
                fmt.Println(v7, v7)
                for fi12 := 0; fi12 < 5; fi12++ {
                    switch "" {
                    case "hello":
                        v7 = 42
                    default:
                        fmt.Println(v7, v7)
                    }
                    _ = fi12
                }
            }
            var v13 string = "foo bar"
            switch "foo bar" {
            case "foo bar":
                var v14 float64 = (-2.25 * -2.25)
                _ = v14
            case "hello":
                var v15 int16 = 1000
                _ = v15
            case "":
                v16 := GenS2{F0: 389, F1: 0, F2: v13}
                _ = v16
            }
            _ = v13
        }
        _ = v7
    } else {
        var v17 uint64 = p0
        switch 454 {
        case -9223372036854775808:
            if true {
                var v18 int64 = 719
                v17 += p0
                v17 = v17
                _ = v18
            } else {
                v19 := GenS2{F0: 42, F1: -2147483648, F2: ""}
                for fi20 := 0; fi20 < 12; fi20++ {
                    v21 := GenS1{F0: 127}
                    v22 := []int{32768, fi20, (fi20 * fi20), fi20}
                    var v23 int32 = -1
                    _ = fi20
                    _ = v21
                    _ = v22
                    _ = v23
                }
                _ = v19
            }
        case -862:
            if false {
                v24 := GenS2{F0: 2, F1: 294, F2: ""}
                v25 := []int{9223372036854775807}
                p0 = v17
                _ = v24
                _ = v25
            } else {
                p0 += v17
            }
        default:
            p0 += (p0 >> ((65535) & 63))
        }
        v26 := []int{-413}
        _ = v17
        _ = v26
    }
    for fi27 := 0; fi27 < 18; fi27++ {
        v28 := []int{}
        _ = fi27
        _ = v28
    }
    fmt.Println(p0)
    v29 := []int{-9223372036854775808, 101}
    if (true || false) {
        fmt.Println(p0)
        v30 := GenS1{F0: 255}
        var v31 float32 = -2.25
        _ = v30
        _ = v31
    }
    _ = p0
    _ = v29
    return 209
}

func main() {
    var v1 int8 = -59
    var v2 bool = (0.0 == (1.5 * float64(3.14159)))
    v3 := make(map[int]string)
    var v32 int32 = 48
    v33 := GenS2{F0: 834, F1: 2147483647, F2: "a"}
    switch 707 {
    case 32768:
        v2 = (!(32767 < -124))
    }
    if (!v2) {
        v34 := GenS2{F0: 255, F1: 256, F2: "a"}
        _ = v34
    }
    var v35 int16 = 959
    fmt.Println(v1, v32)
    var v36 int = 4294967295
    var v37 int16 = (v35 << ((255) & 63))
    v37 -= (v37 - v37)
    switch 4294967295 {
    case 2147483647:
        var v38 uint16 = 2
        _ = v38
    case -1:
        var v39 uint8 = 42
        _ = v39
    }
    v32 = (v32 | v32)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v35
    _ = v36
    _ = v37
}
