package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 int; F1 uint16; F2 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint }
func genFn0() int8 {
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 uint16 = 65535
        v8 := GenS1{F0: 42, F1: ("a" + "")}
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 int16 = -423
    var v10 uint8 = 229
    v11 := []int{127, 360, 0}
    _ = v10
    _ = v11
    _ = v9
    return 115
}
func genFn1() string {
    v12 := []int{9223372036854775807, 127}
    if false {
        v13 := GenS2{F0: 127, F1: "foo bar"}
        _ = v13
    } else {
        var v14 bool = ((1e10 * 3.14159) >= ((((float32(3.14159) + float32(float64((float32(-2.25) * 1.5)))) - -2.25) + 1.5) / 1.0))
        v14 = (true && (!false))
        _ = v14
    }
    _ = len(v12)
    if true {
        v15 := GenS3{F0: 1, F1: 65535, F2: 1}
        _ = v15
    }
    v16 := GenS2{F0: 2147483647, F1: ""}
    _ = v12
    _ = v16
    return "a"
}

func main() {
    var v1 float64 = ((3.14159 + (1e10 + -2.25)) / ((float64(float32(-2.25)) - (1e10 * 1e10)) - ((((1.5 / 1.5) + 0.0) - -2.25) + (3.14159 - (-2.25 * 1e10)))))
    var v2 int16 = 177
    var v3 int64 = -849
    v4 := []int{}
    var v5 uint64 = 42
    if true {
        v3 = (v3 | (v3 & 42))
        var v17 int64 = (v3 >> ((18446744073709551615) & 63))
        _ = v17
    }
    var v18 int64 = v3
    fmt.Println(v18, v2)
    fmt.Println(v18, v2, v1)
    var v19 float32 = 0.1
    if false {
        if (32768 < 42) {
            for fi20 := 0; fi20 < 12; fi20++ {
                fmt.Println(fi20, v2)
                fmt.Println(v1, fi20, v19)
                _ = fi20
            }
            switch "hello" {
            case "a":
                switch "foo bar" {
                case "foo bar":
                    if len(v4) > 0 {
                        v4[0] = 2147483647
                    }
                }
            case "foo bar":
                switch 0 {
                case 454:
                    for fi21 := 0; fi21 < 17; fi21++ {
                        var v22 int = fi21
                        _ = fi21
                        _ = v22
                    }
                case -1:
                    v5 = (v5 ^ v5)
                }
            }
            if false {
                v23 := GenS3{F0: 361, F1: 26, F2: -99}
                fmt.Println(v19, v2)
                v5 += v5
                _ = v23
            } else {
                var v24 uint64 = 378
                fmt.Println(v24, v5, v3)
                switch 178 {
                case 127:
                    var v25 uint16 = 42
                    _ = v25
                default:
                    v3 += (v18 & int64(v5))
                }
                _ = v24
            }
        } else {
            v18 -= 531
            v3 += v18
        }
        for fi26 := 0; fi26 < 15; fi26++ {
            for fi27 := 0; fi27 < 12; fi27++ {
                v5 += (v5 / (((v5 << ((255) & 63))) | 1))
                _ = fi27
            }
            _ = fi26
        }
        v5 = v5
    } else {
        v28 := GenS3{F0: 1, F1: 64, F2: 100}
        _ = v28
    }
    for fi29 := 0; fi29 < 18; fi29++ {
        fmt.Println(v5)
        if true {
            if len(v4) > 0 {
                v4[0] = fi29
            }
            switch ("\n\t" + "") {
            case "foo bar":
                if true {
                    var v30 uint64 = 4294967295
                    _ = v30
                } else {
                    fmt.Println(v5)
                }
            case "":
                for fi31 := 0; fi31 < 1; fi31++ {
                    var v32 bool = true
                    v1 += (0.0 - v1)
                    var v33 float32 = (v19 - ((v19 * 1e10) - (v19 + (0.0 / (v19 * v19)))))
                    _ = fi31
                    _ = v32
                    _ = v33
                }
            case "\n\t":
                var v34 uint8 = 2
                _ = v34
            default:
                var v35 bool = false
                _ = v35
            }
            v5 += 1
        } else {
            _ = len(v4)
        }
        v5 -= 1
        _ = fi29
    }
    var v36 int = 210
    var v37 bool = ((2 == 4294967295) && false)
    v38 := GenS4{F0: 100}
    v39 := GenS5{F0: 578}
    for fi40 := 0; fi40 < 5; fi40++ {
        v18 -= 20
        _ = fi40
    }
    v4 = append(v4, v36)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
