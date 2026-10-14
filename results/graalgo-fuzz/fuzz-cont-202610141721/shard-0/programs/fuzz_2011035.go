package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 uint64; F1 bool }
func genFn0(p0 uint64, p1 int16, p2 int16) uint8 {
    var v7 float32 = (((1.5 - (1.5 + (((0.1 / 0.1) / 0.1) - ((0.0 - (1e10 * ((float32(3.14159) + 0.1) / -2.25))) - (-2.25 * 0.1))))) - (0.1 - -2.25)) / 1.5)
    v8 := []int{887, 127, 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 2
}
func genFn1() string {
    v9 := []int{1}
    var v10 uint8 = 1
    v11 := GenS1{F0: 2}
    v10 -= 144
    _ = v10
    _ = v11
    _ = v9
    return "a"
}

func main() {
    var v1 int64 = 0
    var v2 float64 = 0.1
    var v3 uint8 = 2
    v4 := []string{"hello", "hello", "\n\t", ""}
    var v5 float32 = -2.25
    v6 := make(map[string]uint32)
    fmt.Println(v3, v5, v3)
    if (true && (v1 == (v1 + (v1 ^ v1)))) {
        fmt.Println(v3, v1)
        v12 := GenS2{F0: "hello", F1: 42}
        _ = v12
    }
    if (false && true) {
        v13 := GenS2{F0: "", F1: 1}
        var v14 uint16 = 42
        var v15 int16 = 461
        _ = v13
        _ = v14
        _ = v15
    }
    switch (("foo bar" + "hello") + ("foo bar" + "a")) {
    case "foo bar":
        v16 := GenS3{F0: -2147483648, F1: 25}
        _ = v16
    case "\n\t":
        var v17 uint8 = v3
        _ = v17
    }
    for fi18 := 0; fi18 < 16; fi18++ {
        v2 += v2
        var v19 uint = 2
        _ = fi18
        _ = v19
    }
    v4 = append(v4, "foo bar")
    fmt.Println(v3, v3)
    for fi20 := 0; fi20 < 3; fi20++ {
        for fi21 := 0; fi21 < 19; fi21++ {
            if (!("" != "")) {
                fmt.Println(v2, v2, fi21)
            } else {
                var v22 uint16 = 460
                _ = len(v4)
                _ = v22
            }
            switch (fi20 / (((fi21 / ((fi20) | 1))) | 1)) {
            case 32768:
                fmt.Println(fi20, v2, fi21)
            case -383:
                v23 := GenS3{F0: -1, F1: 1}
                var v24 uint64 = 1
                var v25 uint64 = (v24 | 1)
                fmt.Println(v25, v24)
                if (true || false) {
                    switch "foo bar" {
                    case "hello":
                        var v26 int16 = 1
                        _ = v26
                    case "":
                        var v27 string = "hello"
                        _ = v27
                    case "foo bar":
                        var v28 int64 = -447
                        _ = v28
                    default:
                        if (2 >= 146) {
                            v1 = -1
                            v3 -= (v3 << ((241) & 63))
                            if true {
                                var v29 float64 = v2
                                var v30 float32 = v5
                                v31 := GenS4{F0: (v25 / (((v24 | v25)) | 1)), F1: false}
                                _ = v29
                                _ = v30
                                _ = v31
                            }
                        }
                    }
                }
                var v32 int8 = 1
                fmt.Println(fi21, v24, v5)
                _ = v23
                _ = v24
                _ = v25
                _ = v32
            case -259:
                var v33 bool = (true && (!(42 > 65535)))
                _ = v33
            }
            fmt.Println(fi20)
            _ = fi21
        }
        for fi34 := 0; fi34 < 14; fi34++ {
            switch ("hello" + ("\n\t" + (((((("hello" + (("" + ("hello" + "foo bar")) + "hello")) + ("hello" + "hello")) + "a") + "hello") + "a") + (("hello" + "a") + "\n\t")))) {
            case "\n\t":
                var v35 uint16 = 498
                _ = v35
            default:
                var v36 bool = (!(false || true))
                _ = v36
            }
            _ = fi34
        }
        if false {
            switch -551 {
            case -2147483648:
                for fi37 := 0; fi37 < 20; fi37++ {
                    if (!true) {
                        var v38 bool = (float32((((v2 + v2) * v2) - v2)) <= v5)
                        v4 = append(v4, "a")
                        v38 = true
                        _ = v38
                    } else {
                        var v39 uint64 = 42
                        v3 -= (v3 - (v3 / ((163) | 1)))
                        fmt.Println(v3, v5, v3)
                        var v40 int16 = 56
                        _ = v39
                        _ = v40
                    }
                    v41 := GenS4{F0: 18446744073709551615, F1: true}
                    _ = fi37
                    _ = v41
                }
            case -9223372036854775808:
                switch fi20 {
                case 32768:
                    v42 := GenS3{F0: -1, F1: 114}
                    _ = v42
                case 32767:
                    if (true || false) {
                        v43 := v6["foo bar"]
                        _ = v43
                    }
                }
            case 32767:
                var v44 string = ("\n\t" + "")
                _ = v44
            default:
                if false {
                    for fi45 := 0; fi45 < 6; fi45++ {
                        if (!(!true)) {
                            var v46 uint = 0
                            _ = v46
                        } else {
                            fmt.Println(fi45, fi45)
                        }
                        if ((((v5 / (3.14159 * float32(-2.25))) + (v5 + (v5 + 1.5))) >= ((1e10 * (0.1 - v5)) * 0.1)) || false) {
                            v1 = v1
                        }
                        _ = fi45
                    }
                }
            }
            v47 := GenS1{F0: 2147483648}
            if true {
                var v48 uint16 = 42
                _ = v48
            } else {
                if true {
                    fmt.Println(v1)
                    switch ("a" + "\n\t") {
                    case "hello":
                        var v49 float32 = v5
                        _ = v49
                    case "\n\t":
                        v2 -= v2
                    default:
                        if false {
                            v5 += 1e10
                            v50 := GenS4{F0: 4294967295, F1: true}
                            _ = v50
                        } else {
                            var v51 uint32 = 2147483647
                            _ = v51
                        }
                    }
                    if len(v4) > 0 {
                        v4[0] = ("hello" + ("\n\t" + "a"))
                    }
                }
                for fi52 := 0; fi52 < 1; fi52++ {
                    fmt.Println(v5, v1)
                    var v53 int = (fi20 - (fi20 - 128))
                    if (false || false) {
                        if (57 == 2147483647) {
                            var v54 int = v53
                            if (255 < v3) {
                                if true {
                                    switch (v54 % ((v54) | 1)) {
                                    case 743:
                                        switch fi52 {
                                        case 128:
                                            v6[(("\n\t" + "hello") + "")] = 4294967295
                                        }
                                    default:
                                        var v55 uint16 = 378
                                        _ = v55
                                    }
                                    var v56 uint8 = genFn0(4294967295, 42, -960)
                                    var v57 float64 = v2
                                    _ = v56
                                    _ = v57
                                }
                                if true {
                                    if true {
                                        if true {
                                            var v58 string = "\n\t"
                                            switch v58 {
                                            case "\n\t":
                                                v59 := GenS3{F0: -687, F1: 0}
                                                _ = v59
                                            case "hello":
                                                var v60 string = v58
                                                v61 := v6[(v58 + ((v60 + v58) + (v60 + v60)))]
                                                _ = len(v6)
                                                _ = v60
                                                _ = v61
                                            }
                                            var v62 uint16 = 2
                                            _ = v58
                                            _ = v62
                                        }
                                        v3 = 2
                                    }
                                } else {
                                    fmt.Println(fi20, v3)
                                    switch fi20 {
                                    case -1:
                                        var v63 string = "foo bar"
                                        _ = v63
                                    case 0:
                                        if len(v4) > 0 {
                                            v4[0] = ("foo bar" + ("\n\t" + ((("a" + "foo bar") + "") + "foo bar")))
                                        }
                                    }
                                }
                            }
                            var v64 uint16 = 0
                            _ = v54
                            _ = v64
                        }
                        if (v53 > v53) {
                            v6[""] = 2147483647
                        }
                    }
                    _ = fi52
                    _ = v53
                }
            }
            _ = v47
        }
        _ = fi20
    }
    v1 += (v1 / ((9223372036854775807) | 1))
    if ((v5 + v5) <= 0.1) {
        for fi65 := 0; fi65 < 7; fi65++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi65
        }
        var v66 string = "foo bar"
        _ = v66
    } else {
        _ = len(v4)
        var v67 int16 = -885
        v68 := GenS3{F0: int32(v67), F1: 106}
        _ = v67
        _ = v68
    }
    v3 = 2
    var v69 uint = 12
    switch "" {
    case "foo bar":
        var v70 int = 9223372036854775807
        _ = v70
    case "hello":
        var v71 float32 = v5
        _ = v71
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v69
}
