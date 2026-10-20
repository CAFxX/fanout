package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64; F1 int16 }
type GenS5 struct { F0 uint16; F1 uint32 }
type GenS6 struct { F0 uint16; F1 int64; F2 uint64 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint32; F1 int8; F2 string }
func genFn0(p0 bool) int64 {
    var v5 uint8 = 2
    v6 := []int{-9223372036854775808, 32767}
    _ = p0
    _ = v5
    _ = v6
    return 9223372036854775807
}
func genFn1() bool {
    for fi7 := 0; fi7 < 15; fi7++ {
        switch fi7 {
        case 9223372036854775807:
            var v8 uint64 = 2
            _ = v8
        case 699:
            for fi9 := 0; fi9 < 6; fi9++ {
                v10 := GenS1{F0: 1, F1: 2}
                var v11 int64 = -9223372036854775808
                if true {
                    var v12 uint32 = 132
                    v11 = (v11 | -755)
                    _ = v12
                }
                _ = fi9
                _ = v10
                _ = v11
            }
        case -949:
            fmt.Println(fi7)
        default:
            v13 := GenS2{F0: false, F1: -9223372036854775808}
            _ = v13
        }
        var v14 int64 = 1
        _ = fi7
        _ = v14
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        var v16 uint64 = 2
        for fi17 := 0; fi17 < 8; fi17++ {
            if true {
                v18 := []int{(fi15 >> ((18446744073709551615) & 63))}
                _ = v18
            }
            v16 -= (v16 & (v16 * v16))
            _ = fi17
        }
        fmt.Println(v16)
        _ = fi15
        _ = v16
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        var v20 uint32 = 4294967295
        for fi21 := 0; fi21 < 1; fi21++ {
            if true {
                v22 := GenS2{F0: false, F1: fi21}
                v20 = 65535
                v20 = (^(v20))
                var v23 int16 = 1000
                _ = v22
                _ = v23
            } else {
                var v24 uint64 = 42
                v25 := []int{(fi21 * (fi19 ^ fi21)), 128, 4294967295, (fi21 | (-(fi19)))}
                _ = v24
                _ = v25
            }
            if true {
                fmt.Println(v20)
            } else {
                v26 := []int{fi19}
                _ = len(v26)
                var v27 string = (("" + ((("hello" + (("" + (("a" + "") + "")) + ((("\n\t" + ("foo bar" + (("\n\t" + "foo bar") + (("" + ("\n\t" + "foo bar")) + "\n\t")))) + "a") + ""))) + "a") + "\n\t")) + "a")
                _ = v26
                _ = v27
            }
            _ = fi21
        }
        if (!(0.0 < (3.14159 + 0.0))) {
            if true {
                if false {
                    v20 += (^(v20))
                    fmt.Println(fi19)
                }
                v28 := []int{}
                var v29 uint16 = 430
                _ = v28
                _ = v29
            }
            var v30 int64 = 71
            for fi31 := 0; fi31 < 3; fi31++ {
                switch "hello" {
                case "foo bar":
                    var v32 int64 = -496
                    _ = v32
                }
                v33 := GenS2{F0: (255 == 2), F1: fi19}
                _ = fi31
                _ = v33
            }
            _ = v30
        } else {
            for fi34 := 0; fi34 < 6; fi34++ {
                v35 := []int{(fi34 | fi19), fi19, (fi34 / ((0) | 1)), -2147483648}
                fmt.Println(fi19, fi34, v20)
                if ((742 < 32767) || false) {
                    fmt.Println(fi19, v20, fi34)
                    v20 = 754
                    if ("foo bar" < "a") {
                        var v36 int32 = 1
                        switch "hello" {
                        case "\n\t":
                            fmt.Println(fi34)
                        case "a":
                            switch "foo bar" {
                            case "hello":
                                v35 = append(v35, (fi34 * fi34))
                            case "":
                                v36 += (v36 ^ v36)
                            case "foo bar":
                                v36 = (v36 - (v36 << ((0) & 63)))
                            default:
                                v37 := GenS3{F0: 313}
                                _ = v37
                            }
                        }
                        _ = v36
                    }
                    var v38 uint8 = 183
                    _ = v38
                } else {
                    v20 += v20
                    var v39 uint = 42
                    _ = v39
                }
                _ = fi34
                _ = v35
            }
            fmt.Println(v20)
        }
        _ = fi19
        _ = v20
    }
    return (true && (false && (1e10 >= 3.14159)))
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = 0.0
    v3 := []string{}
    v4 := make(map[uint32]uint32)
    switch 4294967295 {
    case 42:
        v2 = 3.14159
    case 655:
        v2 -= v2
    case 127:
        fmt.Println(v2)
    }
    if (!true) {
        if (-68 >= 376) {
            var v40 uint32 = 4294967295
            _ = v40
        } else {
            switch "" {
            case "\n\t":
                v2 = (1.5 + 0.0)
            default:
                for fi41 := 0; fi41 < 7; fi41++ {
                    v2 -= v2
                    _ = fi41
                }
            }
        }
    } else {
        v42 := GenS3{F0: -1}
        if true {
            v43 := GenS2{F0: true, F1: 491}
            var v44 uint32 = 2147483648
            v45 := GenS4{F0: 4294967295, F1: 42}
            _ = v43
            _ = v44
            _ = v45
        } else {
            var v46 float32 = 0.0
            _ = v46
        }
        _ = v42
    }
    if true {
        var v47 uint16 = 42
        var v48 int = 0
        for fi49 := 0; fi49 < 14; fi49++ {
            var v50 int = -2147483648
            _ = fi49
            _ = v50
        }
        _ = v47
        _ = v48
    } else {
        v2 = float32((float64(float32((-2.25 + 1.5))) * (float64((0.0 + (float32(3.14159) / ((v2 - v2) / v2)))) * 0.1)))
    }
    fmt.Println(v2)
    for fi51 := 0; fi51 < 1; fi51++ {
        switch (fi51 / ((fi51) | 1)) {
        case 32768:
            for fi52 := 0; fi52 < 6; fi52++ {
                if true {
                    v53 := GenS5{F0: 118, F1: 0}
                    v54 := GenS6{F0: 1, F1: -9223372036854775808, F2: 18446744073709551615}
                    v55 := GenS7{F0: 32767}
                    _ = v53
                    _ = v54
                    _ = v55
                } else {
                    var v56 int8 = -18
                    fmt.Println(fi52, fi51, fi51)
                    v56 = 0
                    _ = v56
                }
                _ = fi52
            }
        default:
            var v57 uint32 = 4294967295
            _ = v57
        }
        var v58 uint16 = 157
        if (("foo bar" + "a") <= "hello") {
            var v59 float32 = v2
            _ = v59
        }
        _ = fi51
        _ = v58
    }
    v2 -= 3.14159
    v60 := GenS8{F0: 620, F1: -128, F2: "foo bar"}
    if true {
        var v61 int = 32767
        _ = v61
    }
    v2 = ((v2 + 1e10) * ((v2 / v2) + -2.25))
    v2 -= (-2.25 / (float32(3.14159) / v2))
    fmt.Println(v2)
    var v62 bool = (!((true || false) || false))
    if (42 >= 1) {
        if len(v3) > 0 {
            v3[0] = "hello"
        }
    } else {
        var v63 bool = (float32(3.14159) < 3.14159)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi64 := 0; fi64 < 18; fi64++ {
            v3 = append(v3, "hello")
            if v63 {
                for fi65 := 0; fi65 < 7; fi65++ {
                    v62 = v63
                    _ = fi65
                }
                var v66 bool = v63
                _ = v66
            } else {
                v67 := GenS8{F0: 369, F1: -1, F2: ""}
                v62 = v62
                _ = v67
            }
            _ = fi64
        }
        _ = v63
    }
    for fi68 := 0; fi68 < 19; fi68++ {
        v2 += v2
        _ = fi68
    }
    var v69 uint32 = 334
    v62 = (false && ((!true) && true))
    _ = len(v4)
    v70 := v4[(v69 * (v69 & 707))]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v62:", v62)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v60
    _ = v62
    _ = v69
    _ = v70
}
