package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int16 }
type GenS2 struct { F0 int; F1 int; F2 uint32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64; F1 int }
type GenS5 struct { F0 int8; F1 int64; F2 uint32 }
type GenS6 struct { F0 bool; F1 uint64 }
type GenS7 struct { F0 int8; F1 int16 }
type GenS8 struct { F0 int64; F1 uint64 }
func genFn0() bool {
    var v7 string = ""
    var v8 int64 = -9223372036854775808
    if ((false || (361 < 0)) && false) {
        v8 -= v8
    } else {
        var v9 int8 = 127
        for fi10 := 0; fi10 < 1; fi10++ {
            fmt.Println(v8, fi10)
            fmt.Println(v8, v8)
            _ = fi10
        }
        fmt.Println(v7)
        _ = v9
    }
    _ = v7
    _ = v8
    return (false || ((!(false || true)) || true))
}

func main() {
    v1 := []int{-1, 4294967295, 42, 32767}
    var v2 uint32 = 4294967295
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]int)
    var v5 float64 = (-2.25 * (1.5 / float64(float32(3.14159))))
    var v6 float32 = 3.14159
    v2 = v2
    fmt.Println(v2, v5)
    v11 := GenS1{F0: "", F1: -23, F2: 1000}
    switch "a" {
    case "\n\t":
        fmt.Println(v2, v2)
    }
    if (2 > 176) {
        var v12 string = ("" + "")
        _ = v12
    }
    var v13 int8 = 0
    for fi14 := 0; fi14 < 18; fi14++ {
        switch (fi14 | fi14) {
        case -9223372036854775808:
            if ((109 < 65535) || false) {
                if (true && true) {
                    if true {
                        for fi15 := 0; fi15 < 12; fi15++ {
                            if (1 < 2147483647) {
                                var v16 int64 = -1
                                fmt.Println(v6, v6, v16)
                                var v17 int64 = -1
                                _ = v16
                                _ = v17
                            } else {
                                var v18 uint32 = v2
                                if false {
                                    var v19 int64 = -249
                                    if true {
                                        var v20 int8 = v13
                                        if (!(12 != 127)) {
                                            delete(v4, 765)
                                            v6 += (0.0 / (((v6 * v6) * float32(-2.25)) + -2.25))
                                            var v21 uint = 18446744073709551615
                                            v4[v2] = (fi15 / ((fi14) | 1))
                                            fmt.Println(v2)
                                            _ = v21
                                        }
                                        var v22 string = (("foo bar" + "foo bar") + "foo bar")
                                        _ = v20
                                        _ = v22
                                    } else {
                                        v1 = append(v1, -867)
                                    }
                                    v18 = 14
                                    var v23 bool = true
                                    _ = v19
                                    _ = v23
                                } else {
                                    if true {
                                        if (0 < 424) {
                                            var v24 uint8 = 0
                                            _ = v24
                                        } else {
                                            _ = len(v3)
                                            var v25 uint = 42
                                            _ = v25
                                        }
                                        var v26 int8 = (v13 | 0)
                                        var v27 int = -834
                                        _ = v26
                                        _ = v27
                                    }
                                    _ = len(v3)
                                    v18 += 0
                                    if true {
                                        if (!false) {
                                            v4[v18] = (fi14 << ((42) & 63))
                                        }
                                    }
                                }
                                var v28 int64 = 9223372036854775807
                                _ = v18
                                _ = v28
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            var v29 int16 = -584
                            _ = fi15
                            _ = v29
                        }
                        delete(v3, 4294967295)
                        for fi30 := 0; fi30 < 18; fi30++ {
                            var v31 uint = 2
                            _ = fi30
                            _ = v31
                        }
                    }
                }
            } else {
                _ = len(v1)
                v2 = (^(v2))
            }
        }
        if true {
            var v32 int = (fi14 / ((-15) | 1))
            var v33 int16 = 32767
            _ = v32
            _ = v33
        } else {
            v2 -= v2
            if (false || (-2147483648 > -1)) {
                var v34 uint8 = 133
                if true {
                    v13 += (-(v13))
                    switch "\n\t" {
                    case "a":
                        v2 = 65535
                    }
                } else {
                    v35 := GenS2{F0: (fi14 & (fi14 % (((-(fi14))) | 1))), F1: -18, F2: v2}
                    _ = v35
                }
                for fi36 := 0; fi36 < 17; fi36++ {
                    v6 = v6
                    var v37 int16 = -32768
                    v38 := GenS2{F0: (fi36 & 32767), F1: fi14, F2: (v2 | v2)}
                    _ = fi36
                    _ = v37
                    _ = v38
                }
                _ = v34
            } else {
                for fi39 := 0; fi39 < 9; fi39++ {
                    v40 := GenS3{F0: -2147483648}
                    if true {
                        switch "\n\t" {
                        case "\n\t":
                            v41 := GenS3{F0: 0}
                            _ = v41
                        case "foo bar":
                            switch "a" {
                            case "a":
                                switch "hello" {
                                case "a":
                                    _ = len(v1)
                                case "hello":
                                    if (fi14 == (fi14 / (((fi14 % ((fi39) | 1))) | 1))) {
                                        v42 := GenS4{F0: 9223372036854775807, F1: (fi39 | (-(fi39)))}
                                        fmt.Println(v5, v13, v2)
                                        v6 = ((1.5 + v6) * 1.5)
                                        _ = v42
                                    } else {
                                        fmt.Println(v6, v6)
                                        var v43 float32 = 0.1
                                        _ = v43
                                    }
                                default:
                                    var v44 int16 = 890
                                    _ = v44
                                }
                            default:
                                fmt.Println(fi39, v13, fi39)
                            }
                        }
                        fmt.Println(v2, v5, v13)
                    }
                    switch "foo bar" {
                    case "":
                        var v45 int16 = 1
                        _ = v45
                    }
                    fmt.Println(v5)
                    _ = fi39
                    _ = v40
                }
            }
        }
        _ = fi14
    }
    v46 := GenS5{F0: (-(v13)), F1: 0, F2: (^(v2))}
    v2 -= 459
    if ((!true) && (false || false)) {
        fmt.Println(v2, v13, v2)
        switch 9223372036854775807 {
        case 0:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        default:
            fmt.Println(v2, v13)
        }
    }
    for fi47 := 0; fi47 < 15; fi47++ {
        for fi48 := 0; fi48 < 18; fi48++ {
            v6 = (3.14159 / 1e10)
            var v49 float32 = float32(((v5 - (v5 * (v5 - v5))) / -2.25))
            _ = fi48
            _ = v49
        }
        v6 += (v6 * ((3.14159 * ((v6 + float32(v5)) - (v6 - v6))) / 1.0))
        var v50 float64 = v5
        _ = fi47
        _ = v50
    }
    v51 := GenS6{F0: (657 >= uint(v2)), F1: 1}
    var v52 uint32 = v2
    var v53 int8 = (v13 & 127)
    if false {
        var v54 uint = 0
        var v55 int32 = -14
        v13 = v53
        _ = v54
        _ = v55
    }
    v5 -= -2.25
    if true {
        if false {
            if false {
                v5 = v5
                if true {
                    var v56 int64 = -420
                    _ = v56
                }
            }
        }
    } else {
        if true {
            var v57 uint = 0
            v2 += 0
            _ = v57
        } else {
            for fi58 := 0; fi58 < 2; fi58++ {
                var v59 uint8 = 242
                _ = fi58
                _ = v59
            }
        }
        v60 := GenS7{F0: 1, F1: -1}
        if false {
            var v61 uint64 = 293
            v62 := GenS1{F0: (("a" + ("foo bar" + "hello")) + ("foo bar" + ("hello" + ("hello" + "\n\t")))), F1: v13, F2: -942}
            v63 := GenS8{F0: 812, F1: (v61 ^ 851)}
            _ = v61
            _ = v62
            _ = v63
        }
        _ = v60
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v51
    _ = v52
    _ = v53
    _ = v6
}
