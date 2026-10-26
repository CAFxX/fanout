package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 uint; F1 uint32; F2 uint }
func genFn0(p0 bool) int {
    var v6 int32 = -149
    v7 := []int{816, 806, -663, 32768}
    var v8 uint16 = 2
    var v9 bool = p0
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 4294967295
}
func genFn1(p0 uint8, p1 uint32, p2 int8) uint32 {
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 int64 = -9223372036854775808
        _ = fi10
        _ = v11
    }
    p1 = 2
    v12 := []int{}
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return p1
}

func main() {
    var v1 int64 = 1
    v2 := make(map[uint32]uint32)
    var v3 uint16 = 666
    v4 := make(map[int]string)
    v5 := []uint32{42, 764, 65535, 1}
    v3 = 158
    var v13 float64 = 0.1
    switch "" {
    case "a":
        for fi14 := 0; fi14 < 6; fi14++ {
            var v15 uint16 = 798
            v3 -= v3
            _ = fi14
            _ = v15
        }
    default:
        if (false || (("" + "\n\t") >= ("hello" + ("a" + "foo bar")))) {
            v13 -= v13
            v1 += v1
        } else {
            v3 -= (v3 - v3)
            if false {
                var v16 uint32 = 245
                _ = v16
            } else {
                switch genFn0(false) {
                case 9223372036854775807:
                    v3 = (v3 - (v3 / ((v3) | 1)))
                case 509:
                    for fi17 := 0; fi17 < 11; fi17++ {
                        v3 = (v3 - v3)
                        v2[1] = 2147483647
                        v18 := GenS1{F0: 42, F1: 0}
                        _ = fi17
                        _ = v18
                    }
                case -102:
                    v2[4294967295] = 42
                default:
                    if true {
                        for fi19 := 0; fi19 < 4; fi19++ {
                            var v20 int8 = -128
                            var v21 bool = true
                            for fi22 := 0; fi22 < 18; fi22++ {
                                if v21 {
                                    v20 = (v20 | -82)
                                }
                                v2[555] = 1
                                fmt.Println(fi22, v1, fi22)
                                _ = fi22
                            }
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                        var v23 int16 = -1
                        _ = v23
                    }
                }
                if (false && false) {
                    fmt.Println(v1, v1, v1)
                }
            }
            v24 := GenS1{F0: 432, F1: -110}
            _ = v24
        }
    }
    if (0 != 65535) {
        v5 = append(v5, 2147483648)
        var v25 string = "foo bar"
        _ = v25
    } else {
        v13 += (v13 * (v13 / (v13 - v13)))
    }
    v3 += 801
    fmt.Println(v1)
    v13 = (3.14159 * 1.5)
    var v26 int = 127
    if false {
        if ("foo bar" <= ("foo bar" + "\n\t")) {
            if false {
                for fi27 := 0; fi27 < 17; fi27++ {
                    v28 := GenS1{F0: 613, F1: -1}
                    if true {
                        if true {
                            delete(v4, fi27)
                        } else {
                            fmt.Println(v13, fi27)
                        }
                        var v29 uint32 = 4294967295
                        v30 := GenS1{F0: 42, F1: 74}
                        _ = v29
                        _ = v30
                    }
                    _ = fi27
                    _ = v28
                }
                var v31 uint8 = 42
                fmt.Println(v31, v1, v26)
                _ = v31
            } else {
                v3 += v3
                v1 -= 42
                v26 = v26
            }
            switch v26 {
            case 0:
                fmt.Println(v26, v1, v13)
            case -708:
                var v32 string = ("hello" + "hello")
                _ = v32
            }
        } else {
            switch "a" {
            case "":
                _ = len(v5)
            default:
                var v33 uint = 51
                _ = v33
            }
        }
    } else {
        for fi34 := 0; fi34 < 12; fi34++ {
            v35 := GenS2{F0: 474, F1: 65535, F2: 18446744073709551615}
            _ = fi34
            _ = v35
        }
    }
    switch "" {
    case "\n\t":
        var v36 uint16 = v3
        _ = v36
    default:
        var v37 uint32 = 2147483648
        _ = v37
    }
    fmt.Println(v13, v3, v13)
    if (!true) {
        v13 = v13
    }
    var v38 int32 = -1
    v5 = append(v5, genFn1(255, 2, 0))
    if true {
        var v39 uint32 = 2147483647
        _ = v39
    }
    var v40 int16 = 32767
    delete(v2, genFn1(1, 0, 96))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v26
    _ = v3
    _ = v38
    _ = v4
    _ = v40
    _ = v5
}
