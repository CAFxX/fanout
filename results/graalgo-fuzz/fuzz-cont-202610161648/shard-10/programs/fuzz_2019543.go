package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8; F1 int8; F2 int8 }
type GenS5 struct { F0 int8; F1 int8; F2 uint16 }
type GenS6 struct { F0 string }
func genFn0() int32 {
    v7 := GenS1{F0: 1000, F1: 609}
    v8 := GenS2{F0: 988}
    var v9 float64 = ((0.0 + 3.14159) * float64(-2.25))
    _ = v7
    _ = v8
    _ = v9
    return -1
}

func main() {
    v1 := make(map[int]string)
    v2 := []int{}
    v3 := []string{(("" + ("hello" + "\n\t")) + ""), ("a" + "foo bar"), ("a" + "")}
    v4 := make(map[uint32]bool)
    var v5 uint32 = 4294967295
    v6 := []uint32{v5, (v5 - v5), 78, v5}
    v5 = v5
    v10 := GenS2{F0: 2}
    var v11 float32 = 1e10
    switch -2147483648 {
    case 2147483647:
        if false {
            fmt.Println(v11)
            for fi12 := 0; fi12 < 18; fi12++ {
                var v13 uint16 = 555
                _ = fi12
                _ = v13
            }
        }
    case 9223372036854775807:
        v14 := GenS2{F0: 279}
        _ = v14
    default:
        v5 -= v5
    }
    if (("\n\t" + "a") < (("\n\t" + "") + ("a" + (("foo bar" + "foo bar") + "\n\t")))) {
        if (true && false) {
            v3 = append(v3, ("hello" + ""))
        }
        if true {
            for fi15 := 0; fi15 < 6; fi15++ {
                if true {
                    for fi16 := 0; fi16 < 17; fi16++ {
                        _ = len(v2)
                        switch (fi16 >> ((1) & 63)) {
                        case -397:
                            if (!(!true)) {
                                v17 := GenS2{F0: 0}
                                _ = v17
                            } else {
                                v5 -= (v5 - v5)
                                switch "\n\t" {
                                case "\n\t":
                                    var v18 string = ("\n\t" + "")
                                    _ = v18
                                case "foo bar":
                                    if true {
                                        fmt.Println(fi16, fi16)
                                        var v19 uint64 = 1
                                        if false {
                                            fmt.Println(v5, fi15)
                                            v20 := v4[v5]
                                            _ = v20
                                        } else {
                                            if false {
                                                var v21 int64 = 42
                                                v5 = (v5 | (v5 % (((v5 >> ((2) & 63))) | 1)))
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                _ = v21
                                            }
                                            switch "foo bar" {
                                            case "foo bar":
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            case "":
                                                switch fi16 {
                                                case 42:
                                                    v5 -= 1
                                                case -194:
                                                    if len(v2) > 0 {
                                                        v2[0] = int(v5)
                                                    }
                                                }
                                            }
                                            if ((false || false) && false) {
                                                delete(v1, (^(fi15)))
                                                v11 = ((v11 + float32(3.14159)) - (((v11 - v11) / (float32(float64(v11)) * v11)) / (v11 - 3.14159)))
                                                v22 := GenS3{F0: (!(!true))}
                                                v6 = append(v6, v5)
                                                var v23 int16 = 1
                                                _ = v22
                                                _ = v23
                                            }
                                        }
                                        _ = v19
                                    } else {
                                        v24 := GenS2{F0: 212}
                                        _ = v24
                                    }
                                }
                            }
                        case -2147483648:
                            switch "" {
                            case "":
                                fmt.Println(v11, fi16, fi16)
                            case "\n\t":
                                v11 -= float32(((1e10 / 3.14159) / 3.14159))
                            }
                        default:
                            if (genFn0() >= -434) {
                                if len(v2) > 0 {
                                    v2[0] = (fi15 * (fi16 | fi16))
                                }
                                switch "a" {
                                case "\n\t":
                                    var v25 uint16 = 922
                                    _ = v25
                                case "foo bar":
                                    v11 -= (1.5 + (v11 + v11))
                                }
                                var v26 int8 = 0
                                _ = v26
                            }
                        }
                        var v27 uint = 171
                        _ = fi16
                        _ = v27
                    }
                } else {
                    fmt.Println(fi15, fi15)
                    for fi28 := 0; fi28 < 20; fi28++ {
                        var v29 int16 = 1000
                        _ = fi28
                        _ = v29
                    }
                    fmt.Println(v11, v5, v11)
                }
                _ = fi15
            }
            for fi30 := 0; fi30 < 10; fi30++ {
                if (!(((fi30 | (fi30 << ((4294967295) & 63))) != 32767) || (true || (-328 < int16(fi30))))) {
                    if len(v6) > 0 {
                        v6[0] = 4294967295
                    }
                }
                v5 = 4294967295
                v5 = v5
                _ = fi30
            }
        }
        switch -1 {
        case 32767:
            if true {
                v5 = 4294967295
                var v31 int16 = 57
                _ = v31
            }
        default:
            for fi32 := 0; fi32 < 14; fi32++ {
                fmt.Println(v11)
                v33 := GenS2{F0: 952}
                _ = fi32
                _ = v33
            }
        }
    } else {
        var v34 uint16 = 551
        v11 += ((v11 + 1.5) / v11)
        _ = v34
    }
    fmt.Println(v5, v11)
    v11 -= 3.14159
    if (!(2 < 184)) {
        for fi35 := 0; fi35 < 2; fi35++ {
            for fi36 := 0; fi36 < 11; fi36++ {
                var v37 uint64 = 18446744073709551615
                if (42 < 0) {
                    v11 = v11
                    v11 = (float32(1.5) - v11)
                } else {
                    if (fi36 <= (fi35 * fi36)) {
                        v2 = append(v2, fi36)
                        v4[v5] = true
                        v11 = (-2.25 / (v11 - v11))
                    } else {
                        switch (fi35 % ((1) | 1)) {
                        case 1:
                            var v38 float32 = 0.0
                            _ = v38
                        case 42:
                            _ = len(v2)
                        default:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        v37 = (^(v37))
                        delete(v1, 1)
                    }
                    if true {
                        if (!true) {
                            fmt.Println(v37, v11, fi35)
                            if (true || true) {
                                if false {
                                    var v39 int16 = 1000
                                    if len(v3) > 0 {
                                        v3[0] = ("foo bar" + ("foo bar" + ""))
                                    }
                                    fmt.Println(v5, v37, v39)
                                    _ = v39
                                }
                            } else {
                                v11 = (v11 - float32(1e10))
                                var v40 string = (("" + ("" + "hello")) + ("a" + ("foo bar" + ("" + ("" + "foo bar")))))
                                _ = v40
                            }
                            fmt.Println(v5)
                        }
                        v41 := GenS4{F0: 42, F1: 100, F2: -128}
                        fmt.Println(v37, fi36)
                        if false {
                            fmt.Println(v11, v5)
                            v42 := GenS5{F0: -1, F1: 17, F2: 1}
                            _ = v42
                        }
                        _ = v41
                    } else {
                        if true {
                            var v43 uint16 = 541
                            var v44 uint16 = 42
                            _ = v43
                            _ = v44
                        } else {
                            if (true && ((1 <= 0) && false)) {
                                var v45 int8 = 127
                                v37 = v37
                                var v46 float64 = 3.14159
                                v11 = v11
                                _ = v45
                                _ = v46
                            }
                            if false {
                                v37 = v37
                            }
                        }
                    }
                    if false {
                        switch 32767 {
                        case 1:
                            var v47 float64 = -2.25
                            _ = v47
                        case -2147483648:
                            v5 = (v5 << ((42) & 63))
                            if true {
                                if (0 >= -32768) {
                                    var v48 int8 = 80
                                    fmt.Println(fi35)
                                    if (!("foo bar" >= "hello")) {
                                        var v49 int8 = -1
                                        fmt.Println(v37, v5, v48)
                                        _ = v49
                                    }
                                    _ = v48
                                }
                            }
                        }
                        switch fi35 {
                        case 32767:
                            switch fi36 {
                            case -388:
                                switch ("" + "hello") {
                                case "\n\t":
                                    if (442 < 255) {
                                        if ((("a" < "") && true) && false) {
                                            if (v5 < (v5 & 805)) {
                                                v37 -= 18446744073709551615
                                            } else {
                                                if (!(false && ("" >= ("foo bar" + ((("foo bar" + "") + ("hello" + "a")) + "a"))))) {
                                                    v3 = append(v3, "hello")
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                } else {
                                                    var v50 uint8 = 207
                                                    v11 += float32((-2.25 * ((3.14159 * 0.1) * 1e10)))
                                                    _ = v50
                                                }
                                                var v51 int = int(v5)
                                                _ = v51
                                            }
                                            _ = len(v4)
                                        }
                                    }
                                case "foo bar":
                                    fmt.Println(fi36)
                                default:
                                    fmt.Println(fi36, fi35)
                                }
                            case 0:
                                if (-1 <= -1) {
                                    var v52 int32 = 263
                                    _ = v52
                                } else {
                                    v5 = v5
                                    _ = len(v3)
                                }
                            }
                        case 635:
                            v53 := GenS1{F0: -269, F1: -456}
                            _ = v53
                        case -9223372036854775808:
                            v11 += ((v11 * v11) * v11)
                        }
                    } else {
                        v54 := GenS6{F0: ("a" + "a")}
                        var v55 uint16 = 366
                        _ = v54
                        _ = v55
                    }
                }
                var v56 uint64 = 1
                _ = fi36
                _ = v37
                _ = v56
            }
            _ = len(v3)
            v57 := GenS2{F0: 32767}
            _ = fi35
            _ = v57
        }
        for fi58 := 0; fi58 < 2; fi58++ {
            v5 -= 2147483648
            v59 := GenS6{F0: "a"}
            _ = fi58
            _ = v59
        }
        if true {
            v11 -= v11
            v5 += 2147483647
        }
    }
    var v60 int32 = -2147483648
    for fi61 := 0; fi61 < 3; fi61++ {
        if false {
            _ = len(v1)
            delete(v4, 555)
        } else {
            v5 -= (v5 - (v5 & v5))
        }
        _ = fi61
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
}
