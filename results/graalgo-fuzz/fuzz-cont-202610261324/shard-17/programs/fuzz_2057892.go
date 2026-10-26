package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 int16 }
type GenS2 struct { F0 int8; F1 uint8; F2 int64 }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 bool; F1 uint }
type GenS7 struct { F0 int16; F1 uint64 }
func genFn0() bool {
    switch "hello" {
    case "\n\t":
        var v6 int = -9223372036854775808
        _ = v6
    default:
        var v7 float32 = ((0.0 - 0.1) + 0.0)
        _ = v7
    }
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 uint8 = 205
        var v10 uint = 483
        if false {
            var v11 float64 = 0.0
            _ = v11
        } else {
            for fi12 := 0; fi12 < 11; fi12++ {
                var v13 uint32 = 2147483648
                _ = fi12
                _ = v13
            }
            var v14 uint = (v10 ^ v10)
            _ = v14
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    var v15 float64 = 1.5
    _ = v15
    return (2 != 4294967295)
}
func genFn1() uint16 {
    var v16 uint16 = 32767
    v16 += (v16 << ((18446744073709551615) & 63))
    v17 := GenS1{F0: true, F1: 0, F2: -1}
    _ = v16
    _ = v17
    return (^(v16))
}

func main() {
    v1 := []int{0, 949, 9223372036854775807}
    var v2 int8 = 0
    var v3 float64 = 1e10
    v4 := []uint32{4294967295}
    v5 := []string{"foo bar", "hello", "\n\t", "a"}
    for fi18 := 0; fi18 < 11; fi18++ {
        v19 := GenS2{F0: 0, F1: 42, F2: 42}
        if len(v4) > 0 {
            v4[0] = 0
        }
        switch "" {
        case "a":
            fmt.Println(v3, fi18)
        case "foo bar":
            for fi20 := 0; fi20 < 15; fi20++ {
                if true {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if false {
                        var v21 int8 = v2
                        _ = v21
                    } else {
                        fmt.Println(fi18, v2)
                    }
                    var v22 uint8 = 0
                    _ = v22
                } else {
                    v3 = v3
                    v3 = ((((v3 * v3) * -2.25) * v3) + v3)
                }
                v23 := GenS1{F0: (!(140 < 65535)), F1: 2, F2: 42}
                fmt.Println(v3, v3, fi20)
                _ = fi20
                _ = v23
            }
        case "":
            if ((65535 == 921) || false) {
                var v24 float32 = ((float32(((v3 + 0.0) - (v3 / float64(1e10)))) + (1e10 / 0.1)) + 0.1)
                var v25 uint32 = 1
                _ = v24
                _ = v25
            } else {
                switch (("hello" + "hello") + "a") {
                case "a":
                    switch fi18 {
                    case -40:
                        if true {
                            if true {
                                for fi26 := 0; fi26 < 1; fi26++ {
                                    fmt.Println(v3, v3)
                                    switch "a" {
                                    case "a":
                                        v2 = -48
                                    case "foo bar":
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    }
                                    _ = len(v5)
                                    _ = fi26
                                }
                                for fi27 := 0; fi27 < 2; fi27++ {
                                    v4 = append(v4, 2147483648)
                                    _ = fi27
                                }
                                v28 := GenS3{F0: 236, F1: (true || true)}
                                _ = v28
                            } else {
                                v3 = -2.25
                            }
                        }
                    }
                case "foo bar":
                    for fi29 := 0; fi29 < 9; fi29++ {
                        switch "foo bar" {
                        case "hello":
                            _ = len(v5)
                        case "a":
                            fmt.Println(v3)
                        }
                        if len(v5) > 0 {
                            v5[0] = ((("" + "\n\t") + "\n\t") + "foo bar")
                        }
                        v30 := GenS1{F0: false, F1: 140, F2: 0}
                        var v31 string = (("" + "\n\t") + (("foo bar" + "") + "hello"))
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                case "":
                    v32 := GenS4{F0: 2}
                    _ = v32
                default:
                    var v33 int = fi18
                    _ = v33
                }
            }
        }
        _ = fi18
        _ = v19
    }
    v34 := GenS4{F0: 392}
    switch -629 {
    case -317:
        if false {
            v3 = 0.0
        }
    }
    for fi35 := 0; fi35 < 17; fi35++ {
        v3 += 3.14159
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v36 := GenS5{F0: 42}
        _ = fi35
        _ = v36
    }
    var v37 int64 = 9223372036854775807
    if true {
        v38 := GenS6{F0: true, F1: 403}
        v3 = (((1e10 + v3) / (float64(0.0) + (v3 / 1.5))) + float64((-2.25 / (0.1 + 1.5))))
        _ = v38
    } else {
        if true {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v39 := GenS4{F0: 489}
            var v40 int8 = v2
            _ = v39
            _ = v40
        } else {
            fmt.Println(v2, v2, v37)
            if (v3 <= v3) {
                for fi41 := 0; fi41 < 4; fi41++ {
                    if (true || (true || (!(!false)))) {
                        v37 = 0
                        switch fi41 {
                        case -81:
                            v42 := GenS4{F0: 683}
                            _ = v42
                        case -954:
                            v43 := GenS7{F0: 1000, F1: 0}
                            _ = v43
                        case -2147483648:
                            if (2 > 609) {
                                if (4294967295 < 2) {
                                    var v44 uint32 = 2147483648
                                    fmt.Println(v44)
                                    var v45 int32 = -1
                                    _ = v44
                                    _ = v45
                                }
                                var v46 int32 = -2147483648
                                for fi47 := 0; fi47 < 11; fi47++ {
                                    v4 = append(v4, 42)
                                    _ = fi47
                                }
                                _ = v46
                            } else {
                                if (127 > 200) {
                                    var v48 int32 = -1
                                    v49 := GenS4{F0: 63}
                                    switch (fi41 * (-(fi41))) {
                                    case 128:
                                        if true {
                                            v1 = append(v1, 831)
                                        }
                                    case 5:
                                        if len(v4) > 0 {
                                            v4[0] = 65535
                                        }
                                    default:
                                        var v50 int32 = 1
                                        _ = v50
                                    }
                                    _ = v48
                                    _ = v49
                                } else {
                                    for fi51 := 0; fi51 < 12; fi51++ {
                                        fmt.Println(v3)
                                        fmt.Println(v3, fi41, v3)
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        _ = fi51
                                    }
                                    if ((fi41 & (^(fi41))) >= fi41) {
                                        fmt.Println(fi41, v37)
                                    } else {
                                        var v52 uint16 = 851
                                        var v53 float32 = -2.25
                                        if false {
                                            var v54 int8 = -128
                                            v55 := GenS1{F0: true, F1: 42, F2: 0}
                                            v54 += -100
                                            _ = v54
                                            _ = v55
                                        } else {
                                            v56 := GenS4{F0: (^(v52))}
                                            var v57 float64 = -2.25
                                            _ = v56
                                            _ = v57
                                        }
                                        _ = v52
                                        _ = v53
                                    }
                                }
                                switch "hello" {
                                case "a":
                                    v58 := GenS5{F0: 2147483647}
                                    _ = v58
                                }
                                if false {
                                    v2 += 127
                                } else {
                                    v37 += v37
                                }
                            }
                        default:
                            v37 = v37
                        }
                    }
                    _ = fi41
                }
            }
        }
    }
    var v59 uint = 42
    var v60 string = ("\n\t" + (("" + ("a" + ("hello" + "foo bar"))) + ("\n\t" + "\n\t")))
    var v61 uint = 553
    v62 := GenS1{F0: false, F1: 2, F2: 1000}
    var v63 float64 = v3
    if true {
        v5 = append(v5, v60)
        var v64 int64 = -1
        fmt.Println(v60, v37)
        _ = v64
    } else {
        v2 -= v2
        if true {
            var v65 int16 = 1
            var v66 int16 = (v65 & -1)
            if true {
                v60 = v60
            }
            _ = v65
            _ = v66
        } else {
            v61 = v59
            fmt.Println(v61)
        }
        v60 = (("hello" + v60) + v60)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v37
    _ = v4
    _ = v5
    _ = v59
    _ = v60
    _ = v61
    _ = v62
    _ = v63
}
