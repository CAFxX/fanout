package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 uint }
func genFn0(p0 int16, p1 string, p2 int32) int16 {
    if false {
        p2 = -1
        v5 := []int{9223372036854775807, 9223372036854775807, -953, -9223372036854775808}
        _ = v5
    } else {
        p2 -= 42
    }
    v6 := GenS1{F0: 20, F1: 255}
    fmt.Println(p2, p1, p2)
    var v7 int = -2147483648
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 uint32 = 723
    v2 := make(map[uint32]bool)
    var v3 int = 128
    v4 := []uint32{}
    if true {
        v8 := GenS2{F0: 0, F1: 0}
        var v9 int32 = -1
        v10 := GenS2{F0: v9, F1: -32768}
        _ = v10
        _ = v8
        _ = v9
    }
    v3 += v3
    switch v3 {
    case 127:
        v2[v1] = true
    case 1:
        if (0.1 < 1.5) {
            v4 = append(v4, v1)
            v2[v1] = false
        } else {
            v11 := GenS3{F0: 2147483647, F1: v1}
            _ = v11
        }
    default:
        v1 += 2147483648
    }
    v12 := GenS4{F0: 4294967295}
    var v13 int32 = -2147483648
    for fi14 := 0; fi14 < 19; fi14++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi15 := 0; fi15 < 1; fi15++ {
            var v16 int16 = genFn0(1000, "foo bar", v13)
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    switch "" {
    case "\n\t":
        fmt.Println(v13, v1)
    case "a":
        var v17 uint8 = 0
        _ = v17
    case "foo bar":
        fmt.Println(v1, v13)
    }
    var v18 uint = 65535
    var v19 uint32 = uint32(v18)
    switch "foo bar" {
    case "":
        delete(v2, 2)
    case "\n\t":
        var v20 uint16 = 414
        _ = v20
    case "foo bar":
        fmt.Println(v13)
    default:
        if (84 < 209) {
            for fi21 := 0; fi21 < 4; fi21++ {
                switch (v3 & 42) {
                case -86:
                    v1 = 439
                case 9223372036854775807:
                    if true {
                        v19 += (v19 * 65535)
                    } else {
                        var v22 float64 = float64(1.5)
                        if ((v13 - 919) == v13) {
                            var v23 float64 = ((v22 - ((((v22 + (v22 / 1.0)) + v22) + (0.0 / v22)) + v22)) + 0.0)
                            _ = v23
                        } else {
                            for fi24 := 0; fi24 < 9; fi24++ {
                                if true {
                                    var v25 int32 = 0
                                    _ = v25
                                } else {
                                    v26 := GenS2{F0: 2147483647, F1: genFn0(42, "foo bar", -217)}
                                    v22 -= v22
                                    _ = len(v4)
                                    _ = v26
                                }
                                v22 += (v22 + 0.0)
                                _ = fi24
                            }
                            for fi27 := 0; fi27 < 10; fi27++ {
                                var v28 uint8 = 60
                                _ = fi27
                                _ = v28
                            }
                            v18 = (^(v18))
                        }
                        if (257 >= 0) {
                            _ = len(v4)
                            v29 := GenS3{F0: v13, F1: uint32(v18)}
                            _ = v29
                        } else {
                            for fi30 := 0; fi30 < 1; fi30++ {
                                v22 = v22
                                _ = fi30
                            }
                        }
                        _ = v22
                    }
                case 297:
                    for fi31 := 0; fi31 < 8; fi31++ {
                        v18 = v18
                        _ = len(v2)
                        fmt.Println(v19, fi21)
                        _ = fi31
                    }
                default:
                    var v32 uint8 = 89
                    _ = v32
                }
                _ = fi21
            }
            var v33 string = ("" + "a")
            v19 += v1
            _ = v33
        } else {
            _ = len(v2)
            if ((!((false && (!true)) || ((v13 ^ (v13 ^ v13)) > v13))) && true) {
                var v34 int8 = int8(v1)
                switch (v3 * v3) {
                case 419:
                    fmt.Println(v18, v18, v18)
                default:
                    for fi35 := 0; fi35 < 16; fi35++ {
                        v13 += (v13 + int32(v34))
                        _ = fi35
                    }
                }
                _ = v34
            } else {
                var v36 uint64 = 0
                if false {
                    var v37 string = "a"
                    fmt.Println(v36)
                    var v38 uint = uint(v19)
                    _ = v37
                    _ = v38
                } else {
                    var v39 uint = v18
                    _ = v39
                }
                _ = v36
            }
            if false {
                v19 += v1
                var v40 uint = v18
                _ = v40
            } else {
                for fi41 := 0; fi41 < 7; fi41++ {
                    var v42 float64 = (1.5 - -2.25)
                    _ = fi41
                    _ = v42
                }
            }
        }
    }
    v19 = v19
    fmt.Println(v19, v13, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
