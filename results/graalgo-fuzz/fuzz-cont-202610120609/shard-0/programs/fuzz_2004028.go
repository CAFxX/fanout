package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint }
type GenS2 struct { F0 uint16; F1 uint32; F2 int8 }
type GenS3 struct { F0 int64; F1 uint8; F2 int64 }
type GenS4 struct { F0 uint64; F1 uint8; F2 int8 }
type GenS5 struct { F0 uint16; F1 uint16; F2 int8 }
type GenS6 struct { F0 int32; F1 string }
func genFn0(p0 int64, p1 uint) int16 {
    if (!((p1 * p1) <= 1)) {
        var v4 float64 = 1.5
        _ = v4
    } else {
        v5 := []int{}
        _ = v5
    }
    fmt.Println(p0)
    var v6 uint8 = 42
    _ = p0
    _ = p1
    _ = v6
    return 690
}

func main() {
    var v1 int = 127
    var v2 bool = (true && ("hello" >= ("a" + "")))
    v3 := []uint32{1}
    fmt.Println(v2, v1)
    v2 = v2
    var v7 uint8 = 255
    var v8 int8 = -128
    if v2 {
        v8 -= (-(v8))
    } else {
        if v2 {
            if len(v3) > 0 {
                v3[0] = 436
            }
            var v9 int64 = -1
            for fi10 := 0; fi10 < 19; fi10++ {
                switch ("\n\t" + "foo bar") {
                case "foo bar":
                    switch "\n\t" {
                    case "a":
                        if false {
                            v1 = 4294967295
                            var v11 uint8 = 1
                            v11 += (v11 << ((255) & 63))
                            _ = v11
                        } else {
                            var v12 int16 = -32768
                            var v13 uint32 = 235
                            _ = v12
                            _ = v13
                        }
                    default:
                        if v2 {
                            for fi14 := 0; fi14 < 17; fi14++ {
                                if len(v3) > 0 {
                                    v3[0] = 4294967295
                                }
                                v3 = append(v3, 2147483647)
                                _ = fi14
                            }
                        } else {
                            switch "\n\t" {
                            case "hello":
                                for fi15 := 0; fi15 < 9; fi15++ {
                                    if true {
                                        v8 -= -1
                                    } else {
                                        if false {
                                            if true {
                                                var v16 int = fi10
                                                if v2 {
                                                    var v17 int8 = v8
                                                    if (127 == 127) {
                                                        var v18 string = ("" + ("\n\t" + "\n\t"))
                                                        _ = v18
                                                    }
                                                    v19 := GenS1{F0: -32768, F1: 42, F2: 255}
                                                    fmt.Println(v7, v2)
                                                    v20 := GenS2{F0: 0, F1: 2147483648, F2: v8}
                                                    _ = v17
                                                    _ = v19
                                                    _ = v20
                                                }
                                                _ = v16
                                            }
                                            var v21 float64 = float64(-2.25)
                                            _ = v21
                                        } else {
                                            switch fi15 {
                                            case -2147483648:
                                                v22 := GenS1{F0: 1, F1: -1, F2: 155}
                                                if false {
                                                    var v23 int = (v1 | (fi10 | (v1 % ((fi15) | 1))))
                                                    if (((0.0 + (0.0 / 1.5)) + 3.14159) > 1e10) {
                                                        var v24 uint16 = 65535
                                                        var v25 int8 = -128
                                                        v26 := GenS1{F0: 42, F1: 1, F2: 570}
                                                        _ = v24
                                                        _ = v25
                                                        _ = v26
                                                    } else {
                                                        var v27 uint64 = 1
                                                        v7 -= (^(v7))
                                                        _ = v27
                                                    }
                                                    var v28 int8 = (v8 | (v8 + (v8 - (v8 + v8))))
                                                    _ = v23
                                                    _ = v28
                                                } else {
                                                    v7 -= (v7 - 128)
                                                }
                                                _ = v22
                                            case 0:
                                                v29 := GenS3{F0: 1, F1: (v7 << ((2) & 63)), F2: v9}
                                                _ = v29
                                            case 64:
                                                v2 = v2
                                            default:
                                                fmt.Println(v8)
                                            }
                                        }
                                        v30 := GenS4{F0: 848, F1: (v7 / (((^(v7))) | 1)), F2: 1}
                                        _ = v30
                                    }
                                    v8 = (v8 << ((65535) & 63))
                                    if v2 {
                                        if v2 {
                                            fmt.Println(v7, fi10)
                                            var v31 int = fi15
                                            v1 -= fi10
                                            _ = v31
                                        }
                                        v32 := GenS4{F0: 216, F1: (v7 << ((288) & 63)), F2: (v8 / (((-(v8))) | 1))}
                                        _ = v32
                                    }
                                    _ = fi15
                                }
                            default:
                                v3 = append(v3, 2147483647)
                            }
                            v1 = fi10
                            v33 := GenS5{F0: 65535, F1: 353, F2: v8}
                            _ = v33
                        }
                    }
                case "hello":
                    v34 := GenS6{F0: 0, F1: ("" + "\n\t")}
                    _ = v34
                }
                if true {
                    var v35 int8 = (v8 % ((v8) | 1))
                    var v36 uint64 = 42
                    v8 -= v35
                    _ = v35
                    _ = v36
                }
                v37 := GenS1{F0: 42, F1: -2147483648, F2: 0}
                _ = fi10
                _ = v37
            }
            _ = v9
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    switch "a" {
    case "a":
        v7 = 200
    case "\n\t":
        fmt.Println(v7, v7, v1)
    case "":
        var v38 string = ""
        _ = v38
    default:
        var v39 int = v1
        _ = v39
    }
    if len(v3) > 0 {
        v3[0] = 42
    }
    if (1 <= -1) {
        if false {
            for fi40 := 0; fi40 < 9; fi40++ {
                switch fi40 {
                case 126:
                    if (174 >= 645) {
                        v7 = 0
                    }
                case -9223372036854775808:
                    fmt.Println(v8, v1, fi40)
                default:
                    v3 = append(v3, 42)
                }
                _ = fi40
            }
            v3 = append(v3, 663)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
