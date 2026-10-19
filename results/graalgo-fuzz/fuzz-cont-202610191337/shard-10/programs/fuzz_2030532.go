package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8; F1 uint16; F2 uint }
type GenS4 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS5 struct { F0 uint }
func genFn0(p0 uint16, p1 int16) int8 {
    if ((1e10 - float32(float64((0.1 * (((float32(1.5) / (0.1 + -2.25)) + (0.0 / 1.5)) * -2.25))))) < (float32(1e10) / (-2.25 + (1e10 + 0.1)))) {
        if false {
            var v6 int64 = -1
            _ = v6
        } else {
            p1 = (p1 >> ((255) & 63))
            var v7 uint8 = 2
            v8 := GenS1{F0: (("foo bar" + ("hello" + (("hello" + "hello") + (("hello" + "hello") + "")))) + "\n\t")}
            _ = v7
            _ = v8
        }
        switch "" {
        case "":
            v9 := GenS1{F0: (("" + "a") + "hello")}
            _ = v9
        case "hello":
            var v10 bool = (false || (!false))
            _ = v10
        default:
            var v11 uint64 = 42
            _ = v11
        }
        var v12 int32 = -457
        _ = v12
    } else {
        v13 := GenS2{F0: (p1 / ((1) | 1))}
        for fi14 := 0; fi14 < 13; fi14++ {
            switch fi14 {
            case 32767:
                fmt.Println(fi14, fi14)
            default:
                v15 := GenS2{F0: (p1 | 924)}
                _ = v15
            }
            _ = fi14
        }
        if false {
            v16 := GenS3{F0: 128, F1: (p0 << ((1) & 63)), F2: 715}
            fmt.Println(p0)
            var v17 float64 = 1.5
            _ = v16
            _ = v17
        }
        _ = v13
    }
    var v18 uint64 = 0
    switch 32768 {
    case 649:
        var v19 uint16 = (p0 ^ (p0 >> ((333) & 63)))
        _ = v19
    case -2147483648:
        v20 := []int{-116, -1}
        _ = v20
    }
    v21 := GenS3{F0: 2, F1: p0, F2: 136}
    v22 := GenS2{F0: (p1 << ((718) & 63))}
    _ = p0
    _ = p1
    _ = v18
    _ = v21
    _ = v22
    return 93
}
func genFn1(p0 int, p1 uint8, p2 int64) int64 {
    p1 -= p1
    var v23 int = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    return (p2 & p2)
}

func main() {
    var v1 int16 = 42
    var v2 uint8 = 1
    var v3 bool = true
    var v4 uint64 = 2
    v5 := make(map[uint32]uint32)
    for fi24 := 0; fi24 < 19; fi24++ {
        v25 := []int{fi24, (-(fi24))}
        _ = fi24
        _ = v25
    }
    v2 = v2
    v26 := GenS2{F0: (v1 << ((18446744073709551615) & 63))}
    fmt.Println(v1)
    v4 -= (v4 | v4)
    _ = len(v5)
    v27 := []int{2147483647}
    if len(v27) >= 2 {
        v27 = v27[1:len(v27)]
    }
    v1 = (-(v1))
    if false {
        fmt.Println(v3, v2, v3)
        if len(v27) > 0 {
            v27[0] = 0
        }
        v4 = v4
    }
    v1 = v1
    v5[168] = 628
    v4 = (^(v4))
    v1 = v1
    switch ("a" + "") {
    case "\n\t":
        if (-1 >= -906) {
            fmt.Println(v3, v3, v4)
            v2 = v2
            for fi28 := 0; fi28 < 7; fi28++ {
                switch fi28 {
                case 32767:
                    _ = len(v5)
                case -117:
                    if false {
                        v3 = ((((v4 << ((0) & 63)) == 1) && (827 != 1)) && (v2 > 2))
                        v2 = v2
                    } else {
                        if (("a" < "\n\t") || v3) {
                            if (65535 < 1) {
                                delete(v5, 2147483647)
                                v29 := GenS2{F0: 42}
                                delete(v5, 426)
                                _ = v29
                            } else {
                                for fi30 := 0; fi30 < 13; fi30++ {
                                    v31 := v5[1]
                                    v32 := GenS4{F0: 1, F1: v4, F2: 255}
                                    var v33 uint8 = v2
                                    _ = fi30
                                    _ = v31
                                    _ = v32
                                    _ = v33
                                }
                                v34 := GenS2{F0: v1}
                                _ = v34
                            }
                        }
                        if v3 {
                            if len(v27) >= 2 {
                                v27 = v27[1:len(v27)]
                            }
                            switch 202 {
                            case -1:
                                var v35 int16 = 1
                                _ = v35
                            case 295:
                                v3 = false
                            default:
                                var v36 int8 = 0
                                _ = v36
                            }
                            var v37 int64 = 0
                            _ = v37
                        }
                        v38 := GenS1{F0: "foo bar"}
                        _ = v38
                    }
                }
                _ = fi28
            }
        } else {
            v39 := GenS5{F0: 1}
            _ = v39
        }
    case "a":
        fmt.Println(v3, v3, v4)
    case "":
        fmt.Println(v3)
    }
    v40 := GenS2{F0: v1}
    _ = len(v27)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v40
    _ = v5
}
