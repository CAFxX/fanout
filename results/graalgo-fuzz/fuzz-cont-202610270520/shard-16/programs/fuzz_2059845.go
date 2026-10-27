package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 uint32; F1 int16; F2 string }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint32 }
func genFn0(p0 uint) uint64 {
    for fi4 := 0; fi4 < 20; fi4++ {
        fmt.Println(fi4, fi4)
        p0 = p0
        _ = fi4
    }
    v5 := GenS1{F0: true, F1: 1}
    _ = p0
    _ = v5
    return 42
}
func genFn1(p0 uint8, p1 int16) uint32 {
    if true {
        var v6 uint8 = (p0 % ((p0) | 1))
        var v7 uint = 1
        v8 := []int{1, 0, 688}
        _ = v6
        _ = v7
        _ = v8
    } else {
        v9 := GenS1{F0: false, F1: (p0 % (((p0 & p0)) | 1))}
        for fi10 := 0; fi10 < 10; fi10++ {
            if (-2.25 >= float64(3.14159)) {
                if (true || true) {
                    v11 := []int{fi10, (fi10 % (((^(fi10))) | 1))}
                    _ = v11
                }
            }
            switch "a" {
            case "a":
                v12 := GenS1{F0: true, F1: p0}
                _ = v12
            case "foo bar":
                p1 = 32767
            case "hello":
                v13 := GenS1{F0: false, F1: (^(p0))}
                _ = v13
            }
            _ = fi10
        }
        p1 -= p1
        _ = v9
    }
    for fi14 := 0; fi14 < 2; fi14++ {
        p0 = p0
        switch "a" {
        case "hello":
            if false {
                switch (fi14 ^ -836) {
                case 2147483647:
                    p1 = (p1 * p1)
                case 32767:
                    if (4294967295 != 18446744073709551615) {
                        fmt.Println(p0, p0, p1)
                    }
                case 0:
                    var v15 bool = true
                    _ = v15
                default:
                    v16 := []int{128, 32768, -654, (fi14 << ((18446744073709551615) & 63))}
                    _ = v16
                }
                for fi17 := 0; fi17 < 15; fi17++ {
                    var v18 float64 = ((0.1 - (((1.5 + 1.5) + ((0.0 - (1.5 - 1.5)) / (1e10 + 0.0))) / 1.5)) / 1.0)
                    _ = fi17
                    _ = v18
                }
            }
        case "a":
            fmt.Println(p0, p1, fi14)
        }
        fmt.Println(p1, fi14, p0)
        _ = fi14
    }
    p0 = 0
    p0 -= (p0 >> ((81) & 63))
    switch "\n\t" {
    case "":
        var v19 uint64 = 1
        _ = v19
    }
    _ = p0
    _ = p1
    return 2
}

