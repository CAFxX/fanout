package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 uint8; F1 uint8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint64; F1 bool; F2 bool }
type GenS7 struct { F0 int8; F1 int64; F2 bool }
type GenS8 struct { F0 uint8 }
func genFn0() uint64 {
    var v7 int32 = 321
    if (0.1 < 0.0) {
        fmt.Println(v7)
    } else {
        for fi8 := 0; fi8 < 12; fi8++ {
            var v9 string = ("\n\t" + "foo bar")
            _ = fi8
            _ = v9
        }
        if true {
            if true {
                v7 = (v7 * v7)
                v10 := GenS1{F0: 65535, F1: true}
                _ = v10
            }
        }
    }
    switch "\n\t" {
    case "hello":
        switch "a" {
        case "a":
            v11 := []int{}
            _ = v11
        default:
            switch 42 {
            case 127:
                var v12 float64 = (3.14159 + 0.0)
                _ = v12
            case 765:
                fmt.Println(v7)
            case 42:
                var v13 uint = 1
                _ = v13
            }
        }
    case "foo bar":
        v7 += 1
    case "a":
        v14 := []int{-2147483648}
        _ = v14
    }
    if (!(-2.25 > (float32(((3.14159 / -2.25) / 3.14159)) - ((1.5 - -2.25) * 0.0)))) {
        fmt.Println(v7)
        v7 = (^(v7))
        v7 = 42
    } else {
        v7 += 41
    }
    v15 := GenS1{F0: 854, F1: ((!(255 <= 255)) && (false || (1 <= 1)))}
    _ = v15
    _ = v7
    return 869
}
func genFn1() uint8 {
    var v16 int16 = 32767
    v17 := []int{}
    for fi18 := 0; fi18 < 7; fi18++ {
        for fi19 := 0; fi19 < 15; fi19++ {
            v20 := GenS2{F0: 9223372036854775807, F1: 0}
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    if true {
        v16 = (v16 / (((v16 | 42)) | 1))
        fmt.Println(v16)
        var v21 float64 = (-2.25 / float64(float32((1e10 / (3.14159 - -2.25)))))
        _ = v21
    }
    _ = v16
    _ = v17
    return 1
}

func main() {
    var v1 int32 = 42
    var v2 int16 = 167
    var v3 string = "hello"
    v4 := []string{"hello"}
    var v5 int32 = (-(v1))
    var v6 uint64 = 4294967295
    switch 127 {
    case -20:
        if ((v6 - 4294967295) >= (^(v6))) {
            v22 := GenS1{F0: 65535, F1: false}
            if true {
                if (319 < -293) {
                    switch -222 {
                    case 32767:
                        for fi23 := 0; fi23 < 6; fi23++ {
                            var v24 uint64 = v6
                            v25 := GenS3{F0: 838, F1: v2}
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                    case -9223372036854775808:
                        v4 = append(v4, v3)
                    case -562:
                        fmt.Println(v5)
                    default:
                        v3 = v3
                    }
                    v2 = (v2 | (v2 * v2))
                    var v26 uint32 = 509
                    _ = v26
                } else {
                    var v27 float64 = (((0.1 / (float64(0.0) + (1.5 / ((0.1 / -2.25) * 1e10)))) / 1.5) - float64(3.14159))
                    var v28 uint = 0
                    _ = v27
                    _ = v28
                }
                if true {
                    if (0.1 >= -2.25) {
                        _ = len(v4)
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    } else {
                        fmt.Println(v2, v5, v6)
                    }
                    var v29 uint16 = 32767
                    var v30 uint16 = v29
                    _ = v29
                    _ = v30
                } else {
                    switch v3 {
                    case "hello":
                        fmt.Println(v5, v5)
                    case "foo bar":
                        for fi31 := 0; fi31 < 3; fi31++ {
                            v2 = v2
                            _ = fi31
                        }
                    default:
                        for fi32 := 0; fi32 < 16; fi32++ {
                            v2 = (-(v2))
                            fmt.Println(fi32, v3)
                            var v33 uint = 42
                            _ = fi32
                            _ = v33
                        }
                    }
                }
                v34 := GenS4{F0: 156, F1: 42}
                _ = v34
            } else {
                if false {
                    var v35 bool = (true && ((-2.25 != 0.1) && (true && (!(0 >= 42)))))
                    v3 = (((v3 + v3) + v3) + v3)
                    _ = v35
                }
            }
            _ = v22
        } else {
            switch v3 {
            case "":
                if false {
                    var v36 uint = 18446744073709551615
                    _ = v36
                }
            default:
                var v37 uint16 = 42
                _ = v37
            }
            var v38 uint16 = 42
            _ = v38
        }
    default:
        v39 := GenS5{F0: 109}
        _ = v39
    }
    v40 := GenS2{F0: 127, F1: (^(v1))}
    if ((v2 >= (^(v2))) && (((18446744073709551615 == 1) && true) || true)) {
        fmt.Println(v6)
        for fi41 := 0; fi41 < 17; fi41++ {
            v1 = (-(v1))
            v1 += (v1 / (((v1 * (v5 & (v5 << ((65535) & 63))))) | 1))
            if (1 != 97) {
                if len(v4) > 0 {
                    v4[0] = v3
                }
                fmt.Println(v1, v2, v6)
            }
            _ = fi41
        }
    }
    if true {
        if ("hello" > v3) {
            fmt.Println(v1)
            v3 = ("a" + ((v3 + v3) + ("" + v3)))
            fmt.Println(v5)
        }
        var v42 uint64 = genFn0()
        _ = v42
    } else {
        if true {
            fmt.Println(v3, v6)
            var v43 float32 = 1e10
            v3 = ("a" + v3)
            _ = v43
        }
        v2 += (-(v2))
        var v44 uint16 = 0
        _ = v44
    }
    v4 = append(v4, v3)
    if (!false) {
        var v45 float32 = (3.14159 + 0.0)
        var v46 string = "foo bar"
        fmt.Println(v46, v45, v45)
        _ = v45
        _ = v46
    }
    v47 := GenS6{F0: (v6 % (((^(v6))) | 1)), F1: (false || true), F2: (v2 >= (^(v2)))}
    var v48 int16 = v2
    v49 := GenS7{F0: 0, F1: -1, F2: (v3 >= v3)}
    if false {
        fmt.Println(v1, v3)
    } else {
        var v50 uint = 1
        _ = v50
    }
    v1 -= (v5 << ((0) & 63))
    var v51 int64 = -644
    for fi52 := 0; fi52 < 4; fi52++ {
        v1 = (v5 % ((v5) | 1))
        if true {
            var v53 uint64 = v6
            var v54 uint8 = 31
            for fi55 := 0; fi55 < 18; fi55++ {
                v5 += (v5 | v5)
                _ = fi55
            }
            _ = v53
            _ = v54
        }
        _ = fi52
    }
    v56 := GenS6{F0: (v6 * genFn0()), F1: false, F2: true}
    if false {
        var v57 uint32 = 684
        _ = v57
    }
    var v58 int8 = 127
    switch int(v48) {
    case -530:
        v59 := GenS8{F0: 168}
        _ = v59
    case 32767:
        fmt.Println(v6)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v51
    _ = v56
    _ = v58
    _ = v6
}
