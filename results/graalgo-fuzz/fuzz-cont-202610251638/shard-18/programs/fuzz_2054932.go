package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 int8; F1 uint64; F2 uint64 }
type GenS5 struct { F0 uint; F1 uint16 }
type GenS6 struct { F0 uint8; F1 uint16 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int; F1 uint64 }
type GenS9 struct { F0 int32 }

func main() {
    v1 := []int{32768}
    v2 := []string{(("hello" + "hello") + "hello"), ""}
    var v3 int8 = -21
    var v4 int16 = -243
    var v5 int32 = 371
    if true {
        var v6 float32 = 1.5
        if len(v2) > 0 {
            v2[0] = ("foo bar" + "hello")
        }
        var v7 float32 = v6
        _ = v6
        _ = v7
    } else {
        _ = len(v1)
    }
    v5 = (v5 - -1)
    fmt.Println(v4, v4, v3)
    for fi8 := 0; fi8 < 18; fi8++ {
        for fi9 := 0; fi9 < 3; fi9++ {
            v4 += v4
            if (30 >= (v3 >> ((0) & 63))) {
                v3 = v3
                v4 = v4
            } else {
                v5 = v5
                var v10 uint8 = 128
                if true {
                    var v11 int8 = v3
                    switch ((("a" + "hello") + ("foo bar" + "a")) + "a") {
                    case "a":
                        fmt.Println(fi9, v5, v3)
                    case "\n\t":
                        switch "a" {
                        case "\n\t":
                            if false {
                                var v12 string = "\n\t"
                                _ = v12
                            } else {
                                v13 := GenS1{F0: (v10 + v10)}
                                if (-358 == 739) {
                                    v14 := GenS2{F0: "a"}
                                    fmt.Println(fi8, v11, fi9)
                                    var v15 int32 = (v5 % (((v5 << ((255) & 63))) | 1))
                                    _ = v14
                                    _ = v15
                                } else {
                                    v4 = v4
                                    v3 += -1
                                    v2 = append(v2, "foo bar")
                                }
                                _ = v13
                            }
                        default:
                            var v16 int8 = 0
                            _ = v16
                        }
                        if true {
                            switch "foo bar" {
                            case "hello":
                                v5 = v5
                            }
                        } else {
                            switch "\n\t" {
                            case "a":
                                var v17 uint64 = 226
                                _ = v17
                            default:
                                v3 += -128
                            }
                        }
                        if ("hello" <= ("\n\t" + (("foo bar" + ("\n\t" + "foo bar")) + ""))) {
                            v3 -= int8(v5)
                            v11 = (v3 << ((255) & 63))
                            v10 -= (v10 ^ v10)
                            if len(v2) > 0 {
                                v2[0] = "\n\t"
                            }
                            v10 -= v10
                        } else {
                            var v18 int8 = (v3 * v3)
                            v10 = v10
                            _ = v18
                        }
                        fmt.Println(v4, v5)
                    default:
                        fmt.Println(fi8)
                    }
                    if true {
                        v4 += (v4 & v4)
                        var v19 float32 = (1.5 - 0.1)
                        if true {
                            if false {
                                var v20 uint8 = 238
                                var v21 int = (fi9 >> ((18446744073709551615) & 63))
                                _ = v20
                                _ = v21
                            }
                            if (v10 >= (^(v10))) {
                                if ((v3 | (v3 * (v11 + (v3 >> ((42) & 63))))) > (v3 << ((213) & 63))) {
                                    if (!(-2.25 < float64(v19))) {
                                        var v22 int = fi9
                                        fmt.Println(v19)
                                        switch v22 {
                                        case 9223372036854775807:
                                            v19 = float32(-2.25)
                                        case -9223372036854775808:
                                            v4 = (v4 & (v4 << ((42) & 63)))
                                        case 4294967295:
                                            switch "a" {
                                            case "":
                                                fmt.Println(v19, fi8, v22)
                                            case "a":
                                                if false {
                                                    if (0 < v5) {
                                                        v11 = (v3 * -128)
                                                        v23 := GenS3{F0: (fi8 >> ((299) & 63)), F1: v11}
                                                        if false {
                                                            if ((("foo bar" + "\n\t") + "foo bar") != "foo bar") {
                                                                var v24 float32 = float32(3.14159)
                                                                if len(v1) > 0 {
                                                                    v1[0] = -1
                                                                }
                                                                _ = v24
                                                            } else {
                                                                if false {
                                                                    v25 := GenS4{F0: (v3 ^ v11), F1: 2, F2: 57}
                                                                    var v26 string = "hello"
                                                                    v27 := GenS1{F0: (v10 % ((106) | 1))}
                                                                    _ = v25
                                                                    _ = v26
                                                                    _ = v27
                                                                } else {
                                                                    if len(v2) > 0 {
                                                                        v2[0] = "a"
                                                                    }
                                                                    v22 += (v22 % (((fi8 & (fi8 / (((fi9 * -681)) | 1)))) | 1))
                                                                }
                                                            }
                                                            fmt.Println(v3, v4, v4)
                                                            if true {
                                                                fmt.Println(v19)
                                                                if ((v5 >= 815) || (fi8 < 42)) {
                                                                    var v28 uint8 = 1
                                                                    _ = v28
                                                                }
                                                                if false {
                                                                    switch fi8 {
                                                                    case 33:
                                                                        fmt.Println(v22)
                                                                        if true {
                                                                            if (!(false || false)) {
                                                                                v10 += (v10 ^ (v10 / (((v10 & 127)) | 1)))
                                                                                if false {
                                                                                    fmt.Println(fi8)
                                                                                }
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = (("\n\t" + ("\n\t" + (("hello" + "a") + "a"))) + "foo bar")
                                                                                }
                                                                            } else {
                                                                                var v29 float32 = v19
                                                                                var v30 uint16 = 32767
                                                                                _ = v29
                                                                                _ = v30
                                                                            }
                                                                        }
                                                                    }
                                                                    if (0 != 613) {
                                                                        v31 := GenS2{F0: ((("\n\t" + (("\n\t" + "") + ("" + "a"))) + "\n\t") + "\n\t")}
                                                                        var v32 uint16 = 13
                                                                        _ = v31
                                                                        _ = v32
                                                                    } else {
                                                                        var v33 uint64 = 18446744073709551615
                                                                        _ = len(v1)
                                                                        v19 = float32(0.0)
                                                                        fmt.Println(v19, v10, fi9)
                                                                        _ = v33
                                                                    }
                                                                } else {
                                                                    v34 := GenS1{F0: 93}
                                                                    _ = len(v1)
                                                                    v35 := GenS3{F0: (fi9 + fi9), F1: v3}
                                                                    v2 = append(v2, "\n\t")
                                                                    _ = v34
                                                                    _ = v35
                                                                }
                                                            }
                                                        } else {
                                                            _ = len(v2)
                                                            v36 := GenS5{F0: 4294967295, F1: 1}
                                                            var v37 int16 = int16(fi8)
                                                            v38 := GenS4{F0: v11, F1: 996, F2: 732}
                                                            _ = v36
                                                            _ = v37
                                                            _ = v38
                                                        }
                                                        _ = v23
                                                    } else {
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    }
                                                    var v39 int16 = 1000
                                                    _ = v39
                                                }
                                            default:
                                                v40 := GenS6{F0: 128, F1: 65535}
                                                _ = v40
                                            }
                                        default:
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                        _ = v22
                                    }
                                }
                                v41 := GenS4{F0: (v11 & v11), F1: 0, F2: 2}
                                var v42 uint16 = 992
                                _ = v41
                                _ = v42
                            }
                            var v43 uint64 = 2
                            _ = v43
                        }
                        _ = v19
                    }
                    switch "" {
                    case "a":
                        var v44 int = fi9
                        _ = v44
                    default:
                        switch "a" {
                        case "hello":
                            var v45 uint64 = 208
                            _ = v45
                        case "":
                            if true {
                                v46 := GenS4{F0: v11, F1: 4294967295, F2: 4294967295}
                                if (!true) {
                                    var v47 uint = uint(v10)
                                    fmt.Println(fi8)
                                    _ = len(v1)
                                    _ = v47
                                } else {
                                    if false {
                                        v5 -= (v5 | v5)
                                    } else {
                                        var v48 int8 = (v3 + (v3 * -128))
                                        fmt.Println(v48, v11, fi8)
                                        _ = v48
                                    }
                                }
                                v11 = 42
                                if true {
                                    v2 = append(v2, "\n\t")
                                }
                                _ = v46
                            }
                        case "foo bar":
                            if false {
                                v49 := GenS7{F0: true}
                                v3 = (v3 >> ((2) & 63))
                                _ = v49
                            }
                        }
                    }
                    v5 = (v5 * v5)
                    _ = v11
                }
                _ = v10
            }
            var v50 uint32 = 0
            _ = fi9
            _ = v50
        }
        var v51 float64 = ((3.14159 * float64(0.0)) - (0.1 / (-2.25 / -2.25)))
        v3 += 82
        _ = fi8
        _ = v51
    }
    for fi52 := 0; fi52 < 10; fi52++ {
        switch "foo bar" {
        case "a":
            if (!((!((0.1 <= 0.0) && (1 < 18446744073709551615))) && true)) {
                v53 := GenS6{F0: 0, F1: 0}
                if len(v2) > 0 {
                    v2[0] = "hello"
                }
                _ = v53
            } else {
                for fi54 := 0; fi54 < 10; fi54++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi54
                }
                var v55 int32 = 2147483647
                _ = v55
            }
        case "hello":
            for fi56 := 0; fi56 < 8; fi56++ {
                v57 := GenS8{F0: int(v4), F1: 1}
                _ = fi56
                _ = v57
            }
        case "foo bar":
            if true {
                v3 += int8(fi52)
                v2 = append(v2, ("foo bar" + "hello"))
                v5 = 0
            }
        default:
            for fi58 := 0; fi58 < 17; fi58++ {
                v3 -= (v3 & v3)
                var v59 bool = true
                _ = fi58
                _ = v59
            }
        }
        _ = fi52
    }
    v5 += 0
    v4 = (v4 & (v4 << ((2) & 63)))
    var v60 uint8 = 0
    if true {
        for fi61 := 0; fi61 < 13; fi61++ {
            v1 = append(v1, (fi61 >> ((18446744073709551615) & 63)))
            var v62 float32 = 3.14159
            var v63 int16 = v4
            _ = fi61
            _ = v62
            _ = v63
        }
        v4 = (v4 | (v4 * v4))
    }
    v5 = (v5 % (((v5 << ((740) & 63))) | 1))
    v3 = 0
    v64 := GenS9{F0: 42}
    var v65 int8 = 101
    var v66 bool = (!true)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v60:", v60)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v60
    _ = v64
    _ = v65
    _ = v66
}
