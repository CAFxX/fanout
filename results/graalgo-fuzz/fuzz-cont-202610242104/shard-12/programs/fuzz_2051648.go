package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int16 }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int32; F1 string; F2 int }

func main() {
    var v1 float64 = 1.5
    var v2 int = 32768
    var v3 int64 = 648
    var v4 uint32 = 42
    var v5 uint = 840
    var v6 float32 = float32((((0.1 / v1) + float64((3.14159 + float32(-2.25)))) + v1))
    v7 := GenS1{F0: v3, F1: 7, F2: -32768}
    fmt.Println(v2, v1, v6)
    v8 := GenS1{F0: (v3 << (((^(v5))) & 63)), F1: -128, F2: -32768}
    if true {
        if true {
            var v9 uint = 382
            _ = v9
        } else {
            v10 := []int{}
            var v11 float64 = v1
            _ = v10
            _ = v11
        }
    } else {
        for fi12 := 0; fi12 < 13; fi12++ {
            switch fi12 {
            case 127:
                v13 := []int{(^(fi12)), (v2 | -1), (v2 + (fi12 & (v2 - fi12)))}
                _ = v13
            case -9223372036854775808:
                for fi14 := 0; fi14 < 17; fi14++ {
                    var v15 uint64 = 0
                    _ = fi14
                    _ = v15
                }
            case 329:
                if false {
                    for fi16 := 0; fi16 < 13; fi16++ {
                        v5 = 255
                        var v17 uint = (v5 % (((v5 - 255)) | 1))
                        var v18 int32 = 42
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                }
            }
            if false {
                for fi19 := 0; fi19 < 1; fi19++ {
                    var v20 string = "foo bar"
                    _ = fi19
                    _ = v20
                }
                var v21 string = ("foo bar" + "hello")
                _ = v21
            } else {
                var v22 float64 = 0.1
                _ = v22
            }
            v5 -= v5
            _ = fi12
        }
        v4 += v4
    }
    if (false && ((v4 > 678) || true)) {
        if (!(v4 > v4)) {
            if (!false) {
                v6 = 0.0
                fmt.Println(v4)
                if (255 != 200) {
                    v23 := []int{32767}
                    if false {
                        for fi24 := 0; fi24 < 17; fi24++ {
                            v4 += 523
                            var v25 uint16 = uint16(v5)
                            v4 = 2
                            _ = fi24
                            _ = v25
                        }
                    } else {
                        var v26 int = int(v4)
                        v26 -= v26
                        _ = v26
                    }
                    _ = v23
                } else {
                    if false {
                        if true {
                            var v27 uint16 = 571
                            v28 := []int{2147483647, (v2 & -369)}
                            v1 = 0.1
                            _ = v27
                            _ = v28
                        } else {
                            var v29 float64 = v1
                            _ = v29
                        }
                    } else {
                        v3 = int64(v5)
                        for fi30 := 0; fi30 < 7; fi30++ {
                            fmt.Println(v3, v2)
                            _ = fi30
                        }
                        var v31 uint = (v5 + (v5 & (^(v5))))
                        _ = v31
                    }
                    switch (^(v2)) {
                    case 80:
                        v32 := GenS1{F0: 56, F1: 0, F2: 32767}
                        _ = v32
                    case -2147483648:
                        if true {
                            var v33 float64 = float64(v6)
                            if (!(true || (true && true))) {
                                v2 += (-(v2))
                                v34 := GenS2{F0: 42, F1: 42}
                                var v35 uint = v5
                                _ = v34
                                _ = v35
                            }
                            _ = v33
                        } else {
                            var v36 uint32 = 2
                            fmt.Println(v36, v5)
                            _ = v36
                        }
                    case 4294967295:
                        fmt.Println(v5, v3, v5)
                    }
                }
            }
        } else {
            if false {
                var v37 int = 127
                _ = v37
            } else {
                for fi38 := 0; fi38 < 5; fi38++ {
                    if (v6 <= float32((v1 - (v1 - 0.1)))) {
                        if true {
                            var v39 uint = v5
                            _ = v39
                        } else {
                            if false {
                                fmt.Println(v3, v3, v2)
                                for fi40 := 0; fi40 < 12; fi40++ {
                                    fmt.Println(v5, v3, v4)
                                    v1 = 1.5
                                    var v41 string = (("foo bar" + "\n\t") + ("foo bar" + ("foo bar" + "hello")))
                                    _ = fi40
                                    _ = v41
                                }
                                switch v2 {
                                case 790:
                                    switch "" {
                                    case "hello":
                                        switch ("hello" + (("" + "foo bar") + "foo bar")) {
                                        case "\n\t":
                                            for fi42 := 0; fi42 < 9; fi42++ {
                                                fmt.Println(fi38, v2)
                                                v4 += (v4 + 165)
                                                _ = fi42
                                            }
                                        case "hello":
                                            v1 += v1
                                        }
                                    default:
                                        v43 := GenS2{F0: 11, F1: 255}
                                        _ = v43
                                    }
                                case 2147483647:
                                    var v44 int32 = 2147483647
                                    _ = v44
                                }
                            }
                            if (!(false && (!((v3 % (((v3 ^ (v3 >> ((v5) & 63)))) | 1)) >= int64(v2))))) {
                                v45 := []int{}
                                _ = v45
                            } else {
                                v46 := GenS3{F0: -32768}
                                _ = v46
                            }
                            for fi47 := 0; fi47 < 16; fi47++ {
                                v48 := GenS2{F0: 255, F1: 42}
                                fmt.Println(v1)
                                _ = fi47
                                _ = v48
                            }
                        }
                    } else {
                        if false {
                            v4 = v4
                            var v49 uint32 = v4
                            _ = v49
                        } else {
                            for fi50 := 0; fi50 < 5; fi50++ {
                                var v51 uint = 629
                                v6 += (v6 / ((v6 / 1e10) * -2.25))
                                _ = fi50
                                _ = v51
                            }
                        }
                    }
                    _ = fi38
                }
            }
        }
        v6 += v6
    }
    if false {
        fmt.Println(v5, v1, v5)
        switch ((("" + "\n\t") + "foo bar") + "a") {
        case "foo bar":
            for fi52 := 0; fi52 < 11; fi52++ {
                v53 := GenS4{F0: v5}
                v54 := []int{(v2 / ((fi52) | 1))}
                if ((0 != -571) && (((v2 * (-(fi52))) != 128) || true)) {
                    if (1e10 != -2.25) {
                        fmt.Println(v6, fi52)
                        var v55 bool = (148 < 9223372036854775807)
                        _ = v55
                    }
                    var v56 int16 = -32768
                    _ = v56
                }
                _ = fi52
                _ = v53
                _ = v54
            }
        case "hello":
            for fi57 := 0; fi57 < 1; fi57++ {
                var v58 uint32 = (v4 - v4)
                v59 := GenS5{F0: -1, F1: "", F2: (v2 << (((^(v5))) & 63))}
                fmt.Println(v1, v58, fi57)
                _ = fi57
                _ = v58
                _ = v59
            }
        default:
            var v60 int = v2
            _ = v60
        }
        if false {
            v61 := GenS1{F0: (v3 ^ (v3 << ((v5) & 63))), F1: 75, F2: -32768}
            _ = v61
        } else {
            if false {
                var v62 int32 = 1
                fmt.Println(v2, v4)
                _ = v62
            }
        }
    } else {
        var v63 int8 = 100
        _ = v63
    }
    switch ("" + ("hello" + "")) {
    case "foo bar":
        var v64 uint8 = 127
        _ = v64
    case "hello":
        var v65 float64 = (1.5 * v1)
        _ = v65
    }
    for fi66 := 0; fi66 < 15; fi66++ {
        var v67 uint64 = 18446744073709551615
        fmt.Println(v1, v1)
        _ = fi66
        _ = v67
    }
    fmt.Println(v6, v1, v2)
    var v68 string = "foo bar"
    if false {
        if false {
            var v69 float32 = ((((v6 + v6) - 0.0) + 1e10) - float32(v1))
            switch ((("foo bar" + v68) + v68) + (v68 + v68)) {
            case "foo bar":
                var v70 uint64 = 0
                _ = v70
            }
            if true {
                v71 := []int{v2, (v2 ^ 127)}
                v69 = 0.0
                _ = v71
            }
            _ = v69
        } else {
            v1 -= (((0.1 / 3.14159) * 0.1) / 1.5)
        }
    } else {
        for fi72 := 0; fi72 < 6; fi72++ {
            v73 := GenS5{F0: 2147483647, F1: "\n\t", F2: fi72}
            v74 := []int{v2, 2147483647, -9223372036854775808, (^(fi72))}
            _ = fi72
            _ = v73
            _ = v74
        }
        var v75 int = v2
        v76 := []int{-14}
        _ = v75
        _ = v76
    }
    switch ("hello" + "") {
    case "hello":
        for fi77 := 0; fi77 < 4; fi77++ {
            v78 := []int{(v2 & (fi77 + (fi77 & -849)))}
            switch (^(v2)) {
            case 42:
                fmt.Println(v68)
            case 736:
                if len(v78) > 0 {
                    v78[0] = -9223372036854775808
                }
            case 2147483647:
                fmt.Println(v4, v6, v5)
            default:
                fmt.Println(v68, v68, v5)
            }
            v2 = (-(fi77))
            _ = fi77
            _ = v78
        }
    case "\n\t":
        var v79 int = 2147483647
        _ = v79
    }
    for fi80 := 0; fi80 < 10; fi80++ {
        if (42 <= 65535) {
            if true {
                v4 -= (v4 | 42)
                switch v68 {
                case "a":
                    for fi81 := 0; fi81 < 16; fi81++ {
                        v82 := GenS3{F0: 1000}
                        _ = fi81
                        _ = v82
                    }
                case "":
                    for fi83 := 0; fi83 < 8; fi83++ {
                        switch v68 {
                        case "hello":
                            v84 := []int{0}
                            _ = v84
                        case "foo bar":
                            v85 := []int{}
                            _ = v85
                        default:
                            var v86 float64 = ((((0.0 + v1) - v1) - (v1 / (v1 * ((v1 * (float64(v6) + ((v1 + 0.0) + float64(v6)))) * (v1 / (1e10 - (v1 - ((v1 * v1) + v1)))))))) - v1)
                            _ = v86
                        }
                        var v87 bool = ((1e10 - ((v1 / (float64(v6) - (v1 - (v1 / v1)))) + float64((v6 + v6)))) == (1.5 - float64((v6 / v6))))
                        v1 += 0.0
                        _ = fi83
                        _ = v87
                    }
                case "hello":
                    v4 = v4
                default:
                    v1 = float64(1e10)
                }
                fmt.Println(v4)
            }
            var v88 int64 = (v3 << ((v5) & 63))
            _ = v88
        }
        var v89 uint32 = (v4 | (v4 + v4))
        _ = fi80
        _ = v89
    }
    v4 = (v4 / (((v4 % (((v4 / (((^(v4))) | 1))) | 1))) | 1))
    v6 += 0.1
    var v90 uint32 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v68:", v68)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v68
    _ = v7
    _ = v8
    _ = v90
}
