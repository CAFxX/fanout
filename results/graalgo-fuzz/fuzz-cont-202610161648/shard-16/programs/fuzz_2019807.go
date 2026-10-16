package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int16 }
type GenS2 struct { F0 string; F1 int64; F2 int16 }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint; F1 uint64 }
type GenS6 struct { F0 uint64; F1 int64 }
type GenS7 struct { F0 int32; F1 string }
type GenS8 struct { F0 bool; F1 uint }
type GenS9 struct { F0 string; F1 int64; F2 bool }
func genFn0() int16 {
    v4 := []int{447}
    var v5 string = ("foo bar" + ((("\n\t" + "hello") + "a") + "\n\t"))
    v6 := GenS1{F0: 127, F1: false, F2: 1000}
    _ = v4
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 bool = false
    var v2 string = (("a" + "foo bar") + ("\n\t" + (("" + "foo bar") + "\n\t")))
    var v3 float32 = 1e10
    fmt.Println(v2, v3, v1)
    v7 := []int{127}
    for fi8 := 0; fi8 < 3; fi8++ {
        switch (fi8 << ((65535) & 63)) {
        case -327:
            v2 = ""
        case -1:
            v2 = ""
        case 48:
            v1 = false
        default:
            var v9 int = (-(fi8))
            _ = v9
        }
        v10 := GenS1{F0: 100, F1: v1, F2: 1000}
        _ = fi8
        _ = v10
    }
    var v11 bool = false
    for fi12 := 0; fi12 < 12; fi12++ {
        v1 = (v11 || v11)
        _ = fi12
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        for fi14 := 0; fi14 < 12; fi14++ {
            var v15 uint8 = 1
            fmt.Println(fi14, v3)
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    fmt.Println(v2)
    v2 = ("\n\t" + v2)
    if (true && true) {
        v11 = (32767 < 9223372036854775807)
        var v16 string = "hello"
        if (4294967295 < 0) {
            for fi17 := 0; fi17 < 12; fi17++ {
                if v11 {
                    var v18 uint32 = 42
                    _ = v18
                } else {
                    v19 := GenS1{F0: 1, F1: v1, F2: -741}
                    v20 := GenS2{F0: v16, F1: 42, F2: 1}
                    if v11 {
                        var v21 uint = 255
                        _ = v21
                    }
                    _ = v19
                    _ = v20
                }
                switch ((v2 + v16) + "\n\t") {
                case "\n\t":
                    if v1 {
                        if (v11 && v1) {
                            for fi22 := 0; fi22 < 13; fi22++ {
                                var v23 int = -638
                                _ = fi22
                                _ = v23
                            }
                            switch v2 {
                            case "hello":
                                if v1 {
                                    for fi24 := 0; fi24 < 3; fi24++ {
                                        var v25 float32 = 3.14159
                                        _ = fi24
                                        _ = v25
                                    }
                                }
                            case "foo bar":
                                fmt.Println(v16, v11, v16)
                            case "":
                                v2 = v2
                            }
                            fmt.Println(v2, v1, v16)
                        } else {
                            for fi26 := 0; fi26 < 19; fi26++ {
                                v7 = append(v7, 4294967295)
                                if false {
                                    switch ((v2 + ("\n\t" + v16)) + "foo bar") {
                                    case "foo bar":
                                        switch fi17 {
                                        case -94:
                                            if len(v7) > 0 {
                                                v7[0] = (^(fi26))
                                            }
                                        default:
                                            var v27 uint = 4294967295
                                            _ = v27
                                        }
                                    }
                                } else {
                                    switch -85 {
                                    case 128:
                                        var v28 uint32 = 626
                                        _ = v28
                                    default:
                                        v11 = (4294967295 > 983)
                                    }
                                    if false {
                                        var v29 uint = 854
                                        var v30 float64 = ((((0.0 * -2.25) + 1e10) - 3.14159) - 1e10)
                                        v30 -= v30
                                        _ = v29
                                        _ = v30
                                    } else {
                                        v11 = v11
                                        v1 = (2 < 37)
                                        v11 = (2147483648 <= 2)
                                    }
                                    var v31 bool = v1
                                    _ = v31
                                }
                                v32 := GenS3{F0: 0, F1: fi17}
                                _ = fi26
                                _ = v32
                            }
                            v3 = (((0.1 * float32(1.5)) / (v3 - ((v3 - v3) / -2.25))) - 0.1)
                        }
                        v33 := GenS1{F0: 127, F1: (9223372036854775807 <= 42), F2: 32767}
                        if (!(1 >= 42)) {
                            var v34 int = fi17
                            for fi35 := 0; fi35 < 8; fi35++ {
                                var v36 int64 = 42
                                fmt.Println(v1, v36, v34)
                                _ = fi35
                                _ = v36
                            }
                            if true {
                                v37 := GenS4{F0: 58}
                                var v38 uint8 = 127
                                var v39 uint64 = 4294967295
                                _ = v37
                                _ = v38
                                _ = v39
                            } else {
                                for fi40 := 0; fi40 < 6; fi40++ {
                                    if v11 {
                                        v1 = (1000 > genFn0())
                                        if false {
                                            if true {
                                                switch 1 {
                                                case 127:
                                                    v41 := GenS4{F0: 104}
                                                    _ = v41
                                                }
                                                v1 = true
                                                if (v1 || ((!((v16 + v2) >= "hello")) && v1)) {
                                                    v1 = (!false)
                                                }
                                                v42 := GenS5{F0: 255, F1: 521}
                                                _ = v42
                                            } else {
                                                var v43 uint8 = 127
                                                if len(v7) > 0 {
                                                    v7[0] = 363
                                                }
                                                var v44 uint64 = 18446744073709551615
                                                _ = v43
                                                _ = v44
                                            }
                                            switch fi40 {
                                            case 688:
                                                var v45 uint32 = 2
                                                _ = v45
                                            default:
                                                v16 = ""
                                            }
                                            v3 -= float32(0.0)
                                        }
                                        var v46 int32 = -334
                                        _ = v46
                                    } else {
                                        v34 = 649
                                        var v47 uint32 = 540
                                        if true {
                                            var v48 float32 = ((((0.1 * v3) / v3) * ((v3 * 0.1) + (v3 + 0.0))) / v3)
                                            if ((("hello" + "hello") + ("a" + v16)) > v2) {
                                                var v49 int = v34
                                                _ = v49
                                            } else {
                                                v50 := GenS6{F0: 158, F1: 0}
                                                v7 = append(v7, (^(fi17)))
                                                _ = v50
                                            }
                                            _ = v48
                                        } else {
                                            var v51 uint16 = 32767
                                            if v11 {
                                                v52 := GenS5{F0: 1, F1: 42}
                                                var v53 uint64 = 0
                                                fmt.Println(v3, v53)
                                                _ = v52
                                                _ = v53
                                            }
                                            var v54 uint8 = 255
                                            _ = v51
                                            _ = v54
                                        }
                                        _ = v47
                                    }
                                    _ = fi40
                                }
                                var v55 int8 = 52
                                if v11 {
                                    for fi56 := 0; fi56 < 9; fi56++ {
                                        v3 = 1.5
                                        v2 = v2
                                        v1 = (fi17 <= fi17)
                                        v1 = (32767 <= -1)
                                        _ = fi56
                                    }
                                    var v57 uint16 = 2
                                    _ = v57
                                }
                                _ = v55
                            }
                            _ = v34
                        }
                        _ = v33
                    }
                }
                v58 := GenS7{F0: -2147483648, F1: "\n\t"}
                _ = fi17
                _ = v58
            }
        } else {
            for fi59 := 0; fi59 < 10; fi59++ {
                if (((907 <= 255) && false) && (!(v2 < ((v2 + "foo bar") + v16)))) {
                    var v60 float32 = float32((-2.25 - 0.1))
                    _ = v60
                }
                v61 := GenS8{F0: false, F1: 2}
                _ = fi59
                _ = v61
            }
        }
        _ = v16
    }
    var v62 int16 = 32767
    fmt.Println(v11)
    if ("\n\t" > v2) {
        if true {
            if v1 {
                var v63 int16 = (-(v62))
                _ = v63
            } else {
                var v64 float64 = 3.14159
                _ = v64
            }
            for fi65 := 0; fi65 < 6; fi65++ {
                if len(v7) > 0 {
                    v7[0] = (fi65 * (fi65 % (((fi65 | 1)) | 1)))
                }
                if v11 {
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    v7 = append(v7, -2147483648)
                } else {
                    v7 = append(v7, fi65)
                    var v66 bool = false
                    _ = v66
                }
                v67 := GenS7{F0: 1, F1: ""}
                _ = fi65
                _ = v67
            }
            var v68 uint32 = 1
            _ = v68
        }
    }
    v2 = "a"
    v69 := GenS9{F0: v2, F1: -1, F2: false}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v62:", v62)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v62
    _ = v69
    _ = v7
}
