package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 int; F2 uint32 }
type GenS4 struct { F0 int16; F1 bool; F2 uint32 }
type GenS5 struct { F0 int64; F1 uint8; F2 string }
type GenS6 struct { F0 string; F1 bool }
type GenS7 struct { F0 uint32; F1 int32 }

func main() {
    var v1 uint64 = 153
    var v2 uint16 = 32767
    var v3 string = ""
    var v4 int32 = 42
    var v5 uint = 18446744073709551615
    for fi6 := 0; fi6 < 3; fi6++ {
        for fi7 := 0; fi7 < 1; fi7++ {
            if false {
                if (true && false) {
                    fmt.Println(v1)
                    v8 := GenS1{F0: true, F1: 127}
                    _ = v8
                } else {
                    v9 := GenS2{F0: 4294967295}
                    v2 = (v2 * v2)
                    _ = v9
                }
                var v10 uint8 = uint8(v2)
                var v11 int32 = v4
                var v12 bool = (!true)
                _ = v10
                _ = v11
                _ = v12
            } else {
                v13 := GenS2{F0: 635}
                var v14 uint32 = 128
                _ = v13
                _ = v14
            }
            var v15 float64 = 1.5
            v16 := GenS3{F0: ((v2 * (v2 + 42)) != 506), F1: fi7, F2: 2}
            if (42 == 127) {
                var v17 uint16 = (v2 * 571)
                _ = v17
            }
            if ((!false) && (!(-2147483648 > 0))) {
                if (true && (true || true)) {
                    switch "a" {
                    case "foo bar":
                        var v18 uint = 65535
                        _ = v18
                    default:
                        v2 = (v2 - v2)
                    }
                    var v19 bool = (false && ((!false) || ((v4 > v4) && ((1 == 100) && (-1 == -967)))))
                    var v20 uint8 = 42
                    v1 -= v1
                    _ = v19
                    _ = v20
                }
                var v21 int64 = 0
                _ = v21
            }
            if true {
                v22 := []int{(-(fi6)), (fi6 % ((fi6) | 1)), (fi7 << ((508) & 63))}
                _ = v22
            } else {
                v15 += float64((0.0 * 3.14159))
                var v23 uint = v5
                _ = v23
            }
            _ = fi7
            _ = v15
            _ = v16
        }
        if (int32(v1) <= (v4 << (((v5 & (^(v5)))) & 63))) {
            var v24 int = fi6
            v4 += v4
            _ = v24
        }
        _ = fi6
    }
    if (!(!true)) {
        for fi25 := 0; fi25 < 20; fi25++ {
            if false {
                v5 += (^(v5))
                v26 := []int{(fi25 << ((v5) & 63)), 128}
                _ = v26
            } else {
                v27 := GenS4{F0: int16(fi25), F1: false, F2: 4294967295}
                _ = v27
            }
            _ = fi25
        }
    }
    var v28 float32 = 3.14159
    for fi29 := 0; fi29 < 9; fi29++ {
        if true {
            v5 += 255
            var v30 uint = v5
            fmt.Println(v4, fi29)
            _ = v30
        } else {
            v2 += (v2 - v2)
        }
        _ = fi29
    }
    for fi31 := 0; fi31 < 6; fi31++ {
        if (!false) {
            v28 -= (v28 * (v28 * (1.5 - (-2.25 + v28))))
            if true {
                v32 := []int{fi31, 127}
                switch v3 {
                case "foo bar":
                    var v33 uint16 = v2
                    _ = v33
                case "\n\t":
                    v3 = ("" + v3)
                default:
                    for fi34 := 0; fi34 < 16; fi34++ {
                        if ((0.0 * 1e10) >= 1e10) {
                            if (1e10 < v28) {
                                switch int(v2) {
                                case 400:
                                    v32 = append(v32, -97)
                                case 468:
                                    v28 = 0.0
                                default:
                                    switch (fi34 / ((fi31) | 1)) {
                                    case 652:
                                        fmt.Println(v2)
                                    }
                                }
                                if (42 == v2) {
                                    if ((true || (v1 > v1)) && true) {
                                        v2 = (v2 & v2)
                                        var v35 uint16 = v2
                                        _ = v35
                                    } else {
                                        v1 -= (v1 % (((^(v1))) | 1))
                                        if len(v32) > 0 {
                                            v32[0] = 0
                                        }
                                    }
                                    if false {
                                        fmt.Println(v4)
                                        v36 := GenS5{F0: 957, F1: 127, F2: (((v3 + v3) + "foo bar") + ((((v3 + v3) + v3) + v3) + "foo bar"))}
                                        v28 = (0.0 * (v28 + (3.14159 - v28)))
                                        fmt.Println(v1, v4)
                                        var v37 uint64 = (v1 + v1)
                                        _ = v36
                                        _ = v37
                                    } else {
                                        _ = len(v32)
                                        v38 := GenS3{F0: true, F1: 0, F2: 163}
                                        _ = v38
                                    }
                                    if false {
                                        v39 := GenS3{F0: false, F1: (fi34 / ((fi31) | 1)), F2: 0}
                                        v40 := GenS6{F0: v3, F1: false}
                                        var v41 int = fi31
                                        _ = v39
                                        _ = v40
                                        _ = v41
                                    }
                                    if (fi31 <= (fi31 / (((fi34 % (((fi31 >> (((v5 & (v5 ^ (^(v5))))) & 63))) | 1))) | 1))) {
                                        if true {
                                            v5 += v5
                                            if false {
                                                var v42 int = 2147483647
                                                fmt.Println(v2, v5, v42)
                                                var v43 uint8 = 255
                                                _ = v42
                                                _ = v43
                                            } else {
                                                v5 -= (^(v5))
                                                switch "hello" {
                                                case "":
                                                    v44 := GenS4{F0: int16(v2), F1: false, F2: 2147483647}
                                                    _ = v44
                                                case "\n\t":
                                                    var v45 int8 = 0
                                                    v5 = v5
                                                    v46 := GenS7{F0: 65535, F1: (v4 + 1)}
                                                    _ = v45
                                                    _ = v46
                                                case "hello":
                                                    var v47 uint16 = 239
                                                    _ = v47
                                                }
                                                var v48 uint32 = 318
                                                _ = v48
                                            }
                                        } else {
                                            fmt.Println(v2)
                                            v5 -= (v5 + (v5 >> ((4294967295) & 63)))
                                            if ((v5 ^ (^(v5))) >= v5) {
                                                var v49 bool = true
                                                v2 -= v2
                                                _ = v49
                                            }
                                        }
                                        var v50 int32 = v4
                                        _ = v50
                                    } else {
                                        fmt.Println(fi31, fi31, fi34)
                                        var v51 int = -44
                                        var v52 bool = (v4 < (v4 - (v4 | v4)))
                                        var v53 uint32 = 4294967295
                                        _ = v51
                                        _ = v52
                                        _ = v53
                                    }
                                    v32 = append(v32, fi34)
                                } else {
                                    fmt.Println(fi31, v3, fi31)
                                    _ = len(v32)
                                }
                                if true {
                                    v2 = v2
                                    var v54 int = fi31
                                    fmt.Println(v4, v1)
                                    _ = v54
                                }
                            }
                            var v55 uint16 = v2
                            _ = v55
                        }
                        v3 = v3
                        if len(v32) >= 2 {
                            v32 = v32[1:len(v32)]
                        }
                        _ = fi34
                    }
                }
                _ = v32
            }
            v56 := GenS7{F0: 769, F1: v4}
            _ = v56
        } else {
            v57 := GenS1{F0: false, F1: 42}
            _ = v57
        }
        _ = fi31
    }
    v1 = v1
    if (false && (-1 < -1)) {
        var v58 float64 = (3.14159 * 0.1)
        _ = v58
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
