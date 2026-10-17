package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int8 }
type GenS2 struct { F0 int8; F1 uint32; F2 uint64 }
type GenS3 struct { F0 int16; F1 uint64; F2 int8 }
type GenS4 struct { F0 uint16; F1 bool }
func genFn0(p0 int8, p1 uint8, p2 uint) uint32 {
    switch 0 {
    case 9223372036854775807:
        var v4 string = "hello"
        _ = v4
    case 1:
        var v5 float64 = -2.25
        _ = v5
    case 42:
        fmt.Println(p1, p1)
    }
    if (false && true) {
        v6 := []int{-72, 42, 128, 30}
        var v7 int32 = -599
        _ = v6
        _ = v7
    } else {
        var v8 int16 = -366
        v9 := GenS1{F0: 65535, F1: 2147483648, F2: -128}
        if false {
            if (!true) {
                var v10 string = "foo bar"
                var v11 bool = false
                for fi12 := 0; fi12 < 10; fi12++ {
                    v11 = (v10 > v10)
                    _ = fi12
                }
                _ = v10
                _ = v11
            } else {
                var v13 int8 = p0
                var v14 int8 = p0
                v13 = (p0 << ((p2) & 63))
                _ = v13
                _ = v14
            }
            for fi15 := 0; fi15 < 3; fi15++ {
                var v16 float64 = 0.1
                _ = fi15
                _ = v16
            }
            v17 := []int{-881, -100}
            _ = v17
        } else {
            if (!(((1e10 * -2.25) * 1.5) >= (1e10 - (0.0 / (3.14159 / 1.5))))) {
                var v18 int8 = 123
                switch 32768 {
                case -2147483648:
                    var v19 int32 = 1
                    _ = v19
                case 32768:
                    p1 -= p1
                case -9223372036854775808:
                    v20 := GenS2{F0: p0, F1: uint32(p1), F2: 42}
                    _ = v20
                default:
                    for fi21 := 0; fi21 < 2; fi21++ {
                        v22 := GenS3{F0: int16(v18), F1: 18446744073709551615, F2: p0}
                        if true {
                            fmt.Println(v8, p0, p1)
                            p0 = (p0 / (((v18 - (v18 - (-(v18))))) | 1))
                        }
                        _ = fi21
                        _ = v22
                    }
                }
                _ = v18
            }
        }
        _ = v8
        _ = v9
    }
    v23 := []int{}
    if true {
        p1 += p1
    } else {
        if (false && ("foo bar" == "")) {
            p2 = 1
            if (!(42 <= 0)) {
                p1 = p1
                v23 = append(v23, 42)
            }
        }
        fmt.Println(p0)
        v24 := GenS1{F0: 603, F1: 0, F2: 100}
        _ = v24
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    return 2
}
func genFn1() int32 {
    v25 := []int{}
    if false {
        _ = len(v25)
    } else {
        if true {
            var v26 float32 = ((-2.25 - (1e10 + ((-2.25 - (-2.25 * 0.0)) / (((0.1 - -2.25) - ((-2.25 / ((0.0 * 3.14159) + -2.25)) - 0.0)) * (1.5 - -2.25))))) / 3.14159)
            if (2 == 32767) {
                _ = len(v25)
            } else {
                var v27 bool = true
                fmt.Println(v26)
                for fi28 := 0; fi28 < 16; fi28++ {
                    for fi29 := 0; fi29 < 18; fi29++ {
                        if len(v25) >= 2 {
                            v25 = v25[1:len(v25)]
                        }
                        v27 = v27
                        _ = fi29
                    }
                    v26 = float32(1.5)
                    _ = fi28
                }
                _ = v27
            }
            _ = v26
        } else {
            var v30 float64 = ((1e10 * -2.25) / -2.25)
            switch ("hello" + ("" + ("foo bar" + ""))) {
            case "hello":
                v30 = (v30 - ((v30 / v30) - 0.0))
            case "a":
                v30 -= (0.0 + v30)
            default:
                v31 := GenS3{F0: -241, F1: 397, F2: -21}
                _ = v31
            }
            _ = v30
        }
        var v32 float32 = 0.0
        _ = v32
    }
    if len(v25) > 0 {
        v25[0] = 9223372036854775807
    }
    _ = v25
    return 0
}

func main() {
    var v1 uint16 = 65535
    var v2 int64 = 813
    v3 := make(map[string]uint32)
    var v33 int8 = 48
    v2 = (v2 | (v2 - -266))
    v34 := GenS4{F0: 413, F1: true}
    for fi35 := 0; fi35 < 15; fi35++ {
        var v36 int32 = int32(v2)
        _ = fi35
        _ = v36
    }
    var v37 uint = 252
    fmt.Println(v33, v33, v2)
    for fi38 := 0; fi38 < 13; fi38++ {
        v33 -= (-(v33))
        if ((495 <= 1) && (true && true)) {
            var v39 float64 = ((0.1 / 1e10) - (1.5 - 1.5))
            _ = v39
        }
        if (true || (false && true)) {
            v40 := GenS1{F0: 0, F1: 105, F2: 42}
            v2 = -1
            _ = v40
        } else {
            v2 += -1
            v33 -= -128
            fmt.Println(v2, v33)
        }
        _ = fi38
    }
    fmt.Println(v37)
    for fi41 := 0; fi41 < 20; fi41++ {
        v42 := GenS3{F0: -32768, F1: 464, F2: v33}
        var v43 uint64 = 0
        _ = fi41
        _ = v42
        _ = v43
    }
    switch 396 {
    case 42:
        v44 := []int{42, 9223372036854775807, 1, -642}
        _ = v44
    }
    fmt.Println(v1, v33, v1)
    v37 = v37
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v37
}
