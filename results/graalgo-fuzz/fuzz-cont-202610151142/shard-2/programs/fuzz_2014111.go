package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int64 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 int; F1 int }
func genFn0(p0 int16, p1 int16) uint {
    if false {
        fmt.Println(p1, p0)
    }
    if ((true || false) && (21 > 19)) {
        var v6 uint32 = 897
        v7 := []int{9223372036854775807, 128, -615, -519}
        _ = v6
        _ = v7
    }
    p1 = (p1 - (p1 & p0))
    if ((-194 < 32767) && true) {
        v8 := []int{44, 360, 935}
        for fi9 := 0; fi9 < 9; fi9++ {
            var v10 int16 = 1
            _ = fi9
            _ = v10
        }
        p0 += 1
        _ = v8
    } else {
        v11 := []int{-1, 32768, -983}
        var v12 float32 = 3.14159
        fmt.Println(p1)
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    return 2
}
func genFn1(p0 uint, p1 int32) bool {
    var v13 uint64 = 42
    p1 = p1
    _ = p0
    _ = p1
    _ = v13
    return (3.14159 != 0.1)
}

func main() {
    var v1 float32 = 3.14159
    var v2 float64 = (((float64(v1) / 1.0) + (float64(float32(0.0)) / -2.25)) / 1.5)
    var v3 int8 = -128
    v4 := []string{"hello", "", "\n\t", ("\n\t" + "foo bar")}
    var v5 float64 = (((((1e10 - 0.1) * (v2 + 0.1)) - 0.1) - float64(float32(v2))) * (v2 / v2))
    for fi14 := 0; fi14 < 10; fi14++ {
        var v15 uint8 = 127
        var v16 uint = 946
        v1 += float32(v5)
        _ = fi14
        _ = v15
        _ = v16
    }
    if (!(568 > 255)) {
        v1 = v1
    } else {
        for fi17 := 0; fi17 < 8; fi17++ {
            var v18 int = fi17
            switch "\n\t" {
            case "foo bar":
                v19 := GenS1{F0: v18, F1: 485, F2: 1}
                _ = v19
            case "hello":
                for fi20 := 0; fi20 < 19; fi20++ {
                    var v21 int32 = 0
                    v18 -= v18
                    _ = fi20
                    _ = v21
                }
            default:
                var v22 int8 = v3
                _ = v22
            }
            if (428 == 2147483648) {
                var v23 int64 = 1
                fmt.Println(v5, v3)
                var v24 uint64 = 282
                _ = v23
                _ = v24
            }
            _ = fi17
            _ = v18
        }
        for fi25 := 0; fi25 < 2; fi25++ {
            v26 := GenS1{F0: -1, F1: 18446744073709551615, F2: -1}
            for fi27 := 0; fi27 < 14; fi27++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v28 float64 = float64((float32(float64(v1)) + v1))
                _ = fi27
                _ = v28
            }
            _ = fi25
            _ = v26
        }
    }
    v1 -= v1
    var v29 float32 = ((v1 - (0.0 - ((0.0 * 3.14159) - float32(v5)))) - v1)
    var v30 int = 9223372036854775807
    if ((("foo bar" + ("foo bar" + ((("hello" + "a") + "") + "a"))) + ("" + ("hello" + "\n\t"))) <= "hello") {
        for fi31 := 0; fi31 < 18; fi31++ {
            v3 += (v3 ^ (v3 >> ((4294967295) & 63)))
            if (9223372036854775807 <= 1) {
                if true {
                    v3 += 99
                } else {
                    switch (v30 * -2147483648) {
                    case 1:
                        v32 := GenS2{F0: 4294967295, F1: 704}
                        _ = v32
                    default:
                        switch "a" {
                        case "":
                            fmt.Println(v5, v29, v30)
                        case "foo bar":
                            switch ("foo bar" + ("foo bar" + "\n\t")) {
                            case "\n\t":
                                for fi33 := 0; fi33 < 12; fi33++ {
                                    if true {
                                        var v34 uint = 2
                                        _ = v34
                                    } else {
                                        v1 += v1
                                        var v35 int = (fi33 ^ 32768)
                                        v36 := GenS3{F0: (v30 & int(v3)), F1: fi33}
                                        _ = v35
                                        _ = v36
                                    }
                                    _ = fi33
                                }
                            }
                        case "\n\t":
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        default:
                            v2 = v2
                        }
                    }
                }
                switch "\n\t" {
                case "":
                    v37 := GenS1{F0: (v30 | fi31), F1: 904, F2: 1}
                    _ = v37
                case "foo bar":
                    if false {
                        var v38 uint = 42
                        _ = v38
                    }
                case "a":
                    switch (v30 * v30) {
                    case 42:
                        var v39 string = "hello"
                        _ = v39
                    case 32768:
                        switch "hello" {
                        case "a":
                            v5 += v2
                        case "":
                            if len(v4) > 0 {
                                v4[0] = "hello"
                            }
                        case "foo bar":
                            v3 = -94
                        }
                    default:
                        v1 = 0.1
                    }
                }
                if (true || ((!(255 <= 1)) && true)) {
                    v29 -= v29
                    fmt.Println(v30)
                    var v40 uint64 = 0
                    _ = v40
                } else {
                    if (genFn0(int16(v3), 42) > genFn0(-832, 42)) {
                        v1 += ((v1 - 3.14159) - (1.5 + v1))
                    }
                    var v41 int = (v30 + v30)
                    _ = v41
                }
            } else {
                if len(v4) > 0 {
                    v4[0] = "a"
                }
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi31
        }
        v30 = v30
        var v42 string = "hello"
        _ = v42
    } else {
        var v43 int = -589
        fmt.Println(v30)
        _ = v43
    }
    v30 -= 680
    for fi44 := 0; fi44 < 18; fi44++ {
        for fi45 := 0; fi45 < 15; fi45++ {
            var v46 float64 = (0.1 + (1e10 * v2))
            var v47 uint32 = 2
            _ = fi45
            _ = v46
            _ = v47
        }
        for fi48 := 0; fi48 < 20; fi48++ {
            var v49 int = v30
            var v50 uint = 422
            v50 += genFn0(-32768, -1)
            v51 := GenS3{F0: (^(fi44)), F1: (fi48 / ((127) | 1))}
            v2 += v5
            v52 := GenS3{F0: 4294967295, F1: v49}
            var v53 uint = v50
            _ = fi48
            _ = v49
            _ = v50
            _ = v51
            _ = v52
            _ = v53
        }
        var v54 float32 = (float32(v5) - 3.14159)
        _ = fi44
        _ = v54
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
