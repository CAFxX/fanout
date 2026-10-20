package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 uint; F1 int; F2 int16 }
type GenS3 struct { F0 uint32; F1 uint64; F2 string }
type GenS4 struct { F0 int16; F1 int16 }
type GenS5 struct { F0 uint32; F1 int64 }
func genFn0(p0 bool, p1 int8) uint16 {
    for fi6 := 0; fi6 < 1; fi6++ {
        if ((-2.25 / float32((((-2.25 * (((float64(1.5) - (0.1 * 1e10)) + (1.5 + 3.14159)) * -2.25)) - (-2.25 + (((1.5 + (3.14159 + -2.25)) / 0.1) / -2.25))) - float64((3.14159 + -2.25))))) != (-2.25 * float32((-2.25 * 3.14159)))) {
            fmt.Println(p0, fi6, p1)
        } else {
            var v7 uint = 255
            _ = v7
        }
        for fi8 := 0; fi8 < 10; fi8++ {
            if ((!p0) || (-2.25 == (((0.1 / 1.5) * 1e10) / 1.5))) {
                fmt.Println(p1)
                fmt.Println(fi8, fi8, p1)
            } else {
                var v9 uint64 = 480
                _ = v9
            }
            v10 := GenS1{F0: -110, F1: 42}
            switch ("foo bar" + "") {
            case "a":
                switch "" {
                case "":
                    p1 += (p1 ^ (^(p1)))
                case "a":
                    switch "" {
                    case "":
                        v11 := GenS1{F0: (p1 + (p1 | p1)), F1: 18446744073709551615}
                        _ = v11
                    case "foo bar":
                        if false {
                            v12 := []int{(fi6 % ((fi6) | 1)), fi6}
                            var v13 uint16 = 32767
                            v14 := GenS2{F0: 933, F1: (^(fi6)), F2: 42}
                            _ = v12
                            _ = v13
                            _ = v14
                        }
                    default:
                        var v15 int = (^(fi8))
                        _ = v15
                    }
                case "foo bar":
                    v16 := []int{}
                    v16 = append(v16, fi8)
                    _ = v16
                default:
                    p0 = (p0 && (127 < -287))
                    p1 -= (p1 - p1)
                }
            }
            _ = fi8
            _ = v10
        }
        _ = fi6
    }
    v17 := GenS3{F0: 555, F1: 18446744073709551615, F2: ""}
    for fi18 := 0; fi18 < 11; fi18++ {
        v19 := GenS4{F0: 1, F1: 1000}
        _ = fi18
        _ = v19
    }
    v20 := []int{}
    _ = p0
    _ = p1
    _ = v17
    _ = v20
    return 791
}

func main() {
    v1 := make(map[string]int)
    var v2 int8 = 127
    var v3 int32 = 2147483647
    var v4 uint16 = 1
    v5 := make(map[int]int)
    for fi21 := 0; fi21 < 1; fi21++ {
        var v22 uint32 = 2147483648
        v4 = (v4 * v4)
        _ = fi21
        _ = v22
    }
    v23 := v5[42]
    var v24 int8 = v2
    v23 = 127
    v4 += (v4 + (v4 ^ genFn0((!true), 127)))
    var v25 uint64 = 42
    v2 = 1
    var v26 uint64 = v25
    var v27 int8 = (-(v2))
    if false {
        v3 = 0
        fmt.Println(v27, v25, v23)
    } else {
        var v28 uint8 = 0
        var v29 float32 = 1e10
        _ = v28
        _ = v29
    }
    if true {
        var v30 float64 = 1.5
        _ = v30
    } else {
        var v31 int = (v23 * (v23 * 128))
        for fi32 := 0; fi32 < 20; fi32++ {
            if true {
                var v33 int = fi32
                _ = v33
            }
            v34 := GenS1{F0: v24, F1: (^(v26))}
            var v35 float64 = (1.5 + 1.5)
            _ = fi32
            _ = v34
            _ = v35
        }
        _ = v31
    }
    v36 := GenS5{F0: 2147483647, F1: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
