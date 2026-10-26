package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64 }
func genFn0(p0 int8, p1 int16, p2 uint8) int16 {
    p1 -= p1
    var v5 int8 = -16
    var v6 int32 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (p1 << ((0) & 63))
}
func genFn1(p0 uint64, p1 uint64) bool {
    switch "\n\t" {
    case "hello":
        p1 += 42
    case "foo bar":
        if (true && (false && true)) {
            fmt.Println(p1, p0)
            p1 += (p0 * (p1 + p1))
            p1 = 4294967295
        } else {
            v7 := []int{}
            _ = v7
        }
    case "":
        v8 := []int{9223372036854775807, 965, -9223372036854775808}
        _ = v8
    default:
        for fi9 := 0; fi9 < 1; fi9++ {
            for fi10 := 0; fi10 < 7; fi10++ {
                p1 -= (^(p0))
                if true {
                    fmt.Println(p1, fi10, fi10)
                    var v11 float64 = ((((float64(float32((0.0 - ((-2.25 / 1.0) / -2.25)))) - 1.5) * ((float64(1.5) - 3.14159) / float64(1e10))) + (float64(0.1) + (1.5 / 3.14159))) - (3.14159 / 1.5))
                    _ = v11
                }
                _ = fi10
            }
            p1 = 920
            _ = fi9
        }
    }
    var v12 uint8 = 255
    switch ("foo bar" + "a") {
    case "foo bar":
        fmt.Println(p1)
    case "hello":
        for fi13 := 0; fi13 < 13; fi13++ {
            if (!false) {
                v14 := []int{(fi13 & 490), -2147483648, fi13}
                _ = v14
            }
            _ = fi13
        }
    }
    _ = p0
    _ = p1
    _ = v12
    return false
}

func main() {
    var v1 string = "a"
    var v2 int64 = 96
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]string)
    if false {
        v15 := GenS1{F0: 2147483647, F1: 32767, F2: (v1 + "")}
        for fi16 := 0; fi16 < 20; fi16++ {
            for fi17 := 0; fi17 < 8; fi17++ {
                v1 = ("hello" + v1)
                _ = fi17
            }
            v18 := GenS1{F0: 162, F1: int16(fi16), F2: (((("a" + v1) + "a") + v1) + v1)}
            _ = fi16
            _ = v18
        }
        _ = v15
    } else {
        var v19 uint = 975
        v20 := []int{127, 9223372036854775807}
        _ = v19
        _ = v20
    }
    v1 = (v1 + v1)
    v21 := GenS1{F0: 0, F1: -831, F2: ""}
    v1 = (v1 + ("" + ((v1 + v1) + ("\n\t" + v1))))
    for fi22 := 0; fi22 < 12; fi22++ {
        var v23 uint64 = 2
        var v24 int64 = 843
        fmt.Println(fi22, v24)
        _ = fi22
        _ = v23
        _ = v24
    }
    var v25 int = int(v2)
    var v26 float32 = -2.25
    v2 += (v2 + (v2 - (-(v2))))
    switch v25 {
    case 804:
        switch v1 {
        case "\n\t":
            if (!true) {
                if true {
                    if true {
                        fmt.Println(v2, v1, v1)
                        v25 = v25
                        switch (v1 + v1) {
                        case "\n\t":
                            v25 = 32768
                        case "":
                            v27 := GenS2{F0: v2}
                            _ = v27
                        case "a":
                            v28 := GenS3{F0: (v25 / ((32767) | 1))}
                            _ = v28
                        default:
                            if (false && (!(genFn0(70, -155, 42) < -668))) {
                                v29 := []int{-2147483648, 32768, v25, (v25 & v25)}
                                if (!true) {
                                    fmt.Println(v1)
                                    var v30 uint = 42
                                    switch (-(v25)) {
                                    case -552:
                                        v3["a"] = 354
                                    case 717:
                                        fmt.Println(v2, v1)
                                    default:
                                        fmt.Println(v30)
                                    }
                                    _ = v30
                                }
                                var v31 int16 = 1000
                                _ = v29
                                _ = v31
                            }
                        }
                    } else {
                        fmt.Println(v1)
                        v32 := GenS4{F0: (v2 % ((1) | 1))}
                        _ = v32
                    }
                    v2 += (v2 & (v2 | -35))
                }
            } else {
                v4[938] = v1
                fmt.Println(v1, v1, v2)
            }
        default:
            v33 := v3[((v1 + ((("" + (v1 + v1)) + ("\n\t" + "foo bar")) + ("\n\t" + v1))) + "\n\t")]
            _ = v33
        }
    case 9223372036854775807:
        v34 := GenS4{F0: 9223372036854775807}
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v21
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
