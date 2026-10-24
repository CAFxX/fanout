package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 uint64; F1 int64; F2 uint8 }
type GenS3 struct { F0 uint16; F1 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int16; F1 int32 }
type GenS6 struct { F0 uint16; F1 int16; F2 uint8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int32; F1 uint8 }
func genFn0(p0 uint8, p1 int8, p2 int64) uint8 {
    v5 := []int{-64, 933, 128, 85}
    if (false || (0 > 65535)) {
        for fi6 := 0; fi6 < 9; fi6++ {
            p0 = p0
            _ = fi6
        }
        var v7 int32 = -883
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (^(p0))
}
func genFn1() uint16 {
    var v8 int8 = 0
    v9 := GenS1{F0: 280, F1: 0}
    var v10 string = ""
    var v11 uint = 72
    var v12 int8 = (v8 >> ((v11) & 63))
    _ = v10
    _ = v11
    _ = v12
    _ = v8
    _ = v9
    return 32767
}

func main() {
    var v1 int = 2147483647
    var v2 int8 = -31
    var v3 int = 0
    v4 := make(map[string]bool)
    fmt.Println(v1)
    switch "hello" {
    case "a":
        var v13 uint32 = 4294967295
        _ = v13
    default:
        v14 := v4["\n\t"]
        _ = v14
    }
    if false {
        if true {
            for fi15 := 0; fi15 < 4; fi15++ {
                v1 += (fi15 - (v1 | 42))
                v16 := GenS2{F0: 4294967295, F1: 669, F2: 70}
                _ = fi15
                _ = v16
            }
            for fi17 := 0; fi17 < 10; fi17++ {
                v18 := GenS3{F0: 1, F1: 774}
                fmt.Println(v3, v1)
                v19 := v4["hello"]
                _ = fi17
                _ = v18
                _ = v19
            }
            switch (("\n\t" + "hello") + ((("" + "\n\t") + "foo bar") + "")) {
            case "foo bar":
                if (!true) {
                    v1 -= v3
                    if ("a" == (("a" + "a") + ("a" + "\n\t"))) {
                        for fi20 := 0; fi20 < 17; fi20++ {
                            if true {
                                var v21 float32 = (0.1 * 0.1)
                                _ = v21
                            } else {
                                v22 := GenS3{F0: 32767, F1: 65535}
                                _ = v22
                            }
                            v23 := GenS1{F0: 354, F1: 926}
                            _ = fi20
                            _ = v23
                        }
                        var v24 uint = 255
                        switch -1 {
                        case 0:
                            v25 := GenS4{F0: ("hello" + "hello")}
                            _ = v25
                        case -634:
                            v26 := v4["\n\t"]
                            _ = v26
                        default:
                            v2 -= -126
                        }
                        _ = v24
                    }
                    fmt.Println(v2, v2)
                } else {
                    v27 := []int{(v1 >> ((0) & 63)), v1, v3, v3}
                    var v28 uint32 = 2147483648
                    _ = len(v4)
                    _ = v27
                    _ = v28
                }
            case "\n\t":
                fmt.Println(v3)
            case "hello":
                switch 127 {
                case 0:
                    if true {
                        fmt.Println(v3, v1, v1)
                        v3 += (v3 % (((-(v3))) | 1))
                        for fi29 := 0; fi29 < 6; fi29++ {
                            var v30 int8 = v2
                            _ = fi29
                            _ = v30
                        }
                    }
                case 35:
                    if (((^(v1)) >= (-(v3))) && false) {
                        v1 -= v3
                        v2 = (v2 | 0)
                        if (false && (!(-627 <= 32767))) {
                            v31 := GenS4{F0: ((("a" + "a") + "") + "\n\t")}
                            _ = v31
                        }
                    } else {
                        v32 := GenS4{F0: ("" + ("\n\t" + ""))}
                        _ = v32
                    }
                default:
                    v33 := []int{v1, v3, v1, v1}
                    _ = v33
                }
            default:
                for fi34 := 0; fi34 < 4; fi34++ {
                    if false {
                        var v35 int = (fi34 / (((v1 - -1)) | 1))
                        _ = v35
                    } else {
                        var v36 int = (fi34 % ((-9223372036854775808) | 1))
                        v1 += fi34
                        _ = v36
                    }
                    var v37 uint16 = 898
                    _ = fi34
                    _ = v37
                }
            }
        } else {
            fmt.Println(v1, v1, v1)
            if (1 != 42) {
                v3 -= (v1 >> ((0) & 63))
            } else {
                for fi38 := 0; fi38 < 8; fi38++ {
                    v39 := GenS5{F0: 324, F1: -230}
                    var v40 uint8 = 128
                    v2 += v2
                    _ = fi38
                    _ = v39
                    _ = v40
                }
                var v41 int64 = 42
                _ = v41
            }
            fmt.Println(v2, v2, v1)
        }
        fmt.Println(v3, v2)
        for fi42 := 0; fi42 < 5; fi42++ {
            v2 = v2
            _ = fi42
        }
    }
    if true {
        v43 := v4[""]
        v44 := GenS6{F0: 32767, F1: 1000, F2: genFn0(127, 1, 42)}
        v45 := GenS7{F0: 2147483647}
        _ = v43
        _ = v44
        _ = v45
    } else {
        v46 := []int{-552, -2147483648}
        _ = v46
    }
    var v47 uint8 = 0
    v48 := GenS8{F0: -109, F1: 30}
    var v49 bool = false
    v49 = (false && (((0 > v2) || v49) && v49))
    for fi50 := 0; fi50 < 4; fi50++ {
        v2 = (v2 / ((v2) | 1))
        for fi51 := 0; fi51 < 10; fi51++ {
            _ = len(v4)
            v52 := []int{9223372036854775807}
            v2 -= int8(v47)
            _ = fi51
            _ = v52
        }
        _ = fi50
    }
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
}
