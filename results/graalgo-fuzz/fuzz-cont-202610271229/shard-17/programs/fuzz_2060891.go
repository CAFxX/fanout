package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8; F1 uint64; F2 int64 }
type GenS5 struct { F0 string; F1 uint }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint16 }
func genFn0() int {
    for fi6 := 0; fi6 < 3; fi6++ {
        if (!true) {
            var v7 float32 = (1.5 - ((((((-2.25 / 1.0) + float32(((float64(float32(1e10)) + 0.1) + float64(0.0)))) - (1.5 - ((0.1 * (-2.25 * 0.0)) + float32(float64(-2.25))))) / (3.14159 + 1e10)) * 3.14159) - 0.1))
            v7 -= (1e10 - v7)
            _ = v7
        }
        if ((74 >= 255) || false) {
            v8 := GenS1{F0: 32767}
            var v9 float32 = 0.1
            _ = v8
            _ = v9
        }
        v10 := GenS1{F0: 65535}
        _ = fi6
        _ = v10
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        v12 := GenS1{F0: 255}
        _ = fi11
        _ = v12
    }
    v13 := GenS2{F0: 254}
    v14 := []int{127, 9223372036854775807}
    _ = v13
    _ = v14
    return -385
}
func genFn1() int32 {
    var v15 bool = (529 == -2147483648)
    v15 = v15
    v15 = v15
    var v16 float64 = 3.14159
    _ = v15
    _ = v16
    return 0
}

