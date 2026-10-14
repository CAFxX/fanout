package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 bool }
func genFn0() uint64 {
    var v6 int = 32767
    var v7 float64 = 0.0
    var v8 float64 = (v7 + 0.1)
    var v9 int8 = 100
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 979
}
func genFn1(p0 uint16, p1 uint64) uint32 {
    var v10 uint = 65535
    switch ("foo bar" + "") {
    case "hello":
        var v11 uint16 = (p0 ^ p0)
        _ = v11
    case "":
        v12 := []int{60}
        _ = v12
    default:
        switch 4294967295 {
        case 4294967295:
            if false {
                fmt.Println(p1)
                v13 := []int{127}
                v14 := GenS1{F0: -22}
                _ = v13
                _ = v14
            }
        case -275:
            v15 := GenS1{F0: 1}
            _ = v15
        default:
            var v16 int16 = 0
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 16; fi17++ {
        fmt.Println(p0)
        var v18 uint8 = 50
        for fi19 := 0; fi19 < 15; fi19++ {
            v20 := GenS2{F0: -468}
            v10 = uint(v18)
            _ = fi19
            _ = v20
        }
        _ = fi17
        _ = v18
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v10
    return 0
}

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]int)
    v3 := make(map[int]bool)
    v4 := make(map[string]string)
    var v5 uint32 = 366
    if (((!(42 != -2147483648)) && true) || (480 < -337)) {
        var v21 uint32 = genFn1(940, 18446744073709551615)
        _ = v21
    } else {
        v22 := GenS2{F0: -1}
        v23 := GenS3{F0: 2147483647, F1: v5}
        var v24 int64 = 388
        _ = v22
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 7; fi25++ {
        for fi26 := 0; fi26 < 2; fi26++ {
            var v27 uint8 = 127
            if false {
                v27 = (v27 - (v27 % (((^(v27))) | 1)))
            }
            if (122 <= 1000) {
                v27 += v27
            } else {
                var v28 uint64 = genFn0()
                _ = v28
            }
            _ = fi26
            _ = v27
        }
        v5 = (v5 - 775)
        _ = fi25
    }
    fmt.Println(v5)
    switch "a" {
    case "a":
        v5 = (v5 / (((^(v5))) | 1))
    case "\n\t":
        _ = len(v1)
    }
    var v29 string = ("" + ("" + ("a" + "foo bar")))
    for fi30 := 0; fi30 < 14; fi30++ {
        if (1e10 != 0.1) {
            v29 = v29
            v5 -= v5
        }
        fmt.Println(fi30, v5)
        v29 = v29
        _ = fi30
    }
    var v31 float64 = (1.5 + (3.14159 / -2.25))
    switch "a" {
    case "a":
        v29 = (v29 + "foo bar")
    default:
        if (((!false) && false) && (true || false)) {
            v1 = append(v1, (^(v5)))
            if true {
                var v32 int8 = 0
                fmt.Println(v5, v5)
                for fi33 := 0; fi33 < 8; fi33++ {
                    fmt.Println(fi33, fi33, v32)
                    if false {
                        fmt.Println(v5)
                        v4[v29] = (v29 + "foo bar")
                        fmt.Println(v32, fi33)
                    } else {
                        var v34 int32 = 42
                        switch "foo bar" {
                        case "":
                            v35 := GenS1{F0: v32}
                            _ = v35
                        case "hello":
                            fmt.Println(v34, v29)
                        default:
                            if false {
                                if ((float32(float64(3.14159)) / 0.1) >= 1e10) {
                                    fmt.Println(v29, v31)
                                } else {
                                    for fi36 := 0; fi36 < 4; fi36++ {
                                        var v37 uint16 = 0
                                        if false {
                                            var v38 int8 = v32
                                            _ = v38
                                        }
                                        v5 += (v5 + (v5 & uint32(v37)))
                                        _ = fi36
                                        _ = v37
                                    }
                                    for fi39 := 0; fi39 < 20; fi39++ {
                                        switch (((((v29 + v29) + v29) + "a") + (((v29 + ("a" + "hello")) + "\n\t") + v29)) + v29) {
                                        case "a":
                                            v40 := GenS1{F0: v32}
                                            _ = v40
                                        }
                                        _ = fi39
                                    }
                                }
                                if len(v1) > 0 {
                                    v1[0] = (^(v5))
                                }
                                for fi41 := 0; fi41 < 19; fi41++ {
                                    fmt.Println(v29, v34)
                                    var v42 float32 = 0.1
                                    _ = fi41
                                    _ = v42
                                }
                            } else {
                                v43 := GenS4{F0: false}
                                v5 = (v5 ^ v5)
                                _ = v43
                            }
                        }
                        _ = v34
                    }
                    switch (fi33 + 9223372036854775807) {
                    case -9223372036854775808:
                        v44 := v2[(v5 ^ v5)]
                        _ = v44
                    case 32768:
                        var v45 uint16 = 65535
                        _ = v45
                    default:
                        var v46 int32 = -2147483648
                        _ = v46
                    }
                    _ = fi33
                }
                _ = v32
            }
            v47 := GenS5{F0: (false && false)}
            _ = v47
        } else {
            var v48 uint64 = 864
            _ = len(v1)
            _ = v48
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
