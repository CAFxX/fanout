package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 uint32; F1 int32; F2 int32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8; F1 uint8; F2 uint16 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint32; F1 int64 }
type GenS8 struct { F0 uint }
func genFn0(p0 int32, p1 string) uint8 {
    v6 := GenS1{F0: 0, F1: 56}
    v7 := GenS2{F0: true, F1: 144}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 2
}

func main() {
    v1 := make(map[uint32]string)
    var v2 bool = false
    v3 := []string{}
    v4 := make(map[string]int)
    v5 := make(map[uint32]bool)
    v2 = false
    var v8 uint64 = 341
    if len(v3) > 0 {
        v3[0] = ("hello" + (("foo bar" + "\n\t") + "a"))
    }
    switch "foo bar" {
    case "foo bar":
        v2 = v2
    default:
        _ = len(v3)
    }
    switch "a" {
    case "\n\t":
        var v9 bool = ((!false) || v2)
        _ = v9
    case "":
        v10 := GenS3{F0: 4294967295, F1: 2147483647, F2: -1}
        _ = v10
    default:
        if (((1 >= 116) || v2) && true) {
            v2 = (v2 && (174 != 220))
        } else {
            v2 = v2
            v11 := GenS1{F0: -9223372036854775808, F1: 243}
            _ = v11
        }
    }
    switch "\n\t" {
    case "\n\t":
        switch "hello" {
        case "hello":
            switch ("foo bar" + "\n\t") {
            case "a":
                var v12 uint16 = 32767
                _ = v12
            case "hello":
                var v13 float32 = 3.14159
                _ = v13
            }
        default:
            if (false && v2) {
                fmt.Println(v8, v2)
                var v14 float64 = (((-2.25 / (-2.25 / 1e10)) / float64(0.1)) * (1e10 * ((1e10 * 0.0) / 1.0)))
                v8 -= v8
                _ = v14
            }
        }
    case "foo bar":
        v15 := v4["foo bar"]
        _ = v15
    default:
        v16 := v5[4294967295]
        _ = v16
    }
    switch (((("\n\t" + "hello") + ("" + "foo bar")) + "\n\t") + "\n\t") {
    case "hello":
        for fi17 := 0; fi17 < 20; fi17++ {
            v3 = append(v3, "hello")
            if (!v2) {
                var v18 int32 = 1
                for fi19 := 0; fi19 < 10; fi19++ {
                    if (292 < 1) {
                        v20 := GenS4{F0: 0}
                        if (!true) {
                            v21 := GenS3{F0: 1, F1: v18, F2: (v18 * 841)}
                            _ = v21
                        }
                        if v2 {
                            v22 := GenS5{F0: 0, F1: 128, F2: 42}
                            fmt.Println(fi17, fi19, v8)
                            v8 = v8
                            _ = v22
                        } else {
                            switch (("\n\t" + "hello") + ("hello" + "hello")) {
                            case "hello":
                                v8 += 2
                            case "a":
                                var v23 float32 = ((((0.1 - 0.1) / 1e10) - (1e10 + ((-2.25 - 3.14159) - (1e10 + (1.5 - (float32((1e10 - (1.5 * 1e10))) + ((1e10 + float32(3.14159)) + -2.25))))))) * (0.1 * -2.25))
                                _ = v23
                            case "\n\t":
                                fmt.Println(v2, v18)
                            }
                            fmt.Println(v2, v2)
                            switch 9223372036854775807 {
                            case -9223372036854775808:
                                switch "" {
                                case "hello":
                                    v24 := GenS4{F0: 18446744073709551615}
                                    _ = v24
                                }
                            case 32768:
                                if v2 {
                                    if len(v3) > 0 {
                                        v3[0] = "a"
                                    }
                                    v8 += (v8 & 541)
                                    v4["hello"] = fi19
                                    var v25 int = 32767
                                    _ = v25
                                } else {
                                    var v26 uint = 660
                                    v27 := GenS2{F0: true, F1: 128}
                                    v28 := GenS6{F0: uint16(v26)}
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                }
                            default:
                                if v2 {
                                    var v29 int32 = v18
                                    _ = v29
                                }
                            }
                        }
                        _ = v20
                    } else {
                        v2 = (false && v2)
                        var v30 uint64 = (v8 * v8)
                        var v31 uint8 = 1
                        switch (fi17 | fi19) {
                        case 919:
                            switch ("foo bar" + "hello") {
                            case "":
                                var v32 int64 = 9223372036854775807
                                _ = v32
                            case "\n\t":
                                v30 -= uint64(v31)
                            }
                        case 32767:
                            v8 = (v30 | (v30 - 18446744073709551615))
                        default:
                            fmt.Println(v30, v2, fi17)
                        }
                        _ = v30
                        _ = v31
                    }
                    v33 := GenS4{F0: 4294967295}
                    v8 += v8
                    v34 := GenS7{F0: 2147483647, F1: -298}
                    var v35 int8 = 1
                    _ = fi19
                    _ = v33
                    _ = v34
                    _ = v35
                }
                v36 := GenS3{F0: 2147483647, F1: (-(v18)), F2: int32(v8)}
                _ = v18
                _ = v36
            } else {
                v37 := GenS1{F0: 0, F1: 127}
                _ = v37
            }
            v38 := v5[350]
            _ = fi17
            _ = v38
        }
    case "a":
        switch 4294967295 {
        case 32768:
            switch 4294967295 {
            case 4294967295:
                var v39 int16 = 679
                _ = v39
            default:
                if v2 {
                    fmt.Println(v2, v2)
                } else {
                    for fi40 := 0; fi40 < 10; fi40++ {
                        var v41 uint32 = 157
                        var v42 int = (-(fi40))
                        _ = fi40
                        _ = v41
                        _ = v42
                    }
                    var v43 float64 = 0.0
                    _ = v43
                }
            }
        case -1:
            var v44 int32 = 0
            _ = v44
        }
    default:
        if v2 {
            if (1 != 42) {
                var v45 int64 = -129
                if (1e10 == -2.25) {
                    if len(v3) > 0 {
                        v3[0] = ""
                    }
                    var v46 bool = (100 <= 0)
                    _ = v46
                }
                _ = v45
            } else {
                var v47 bool = false
                var v48 uint = 731
                v49 := GenS5{F0: 189, F1: genFn0(-539, ""), F2: 255}
                _ = v47
                _ = v48
                _ = v49
            }
            var v50 float64 = ((1.5 + float64(0.1)) / 1.0)
            _ = v50
        } else {
            _ = len(v4)
            if (42 != 2147483647) {
                var v51 uint = 652
                _ = v51
            }
            var v52 uint8 = 2
            _ = v52
        }
    }
    switch 127 {
    case -2147483648:
        v8 -= (v8 % (((v8 % (((v8 ^ v8)) | 1))) | 1))
    default:
        _ = len(v5)
    }
    var v53 bool = false
    var v54 int64 = 42
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v55 uint = 65535
    if true {
        for fi56 := 0; fi56 < 20; fi56++ {
            v57 := GenS5{F0: 127, F1: 42, F2: 2}
            if len(v3) > 0 {
                v3[0] = ("\n\t" + "")
            }
            v53 = v53
            _ = fi56
            _ = v57
        }
        v2 = false
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if ((373 == -9223372036854775808) || v53) {
            var v58 float64 = 3.14159
            switch "foo bar" {
            case "\n\t":
                v54 += (v54 << (((v55 >> ((v55) & 63))) & 63))
            default:
                v55 -= (v55 | (v55 + 4294967295))
            }
            if false {
                v59 := GenS8{F0: (v55 % ((0) | 1))}
                fmt.Println(v55, v54)
                _ = v59
            }
            _ = v58
        }
    }
    var v60 uint64 = v8
    fmt.Println(v54)
    var v61 uint8 = 1
    v61 += 119
    var v62 int = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v55
    _ = v60
    _ = v61
    _ = v62
    _ = v8
}
