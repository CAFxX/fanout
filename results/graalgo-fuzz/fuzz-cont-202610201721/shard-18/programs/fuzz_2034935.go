package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 string; F2 int32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int32 }
func genFn0() uint16 {
    var v6 int = 88
    for fi7 := 0; fi7 < 6; fi7++ {
        v8 := GenS1{F0: 250, F1: 0}
        for fi9 := 0; fi9 < 19; fi9++ {
            fmt.Println(fi7, fi7)
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    switch v6 {
    case -146:
        for fi10 := 0; fi10 < 15; fi10++ {
            for fi11 := 0; fi11 < 11; fi11++ {
                fmt.Println(fi10)
                v12 := []int{}
                v6 += (fi10 | v6)
                if ((float32((float64(1e10) + 0.1)) + float32(1.5)) > 1.5) {
                    v6 = (-(fi10))
                    v13 := GenS1{F0: 127, F1: -294}
                    _ = v13
                }
                _ = fi11
                _ = v12
            }
            for fi14 := 0; fi14 < 20; fi14++ {
                fmt.Println(fi10, fi14, fi10)
                var v15 int = (v6 / ((fi10) | 1))
                var v16 uint16 = 1
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = fi10
        }
    case 279:
        v17 := GenS2{F0: (("foo bar" + ("a" + ((("" + "\n\t") + "foo bar") + "foo bar"))) + ("a" + ("" + (((("foo bar" + ((("hello" + "hello") + "a") + "")) + ("hello" + "")) + "hello") + ("hello" + "\n\t")))))}
        _ = v17
    }
    v6 -= 879
    var v18 uint32 = 42
    _ = v18
    _ = v6
    return uint16(v18)
}

func main() {
    var v1 int32 = 42
    var v2 uint32 = 2147483648
    v3 := []uint32{(v2 >> ((2) & 63)), v2}
    var v4 int32 = -1
    v5 := make(map[string]bool)
    fmt.Println(v4, v1)
    if false {
        var v19 int8 = 126
        _ = v19
    } else {
        switch int(v2) {
        case 0:
            switch "\n\t" {
            case "a":
                if true {
                    v2 = (v2 % (((v2 << ((65535) & 63))) | 1))
                } else {
                    var v20 bool = (genFn0() == 1)
                    switch (((("foo bar" + ("\n\t" + "a")) + ((("hello" + (("foo bar" + "\n\t") + "a")) + "hello") + "hello")) + ("a" + "foo bar")) + "a") {
                    case "\n\t":
                        _ = len(v5)
                    case "a":
                        var v21 uint16 = 362
                        _ = v21
                    case "hello":
                        for fi22 := 0; fi22 < 2; fi22++ {
                            v5["a"] = (("\n\t" + "") <= "hello")
                            v1 -= v1
                            v23 := GenS1{F0: 2, F1: 1}
                            _ = fi22
                            _ = v23
                        }
                    }
                    _ = v20
                }
            case "":
                v24 := GenS1{F0: 0, F1: 0}
                _ = v24
            case "\n\t":
                if true {
                    v1 = v1
                } else {
                    var v25 float32 = float32(0.1)
                    _ = len(v5)
                    _ = v25
                }
            default:
                for fi26 := 0; fi26 < 14; fi26++ {
                    var v27 int16 = 0
                    _ = fi26
                    _ = v27
                }
            }
        }
    }
    var v28 uint64 = 18446744073709551615
    var v29 int32 = v4
    delete(v5, "hello")
    fmt.Println(v2)
    var v30 uint32 = (v2 * v2)
    if false {
        v28 = 21
        v31 := GenS2{F0: "a"}
        for fi32 := 0; fi32 < 18; fi32++ {
            if true {
                if true {
                    if (!(int(v28) == fi32)) {
                        v28 = (v28 + 817)
                    } else {
                        v1 += (v1 * (v29 ^ (-(v1))))
                        var v33 bool = (!false)
                        _ = v33
                    }
                    for fi34 := 0; fi34 < 4; fi34++ {
                        v29 = v29
                        var v35 uint = 255
                        _ = fi34
                        _ = v35
                    }
                } else {
                    v30 = 65535
                    v2 = v2
                }
            } else {
                v29 += v1
                if false {
                    v4 += v4
                } else {
                    if true {
                        switch ("foo bar" + "a") {
                        case "":
                            v4 -= v4
                        default:
                            for fi36 := 0; fi36 < 6; fi36++ {
                                var v37 uint = 42
                                if false {
                                    if (!false) {
                                        var v38 uint64 = (v28 & uint64(v2))
                                        _ = v38
                                    }
                                    fmt.Println(v37, v1)
                                } else {
                                    v39 := GenS1{F0: (v28 * v28), F1: -1}
                                    if true {
                                        if (1.5 != ((float64(0.1) / 0.1) * 1e10)) {
                                            v29 -= 42
                                        }
                                        var v40 int16 = -32768
                                        v41 := GenS3{F0: v1, F1: (((("" + "a") + "hello") + "hello") + "a"), F2: (v1 + (v1 % ((-1) | 1)))}
                                        _ = v40
                                        _ = v41
                                    }
                                    var v42 int32 = v4
                                    if (1 != 128) {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        if len(v3) > 0 {
                                            v3[0] = v30
                                        }
                                    } else {
                                        if true {
                                            var v43 int32 = v1
                                            v44 := GenS1{F0: (v28 ^ 338), F1: 42}
                                            if ((true || false) || ((fi36 & fi36) >= (-(fi32)))) {
                                                var v45 int = (fi32 & fi32)
                                                fmt.Println(v2)
                                                _ = v45
                                            }
                                            var v46 int = fi32
                                            switch (("\n\t" + "hello") + (("hello" + "") + "foo bar")) {
                                            case "foo bar":
                                                var v47 int32 = (v43 % (((v43 ^ v1)) | 1))
                                                _ = v47
                                            }
                                            _ = v43
                                            _ = v44
                                            _ = v46
                                        } else {
                                            if (!(9223372036854775807 >= 422)) {
                                                v4 = v4
                                                fmt.Println(v1, v42)
                                                v42 = (v42 & (v4 >> ((1) & 63)))
                                            }
                                            var v48 int = 412
                                            v49 := GenS2{F0: ("hello" + (("" + ("\n\t" + "hello")) + "foo bar"))}
                                            v2 -= 0
                                            switch "a" {
                                            case "hello":
                                                var v50 int8 = 116
                                                _ = v50
                                            case "a":
                                                v51 := GenS3{F0: (-(v1)), F1: "foo bar", F2: -954}
                                                _ = v51
                                            case "foo bar":
                                                v52 := GenS2{F0: ""}
                                                _ = v52
                                            default:
                                                switch fi32 {
                                                case 1:
                                                    v53 := GenS4{F0: -128}
                                                    if true {
                                                        fmt.Println(fi32, v42)
                                                        var v54 float64 = float64((1.5 + (((1e10 / 0.1) * (0.1 + 3.14159)) * (1.5 + (3.14159 - 0.1)))))
                                                        _ = v54
                                                    }
                                                    v55 := GenS2{F0: "a"}
                                                    _ = v53
                                                    _ = v55
                                                default:
                                                    v29 = (v4 / (((v29 - (v1 - (^(v1))))) | 1))
                                                }
                                            }
                                            _ = v48
                                            _ = v49
                                        }
                                        delete(v5, "\n\t")
                                        v30 = 4294967295
                                    }
                                    _ = v39
                                    _ = v42
                                }
                                var v56 bool = false
                                _ = fi36
                                _ = v37
                                _ = v56
                            }
                        }
                    }
                    var v57 uint8 = 209
                    v29 = 1
                    _ = v57
                }
                for fi58 := 0; fi58 < 10; fi58++ {
                    switch "\n\t" {
                    case "hello":
                        if false {
                            var v59 bool = (v1 < 442)
                            switch "\n\t" {
                            case "a":
                                fmt.Println(v29, v59, v1)
                            case "\n\t":
                                var v60 float64 = ((((1.5 + ((0.0 - 0.1) + (float64(3.14159) / -2.25))) / -2.25) * float64(1.5)) - 0.1)
                                _ = v60
                            default:
                                if ((fi32 % ((fi58) | 1)) != (fi58 / (((fi58 & (fi58 + fi58))) | 1))) {
                                    if v59 {
                                        if v59 {
                                            if false {
                                                _ = len(v3)
                                                if (1e10 != 3.14159) {
                                                    if false {
                                                        v2 = v30
                                                        fmt.Println(v4)
                                                    }
                                                    var v61 uint32 = (^(v2))
                                                    fmt.Println(fi58)
                                                    _ = v61
                                                }
                                            }
                                        }
                                        v29 = (v1 + v29)
                                        var v62 float64 = -2.25
                                        fmt.Println(v28, v1, v4)
                                        switch "foo bar" {
                                        case "hello":
                                            if v59 {
                                                fmt.Println(v4, v2, v28)
                                                if true {
                                                    var v63 uint32 = (^(v30))
                                                    _ = v63
                                                } else {
                                                    switch "hello" {
                                                    case "hello":
                                                        v2 -= v30
                                                    case "foo bar":
                                                        var v64 int = fi58
                                                        _ = v64
                                                    case "a":
                                                        v1 += (v4 * (v29 | -503))
                                                    default:
                                                        v30 = 895
                                                    }
                                                }
                                            }
                                        default:
                                            v5["\n\t"] = (v59 && false)
                                        }
                                        _ = v62
                                    } else {
                                        fmt.Println(v29, fi58, v29)
                                        switch ("foo bar" + (((("" + "") + "a") + (("" + "hello") + "\n\t")) + "hello")) {
                                        case "":
                                            v28 = v28
                                        default:
                                            var v65 uint64 = v28
                                            _ = v65
                                        }
                                    }
                                    if (3.14159 > 0.1) {
                                        _ = len(v5)
                                    } else {
                                        _ = len(v3)
                                    }
                                    v66 := GenS3{F0: -1, F1: ("hello" + (("\n\t" + "a") + "")), F2: v4}
                                    _ = v66
                                }
                            }
                            if true {
                                switch (fi32 << ((1) & 63)) {
                                case 42:
                                    fmt.Println(v1, v59)
                                case 2147483647:
                                    v67 := GenS5{F0: v29}
                                    _ = v67
                                case 32768:
                                    v30 = v30
                                default:
                                    v5["hello"] = (!(303 > 0))
                                }
                                v68 := v5[""]
                                _ = v68
                            } else {
                                var v69 int64 = 1
                                _ = v69
                            }
                            _ = v59
                        }
                    case "\n\t":
                        var v70 int8 = -105
                        _ = v70
                    }
                    _ = fi58
                }
            }
            v71 := v5["\n\t"]
            for fi72 := 0; fi72 < 7; fi72++ {
                fmt.Println(v29, v28)
                v4 += v4
                var v73 int32 = (v1 ^ (v29 << ((1) & 63)))
                var v74 float32 = (((-2.25 / 1e10) / (float32((1e10 + 1e10)) - (3.14159 - 0.1))) * float32(float64(float32(1.5))))
                _ = fi72
                _ = v73
                _ = v74
            }
            _ = fi32
            _ = v71
        }
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
