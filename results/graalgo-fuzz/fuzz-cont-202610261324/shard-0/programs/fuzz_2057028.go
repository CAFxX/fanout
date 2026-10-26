package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint16 }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 bool; F1 uint16; F2 int16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint16; F1 uint8; F2 int32 }
type GenS8 struct { F0 uint }

func main() {
    var v1 float64 = 1.5
    v2 := []uint32{2, 0, 65535}
    v3 := make(map[string]string)
    v4 := make(map[uint32]uint32)
    v5 := GenS1{F0: -598, F1: 163, F2: 669}
    if (42 < -1) {
        switch "hello" {
        case "foo bar":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case "hello":
            v1 -= v1
        default:
            var v6 int16 = 0
            _ = v6
        }
        var v7 int32 = -2147483648
        if (0 < 42) {
            for fi8 := 0; fi8 < 12; fi8++ {
                v9 := GenS1{F0: -856, F1: 128, F2: 0}
                switch "foo bar" {
                case "a":
                    var v10 uint64 = 4294967295
                    _ = v10
                case "hello":
                    v11 := GenS2{F0: false, F1: 18446744073709551615}
                    _ = v11
                case "foo bar":
                    delete(v3, "foo bar")
                default:
                    for fi12 := 0; fi12 < 14; fi12++ {
                        if false {
                            v7 -= -728
                            v7 += (-(v7))
                        }
                        var v13 int16 = 0
                        v7 -= -309
                        v1 += ((3.14159 / v1) * v1)
                        v14 := GenS3{F0: 0, F1: "hello"}
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                }
                _ = fi8
                _ = v9
            }
            if len(v2) > 0 {
                v2[0] = 65535
            }
        } else {
            var v15 uint8 = 128
            v2 = append(v2, 0)
            if false {
                v1 -= v1
                for fi16 := 0; fi16 < 19; fi16++ {
                    fmt.Println(v1, v1, v1)
                    _ = fi16
                }
                for fi17 := 0; fi17 < 11; fi17++ {
                    switch fi17 {
                    case 128:
                        fmt.Println(v1, v7, fi17)
                    default:
                        var v18 int64 = -1
                        _ = v18
                    }
                    _ = fi17
                }
            } else {
                for fi19 := 0; fi19 < 11; fi19++ {
                    fmt.Println(fi19)
                    var v20 string = "foo bar"
                    var v21 int16 = 1
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            }
            _ = v15
        }
        _ = v7
    }
    v22 := GenS2{F0: (0.1 <= ((0.1 - -2.25) + (1e10 + 1e10))), F1: 4294967295}
    var v23 uint8 = 128
    v24 := v3["hello"]
    fmt.Println(v23)
    v25 := GenS4{F0: (false || ((true && false) || (!true))), F1: 764, F2: 186}
    for fi26 := 0; fi26 < 17; fi26++ {
        for fi27 := 0; fi27 < 17; fi27++ {
            fmt.Println(fi27)
            switch v24 {
            case "\n\t":
                v28 := GenS5{F0: 389}
                _ = v28
            }
            var v29 int8 = 1
            switch "hello" {
            case "foo bar":
                v1 -= float64((0.0 - (3.14159 / 1e10)))
            case "":
                var v30 int32 = -176
                _ = v30
            case "hello":
                var v31 int8 = v29
                _ = v31
            }
            switch (fi27 * 32767) {
            case 1:
                switch "foo bar" {
                case "":
                    v29 = 42
                case "foo bar":
                    var v32 int64 = -630
                    _ = v32
                case "a":
                    v33 := GenS6{F0: 936}
                    _ = v33
                default:
                    switch -9223372036854775808 {
                    case -302:
                        if (((v1 / 1.0) * ((v1 - 3.14159) - v1)) != (v1 + v1)) {
                            v24 = ""
                            v23 -= v23
                            if (0 <= 42) {
                                fmt.Println(v24)
                                v34 := GenS1{F0: int16(fi27), F1: (v23 % ((v23) | 1)), F2: 32767}
                                switch (^(fi27)) {
                                case 193:
                                    if (true || (2 >= 567)) {
                                        v29 += (v29 & (v29 << ((777) & 63)))
                                        fmt.Println(v1, fi26)
                                    } else {
                                        if false {
                                            v24 = v24
                                            var v35 uint64 = 4294967295
                                            v3[v24] = v24
                                            _ = v35
                                        } else {
                                            v1 = (1e10 / 1.5)
                                            switch (v24 + v24) {
                                            case "a":
                                                fmt.Println(v1, v1)
                                            case "foo bar":
                                                if ((!(-79 == -650)) && (156 > 42)) {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    v24 = "\n\t"
                                                }
                                            default:
                                                _ = len(v3)
                                            }
                                            fmt.Println(v29, v24, v29)
                                        }
                                        if (true && true) {
                                            if true {
                                                var v36 float64 = (-2.25 - ((v1 * v1) * float64(((float32(float64((0.1 * 1.5))) - -2.25) + -2.25))))
                                                fmt.Println(fi27)
                                                v37 := GenS7{F0: 0, F1: v23, F2: 42}
                                                _ = v36
                                                _ = v37
                                            }
                                            v38 := v3[v24]
                                            delete(v4, 682)
                                            var v39 int = (fi27 & fi26)
                                            v23 = v23
                                            _ = v38
                                            _ = v39
                                        } else {
                                            v23 = v23
                                            v1 = (-2.25 * (float64(-2.25) + 1e10))
                                        }
                                        fmt.Println(fi27, v1)
                                    }
                                }
                                if (9223372036854775807 == -9223372036854775808) {
                                    v40 := GenS8{F0: 65535}
                                    _ = v40
                                } else {
                                    var v41 int16 = -1
                                    v42 := GenS3{F0: v29, F1: v24}
                                    v4[4294967295] = 2
                                    fmt.Println(v1, v24, v23)
                                    _ = v41
                                    _ = v42
                                }
                                _ = v34
                            } else {
                                v43 := GenS5{F0: 0}
                                _ = v43
                            }
                        }
                    case -9223372036854775808:
                        v44 := GenS8{F0: 7}
                        _ = v44
                    }
                }
            }
            _ = fi27
            _ = v29
        }
        switch 127 {
        case 4294967295:
            if false {
                fmt.Println(v23)
            }
        case -570:
            if true {
                for fi45 := 0; fi45 < 14; fi45++ {
                    var v46 uint32 = 65535
                    _ = fi45
                    _ = v46
                }
                v1 += v1
            } else {
                var v47 string = "hello"
                _ = v47
            }
        case -1:
            if false {
                v23 = (v23 ^ v23)
                fmt.Println(v1, v1, v1)
            }
        default:
            v23 -= (v23 | 127)
        }
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
