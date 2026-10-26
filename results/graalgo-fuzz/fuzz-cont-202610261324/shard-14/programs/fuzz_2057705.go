package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 string; F1 int32; F2 uint64 }
func genFn0() int16 {
    if (true && (!false)) {
        if (!(-851 >= 1)) {
            var v4 bool = false
            if true {
                var v5 int32 = 1
                v4 = (((("\n\t" + "\n\t") + "hello") + "a") == "\n\t")
                fmt.Println(v4, v5)
                _ = v5
            }
            _ = v4
        } else {
            switch -806 {
            case 2147483647:
                var v6 int = 32768
                _ = v6
            }
        }
    } else {
        var v7 string = ("a" + "hello")
        _ = v7
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        if (2147483647 >= 1) {
            var v9 uint8 = 179
            _ = v9
        }
        if false {
            switch -10 {
            case 2147483647:
                for fi10 := 0; fi10 < 19; fi10++ {
                    fmt.Println(fi10, fi8)
                    _ = fi10
                }
            case -387:
                switch (fi8 * (fi8 & (^(fi8)))) {
                case -856:
                    v11 := []int{(fi8 ^ fi8), (fi8 - (fi8 * fi8)), (fi8 >> ((40) & 63)), (fi8 ^ (fi8 + (fi8 - fi8)))}
                    _ = v11
                }
            case -75:
                var v12 int = -168
                _ = v12
            default:
                var v13 float64 = 3.14159
                _ = v13
            }
        }
        for fi14 := 0; fi14 < 16; fi14++ {
            var v15 uint16 = 258
            v16 := GenS1{F0: (v15 & (v15 | (v15 / ((v15) | 1)))), F1: 255}
            v15 = (v15 % ((v15) | 1))
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi8
    }
    if (-1 < 42) {
        if (false && (1 > 2)) {
            v17 := GenS2{F0: 32767}
            _ = v17
        }
        var v18 float32 = 1e10
        _ = v18
    }
    var v19 uint16 = 1
    var v20 int32 = -835
    _ = v19
    _ = v20
    return 441
}
func genFn1(p0 uint16, p1 bool) uint16 {
    var v21 int16 = 909
    v21 -= 1
    v21 += v21
    _ = p0
    _ = p1
    _ = v21
    return p0
}

func main() {
    v1 := []string{"foo bar", "\n\t", "hello"}
    v2 := []string{(((((("a" + "hello") + "foo bar") + "a") + ((("foo bar" + "hello") + ("foo bar" + "a")) + "")) + ("hello" + "hello")) + "a"), ""}
    v3 := []string{}
    for fi22 := 0; fi22 < 3; fi22++ {
        fmt.Println(fi22)
        _ = fi22
    }
    var v23 int32 = 0
    v23 = v23
    if true {
        for fi24 := 0; fi24 < 16; fi24++ {
            fmt.Println(v23, v23)
            _ = fi24
        }
        v23 = (^(v23))
    }
    if ((42 != -1) || false) {
        var v25 uint32 = 4294967295
        v1 = append(v1, ("a" + ("hello" + "foo bar")))
        if ((!((!(-160 < -9223372036854775808)) && (328 != 127))) && (1.5 != 0.0)) {
            var v26 bool = true
            if v26 {
                v23 += v23
                v26 = v26
            }
            fmt.Println(v26, v23, v23)
            _ = v26
        } else {
            v23 -= v23
            var v27 uint16 = 32767
            v2 = append(v2, ("hello" + "hello"))
            _ = v27
        }
        _ = v25
    }
    var v28 uint = 349
    for fi29 := 0; fi29 < 12; fi29++ {
        v30 := GenS1{F0: 1, F1: 65535}
        for fi31 := 0; fi31 < 6; fi31++ {
            if len(v1) > 0 {
                v1[0] = ("" + (("\n\t" + ("hello" + ("foo bar" + (("hello" + ("\n\t" + "hello")) + "")))) + "a"))
            }
            switch "foo bar" {
            case "a":
                if true {
                    v32 := GenS1{F0: 455, F1: 42}
                    var v33 uint = v28
                    _ = v32
                    _ = v33
                } else {
                    var v34 uint32 = 42
                    var v35 int = fi31
                    if (true || (42 <= -824)) {
                        v23 = v23
                        var v36 float64 = 0.1
                        v37 := GenS3{F0: (fi31 % (((-(fi29))) | 1))}
                        _ = v36
                        _ = v37
                    }
                    _ = v34
                    _ = v35
                }
            }
            fmt.Println(v28)
            switch ("hello" + "foo bar") {
            case "a":
                if false {
                    if (!(!((fi29 % (((fi29 & fi29)) | 1)) <= fi31))) {
                        v38 := GenS1{F0: 255, F1: 42}
                        fmt.Println(fi29)
                        _ = v38
                    }
                } else {
                    v39 := GenS2{F0: 1}
                    if true {
                        if (0 > 13) {
                            v28 = (v28 & (v28 * 42))
                        }
                        switch "a" {
                        case "hello":
                            v40 := GenS4{F0: 4294967295}
                            _ = v40
                        }
                    } else {
                        switch (fi29 >> ((v28) & 63)) {
                        case 312:
                            if false {
                                fmt.Println(v23, v23)
                                _ = len(v2)
                                v23 = v23
                            }
                        default:
                            v23 = v23
                        }
                    }
                    _ = v39
                }
            case "hello":
                var v41 uint8 = 128
                _ = v41
            }
            _ = fi31
        }
        _ = fi29
        _ = v30
    }
    if (((0.1 - 1.5) >= -2.25) && (genFn1(252, (!((!true) && true))) <= 1)) {
        if true {
            switch 9223372036854775807 {
            case 32767:
                v42 := GenS4{F0: 42}
                _ = v42
            case -9223372036854775808:
                switch 127 {
                case -807:
                    v23 += (v23 + v23)
                case 234:
                    if false {
                        _ = len(v2)
                        var v43 bool = (-296 <= -32768)
                        _ = v43
                    } else {
                        for fi44 := 0; fi44 < 13; fi44++ {
                            v28 -= (v28 << (((v28 ^ v28)) & 63))
                            v45 := GenS5{F0: 100}
                            _ = fi44
                            _ = v45
                        }
                    }
                default:
                    fmt.Println(v23)
                }
            case -2147483648:
                var v46 float64 = -2.25
                _ = v46
            default:
                v23 = v23
            }
        } else {
            v28 = v28
            var v47 int64 = 42
            _ = len(v3)
            _ = v47
        }
        if true {
            if len(v2) > 0 {
                v2[0] = "hello"
            }
        } else {
            var v48 int16 = 214
            if (false && (true && false)) {
                var v49 uint8 = 1
                v28 -= (v28 / ((1) | 1))
                _ = v49
            } else {
                if ((false || false) && false) {
                    for fi50 := 0; fi50 < 8; fi50++ {
                        var v51 int64 = 9223372036854775807
                        for fi52 := 0; fi52 < 20; fi52++ {
                            if (!true) {
                                v23 = (v23 * 982)
                            }
                            v53 := GenS6{F0: "foo bar", F1: v23, F2: 2}
                            if ((1 == 127) && true) {
                                var v54 float64 = ((float64(3.14159) / (1.5 * ((((3.14159 / 1.0) * 3.14159) / 1.5) + (3.14159 + 1e10)))) + 1.5)
                                v23 += v23
                                var v55 uint = 18446744073709551615
                                switch "foo bar" {
                                case "a":
                                    v1 = append(v1, (((((("\n\t" + ("a" + "\n\t")) + "") + ("a" + "foo bar")) + (("" + "hello") + ("foo bar" + "foo bar"))) + "a") + ("hello" + ("foo bar" + ("\n\t" + ("" + ""))))))
                                }
                                _ = len(v3)
                                _ = v54
                                _ = v55
                            } else {
                                switch fi50 {
                                case 338:
                                    var v56 float32 = 1.5
                                    _ = v56
                                case 2147483647:
                                    v23 += (v23 & (v23 & (v23 << ((v28) & 63))))
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v48 = 1
                                v51 = 846
                            }
                            _ = fi52
                            _ = v53
                        }
                        _ = fi50
                        _ = v51
                    }
                    var v57 uint64 = 2
                    _ = v57
                }
            }
            v23 = (v23 * v23)
            _ = v48
        }
        var v58 float32 = float32(((0.0 - 1.5) * 0.1))
        _ = v58
    } else {
        var v59 uint = (v28 << ((1) & 63))
        _ = v59
    }
    _ = len(v3)
    if (248 < 9223372036854775807) {
        fmt.Println(v23)
        if true {
            _ = len(v1)
            _ = len(v1)
        }
    } else {
        v23 = v23
        v60 := GenS1{F0: 65535, F1: 42}
        _ = v60
    }
    var v61 int = 4294967295
    var v62 float32 = (0.1 - 1e10)
    fmt.Println(v23)
    fmt.Println(v28, v62, v61)
    var v63 uint16 = 32767
    switch (v61 | -944) {
    case -2147483648:
        var v64 string = ""
        _ = v64
    case -120:
        var v65 int8 = -1
        _ = v65
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v23
    _ = v28
    _ = v3
    _ = v61
    _ = v62
    _ = v63
}
