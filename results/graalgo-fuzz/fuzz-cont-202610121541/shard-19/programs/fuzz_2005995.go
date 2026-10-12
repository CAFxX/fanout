package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 int16; F2 int }
type GenS4 struct { F0 int }
type GenS5 struct { F0 string; F1 int64; F2 int32 }
func genFn0(p0 int) string {
    p0 += (-(p0))
    switch ("foo bar" + ("foo bar" + "a")) {
    case "foo bar":
        var v6 uint64 = 1
        _ = v6
    case "a":
        var v7 float64 = 0.0
        _ = v7
    case "\n\t":
        var v8 float32 = ((0.0 + 1e10) - ((3.14159 - ((((3.14159 * 1.5) * (0.1 + -2.25)) - float32(1.5)) - (0.0 / (3.14159 - (float32((1e10 * -2.25)) / 1.0))))) / 1.0))
        _ = v8
    }
    v9 := GenS1{F0: (("a" + "a") + ""), F1: 162, F2: 307}
    _ = p0
    _ = v9
    return "\n\t"
}
func genFn1(p0 uint8, p1 uint) int {
    p0 -= (p0 / ((9) | 1))
    for fi10 := 0; fi10 < 8; fi10++ {
        for fi11 := 0; fi11 < 9; fi11++ {
            if true {
                switch ("hello" + ("foo bar" + "a")) {
                case "\n\t":
                    switch "\n\t" {
                    case "\n\t":
                        v12 := GenS2{F0: fi10}
                        _ = v12
                    }
                case "a":
                    p1 = p1
                default:
                    if false {
                        var v13 uint8 = (p0 - p0)
                        var v14 bool = false
                        _ = v13
                        _ = v14
                    }
                }
            }
            _ = fi11
        }
        _ = fi10
    }
    v15 := GenS2{F0: 808}
    _ = p0
    _ = p1
    _ = v15
    return 1
}

func main() {
    var v1 int32 = -206
    var v2 string = "foo bar"
    var v3 int16 = 574
    v4 := make(map[uint32]string)
    var v5 int32 = (^(v1))
    var v16 int = 4294967295
    v17 := GenS2{F0: (v16 - (v16 - 255))}
    v18 := []int{(v16 % (((v16 / ((64) | 1))) | 1)), (v16 >> ((98) & 63))}
    fmt.Println(v5)
    v1 = v5
    v4[1] = (((v2 + v2) + "hello") + "hello")
    v19 := GenS3{F0: v2, F1: (v3 / ((v3) | 1)), F2: (v16 - (v16 / ((v16) | 1)))}
    v16 += 32768
    var v20 uint64 = 1
    fmt.Println(v3, v16, v5)
    var v21 float64 = ((1e10 * -2.25) * 0.1)
    for fi22 := 0; fi22 < 4; fi22++ {
        delete(v4, 2)
        if (-1 < -128) {
            if false {
                v23 := GenS4{F0: genFn1(216, 203)}
                _ = v23
            } else {
                v3 = (-(v3))
                for fi24 := 0; fi24 < 13; fi24++ {
                    v4[2147483648] = v2
                    if false {
                        v25 := GenS5{F0: v2, F1: 1, F2: -2147483648}
                        v5 = (v1 >> ((523) & 63))
                        var v26 uint16 = 2
                        _ = v25
                        _ = v26
                    }
                    _ = fi24
                }
            }
            if ("\n\t" <= (v2 + "\n\t")) {
                var v27 int = (fi22 + (v16 % ((53) | 1)))
                _ = v27
            } else {
                v28 := GenS5{F0: ((("a" + "foo bar") + "foo bar") + ((v2 + v2) + (v2 + v2))), F1: -1, F2: int32(v16)}
                delete(v4, 2147483648)
                _ = v28
            }
        }
        _ = fi22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
