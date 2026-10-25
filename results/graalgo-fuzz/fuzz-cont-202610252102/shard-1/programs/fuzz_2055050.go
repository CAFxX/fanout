package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 string; F2 bool }
func genFn0() int8 {
    var v4 int = 87
    v5 := []int{9223372036854775807, v4, (v4 | v4)}
    v5 = append(v5, 25)
    _ = v4
    _ = v5
    return 7
}
func genFn1(p0 uint32, p1 int8) uint8 {
    for fi6 := 0; fi6 < 15; fi6++ {
        if false {
            if false {
                for fi7 := 0; fi7 < 15; fi7++ {
                    v8 := GenS1{F0: 1000, F1: 1}
                    fmt.Println(p1, fi7)
                    _ = fi7
                    _ = v8
                }
                if true {
                    var v9 uint8 = 255
                    if false {
                        p0 -= (p0 & p0)
                    }
                    _ = v9
                } else {
                    p1 = (^(p1))
                }
                p0 -= p0
            } else {
                if ((32767 != 88) || ((3.14159 * 1.5) != 3.14159)) {
                    v10 := GenS1{F0: -228, F1: 625}
                    p0 -= (^(p0))
                    v11 := []int{}
                    _ = v10
                    _ = v11
                }
                switch "hello" {
                case "hello":
                    v12 := GenS1{F0: -1, F1: 65535}
                    _ = v12
                case "":
                    fmt.Println(p1, fi6)
                default:
                    fmt.Println(p0, p1, p1)
                }
            }
            var v13 uint64 = 702
            if false {
                p0 = (p0 | (p0 | p0))
                for fi14 := 0; fi14 < 5; fi14++ {
                    v15 := GenS2{F0: 2147483647}
                    _ = fi14
                    _ = v15
                }
                v16 := []int{128, -820}
                _ = v16
            } else {
                var v17 int8 = p1
                var v18 bool = ((1.5 / 1.5) != (float64((-2.25 - 1e10)) - 1.5))
                v19 := GenS3{F0: (fi6 | fi6), F1: (("\n\t" + (("hello" + "\n\t") + ("\n\t" + "foo bar"))) + "a"), F2: v18}
                _ = v17
                _ = v18
                _ = v19
            }
            _ = v13
        } else {
            if (!false) {
                fmt.Println(p0)
                p1 += 3
            }
        }
        switch "a" {
        case "":
            for fi20 := 0; fi20 < 1; fi20++ {
                fmt.Println(p0)
                v21 := []int{1, fi20}
                fmt.Println(fi6, p1, p0)
                _ = fi20
                _ = v21
            }
        case "foo bar":
            v22 := []int{0, (-(fi6)), 32767}
            _ = v22
        case "\n\t":
            v23 := GenS3{F0: -2147483648, F1: "", F2: (0.0 > (-2.25 * (0.0 + (float64(float32(float64(((-2.25 * 1.5) + 0.1)))) - 0.1))))}
            _ = v23
        default:
            switch fi6 {
            case -2147483648:
                v24 := GenS3{F0: 960, F1: "\n\t", F2: false}
                _ = v24
            case 4294967295:
                var v25 float64 = (0.1 - (1e10 * float64((1.5 * 0.1))))
                _ = v25
            default:
                var v26 int16 = 42
                _ = v26
            }
        }
        _ = fi6
    }
    for fi27 := 0; fi27 < 17; fi27++ {
        p0 = 65535
        _ = fi27
    }
    if true {
        switch 32767 {
        case 32768:
            if ((p0 / ((p0) | 1)) > 525) {
                p0 -= (p0 << ((255) & 63))
            }
        default:
            v28 := GenS3{F0: 32768, F1: "hello", F2: (!false)}
            _ = v28
        }
        p1 = genFn0()
    } else {
        v29 := GenS1{F0: -133, F1: 32767}
        _ = v29
    }
    v30 := []int{-289, 127, -265}
    var v31 uint16 = 828
    _ = p0
    _ = p1
    _ = v30
    _ = v31
    return uint8(v31)
}

func main() {
    var v1 uint = 65535
    v2 := make(map[string]bool)
    var v3 uint32 = 2147483647
    v3 = v3
    v3 = (^(v3))
    var v32 int8 = genFn0()
    v33 := []int{}
    switch "hello" {
    case "":
        for fi34 := 0; fi34 < 19; fi34++ {
            v3 = 4294967295
            var v35 int8 = (-(v32))
            _ = fi34
            _ = v35
        }
    default:
        v32 = (v32 ^ (v32 << ((v1) & 63)))
    }
    var v36 uint = v1
    if false {
        v37 := GenS1{F0: 42, F1: 65535}
        var v38 uint64 = 414
        if len(v33) >= 2 {
            v33 = v33[1:len(v33)]
        }
        _ = v37
        _ = v38
    }
    var v39 uint = v36
    v39 -= (v39 * uint(v3))
    fmt.Println(v32, v39, v32)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v36
    _ = v39
}
