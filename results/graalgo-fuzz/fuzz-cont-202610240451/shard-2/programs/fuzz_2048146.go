package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 bool; F1 int64; F2 int32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint64 }
func genFn0(p0 int, p1 uint32) string {
    p1 += 4294967295
    p0 = 9223372036854775807
    _ = p0
    _ = p1
    return "a"
}
func genFn1(p0 bool, p1 int64) uint {
    var v4 uint16 = 165
    var v5 int32 = 446
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 1
}

func main() {
    var v1 float64 = float64((1.5 / -2.25))
    var v2 int8 = 42
    v3 := []string{}
    switch "foo bar" {
    case "":
        if (-556 > -377) {
            if (true || (!(65535 <= 32767))) {
                var v6 int = 128
                _ = v6
            }
            if true {
                v7 := GenS1{F0: 255, F1: 155}
                for fi8 := 0; fi8 < 16; fi8++ {
                    if len(v3) > 0 {
                        v3[0] = "foo bar"
                    }
                    fmt.Println(fi8)
                    var v9 string = ("a" + "hello")
                    _ = fi8
                    _ = v9
                }
                for fi10 := 0; fi10 < 9; fi10++ {
                    var v11 float32 = 1.5
                    for fi12 := 0; fi12 < 10; fi12++ {
                        var v13 int = fi10
                        _ = fi12
                        _ = v13
                    }
                    fmt.Println(v2)
                    _ = fi10
                    _ = v11
                }
                _ = v7
            }
        } else {
            fmt.Println(v2)
            var v14 float32 = 1.5
            switch "a" {
            case "hello":
                if (42 < 65535) {
                    var v15 int = -868
                    fmt.Println(v2)
                    _ = v15
                } else {
                    v14 = -2.25
                    var v16 string = (((("\n\t" + "foo bar") + ("hello" + "a")) + "a") + ("\n\t" + ("" + "\n\t")))
                    _ = v16
                }
            }
            _ = v14
        }
    default:
        for fi17 := 0; fi17 < 6; fi17++ {
            v1 += (v1 - v1)
            var v18 uint64 = 0
            v18 = v18
            _ = fi17
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 7; fi19++ {
        if true {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v20 uint = 487
            _ = v20
        }
        v1 += v1
        _ = fi19
    }
    if true {
        if ((0.0 != (v1 - (v1 + (v1 / -2.25)))) || (1 < 146)) {
            if true {
                var v21 string = "\n\t"
                var v22 float32 = 3.14159
                if ((v2 << ((1) & 63)) > (v2 / (((v2 | v2)) | 1))) {
                    for fi23 := 0; fi23 < 13; fi23++ {
                        fmt.Println(fi23)
                        for fi24 := 0; fi24 < 2; fi24++ {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            fmt.Println(v2)
                            fmt.Println(v22, v1, fi23)
                            _ = fi24
                        }
                        _ = fi23
                    }
                } else {
                    if (genFn1(true, -445) == 18446744073709551615) {
                        v21 = "a"
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v25 uint = 799
                        _ = v25
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                _ = v21
                _ = v22
            }
            for fi26 := 0; fi26 < 13; fi26++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                if true {
                    fmt.Println(v2, fi26)
                    v3 = append(v3, "")
                    v2 = (^(v2))
                } else {
                    if (255 == genFn1(false, -328)) {
                        var v27 uint = 300
                        _ = v27
                    } else {
                        v1 = v1
                        var v28 uint64 = 228
                        if len(v3) > 0 {
                            v3[0] = "hello"
                        }
                        _ = v28
                    }
                    var v29 float64 = (v1 / 1.5)
                    _ = v29
                }
                v1 = 1e10
                _ = fi26
            }
        }
    } else {
        v2 = 100
        v30 := GenS2{F0: "hello", F1: 14}
        var v31 float64 = 0.0
        _ = v30
        _ = v31
    }
    switch 9223372036854775807 {
    case 32767:
        v32 := GenS2{F0: (("\n\t" + "\n\t") + ""), F1: 0}
        _ = v32
    case 767:
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
    case 0:
        v2 -= v2
    default:
        if ("a" <= "\n\t") {
            v2 -= 100
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            switch -349 {
            case 32767:
                var v33 uint = 255
                _ = v33
            case 128:
                v1 = v1
            case 550:
                v3 = append(v3, "")
            default:
                var v34 uint = genFn1((!(!false)), 42)
                _ = v34
            }
        } else {
            v1 += 1e10
            var v35 int64 = -426
            var v36 uint64 = 2
            _ = v35
            _ = v36
        }
    }
    if (232 < 255) {
        v37 := GenS3{F0: (false || false), F1: -1, F2: 56}
        fmt.Println(v2, v1)
        v38 := GenS4{F0: (584 > 32767)}
        _ = v37
        _ = v38
    }
    v39 := GenS5{F0: 128}
    if true {
        var v40 string = "a"
        _ = v40
    }
    for fi41 := 0; fi41 < 9; fi41++ {
        var v42 string = "hello"
        _ = fi41
        _ = v42
    }
    v2 -= (v2 * (v2 % ((-108) | 1)))
    if false {
        if (!false) {
            v43 := GenS3{F0: (-2147483648 != 378), F1: 1, F2: 2147483647}
            _ = v43
        }
    }
    _ = len(v3)
    if true {
        var v44 uint = genFn1(true, -673)
        _ = v44
    } else {
        v2 -= -1
        if true {
            if ((-2.25 * 1.5) > (-2.25 - float64(3.14159))) {
                v2 -= -128
                if (!(!false)) {
                    var v45 uint32 = 42
                    var v46 uint = genFn1((!((2 >= 1) || true)), 42)
                    switch "hello" {
                    case "foo bar":
                        switch ("foo bar" + "hello") {
                        case "\n\t":
                            v46 -= (v46 ^ v46)
                        case "foo bar":
                            if len(v3) > 0 {
                                v3[0] = "\n\t"
                            }
                        case "a":
                            fmt.Println(v46, v2)
                        }
                    case "a":
                        for fi47 := 0; fi47 < 6; fi47++ {
                            _ = len(v3)
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            if (true && ((float64(0.1) * (v1 / 0.1)) == 1e10)) {
                                _ = len(v3)
                                _ = len(v3)
                            } else {
                                v48 := GenS2{F0: "hello", F1: 0}
                                _ = v48
                            }
                            _ = fi47
                        }
                    case "":
                        v2 = (v2 << (((v46 ^ 255)) & 63))
                    }
                    _ = v45
                    _ = v46
                } else {
                    var v49 int64 = -1
                    for fi50 := 0; fi50 < 7; fi50++ {
                        fmt.Println(v49, fi50, v49)
                        _ = fi50
                    }
                    _ = v49
                }
            } else {
                if false {
                    v51 := GenS2{F0: (("" + "hello") + "\n\t"), F1: 13}
                    _ = v51
                }
                var v52 int32 = -872
                _ = v52
            }
            if (!(!true)) {
                v1 -= 0.0
                for fi53 := 0; fi53 < 4; fi53++ {
                    if true {
                        for fi54 := 0; fi54 < 1; fi54++ {
                            v2 = v2
                            v55 := GenS5{F0: 214}
                            var v56 int16 = 42
                            _ = len(v3)
                            var v57 float64 = ((-2.25 - v1) - v1)
                            _ = fi54
                            _ = v55
                            _ = v56
                            _ = v57
                        }
                        v58 := GenS5{F0: 238}
                        _ = v58
                    } else {
                        for fi59 := 0; fi59 < 20; fi59++ {
                            var v60 int32 = 1
                            _ = fi59
                            _ = v60
                        }
                    }
                    for fi61 := 0; fi61 < 18; fi61++ {
                        var v62 int8 = (v2 & v2)
                        v62 = v2
                        v3 = append(v3, "foo bar")
                        _ = fi61
                        _ = v62
                    }
                    if true {
                        var v63 int16 = 671
                        _ = len(v3)
                        var v64 uint32 = 65535
                        _ = v63
                        _ = v64
                    } else {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v3 = append(v3, "")
                        switch "foo bar" {
                        case "foo bar":
                            for fi65 := 0; fi65 < 20; fi65++ {
                                v2 = 42
                                _ = fi65
                            }
                        case "\n\t":
                            v66 := GenS6{F0: 42}
                            _ = v66
                        }
                    }
                    _ = fi53
                }
            }
            v1 = float64(0.1)
        }
    }
    v1 += v1
    var v67 int64 = -193
    if true {
        if true {
            v67 = (v67 | 0)
            v2 = (-(v2))
        }
    } else {
        v68 := GenS6{F0: 2}
        v69 := GenS5{F0: 0}
        _ = v68
        _ = v69
    }
    var v70 bool = (-539 > 128)
    v70 = (0 > 100)
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v67:", v67)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v67
    _ = v70
}
