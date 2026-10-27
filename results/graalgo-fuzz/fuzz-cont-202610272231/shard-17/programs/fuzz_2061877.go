package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 string; F1 uint64; F2 string }
type GenS3 struct { F0 int16; F1 bool; F2 uint8 }
type GenS4 struct { F0 uint8; F1 uint8; F2 int32 }
func genFn0() uint64 {
    var v6 int16 = 122
    fmt.Println(v6)
    if (3.14159 <= 1e10) {
        var v7 uint = 771
        _ = v7
    }
    v6 = (v6 * 32767)
    _ = v6
    return 42
}

func main() {
    var v1 uint64 = 2
    var v2 uint8 = 255
    var v3 float32 = (0.0 / 1.5)
    v4 := []uint32{42, 65535, 2147483648, 42}
    var v5 uint32 = 65535
    var v8 uint16 = 348
    v2 -= v2
    v2 = (v2 & (v2 & 255))
    var v9 string = (("hello" + "hello") + "\n\t")
    var v10 int32 = 1
    fmt.Println(v5, v8, v10)
    for fi11 := 0; fi11 < 10; fi11++ {
        if true {
            if false {
                for fi12 := 0; fi12 < 16; fi12++ {
                    var v13 bool = false
                    _ = fi12
                    _ = v13
                }
                for fi14 := 0; fi14 < 13; fi14++ {
                    if len(v4) > 0 {
                        v4[0] = (v5 | (v5 / (((v5 / (((v5 ^ (^(v5)))) | 1))) | 1)))
                    }
                    v10 = v10
                    v15 := GenS1{F0: 4294967295, F1: true}
                    switch -1 {
                    case -645:
                        switch ((("" + "a") + v9) + v9) {
                        case "foo bar":
                            switch "foo bar" {
                            case "":
                                if (uint8(v5) < uint8(v8)) {
                                    if len(v4) > 0 {
                                        v4[0] = (v5 % (((v5 & (v5 / ((v5) | 1)))) | 1))
                                    }
                                    switch -893 {
                                    case -666:
                                        var v16 string = ("hello" + "a")
                                        _ = v16
                                    }
                                } else {
                                    fmt.Println(v9)
                                    if (4294967295 >= 0) {
                                        fmt.Println(v5, v10, v2)
                                        v5 = (^(v5))
                                    }
                                    var v17 int8 = 1
                                    v8 = (v8 ^ (v8 / (((v8 % ((32767) | 1))) | 1)))
                                    if (921 < v1) {
                                        v9 = ""
                                    }
                                    _ = v17
                                }
                            case "a":
                                var v18 string = ((v9 + "\n\t") + v9)
                                _ = v18
                            case "hello":
                                var v19 string = "hello"
                                _ = v19
                            default:
                                if true {
                                    if ((!false) || false) {
                                        v3 = v3
                                        if true {
                                            var v20 uint8 = (v2 - 255)
                                            _ = v20
                                        }
                                        var v21 uint16 = v8
                                        _ = v21
                                    } else {
                                        switch v9 {
                                        case "\n\t":
                                            if false {
                                                if false {
                                                    var v22 uint8 = (v2 >> ((1) & 63))
                                                    _ = v22
                                                }
                                                var v23 int = (^(fi11))
                                                _ = v23
                                            }
                                        default:
                                            fmt.Println(v2)
                                        }
                                        v4 = append(v4, v5)
                                    }
                                    var v24 int8 = 42
                                    fmt.Println(v8)
                                    v1 += (^(v1))
                                    _ = v24
                                }
                            }
                        case "\n\t":
                            var v25 int64 = 42
                            _ = v25
                        }
                    case 1:
                        var v26 uint64 = genFn0()
                        _ = v26
                    case 867:
                        var v27 uint64 = (v1 >> ((919) & 63))
                        _ = v27
                    }
                    _ = fi14
                    _ = v15
                }
                var v28 int8 = 95
                _ = v28
            } else {
                if (!true) {
                    _ = len(v4)
                    var v29 uint16 = v8
                    if true {
                        var v30 bool = true
                        _ = v30
                    } else {
                        v31 := GenS2{F0: ("" + v9), F1: (v1 & (v1 - 803)), F2: ("a" + "foo bar")}
                        _ = v31
                    }
                    _ = v29
                }
                var v32 int = fi11
                v4 = append(v4, (v5 * v5))
                _ = v32
            }
        } else {
            if false {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            } else {
                v33 := GenS1{F0: uint(v2), F1: ((!true) || true)}
                var v34 string = "\n\t"
                _ = v33
                _ = v34
            }
            var v35 uint16 = uint16(v2)
            v36 := GenS2{F0: "", F1: v1, F2: v9}
            _ = v35
            _ = v36
        }
        v8 += v8
        for fi37 := 0; fi37 < 3; fi37++ {
            if (!false) {
                v38 := GenS3{F0: -237, F1: true, F2: (v2 * 133)}
                if len(v4) > 0 {
                    v4[0] = (v5 / ((v5) | 1))
                }
                var v39 int8 = 1
                _ = v38
                _ = v39
            } else {
                fmt.Println(v3, v10)
                var v40 string = v9
                v10 -= v10
                _ = v40
            }
            var v41 float32 = 3.14159
            if len(v4) > 0 {
                v4[0] = (^(v5))
            }
            if false {
                var v42 uint8 = (v2 | (v2 ^ v2))
                _ = v42
            }
            _ = fi37
            _ = v41
        }
        _ = fi11
    }
    fmt.Println(v5, v1, v8)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi43 := 0; fi43 < 13; fi43++ {
        switch 42 {
        case 629:
            v5 -= (v5 + v5)
        case 127:
            fmt.Println(v8)
        case 128:
            v8 = 1
        }
        var v44 uint16 = (v8 - v8)
        var v45 float32 = 0.0
        _ = fi43
        _ = v44
        _ = v45
    }
    if (false && ((v2 & v2) == v2)) {
        if true {
            v9 = "a"
        } else {
            v3 = v3
            if len(v4) > 0 {
                v4[0] = v5
            }
            v46 := GenS4{F0: (v2 % (((v2 >> ((758) & 63))) | 1)), F1: (^(v2)), F2: v10}
            _ = v46
        }
    } else {
        v1 += v1
        for fi47 := 0; fi47 < 3; fi47++ {
            v9 = v9
            _ = fi47
        }
        if true {
            v2 += v2
        } else {
            _ = len(v4)
        }
    }
    v10 = v10
    for fi48 := 0; fi48 < 15; fi48++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi49 := 0; fi49 < 15; fi49++ {
            v1 = v1
            _ = fi49
        }
        v4 = append(v4, (v5 / ((v5) | 1)))
        _ = fi48
    }
    fmt.Println(v10, v3, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
