package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 int64; F1 int8; F2 bool }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint; F1 uint; F2 uint8 }
type GenS6 struct { F0 uint16 }
func genFn0(p0 uint32, p1 uint16) int8 {
    for fi5 := 0; fi5 < 18; fi5++ {
        v6 := GenS1{F0: p1, F1: 780, F2: 1}
        _ = fi5
        _ = v6
    }
    p0 = (^(p0))
    var v7 bool = (false || (!true))
    _ = p0
    _ = p1
    _ = v7
    return 22
}
func genFn1(p0 uint16, p1 uint8) int {
    if (2147483647 <= 4294967295) {
        if (65535 != 0) {
            v8 := GenS1{F0: p0, F1: 0, F2: -1}
            _ = v8
        }
        v9 := GenS2{F0: -108, F1: -128}
        var v10 float32 = 3.14159
        _ = v10
        _ = v9
    }
    var v11 bool = false
    var v12 bool = true
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 9223372036854775807
}

func main() {
    var v1 int8 = -1
    v2 := make(map[int]string)
    var v3 float64 = 3.14159
    v4 := make(map[uint32]int)
    v1 -= v1
    if false {
        v13 := GenS1{F0: 0, F1: 979, F2: 0}
        for fi14 := 0; fi14 < 1; fi14++ {
            var v15 bool = (((float64((float32(v3) * (0.0 * float32((v3 - -2.25))))) > v3) && (522 >= 697)) && ((-2.25 / (((float64((1.5 + 0.0)) + (v3 - 0.0)) - (v3 / 3.14159)) + 1e10)) > -2.25))
            if (!((-2.25 + (((v3 + 1e10) / v3) + v3)) == (v3 * (1.5 * 3.14159)))) {
                switch fi14 {
                case 32768:
                    var v16 float64 = v3
                    _ = v16
                }
            }
            for fi17 := 0; fi17 < 3; fi17++ {
                var v18 float32 = 0.0
                _ = fi17
                _ = v18
            }
            _ = fi14
            _ = v15
        }
        _ = v13
    }
    switch "foo bar" {
    case "a":
        v19 := GenS1{F0: 65535, F1: 948, F2: 0}
        _ = v19
    case "foo bar":
        v1 -= -74
    }
    for fi20 := 0; fi20 < 10; fi20++ {
        v21 := v4[0]
        for fi22 := 0; fi22 < 11; fi22++ {
            v21 = (-(fi20))
            fmt.Println(v1)
            _ = fi22
        }
        if (!(42 != 65535)) {
            var v23 uint32 = 42
            _ = v23
        }
        _ = fi20
        _ = v21
    }
    switch -683 {
    case 32768:
        if false {
            v1 += genFn0(2147483647, 472)
            v24 := GenS3{F0: 1, F1: (v1 - v1), F2: (false && true)}
            _ = v24
        } else {
            v3 += (((v3 / v3) + ((0.1 + (v3 / (v3 + 0.0))) * (v3 * v3))) * (((0.0 / (0.0 * v3)) / v3) - (1e10 * v3)))
            v25 := GenS3{F0: -454, F1: 127, F2: false}
            var v26 int16 = 376
            _ = v25
            _ = v26
        }
    case 9223372036854775807:
        fmt.Println(v3, v1)
    default:
        v3 += 1.5
    }
    v27 := GenS4{F0: 9223372036854775807}
    if (!(!((("hello" + ("a" + "")) + "foo bar") == ""))) {
        if true {
            var v28 uint = 0
            _ = v28
        } else {
            for fi29 := 0; fi29 < 16; fi29++ {
                if true {
                    var v30 uint64 = 18446744073709551615
                    v31 := GenS2{F0: 1, F1: genFn0(457, 855)}
                    var v32 int16 = 1
                    _ = v30
                    _ = v31
                    _ = v32
                }
                _ = fi29
            }
            v33 := GenS4{F0: -2147483648}
            _ = v33
        }
        v34 := GenS1{F0: 811, F1: 4294967295, F2: 1}
        if ((!true) || false) {
            var v35 uint32 = 1
            var v36 int64 = 42
            _ = v35
            _ = v36
        } else {
            if ((1e10 <= (((1e10 + 0.1) - 1.5) * (float32((1e10 - 3.14159)) / 3.14159))) && ((2 != 1) && true)) {
                v37 := GenS5{F0: 1, F1: 42, F2: 128}
                _ = v37
            } else {
                for fi38 := 0; fi38 < 17; fi38++ {
                    var v39 float32 = -2.25
                    _ = fi38
                    _ = v39
                }
                switch 32768 {
                case 121:
                    v40 := GenS5{F0: 433, F1: 929, F2: 0}
                    _ = v40
                case -1:
                    var v41 uint64 = 42
                    _ = v41
                case -684:
                    var v42 uint = 42
                    _ = v42
                }
            }
        }
        _ = v34
    }
    v43 := GenS6{F0: 65535}
    var v44 float32 = ((float32(1.5) * (3.14159 / 1e10)) * 0.0)
    var v45 uint32 = 754
    fmt.Println(v1, v3, v1)
    var v46 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
}
