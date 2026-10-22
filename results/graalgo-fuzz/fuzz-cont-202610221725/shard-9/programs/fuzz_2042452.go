package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 uint64; F2 uint64 }
func genFn0() uint8 {
    if false {
        var v4 int = 9223372036854775807
        fmt.Println(v4)
        _ = v4
    } else {
        switch 9223372036854775807 {
        case 9223372036854775807:
            for fi5 := 0; fi5 < 4; fi5++ {
                var v6 uint8 = 42
                switch fi5 {
                case -800:
                    if false {
                        var v7 int16 = -556
                        v7 -= 0
                        var v8 int16 = 42
                        _ = v7
                        _ = v8
                    } else {
                        v6 = v6
                        v9 := GenS1{F0: 42}
                        _ = v9
                    }
                default:
                    var v10 int64 = 42
                    _ = v10
                }
                switch ("hello" + "") {
                case "hello":
                    fmt.Println(v6)
                case "foo bar":
                    for fi11 := 0; fi11 < 4; fi11++ {
                        var v12 int16 = -1
                        _ = fi11
                        _ = v12
                    }
                default:
                    v13 := []int{2147483647, (-(fi5)), 42}
                    _ = v13
                }
                _ = fi5
                _ = v6
            }
        default:
            switch "\n\t" {
            case "\n\t":
                if true {
                    if false {
                        v14 := []int{}
                        for fi15 := 0; fi15 < 16; fi15++ {
                            v16 := GenS1{F0: 2}
                            _ = fi15
                            _ = v16
                        }
                        if true {
                            var v17 float64 = ((1.5 - 1e10) + 0.0)
                            if false {
                                v17 = 1.5
                            } else {
                                if true {
                                    var v18 int = 51
                                    _ = v18
                                }
                                for fi19 := 0; fi19 < 5; fi19++ {
                                    if false {
                                        var v20 int32 = 2147483647
                                        _ = v20
                                    }
                                    var v21 bool = (32767 != 32767)
                                    _ = fi19
                                    _ = v21
                                }
                                _ = len(v14)
                            }
                            for fi22 := 0; fi22 < 18; fi22++ {
                                v23 := GenS2{F0: 81}
                                v17 -= 1e10
                                var v24 uint8 = 255
                                _ = fi22
                                _ = v23
                                _ = v24
                            }
                            _ = v17
                        } else {
                            if false {
                                var v25 uint = 42
                                _ = v25
                            }
                            v14 = append(v14, 32767)
                        }
                        _ = v14
                    }
                }
            case "a":
                for fi26 := 0; fi26 < 12; fi26++ {
                    if true {
                        var v27 float64 = 1e10
                        _ = v27
                    }
                    var v28 int = (-(fi26))
                    _ = fi26
                    _ = v28
                }
            case "hello":
                v29 := []int{}
                _ = v29
            default:
                var v30 int8 = 0
                _ = v30
            }
        }
        var v31 int16 = 0
        v31 -= v31
        _ = v31
    }
    switch 2147483647 {
    case 9223372036854775807:
        for fi32 := 0; fi32 < 16; fi32++ {
            var v33 float64 = 0.1
            _ = fi32
            _ = v33
        }
    default:
        var v34 uint = 654
        _ = v34
    }
    return 0
}
func genFn1(p0 uint16) int64 {
    p0 = p0
    for fi35 := 0; fi35 < 17; fi35++ {
        fmt.Println(fi35)
        _ = fi35
    }
    for fi36 := 0; fi36 < 4; fi36++ {
        if false {
            for fi37 := 0; fi37 < 14; fi37++ {
                var v38 uint = 65535
                var v39 int16 = -1
                v40 := GenS2{F0: 840}
                if false {
                    v41 := []int{(fi36 >> (((^(v38))) & 63)), 4294967295, (^(fi37)), fi37}
                    _ = len(v41)
                    v39 += v39
                    p0 -= p0
                    fmt.Println(p0)
                    var v42 int = (^(fi36))
                    _ = v41
                    _ = v42
                } else {
                    var v43 int8 = -99
                    var v44 int64 = 1
                    _ = v43
                    _ = v44
                }
                _ = fi37
                _ = v38
                _ = v39
                _ = v40
            }
        }
        _ = fi36
    }
    _ = p0
    return 0
}

func main() {
    var v1 float32 = float32(1e10)
    v2 := make(map[uint32]int)
    var v3 bool = (!true)
    if v3 {
        for fi45 := 0; fi45 < 7; fi45++ {
            if (-1 < 32767) {
                v46 := GenS3{F0: 1, F1: 18446744073709551615, F2: 0}
                var v47 int16 = -924
                v1 -= v1
                _ = v46
                _ = v47
            } else {
                fmt.Println(v3, v1)
                var v48 bool = (!(255 == 0))
                _ = len(v2)
                _ = v48
            }
            v49 := GenS3{F0: 100, F1: 0, F2: 0}
            _ = fi45
            _ = v49
        }
        var v50 uint32 = 65535
        var v51 uint = 42
        _ = v50
        _ = v51
    } else {
        for fi52 := 0; fi52 < 12; fi52++ {
            v53 := GenS2{F0: 2}
            fmt.Println(v3, v3, v1)
            _ = fi52
            _ = v53
        }
        if (v3 && v3) {
            v3 = (42 > 880)
            v3 = v3
        }
    }
    v3 = ((!("a" <= "foo bar")) || v3)
    var v54 float32 = (v1 - v1)
    for fi55 := 0; fi55 < 6; fi55++ {
        fmt.Println(fi55, v3)
        var v56 int64 = genFn1(0)
        v57 := v2[348]
        _ = fi55
        _ = v56
        _ = v57
    }
    if (v54 == (v1 * v54)) {
        var v58 int8 = -128
        var v59 uint16 = 1
        switch (("" + "foo bar") + (("foo bar" + ("" + ((("hello" + "\n\t") + "a") + "\n\t"))) + "foo bar")) {
        case "":
            v60 := v2[979]
            _ = v60
        case "hello":
            var v61 bool = false
            _ = v61
        }
        _ = v58
        _ = v59
    }
    v62 := GenS2{F0: 97}
    if (524 == 42) {
        for fi63 := 0; fi63 < 4; fi63++ {
            var v64 uint32 = 803
            if v3 {
                var v65 uint8 = 2
                _ = v65
            } else {
                for fi66 := 0; fi66 < 15; fi66++ {
                    v54 += (v1 * v1)
                    _ = fi66
                }
                v64 = (v64 % (((v64 + v64)) | 1))
            }
            _ = fi63
            _ = v64
        }
        v67 := []int{0, 411, 265}
        _ = v67
    }
    if ("" >= "a") {
        var v68 float32 = ((v54 / v1) + -2.25)
        v69 := GenS2{F0: 2}
        _ = v68
        _ = v69
    }
    for fi70 := 0; fi70 < 20; fi70++ {
        v3 = ((("hello" + ("a" + (("" + "") + ""))) > (("" + "a") + "\n\t")) || false)
        _ = fi70
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v54
    _ = v62
}
