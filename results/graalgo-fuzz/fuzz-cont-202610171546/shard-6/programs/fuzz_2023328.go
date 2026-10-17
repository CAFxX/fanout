package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 uint; F1 uint16; F2 uint16 }
func genFn0(p0 uint64, p1 uint32, p2 int) int16 {
    v4 := []int{128, p2}
    p0 = (^(p0))
    p0 += (p0 | p0)
    var v5 string = "a"
    if false {
        v6 := GenS1{F0: 0, F1: -1}
        _ = len(v4)
        _ = v6
    } else {
        switch "" {
        case "foo bar":
            fmt.Println(p2, p1)
        case "a":
            v4 = append(v4, 42)
        case "":
            var v7 float32 = 0.1
            _ = v7
        default:
            p2 = p2
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 1
}
func genFn1() int {
    var v8 int16 = -32768
    var v9 uint8 = 2
    if true {
        if false {
            v9 = v9
            v10 := GenS1{F0: 23, F1: 1}
            v8 -= (v8 * v8)
            _ = v10
        } else {
            var v11 uint64 = 52
            _ = v11
        }
        if (!true) {
            var v12 uint8 = v9
            v13 := GenS2{F0: 2, F1: 255}
            _ = v12
            _ = v13
        }
    }
    var v14 int64 = -77
    _ = v14
    _ = v8
    _ = v9
    return -655
}

func main() {
    v1 := []int{127, -9223372036854775808, 9223372036854775807, 32767}
    var v2 float32 = 0.1
    var v3 bool = true
    fmt.Println(v2, v3)
    var v15 uint = 712
    var v16 uint64 = 42
    if (861 >= uint64(v15)) {
        v17 := GenS1{F0: 128, F1: 1}
        var v18 float32 = float32(((1.5 * ((1.5 + 0.1) + (3.14159 / (0.1 * 3.14159)))) * (1e10 + ((float64(v2) / -2.25) / float64(v2)))))
        _ = v17
        _ = v18
    } else {
        for fi19 := 0; fi19 < 3; fi19++ {
            v20 := GenS3{F0: 65535, F1: int8(v16)}
            _ = fi19
            _ = v20
        }
        if (941 == -2147483648) {
            v15 = (v15 | (v15 | (v15 & 873)))
            var v21 int = 2147483647
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v21
        } else {
            var v22 uint = (v15 + 451)
            v16 += v16
            _ = v22
        }
        v23 := GenS1{F0: 157, F1: 1}
        _ = v23
    }
    var v24 float32 = 3.14159
    var v25 uint16 = 1
    if v3 {
        switch 4294967295 {
        case 2147483647:
            var v26 string = "a"
            _ = v26
        case -2147483648:
            var v27 float32 = v24
            _ = v27
        case 128:
            if true {
                var v28 int32 = 42
                _ = v28
            } else {
                for fi29 := 0; fi29 < 14; fi29++ {
                    if v3 {
                        v30 := GenS3{F0: 2147483648, F1: 9}
                        _ = v30
                    } else {
                        v2 -= 1e10
                        v31 := GenS4{F0: 1, F1: 992}
                        switch (fi29 / ((389) | 1)) {
                        case -9223372036854775808:
                            fmt.Println(v25)
                        case 32767:
                            var v32 int8 = 127
                            _ = v32
                        }
                        _ = v31
                    }
                    v15 = (v15 << ((v15) & 63))
                    _ = fi29
                }
                for fi33 := 0; fi33 < 14; fi33++ {
                    for fi34 := 0; fi34 < 10; fi34++ {
                        v2 = (0.0 + 0.1)
                        v1 = append(v1, 127)
                        _ = fi34
                    }
                    _ = fi33
                }
            }
        }
        switch (("hello" + "hello") + ("" + "\n\t")) {
        case "\n\t":
            _ = len(v1)
        case "":
            switch "\n\t" {
            case "hello":
                _ = len(v1)
            case "\n\t":
                if (0 == v25) {
                    for fi35 := 0; fi35 < 11; fi35++ {
                        v24 = 1e10
                        _ = fi35
                    }
                }
            case "foo bar":
                v24 -= v24
            }
        case "hello":
            v25 = 640
        }
        if (-151 < -384) {
            if v3 {
                _ = len(v1)
            }
            if v3 {
                fmt.Println(v2, v24, v15)
                for fi36 := 0; fi36 < 20; fi36++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi36
                }
            } else {
                switch genFn1() {
                case 1:
                    var v37 uint8 = 205
                    _ = v37
                default:
                    v16 = (^(v16))
                }
                var v38 float32 = (-2.25 / (((v24 * v24) - v24) - float32(-2.25)))
                _ = v38
            }
            fmt.Println(v3, v3)
        } else {
            v39 := GenS4{F0: 0, F1: (v16 * v16)}
            var v40 int8 = int8(v25)
            switch 128 {
            case 42:
                for fi41 := 0; fi41 < 18; fi41++ {
                    if len(v1) > 0 {
                        v1[0] = (fi41 - (fi41 << ((v15) & 63)))
                    }
                    _ = len(v1)
                    var v42 float64 = (-2.25 / 1.0)
                    _ = fi41
                    _ = v42
                }
            case -2147483648:
                switch -495 {
                case 745:
                    if false {
                        v16 -= (v16 - (v16 >> ((v15) & 63)))
                    } else {
                        v3 = v3
                        v24 = (v2 / (v2 - (float32(0.1) / ((((0.1 - 1e10) * 3.14159) * 0.1) * 3.14159))))
                    }
                default:
                    v16 = 170
                }
            }
            _ = v39
            _ = v40
        }
    } else {
        var v43 float64 = 0.1
        if ((!(false || v3)) && true) {
            v44 := GenS5{F0: v15, F1: (v25 >> (((v15 - (v15 / (((v15 / (((v15 >> (((^(v15))) & 63))) | 1))) | 1)))) & 63)), F2: v25}
            switch (((("\n\t" + ("\n\t" + ("" + ("foo bar" + "\n\t")))) + "\n\t") + "foo bar") + "") {
            case "foo bar":
                if true {
                    v25 += 2
                    var v45 int64 = -853
                    fmt.Println(v2)
                    _ = v45
                }
            default:
                if false {
                    v25 = (v25 - v25)
                    var v46 uint16 = 972
                    v3 = (1 >= uint16(v16))
                    _ = v46
                }
            }
            var v47 uint8 = 0
            _ = v44
            _ = v47
        }
        v24 += ((((v2 + 0.0) * ((0.1 / v24) + v2)) + (-2.25 / v24)) * v24)
        _ = v43
    }
    v48 := GenS3{F0: 54, F1: 106}
    var v49 int = 239
    v24 = float32((-2.25 + 1e10))
    var v50 float32 = -2.25
    v1 = append(v1, genFn1())
    switch (^(v49)) {
    case 1:
        switch (v49 + v49) {
        case 32768:
            v2 = 3.14159
        case 155:
            v49 += 32768
        default:
            v50 -= v24
        }
    case 42:
        for fi51 := 0; fi51 < 9; fi51++ {
            var v52 int = (v49 / ((v49) | 1))
            _ = fi51
            _ = v52
        }
    }
    v53 := GenS3{F0: 42, F1: -102}
    v24 -= 3.14159
    v54 := GenS2{F0: (v25 / ((42) | 1)), F1: (v15 << (((v15 & (v15 << (((^(v15))) & 63)))) & 63))}
    var v55 uint16 = v25
    v56 := GenS4{F0: -1, F1: v16}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v48
    _ = v49
    _ = v50
    _ = v53
    _ = v54
    _ = v55
    _ = v56
}
