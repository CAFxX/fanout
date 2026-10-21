package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 int8; F1 int64; F2 uint16 }
type GenS4 struct { F0 int16; F1 int8; F2 int }
type GenS5 struct { F0 uint64; F1 uint }
type GenS6 struct { F0 uint8; F1 int8 }
type GenS7 struct { F0 bool }
func genFn0(p0 uint8, p1 int32, p2 uint64) uint8 {
    var v5 string = "foo bar"
    v6 := []int{-2147483648, -414, 9223372036854775807, 32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return p0
}

func main() {
    v1 := make(map[string]int)
    var v2 bool = true
    var v3 uint64 = 4294967295
    var v4 int = -1
    var v7 int8 = 42
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 uint8 = 115
        _ = fi8
        _ = v9
    }
    if v2 {
        for fi10 := 0; fi10 < 10; fi10++ {
            delete(v1, "\n\t")
            v4 = (fi10 / ((-9223372036854775808) | 1))
            fmt.Println(v2, v4)
            _ = fi10
        }
    } else {
        var v11 float64 = 3.14159
        for fi12 := 0; fi12 < 12; fi12++ {
            v13 := GenS1{F0: 128}
            _ = fi12
            _ = v13
        }
        _ = v11
    }
    v14 := GenS1{F0: 127}
    v1[("\n\t" + ("foo bar" + (("" + (("\n\t" + ("" + ("foo bar" + ("hello" + "\n\t")))) + (("\n\t" + "hello") + "foo bar"))) + "\n\t")))] = v4
    for fi15 := 0; fi15 < 19; fi15++ {
        var v16 bool = false
        _ = fi15
        _ = v16
    }
    v17 := GenS2{F0: 4294967295, F1: ("" + "\n\t")}
    for fi18 := 0; fi18 < 16; fi18++ {
        v19 := GenS1{F0: 255}
        if v2 {
            switch "a" {
            case "foo bar":
                var v20 float32 = float32((0.0 * 1e10))
                _ = v20
            case "hello":
                if (!v2) {
                    if false {
                        var v21 uint8 = 42
                        var v22 bool = (!(uint32(v3) > 49))
                        v22 = v22
                        _ = v21
                        _ = v22
                    } else {
                        var v23 float64 = 3.14159
                        fmt.Println(v7, v2)
                        for fi24 := 0; fi24 < 14; fi24++ {
                            if (3.14159 >= (0.1 - 1e10)) {
                                if (1 < 424) {
                                    if v2 {
                                        var v25 int64 = -1
                                        v26 := GenS3{F0: (v7 | v7), F1: -9223372036854775808, F2: 0}
                                        var v27 uint32 = 531
                                        _ = v25
                                        _ = v26
                                        _ = v27
                                    } else {
                                        v2 = (42 == 9223372036854775807)
                                        v28 := GenS4{F0: -124, F1: -42, F2: (fi18 * (fi24 >> ((255) & 63)))}
                                        _ = v28
                                    }
                                    v23 = (((((float64(-2.25) * v23) + (((v23 * 1.5) + (v23 + v23)) / float64((0.1 * (((0.1 * ((0.1 - 3.14159) + (1e10 - (0.0 / (0.1 * 1e10))))) + ((1.5 + 0.1) / -2.25)) + 1e10))))) - -2.25) + v23) + 0.1)
                                    var v29 uint = 914
                                    _ = v29
                                } else {
                                    v7 -= -47
                                    var v30 uint32 = 42
                                    _ = v30
                                }
                            }
                            switch (v4 % ((32768) | 1)) {
                            case -2147483648:
                                v31 := GenS5{F0: 67, F1: 950}
                                _ = v31
                            case -9223372036854775808:
                                fmt.Println(fi24, v7)
                            case 0:
                                v32 := []int{(v4 % ((fi18) | 1))}
                                _ = v32
                            }
                            _ = fi24
                        }
                        _ = v23
                    }
                    v4 -= fi18
                }
            case "\n\t":
                switch -9223372036854775808 {
                case 9223372036854775807:
                    var v33 bool = ((v7 % (((v7 * (v7 * (v7 % ((0) | 1))))) | 1)) < v7)
                    _ = v33
                default:
                    if (genFn0(0, 42, v3) != 2) {
                        v7 -= (v7 & (v7 << ((255) & 63)))
                        v7 = (v7 >> ((30) & 63))
                        v7 = -128
                    } else {
                        for fi34 := 0; fi34 < 16; fi34++ {
                            v35 := []int{}
                            if (-9223372036854775808 > -9223372036854775808) {
                                var v36 float64 = 3.14159
                                _ = v36
                            }
                            v37 := GenS6{F0: genFn0(243, -108, v3), F1: (v7 & v7)}
                            _ = fi34
                            _ = v35
                            _ = v37
                        }
                    }
                }
            default:
                v38 := []int{(fi18 ^ 4), v4}
                _ = v38
            }
        } else {
            switch ("a" + "a") {
            case "hello":
                if (true || true) {
                    _ = len(v1)
                    v3 = (v3 << ((65535) & 63))
                } else {
                    fmt.Println(v3, v3, v2)
                }
            case "a":
                fmt.Println(fi18, v2)
            }
            v39 := GenS7{F0: (!(!(2 < 127)))}
            _ = v39
        }
        _ = fi18
        _ = v19
    }
    for fi40 := 0; fi40 < 14; fi40++ {
        v2 = false
        fmt.Println(v4)
        if (v2 || v2) {
            v41 := GenS5{F0: v3, F1: 845}
            _ = v41
        }
        _ = fi40
    }
    var v42 uint32 = 42
    v7 += (v7 - v7)
    for fi43 := 0; fi43 < 4; fi43++ {
        for fi44 := 0; fi44 < 15; fi44++ {
            var v45 uint16 = 21
            v46 := v1[("\n\t" + "hello")]
            var v47 int64 = 807
            _ = fi44
            _ = v45
            _ = v46
            _ = v47
        }
        _ = fi43
    }
    v1["foo bar"] = (v4 << ((42) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v7
}
