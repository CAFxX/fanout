package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint8; F1 uint64; F2 string }
type GenS7 struct { F0 int8; F1 uint; F2 uint16 }
func genFn0(p0 uint) uint64 {
    p0 -= p0
    for fi6 := 0; fi6 < 6; fi6++ {
        p0 = 0
        _ = fi6
    }
    p0 = p0
    fmt.Println(p0)
    p0 -= (p0 % (((p0 + (p0 ^ 0))) | 1))
    _ = p0
    return 946
}
func genFn1() string {
    var v7 int8 = -128
    if true {
        v8 := []int{-734, 42, 127, 32767}
        v9 := GenS1{F0: 42, F1: 369, F2: -1}
        _ = v8
        _ = v9
    } else {
        fmt.Println(v7)
        fmt.Println(v7)
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        var v11 uint = 1
        _ = fi10
        _ = v11
    }
    _ = v7
    return ("foo bar" + "hello")
}

func main() {
    v1 := []uint32{65535, 4294967295, 4294967295}
    var v2 uint = 255
    var v3 bool = (true || (853 < 26))
    var v4 string = ("foo bar" + "\n\t")
    var v5 uint16 = 255
    var v12 uint = 18446744073709551615
    v13 := GenS2{F0: 128}
    var v14 uint64 = 1
    v2 -= 65535
    v5 = (v5 << ((v12) & 63))
    v12 = (^(v12))
    for fi15 := 0; fi15 < 1; fi15++ {
        switch "hello" {
        case "hello":
            var v16 int32 = int32(fi15)
            _ = v16
        case "\n\t":
            fmt.Println(v5, v4, v4)
        default:
            fmt.Println(v2, v12, fi15)
        }
        var v17 uint16 = 42
        _ = fi15
        _ = v17
    }
    if v3 {
        if v3 {
            var v18 uint8 = 99
            _ = v18
        } else {
            v2 = v12
            v1 = append(v1, 65535)
        }
        if ((-2.25 * (0.0 * 1.5)) < (float32((0.0 * 1e10)) - 0.0)) {
            v19 := GenS1{F0: 1, F1: (v2 * (^(v12))), F2: -632}
            v1 = append(v1, 239)
            v5 -= v5
            _ = v19
        }
    } else {
        v20 := GenS1{F0: 42, F1: uint(v5), F2: -2147483648}
        _ = v20
    }
    var v21 int16 = 1000
    if v3 {
        var v22 float32 = 0.1
        _ = v22
    } else {
        fmt.Println(v4, v12, v21)
        if false {
            switch v4 {
            case "":
                for fi23 := 0; fi23 < 14; fi23++ {
                    var v24 int32 = 456
                    v1 = append(v1, 65535)
                    switch (fi23 * 0) {
                    case -528:
                        v25 := GenS1{F0: v5, F1: 4294967295, F2: int32(fi23)}
                        _ = v25
                    case 134:
                        v24 -= v24
                    case 32767:
                        for fi26 := 0; fi26 < 19; fi26++ {
                            v27 := GenS2{F0: fi26}
                            _ = fi26
                            _ = v27
                        }
                    default:
                        for fi28 := 0; fi28 < 20; fi28++ {
                            v1 = append(v1, 2147483647)
                            _ = fi28
                        }
                    }
                    _ = fi23
                    _ = v24
                }
            case "foo bar":
                var v29 float32 = (-2.25 + ((3.14159 * 1.5) - 1.5))
                _ = v29
            case "\n\t":
                for fi30 := 0; fi30 < 18; fi30++ {
                    v1 = append(v1, 1)
                    _ = fi30
                }
            }
            if (v3 && (false && (false && ((!true) && (286 >= 42))))) {
                v31 := GenS3{F0: -61, F1: 166}
                for fi32 := 0; fi32 < 11; fi32++ {
                    v33 := GenS4{F0: 311, F1: (^(v21)), F2: (v5 >> ((v2) & 63))}
                    var v34 uint64 = (v14 << ((v12) & 63))
                    _ = fi32
                    _ = v33
                    _ = v34
                }
                _ = v31
            } else {
                if (118 <= 4294967295) {
                    var v35 uint8 = 0
                    _ = len(v1)
                    switch 42 {
                    case 42:
                        fmt.Println(v3, v35)
                    case -1:
                        v21 = 32767
                    case -9223372036854775808:
                        if true {
                            if (!(!((0.1 - float32(0.0)) >= 1e10))) {
                                fmt.Println(v2)
                                v36 := GenS5{F0: uint64(v5)}
                                switch 479 {
                                case 502:
                                    if ((v12 << ((v2) & 63)) != v12) {
                                        v5 = v5
                                        var v37 int32 = -136
                                        switch 754 {
                                        case 9223372036854775807:
                                            var v38 uint32 = 0
                                            _ = v38
                                        }
                                        _ = v37
                                    } else {
                                        v2 = 632
                                        if (-83 == -1) {
                                            var v39 float64 = -2.25
                                            var v40 float32 = 0.1
                                            _ = v39
                                            _ = v40
                                        } else {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            var v41 string = v4
                                            v42 := GenS6{F0: 128, F1: 630, F2: "\n\t"}
                                            _ = v41
                                            _ = v42
                                        }
                                    }
                                case -9223372036854775808:
                                    v21 = v21
                                default:
                                    if v3 {
                                        if (3.14159 != 0.0) {
                                            switch -984 {
                                            case -9223372036854775808:
                                                v1 = append(v1, uint32(v2))
                                            default:
                                                fmt.Println(v2)
                                            }
                                        }
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                }
                                _ = v36
                            }
                            if (uint32(v35) >= 2147483648) {
                                fmt.Println(v12)
                            }
                        } else {
                            var v43 uint16 = (^(v5))
                            _ = v43
                        }
                    default:
                        v12 = v12
                    }
                    _ = v35
                }
                v44 := GenS7{F0: 127, F1: 0, F2: v5}
                _ = v44
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
