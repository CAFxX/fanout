package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 uint16; F1 bool; F2 uint16 }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 uint8; F1 uint8; F2 bool }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 bool; F1 string }
type GenS8 struct { F0 int }
func genFn0(p0 int, p1 uint8, p2 int16) bool {
    v4 := GenS1{F0: 18446744073709551615, F1: 18446744073709551615}
    v5 := GenS2{F0: 255, F1: true, F2: 32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return false
}

func main() {
    var v1 uint32 = 2
    v2 := []string{("a" + "foo bar")}
    var v3 float64 = -2.25
    var v6 uint64 = 4294967295
    v7 := GenS1{F0: (v6 % (((v6 >> ((1) & 63))) | 1)), F1: (v6 | 18446744073709551615)}
    var v8 int = 32767
    var v9 int = v8
    switch "" {
    case "":
        v2 = append(v2, ("" + ""))
    case "hello":
        switch v9 {
        case 1:
            v10 := GenS3{F0: (v6 % ((4294967295) | 1)), F1: 0}
            _ = v10
        default:
            var v11 float32 = 1e10
            _ = v11
        }
    }
    v12 := GenS2{F0: 261, F1: ((936 < 335) && (4294967295 <= 42)), F2: 2}
    v13 := GenS4{F0: 42, F1: 131, F2: true}
    v9 += v9
    for fi14 := 0; fi14 < 9; fi14++ {
        switch ((("" + "foo bar") + "a") + ("foo bar" + "")) {
        case "\n\t":
            switch v8 {
            case 0:
                var v15 uint32 = v1
                _ = v15
            }
        case "hello":
            if (42 != v6) {
                v3 = v3
            }
        }
        _ = fi14
    }
    switch "\n\t" {
    case "a":
        v6 = v6
    default:
        fmt.Println(v8, v1)
    }
    v16 := GenS1{F0: (v6 >> ((65535) & 63)), F1: 42}
    if false {
        v2 = append(v2, ("hello" + "hello"))
        if false {
            fmt.Println(v6, v6, v6)
            for fi17 := 0; fi17 < 6; fi17++ {
                var v18 float32 = 0.1
                var v19 uint16 = 255
                for fi20 := 0; fi20 < 11; fi20++ {
                    switch fi20 {
                    case 128:
                        var v21 uint8 = 24
                        v21 = v21
                        _ = v21
                    case 4294967295:
                        v6 = v6
                    }
                    switch "a" {
                    case "":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case "hello":
                        var v22 uint16 = 889
                        _ = v22
                    }
                    _ = fi20
                }
                _ = fi17
                _ = v18
                _ = v19
            }
            if (false || false) {
                var v23 float64 = v3
                if false {
                    fmt.Println(v23)
                } else {
                    v24 := GenS4{F0: 1, F1: 117, F2: true}
                    if false {
                        if false {
                            if len(v2) > 0 {
                                v2[0] = "hello"
                            }
                            var v25 int = (v8 ^ v9)
                            switch "a" {
                            case "\n\t":
                                v9 += v9
                            case "":
                                switch v9 {
                                case 63:
                                    var v26 int32 = 42
                                    _ = v26
                                case 42:
                                    var v27 uint64 = (v6 - (v6 & (^(v6))))
                                    _ = v27
                                case -400:
                                    fmt.Println(v25, v23)
                                }
                            case "a":
                                v6 = v6
                            }
                            _ = v25
                        } else {
                            if true {
                                if false {
                                    v2 = append(v2, "a")
                                    v2 = append(v2, ("" + ""))
                                } else {
                                    v28 := GenS5{F0: 822}
                                    var v29 bool = (((!((v1 << ((65535) & 63)) >= 752)) && (!true)) && (0 < -2147483648))
                                    var v30 uint64 = 0
                                    _ = v28
                                    _ = v29
                                    _ = v30
                                }
                                if true {
                                    v31 := GenS1{F0: v6, F1: 1}
                                    for fi32 := 0; fi32 < 13; fi32++ {
                                        var v33 uint64 = (v6 << ((238) & 63))
                                        _ = fi32
                                        _ = v33
                                    }
                                    v34 := GenS6{F0: -193}
                                    _ = v31
                                    _ = v34
                                }
                            } else {
                                for fi35 := 0; fi35 < 6; fi35++ {
                                    v6 = v6
                                    if true {
                                        v1 = (v1 ^ v1)
                                        v36 := GenS4{F0: 1, F1: 10, F2: true}
                                        var v37 uint8 = 90
                                        _ = v36
                                        _ = v37
                                    } else {
                                        if false {
                                            _ = len(v2)
                                            for fi38 := 0; fi38 < 15; fi38++ {
                                                v9 -= (-(fi38))
                                                _ = fi38
                                            }
                                            if true {
                                                v39 := GenS7{F0: true, F1: ("hello" + "\n\t")}
                                                for fi40 := 0; fi40 < 1; fi40++ {
                                                    switch "hello" {
                                                    case "foo bar":
                                                        var v41 int16 = int16(v9)
                                                        _ = v41
                                                    case "\n\t":
                                                        var v42 int16 = -1
                                                        _ = v42
                                                    case "hello":
                                                        v8 -= v9
                                                    }
                                                    if (1 != 834) {
                                                        if true {
                                                            v43 := GenS8{F0: (v8 / ((fi35) | 1))}
                                                            var v44 bool = true
                                                            switch "" {
                                                            case "\n\t":
                                                                v8 = fi40
                                                            case "":
                                                                v2 = append(v2, (("\n\t" + "") + ("foo bar" + ("hello" + "foo bar"))))
                                                            case "hello":
                                                                fmt.Println(v1)
                                                            default:
                                                                v8 = -422
                                                            }
                                                            _ = v43
                                                            _ = v44
                                                        } else {
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            var v45 uint64 = (v6 * v6)
                                                            _ = v45
                                                        }
                                                        v46 := GenS8{F0: 0}
                                                        v2 = append(v2, "\n\t")
                                                        _ = v46
                                                    }
                                                    _ = fi40
                                                }
                                                v23 = v3
                                                _ = v39
                                            }
                                        } else {
                                            fmt.Println(fi35, v8, v8)
                                        }
                                    }
                                    _ = fi35
                                }
                            }
                        }
                        v47 := GenS5{F0: 1}
                        if true {
                            if false {
                                v48 := GenS3{F0: (v6 + v6), F1: int16(v6)}
                                _ = v48
                            }
                            fmt.Println(v9, v8)
                            v49 := GenS8{F0: 127}
                            _ = v49
                        }
                        _ = v47
                    }
                    _ = v24
                }
                _ = v23
            }
        }
    }
    var v50 uint32 = (v1 - (v1 / (((^(v1))) | 1)))
    v51 := GenS6{F0: -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v50
    _ = v51
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
