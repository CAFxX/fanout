package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 int8; F1 uint8; F2 int16 }
type GenS3 struct { F0 int8; F1 int64; F2 bool }
type GenS4 struct { F0 uint16; F1 uint32; F2 uint }
type GenS5 struct { F0 int16 }
func genFn0() int8 {
    var v6 float64 = -2.25
    var v7 int32 = -380
    v6 = v6
    var v8 uint8 = 0
    _ = v6
    _ = v7
    _ = v8
    return -128
}
func genFn1(p0 uint8) int64 {
    if (false || (false || (true || (-758 <= -9223372036854775808)))) {
        var v9 int = 878
        fmt.Println(v9)
        _ = v9
    } else {
        var v10 uint32 = 0
        switch "a" {
        case "\n\t":
            var v11 int = -330
            _ = v11
        case "hello":
            for fi12 := 0; fi12 < 6; fi12++ {
                var v13 int = fi12
                _ = fi12
                _ = v13
            }
        default:
            var v14 uint = 266
            _ = v14
        }
        v15 := GenS1{F0: ("" + ("hello" + "hello")), F1: 255}
        _ = v10
        _ = v15
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        var v17 int16 = 42
        v18 := GenS2{F0: genFn0(), F1: p0, F2: (v17 / ((-1) | 1))}
        _ = fi16
        _ = v17
        _ = v18
    }
    v19 := []int{-2147483648, 32767}
    for fi20 := 0; fi20 < 20; fi20++ {
        var v21 float32 = (1e10 - 0.0)
        _ = fi20
        _ = v21
    }
    _ = p0
    _ = v19
    return 42
}

func main() {
    var v1 uint64 = 731
    var v2 bool = (-2.25 != 1.5)
    var v3 int = 108
    var v4 float64 = 0.1
    v5 := make(map[string]int)
    v22 := GenS2{F0: -73, F1: 235, F2: 42}
    fmt.Println(v3, v2)
    switch "hello" {
    case "\n\t":
        v4 -= v4
    case "a":
        var v23 uint16 = 65535
        _ = v23
    default:
        var v24 uint16 = 886
        _ = v24
    }
    for fi25 := 0; fi25 < 16; fi25++ {
        v26 := GenS3{F0: 100, F1: -9223372036854775808, F2: true}
        _ = fi25
        _ = v26
    }
    if false {
        for fi27 := 0; fi27 < 14; fi27++ {
            for fi28 := 0; fi28 < 4; fi28++ {
                switch v3 {
                case 1:
                    fmt.Println(v1)
                case 4294967295:
                    var v29 uint64 = v1
                    var v30 int32 = -69
                    _ = v29
                    _ = v30
                case -1:
                    v31 := GenS4{F0: 32767, F1: 1, F2: 255}
                    _ = v31
                }
                v3 += 128
                _ = fi28
            }
            _ = fi27
        }
        fmt.Println(v2, v2, v3)
        fmt.Println(v3, v4)
    }
    v32 := GenS5{F0: -1}
    var v33 string = ((("\n\t" + "\n\t") + "foo bar") + "foo bar")
    if v2 {
        var v34 float32 = 1.5
        _ = v34
    }
    var v35 uint = 2
    var v36 string = (((v33 + (((v33 + "a") + v33) + v33)) + v33) + v33)
    var v37 bool = ((((v35 * 18446744073709551615) >= (v35 ^ (^(v35)))) && true) || v2)
    var v38 uint = (v35 + (v35 | v35))
    v39 := v5[(v33 + ("" + ("\n\t" + v36)))]
    for fi40 := 0; fi40 < 11; fi40++ {
        v41 := GenS4{F0: 2, F1: 0, F2: (^(v38))}
        _ = fi40
        _ = v41
    }
    v42 := []int{(v3 % (((v3 % ((v3) | 1))) | 1)), v39, -9223372036854775808}
    v36 = v33
    v43 := GenS5{F0: 1}
    var v44 int64 = genFn1(127)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v32
    _ = v33
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
}
