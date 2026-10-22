package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int; F1 int }
type GenS5 struct { F0 uint16; F1 uint8; F2 bool }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint64; F1 uint }
type GenS8 struct { F0 uint; F1 int16; F2 uint8 }
type GenS9 struct { F0 int; F1 int32 }
func genFn0(p0 int) uint {
    v6 := []int{(p0 & (p0 % ((p0) | 1))), 1, 0, p0}
    fmt.Println(p0)
    for fi7 := 0; fi7 < 18; fi7++ {
        p0 += fi7
        fmt.Println(p0, p0)
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi7
    }
    p0 += 1
    p0 = p0
    _ = p0
    _ = v6
    return 255
}
func genFn1(p0 string, p1 string, p2 int) int32 {
    var v8 uint64 = 0
    v9 := GenS1{F0: 1000}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 2147483647
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{}
    var v3 bool = false
    var v4 uint8 = 80
    var v5 float32 = (((0.1 - -2.25) + (1e10 + 1.5)) - ((((float32((1e10 + -2.25)) + 0.0) / 1e10) + 3.14159) + -2.25))
    v4 = v4
    for fi10 := 0; fi10 < 12; fi10++ {
        switch fi10 {
        case -462:
            switch "" {
            case "\n\t":
                fmt.Println(fi10)
            case "hello":
                v5 -= float32((1e10 + (-2.25 + -2.25)))
            default:
                if len(v2) > 0 {
                    v2[0] = ("hello" + "")
                }
            }
        case 1:
            if false {
                for fi11 := 0; fi11 < 2; fi11++ {
                    switch "hello" {
                    case "\n\t":
                        v12 := GenS2{F0: 83, F1: -128}
                        _ = v12
                    case "hello":
                        var v13 uint64 = 4294967295
                        _ = v13
                    case "foo bar":
                        var v14 uint32 = 356
                        _ = v14
                    default:
                        switch 514 {
                        case -9223372036854775808:
                            v3 = ((("foo bar" + "") + "foo bar") == (("hello" + (("foo bar" + "") + (("hello" + ("" + "a")) + "hello"))) + "foo bar"))
                        case 971:
                            var v15 int64 = 0
                            _ = v15
                        default:
                            v4 -= (v4 - 122)
                        }
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi11
                }
            }
        }
        var v16 int8 = 99
        _ = fi10
        _ = v16
    }
    switch -567 {
    case -9223372036854775808:
        var v17 int = 32768
        _ = v17
    case 564:
        if v3 {
            var v18 uint64 = 477
            if v3 {
                _ = len(v1)
            } else {
                var v19 float32 = ((-2.25 / float32(1.5)) / v5)
                if (((((1e10 * (((-2.25 - (float64((((1.5 / (float32(-2.25) / 1e10)) + ((v5 - v19) - (v19 * -2.25))) + 0.0)) + ((3.14159 / 1.0) / -2.25))) * float64(-2.25)) - 1e10)) + (-2.25 / float64(v5))) - (-2.25 + -2.25)) / 1.5) <= -2.25) {
                    if true {
                        v3 = v3
                        if v3 {
                            v1["a"] = 1
                            var v20 uint64 = v18
                            _ = v20
                        }
                        for fi21 := 0; fi21 < 8; fi21++ {
                            v4 = (v4 * 32)
                            switch "a" {
                            case "a":
                                if (-2147483648 <= -701) {
                                    switch fi21 {
                                    case 127:
                                        v22 := GenS3{F0: (v4 & v4)}
                                        _ = v22
                                    }
                                    switch "" {
                                    case "":
                                        v19 = ((float32(1.5) - float32(1e10)) * (1e10 / (float32(0.1) - -2.25)))
                                    case "hello":
                                        v5 = (v19 * v19)
                                    default:
                                        v18 = v18
                                    }
                                }
                            }
                            var v23 bool = v3
                            _ = fi21
                            _ = v23
                        }
                    }
                    var v24 bool = v3
                    if (770 != 255) {
                        if v3 {
                            if ("" != "") {
                                var v25 int64 = 42
                                _ = v25
                            } else {
                                v4 -= v4
                                var v26 int64 = 1
                                fmt.Println(v3, v26)
                                _ = v26
                            }
                            v3 = (0 < -2147483648)
                            fmt.Println(v24, v19, v18)
                        } else {
                            v2 = append(v2, ("a" + "a"))
                            v27 := v1["foo bar"]
                            _ = v27
                        }
                        v19 += float32(float64((v19 + float32(-2.25))))
                    } else {
                        v28 := GenS4{F0: 556, F1: 873}
                        _ = v28
                    }
                    _ = v24
                } else {
                    if ((3.14159 != (3.14159 + (float64(v19) / -2.25))) || ((v3 || v3) || false)) {
                        v3 = ((!(42 <= -1)) && v3)
                        if (!v3) {
                            v18 += v18
                            var v29 int8 = 102
                            v30 := GenS4{F0: 32768, F1: -851}
                            _ = v29
                            _ = v30
                        } else {
                            v3 = (42 == 42)
                        }
                        v1[("hello" + "")] = 2
                    }
                }
                switch ("\n\t" + "") {
                case "":
                    for fi31 := 0; fi31 < 10; fi31++ {
                        switch ("a" + "a") {
                        case "hello":
                            if false {
                                var v32 int = -471
                                _ = v32
                            } else {
                                v5 -= ((v19 + v19) / (-2.25 + 3.14159))
                                var v33 int = 32767
                                var v34 int32 = 461
                                _ = v33
                                _ = v34
                            }
                        }
                        if v3 {
                            var v35 int8 = 42
                            _ = v35
                        }
                        _ = fi31
                    }
                case "hello":
                    var v36 uint16 = 523
                    _ = v36
                case "\n\t":
                    v37 := GenS5{F0: 32767, F1: uint8(v18), F2: (!(!v3))}
                    _ = v37
                default:
                    v38 := GenS6{F0: 1}
                    _ = v38
                }
                _ = v19
            }
            _ = v18
        }
    case -272:
        switch 9223372036854775807 {
        case 32768:
            v39 := GenS7{F0: 113, F1: genFn0(2147483647)}
            _ = v39
        case 127:
            var v40 uint32 = 2
            _ = v40
        default:
            delete(v1, ((("foo bar" + ((("\n\t" + "hello") + ("\n\t" + "hello")) + "a")) + "a") + ""))
        }
    }
    var v41 float64 = 0.0
    v5 = (v5 - 1.5)
    for fi42 := 0; fi42 < 6; fi42++ {
        v43 := GenS2{F0: 255, F1: 42}
        _ = fi42
        _ = v43
    }
    var v44 uint64 = 18446744073709551615
    v3 = v3
    if v3 {
        for fi45 := 0; fi45 < 5; fi45++ {
            for fi46 := 0; fi46 < 6; fi46++ {
                if (0 > v44) {
                    _ = len(v1)
                }
                v47 := v1["foo bar"]
                _ = fi46
                _ = v47
            }
            if false {
                var v48 int64 = 839
                for fi49 := 0; fi49 < 17; fi49++ {
                    switch fi49 {
                    case 128:
                        var v50 bool = true
                        _ = v50
                    }
                    switch (fi49 | fi45) {
                    case -231:
                        switch fi49 {
                        case -967:
                            v51 := GenS8{F0: genFn0(fi45), F1: 1000, F2: 155}
                            _ = v51
                        }
                    case 622:
                        v52 := GenS9{F0: 128, F1: 0}
                        _ = v52
                    default:
                        if v3 {
                            switch (-(fi45)) {
                            case -130:
                                var v53 int32 = int32(v48)
                                _ = v53
                            case 845:
                                var v54 int64 = (^(v48))
                                _ = v54
                            case -348:
                                v5 = 3.14159
                            }
                        }
                    }
                    _ = fi49
                }
                _ = v48
            }
            _ = fi45
        }
        v3 = v3
        if len(v2) > 0 {
            v2[0] = "a"
        }
    } else {
        if v3 {
            for fi55 := 0; fi55 < 1; fi55++ {
                v3 = ((!false) || ((1 <= int16(v44)) && v3))
                _ = fi55
            }
        }
        v44 = v44
    }
    _ = len(v1)
    v44 += (^(v44))
    var v56 uint8 = (^(v4))
    fmt.Println(v41, v56, v56)
    var v57 float32 = (v5 + v5)
    var v58 int64 = 1
    v59 := GenS3{F0: (v56 * (v4 ^ (v4 % (((^(v4))) | 1))))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v44
    _ = v5
    _ = v56
    _ = v57
    _ = v58
    _ = v59
}