func main() {
    var v1 string = "foo bar"
    var v2 float32 = 1e10
    var v3 uint32 = 1
    v20 := GenS2{F0: ("hello" + v1), F1: (v1 + "a")}
    v21 := GenS2{F0: "\n\t", F1: v1}
    v22 := GenS1{F0: (42 != 877), F1: 42}
    for fi23 := 0; fi23 < 12; fi23++ {
        if (true && (!false)) {
            var v24 bool = (128 != 250)
            v25 := GenS2{F0: ((("" + (("foo bar" + "\n\t") + ("foo bar" + v1))) + "") + v1), F1: v1}
            for fi26 := 0; fi26 < 14; fi26++ {
                fmt.Println(v3, fi23, fi26)
                if v24 {
                    v2 -= 0.1
                    v1 = (v1 + (v1 + v1))
                } else {
                    var v27 float32 = ((v2 + (v2 / v2)) - v2)
                    switch fi23 {
                    case 42:
                        var v28 float64 = (float64(v27) - 3.14159)
                        _ = v28
                    case 0:
                        v29 := []int{(fi26 + (fi23 ^ (fi23 * (-(fi23))))), (fi26 + fi23), 680}
                        _ = v29
                    case 128:
                        var v30 uint8 = 171
                        _ = v30
                    }
                    _ = v27
                }
                _ = fi26
            }
            _ = v24
            _ = v25
        }
        v31 := []int{(fi23 << ((4294967295) & 63)), fi23, (fi23 | (fi23 & 9223372036854775807)), 0}
        _ = fi23
        _ = v31
    }
    var v32 uint32 = v3
    v33 := []int{127, 1, 42}
    v1 = (("" + v1) + (v1 + (v1 + "")))
    var v34 int16 = 0
    switch 42 {
    case 32767:
        for fi35 := 0; fi35 < 2; fi35++ {
            if (v1 > v1) {
                for fi36 := 0; fi36 < 6; fi36++ {
                    v37 := GenS2{F0: "hello", F1: (v1 + "")}
                    var v38 bool = false
                    fmt.Println(fi36)
                    _ = fi36
                    _ = v37
                    _ = v38
                }
                v39 := GenS3{F0: (v3 ^ (v3 + (v3 * (^(v3))))), F1: v34, F2: v1}
                _ = v39
            } else {
                for fi40 := 0; fi40 < 8; fi40++ {
                    if true {
                        v41 := GenS4{F0: ((v34 > 32767) || (827 < v32))}
                        var v42 float32 = float32(1e10)
                        _ = v41
                        _ = v42
                    } else {
                        var v43 float32 = 1e10
                        var v44 bool = false
                        var v45 bool = (1 >= uint8(v32))
                        _ = v43
                        _ = v44
                        _ = v45
                    }
                    if false {
                        v33 = append(v33, (fi40 % ((fi40) | 1)))
                        var v46 uint8 = 128
                        switch (fi40 % ((fi35) | 1)) {
                        case -2147483648:
                            if len(v33) >= 2 {
                                v33 = v33[1:len(v33)]
                            }
                        default:
                            _ = len(v33)
                        }
                        _ = v46
                    } else {
                        v47 := GenS5{F0: (v32 + (v32 + v32))}
                        _ = v47
                    }
                    if (((107 != 958) || false) && false) {
                        v48 := GenS1{F0: (!true), F1: 128}
                        fmt.Println(fi40)
                        if false {
                            fmt.Println(fi40, v32)
                        }
                        _ = v48
                    }
                    _ = fi40
                }
                v49 := GenS3{F0: (v32 ^ v32), F1: (v34 - 42), F2: ""}
                fmt.Println(v32)
                _ = v49
            }
            _ = fi35
        }
    case 128:
        if len(v33) >= 2 {
            v33 = v33[1:len(v33)]
        }
    case 1:
        var v50 uint = 42
        _ = v50
    default:
        var v51 float64 = 1e10
        _ = v51
    }
    switch -11 {
    case -9223372036854775808:
        if ((v34 % ((v34) | 1)) <= v34) {
            var v52 int = 42
            var v53 uint = 2
            _ = v52
            _ = v53
        }
    case -21:
        for fi54 := 0; fi54 < 20; fi54++ {
            v34 -= v34
            v34 -= v34
            fmt.Println(v3)
            _ = fi54
        }
    case -978:
        _ = len(v33)
    default:
        v1 = ""
    }
    if len(v33) > 0 {
        v33[0] = 128
    }
    v55 := GenS2{F0: v1, F1: "\n\t"}
    var v56 uint32 = (v3 % ((42) | 1))
    var v57 uint16 = 42
    for fi58 := 0; fi58 < 2; fi58++ {
        v56 = v3
        fmt.Println(v1)
        var v59 int8 = -128
        _ = fi58
        _ = v59
    }
    var v60 bool = (false && true)
    v60 = v60
    switch v1 {
    case "foo bar":
        var v61 float64 = 0.1
        _ = v61
    case "a":
        for fi62 := 0; fi62 < 7; fi62++ {
            v1 = v1
            v63 := GenS1{F0: false, F1: 127}
            v33 = append(v33, fi62)
            _ = fi62
            _ = v63
        }
    case "":
        v33 = append(v33, int(v32))
    default:
        var v64 string = ("hello" + v1)
        _ = v64
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v55
    _ = v56
    _ = v57
    _ = v60
}
