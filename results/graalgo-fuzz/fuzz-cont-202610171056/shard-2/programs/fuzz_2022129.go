package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32; F1 uint; F2 int64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int; F1 uint32 }
type GenS7 struct { F0 int8; F1 uint64; F2 int32 }
type GenS8 struct { F0 int64; F1 int8 }
func genFn0(p0 uint8, p1 int16, p2 int) uint16 {
    for fi6 := 0; fi6 < 15; fi6++ {
        p2 = p2
        _ = fi6
    }
    var v7 uint16 = 16
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return v7
}
func genFn1() int8 {
    for fi8 := 0; fi8 < 3; fi8++ {
        fmt.Println(fi8)
        v9 := []int{(fi8 ^ fi8), (fi8 >> ((18446744073709551615) & 63)), (fi8 | fi8), (fi8 << ((65535) & 63))}
        _ = fi8
        _ = v9
    }
    var v10 string = "\n\t"
    if true {
        var v11 string = v10
        v10 = v10
        _ = v11
    }
    _ = v10
    return -1
}

func main() {
    v1 := []string{"hello"}
    var v2 int32 = -1
    var v3 uint = 232
    var v4 uint32 = 917
    v5 := make(map[int]string)
    v5[-326] = "foo bar"
    var v12 string = ""
    v13 := GenS1{F0: 42, F1: (^(v4)), F2: 0}
    switch 928 {
    case -9223372036854775808:
        v14 := GenS1{F0: 42, F1: (v4 >> ((v3) & 63)), F2: 195}
        _ = v14
    case -447:
        var v15 bool = true
        _ = v15
    default:
        _ = len(v1)
    }
    var v16 float32 = float32(3.14159)
    v17 := GenS2{F0: v12}
    var v18 uint16 = genFn0(1, -32768, -1)
    if (false && (v2 != 1)) {
        if (!true) {
            switch v12 {
            case "\n\t":
                v5[128] = v12
            case "hello":
                switch 32767 {
                case -537:
                    if (!(v3 > v3)) {
                        switch 4294967295 {
                        case -27:
                            var v19 int32 = (v2 % ((v2) | 1))
                            _ = v19
                        case -650:
                            fmt.Println(v18, v4, v16)
                        case 100:
                            v20 := GenS2{F0: "\n\t"}
                            _ = v20
                        }
                    }
                case 1:
                    var v21 string = ""
                    _ = v21
                case 72:
                    fmt.Println(v16, v2, v3)
                }
            default:
                fmt.Println(v18, v3)
            }
            if false {
                switch -9223372036854775808 {
                case 42:
                    v22 := GenS2{F0: "\n\t"}
                    _ = v22
                case -1:
                    v3 += uint(v4)
                case 9223372036854775807:
                    var v23 int16 = 32767
                    _ = v23
                default:
                    v24 := GenS2{F0: ""}
                    _ = v24
                }
            } else {
                if false {
                    if (776 <= 917) {
                        var v25 uint = v3
                        var v26 int = 127
                        var v27 float64 = 3.14159
                        _ = v25
                        _ = v26
                        _ = v27
                    } else {
                        if false {
                            v3 = uint(v18)
                        } else {
                            var v28 int = -598
                            if (-2.25 != 0.0) {
                                v12 = ((v12 + "a") + v12)
                            } else {
                                if len(v1) > 0 {
                                    v1[0] = v12
                                }
                                v4 += (^(v4))
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                            _ = v28
                        }
                        var v29 int32 = (v2 | 1)
                        if ((v18 * 32767) != genFn0(232, 32767, 32767)) {
                            v2 = (v2 + (v29 << (((v3 ^ 721)) & 63)))
                            if true {
                                for fi30 := 0; fi30 < 12; fi30++ {
                                    fmt.Println(fi30)
                                    _ = fi30
                                }
                                for fi31 := 0; fi31 < 6; fi31++ {
                                    for fi32 := 0; fi32 < 3; fi32++ {
                                        v2 = (-(v2))
                                        fmt.Println(v4)
                                        _ = fi32
                                    }
                                    v2 -= v2
                                    var v33 uint64 = 18446744073709551615
                                    _ = fi31
                                    _ = v33
                                }
                                switch -727 {
                                case 2147483647:
                                    switch -181 {
                                    case 127:
                                        v29 += (v2 * v29)
                                    default:
                                        var v34 uint32 = 146
                                        _ = v34
                                    }
                                case 128:
                                    fmt.Println(v3, v2, v29)
                                }
                            }
                        } else {
                            var v35 int32 = (-(v29))
                            v12 = (v12 + (("" + v12) + ("hello" + ((("hello" + ("foo bar" + v12)) + v12) + "a"))))
                            v36 := GenS3{F0: 0}
                            _ = v35
                            _ = v36
                        }
                        _ = v29
                    }
                    v37 := GenS4{F0: v4, F1: 42, F2: 1}
                    _ = v37
                } else {
                    v38 := GenS4{F0: v4, F1: (^(v3)), F2: 1}
                    _ = v38
                }
                if true {
                    if true {
                        for fi39 := 0; fi39 < 5; fi39++ {
                            if true {
                                v40 := GenS5{F0: v2}
                                var v41 int32 = (v2 - (v2 << ((v3) & 63)))
                                v41 -= (v2 / (((v2 >> ((v3) & 63))) | 1))
                                _ = v40
                                _ = v41
                            }
                            var v42 int = 0
                            v4 -= (v4 >> (((v3 & 42)) & 63))
                            _ = fi39
                            _ = v42
                        }
                    }
                    var v43 float64 = ((0.1 + -2.25) / -2.25)
                    var v44 string = v12
                    _ = v43
                    _ = v44
                }
            }
        }
        v45 := v5[982]
        switch 127 {
        case 0:
            v12 = v45
        }
        _ = v45
    }
    v2 += v2
    switch -507 {
    case 2147483647:
        var v46 uint64 = 1
        _ = v46
    case 0:
        for fi47 := 0; fi47 < 10; fi47++ {
            if true {
                for fi48 := 0; fi48 < 15; fi48++ {
                    var v49 string = v12
                    var v50 float32 = (v16 / v16)
                    _ = fi48
                    _ = v49
                    _ = v50
                }
                v51 := GenS6{F0: fi47, F1: (^(v4))}
                _ = v51
            } else {
                v52 := GenS7{F0: -18, F1: 481, F2: -1}
                _ = v52
            }
            _ = fi47
        }
    }
    var v53 float64 = 3.14159
    var v54 int8 = genFn1()
    v55 := GenS8{F0: 42, F1: v54}
    var v56 int64 = 241
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v55
    _ = v56
}
