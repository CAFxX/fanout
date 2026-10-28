package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 uint8; F1 int8 }
type GenS5 struct { F0 uint64; F1 int64; F2 string }

func main() {
    var v1 int32 = 42
    var v2 uint64 = 237
    var v3 float64 = 0.0
    var v4 bool = false
    var v5 float64 = v3
    v6 := []int{-1}
    var v7 uint8 = 74
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if false {
        v7 += 19
        v4 = true
        v6 = append(v6, 313)
    }
    var v8 int16 = 317
    switch "foo bar" {
    case "hello":
        if len(v6) > 0 {
            v6[0] = 9223372036854775807
        }
    }
    if (-438 <= (^(v8))) {
        var v9 int16 = 0
        v10 := GenS1{F0: 128, F1: (v7 / ((v7) | 1)), F2: (v1 & v1)}
        _ = v10
        _ = v9
    } else {
        v11 := GenS1{F0: 4294967295, F1: 42, F2: int32(v2)}
        _ = v11
    }
    v2 -= (v2 - (v2 >> ((284) & 63)))
    for fi12 := 0; fi12 < 19; fi12++ {
        if v4 {
            if (int(v7) >= fi12) {
                if true {
                    var v13 uint8 = 2
                    fmt.Println(v1, v2, v13)
                    _ = v13
                } else {
                    if ((!v4) && (!(v7 == (v7 & 127)))) {
                        _ = len(v6)
                        for fi14 := 0; fi14 < 13; fi14++ {
                            v1 = (^(v1))
                            _ = fi14
                        }
                        v6 = append(v6, 4294967295)
                    } else {
                        var v15 string = ("foo bar" + ("a" + ("\n\t" + ("foo bar" + "a"))))
                        var v16 int = fi12
                        _ = v15
                        _ = v16
                    }
                    if false {
                        var v17 uint32 = 4294967295
                        _ = v17
                    }
                    if v4 {
                        v18 := GenS2{F0: 4294967295}
                        _ = v18
                    }
                }
            }
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            if len(v6) > 0 {
                v6[0] = fi12
            }
        }
        if v4 {
            _ = len(v6)
            for fi19 := 0; fi19 < 15; fi19++ {
                v7 -= (v7 + (^(v7)))
                v20 := GenS3{F0: 65535, F1: ("\n\t" + "a")}
                _ = fi19
                _ = v20
            }
        } else {
            if v4 {
                var v21 uint16 = 2
                var v22 int16 = v8
                _ = v21
                _ = v22
            } else {
                v4 = ((v2 / ((v2) | 1)) == v2)
                if (((1e10 + -2.25) + 1e10) >= (3.14159 / float32(v3))) {
                    if v4 {
                        fmt.Println(v7, v8, v8)
                    } else {
                        for fi23 := 0; fi23 < 11; fi23++ {
                            if ((0 >= 338) || true) {
                                if v4 {
                                    v24 := GenS4{F0: (v7 ^ (v7 ^ uint8(v2))), F1: -28}
                                    if ((!((3.14159 + float32((float64(1e10) / v3))) >= 3.14159)) && (2147483647 == (v1 - (v1 - v1)))) {
                                        fmt.Println(v1)
                                    } else {
                                        var v25 int16 = v8
                                        _ = v25
                                    }
                                    var v26 int = 2147483647
                                    _ = v24
                                    _ = v26
                                }
                                var v27 uint = 655
                                _ = v27
                            }
                            _ = fi23
                        }
                        v7 += v7
                    }
                    for fi28 := 0; fi28 < 12; fi28++ {
                        if (v8 >= 0) {
                            var v29 uint64 = v2
                            v5 -= 0.1
                            _ = v29
                        }
                        v7 -= 2
                        _ = fi28
                    }
                    for fi30 := 0; fi30 < 15; fi30++ {
                        var v31 uint = 2
                        _ = fi30
                        _ = v31
                    }
                }
            }
            for fi32 := 0; fi32 < 19; fi32++ {
                if (!(34 != 42)) {
                    if ((fi12 & 9223372036854775807) >= (fi12 + (fi12 % (((-(fi12))) | 1)))) {
                        var v33 string = "a"
                        switch "hello" {
                        case "\n\t":
                            if (4294967295 != 2) {
                                var v34 string = ("hello" + "a")
                                var v35 uint8 = 127
                                v8 = (v8 + (v8 >> ((783) & 63)))
                                _ = v34
                                _ = v35
                            }
                        case "hello":
                            if (3.14159 != v5) {
                                switch v33 {
                                case "\n\t":
                                    switch (^(fi12)) {
                                    case -993:
                                        if v4 {
                                            _ = len(v6)
                                            var v36 bool = (v4 && (1 < -1))
                                            if len(v6) > 0 {
                                                v6[0] = fi12
                                            }
                                            _ = v36
                                        } else {
                                            var v37 float32 = (-2.25 - 3.14159)
                                            v38 := GenS1{F0: int(v1), F1: uint8(v2), F2: (v1 ^ 42)}
                                            fmt.Println(v5, v3)
                                            _ = v37
                                            _ = v38
                                        }
                                    case 687:
                                        if (v7 != 117) {
                                            var v39 uint16 = 0
                                            _ = v39
                                        } else {
                                            v2 -= v2
                                            v40 := GenS5{F0: 382, F1: 42, F2: "hello"}
                                            v4 = (v4 || ((914 >= 517) && false))
                                            _ = v40
                                        }
                                    case 128:
                                        if len(v6) > 0 {
                                            v6[0] = fi12
                                        }
                                    default:
                                        v3 = v5
                                    }
                                case "a":
                                    v3 = -2.25
                                case "hello":
                                    var v41 int = (fi32 + (fi12 & (fi32 - (^(fi32)))))
                                    _ = v41
                                }
                                v42 := GenS3{F0: 957, F1: "a"}
                                _ = v42
                            } else {
                                var v43 uint16 = 233
                                v33 = (v33 + v33)
                                if ((^(v2)) >= (v2 + v2)) {
                                    v2 -= (v2 ^ (v2 << ((2) & 63)))
                                    if (!v4) {
                                        v43 -= (v43 | 32767)
                                        var v44 int16 = v8
                                        v6 = append(v6, (fi32 ^ (-(fi12))))
                                        _ = v44
                                    }
                                } else {
                                    v45 := GenS4{F0: 189, F1: 0}
                                    var v46 int16 = -32768
                                    fmt.Println(v43, v46)
                                    _ = v45
                                    _ = v46
                                }
                                _ = v43
                            }
                        default:
                            v47 := GenS1{F0: 438, F1: 0, F2: (v1 ^ (v1 | (-(v1))))}
                            _ = v47
                        }
                        v5 = v3
                        _ = v33
                    }
                }
                if true {
                    v5 = float64(-2.25)
                }
                _ = fi32
            }
        }
        _ = fi12
    }
    var v48 uint64 = 407
    if (0 == 0) {
        for fi49 := 0; fi49 < 18; fi49++ {
            v1 += -981
            _ = fi49
        }
    }
    var v50 float32 = 3.14159
    fmt.Println(v5, v2, v8)
    var v51 uint8 = 128
    _ = len(v6)
    v3 += 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v50
    _ = v51
    _ = v6
    _ = v7
    _ = v8
}
