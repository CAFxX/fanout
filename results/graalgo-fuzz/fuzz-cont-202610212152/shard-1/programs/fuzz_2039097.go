package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 uint8; F1 string }
func genFn0() int64 {
    var v7 float32 = (float32(3.14159) - 1.5)
    v7 = (((0.1 / v7) + v7) * v7)
    v8 := GenS1{F0: 345, F1: false}
    _ = v7
    _ = v8
    return -744
}
func genFn1(p0 int32) string {
    p0 = p0
    p0 = p0
    var v9 uint16 = 32767
    fmt.Println(v9)
    _ = p0
    _ = v9
    return ("" + "\n\t")
}

func main() {
    var v1 string = ""
    var v2 int32 = 2147483647
    v3 := []uint32{65535, 42}
    var v4 uint = 399
    v5 := make(map[uint32]uint32)
    v6 := make(map[int]string)
    if true {
        switch v1 {
        case "a":
            if false {
                if (((1.5 - 3.14159) >= (1.5 + (((0.0 - (1.5 - 3.14159)) - -2.25) / (float32(1e10) - (0.0 / (1e10 * float32(1.5))))))) && (42 != 127)) {
                    if false {
                        var v10 uint64 = 492
                        _ = v10
                    } else {
                        if (v1 != (("foo bar" + ("hello" + ((v1 + "foo bar") + ((v1 + v1) + (("foo bar" + v1) + ("a" + v1)))))) + v1)) {
                            if true {
                                if false {
                                    if (false || false) {
                                        var v11 uint64 = 42
                                        _ = v11
                                    }
                                }
                            } else {
                                v5[2147483647] = 158
                                switch -1 {
                                case -580:
                                    var v12 int = -1
                                    _ = v12
                                }
                                var v13 int64 = -1
                                _ = v13
                            }
                        } else {
                            if ((369 < 1) && true) {
                                for fi14 := 0; fi14 < 17; fi14++ {
                                    switch v1 {
                                    case "\n\t":
                                        if true {
                                            v1 = v1
                                        }
                                    default:
                                        v15 := GenS2{F0: 18446744073709551615}
                                        _ = v15
                                    }
                                    for fi16 := 0; fi16 < 4; fi16++ {
                                        v4 -= v4
                                        if false {
                                            var v17 int8 = 1
                                            _ = v17
                                        }
                                        fmt.Println(v1, v2, v4)
                                        var v18 uint8 = 42
                                        _ = fi16
                                        _ = v18
                                    }
                                    _ = fi14
                                }
                            } else {
                                if false {
                                    if true {
                                        if (255 >= 81) {
                                            v2 -= (v2 & (v2 + (-(v2))))
                                            v19 := GenS3{F0: (v4 + v4), F1: -727}
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            _ = v19
                                        }
                                    } else {
                                        var v20 float64 = 0.0
                                        _ = v20
                                    }
                                }
                                switch -9223372036854775808 {
                                case 128:
                                    switch -571 {
                                    case 32768:
                                        v21 := GenS3{F0: (v4 >> (((v4 * 0)) & 63)), F1: 558}
                                        _ = v21
                                    }
                                case 4294967295:
                                    v1 = "a"
                                default:
                                    for fi22 := 0; fi22 < 10; fi22++ {
                                        v4 = (v4 % (((^(v4))) | 1))
                                        _ = fi22
                                    }
                                }
                                fmt.Println(v1, v2)
                            }
                            v1 = ""
                        }
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v23 := GenS4{F0: 127, F1: "foo bar"}
                    _ = v23
                }
            } else {
                v4 = (v4 | 2)
            }
        case "foo bar":
            v1 = v1
        }
        for fi24 := 0; fi24 < 13; fi24++ {
            v1 = v1
            v1 = v1
            _ = fi24
        }
    } else {
        v1 = v1
        v4 = 0
        if (42 != 988) {
            var v25 int = -1
            v26 := GenS2{F0: v4}
            delete(v6, v25)
            _ = v25
            _ = v26
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if true {
        var v27 uint64 = 4294967295
        _ = v27
    }
    fmt.Println(v4, v1, v4)
    v28 := GenS1{F0: uint16(v4), F1: true}
    v2 = v2
    v4 = v4
    var v29 int8 = 1
    var v30 string = v1
    if false {
        var v31 float64 = -2.25
        _ = v31
    }
    v2 += (v2 << ((v4) & 63))
    v4 += (v4 / ((v4) | 1))
    v30 = v30
    if false {
        fmt.Println(v29, v4)
    }
    v2 += 42
    var v32 uint64 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
