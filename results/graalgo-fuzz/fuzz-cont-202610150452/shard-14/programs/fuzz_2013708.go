package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 int32 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 bool; F1 uint16; F2 int32 }
type GenS4 struct { F0 bool; F1 uint16; F2 int }
func genFn0(p0 bool) uint32 {
    var v5 uint16 = 268
    fmt.Println(v5, v5)
    v6 := GenS1{F0: true, F1: true, F2: -1}
    _ = p0
    _ = v5
    _ = v6
    return 0
}
func genFn1() uint16 {
    switch ("" + ("hello" + "a")) {
    case "hello":
        v7 := GenS2{F0: 2147483647, F1: 1}
        _ = v7
    default:
        var v8 uint = 1
        _ = v8
    }
    if ((!true) && (401 == 1)) {
        var v9 uint = 4294967295
        v9 += (v9 * (v9 & v9))
        var v10 uint32 = 65535
        _ = v10
        _ = v9
    }
    v11 := GenS3{F0: true, F1: 32767, F2: -963}
    var v12 uint8 = 1
    if (810 < 4294967295) {
        for fi13 := 0; fi13 < 8; fi13++ {
            var v14 string = "\n\t"
            _ = fi13
            _ = v14
        }
    } else {
        v15 := []int{-1, -2147483648, 9223372036854775807}
        v12 -= v12
        _ = v15
    }
    _ = v11
    _ = v12
    return 65535
}

func main() {
    var v1 uint8 = 30
    var v2 float64 = (0.1 / 1e10)
    v3 := make(map[uint32]string)
    var v4 uint64 = 964
    for fi16 := 0; fi16 < 6; fi16++ {
        var v17 bool = false
        _ = fi16
        _ = v17
    }
    fmt.Println(v1, v4, v2)
    if (true && (!true)) {
        _ = len(v3)
        if true {
            var v18 float32 = 1.5
            _ = v18
        }
    }
    switch 32767 {
    case 1:
        var v19 uint8 = 42
        _ = v19
    case -410:
        if (37 == genFn0(false)) {
            switch "a" {
            case "hello":
                if (((!(32767 == 790)) || false) && true) {
                    v1 = (v1 ^ uint8(v4))
                } else {
                    v2 = v2
                    if false {
                        v20 := v3[uint32(v1)]
                        v2 -= 1e10
                        v21 := GenS1{F0: true, F1: (644 > 255), F2: -1}
                        _ = v20
                        _ = v21
                    }
                    if (!true) {
                        v4 += v4
                        fmt.Println(v2)
                    } else {
                        v22 := GenS1{F0: (0 == int64(v4)), F1: ((42 > 0) || (false || (-2.25 < 3.14159))), F2: int32(v1)}
                        _ = v22
                    }
                }
            case "":
                v23 := []int{9223372036854775807, -133, -9223372036854775808}
                _ = v23
            }
            fmt.Println(v1, v2)
        } else {
            for fi24 := 0; fi24 < 14; fi24++ {
                for fi25 := 0; fi25 < 17; fi25++ {
                    var v26 int = (-(fi25))
                    v27 := v3[42]
                    _ = fi25
                    _ = v26
                    _ = v27
                }
                _ = fi24
            }
            var v28 uint64 = v4
            _ = v28
        }
    default:
        switch 42 {
        case 1:
            var v29 uint = 741
            _ = v29
        default:
            var v30 float32 = ((1e10 * ((1.5 - (1e10 * -2.25)) / 1.5)) * 0.0)
            _ = v30
        }
    }
    for fi31 := 0; fi31 < 17; fi31++ {
        for fi32 := 0; fi32 < 17; fi32++ {
            v2 -= 1e10
            _ = fi32
        }
        for fi33 := 0; fi33 < 1; fi33++ {
            fmt.Println(v4, fi31)
            var v34 int8 = 7
            fmt.Println(v1, v4, v1)
            _ = fi33
            _ = v34
        }
        v35 := v3[329]
        _ = fi31
        _ = v35
    }
    v36 := GenS4{F0: (true || (!true)), F1: 255, F2: -541}
    v1 -= v1
    if true {
        v2 += v2
        v37 := GenS2{F0: 1, F1: -9223372036854775808}
        var v38 uint64 = v4
        _ = v37
        _ = v38
    }
    var v39 bool = (!true)
    for fi40 := 0; fi40 < 14; fi40++ {
        fmt.Println(v4, fi40)
        _ = fi40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v39
    _ = v4
}
