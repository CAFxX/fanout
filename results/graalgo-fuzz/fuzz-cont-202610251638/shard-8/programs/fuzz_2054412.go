package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 int8; F1 uint64; F2 int32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 bool; F1 int32; F2 int32 }
type GenS7 struct { F0 int16 }
func genFn0(p0 uint8, p1 string, p2 int16) int {
    v4 := []int{399, 4294967295}
    fmt.Println(p2, p1)
    v4 = append(v4, -9223372036854775808)
    if true {
        p1 = p1
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -2147483648
}
func genFn1() int {
    var v5 int32 = -1
    v6 := GenS1{F0: 1, F1: v5}
    if false {
        var v7 int64 = 0
        _ = v7
    }
    var v8 uint64 = 18446744073709551615
    for fi9 := 0; fi9 < 17; fi9++ {
        v10 := []int{fi9, 1}
        var v11 int32 = (v5 % ((v5) | 1))
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = v5
    _ = v6
    _ = v8
    return -2147483648
}

func main() {
    var v1 int8 = 42
    v2 := make(map[uint32]int)
    v3 := []int{1, 34, 127}
    for fi12 := 0; fi12 < 20; fi12++ {
        fmt.Println(fi12, fi12)
        fmt.Println(v1)
        _ = fi12
    }
    var v13 string = "hello"
    fmt.Println(v1, v13)
    fmt.Println(v13, v1)
    var v14 uint16 = 0
    for fi15 := 0; fi15 < 1; fi15++ {
        var v16 int32 = -866
        if (false || (2 > 81)) {
            for fi17 := 0; fi17 < 13; fi17++ {
                v1 += v1
                v1 = v1
                if ((fi17 ^ genFn0(1, (v13 + ((("foo bar" + v13) + "a") + (((((v13 + (("a" + "") + v13)) + v13) + "a") + (v13 + v13)) + v13))), -32768)) != (-(fi15))) {
                    v14 += (v14 | (v14 % ((65535) | 1)))
                    if true {
                        fmt.Println(v13, v13)
                        var v18 uint64 = 2
                        _ = v18
                    } else {
                        delete(v2, 1)
                        var v19 uint16 = 42
                        _ = v19
                    }
                    v3 = append(v3, (fi15 % ((fi15) | 1)))
                } else {
                    var v20 int8 = (^(v1))
                    v21 := GenS1{F0: (v14 ^ 2), F1: (v16 / ((v16) | 1))}
                    v1 -= v20
                    var v22 float32 = 1e10
                    _ = v20
                    _ = v21
                    _ = v22
                }
                _ = fi17
            }
            if true {
                var v23 float32 = float32((1.5 / 1.5))
                switch v13 {
                case "":
                    if (-623 != 9223372036854775807) {
                        if true {
                            var v24 int16 = -1
                            delete(v2, 2147483647)
                            for fi25 := 0; fi25 < 19; fi25++ {
                                var v26 uint32 = 736
                                _ = fi25
                                _ = v26
                            }
                            _ = v24
                        } else {
                            var v27 uint = 255
                            var v28 int = (fi15 % ((0) | 1))
                            _ = v27
                            _ = v28
                        }
                        v2[2147483648] = (fi15 & fi15)
                        if true {
                            for fi29 := 0; fi29 < 19; fi29++ {
                                var v30 int = (fi15 - fi29)
                                v1 -= (v1 - -108)
                                v31 := GenS2{F0: (v1 >> ((2) & 63))}
                                _ = fi29
                                _ = v30
                                _ = v31
                            }
                            v32 := GenS2{F0: v1}
                            _ = v32
                        }
                    } else {
                        v33 := GenS3{F0: -1, F1: v14}
                        var v34 int32 = 452
                        _ = v33
                        _ = v34
                    }
                case "hello":
                    if (true || (float32(3.14159) < 3.14159)) {
                        fmt.Println(v13, fi15, fi15)
                        var v35 string = v13
                        var v36 string = v35
                        _ = v35
                        _ = v36
                    } else {
                        v37 := GenS3{F0: 42, F1: 944}
                        var v38 uint32 = 0
                        _ = v37
                        _ = v38
                    }
                case "\n\t":
                    for fi39 := 0; fi39 < 2; fi39++ {
                        var v40 int = 9223372036854775807
                        _ = fi39
                        _ = v40
                    }
                }
                _ = v23
            } else {
                switch v13 {
                case "a":
                    delete(v2, 2147483647)
                default:
                    if false {
                        for fi41 := 0; fi41 < 10; fi41++ {
                            v14 -= 65535
                            v16 = int32(fi41)
                            _ = fi41
                        }
                        var v42 uint = 4294967295
                        v42 -= 612
                        _ = v42
                    } else {
                        fmt.Println(v1, v16)
                        v1 = v1
                        var v43 bool = true
                        _ = v43
                    }
                }
                for fi44 := 0; fi44 < 15; fi44++ {
                    v13 = (v13 + (v13 + ("foo bar" + v13)))
                    if (172 > 1) {
                        if false {
                            var v45 int64 = int64(v16)
                            if false {
                                if true {
                                    v46 := GenS4{F0: (v1 % ((111) | 1)), F1: 18446744073709551615, F2: (v16 - v16)}
                                    v47 := GenS5{F0: 4294967295}
                                    _ = v46
                                    _ = v47
                                }
                                v14 = 42
                            } else {
                                v13 = (v13 + v13)
                                v45 -= 9223372036854775807
                            }
                            _ = v45
                        }
                    } else {
                        v48 := GenS6{F0: (((!false) && false) || ((v14 ^ (v14 | 692)) == v14)), F1: v16, F2: v16}
                        v16 = (-(v16))
                        if (v14 <= (v14 >> ((708) & 63))) {
                            v13 = (v13 + v13)
                            if (!false) {
                                v14 = v14
                            }
                            v49 := GenS3{F0: -1, F1: v14}
                            _ = v49
                        } else {
                            v16 = v16
                        }
                        _ = v48
                    }
                    var v50 uint64 = 2
                    _ = fi44
                    _ = v50
                }
                for fi51 := 0; fi51 < 15; fi51++ {
                    var v52 bool = true
                    v52 = (64 <= 629)
                    var v53 int16 = 1000
                    _ = fi51
                    _ = v52
                    _ = v53
                }
            }
        }
        _ = fi15
        _ = v16
    }
    var v54 int = 0
    switch v13 {
    case "a":
        delete(v2, 4294967295)
    case "hello":
        v3 = append(v3, (v54 - (v54 | (v54 << ((18446744073709551615) & 63)))))
    }
    v55 := GenS7{F0: -321}
    v13 = ("" + (v13 + v13))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v56 int16 = 419
    if false {
        v14 -= v14
        var v57 uint64 = 878
        _ = v57
    }
    fmt.Println(v54, v54)
    if len(v3) > 0 {
        v3[0] = -55
    }
    v54 += v54
    var v58 float32 = ((-2.25 / 1e10) * 1e10)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v54:", v54)
    fmt.Println("v56:", v56)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v54
    _ = v55
    _ = v56
    _ = v58
}
