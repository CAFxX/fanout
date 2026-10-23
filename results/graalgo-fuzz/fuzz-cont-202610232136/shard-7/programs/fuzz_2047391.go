package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 uint }
func genFn0(p0 uint16, p1 int, p2 int64) int32 {
    v4 := []int{(p1 / ((-1) | 1)), (p1 >> ((42) & 63)), p1}
    if len(v4) > 0 {
        v4[0] = 1
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -2147483648
}
func genFn1(p0 int16, p1 string, p2 int16) string {
    for fi5 := 0; fi5 < 14; fi5++ {
        p2 -= p0
        _ = fi5
    }
    var v6 int = 127
    switch -1 {
    case 809:
        p0 -= p0
    case -913:
        for fi7 := 0; fi7 < 4; fi7++ {
            switch "hello" {
            case "":
                v8 := GenS1{F0: 2}
                _ = v8
            default:
                fmt.Println(v6, v6)
            }
            v9 := []int{-793, 32767, 953}
            _ = fi7
            _ = v9
        }
    case 0:
        p2 -= (^(p2))
    default:
        switch (-(v6)) {
        case -839:
            switch v6 {
            case -2147483648:
                var v10 bool = false
                _ = v10
            case -607:
                v11 := []int{}
                _ = v11
            case -1:
                v12 := []int{}
                _ = v12
            }
        case 339:
            v6 = 128
        default:
            var v13 int16 = p0
            _ = v13
        }
    }
    p0 += p0
    v14 := GenS2{F0: 42, F1: 18446744073709551615}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v6
    return ("hello" + ("foo bar" + p1))
}

func main() {
    var v1 uint = 42
    v2 := make(map[string]string)
    v3 := make(map[int]uint32)
    var v15 uint32 = 2147483648
    v16 := []int{}
    _ = len(v16)
    v1 += (v1 << (((v1 ^ 65535)) & 63))
    v1 -= (v1 | 2)
    var v17 float32 = 1e10
    v15 += v15
    var v18 uint32 = (v15 % ((v15) | 1))
    switch -1 {
    case -1:
        v19 := v3[1]
        _ = v19
    case 175:
        v15 = v15
    case 2147483647:
        v17 = 0.0
    default:
        if true {
            switch -2147483648 {
            case 0:
                var v20 int64 = 1
                _ = v20
            case 32767:
                for fi21 := 0; fi21 < 18; fi21++ {
                    if ((1 >= v18) && (true && true)) {
                        for fi22 := 0; fi22 < 12; fi22++ {
                            v18 -= 669
                            var v23 int8 = 100
                            _ = fi22
                            _ = v23
                        }
                    }
                    fmt.Println(v1, v17)
                    v17 = (v17 + v17)
                    _ = fi21
                }
            default:
                v17 = (1.5 / v17)
            }
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
        } else {
            for fi24 := 0; fi24 < 4; fi24++ {
                fmt.Println(fi24)
                switch "a" {
                case "a":
                    for fi25 := 0; fi25 < 19; fi25++ {
                        var v26 bool = true
                        if (genFn0(376, 32768, -1) >= -2147483648) {
                            v18 = v15
                            var v27 uint = 1
                            _ = v27
                        }
                        if (((0.1 + 0.0) - 0.0) != (float64(v17) / 1.0)) {
                            v16 = append(v16, 128)
                            v28 := GenS2{F0: -2147483648, F1: (v1 | (v1 / ((v1) | 1)))}
                            _ = v28
                        } else {
                            if v26 {
                                if v26 {
                                    fmt.Println(v15, v1)
                                    fmt.Println(fi24, v15, fi24)
                                    var v29 float64 = 3.14159
                                    _ = v29
                                }
                                _ = len(v2)
                            } else {
                                var v30 string = "\n\t"
                                fmt.Println(v18)
                                _ = v30
                            }
                            if v26 {
                                var v31 bool = false
                                if (18446744073709551615 != 896) {
                                    v17 = ((((v17 - v17) / (float32(-2.25) + (float32(float64(v17)) + -2.25))) - v17) - float32(float64(0.0)))
                                    var v32 uint8 = 1
                                    var v33 float32 = 0.1
                                    _ = v32
                                    _ = v33
                                } else {
                                    var v34 int64 = 150
                                    _ = v34
                                }
                                if (42 >= -128) {
                                    fmt.Println(fi25, v1)
                                    var v35 float64 = (1.5 / (1.5 + -2.25))
                                    _ = v35
                                } else {
                                    var v36 int = (fi25 & fi25)
                                    fmt.Println(v15)
                                    var v37 uint32 = v15
                                    var v38 uint64 = uint64(v1)
                                    _ = v36
                                    _ = v37
                                    _ = v38
                                }
                                var v39 int8 = 57
                                _ = v31
                                _ = v39
                            } else {
                                var v40 uint16 = 2
                                _ = v40
                            }
                            v18 = (v15 - v18)
                        }
                        _ = fi25
                        _ = v26
                    }
                }
                var v41 bool = false
                _ = fi24
                _ = v41
            }
            var v42 float64 = 0.1
            _ = v42
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
