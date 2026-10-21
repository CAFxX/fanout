package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint32; F1 string; F2 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int32; F1 bool; F2 uint64 }
type GenS6 struct { F0 uint16; F1 uint16; F2 uint64 }
func genFn0(p0 int32, p1 int8, p2 int64) int64 {
    p1 -= p1
    fmt.Println(p2, p0)
    if (!true) {
        p0 = p0
    }
    var v5 uint64 = 18446744073709551615
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 float64 = 0.1
        if true {
            fmt.Println(p0)
            for fi8 := 0; fi8 < 15; fi8++ {
                fmt.Println(fi8, p1, p0)
                var v9 uint8 = 102
                v10 := GenS1{F0: fi6, F1: 4294967295}
                _ = fi8
                _ = v10
                _ = v9
            }
        }
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return int64(p1)
}

func main() {
    var v1 int = -542
    var v2 uint64 = 0
    var v3 int16 = 1000
    var v4 uint = 1
    var v11 bool = false
    v12 := []int{v1}
    for fi13 := 0; fi13 < 17; fi13++ {
        fmt.Println(v3)
        v14 := GenS2{F0: 251, F1: "\n\t", F2: -510}
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        v11 = (393 != 65535)
        fmt.Println(v11, v4)
        _ = fi15
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        var v17 int32 = 42
        fmt.Println(v4)
        var v18 float32 = 1e10
        _ = fi16
        _ = v17
        _ = v18
    }
    if false {
        var v19 int32 = 441
        _ = v19
    }
    for fi20 := 0; fi20 < 5; fi20++ {
        switch "hello" {
        case "":
            v3 = (^(v3))
        case "hello":
            for fi21 := 0; fi21 < 3; fi21++ {
                switch "foo bar" {
                case "\n\t":
                    v22 := GenS3{F0: ("hello" + "hello")}
                    _ = v22
                case "":
                    if len(v12) > 0 {
                        v12[0] = 4294967295
                    }
                    var v23 float32 = ((float32(1e10) * (((((1.5 + float32(0.1)) - -2.25) / 0.1) * 0.0) / -2.25)) + 0.0)
                    if ((v4 + (v4 + (v4 % ((v4) | 1)))) != (v4 & (v4 + v4))) {
                        if (1 == 843) {
                            v11 = true
                        }
                        if len(v12) >= 2 {
                            v12 = v12[1:len(v12)]
                        }
                    } else {
                        var v24 float32 = -2.25
                        var v25 uint = v4
                        _ = v24
                        _ = v25
                    }
                    _ = v23
                case "a":
                    var v26 bool = (!(v3 == v3))
                    _ = v26
                }
                _ = fi21
            }
        default:
            fmt.Println(v2, v4, v3)
        }
        var v27 uint8 = 74
        _ = fi20
        _ = v27
    }
    if (v11 && true) {
        var v28 int16 = 1000
        switch ("foo bar" + "hello") {
        case "hello":
            v12 = append(v12, v1)
        }
        _ = v28
    }
    var v29 uint64 = v2
    v30 := GenS1{F0: 1, F1: 42}
    if (42 > 2) {
        if (71 != 1) {
            for fi31 := 0; fi31 < 6; fi31++ {
                switch 32767 {
                case 771:
                    for fi32 := 0; fi32 < 13; fi32++ {
                        if v11 {
                            if false {
                                switch ((("foo bar" + ("a" + ((("" + "foo bar") + "foo bar") + ("a" + "a")))) + ("hello" + (("hello" + "\n\t") + ("hello" + "hello")))) + "") {
                                case "\n\t":
                                    v2 -= v2
                                case "foo bar":
                                    if (1.5 >= 0.0) {
                                        var v33 string = "foo bar"
                                        _ = v33
                                    }
                                default:
                                    switch 42 {
                                    case 32768:
                                        v3 = (v3 / ((146) | 1))
                                        v11 = v11
                                    }
                                }
                                switch "" {
                                case "a":
                                    v3 -= 487
                                case "":
                                    switch "a" {
                                    case "hello":
                                        switch v1 {
                                        case 9223372036854775807:
                                            v12 = append(v12, fi31)
                                        case -9223372036854775808:
                                            if len(v12) >= 2 {
                                                v12 = v12[1:len(v12)]
                                            }
                                        case 42:
                                            if len(v12) >= 2 {
                                                v12 = v12[1:len(v12)]
                                            }
                                        default:
                                            switch 123 {
                                            case 127:
                                                v29 -= v29
                                            case 376:
                                                var v34 string = "a"
                                                _ = v34
                                            default:
                                                switch "a" {
                                                case "a":
                                                    _ = len(v12)
                                                default:
                                                    v4 += (v4 << ((v4) & 63))
                                                }
                                            }
                                            if len(v12) > 0 {
                                                v12[0] = 9223372036854775807
                                            }
                                            v35 := GenS2{F0: 2, F1: "a", F2: genFn0(-2147483648, 127, -9223372036854775808)}
                                            _ = v35
                                        }
                                    }
                                case "hello":
                                    var v36 uint8 = 1
                                    _ = v36
                                default:
                                    if len(v12) > 0 {
                                        v12[0] = fi31
                                    }
                                }
                                v37 := GenS4{F0: ("\n\t" + "a")}
                                _ = v37
                            } else {
                                if v11 {
                                    v3 += 1
                                } else {
                                    fmt.Println(v4)
                                    _ = len(v12)
                                    if v11 {
                                        switch ("\n\t" + "") {
                                        case "hello":
                                            _ = len(v12)
                                        case "":
                                            v38 := GenS5{F0: 1, F1: v11, F2: v29}
                                            _ = v38
                                        }
                                        v4 = (v4 - 961)
                                    } else {
                                        switch fi31 {
                                        case -332:
                                            switch fi32 {
                                            case -293:
                                                fmt.Println(v3)
                                            default:
                                                if ("" >= ("" + "a")) {
                                                    v1 -= fi31
                                                }
                                            }
                                        case 32767:
                                            v12 = append(v12, fi31)
                                        }
                                        var v39 uint64 = 962
                                        _ = len(v12)
                                        _ = v39
                                    }
                                }
                                var v40 string = "foo bar"
                                _ = v40
                            }
                            v12 = append(v12, 1)
                        } else {
                            v41 := GenS6{F0: 32767, F1: 32767, F2: 1}
                            _ = len(v12)
                            var v42 uint32 = 197
                            _ = v41
                            _ = v42
                        }
                        v11 = v11
                        _ = fi32
                    }
                default:
                    v4 = (v4 - (v4 ^ v4))
                }
                _ = fi31
            }
            var v43 uint16 = 1
            v3 = (-(v3))
            _ = v43
        } else {
            v3 = v3
            for fi44 := 0; fi44 < 17; fi44++ {
                if len(v12) > 0 {
                    v12[0] = (v1 ^ (v1 ^ (fi44 + -1)))
                }
                _ = fi44
            }
            var v45 uint16 = 255
            _ = v45
        }
        var v46 uint = (v4 | (v4 % (((v4 + v4)) | 1)))
        _ = v46
    }
    _ = len(v12)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
