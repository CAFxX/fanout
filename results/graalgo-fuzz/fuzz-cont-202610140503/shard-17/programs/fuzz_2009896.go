package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int8; F1 uint16; F2 string }
func genFn0(p0 int, p1 int, p2 int64) int8 {
    if (!((true && true) || true)) {
        var v5 string = "\n\t"
        if (!(!true)) {
            v6 := GenS1{F0: p1}
            p1 = p0
            v7 := []int{p0}
            _ = v6
            _ = v7
        } else {
            if false {
                v8 := GenS2{F0: 4294967295}
                if ((-2147483648 >= 42) || (-8 < -128)) {
                    for fi9 := 0; fi9 < 7; fi9++ {
                        if (1 >= 1) {
                            var v10 uint16 = 65535
                            _ = v10
                        } else {
                            for fi11 := 0; fi11 < 18; fi11++ {
                                switch (p1 % (((-(p1))) | 1)) {
                                case -1:
                                    var v12 int8 = -9
                                    _ = v12
                                case 829:
                                    var v13 int32 = 377
                                    _ = v13
                                }
                                fmt.Println(p0, p1)
                                p2 = p2
                                _ = fi11
                            }
                        }
                        v14 := GenS2{F0: 147}
                        for fi15 := 0; fi15 < 15; fi15++ {
                            var v16 uint = 2
                            _ = fi15
                            _ = v16
                        }
                        _ = fi9
                        _ = v14
                    }
                }
                switch v5 {
                case "foo bar":
                    if (42 != 227) {
                        p2 = p2
                        p1 = p1
                    } else {
                        p2 = (^(p2))
                    }
                case "\n\t":
                    v17 := []int{}
                    _ = v17
                case "hello":
                    switch "a" {
                    case "\n\t":
                        v5 = ("\n\t" + v5)
                    }
                }
                _ = v8
            }
            p0 += 128
            for fi18 := 0; fi18 < 2; fi18++ {
                v19 := GenS1{F0: (p1 - -796)}
                _ = fi18
                _ = v19
            }
        }
        v20 := GenS2{F0: 4294967295}
        _ = v20
        _ = v5
    }
    v21 := GenS1{F0: (p1 * (p1 | 1))}
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    return -26
}
func genFn1(p0 uint32, p1 int64, p2 int8) uint16 {
    if false {
        var v22 bool = true
        v23 := GenS2{F0: 18446744073709551615}
        p0 -= 2147483648
        _ = v22
        _ = v23
    }
    p0 -= (^(p0))
    v24 := GenS1{F0: 128}
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    return 65535
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint16 = 2
    var v3 string = ""
    var v4 uint8 = 255
    v25 := GenS1{F0: 975}
    var v26 int8 = 100
    if false {
        for fi27 := 0; fi27 < 11; fi27++ {
            v28 := GenS2{F0: 65535}
            v29 := GenS3{F0: 18446744073709551615, F1: (v4 % (((v4 ^ 66)) | 1))}
            v30 := GenS4{F0: 610}
            _ = fi27
            _ = v28
            _ = v29
            _ = v30
        }
        var v31 int = -795
        _ = v31
    } else {
        var v32 float64 = (0.1 / 0.1)
        var v33 uint16 = v2
        _ = v32
        _ = v33
    }
    if true {
        var v34 uint8 = (v4 / ((v4) | 1))
        _ = v34
    } else {
        v35 := []int{42}
        _ = len(v35)
        _ = v35
    }
    var v36 float64 = ((((3.14159 * 0.1) / ((0.0 * (3.14159 - (0.1 + -2.25))) + 1.5)) + (1e10 - 1e10)) + (1.5 * (-2.25 - 1e10)))
    if true {
        var v37 float32 = 1.5
        if (-322 == 270) {
            for fi38 := 0; fi38 < 15; fi38++ {
                fmt.Println(v26, v4, v4)
                _ = fi38
            }
            v39 := GenS1{F0: int(v4)}
            switch 2147483647 {
            case -484:
                fmt.Println(v2)
            case 2147483647:
                for fi40 := 0; fi40 < 2; fi40++ {
                    if (!false) {
                        v41 := GenS1{F0: fi40}
                        _ = v41
                    }
                    _ = fi40
                }
            }
            _ = v39
        } else {
            for fi42 := 0; fi42 < 10; fi42++ {
                for fi43 := 0; fi43 < 15; fi43++ {
                    var v44 int8 = v26
                    _ = fi43
                    _ = v44
                }
                var v45 int64 = 952
                v46 := GenS1{F0: fi42}
                _ = fi42
                _ = v45
                _ = v46
            }
        }
        _ = v37
    } else {
        v36 -= v36
    }
    v2 -= v2
    delete(v1, v3)
    if true {
        switch ("foo bar" + (v3 + "foo bar")) {
        case "foo bar":
            switch v3 {
            case "hello":
                if false {
                    if (1 <= 4294967295) {
                        var v47 string = "a"
                        _ = v47
                    }
                    var v48 uint16 = genFn1(310, 42, (v26 * (v26 >> ((65535) & 63))))
                    var v49 int8 = v26
                    _ = v48
                    _ = v49
                } else {
                    var v50 int8 = genFn0(-291, 1, 1)
                    fmt.Println(v4)
                    _ = v50
                }
            case "\n\t":
                delete(v1, "")
            }
        }
        if ((v26 << ((2) & 63)) <= (v26 / ((-128) | 1))) {
            var v51 float32 = 1.5
            _ = v51
        }
    }
    switch "\n\t" {
    case "hello":
        v52 := GenS4{F0: 58}
        _ = v52
    case "a":
        if true {
            v3 = v3
            var v53 float32 = 0.1
            _ = v53
        } else {
            for fi54 := 0; fi54 < 8; fi54++ {
                if (592 < -635) {
                    switch v3 {
                    case "foo bar":
                        var v55 int = (^(fi54))
                        _ = v55
                    case "\n\t":
                        v56 := GenS3{F0: uint(v2), F1: v4}
                        _ = v56
                    case "hello":
                        var v57 uint16 = 1
                        _ = v57
                    }
                }
                switch v3 {
                case "":
                    v2 -= 219
                default:
                    if (1e10 != (3.14159 * 0.1)) {
                        if false {
                            for fi58 := 0; fi58 < 6; fi58++ {
                                var v59 uint16 = v2
                                var v60 float64 = 0.1
                                _ = fi58
                                _ = v59
                                _ = v60
                            }
                        }
                    } else {
                        switch "" {
                        case "a":
                            for fi61 := 0; fi61 < 20; fi61++ {
                                var v62 uint32 = 2147483648
                                _ = fi61
                                _ = v62
                            }
                        }
                        v63 := GenS4{F0: 1}
                        _ = v63
                    }
                }
                var v64 uint = 255
                _ = fi54
                _ = v64
            }
        }
    default:
        if (255 == genFn1(2, -906, (v26 ^ 127))) {
            for fi65 := 0; fi65 < 20; fi65++ {
                if (0 < -32768) {
                    if (!(!(v3 <= ""))) {
                        v66 := GenS1{F0: fi65}
                        v26 += (v26 ^ (v26 * (v26 / (((v26 - int8(v4))) | 1))))
                        _ = v66
                    } else {
                        v67 := GenS1{F0: (^(fi65))}
                        v68 := []int{}
                        _ = v67
                        _ = v68
                    }
                    var v69 uint32 = 0
                    v26 += v26
                    _ = v69
                } else {
                    v2 = 0
                }
                v70 := GenS1{F0: (fi65 ^ 128)}
                if (42 != 4294967295) {
                    if (0 < -215) {
                        for fi71 := 0; fi71 < 5; fi71++ {
                            v72 := GenS5{F0: 0, F1: (^(v2)), F2: (v3 + ("foo bar" + v3))}
                            v26 = v26
                            var v73 uint8 = 128
                            _ = fi71
                            _ = v72
                            _ = v73
                        }
                    } else {
                        switch "" {
                        case "foo bar":
                            if false {
                                v74 := []int{(fi65 - (fi65 >> ((0) & 63)))}
                                v2 = (v2 - (v2 & (v2 ^ v2)))
                                _ = v74
                            }
                        default:
                            if false {
                                var v75 uint8 = v4
                                _ = v75
                            } else {
                                v2 -= v2
                            }
                        }
                    }
                    v26 = 42
                }
                _ = fi65
                _ = v70
            }
            v76 := GenS1{F0: 128}
            _ = v76
        }
    }
    switch v3 {
    case "hello":
        if (!false) {
            v3 = ("\n\t" + "")
            switch (((v3 + "foo bar") + v3) + (v3 + v3)) {
            case "\n\t":
                v3 = "a"
            default:
                var v77 float32 = 1.5
                _ = v77
            }
            fmt.Println(v2, v3, v36)
        } else {
            var v78 bool = (!(936 == 718))
            var v79 int8 = v26
            var v80 int = -371
            _ = v78
            _ = v79
            _ = v80
        }
    default:
        v3 = (("hello" + ("hello" + v3)) + v3)
    }
    var v81 uint8 = v4
    fmt.Println(v36, v26)
    v82 := v1[("hello" + "hello")]
    var v83 int32 = -413
    fmt.Println(v4)
    switch (v3 + v3) {
    case "a":
        v1["hello"] = true
    case "\n\t":
        switch 4294967295 {
        case -9223372036854775808:
            v4 = (^(v4))
        case -1:
            var v84 uint8 = v4
            _ = v84
        }
    case "":
        var v85 int32 = -137
        _ = v85
    default:
        v26 = 1
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v36
    _ = v4
    _ = v81
    _ = v82
    _ = v83
}
