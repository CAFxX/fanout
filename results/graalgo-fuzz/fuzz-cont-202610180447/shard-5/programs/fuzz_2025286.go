package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 string; F1 uint64 }
type GenS6 struct { F0 int32; F1 uint32 }
func genFn0() uint32 {
    var v7 uint16 = 2
    v8 := GenS1{F0: true, F1: (("a" + "\n\t") + "foo bar")}
    if (1 <= 1) {
        v9 := []int{128}
        v7 = (v7 + 2)
        _ = v9
    } else {
        switch 128 {
        case -227:
            switch -469 {
            case -9223372036854775808:
                for fi10 := 0; fi10 < 15; fi10++ {
                    if (0 >= 1) {
                        v7 = 1
                        v11 := GenS2{F0: 1000}
                        var v12 int = fi10
                        _ = v11
                        _ = v12
                    }
                    _ = fi10
                }
            case 127:
                if (206 <= 821) {
                    v13 := []int{-349, -912}
                    _ = v13
                } else {
                    switch 127 {
                    case 127:
                        v14 := []int{-9223372036854775808, -936}
                        _ = v14
                    default:
                        v15 := GenS3{F0: 42, F1: (v7 / ((v7) | 1))}
                        _ = v15
                    }
                }
            }
        case -618:
            for fi16 := 0; fi16 < 18; fi16++ {
                for fi17 := 0; fi17 < 4; fi17++ {
                    v18 := []int{}
                    v7 = v7
                    _ = fi17
                    _ = v18
                }
                _ = fi16
            }
        }
        for fi19 := 0; fi19 < 15; fi19++ {
            fmt.Println(v7, fi19)
            _ = fi19
        }
        var v20 int32 = -1
        _ = v20
    }
    _ = v7
    _ = v8
    return 2
}
func genFn1(p0 int) uint {
    switch "hello" {
    case "":
        v21 := []int{p0}
        _ = v21
    case "hello":
        fmt.Println(p0)
    case "foo bar":
        if (!false) {
            p0 = p0
            if false {
                if (2 >= 1) {
                    v22 := GenS4{F0: 1}
                    for fi23 := 0; fi23 < 11; fi23++ {
                        fmt.Println(p0, fi23)
                        _ = fi23
                    }
                    v24 := GenS4{F0: -1}
                    _ = v22
                    _ = v24
                } else {
                    for fi25 := 0; fi25 < 19; fi25++ {
                        for fi26 := 0; fi26 < 20; fi26++ {
                            var v27 int8 = 73
                            v28 := []int{(fi26 - fi25), fi26, (fi25 * (^(fi26))), (fi25 + 0)}
                            if len(v28) >= 2 {
                                v28 = v28[1:len(v28)]
                            }
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                        v29 := []int{}
                        _ = fi25
                        _ = v29
                    }
                }
                v30 := []int{42, 390}
                _ = v30
            } else {
                fmt.Println(p0)
                if true {
                    v31 := []int{(p0 & p0), (p0 * p0), p0}
                    _ = v31
                }
            }
        }
    }
    var v32 float64 = -2.25
    v32 -= v32
    if true {
        fmt.Println(p0)
    } else {
        p0 += p0
        p0 += (p0 | 32768)
        for fi33 := 0; fi33 < 13; fi33++ {
            var v34 int8 = 68
            p0 += 4294967295
            if (v32 > v32) {
                p0 = (fi33 + fi33)
            } else {
                for fi35 := 0; fi35 < 11; fi35++ {
                    var v36 int = fi35
                    v32 += v32
                    _ = fi35
                    _ = v36
                }
                switch "" {
                case "hello":
                    var v37 int8 = (^(v34))
                    _ = v37
                }
                switch 9223372036854775807 {
                case 55:
                    v38 := []int{}
                    _ = v38
                case 478:
                    var v39 float64 = (v32 * v32)
                    _ = v39
                default:
                    v40 := []int{fi33, (p0 >> ((255) & 63)), (p0 & fi33), (p0 & (p0 & fi33))}
                    _ = v40
                }
            }
            _ = fi33
            _ = v34
        }
    }
    _ = p0
    _ = v32
    return 2
}

func main() {
    var v1 int = 2147483647
    v2 := []string{((("hello" + ((("\n\t" + "hello") + "") + (("" + ("\n\t" + (("hello" + ((("" + "hello") + "hello") + ("\n\t" + "hello"))) + "\n\t"))) + ""))) + "foo bar") + ((("a" + "hello") + (("\n\t" + ("" + "hello")) + ("foo bar" + "\n\t"))) + "foo bar"))}
    v3 := []uint32{}
    var v4 uint32 = 65535
    var v5 uint32 = (v4 ^ v4)
    var v6 string = ""
    if false {
        fmt.Println(v5, v4)
        v41 := GenS5{F0: "a", F1: 362}
        var v42 uint64 = 789
        _ = v41
        _ = v42
    }
    var v43 string = "hello"
    v43 = (v43 + "\n\t")
    fmt.Println(v43)
    if (!(float32(-2.25) < ((float32((1.5 / 3.14159)) - -2.25) * (1e10 / float32(((-2.25 - 1e10) + 0.0)))))) {
        _ = len(v2)
        var v44 int8 = 67
        _ = v44
    } else {
        if false {
            var v45 int32 = 2147483647
            _ = v45
        } else {
            if ((0.0 * (1.5 + 3.14159)) < (0.0 / 1.0)) {
                v5 -= 690
                v43 = v43
            }
            v43 = ((v43 + v43) + "a")
            var v46 uint32 = (^(v4))
            _ = v46
        }
    }
    v47 := GenS1{F0: (!true), F1: v6}
    var v48 bool = (!false)
    var v49 uint64 = 930
    fmt.Println(v4, v4, v48)
    for fi50 := 0; fi50 < 11; fi50++ {
        if len(v3) > 0 {
            v3[0] = (v5 | v4)
        }
        _ = fi50
    }
    v51 := GenS6{F0: -564, F1: genFn0()}
    if ((v1 & 128) >= v1) {
        var v52 string = ((v43 + (v43 + ("a" + (v6 + v6)))) + "foo bar")
        var v53 int64 = 0
        fmt.Println(v53, v49)
        _ = v52
        _ = v53
    } else {
        v1 = 9223372036854775807
    }
    for fi54 := 0; fi54 < 17; fi54++ {
        var v55 uint32 = (v4 & 65535)
        fmt.Println(v4)
        fmt.Println(v1)
        _ = fi54
        _ = v55
    }
    var v56 uint32 = v4
    v2 = append(v2, ("" + v6))
    var v57 uint64 = (^(v49))
    var v58 int = (v1 / ((485) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v51
    _ = v56
    _ = v57
    _ = v58
    _ = v6
}
