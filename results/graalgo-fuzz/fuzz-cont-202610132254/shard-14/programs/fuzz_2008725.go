package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 uint16; F2 int }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 int8; F1 string; F2 int8 }
type GenS6 struct { F0 uint8; F1 string; F2 int64 }
type GenS7 struct { F0 int8; F1 int32; F2 int16 }
type GenS8 struct { F0 uint; F1 bool }
func genFn0(p0 uint, p1 string, p2 uint64) uint8 {
    p0 += (p0 % ((p0) | 1))
    for fi4 := 0; fi4 < 13; fi4++ {
        if (255 != 368) {
            fmt.Println(fi4, p2, fi4)
            v5 := []int{}
            _ = v5
        } else {
            if (!false) {
                v6 := []int{-2147483648, (fi4 | 1)}
                switch fi4 {
                case 9223372036854775807:
                    if true {
                        var v7 uint = 65535
                        _ = v7
                    }
                }
                _ = v6
            }
            p1 = (p1 + "")
            for fi8 := 0; fi8 < 10; fi8++ {
                v9 := GenS1{F0: -9223372036854775808}
                fmt.Println(fi8, p0, fi4)
                _ = fi8
                _ = v9
            }
        }
        var v10 uint = (^(p0))
        v10 = v10
        _ = fi4
        _ = v10
    }
    if true {
        for fi11 := 0; fi11 < 9; fi11++ {
            for fi12 := 0; fi12 < 18; fi12++ {
                p2 = p2
                var v13 bool = (!false)
                _ = fi12
                _ = v13
            }
            if true {
                fmt.Println(p2)
                var v14 string = p1
                _ = v14
            } else {
                p1 = ("foo bar" + ("a" + ("hello" + p1)))
            }
            var v15 uint16 = 0
            _ = fi11
            _ = v15
        }
    } else {
        v16 := []int{0, 873}
        var v17 uint64 = 807
        var v18 int64 = -425
        _ = v16
        _ = v17
        _ = v18
    }
    v19 := GenS1{F0: 42}
    var v20 bool = (!((0 < 8) || true))
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v20
    return 255
}
func genFn1(p0 int) string {
    v21 := GenS1{F0: 42}
    var v22 uint64 = 6
    p0 -= (p0 >> ((4294967295) & 63))
    v23 := GenS2{F0: true}
    p0 = p0
    _ = p0
    _ = v21
    _ = v22
    _ = v23
    return (("" + "") + "")
}

