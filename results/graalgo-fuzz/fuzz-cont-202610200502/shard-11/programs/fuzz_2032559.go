package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 string; F1 bool; F2 int32 }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 int32; F1 uint }
type GenS5 struct { F0 uint }
func genFn0(p0 int64, p1 uint, p2 string) uint8 {
    fmt.Println(p0, p1, p2)
    switch -1 {
    case 32768:
        v5 := GenS1{F0: 32767, F1: "a"}
        _ = v5
    case 2147483647:
        var v6 int16 = -575
        _ = v6
    default:
        v7 := []int{12, -9223372036854775808}
        _ = v7
    }
    switch 2147483647 {
    case -9223372036854775808:
        p2 = p2
    case -2147483648:
        for fi8 := 0; fi8 < 3; fi8++ {
            if ((false || true) || (p1 == (p1 - (^(p1))))) {
                v9 := []int{fi8, 127, (fi8 >> ((p1) & 63))}
                _ = v9
            } else {
                fmt.Println(p1, p0, p2)
            }
            _ = fi8
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}
func genFn1() string {
    v10 := []int{32768, 4294967295, 32768}
    var v11 string = "hello"
    v11 = "foo bar"
    switch (((v11 + (((v11 + (((v11 + "\n\t") + "\n\t") + "a")) + v11) + v11)) + v11) + v11) {
    case "a":
        if (1 <= 32768) {
            v11 = v11
        }
    default:
        for fi12 := 0; fi12 < 12; fi12++ {
            v13 := GenS2{F0: (v11 + v11), F1: ((1 >= 2147483647) || (v11 == (v11 + v11))), F2: 296}
            _ = fi12
            _ = v13
        }
    }
    _ = v10
    _ = v11
    return (v11 + "foo bar")
}

func main() {
    var v1 float32 = 0.0
    var v2 float32 = 3.14159
    var v3 int32 = 295
    var v4 uint8 = 2
    var v14 int32 = (v3 >> ((138) & 63))
    for fi15 := 0; fi15 < 16; fi15++ {
        v16 := []int{}
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        fmt.Println(v14, v1)
        fmt.Println(v4, v1)
        if (fi17 >= fi17) {
            var v18 int32 = v3
            v19 := []int{fi17}
            _ = v18
            _ = v19
        } else {
            v14 += (v3 >> ((18446744073709551615) & 63))
            var v20 uint64 = 4294967295
            _ = v20
        }
        _ = fi17
    }
    var v21 bool = false
    if (v2 <= v1) {
        v22 := []int{128}
        _ = v22
    }
    var v23 uint = 65535
    v2 += v2
    v24 := []int{}
    v3 = v14
    switch 4294967295 {
    case -570:
        v23 -= (v23 * (v23 + (^(v23))))
    }
    v25 := GenS1{F0: 2, F1: "a"}
    if false {
        var v26 int = 1
        if (true && ((v1 * (v1 / v1)) <= 3.14159)) {
            switch (("hello" + ("foo bar" + "")) + ("hello" + (("a" + "a") + "foo bar"))) {
            case "hello":
                v27 := GenS1{F0: 32767, F1: ""}
                _ = v27
            default:
                var v28 int64 = -696
                _ = v28
            }
            for fi29 := 0; fi29 < 17; fi29++ {
                var v30 uint = 255
                fmt.Println(v1)
                v31 := GenS2{F0: ("" + "hello"), F1: v21, F2: v14}
                _ = fi29
                _ = v30
                _ = v31
            }
        }
        for fi32 := 0; fi32 < 12; fi32++ {
            switch "foo bar" {
            case "":
                v1 = v1
            case "\n\t":
                var v33 int32 = (v3 * 2147483647)
                _ = v33
            case "a":
                for fi34 := 0; fi34 < 18; fi34++ {
                    var v35 float64 = (-2.25 - -2.25)
                    if true {
                        switch "foo bar" {
                        case "a":
                            v3 = int32(fi32)
                        default:
                            _ = len(v24)
                        }
                    }
                    _ = fi34
                    _ = v35
                }
            default:
                if true {
                    v23 = 4294967295
                    v2 += ((((float32(-2.25) + v2) * 1e10) - (((v2 * ((v2 + float32(0.1)) * 1e10)) + v1) + 1.5)) / v2)
                    if (!(!false)) {
                        switch "hello" {
                        case "\n\t":
                            if ((!(v21 && ((^(v23)) == v23))) && false) {
                                _ = len(v24)
                                for fi36 := 0; fi36 < 4; fi36++ {
                                    if (v21 || (0.1 > ((3.14159 - 0.1) * (((0.1 - 1.5) / 3.14159) - (float64(v1) - ((3.14159 + 0.1) - -2.25)))))) {
                                        v21 = (true && (4294967295 <= (-(fi36))))
                                    } else {
                                        var v37 string = ("\n\t" + "")
                                        _ = v37
                                    }
                                    v38 := GenS3{F0: 2, F1: -1}
                                    var v39 int8 = -128
                                    _ = fi36
                                    _ = v38
                                    _ = v39
                                }
                            }
                        case "hello":
                            if false {
                                var v40 int64 = -10
                                _ = v40
                            }
                        case "":
                            v41 := GenS4{F0: v14, F1: 0}
                            _ = v41
                        default:
                            switch fi32 {
                            case 42:
                                if len(v24) > 0 {
                                    v24[0] = fi32
                                }
                            case 9223372036854775807:
                                v42 := GenS5{F0: (v23 - v23)}
                                _ = v42
                            }
                        }
                        var v43 int8 = 1
                        _ = v43
                    } else {
                        for fi44 := 0; fi44 < 5; fi44++ {
                            v26 = (fi32 % ((4294967295) | 1))
                            _ = fi44
                        }
                    }
                }
            }
            _ = fi32
        }
        _ = v26
    } else {
        switch 299 {
        case 161:
            if false {
                var v45 float32 = 1e10
                var v46 uint16 = 1
                if v21 {
                    var v47 int8 = 46
                    _ = v47
                } else {
                    for fi48 := 0; fi48 < 5; fi48++ {
                        if v21 {
                            v3 += -2147483648
                            var v49 float32 = (((((v2 - v45) / ((v45 * 1.5) / (((v2 - v2) / v2) / (float32((0.0 - 1.5)) - v1)))) - v45) * float32(0.1)) - (v1 + v1))
                            _ = v49
                        } else {
                            v23 += uint(v4)
                            if len(v24) > 0 {
                                v24[0] = (fi48 * fi48)
                            }
                            var v50 uint = v23
                            _ = v50
                        }
                        var v51 int = (fi48 >> ((v23) & 63))
                        var v52 string = "\n\t"
                        _ = fi48
                        _ = v51
                        _ = v52
                    }
                    v53 := GenS3{F0: 18446744073709551615, F1: 9223372036854775807}
                    _ = v53
                }
                _ = v45
                _ = v46
            }
        }
    }
    v24 = append(v24, 32768)
    fmt.Println(v23)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
