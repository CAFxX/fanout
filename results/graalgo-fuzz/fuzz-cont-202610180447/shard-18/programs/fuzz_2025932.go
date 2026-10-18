package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int16; F1 int32; F2 int32 }
type GenS3 struct { F0 int16; F1 int16; F2 bool }
type GenS4 struct { F0 bool; F1 uint }
type GenS5 struct { F0 int8; F1 string }
type GenS6 struct { F0 uint16; F1 int; F2 uint8 }
type GenS7 struct { F0 bool; F1 string; F2 int }
type GenS8 struct { F0 string; F1 bool }
func genFn0() string {
    var v4 string = "a"
    v4 = "a"
    v4 = "hello"
    fmt.Println(v4)
    for fi5 := 0; fi5 < 7; fi5++ {
        var v6 uint = 4294967295
        v6 += (v6 | 42)
        switch (^(fi5)) {
        case -773:
            v4 = ("" + (v4 + v4))
        default:
            v7 := []int{(fi5 % (((fi5 ^ -275)) | 1)), fi5, (fi5 % ((127) | 1))}
            _ = v7
        }
        _ = fi5
        _ = v6
    }
    _ = v4
    return "foo bar"
}
func genFn1() uint16 {
    var v8 uint16 = 864
    v9 := []int{-517, 42}
    _ = v8
    _ = v9
    return v8
}

