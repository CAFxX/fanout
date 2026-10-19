package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 int64 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 uint64; F1 uint8; F2 int8 }
type GenS4 struct { F0 uint64; F1 int16; F2 uint }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int; F1 uint; F2 int32 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 bool; F1 uint8 }
type GenS10 struct { F0 uint32; F1 uint }

func main() {
    var v1 int = 4294967295
    var v2 uint = 0
    var v3 int16 = 1000
    v4 := []int{(v1 ^ 150), v1, 127, (-(v1))}
    var v5 uint16 = 42
    v1 -= 2147483647
    v1 -= (^(v1))
    v2 -= (v2 * (v2 % ((v2) | 1)))
    v6 := GenS1{F0: (false || true), F1: 1, F2: 42}
    if true {
        v7 := GenS1{F0: ((false && (255 > 1)) && true), F1: -128, F2: 42}
        var v8 uint64 = 1
        fmt.Println(v1)
        _ = v7
        _ = v8
    }
    var v9 float32 = 3.14159
    var v10 float64 = 3.14159
    v10 -= v10
    if (!true) {
        v11 := GenS1{F0: false, F1: -39, F2: -489}
        _ = v11
    } else {
        switch v1 {
        case 32768:
            v1 += -671
        case -925:
            v1 = v1
        default:
            v12 := GenS2{F0: 42, F1: 1}
            _ = v12
        }
        var v13 string = "foo bar"
        if false {
            v14 := GenS3{F0: 1, F1: 42, F2: 100}
            v1 = (^(v1))
            var v15 int16 = -617
            _ = v14
            _ = v15
        }
        _ = v13
    }
    var v16 int32 = 0
    switch ("" + (("" + "") + ((("\n\t" + "\n\t") + "\n\t") + "\n\t"))) {
    case "foo bar":
        switch ("" + "a") {
        case "foo bar":
            v1 -= (v1 | v1)
        case "hello":
            if true {
                var v17 int32 = (v16 | v16)
                _ = v17
            } else {
                if false {
                    v4 = append(v4, (v1 / (((v1 % ((v1) | 1))) | 1)))
                    v2 += (^(v2))
                    v18 := GenS1{F0: (true || (v1 <= 2147483647)), F1: 0, F2: 0}
                    _ = v18
                } else {
                    var v19 uint8 = 0
                    _ = v19
                }
                for fi20 := 0; fi20 < 3; fi20++ {
                    if (!(false || (46 > 21))) {
                        fmt.Println(v1, v2)
                    } else {
                        if (!(42 <= 35)) {
                            var v21 int = (v1 | v1)
                            if (v5 < (v5 / (((^(v5))) | 1))) {
                                if len(v4) > 0 {
                                    v4[0] = v21
                                }
                                fmt.Println(v21)
                                if false {
                                    var v22 string = ("" + "a")
                                    v22 = "\n\t"
                                    _ = v22
                                } else {
                                    for fi23 := 0; fi23 < 7; fi23++ {
                                        if len(v4) > 0 {
                                            v4[0] = (-(fi20))
                                        }
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        _ = fi23
                                    }
                                    if len(v4) > 0 {
                                        v4[0] = fi20
                                    }
                                }
                            }
                            _ = v21
                        } else {
                            v24 := GenS3{F0: 42, F1: 1, F2: 1}
                            if false {
                                _ = len(v4)
                                switch "a" {
                                case "a":
                                    if true {
                                        switch v1 {
                                        case 0:
                                            if (int64(fi20) <= -1) {
                                                v10 += v10
                                            } else {
                                                fmt.Println(v9)
                                                for fi25 := 0; fi25 < 5; fi25++ {
                                                    if (fi25 > fi20) {
                                                        if (!(v9 != 0.0)) {
                                                            if true {
                                                                if ((-2.25 - float32((v10 / v10))) != v9) {
                                                                    v5 = v5
                                                                    if true {
                                                                        _ = len(v4)
                                                                        v26 := GenS2{F0: -9223372036854775808, F1: 0}
                                                                        _ = v26
                                                                    } else {
                                                                        switch ("a" + "a") {
                                                                        case "foo bar":
                                                                            var v27 float64 = 0.0
                                                                            _ = v27
                                                                        case "\n\t":
                                                                            v28 := GenS4{F0: 586, F1: (v3 | -281), F2: v2}
                                                                            _ = v28
                                                                        case "hello":
                                                                            v3 -= int16(fi25)
                                                                        default:
                                                                            var v29 int16 = v3
                                                                            _ = v29
                                                                        }
                                                                    }
                                                                    if ((!true) || true) {
                                                                        if len(v4) > 0 {
                                                                            v4[0] = (v1 & fi25)
                                                                        }
                                                                    }
                                                                    switch ("hello" + "\n\t") {
                                                                    case "hello":
                                                                        v30 := GenS5{F0: 752}
                                                                        _ = v30
                                                                    case "":
                                                                        v2 = v2
                                                                    case "a":
                                                                        v31 := GenS1{F0: false, F1: 127, F2: int64(v16)}
                                                                        _ = v31
                                                                    default:
                                                                        fmt.Println(v16)
                                                                    }
                                                                    var v32 uint32 = 4294967295
                                                                    _ = v32
                                                                }
                                                                if ((!false) && false) {
                                                                    v3 = v3
                                                                } else {
                                                                    if false {
                                                                        var v33 int32 = v16
                                                                        if (1e10 > float32(float64(-2.25))) {
                                                                            var v34 uint32 = 2147483647
                                                                            _ = v34
                                                                        }
                                                                        _ = v33
                                                                    } else {
                                                                        v5 -= 42
                                                                        var v35 uint8 = 240
                                                                        _ = v35
                                                                    }
                                                                    var v36 uint32 = 65535
                                                                    _ = v36
                                                                }
                                                            } else {
                                                                v3 += int16(v5)
                                                                v37 := GenS3{F0: 4294967295, F1: 127, F2: -1}
                                                                v38 := GenS6{F0: v1, F1: (v2 | v2), F2: v16}
                                                                v4 = append(v4, -1)
                                                                _ = v37
                                                                _ = v38
                                                            }
                                                            if (1 < 526) {
                                                                var v39 bool = (!("a" <= "hello"))
                                                                v40 := GenS5{F0: 350}
                                                                _ = v39
                                                                _ = v40
                                                            } else {
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                            }
                                                            v2 = uint(v5)
                                                        }
                                                    }
                                                    if (int64(v5) == -490) {
                                                        fmt.Println(v9, fi20, fi20)
                                                        var v41 float64 = v10
                                                        v42 := GenS7{F0: (-(v3))}
                                                        _ = v41
                                                        _ = v42
                                                    }
                                                    var v43 string = "hello"
                                                    _ = fi25
                                                    _ = v43
                                                }
                                                var v44 uint8 = 90
                                                _ = v44
                                            }
                                        }
                                        v45 := GenS7{F0: 1}
                                        fmt.Println(v16, v3, v2)
                                        _ = v45
                                    } else {
                                        var v46 float64 = v10
                                        fmt.Println(fi20, v5, fi20)
                                        _ = v46
                                    }
                                case "foo bar":
                                    switch "a" {
                                    case "foo bar":
                                        v2 = (v2 ^ (v2 >> ((v2) & 63)))
                                    case "\n\t":
                                        switch "" {
                                        case "hello":
                                            v5 -= (v5 >> ((2) & 63))
                                        case "\n\t":
                                            _ = len(v4)
                                        }
                                    }
                                case "hello":
                                    var v47 bool = ((!true) || (v1 >= (fi20 << (((v2 >> ((255) & 63))) & 63))))
                                    _ = v47
                                default:
                                    if (100 != 100) {
                                        var v48 int = int(v3)
                                        v5 -= (v5 | (v5 / ((v5) | 1)))
                                        _ = v48
                                    } else {
                                        if (127 == 128) {
                                            var v49 int32 = (v16 + (v16 - v16))
                                            _ = v49
                                        }
                                    }
                                }
                            } else {
                                if false {
                                    v1 += -329
                                    var v50 int32 = (^(v16))
                                    _ = v50
                                }
                            }
                            v51 := GenS6{F0: (fi20 ^ (fi20 / (((^(fi20))) | 1))), F1: (v2 % ((255) | 1)), F2: v16}
                            _ = v24
                            _ = v51
                        }
                        v52 := GenS8{F0: 65535}
                        _ = v52
                    }
                    if false {
                        if len(v4) > 0 {
                            v4[0] = (v1 & (v1 & v1))
                        }
                    }
                    v10 -= (v10 - v10)
                    _ = fi20
                }
            }
        case "a":
            v4 = append(v4, v1)
        default:
            v2 += v2
        }
    case "a":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case "\n\t":
        fmt.Println(v2, v5, v1)
    default:
        for fi53 := 0; fi53 < 16; fi53++ {
            if true {
                v16 = (-(v16))
                fmt.Println(fi53)
                v54 := GenS9{F0: false, F1: 170}
                _ = v54
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v16 = -510
                var v55 int16 = v3
                _ = v55
            }
            _ = fi53
        }
    }
    _ = len(v4)
    switch (v1 * (v1 % (((v1 * 32767)) | 1))) {
    case 728:
        var v56 string = "hello"
        _ = v56
    case -2147483648:
        fmt.Println(v3, v1, v9)
    }
    if (!false) {
        if true {
            var v57 int8 = 27
            fmt.Println(v10)
            _ = v57
        }
        switch "foo bar" {
        case "a":
            v58 := GenS10{F0: 2, F1: v2}
            _ = v58
        case "\n\t":
            _ = len(v4)
        default:
            if ((true && false) && true) {
                var v59 int = -9223372036854775808
                _ = v59
            } else {
                var v60 float32 = ((v9 - v9) * v9)
                switch (^(v1)) {
                case 677:
                    fmt.Println(v5, v16, v1)
                case 42:
                    fmt.Println(v16, v1)
                case -367:
                    if false {
                        var v61 float64 = float64(float32((((v10 * v10) / 3.14159) - (v10 - 1e10))))
                        var v62 uint16 = 255
                        _ = v61
                        _ = v62
                    }
                }
                _ = v60
            }
        }
        var v63 int32 = (v16 + (v16 << ((v2) & 63)))
        _ = v63
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
