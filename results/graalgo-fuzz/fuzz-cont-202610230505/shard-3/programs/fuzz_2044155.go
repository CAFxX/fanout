package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 int32; F1 int16; F2 int }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int; F1 uint8; F2 int16 }
type GenS7 struct { F0 uint; F1 uint16 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int16 }
func genFn0(p0 uint16, p1 bool, p2 string) int {
    p0 -= (p0 ^ 1)
    p1 = ((((("foo bar" + "hello") + ("\n\t" + "a")) + p2) + p2) > p2)
    _ = p0
    _ = p1
    _ = p2
    return -2147483648
}
func genFn1(p0 bool, p1 uint32, p2 bool) int8 {
    if false {
        v5 := []int{-9223372036854775808}
        var v6 uint64 = 1
        if true {
            var v7 uint16 = 2
            v5 = append(v5, -9223372036854775808)
            _ = v7
        }
        _ = v5
        _ = v6
    }
    var v8 string = "a"
    v9 := GenS1{F0: 65535, F1: 636}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 42
}

func main() {
    v1 := []string{("foo bar" + (("\n\t" + "a") + "foo bar"))}
    var v2 int = -999
    var v3 uint32 = 558
    v4 := []string{"a", "\n\t", "hello", "hello"}
    for fi10 := 0; fi10 < 11; fi10++ {
        for fi11 := 0; fi11 < 19; fi11++ {
            if true {
                if ((!true) && false) {
                    v3 += (v3 << ((681) & 63))
                    var v12 string = ("" + "\n\t")
                    _ = v12
                } else {
                    if true {
                        fmt.Println(v2, v3, v3)
                    }
                    if false {
                        var v13 float64 = float64(0.0)
                        v13 -= (((float64((((0.0 * (0.1 / 3.14159)) / (3.14159 * -2.25)) * -2.25)) / 3.14159) / v13) - v13)
                        _ = v13
                    } else {
                        v1 = append(v1, "foo bar")
                        if false {
                            switch ("\n\t" + "foo bar") {
                            case "a":
                                v3 -= (v3 / ((112) | 1))
                            case "hello":
                                v2 -= 127
                            default:
                                if false {
                                    _ = len(v4)
                                } else {
                                    var v14 uint32 = 1
                                    _ = v14
                                }
                            }
                            switch ("foo bar" + "foo bar") {
                            case "foo bar":
                                v15 := GenS1{F0: 514, F1: 1}
                                _ = v15
                            case "":
                                v3 = v3
                            default:
                                if false {
                                    var v16 uint32 = 870
                                    var v17 float64 = (-2.25 * -2.25)
                                    _ = v16
                                    _ = v17
                                } else {
                                    v18 := GenS1{F0: 2, F1: 42}
                                    if (!true) {
                                        switch "\n\t" {
                                        case "hello":
                                            v19 := GenS1{F0: 42, F1: 255}
                                            _ = v19
                                        }
                                        v2 -= 32767
                                        if true {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            var v20 string = ("\n\t" + "a")
                                            _ = len(v4)
                                            _ = v20
                                        }
                                    }
                                    v21 := GenS2{F0: 42, F1: 1}
                                    v22 := GenS3{F0: -775, F1: -32768, F2: v2}
                                    _ = v18
                                    _ = v21
                                    _ = v22
                                }
                            }
                        } else {
                            if (1.5 == -2.25) {
                                if (42 < 880) {
                                    if false {
                                        switch (v2 >> ((42) & 63)) {
                                        case 1:
                                            v2 += (fi11 - (v2 - (v2 - fi11)))
                                        default:
                                            var v23 uint64 = 18446744073709551615
                                            switch "a" {
                                            case "foo bar":
                                                v23 = 144
                                                if (!(1000 > 1)) {
                                                    v23 += (v23 - (v23 | v23))
                                                    switch "a" {
                                                    case "":
                                                        var v24 uint = 239
                                                        _ = v24
                                                    default:
                                                        var v25 uint8 = 127
                                                        _ = v25
                                                    }
                                                    if false {
                                                        v26 := GenS3{F0: 42, F1: -32768, F2: (v2 ^ (fi10 + (v2 * -1)))}
                                                        _ = v26
                                                    }
                                                }
                                            default:
                                                if true {
                                                    var v27 uint8 = 128
                                                    v28 := GenS4{F0: -762}
                                                    switch "foo bar" {
                                                    case "":
                                                        fmt.Println(fi10)
                                                    case "hello":
                                                        var v29 string = "a"
                                                        _ = v29
                                                    case "a":
                                                        if true {
                                                            v2 += fi11
                                                        } else {
                                                            var v30 float64 = 1.5
                                                            _ = v30
                                                        }
                                                        var v31 uint32 = (v3 % ((v3) | 1))
                                                        _ = v31
                                                    }
                                                    _ = v27
                                                    _ = v28
                                                } else {
                                                    fmt.Println(v2, fi10, fi10)
                                                    if len(v1) > 0 {
                                                        v1[0] = "foo bar"
                                                    }
                                                }
                                            }
                                            if len(v1) > 0 {
                                                v1[0] = "hello"
                                            }
                                            _ = v23
                                        }
                                        v3 = 1
                                        fmt.Println(v2, fi11)
                                    }
                                    if false {
                                        if len(v1) > 0 {
                                            v1[0] = "foo bar"
                                        }
                                        v32 := GenS4{F0: 42}
                                        _ = v32
                                    } else {
                                        v3 = 42
                                    }
                                } else {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                                v3 += (v3 + v3)
                            } else {
                                if len(v4) > 0 {
                                    v4[0] = "hello"
                                }
                                v33 := GenS1{F0: 861, F1: 698}
                                var v34 uint16 = 42
                                _ = v33
                                _ = v34
                            }
                            var v35 int = fi10
                            _ = v35
                        }
                        v3 = (v3 ^ (v3 + v3))
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v36 int32 = -1
                        v37 := GenS3{F0: v36, F1: 32767, F2: genFn0(1, true, "\n\t")}
                        v38 := GenS5{F0: (("a" + ("foo bar" + "")) + "a")}
                        _ = v36
                        _ = v37
                        _ = v38
                    }
                }
            }
            v3 -= v3
            if true {
                v2 = 32768
                v2 = (v2 % ((genFn0(183, ("foo bar" == "a"), "")) | 1))
                if len(v4) > 0 {
                    v4[0] = "\n\t"
                }
            }
            _ = fi11
        }
        var v39 uint32 = v3
        v3 = v39
        _ = fi10
        _ = v39
    }
    for fi40 := 0; fi40 < 16; fi40++ {
        var v41 int8 = 80
        if true {
            for fi42 := 0; fi42 < 7; fi42++ {
                var v43 uint = 447
                switch genFn0(42, (!(1 == 1)), "a") {
                case 1:
                    v44 := GenS6{F0: v2, F1: 255, F2: 648}
                    _ = v44
                default:
                    v43 = v43
                }
                _ = fi42
                _ = v43
            }
        }
        _ = fi40
        _ = v41
    }
    for fi45 := 0; fi45 < 3; fi45++ {
        for fi46 := 0; fi46 < 17; fi46++ {
            v2 -= 4294967295
            v47 := GenS7{F0: 65535, F1: 32767}
            var v48 float32 = (1e10 + 1e10)
            _ = fi46
            _ = v47
            _ = v48
        }
        switch "\n\t" {
        case "":
            v49 := GenS8{F0: 21}
            _ = v49
        case "\n\t":
            v50 := GenS9{F0: 1}
            _ = v50
        case "foo bar":
            var v51 int = (-(v2))
            _ = v51
        default:
            _ = len(v4)
        }
        v52 := GenS8{F0: 100}
        _ = fi45
        _ = v52
    }
    var v53 int32 = -127
    if len(v4) > 0 {
        v4[0] = ("a" + "\n\t")
    }
    if false {
        for fi54 := 0; fi54 < 17; fi54++ {
            if ((v3 ^ 65535) >= 2) {
                v55 := GenS5{F0: "a"}
                if true {
                    v3 = 0
                }
                _ = v55
            } else {
                v53 = 2147483647
                v56 := GenS6{F0: -93, F1: 246, F2: -32768}
                fmt.Println(fi54, fi54, v2)
                _ = v56
            }
            switch fi54 {
            case 365:
                var v57 int32 = 0
                _ = v57
            case 128:
                var v58 bool = true
                _ = v58
            default:
                v2 -= -1
            }
            for fi59 := 0; fi59 < 5; fi59++ {
                if false {
                    if false {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if len(v1) > 0 {
                            v1[0] = ("hello" + "")
                        }
                        var v60 float64 = (1.5 + ((1.5 - 1e10) - 3.14159))
                        _ = v60
                    }
                } else {
                    v3 += 65535
                }
                if (false && ((885 <= 455) && (((("\n\t" + "foo bar") + (("" + "hello") + "")) + "") > (("" + ((("foo bar" + "a") + "foo bar") + "foo bar")) + "foo bar")))) {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v61 uint = uint(v3)
                    _ = v61
                } else {
                    fmt.Println(v53, v3)
                }
                v62 := GenS1{F0: 42, F1: 65535}
                _ = fi59
                _ = v62
            }
            _ = fi54
        }
    }
    var v63 int8 = 0
    for fi64 := 0; fi64 < 3; fi64++ {
        v1 = append(v1, "foo bar")
        _ = fi64
    }
    var v65 uint16 = 934
    switch genFn0(v65, false, "hello") {
    case 920:
        var v66 int = v2
        _ = v66
    case -128:
        if true {
            var v67 uint16 = v65
            v65 = v67
            switch v2 {
            case 32767:
                _ = len(v4)
            case 619:
                for fi68 := 0; fi68 < 18; fi68++ {
                    for fi69 := 0; fi69 < 18; fi69++ {
                        fmt.Println(v2, v2)
                        v70 := GenS3{F0: (v53 ^ 1), F1: 506, F2: genFn0(v67, true, "a")}
                        var v71 int16 = 32767
                        _ = fi69
                        _ = v70
                        _ = v71
                    }
                    if len(v4) > 0 {
                        v4[0] = ""
                    }
                    _ = fi68
                }
            }
            _ = v67
        }
    }
    fmt.Println(v53, v65, v53)
    for fi72 := 0; fi72 < 2; fi72++ {
        var v73 int8 = v63
        var v74 string = "\n\t"
        _ = fi72
        _ = v73
        _ = v74
    }
    for fi75 := 0; fi75 < 5; fi75++ {
        v3 = (v3 - v3)
        v76 := GenS10{F0: 0}
        _ = fi75
        _ = v76
    }
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v53:", v53)
    fmt.Println("v63:", v63)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v53
    _ = v63
    _ = v65
}