func main() {
    var v1 float64 = (0.0 - 0.0)
    v2 := []string{("foo bar" + "foo bar")}
    var v3 int8 = 1
    switch "\n\t" {
    case "foo bar":
        switch 850 {
        case 215:
            if (195 > -1) {
                switch 32767 {
                case 42:
                    var v10 float32 = (((3.14159 + (3.14159 / (0.1 - 3.14159))) + 1e10) + (float32(v1) + (((3.14159 + 1.5) * 1.5) / (1.5 / 1.5))))
                    _ = v10
                case 575:
                    for fi11 := 0; fi11 < 10; fi11++ {
                        v3 = v3
                        switch fi11 {
                        case 128:
                            fmt.Println(fi11, fi11, fi11)
                        }
                        _ = fi11
                    }
                }
                var v12 bool = (1 != genFn1())
                for fi13 := 0; fi13 < 11; fi13++ {
                    v3 += (v3 << ((42) & 63))
                    var v14 int64 = 42
                    _ = fi13
                    _ = v14
                }
                _ = v12
            }
        case 937:
            v15 := GenS1{F0: 2, F1: 42, F2: 98}
            _ = v15
        default:
            switch ("foo bar" + "") {
            case "":
                var v16 string = "hello"
                _ = v16
            case "foo bar":
                var v17 uint = 0
                _ = v17
            }
        }
    }
    if (true || ((1 < 74) || false)) {
        if ((18446744073709551615 == 42) && (255 > 1)) {
            switch "foo bar" {
            case "hello":
                v3 -= (v3 | (-(v3)))
            case "foo bar":
                v18 := GenS1{F0: 1, F1: 0, F2: 2}
                _ = v18
            case "a":
                v19 := GenS1{F0: 1, F1: 446, F2: 127}
                _ = v19
            default:
                var v20 int = -584
                _ = v20
            }
        } else {
            v2 = append(v2, "a")
            if false {
                for fi21 := 0; fi21 < 7; fi21++ {
                    for fi22 := 0; fi22 < 8; fi22++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi22
                    }
                    switch 646 {
                    case -906:
                        v23 := GenS2{F0: 1000, F1: 2147483647, F2: 2147483647}
                        _ = v23
                    case 1:
                        v1 += v1
                    case -1:
                        v24 := GenS3{F0: -32768, F1: 1000, F2: (0.1 != (v1 + 0.0))}
                        _ = v24
                    default:
                        v25 := GenS1{F0: 18446744073709551615, F1: 715, F2: 2}
                        _ = v25
                    }
                    _ = fi21
                }
                fmt.Println(v1)
                v1 -= v1
            }
            v3 = v3
        }
        for fi26 := 0; fi26 < 16; fi26++ {
            var v27 int = (^(fi26))
            _ = len(v2)
            v28 := GenS2{F0: -1, F1: 2147483647, F2: -2147483648}
            _ = fi26
            _ = v27
            _ = v28
        }
    } else {
        fmt.Println(v1, v1)
    }
    switch (("a" + ("foo bar" + "")) + "a") {
    case "":
        if true {
            v29 := GenS1{F0: 1, F1: 42, F2: 2}
            _ = v29
        } else {
            for fi30 := 0; fi30 < 17; fi30++ {
                for fi31 := 0; fi31 < 9; fi31++ {
                    v32 := GenS4{F0: (32767 >= genFn1()), F1: 175}
                    if false {
                        if (v3 > 100) {
                            var v33 int16 = 0
                            if len(v2) > 0 {
                                v2[0] = ("" + "a")
                            }
                            var v34 uint = 65535
                            _ = v33
                            _ = v34
                        }
                        var v35 float64 = v1
                        _ = v35
                    }
                    _ = fi31
                    _ = v32
                }
                var v36 uint16 = 32767
                v37 := GenS2{F0: 988, F1: -1, F2: -1000}
                _ = fi30
                _ = v36
                _ = v37
            }
            var v38 uint64 = 582
            if (v1 == 0.1) {
                for fi39 := 0; fi39 < 10; fi39++ {
                    for fi40 := 0; fi40 < 20; fi40++ {
                        var v41 bool = false
                        _ = fi40
                        _ = v41
                    }
                    v3 = (^(v3))
                    v38 = v38
                    _ = fi39
                }
                if (4294967295 != 65535) {
                    switch "\n\t" {
                    case "":
                        switch "a" {
                        case "":
                            for fi42 := 0; fi42 < 8; fi42++ {
                                v38 = (v38 - (v38 | 18446744073709551615))
                                v43 := GenS2{F0: 32767, F1: 2147483647, F2: 0}
                                _ = fi42
                                _ = v43
                            }
                        case "hello":
                            var v44 int8 = 42
                            _ = v44
                        default:
                            v45 := GenS5{F0: (v3 * v3), F1: "\n\t"}
                            _ = v45
                        }
                    }
                    if false {
                        var v46 int = -2147483648
                        var v47 float32 = ((float32(((3.14159 * float64(0.0)) + float64((0.0 * 0.1)))) - ((float32(1e10) * (float32(1e10) - -2.25)) + 0.1)) - -2.25)
                        _ = len(v2)
                        _ = v46
                        _ = v47
                    }
                }
            } else {
                if (false || true) {
                    var v48 int16 = 42
                    for fi49 := 0; fi49 < 3; fi49++ {
                        if (!(!false)) {
                            v2 = append(v2, "foo bar")
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            fmt.Println(v3, v48, v3)
                        } else {
                            if (0 == 42) {
                                var v50 bool = (255 == 2)
                                _ = v50
                            }
                            for fi51 := 0; fi51 < 12; fi51++ {
                                switch fi51 {
                                case 0:
                                    var v52 int8 = v3
                                    _ = v52
                                case 2147483647:
                                    fmt.Println(v3)
                                }
                                _ = len(v2)
                                _ = fi51
                            }
                            if true {
                                v53 := GenS2{F0: v48, F1: 333, F2: -767}
                                v3 = (v3 >> ((136) & 63))
                                _ = v53
                            } else {
                                v1 = v1
                                v54 := GenS3{F0: 255, F1: v48, F2: ((v38 >> ((455) & 63)) < 42)}
                                _ = v54
                            }
                        }
                        var v55 bool = (v1 != v1)
                        _ = fi49
                        _ = v55
                    }
                    for fi56 := 0; fi56 < 9; fi56++ {
                        var v57 string = "hello"
                        _ = fi56
                        _ = v57
                    }
                    _ = v48
                }
                if (!false) {
                    var v58 int64 = 108
                    _ = v58
                }
            }
            _ = v38
        }
    case "\n\t":
        switch 9223372036854775807 {
        case -170:
            if (1 == 0) {
                v59 := GenS6{F0: genFn1(), F1: -34, F2: 13}
                var v60 float32 = float32(((-2.25 - float64((-2.25 * ((3.14159 * -2.25) + 0.0)))) - v1))
                switch "a" {
                case "":
                    v1 -= ((v1 + ((-2.25 + 0.0) * (v1 / v1))) + 1.5)
                }
                _ = v59
                _ = v60
            }
        case -2147483648:
            if true {
                var v61 int = 154
                _ = len(v2)
                if true {
                    switch -1 {
                    case 42:
                        v62 := GenS7{F0: (42 < 65535), F1: "\n\t", F2: (-(v61))}
                        _ = v62
                    case 32768:
                        v3 = (^(v3))
                    case 2147483647:
                        v63 := GenS6{F0: 65535, F1: v61, F2: 1}
                        _ = v63
                    }
                }
                _ = v61
            } else {
                if len(v2) > 0 {
                    v2[0] = "hello"
                }
                if true {
                    if true {
                        v64 := GenS8{F0: "foo bar", F1: true}
                        if len(v2) > 0 {
                            v2[0] = "a"
                        }
                        v2 = append(v2, "a")
                        _ = v64
                    } else {
                        if len(v2) > 0 {
                            v2[0] = "\n\t"
                        }
                    }
                } else {
                    v65 := GenS2{F0: 1000, F1: -782, F2: -2147483648}
                    v3 = (v3 >> ((0) & 63))
                    v66 := GenS1{F0: 4294967295, F1: 4294967295, F2: 127}
                    _ = v65
                    _ = v66
                }
            }
        case 32768:
            v3 += (v3 | v3)
        default:
            var v67 int16 = -381
            _ = v67
        }
    case "a":
        var v68 uint16 = 488
        _ = v68
    }
    if true {
        for fi69 := 0; fi69 < 2; fi69++ {
            v3 += v3
            var v70 float32 = float32(v1)
            _ = fi69
            _ = v70
        }
    }
    var v71 uint32 = 4294967295
    var v72 uint = 500
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    var v73 uint16 = 1
    v74 := GenS7{F0: (v72 < (v72 % (((v72 & v72)) | 1))), F1: ("" + "hello"), F2: 0}
    var v75 uint = v72
    v72 -= (v75 - uint(v73))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    fmt.Println("v73:", v73)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v71
    _ = v72
    _ = v73
    _ = v74
    _ = v75
}
