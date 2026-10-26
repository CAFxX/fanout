package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 uint8; F1 uint8 }
func genFn0(p0 uint8) string {
    switch 719 {
    case -2147483648:
        var v4 int8 = 42
        _ = v4
    }
    if false {
        var v5 int64 = 1
        var v6 uint16 = 32767
        _ = v5
        _ = v6
    } else {
        v7 := GenS1{F0: ("foo bar" + (("" + "") + (("" + "a") + "\n\t")))}
        var v8 uint64 = 354
        var v9 uint8 = p0
        _ = v7
        _ = v8
        _ = v9
    }
    var v10 uint64 = 4294967295
    p0 += 2
    if false {
        if true {
            v11 := []int{4294967295}
            if (4294967295 != 2147483647) {
                if len(v11) > 0 {
                    v11[0] = -545
                }
            }
            if len(v11) > 0 {
                v11[0] = -9223372036854775808
            }
            _ = v11
        }
        v12 := GenS1{F0: ("" + ("" + "foo bar"))}
        var v13 uint64 = v10
        _ = v12
        _ = v13
    }
    _ = p0
    _ = v10
    return "\n\t"
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{1, 445, 2147483648}
    v3 := make(map[int]int)
    v14 := GenS1{F0: ("a" + "\n\t")}
    var v15 int32 = 752
    for fi16 := 0; fi16 < 15; fi16++ {
        v15 = 1
        if (3.14159 >= 0.0) {
            v1[2] = 2147483648
            var v17 bool = (65535 <= 255)
            _ = v17
        }
        v15 = v15
        _ = fi16
    }
    v15 = v15
    v18 := GenS1{F0: "a"}
    switch 32768 {
    case 127:
        for fi19 := 0; fi19 < 2; fi19++ {
            for fi20 := 0; fi20 < 15; fi20++ {
                switch fi20 {
                case 2147483647:
                    if (!true) {
                        v21 := v3[42]
                        fmt.Println(fi20, fi19, v21)
                        _ = v21
                    } else {
                        if false {
                            if (false && (2147483648 < 2147483647)) {
                                if true {
                                    var v22 int64 = 998
                                    if true {
                                        v15 = (v15 * (v15 / ((982) | 1)))
                                        v23 := GenS2{F0: -746}
                                        _ = v23
                                    } else {
                                        v24 := v1[2147483647]
                                        _ = v24
                                    }
                                    var v25 uint32 = 2
                                    _ = v22
                                    _ = v25
                                }
                                v26 := GenS3{F0: ((!(124 < 100)) && (v15 >= (v15 | v15))), F1: 25}
                                v2 = append(v2, 65535)
                                _ = v26
                            } else {
                                v2 = append(v2, 752)
                            }
                            if true {
                                v15 += 2147483647
                                switch "foo bar" {
                                case "foo bar":
                                    switch "\n\t" {
                                    case "\n\t":
                                        _ = len(v2)
                                    case "":
                                        v15 = (v15 / (((v15 + v15)) | 1))
                                    }
                                case "\n\t":
                                    delete(v3, (-(fi20)))
                                case "hello":
                                    v27 := GenS3{F0: (!((v15 - (v15 ^ 1)) == 42)), F1: -128}
                                    _ = v27
                                default:
                                    if true {
                                        v15 = (v15 & v15)
                                        switch fi19 {
                                        case 127:
                                            var v28 uint16 = 616
                                            _ = v28
                                        case 0:
                                            v29 := v1[4294967295]
                                            _ = v29
                                        default:
                                            v15 = v15
                                        }
                                        if false {
                                            fmt.Println(fi20)
                                        } else {
                                            var v30 int16 = 1
                                            v15 = v15
                                            fmt.Println(v30)
                                            _ = v30
                                        }
                                    }
                                }
                            } else {
                                var v31 int64 = -709
                                if (2147483647 == 275) {
                                    v32 := GenS3{F0: false, F1: 100}
                                    var v33 int64 = 42
                                    v34 := GenS1{F0: ("hello" + "hello")}
                                    _ = v32
                                    _ = v33
                                    _ = v34
                                } else {
                                    var v35 uint = 42
                                    if (true && true) {
                                        var v36 float32 = 3.14159
                                        _ = v36
                                    }
                                    _ = v35
                                }
                                switch (fi19 % (((fi19 >> ((4294967295) & 63))) | 1)) {
                                case -2147483648:
                                    fmt.Println(fi20, fi19, v15)
                                default:
                                    var v37 int64 = v31
                                    _ = v37
                                }
                                _ = v31
                            }
                            switch "hello" {
                            case "":
                                var v38 float64 = 3.14159
                                _ = v38
                            }
                        }
                    }
                    v39 := GenS4{F0: 857, F1: 19}
                    v40 := GenS5{F0: 42, F1: 42}
                    var v41 string = ("\n\t" + "foo bar")
                    _ = v39
                    _ = v40
                    _ = v41
                case 32767:
                    var v42 bool = false
                    _ = v42
                case 32768:
                    if false {
                        if true {
                            delete(v3, fi20)
                        }
                        var v43 string = ("\n\t" + "hello")
                        _ = v43
                    } else {
                        var v44 string = (("" + (((("hello" + "\n\t") + "a") + "") + ("foo bar" + "foo bar"))) + "a")
                        fmt.Println(fi20)
                        if true {
                            var v45 int64 = -1
                            var v46 int32 = (v15 * (-(v15)))
                            _ = v45
                            _ = v46
                        }
                        _ = v44
                    }
                default:
                    switch "\n\t" {
                    case "a":
                        if ((-2.25 + 1.5) <= (0.0 - (1e10 * float64(0.1)))) {
                            var v47 uint64 = 591
                            if false {
                                var v48 string = ("foo bar" + (("" + "foo bar") + "foo bar"))
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                switch v48 {
                                case "foo bar":
                                    var v49 bool = true
                                    switch v48 {
                                    case "\n\t":
                                        switch "hello" {
                                        case "\n\t":
                                            var v50 string = v48
                                            _ = v50
                                        case "a":
                                            v47 -= v47
                                        case "hello":
                                            delete(v1, 1)
                                        default:
                                            if len(v2) > 0 {
                                                v2[0] = 676
                                            }
                                        }
                                    }
                                    fmt.Println(v49, v47)
                                    _ = v49
                                case "":
                                    var v51 int64 = -917
                                    _ = v51
                                }
                                _ = v48
                            }
                            _ = v47
                        }
                    default:
                        var v52 int = (fi20 >> ((65535) & 63))
                        _ = v52
                    }
                }
                _ = fi20
            }
            var v53 uint64 = 167
            v53 = 0
            _ = fi19
            _ = v53
        }
    case 32767:
        delete(v1, 2147483647)
    default:
        switch ("hello" + ("hello" + "a")) {
        case "foo bar":
            fmt.Println(v15)
        case "\n\t":
            for fi54 := 0; fi54 < 6; fi54++ {
                var v55 int8 = -128
                _ = fi54
                _ = v55
            }
        default:
            for fi56 := 0; fi56 < 15; fi56++ {
                fmt.Println(fi56, v15)
                _ = fi56
            }
        }
    }
    switch 0 {
    case 0:
        if len(v2) > 0 {
            v2[0] = 42
        }
    }
    if (42 <= 18446744073709551615) {
        fmt.Println(v15)
        v57 := GenS4{F0: 655, F1: 0}
        _ = v57
    }
    var v58 float32 = -2.25
    var v59 bool = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v58
    _ = v59
}