func main() {
    var v1 int8 = 1
    var v2 int = -9223372036854775808
    var v3 uint32 = 2147483648
    v3 = v3
    v24 := []int{(v2 - v2), 32767}
    v3 = 1
    switch (-(v2)) {
    case 32767:
        if ((((float64(1e10) * -2.25) - ((1e10 - -2.25) / (((float64(1e10) - 0.0) * 3.14159) / float64(1e10)))) - ((-2.25 - 0.1) + 1.5)) != ((0.0 + 1.5) - 0.0)) {
            v1 = (v1 / ((100) | 1))
        }
    case 1:
        var v25 int = (^(v2))
        _ = v25
    case -439:
        var v26 int64 = 1
        _ = v26
    }
    for fi27 := 0; fi27 < 8; fi27++ {
        if len(v24) >= 2 {
            v24 = v24[1:len(v24)]
        }
        _ = fi27
    }
    fmt.Println(v1, v3)
    v3 = (v3 ^ 0)
    for fi28 := 0; fi28 < 15; fi28++ {
        if true {
            v3 += v3
            _ = len(v24)
            fmt.Println(fi28, v3, v1)
        }
        for fi29 := 0; fi29 < 20; fi29++ {
            switch (fi29 ^ v2) {
            case 0:
                v2 = 32767
            case -776:
                switch (v2 / (((v2 * 127)) | 1)) {
                case 931:
                    var v30 bool = (!true)
                    var v31 int64 = 0
                    var v32 uint = 42
                    _ = v30
                    _ = v31
                    _ = v32
                case 722:
                    v2 = (-(fi29))
                case 9223372036854775807:
                    v2 = -176
                    if (1000 >= 32767) {
                        v33 := GenS2{F0: ("\n\t" == "\n\t")}
                        v24 = append(v24, 2147483647)
                        v34 := GenS3{F0: genFn0(4294967295, ("a" + ("hello" + ((("\n\t" + ("" + "foo bar")) + ("" + "hello")) + (("\n\t" + "a") + ("\n\t" + "hello"))))), 1), F1: 65535, F2: (-(v2))}
                        _ = v33
                        _ = v34
                    }
                    fmt.Println(fi28)
                default:
                    var v35 string = (("hello" + (("" + "hello") + "hello")) + (("\n\t" + "foo bar") + ("hello" + "hello")))
                    _ = v35
                }
            }
            _ = fi29
        }
        var v36 int64 = -1
        _ = fi28
        _ = v36
    }
    if false {
        v3 = 2
        v37 := GenS3{F0: 42, F1: 32767, F2: v2}
        var v38 uint8 = 128
        _ = v37
        _ = v38
    } else {
        _ = len(v24)
        fmt.Println(v1, v2)
        switch "" {
        case "a":
            v24 = append(v24, (v2 >> ((670) & 63)))
        case "\n\t":
            fmt.Println(v1, v2, v2)
        case "":
            var v39 uint = 42
            _ = v39
        default:
            fmt.Println(v3, v2)
        }
    }
    var v40 uint16 = 255
    v41 := GenS1{F0: -1}
    if false {
        var v42 uint32 = (v3 - 42)
        v43 := GenS4{F0: "a", F1: genFn0(2, "foo bar", 18446744073709551615)}
        var v44 uint16 = (^(v40))
        _ = v42
        _ = v43
        _ = v44
    } else {
        for fi45 := 0; fi45 < 14; fi45++ {
            switch "a" {
            case "":
                v40 = (^(v40))
            }
            if ((((("\n\t" + "foo bar") + ("\n\t" + "\n\t")) + "") + "\n\t") >= (("a" + ("hello" + "foo bar")) + "\n\t")) {
                for fi46 := 0; fi46 < 14; fi46++ {
                    switch "a" {
                    case "hello":
                        var v47 float32 = (0.1 * -2.25)
                        _ = v47
                    case "\n\t":
                        var v48 uint16 = (v40 % (((v40 ^ (v40 % (((v40 ^ (v40 * v40))) | 1)))) | 1))
                        _ = v48
                    case "":
                        var v49 int16 = 0
                        _ = v49
                    default:
                        var v50 int64 = 1
                        _ = v50
                    }
                    _ = fi46
                }
            } else {
                if len(v24) > 0 {
                    v24[0] = 128
                }
                for fi51 := 0; fi51 < 3; fi51++ {
                    v40 -= v40
                    fmt.Println(fi45)
                    _ = fi51
                }
            }
            _ = fi45
        }
        for fi52 := 0; fi52 < 8; fi52++ {
            if (-2147483648 == 1) {
                v53 := GenS5{F0: (v1 >> ((4294967295) & 63)), F1: ("\n\t" + (((("a" + "") + ("hello" + "foo bar")) + ("a" + "")) + "\n\t")), F2: v1}
                switch "\n\t" {
                case "hello":
                    v2 = -1
                case "\n\t":
                    fmt.Println(v2)
                default:
                    var v54 int64 = 9223372036854775807
                    _ = v54
                }
                _ = v53
            } else {
                if false {
                    var v55 int16 = 783
                    _ = v55
                } else {
                    for fi56 := 0; fi56 < 10; fi56++ {
                        switch 32768 {
                        case -44:
                            switch "hello" {
                            case "foo bar":
                                if len(v24) >= 2 {
                                    v24 = v24[1:len(v24)]
                                }
                                var v57 uint8 = 255
                                fmt.Println(fi56, v1, v3)
                                _ = v57
                            case "a":
                                v58 := GenS6{F0: 0, F1: (("hello" + ("\n\t" + "hello")) + "\n\t"), F2: 42}
                                _ = v58
                            default:
                                switch "\n\t" {
                                case "foo bar":
                                    var v59 float32 = -2.25
                                    _ = v59
                                default:
                                    var v60 uint = 2
                                    _ = v60
                                }
                            }
                        case 32768:
                            v40 = 335
                        case 9223372036854775807:
                            v40 += v40
                        }
                        _ = fi56
                    }
                }
                v3 = (v3 + v3)
            }
            if len(v24) > 0 {
                v24[0] = (v2 % ((42) | 1))
            }
            if len(v24) > 0 {
                v24[0] = 32767
            }
            _ = fi52
        }
        var v61 int16 = 1
        _ = v61
    }
    var v62 float32 = 1e10
    fmt.Println(v2, v2)
    if false {
        for fi63 := 0; fi63 < 13; fi63++ {
            v64 := GenS7{F0: 25, F1: 549, F2: 764}
            _ = fi63
            _ = v64
        }
        if true {
            v65 := GenS3{F0: 1, F1: v40, F2: (^(v2))}
            switch (v2 >> ((4294967295) & 63)) {
            case 718:
                v66 := GenS8{F0: 65535, F1: false}
                _ = v66
            case -306:
                var v67 bool = (1e10 <= -2.25)
                _ = v67
            default:
                var v68 int = (-(v2))
                _ = v68
            }
            v62 -= (1.5 / v62)
            _ = v65
        }
    }
    var v69 int16 = -1
    for fi70 := 0; fi70 < 3; fi70++ {
        if (18446744073709551615 > 2) {
            if len(v24) >= 2 {
                v24 = v24[1:len(v24)]
            }
        } else {
            for fi71 := 0; fi71 < 3; fi71++ {
                v62 = v62
                var v72 bool = (false || true)
                _ = fi71
                _ = v72
            }
            v3 = (^(v3))
        }
        if ("foo bar" == ("hello" + "")) {
            _ = len(v24)
        }
        var v73 int64 = 0
        _ = fi70
        _ = v73
    }
    for fi74 := 0; fi74 < 17; fi74++ {
        _ = len(v24)
        _ = fi74
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v62:", v62)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v40
    _ = v41
    _ = v62
    _ = v69
}
