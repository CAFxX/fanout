package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int64 }
type GenS2 struct { F0 int; F1 uint; F2 int64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int; F1 bool; F2 int8 }
type GenS6 struct { F0 uint16 }
func genFn0(p0 uint8, p1 int32, p2 bool) uint64 {
    v5 := GenS1{F0: p2, F1: 865, F2: -503}
    var v6 uint64 = 0
    p1 -= p1
    v7 := GenS1{F0: p2, F1: 255, F2: -556}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return (v6 & v6)
}

func main() {
    v1 := []string{"\n\t", "hello"}
    v2 := make(map[uint32]int)
    var v3 uint32 = 4294967295
    var v4 uint64 = 940
    if (!(("foo bar" + ((((("\n\t" + "hello") + ("hello" + ("" + "\n\t"))) + "") + "") + "hello")) > ("\n\t" + "hello"))) {
        switch 9223372036854775807 {
        case 949:
            fmt.Println(v3)
        case 0:
            if (true || (!false)) {
                _ = len(v1)
                var v8 int = 4294967295
                _ = v8
            }
        default:
            v3 = v3
        }
        for fi9 := 0; fi9 < 4; fi9++ {
            var v10 uint64 = (v4 ^ (v4 ^ (^(v4))))
            v10 = v4
            _ = fi9
            _ = v10
        }
        switch -2147483648 {
        case 758:
            if true {
                fmt.Println(v4)
                for fi11 := 0; fi11 < 14; fi11++ {
                    var v12 float32 = (0.1 - 1e10)
                    _ = fi11
                    _ = v12
                }
            } else {
                var v13 uint8 = 127
                _ = v13
            }
        case 127:
            fmt.Println(v4)
        case -1:
            var v14 uint32 = 632
            _ = v14
        default:
            if ((!false) || (!((v3 % (((v3 / ((v3) | 1))) | 1)) <= (^(v3))))) {
                v4 += (v4 << ((0) & 63))
                var v15 float32 = 0.1
                if false {
                    var v16 int8 = 16
                    _ = v16
                } else {
                    v17 := GenS1{F0: ((true && (((3.14159 / 1.0) <= -2.25) || (true && (128 < 255)))) && (((v4 > (^(v4))) || true) || (255 != 2))), F1: 311, F2: 239}
                    if (true || (-1 > -128)) {
                        switch -9223372036854775808 {
                        case 453:
                            v3 += 42
                        }
                        v18 := GenS1{F0: (!(1 > genFn0(4, 42, false))), F1: 1, F2: 1}
                        _ = v18
                    }
                    _ = v17
                }
                _ = v15
            }
        }
    } else {
        delete(v2, (v3 / ((v3) | 1)))
        if true {
            var v19 uint16 = 2
            _ = v19
        }
        for fi20 := 0; fi20 < 15; fi20++ {
            var v21 float32 = ((-2.25 + ((float32((1.5 + (((0.0 - 1e10) / -2.25) * float64((1e10 - 1.5))))) + 1.5) * float32(3.14159))) - 1e10)
            v22 := GenS1{F0: true, F1: 218, F2: 0}
            _ = fi20
            _ = v21
            _ = v22
        }
    }
    if (51 < 0) {
        if true {
            _ = len(v2)
            v23 := GenS2{F0: 128, F1: 109, F2: -593}
            _ = v23
        }
    }
    fmt.Println(v4, v3)
    if (v3 != v3) {
        for fi24 := 0; fi24 < 12; fi24++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v1 = append(v1, "hello")
            v25 := GenS3{F0: 4294967295}
            _ = fi24
            _ = v25
        }
        v26 := GenS4{F0: 1}
        v27 := GenS5{F0: -118, F1: (!false), F2: 1}
        _ = v26
        _ = v27
    }
    v4 = (v4 << ((18446744073709551615) & 63))
    v4 += (v4 / ((v4) | 1))
    var v28 float64 = float64(((-2.25 - 0.1) / 3.14159))
    fmt.Println(v28)
    switch 0 {
    case -2147483648:
        switch "\n\t" {
        case "\n\t":
            if true {
                v3 = (v3 * 0)
                var v29 uint32 = (v3 % ((v3) | 1))
                v4 = (v4 % (((^(v4))) | 1))
                _ = v29
            }
        case "":
            delete(v2, v3)
        case "a":
            var v30 uint64 = 155
            _ = v30
        default:
            delete(v2, (v3 / ((2) | 1)))
        }
    case -1:
        v3 -= (v3 % (((^(v3))) | 1))
    case 32767:
        for fi31 := 0; fi31 < 11; fi31++ {
            for fi32 := 0; fi32 < 10; fi32++ {
                var v33 int32 = 1
                v34 := GenS6{F0: 655}
                var v35 int = int(v3)
                _ = fi32
                _ = v33
                _ = v34
                _ = v35
            }
            _ = fi31
        }
    default:
        switch ("hello" + ((("\n\t" + "foo bar") + ("hello" + "")) + "\n\t")) {
        case "hello":
            var v36 int16 = -479
            _ = v36
        }
    }
    v3 = uint32(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
}
