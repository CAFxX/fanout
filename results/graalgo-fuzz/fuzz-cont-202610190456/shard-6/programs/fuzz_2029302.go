package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8; F1 string; F2 int8 }
type GenS5 struct { F0 uint; F1 uint }
type GenS6 struct { F0 bool; F1 int32 }
type GenS7 struct { F0 uint32; F1 uint8; F2 int32 }
type GenS8 struct { F0 int32; F1 int8; F2 int32 }
type GenS9 struct { F0 uint16; F1 int16; F2 int16 }
func genFn0(p0 uint64, p1 bool) uint {
    v6 := []int{127}
    switch 841 {
    case 4294967295:
        fmt.Println(p0)
    case 445:
        v7 := GenS1{F0: -176, F1: p1}
        _ = v7
    case 459:
        switch 127 {
        case -1:
            if p1 {
                var v8 int64 = -938
                for fi9 := 0; fi9 < 1; fi9++ {
                    if (-2.25 > float32(float64(-2.25))) {
                        v10 := GenS2{F0: 303, F1: p0, F2: 65535}
                        v8 += v8
                        _ = v10
                    }
                    _ = fi9
                }
                fmt.Println(p0, v8)
                _ = v8
            }
        default:
            var v11 int64 = -1
            _ = v11
        }
    default:
        switch -2147483648 {
        case 208:
            p1 = p1
        case 9223372036854775807:
            v12 := GenS2{F0: -252, F1: 480, F2: 1}
            _ = v12
        default:
            if true {
                for fi13 := 0; fi13 < 17; fi13++ {
                    _ = len(v6)
                    _ = fi13
                }
                v14 := GenS3{F0: -1}
                for fi15 := 0; fi15 < 8; fi15++ {
                    switch int(p0) {
                    case 128:
                        p1 = false
                    case 32767:
                        p0 -= (p0 * (p0 - (^(p0))))
                    default:
                        v16 := GenS4{F0: 42, F1: "", F2: 42}
                        _ = v16
                    }
                    _ = fi15
                }
                _ = v14
            } else {
                for fi17 := 0; fi17 < 18; fi17++ {
                    _ = len(v6)
                    _ = fi17
                }
                fmt.Println(p1, p0)
            }
        }
    }
    v18 := GenS5{F0: 42, F1: 974}
    fmt.Println(p0)
    v6 = append(v6, 32768)
    _ = p0
    _ = p1
    _ = v18
    _ = v6
    return 18446744073709551615
}

func main() {
    var v1 bool = (75 >= -1)
    v2 := make(map[uint32]bool)
    v3 := make(map[string]int)
    var v4 float64 = 1e10
    var v5 uint32 = 65535
    if ((1000 >= 1000) || v1) {
        var v19 bool = v1
        v20 := GenS6{F0: (v19 || v19), F1: -527}
        _ = v19
        _ = v20
    } else {
        switch 32768 {
        case 32767:
            fmt.Println(v4, v1, v1)
        case 329:
            var v21 float64 = ((v4 * (float64(float32(v4)) - ((float64(((-2.25 - 3.14159) * -2.25)) - v4) + ((v4 * v4) / v4)))) * v4)
            _ = v21
        case 715:
            var v22 int8 = 100
            _ = v22
        default:
            v4 -= v4
        }
    }
    var v23 string = (("hello" + "foo bar") + "a")
    var v24 float32 = 3.14159
    v25 := GenS7{F0: v5, F1: 245, F2: 1}
    switch -2147483648 {
    case 9223372036854775807:
        if false {
            v26 := GenS8{F0: -342, F1: 1, F2: 0}
            _ = v26
        }
    case 939:
        var v27 uint64 = 375
        _ = v27
    case 447:
        v23 = ((("\n\t" + "") + (("hello" + "\n\t") + "foo bar")) + v23)
    default:
        v4 = 1e10
    }
    v4 += ((v4 + v4) + (1e10 - v4))
    v28 := GenS2{F0: 0, F1: 861, F2: 867}
    for fi29 := 0; fi29 < 4; fi29++ {
        switch v23 {
        case "":
            v30 := GenS1{F0: int64(fi29), F1: false}
            _ = v30
        case "\n\t":
            v31 := GenS2{F0: 1000, F1: 927, F2: 2}
            _ = v31
        case "hello":
            var v32 int16 = 32767
            _ = v32
        }
        if (v1 && ((-(fi29)) < fi29)) {
            var v33 float32 = v24
            _ = v33
        } else {
            v34 := []int{int(v5), 1, fi29, (-(fi29))}
            v35 := GenS4{F0: 0, F1: ("hello" + "hello"), F2: -58}
            _ = v34
            _ = v35
        }
        _ = fi29
    }
    var v36 float64 = float64(0.1)
    v37 := GenS9{F0: 588, F1: 172, F2: 1000}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
