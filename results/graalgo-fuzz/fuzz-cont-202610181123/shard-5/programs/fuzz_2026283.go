package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 bool }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 bool; F1 int64; F2 string }
type GenS5 struct { F0 int16; F1 uint32 }
type GenS6 struct { F0 int64 }
func genFn0() int64 {
    switch 4294967295 {
    case 32767:
        var v4 int64 = -755
        _ = v4
    case 127:
        var v5 uint16 = 677
        _ = v5
    case -2147483648:
        var v6 string = "foo bar"
        _ = v6
    }
    var v7 int32 = 1
    var v8 uint32 = 65535
    v9 := GenS1{F0: 42, F1: 18446744073709551615, F2: false}
    switch 1 {
    case 32768:
        switch 0 {
        case 4294967295:
            v10 := []int{32768, -1, 128, -103}
            _ = v10
        case 32767:
            var v11 float32 = 0.0
            _ = v11
        case 42:
            v12 := GenS2{F0: v7, F1: v8}
            _ = v12
        default:
            for fi13 := 0; fi13 < 19; fi13++ {
                var v14 uint = uint(v8)
                v7 = (v7 % (((v7 * (v7 | (v7 ^ (^(v7)))))) | 1))
                _ = fi13
                _ = v14
            }
        }
    default:
        v7 += (v7 >> ((986) & 63))
    }
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    v1 := make(map[int]string)
    var v2 uint64 = 18446744073709551615
    v3 := []uint32{42, 0, 1, 443}
    if false {
        v15 := GenS1{F0: -9223372036854775808, F1: v2, F2: ((1.5 * 0.0) != -2.25)}
        v2 = v2
        _ = v15
    } else {
        _ = len(v3)
        if true {
            fmt.Println(v2)
            switch -1 {
            case 475:
                var v16 int64 = -1
                _ = v16
            case 9223372036854775807:
                v17 := GenS2{F0: 917, F1: 4294967295}
                _ = v17
            default:
                if true {
                    if true {
                        fmt.Println(v2)
                        fmt.Println(v2)
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    if true {
                        v18 := v1[1]
                        for fi19 := 0; fi19 < 14; fi19++ {
                            v20 := GenS1{F0: -775, F1: v2, F2: false}
                            var v21 uint8 = 42
                            var v22 int32 = -2147483648
                            _ = fi19
                            _ = v20
                            _ = v21
                            _ = v22
                        }
                        v23 := GenS3{F0: 2, F1: 128}
                        _ = v18
                        _ = v23
                    }
                    var v24 float64 = 1e10
                    _ = v24
                } else {
                    v3 = append(v3, 1)
                }
            }
            switch 0 {
            case 9223372036854775807:
                var v25 int64 = -9223372036854775808
                _ = v25
            case -206:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            default:
                fmt.Println(v2)
            }
        } else {
            switch ("foo bar" + "") {
            case "\n\t":
                var v26 float64 = float64(float32(1.5))
                _ = v26
            default:
                if (501 <= 673) {
                    for fi27 := 0; fi27 < 4; fi27++ {
                        for fi28 := 0; fi28 < 10; fi28++ {
                            v2 += v2
                            if false {
                                v29 := GenS2{F0: -1, F1: 4294967295}
                                v2 -= v2
                                _ = v29
                            } else {
                                if true {
                                    var v30 int32 = 42
                                    v2 = (v2 + v2)
                                    if true {
                                        fmt.Println(fi27)
                                        var v31 int8 = 127
                                        if (255 >= 0) {
                                            var v32 int8 = -128
                                            _ = v32
                                        }
                                        _ = v31
                                    }
                                    _ = v30
                                }
                            }
                            v33 := GenS4{F0: false, F1: int64(v2), F2: "a"}
                            _ = fi28
                            _ = v33
                        }
                        v34 := GenS5{F0: 1000, F1: 1}
                        if (((^(v2)) == 4294967295) || false) {
                            for fi35 := 0; fi35 < 10; fi35++ {
                                var v36 uint16 = 255
                                _ = fi35
                                _ = v36
                            }
                        }
                        _ = fi27
                        _ = v34
                    }
                    var v37 uint = 65535
                    v37 += (v37 / ((65535) | 1))
                    _ = v37
                }
            }
            v2 += v2
            v3 = append(v3, 4294967295)
        }
        for fi38 := 0; fi38 < 20; fi38++ {
            switch "\n\t" {
            case "\n\t":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            _ = fi38
        }
    }
    switch -9223372036854775808 {
    case 0:
        var v39 string = (("a" + "hello") + "foo bar")
        _ = v39
    case -623:
        fmt.Println(v2)
    case 4294967295:
        var v40 float32 = 1.5
        _ = v40
    default:
        switch ("a" + "a") {
        case "foo bar":
            v41 := GenS4{F0: (-885 < 229), F1: -355, F2: ""}
            _ = v41
        default:
            fmt.Println(v2)
        }
    }
    switch 32768 {
    case -374:
        v2 = (v2 | (v2 >> ((0) & 63)))
    case 127:
        v2 = (^(v2))
    case -203:
        v1[2147483647] = (((("" + "") + "") + "") + (("a" + "hello") + (("\n\t" + ("\n\t" + "")) + "foo bar")))
    }
    for fi42 := 0; fi42 < 12; fi42++ {
        v2 = 42
        var v43 int64 = -1
        _ = fi42
        _ = v43
    }
    v2 += v2
    v2 = (v2 % ((v2) | 1))
    var v44 uint32 = 0
    v45 := GenS6{F0: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v44
    _ = v45
}
