package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 int16; F1 int32; F2 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8; F1 int16; F2 uint8 }
func genFn0(p0 uint64, p1 uint8) int16 {
    p1 += p1
    v4 := []int{}
    p1 -= (p1 & 1)
    _ = p0
    _ = p1
    _ = v4
    return -279
}

func main() {
    var v1 int16 = -660
    var v2 uint32 = 2147483648
    var v3 uint64 = 175
    var v5 uint32 = (v2 << ((777) & 63))
    var v6 int64 = 42
    if false {
        var v7 uint16 = 255
        var v8 int = int(v3)
        _ = v7
        _ = v8
    }
    var v9 uint64 = (v3 << ((4294967295) & 63))
    if true {
        for fi10 := 0; fi10 < 12; fi10++ {
            var v11 int32 = 0
            for fi12 := 0; fi12 < 17; fi12++ {
                fmt.Println(fi12, v2, fi10)
                _ = fi12
            }
            var v13 int32 = v11
            _ = fi10
            _ = v11
            _ = v13
        }
        v5 = (v2 ^ 2147483648)
        v14 := []int{354}
        _ = v14
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        var v16 uint8 = 2
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 16; fi17++ {
        for fi18 := 0; fi18 < 6; fi18++ {
            v19 := GenS1{F0: v3}
            _ = fi18
            _ = v19
        }
        v20 := []int{fi17}
        _ = fi17
        _ = v20
    }
    v21 := []int{}
    v3 = (v9 - 42)
    v22 := GenS1{F0: (v3 - (v3 << ((1) & 63)))}
    fmt.Println(v2, v3, v2)
    v6 = v6
    fmt.Println(v9, v3)
    if false {
        if false {
            if true {
                v9 += (^(v3))
                if len(v21) > 0 {
                    v21[0] = -1
                }
                for fi23 := 0; fi23 < 2; fi23++ {
                    var v24 int16 = genFn0((v3 - (v3 * v3)), 1)
                    v24 = (-(v1))
                    _ = fi23
                    _ = v24
                }
            } else {
                switch ("hello" + "foo bar") {
                case "hello":
                    v2 += v5
                case "a":
                    if true {
                        switch "hello" {
                        case "foo bar":
                            v2 -= v5
                        case "":
                            if len(v21) > 0 {
                                v21[0] = 42
                            }
                        case "hello":
                            for fi25 := 0; fi25 < 18; fi25++ {
                                if len(v21) >= 2 {
                                    v21 = v21[1:len(v21)]
                                }
                                fmt.Println(v1, v5)
                                for fi26 := 0; fi26 < 2; fi26++ {
                                    switch "\n\t" {
                                    case "a":
                                        var v27 int = fi26
                                        _ = v27
                                    case "\n\t":
                                        switch fi25 {
                                        case -869:
                                            if true {
                                                if len(v21) >= 2 {
                                                    v21 = v21[1:len(v21)]
                                                }
                                                v2 = (^(v5))
                                                v28 := GenS2{F0: (v1 ^ v1), F1: "\n\t"}
                                                switch (fi26 ^ 4294967295) {
                                                case 1:
                                                    v29 := GenS3{F0: (v1 % ((357) | 1)), F1: 0, F2: "a"}
                                                    _ = v29
                                                case -643:
                                                    v5 = 388
                                                }
                                                v30 := GenS3{F0: 1000, F1: -298, F2: ""}
                                                _ = v28
                                                _ = v30
                                            } else {
                                                if false {
                                                    v31 := GenS3{F0: genFn0(v9, 255), F1: 1, F2: ("" + ("a" + ""))}
                                                    if ((v1 & v1) != 1000) {
                                                        if len(v21) >= 2 {
                                                            v21 = v21[1:len(v21)]
                                                        }
                                                    } else {
                                                        var v32 int32 = 2147483647
                                                        if true {
                                                            v6 -= v6
                                                        }
                                                        v33 := GenS4{F0: (v6 % ((1) | 1))}
                                                        _ = v32
                                                        _ = v33
                                                    }
                                                    if len(v21) > 0 {
                                                        v21[0] = fi26
                                                    }
                                                    _ = v31
                                                } else {
                                                    switch ("" + "foo bar") {
                                                    case "a":
                                                        var v34 int32 = -1
                                                        _ = v34
                                                    case "hello":
                                                        if false {
                                                            v3 += 601
                                                            if len(v21) > 0 {
                                                                v21[0] = fi26
                                                            }
                                                            var v35 int32 = 42
                                                            _ = v35
                                                        }
                                                    case "":
                                                        var v36 uint = 848
                                                        _ = v36
                                                    }
                                                    v21 = append(v21, int(v6))
                                                }
                                                fmt.Println(v2, v5, v2)
                                                var v37 int64 = v6
                                                _ = v37
                                            }
                                        case -11:
                                            var v38 int = 42
                                            _ = v38
                                        case 32767:
                                            v3 = v3
                                        default:
                                            v6 = (v6 ^ -1)
                                        }
                                    default:
                                        var v39 uint16 = 611
                                        _ = v39
                                    }
                                    var v40 uint64 = 5
                                    v41 := GenS3{F0: genFn0(v9, 128), F1: 2147483647, F2: ("foo bar" + ((((("foo bar" + ("hello" + "a")) + ("foo bar" + "hello")) + ((("hello" + "") + (("foo bar" + "hello") + ("a" + "a"))) + "\n\t")) + "") + "a"))}
                                    _ = fi26
                                    _ = v40
                                    _ = v41
                                }
                                _ = fi25
                            }
                        default:
                            var v42 float32 = 1e10
                            _ = v42
                        }
                        if (!false) {
                            if false {
                                v1 = v1
                                if (!true) {
                                    switch ("foo bar" + (((("\n\t" + "\n\t") + "") + (("a" + "foo bar") + "")) + (("\n\t" + "foo bar") + ("\n\t" + "")))) {
                                    case "\n\t":
                                        if false {
                                            if true {
                                                fmt.Println(v1, v3, v5)
                                            }
                                        }
                                    default:
                                        for fi43 := 0; fi43 < 18; fi43++ {
                                            for fi44 := 0; fi44 < 3; fi44++ {
                                                switch 1 {
                                                case 32767:
                                                    v1 -= (v1 | 32767)
                                                default:
                                                    if false {
                                                        var v45 int16 = 522
                                                        var v46 int = -9223372036854775808
                                                        v6 += v6
                                                        v47 := GenS1{F0: (v9 - (v9 * v9))}
                                                        _ = v45
                                                        _ = v46
                                                        _ = v47
                                                    } else {
                                                        var v48 int8 = 1
                                                        fmt.Println(v2)
                                                        var v49 int16 = v1
                                                        _ = v48
                                                        _ = v49
                                                    }
                                                }
                                                var v50 uint16 = 377
                                                _ = fi44
                                                _ = v50
                                            }
                                            _ = fi43
                                        }
                                    }
                                    v51 := GenS1{F0: v3}
                                    _ = v51
                                } else {
                                    fmt.Println(v9, v5)
                                    if len(v21) > 0 {
                                        v21[0] = 9223372036854775807
                                    }
                                }
                                switch "\n\t" {
                                case "hello":
                                    if false {
                                        fmt.Println(v3)
                                        if false {
                                            v9 = (v3 | 2)
                                            var v52 bool = false
                                            v21 = append(v21, 828)
                                            _ = v52
                                        }
                                    } else {
                                        for fi53 := 0; fi53 < 9; fi53++ {
                                            if ((((("foo bar" < "a") || true) || (9223372036854775807 > v6)) || true) || false) {
                                                fmt.Println(v1, fi53)
                                                var v54 int64 = (v6 / ((474) | 1))
                                                if (255 != 1) {
                                                    switch -315 {
                                                    case 4294967295:
                                                        var v55 uint8 = 1
                                                        _ = v55
                                                    case 32767:
                                                        v56 := GenS1{F0: (v9 * v9)}
                                                        _ = v56
                                                    }
                                                    fmt.Println(v1, v1)
                                                }
                                                _ = v54
                                            }
                                            v3 += (v3 % ((90) | 1))
                                            _ = fi53
                                        }
                                        v5 = v2
                                        fmt.Println(v9, v1, v9)
                                    }
                                default:
                                    for fi57 := 0; fi57 < 6; fi57++ {
                                        var v58 float64 = 1.5
                                        _ = fi57
                                        _ = v58
                                    }
                                }
                            } else {
                                if (true || ((0.0 == (float32(-2.25) / (((0.1 + -2.25) + (1.5 / (-2.25 * (0.1 + (((0.1 * (1.5 / (0.0 + 1.5))) / float32(((float64(float32((0.1 / -2.25))) - ((-2.25 - 1.5) + 1.5)) + ((-2.25 + 1e10) / float64((-2.25 / float32(3.14159))))))) - (-2.25 - (3.14159 - (0.0 * 3.14159)))))))) + 1.5))) && true)) {
                                    v6 = v6
                                }
                                v9 -= v9
                            }
                        } else {
                            switch ("foo bar" + "hello") {
                            case "a":
                                v2 += v5
                            case "\n\t":
                                v2 += 65535
                            case "foo bar":
                                switch (("a" + "\n\t") + "\n\t") {
                                case "a":
                                    if len(v21) >= 2 {
                                        v21 = v21[1:len(v21)]
                                    }
                                }
                            default:
                                for fi59 := 0; fi59 < 8; fi59++ {
                                    var v60 string = "a"
                                    _ = fi59
                                    _ = v60
                                }
                            }
                            v61 := GenS2{F0: -343, F1: ""}
                            var v62 int32 = -691
                            _ = v61
                            _ = v62
                        }
                        var v63 uint16 = 32767
                        _ = v63
                    }
                case "foo bar":
                    v1 += genFn0(0, 178)
                }
            }
            var v64 int64 = (v6 * (v6 ^ (^(v6))))
            v1 += v1
            _ = v64
        } else {
            for fi65 := 0; fi65 < 20; fi65++ {
                for fi66 := 0; fi66 < 19; fi66++ {
                    switch "hello" {
                    case "foo bar":
                        v9 -= uint64(v5)
                    case "a":
                        var v67 float64 = (1e10 * 0.1)
                        _ = v67
                    case "":
                        v2 = (^(v5))
                    default:
                        v68 := GenS3{F0: -1, F1: 812, F2: (("a" + "\n\t") + "\n\t")}
                        _ = v68
                    }
                    switch ("\n\t" + ((("foo bar" + "hello") + "\n\t") + "\n\t")) {
                    case "hello":
                        v69 := GenS3{F0: (v1 << ((255) & 63)), F1: -644, F2: "\n\t"}
                        _ = v69
                    case "foo bar":
                        v70 := GenS5{F0: 127, F1: genFn0(1, 45), F2: 71}
                        _ = v70
                    default:
                        v71 := GenS5{F0: 128, F1: genFn0(v9, 128), F2: 2}
                        _ = v71
                    }
                    _ = fi66
                }
                var v72 uint = 18446744073709551615
                _ = fi65
                _ = v72
            }
            v6 = v6
            if len(v21) >= 2 {
                v21 = v21[1:len(v21)]
            }
        }
        v6 = v6
    } else {
        v6 += v6
        v73 := GenS5{F0: 42, F1: 42, F2: 0}
        if len(v21) > 0 {
            v21[0] = 128
        }
        _ = v73
    }
    switch 32767 {
    case 690:
        v9 = (v3 * (v9 / (((v9 + 2)) | 1)))
    case 675:
        v21 = append(v21, 666)
    case 127:
        if len(v21) > 0 {
            v21[0] = 4294967295
        }
    }
    switch "foo bar" {
    case "":
        v21 = append(v21, -1)
    case "hello":
        fmt.Println(v9, v1, v6)
    case "a":
        if (42 == -128) {
            v74 := GenS5{F0: 128, F1: v1, F2: 42}
            switch "hello" {
            case "\n\t":
                var v75 int16 = v1
                _ = v75
            case "hello":
                for fi76 := 0; fi76 < 5; fi76++ {
                    if true {
                        for fi77 := 0; fi77 < 20; fi77++ {
                            var v78 int8 = 127
                            var v79 float64 = 1.5
                            var v80 int = (^(fi76))
                            v81 := GenS4{F0: (v6 / ((v6) | 1))}
                            _ = fi77
                            _ = v78
                            _ = v79
                            _ = v80
                            _ = v81
                        }
                    }
                    _ = fi76
                }
            }
            _ = v74
        }
    }
    fmt.Println(v9, v3, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v5
    _ = v6
    _ = v9
}
