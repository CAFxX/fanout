package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint32; F1 uint16 }
type GenS6 struct { F0 int32; F1 uint; F2 int64 }
type GenS7 struct { F0 int64; F1 int8; F2 uint32 }
type GenS8 struct { F0 int16; F1 uint32; F2 int32 }
type GenS9 struct { F0 uint16; F1 int16 }
type GenS10 struct { F0 int64 }
func genFn0(p0 int, p1 string, p2 uint64) int8 {
    var v5 uint64 = (p2 % ((p2) | 1))
    v6 := GenS1{F0: 255, F1: 227, F2: (((p1 + ("" + ((p1 + p1) + "a"))) + "a") + p1)}
    for fi7 := 0; fi7 < 10; fi7++ {
        v5 = 2
        switch (((("hello" + p1) + "a") + ((("a" + p1) + p1) + p1)) + p1) {
        case "":
            fmt.Println(p0)
        case "hello":
            p1 = "hello"
        default:
            if false {
                var v8 int16 = 685
                for fi9 := 0; fi9 < 5; fi9++ {
                    var v10 uint32 = 712
                    switch p0 {
                    case 9223372036854775807:
                        v11 := GenS1{F0: 233, F1: uint8(v10), F2: p1}
                        switch p1 {
                        case "\n\t":
                            p0 -= (fi7 >> ((0) & 63))
                        }
                        _ = v11
                    default:
                        p2 += (p2 / ((v5) | 1))
                    }
                    _ = fi9
                    _ = v10
                }
                _ = v8
            } else {
                if false {
                    if (3.14159 < 1e10) {
                        v12 := GenS2{F0: 2}
                        if (((1e10 / 3.14159) + 0.1) == ((3.14159 - -2.25) / ((((1e10 / 1.0) - -2.25) + float32(float64(((-2.25 - 1e10) * 1.5)))) + float32(1.5)))) {
                            for fi13 := 0; fi13 < 12; fi13++ {
                                fmt.Println(v5)
                                _ = fi13
                            }
                            var v14 uint16 = 32767
                            _ = v14
                        }
                        _ = v12
                    } else {
                        var v15 int64 = 101
                        switch (p0 % ((fi7) | 1)) {
                        case 32767:
                            if (1 > v15) {
                                fmt.Println(p2, v5)
                            } else {
                                p2 = 18446744073709551615
                            }
                        }
                        if (65535 <= 527) {
                            switch (fi7 % ((9223372036854775807) | 1)) {
                            case 1:
                                v16 := GenS1{F0: 1, F1: 144, F2: "hello"}
                                _ = v16
                            default:
                                v17 := []int{p0, (fi7 - (^(p0))), fi7}
                                _ = v17
                            }
                            var v18 int64 = (v15 * (v15 + -848))
                            _ = v18
                        }
                        _ = v15
                    }
                    switch 241 {
                    case 376:
                        for fi19 := 0; fi19 < 17; fi19++ {
                            if (1.5 < (0.0 / 1e10)) {
                                p2 = (v5 >> ((uint(v5)) & 63))
                            }
                            _ = fi19
                        }
                    case 32768:
                        for fi20 := 0; fi20 < 12; fi20++ {
                            var v21 int32 = -1
                            v22 := GenS3{F0: 851}
                            p2 = v5
                            v23 := []int{}
                            var v24 uint8 = 42
                            _ = fi20
                            _ = v21
                            _ = v22
                            _ = v23
                            _ = v24
                        }
                    }
                } else {
                    var v25 uint64 = (^(p2))
                    _ = v25
                }
                v26 := GenS1{F0: 127, F1: 197, F2: (("foo bar" + "\n\t") + p1)}
                _ = v26
            }
        }
        switch p1 {
        case "hello":
            v27 := []int{-2147483648}
            _ = v27
        case "":
            v28 := GenS1{F0: 128, F1: 42, F2: p1}
            _ = v28
        case "\n\t":
            fmt.Println(p2)
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 127
}
func genFn1(p0 uint32, p1 int32) int16 {
    p1 += (p1 ^ p1)
    if (1 < -574) {
        v29 := GenS4{F0: 255}
        _ = v29
    }
    var v30 uint16 = 605
    if false {
        if true {
            v30 += v30
            v31 := GenS5{F0: 0, F1: 0}
            _ = v31
        } else {
            p0 = 140
        }
        p0 = (p0 - p0)
        v32 := []int{128, int(v30), 9223372036854775807, -2147483648}
        _ = v32
    }
    v30 += v30
    _ = p0
    _ = p1
    _ = v30
    return -1
}

func main() {
    var v1 int8 = 5
    v2 := make(map[string]uint32)
    var v3 int16 = 42
    var v4 int32 = -1
    _ = len(v2)
    var v33 uint16 = 42
    v34 := GenS3{F0: 42}
    for fi35 := 0; fi35 < 7; fi35++ {
        for fi36 := 0; fi36 < 10; fi36++ {
            v37 := GenS2{F0: 1}
            _ = fi36
            _ = v37
        }
        switch fi35 {
        case 343:
            switch (fi35 >> ((387) & 63)) {
            case 1:
                var v38 int32 = 191
                _ = v38
            case 42:
                v39 := GenS6{F0: 1, F1: 42, F2: -355}
                _ = v39
            case 127:
                switch "hello" {
                case "foo bar":
                    if (821 == 154) {
                        var v40 int = int(v33)
                        _ = v40
                    }
                case "\n\t":
                    for fi41 := 0; fi41 < 3; fi41++ {
                        if false {
                            var v42 float32 = (0.1 / -2.25)
                            switch ("" + "") {
                            case "a":
                                if (272 > 10) {
                                    v43 := GenS7{F0: -662, F1: (v1 & v1), F2: 0}
                                    var v44 float32 = 1e10
                                    _ = v43
                                    _ = v44
                                }
                            }
                            _ = v42
                        } else {
                            switch "a" {
                            case "":
                                v45 := GenS8{F0: v3, F1: 42, F2: (v4 | int32(v1))}
                                _ = v45
                            case "foo bar":
                                var v46 int64 = -425
                                _ = v46
                            case "a":
                                var v47 int = 9223372036854775807
                                _ = v47
                            }
                            v48 := GenS9{F0: v33, F1: v3}
                            v2[(("" + "a") + (("a" + "foo bar") + ""))] = 65535
                            _ = v48
                        }
                        v33 = (v33 ^ (v33 % ((v33) | 1)))
                        if (true || (-594 < -9223372036854775808)) {
                            v33 -= 385
                            var v49 uint = 551
                            v3 = int16(v1)
                            _ = v49
                        } else {
                            if true {
                                var v50 float32 = (0.1 + 3.14159)
                                _ = v50
                            } else {
                                switch 1 {
                                case 32768:
                                    v4 = (v4 >> ((4294967295) & 63))
                                }
                            }
                        }
                        _ = fi41
                    }
                }
            default:
                var v51 bool = true
                _ = v51
            }
        default:
            v1 += v1
        }
        _ = fi35
    }
    if false {
        v1 = (v1 >> ((65535) & 63))
        _ = len(v2)
    }
    v52 := GenS10{F0: 0}
    fmt.Println(v3, v4)
    var v53 int = 127
    var v54 uint16 = v33
    for fi55 := 0; fi55 < 12; fi55++ {
        for fi56 := 0; fi56 < 5; fi56++ {
            v1 += v1
            var v57 float64 = 1.5
            var v58 float32 = 0.0
            _ = fi56
            _ = v57
            _ = v58
        }
        _ = fi55
    }
    switch (v53 / ((v53) | 1)) {
    case 127:
        var v59 int64 = -469
        _ = v59
    default:
        v60 := GenS7{F0: -614, F1: (v1 + (v1 % (((-(v1))) | 1))), F2: 42}
        _ = v60
    }
    if (v33 < (v33 / (((v33 * v54)) | 1))) {
        v1 = (v1 & v1)
    } else {
        v33 += v33
        _ = len(v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v52
    _ = v53
    _ = v54
}
