package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 uint }
func genFn0() int16 {
    var v7 string = "a"
    fmt.Println(v7)
    v7 = "a"
    _ = v7
    return 32767
}
func genFn1(p0 uint8) int {
    p0 += p0
    p0 = (p0 << ((0) & 63))
    _ = p0
    return -1
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{32767, 42, -1}
    v3 := make(map[string]bool)
    var v4 float32 = float32(((0.0 * ((float64(3.14159) / ((3.14159 + -2.25) * -2.25)) + 1e10)) - 1e10))
    var v5 int32 = 42
    v6 := make(map[string]uint32)
    v8 := GenS1{F0: 42, F1: (v5 + v5)}
    fmt.Println(v5, v5)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if false {
        for fi9 := 0; fi9 < 6; fi9++ {
            for fi10 := 0; fi10 < 10; fi10++ {
                var v11 uint = 957
                fmt.Println(v11, v4)
                _ = fi10
                _ = v11
            }
            fmt.Println(v4, v5)
            v5 += (v5 / ((v5) | 1))
            _ = fi9
        }
    } else {
        if true {
            for fi12 := 0; fi12 < 3; fi12++ {
                v4 = ((v4 / v4) / float32(float64(float32((-2.25 * float64(0.1))))))
                _ = fi12
            }
            fmt.Println(v4)
        }
        v13 := GenS1{F0: 590, F1: -727}
        fmt.Println(v5, v5)
        _ = v13
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        var v15 bool = false
        v5 = v5
        var v16 int32 = (v5 >> ((4294967295) & 63))
        _ = fi14
        _ = v15
        _ = v16
    }
    fmt.Println(v5)
    if true {
        v4 = v4
        if false {
            switch "" {
            case "hello":
                v17 := GenS2{F0: 1000}
                _ = v17
            case "a":
                if len(v2) > 0 {
                    v2[0] = -2147483648
                }
            case "":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            default:
                var v18 int8 = 1
                _ = v18
            }
            var v19 string = ("hello" + "a")
            fmt.Println(v4, v4, v19)
            _ = v19
        } else {
            switch -2147483648 {
            case 127:
                v2 = append(v2, -284)
            case -29:
                for fi20 := 0; fi20 < 11; fi20++ {
                    if (-604 > 42) {
                        v4 = (v4 / float32(1.5))
                        v21 := GenS2{F0: -1}
                        _ = v21
                    } else {
                        if (genFn1(255) < 0) {
                            if ((-9223372036854775808 == 369) || (((v5 - (v5 / ((1) | 1))) <= (-(v5))) || (1 != 42))) {
                                var v22 int64 = 384
                                for fi23 := 0; fi23 < 11; fi23++ {
                                    _ = len(v3)
                                    _ = fi23
                                }
                                switch "foo bar" {
                                case "":
                                    if true {
                                        var v24 int16 = 1
                                        _ = v24
                                    } else {
                                        v22 = (v22 | 928)
                                        var v25 int64 = (v22 << ((18446744073709551615) & 63))
                                        v5 += v5
                                        _ = v25
                                    }
                                case "foo bar":
                                    switch fi20 {
                                    case -9223372036854775808:
                                        v4 = (v4 / v4)
                                    case 32767:
                                        v5 = v5
                                    }
                                }
                                _ = v22
                            }
                        } else {
                            if (1 >= -32768) {
                                fmt.Println(fi20)
                                if (!true) {
                                    var v26 int32 = (v5 + (^(v5)))
                                    _ = v26
                                }
                                if (1.5 >= (float64(3.14159) / (3.14159 + 0.1))) {
                                    var v27 int64 = 684
                                    v28 := GenS3{F0: 313, F1: 42}
                                    _ = v27
                                    _ = v28
                                } else {
                                    switch "\n\t" {
                                    case "":
                                        fmt.Println(fi20, v4)
                                    default:
                                        v29 := GenS1{F0: 255, F1: -2147483648}
                                        _ = v29
                                    }
                                }
                            } else {
                                v5 = (v5 | (-(v5)))
                            }
                            var v30 bool = true
                            _ = v30
                        }
                    }
                    _ = fi20
                }
            case -148:
                if len(v2) > 0 {
                    v2[0] = genFn1(109)
                }
            default:
                var v31 uint32 = 2147483647
                _ = v31
            }
        }
    } else {
        var v32 int16 = -389
        _ = v32
    }
    for fi33 := 0; fi33 < 18; fi33++ {
        switch 0 {
        case 9223372036854775807:
            if len(v2) > 0 {
                v2[0] = 4294967295
            }
        default:
            v5 -= -2147483648
        }
        _ = fi33
    }
    var v34 uint64 = 767
    v5 = v5
    switch 127 {
    case 32767:
        var v35 bool = (-830 <= 9223372036854775807)
        _ = v35
    }
    var v36 string = (("\n\t" + "") + "hello")
    var v37 int32 = v5
    _ = len(v1)
    for fi38 := 0; fi38 < 14; fi38++ {
        v39 := GenS4{F0: 2}
        _ = fi38
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
