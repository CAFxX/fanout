package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint }
type GenS2 struct { F0 uint; F1 int32; F2 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint16; F1 uint32; F2 uint8 }
type GenS7 struct { F0 int; F1 int8; F2 int8 }
type GenS8 struct { F0 uint8; F1 uint32 }
type GenS9 struct { F0 uint; F1 int16 }
func genFn0() int16 {
    var v5 bool = (!false)
    var v6 int32 = -1
    v7 := []int{1, -1}
    for fi8 := 0; fi8 < 14; fi8++ {
        v7 = append(v7, fi8)
        _ = fi8
    }
    v9 := GenS1{F0: 31, F1: 350, F2: 378}
    _ = v5
    _ = v6
    _ = v7
    _ = v9
    return 651
}

func main() {
    v1 := make(map[string]string)
    var v2 int16 = -32768
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]int)
    v10 := GenS2{F0: 676, F1: 1, F2: 1}
    v2 = (^(v2))
    v11 := []int{963, -712, -726}
    v2 -= genFn0()
    for fi12 := 0; fi12 < 16; fi12++ {
        for fi13 := 0; fi13 < 12; fi13++ {
            fmt.Println(fi13)
            if ((-1 <= -17) || ((fi12 >> ((701) & 63)) < 127)) {
                switch (fi12 - fi13) {
                case -2147483648:
                    if ((fi13 / ((-2147483648) | 1)) <= fi13) {
                        v14 := GenS3{F0: 0}
                        switch "" {
                        case "a":
                            v15 := GenS4{F0: false}
                            _ = v15
                        }
                        _ = v14
                    } else {
                        v16 := GenS4{F0: (!((3.14159 - ((3.14159 / 1.5) / (1e10 + (1.5 - 1e10)))) >= 3.14159))}
                        _ = v16
                    }
                default:
                    v17 := GenS4{F0: false}
                    _ = v17
                }
                var v18 uint64 = 1
                var v19 float64 = 1e10
                _ = v18
                _ = v19
            }
            if len(v11) > 0 {
                v11[0] = (fi13 << ((42) & 63))
            }
            v20 := GenS5{F0: true}
            switch 42 {
            case -972:
                v2 += v2
                if (0 > 16) {
                    v21 := GenS1{F0: 184, F1: 2147483647, F2: 255}
                    if false {
                        if false {
                            fmt.Println(fi12)
                            v2 += genFn0()
                            v2 = (v2 % (((v2 * 231)) | 1))
                            v22 := GenS6{F0: 732, F1: 4294967295, F2: 164}
                            v23 := GenS3{F0: 18446744073709551615}
                            if (v2 <= v2) {
                                v2 += v2
                                v2 += v2
                            } else {
                                var v24 bool = true
                                if v24 {
                                    var v25 uint32 = 4294967295
                                    var v26 int16 = (v2 ^ v2)
                                    _ = v25
                                    _ = v26
                                } else {
                                    v27 := GenS2{F0: 2, F1: 2147483647, F2: 0}
                                    var v28 uint64 = 42
                                    var v29 float64 = 1.5
                                    _ = v27
                                    _ = v28
                                    _ = v29
                                }
                                _ = v24
                            }
                            fmt.Println(fi12, fi12, v2)
                            _ = v22
                            _ = v23
                        } else {
                            v2 = v2
                            var v30 int = 9223372036854775807
                            v30 = (v30 & (v30 * 128))
                            _ = v30
                        }
                        v31 := v4[4294967295]
                        var v32 float32 = ((float32(0.0) * (1.5 + (1e10 - (3.14159 / -2.25)))) + 3.14159)
                        switch "hello" {
                        case "\n\t":
                            if (942 < int64(v2)) {
                                v32 -= float32((1e10 / float64(v32)))
                                if ((0.1 - float64((v32 - -2.25))) > 1e10) {
                                    switch v31 {
                                    case -1:
                                        if false {
                                            v33 := GenS7{F0: v31, F1: -89, F2: 1}
                                            _ = v33
                                        }
                                    case 4294967295:
                                        var v34 int16 = (-(v2))
                                        _ = v34
                                    default:
                                        v35 := GenS5{F0: true}
                                        _ = v35
                                    }
                                }
                                if (((273 == 4294967295) && (true || false)) || ("a" != "a")) {
                                    if false {
                                        if len(v11) > 0 {
                                            v11[0] = (v31 >> ((18446744073709551615) & 63))
                                        }
                                        var v36 bool = (false || (955 > 343))
                                        _ = v36
                                    }
                                    v37 := GenS5{F0: false}
                                    fmt.Println(v2, v2)
                                    v38 := GenS8{F0: 2, F1: 2}
                                    _ = v37
                                    _ = v38
                                }
                                switch ("" + "") {
                                case "foo bar":
                                    v39 := GenS1{F0: 153, F1: 23, F2: 298}
                                    _ = v39
                                case "":
                                    if len(v11) >= 2 {
                                        v11 = v11[1:len(v11)]
                                    }
                                case "a":
                                    var v40 int8 = int8(v2)
                                    _ = v40
                                }
                            }
                        case "":
                            v41 := GenS7{F0: (-(v31)), F1: 8, F2: 127}
                            _ = v41
                        default:
                            fmt.Println(fi13, fi13)
                        }
                        _ = v31
                        _ = v32
                    } else {
                        v42 := GenS4{F0: true}
                        fmt.Println(v2)
                        var v43 uint64 = 433
                        _ = v42
                        _ = v43
                    }
                    _ = v21
                } else {
                    switch "foo bar" {
                    case "hello":
                        v2 -= v2
                    default:
                        var v44 uint32 = 2147483647
                        _ = v44
                    }
                    v45 := GenS6{F0: 2, F1: 889, F2: 2}
                    var v46 float64 = 1e10
                    _ = v45
                    _ = v46
                }
                var v47 float32 = (3.14159 * -2.25)
                if false {
                    v48 := GenS9{F0: 4294967295, F1: v2}
                    _ = v48
                }
                _ = v47
            case -402:
                _ = len(v11)
                v2 -= (-(v2))
            default:
                v2 -= v2
            }
            _ = fi13
            _ = v20
        }
        v2 = (v2 ^ (v2 / (((^(v2))) | 1)))
        if len(v11) > 0 {
            v11[0] = fi12
        }
        _ = fi12
    }
    delete(v3, ("a" + ("\n\t" + "")))
    var v49 int32 = 0
    var v50 int8 = 127
    v50 = v50
    var v51 int32 = -2147483648
    if (1.5 <= (1.5 + 1e10)) {
        var v52 bool = true
        _ = v52
    }
    var v53 int32 = v49
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v50
    _ = v51
    _ = v53
}
