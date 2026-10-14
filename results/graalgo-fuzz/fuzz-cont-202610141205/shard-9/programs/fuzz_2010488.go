package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint }
type GenS2 struct { F0 bool; F1 uint64; F2 uint }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 bool; F1 uint32 }
type GenS5 struct { F0 int8; F1 uint }
type GenS6 struct { F0 uint8; F1 uint; F2 uint8 }
func genFn0(p0 uint8, p1 uint64, p2 int) uint64 {
    v7 := GenS1{F0: 65535, F1: ((((("" + "hello") + ("a" + ((("\n\t" + "a") + ("foo bar" + ("" + (("" + "\n\t") + "hello")))) + "a"))) + ("hello" + "hello")) + "\n\t") + "hello"), F2: 42}
    if (!false) {
        v8 := GenS1{F0: 42, F1: "foo bar", F2: 42}
        if false {
            p2 -= 32768
            p1 -= p1
        } else {
            var v9 int16 = 42
            v10 := []int{p2, (p2 % (((p2 % ((487) | 1))) | 1)), p2, (p2 * 1)}
            _ = v10
            _ = v9
        }
        _ = v8
    }
    if true {
        switch "a" {
        case "a":
            v11 := []int{(p2 >> ((0) & 63)), (p2 >> ((uint(p1)) & 63)), (-(p2)), 2147483647}
            _ = v11
        case "hello":
            switch 42 {
            case 32767:
                p1 = 298
            case 784:
                p0 += p0
            default:
                p2 = -1
            }
        default:
            p0 = 255
        }
        if true {
            v12 := []int{}
            p1 += 1
            _ = v12
        }
    } else {
        var v13 int = p2
        _ = v13
    }
    for fi14 := 0; fi14 < 2; fi14++ {
        if false {
            p0 -= p0
            var v15 uint32 = uint32(p0)
            _ = v15
        } else {
            v16 := GenS2{F0: ((float64((((((float32(float64(3.14159)) + ((0.1 * 1e10) / float32(3.14159))) * ((float32(1.5) - 1.5) * float32(0.1))) * 0.1) + (3.14159 * 1.5)) - 3.14159)) / 1.0) <= (float64(0.0) - (1e10 * ((float64(1e10) + 3.14159) - float64(((1.5 * 0.0) + (0.1 / 0.1))))))), F1: (p1 | (p1 >> ((65535) & 63))), F2: 441}
            if true {
                switch ("" + "foo bar") {
                case "a":
                    v17 := GenS1{F0: 4294967295, F1: "foo bar", F2: 805}
                    _ = v17
                case "":
                    for fi18 := 0; fi18 < 9; fi18++ {
                        fmt.Println(p1, fi18)
                        p0 += 128
                        var v19 int32 = 1
                        _ = fi18
                        _ = v19
                    }
                default:
                    switch fi14 {
                    case 885:
                        p2 = fi14
                    }
                }
            }
            _ = v16
        }
        v20 := []int{}
        _ = fi14
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (p1 | p1)
}

func main() {
    var v1 float64 = float64(float32((3.14159 + 1.5)))
    var v2 int64 = -9223372036854775808
    v3 := []uint32{2147483647, 65535, 2147483648, 2147483647}
    v4 := []uint32{2}
    var v5 string = "a"
    var v6 int16 = 578
    fmt.Println(v6)
    v6 -= (v6 & -32768)
    v21 := GenS1{F0: 2, F1: "a", F2: 18446744073709551615}
    if len(v4) > 0 {
        v4[0] = 831
    }
    _ = len(v4)
    var v22 int32 = -2147483648
    for fi23 := 0; fi23 < 15; fi23++ {
        for fi24 := 0; fi24 < 8; fi24++ {
            v5 = (((v5 + ((v5 + (("hello" + "\n\t") + "a")) + (v5 + "a"))) + "") + "\n\t")
            _ = len(v4)
            if (true && true) {
                v25 := GenS3{F0: -102, F1: "a"}
                var v26 uint8 = 127
                v4 = append(v4, 285)
                _ = v25
                _ = v26
            }
            v6 = (v6 / ((v6) | 1))
            _ = fi24
        }
        _ = fi23
    }
    if false {
        var v27 int8 = 1
        if false {
            v5 = v5
            for fi28 := 0; fi28 < 5; fi28++ {
                v6 = 1
                _ = fi28
            }
            fmt.Println(v1, v1)
        } else {
            for fi29 := 0; fi29 < 4; fi29++ {
                if (v22 != (v22 - (v22 >> ((2) & 63)))) {
                    var v30 float32 = 0.0
                    if (fi29 > (fi29 + 32768)) {
                        var v31 int8 = (v27 * v27)
                        _ = v31
                    }
                    _ = v30
                } else {
                    switch ("" + "hello") {
                    case "\n\t":
                        v27 = v27
                    case "a":
                        v22 = v22
                    case "foo bar":
                        for fi32 := 0; fi32 < 8; fi32++ {
                            if len(v3) > 0 {
                                v3[0] = 2147483647
                            }
                            if true {
                                v33 := GenS4{F0: false, F1: 597}
                                _ = v33
                            }
                            _ = fi32
                        }
                    default:
                        if true {
                            v34 := GenS1{F0: 853, F1: v5, F2: 662}
                            v5 = ((("foo bar" + v5) + v5) + "foo bar")
                            _ = v34
                        }
                    }
                }
                if (!true) {
                    var v35 bool = false
                    if v35 {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v36 int32 = v22
                        v37 := GenS1{F0: 0, F1: ((v5 + v5) + "hello"), F2: 1}
                        _ = v36
                        _ = v37
                    } else {
                        if ((fi29 & (fi29 * (fi29 | (fi29 >> ((39) & 63))))) < 42) {
                            switch (fi29 ^ 380) {
                            case -9223372036854775808:
                                v38 := GenS2{F0: (!(false && (1 <= 42))), F1: 99, F2: 2}
                                _ = v38
                            case -374:
                                v22 += (v22 - v22)
                            case 32767:
                                fmt.Println(v2)
                            default:
                                fmt.Println(fi29)
                            }
                            for fi39 := 0; fi39 < 18; fi39++ {
                                var v40 uint8 = 128
                                var v41 int16 = (v6 | (v6 + 42))
                                var v42 uint32 = 42
                                _ = fi39
                                _ = v40
                                _ = v41
                                _ = v42
                            }
                            var v43 int8 = 6
                            _ = v43
                        }
                        if len(v4) > 0 {
                            v4[0] = 698
                        }
                    }
                    var v44 int64 = 525
                    _ = v35
                    _ = v44
                }
                _ = fi29
            }
        }
        if (true && false) {
            if (v22 == v22) {
                var v45 int32 = 0
                v46 := GenS4{F0: ((2147483648 != 250) && ((^(v2)) < (v2 - (v2 | (v2 ^ 793))))), F1: 2147483648}
                switch v5 {
                case "foo bar":
                    var v47 uint64 = 0
                    _ = v47
                case "":
                    v48 := GenS5{F0: (v27 | -1), F1: 65535}
                    _ = v48
                default:
                    v49 := GenS3{F0: (v27 << ((2) & 63)), F1: v5}
                    _ = v49
                }
                _ = v45
                _ = v46
            }
        }
        _ = v27
    }
    v6 += v6
    v50 := GenS6{F0: 255, F1: 230, F2: 178}
    fmt.Println(v5)
    var v51 int8 = -128
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if false {
        for fi52 := 0; fi52 < 11; fi52++ {
            v51 += -1
            _ = fi52
        }
        v53 := GenS4{F0: (!false), F1: 2147483647}
        _ = len(v4)
        _ = v53
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v6
}
