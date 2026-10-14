package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 uint64; F2 uint16 }
type GenS4 struct { F0 int64; F1 uint16 }
type GenS5 struct { F0 int; F1 uint16; F2 int64 }
func genFn0(p0 string, p1 uint32) int64 {
    v7 := []int{-2147483648, 0, 42, -771}
    for fi8 := 0; fi8 < 2; fi8++ {
        fmt.Println(p1, p1)
        _ = fi8
    }
    p1 -= 547
    _ = p0
    _ = p1
    _ = v7
    return -9223372036854775808
}
func genFn1(p0 bool, p1 int32) int64 {
    fmt.Println(p0, p0)
    var v9 bool = (!true)
    switch "foo bar" {
    case "foo bar":
        v10 := GenS1{F0: 100}
        _ = v10
    case "hello":
        var v11 uint16 = 2
        _ = v11
    default:
        v12 := []int{-2147483648, 668}
        _ = v12
    }
    v13 := GenS1{F0: -86}
    _ = p0
    _ = p1
    _ = v13
    _ = v9
    return 392
}

func main() {
    var v1 bool = false
    v2 := make(map[uint32]string)
    v3 := []int{0}
    var v4 uint8 = 2
    var v5 int8 = 0
    var v6 int8 = -1
    for fi14 := 0; fi14 < 15; fi14++ {
        v4 += (v4 ^ v4)
        _ = fi14
    }
    fmt.Println(v5, v1, v6)
    if v1 {
        v4 = 127
        fmt.Println(v1)
        if v1 {
            v15 := GenS1{F0: -128}
            _ = v15
        } else {
            v6 = v6
            var v16 uint = 18446744073709551615
            for fi17 := 0; fi17 < 7; fi17++ {
                if true {
                    v18 := GenS1{F0: v5}
                    _ = v18
                } else {
                    for fi19 := 0; fi19 < 4; fi19++ {
                        v6 = v5
                        v5 -= v6
                        _ = fi19
                    }
                    v16 -= (^(v16))
                }
                _ = fi17
            }
            _ = v16
        }
    } else {
        if v1 {
            v20 := GenS1{F0: (^(v5))}
            var v21 uint16 = 332
            if v1 {
                v22 := GenS2{F0: (v4 * v4)}
                _ = v22
            } else {
                var v23 uint64 = 2
                v24 := GenS3{F0: 42, F1: 2, F2: 1}
                switch 740 {
                case 127:
                    if v1 {
                        var v25 float64 = (1.5 / 0.1)
                        for fi26 := 0; fi26 < 11; fi26++ {
                            v23 += 2
                            _ = fi26
                        }
                        switch "a" {
                        case "":
                            for fi27 := 0; fi27 < 18; fi27++ {
                                fmt.Println(v23, v25)
                                if v1 {
                                    switch "" {
                                    case "foo bar":
                                        fmt.Println(v4, fi27, fi27)
                                    case "":
                                        v4 = v4
                                    case "a":
                                        for fi28 := 0; fi28 < 18; fi28++ {
                                            delete(v2, 341)
                                            v2[0] = ""
                                            var v29 int64 = 30
                                            _ = fi28
                                            _ = v29
                                        }
                                    }
                                }
                                _ = fi27
                            }
                        case "\n\t":
                            v5 -= v5
                        case "a":
                            for fi30 := 0; fi30 < 11; fi30++ {
                                v6 = v6
                                fmt.Println(fi30, v21)
                                _ = fi30
                            }
                        }
                        _ = v25
                    } else {
                        if (!(835 != uint(v23))) {
                            if v1 {
                                v21 -= (v21 + v21)
                                v3 = append(v3, -878)
                                v1 = (!((v21 % ((v21) | 1)) >= v21))
                            }
                            v1 = (float32(3.14159) >= float32(1.5))
                            if (76 < 165) {
                                var v31 bool = v1
                                v2[65535] = ("\n\t" + "foo bar")
                                _ = v31
                            } else {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                if (((1e10 * float32(float64((0.0 * ((3.14159 - 3.14159) - 3.14159))))) < (float32(0.1) + 3.14159)) || (-1 <= -1)) {
                                    delete(v2, 1)
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                                _ = len(v3)
                            }
                        }
                        switch "hello" {
                        case "hello":
                            var v32 uint16 = 42
                            _ = v32
                        }
                    }
                default:
                    v1 = (-114 < -1)
                }
                _ = v23
                _ = v24
            }
            _ = v20
            _ = v21
        }
        _ = len(v2)
    }
    switch 9223372036854775807 {
    case 9223372036854775807:
        var v33 uint = 42
        _ = v33
    case 1:
        v6 = 100
    case -469:
        switch -9223372036854775808 {
        case 88:
            if (0.1 != -2.25) {
                var v34 uint32 = 893
                v35 := GenS4{F0: -9223372036854775808, F1: 311}
                _ = v34
                _ = v35
            } else {
                var v36 uint = uint(v4)
                _ = v36
            }
        case -2147483648:
            v2[4294967295] = ""
        case 127:
            if v1 {
                var v37 bool = ((v4 % ((v4) | 1)) == (v4 & (v4 << ((255) & 63))))
                _ = v37
            }
        }
    }
    v6 = (v6 / ((v5) | 1))
    if (4294967295 > 197) {
        var v38 int64 = -1
        var v39 uint64 = 2
        v40 := GenS1{F0: v6}
        _ = v38
        _ = v39
        _ = v40
    } else {
        switch 583 {
        case -2147483648:
            v5 -= v6
        case 399:
            v41 := GenS5{F0: 32767, F1: 32767, F2: 42}
            _ = v41
        }
    }
    fmt.Println(v6, v6, v1)
    for fi42 := 0; fi42 < 6; fi42++ {
        v1 = v1
        var v43 uint8 = 0
        fmt.Println(fi42, v5)
        _ = fi42
        _ = v43
    }
    var v44 int8 = (v6 ^ v5)
    var v45 int = 128
    v44 = (v44 / ((-10) | 1))
    if (v1 || false) {
        var v46 uint8 = (v4 | v4)
        _ = v46
    }
    var v47 uint8 = v4
    for fi48 := 0; fi48 < 16; fi48++ {
        v44 = (v6 - (v44 << ((65535) & 63)))
        _ = fi48
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v47
    _ = v5
    _ = v6
}