func main() {
    var v1 uint8 = 127
    v2 := make(map[string]int)
    var v3 float32 = 0.0
    var v4 uint8 = (v1 ^ v1)
    var v5 bool = (0 >= 42)
    for fi17 := 0; fi17 < 5; fi17++ {
        for fi18 := 0; fi18 < 13; fi18++ {
            v19 := GenS1{F0: 1}
            _ = fi18
            _ = v19
        }
        switch "a" {
        case "":
            for fi20 := 0; fi20 < 18; fi20++ {
                v3 = v3
                var v21 uint32 = 2147483647
                var v22 int32 = -2147483648
                _ = fi20
                _ = v21
                _ = v22
            }
        default:
            switch "\n\t" {
            case "foo bar":
                var v23 uint64 = 2
                _ = v23
            }
        }
        _ = fi17
    }
    var v24 uint8 = 128
    switch "a" {
    case "foo bar":
        switch "\n\t" {
        case "foo bar":
            fmt.Println(v5, v3)
        default:
            v5 = ((^(v24)) != v4)
        }
    case "hello":
        if (v3 != ((1.5 / v3) - (v3 + 0.0))) {
            v5 = v5
            var v25 float64 = float64(((float32(3.14159) - 1e10) * 0.0))
            v26 := GenS1{F0: 207}
            _ = v25
            _ = v26
        } else {
            for fi27 := 0; fi27 < 5; fi27++ {
                v28 := GenS2{F0: 0}
                _ = fi27
                _ = v28
            }
            var v29 int16 = -32768
            _ = len(v2)
            _ = v29
        }
    case "a":
        v24 += v24
    default:
        for fi30 := 0; fi30 < 12; fi30++ {
            v24 += (v1 + (v4 / ((v24) | 1)))
            for fi31 := 0; fi31 < 7; fi31++ {
                v4 = 2
                v32 := GenS3{F0: 4294967295}
                v4 += 2
                _ = fi31
                _ = v32
            }
            v33 := []int{}
            _ = fi30
            _ = v33
        }
    }
    v34 := GenS4{F0: 0, F1: 18446744073709551615, F2: -9223372036854775808}
    for fi35 := 0; fi35 < 20; fi35++ {
        v4 += v24
        v1 = v24
        _ = fi35
    }
    if v5 {
        v24 = v24
        var v36 string = "\n\t"
        _ = v36
    }
    v37 := GenS4{F0: 127, F1: 708, F2: -650}
    v38 := GenS4{F0: 185, F1: 18446744073709551615, F2: -311}
    v39 := v2["foo bar"]
    fmt.Println(v1, v39)
    fmt.Println(v4)
    var v40 bool = false
    if true {
        v41 := []int{(v39 - 0), -2147483648, 4294967295, (v39 & 9223372036854775807)}
        delete(v2, "")
        delete(v2, "foo bar")
        _ = v41
    }
    if v40 {
        fmt.Println(v3)
    }
    switch (v39 - v39) {
    case 9223372036854775807:
        fmt.Println(v3, v1, v1)
    }
    v42 := []int{genFn0(), v39, v39}
    if v40 {
        switch v39 {
        case -2147483648:
            if (100 == 100) {
                var v43 float64 = 3.14159
                _ = v43
            } else {
                v44 := GenS5{F0: "\n\t", F1: 18446744073709551615}
                _ = v44
            }
        case 127:
            delete(v2, "hello")
        }
    } else {
        if (-1 == 1) {
            for fi45 := 0; fi45 < 7; fi45++ {
                if true {
                    v39 -= genFn0()
                    switch genFn0() {
                    case 4294967295:
                        if (2 != 4294967295) {
                            if ((979 > 794) || ((v1 & (v24 % (((^(v24))) | 1))) >= v24)) {
                                if false {
                                    v5 = (-9223372036854775808 < 591)
                                    v46 := GenS6{F0: genFn1()}
                                    switch (("foo bar" + "a") + ("\n\t" + "a")) {
                                    case "foo bar":
                                        v24 -= 127
                                    case "a":
                                        var v47 uint16 = 0
                                        _ = v47
                                    default:
                                        var v48 int8 = -85
                                        _ = v48
                                    }
                                    _ = v46
                                } else {
                                    var v49 int64 = -1
                                    switch "a" {
                                    case "":
                                        var v50 int8 = 42
                                        _ = v50
                                    default:
                                        for fi51 := 0; fi51 < 10; fi51++ {
                                            var v52 int16 = -1
                                            v3 = (float32((0.1 * 0.1)) - v3)
                                            _ = fi51
                                            _ = v52
                                        }
                                    }
                                    v24 += (v4 * 127)
                                    _ = v49
                                }
                                var v53 float32 = float32(1e10)
                                v53 -= v3
                                _ = v53
                            }
                            v39 = genFn0()
                            switch "\n\t" {
                            case "hello":
                                v3 += (v3 + (float32(float64(v3)) + ((0.1 + -2.25) / v3)))
                            case "foo bar":
                                var v54 float32 = 1.5
                                _ = v54
                            }
                        }
                    case 42:
                        var v55 bool = (("a" + "") <= (("a" + ((("" + ("foo bar" + ("hello" + "\n\t"))) + ((("\n\t" + ("\n\t" + "a")) + "hello") + (("foo bar" + "hello") + "a"))) + (("foo bar" + "hello") + ("a" + (("hello" + "\n\t") + ("foo bar" + ("foo bar" + ("a" + "foo bar")))))))) + "foo bar"))
                        _ = v55
                    case -339:
                        var v56 uint = 0
                        _ = v56
                    default:
                        v3 = v3
                    }
                    for fi57 := 0; fi57 < 18; fi57++ {
                        switch fi45 {
                        case 32767:
                            v58 := v2["hello"]
                            _ = v58
                        }
                        var v59 string = "\n\t"
                        var v60 uint32 = 2147483647
                        _ = fi57
                        _ = v59
                        _ = v60
                    }
                } else {
                    var v61 float32 = (float32((((1e10 / 3.14159) * 1.5) + ((1e10 * 3.14159) - (0.1 - 1e10)))) - 0.1)
                    _ = v61
                }
                v62 := GenS6{F0: -2147483648}
                v63 := GenS6{F0: -280}
                _ = fi45
                _ = v62
                _ = v63
            }
            v4 -= (v4 & 255)
        } else {
            v64 := GenS7{F0: 0}
            v40 = (v1 <= (v24 ^ 128))
            _ = v64
        }
        if (false && v5) {
            v3 = v3
        }
    }
    if len(v42) >= 2 {
        v42 = v42[1:len(v42)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v34
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v42
    _ = v5
}
