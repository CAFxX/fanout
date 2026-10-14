package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int }
type GenS2 struct { F0 int16; F1 string; F2 uint64 }
type GenS3 struct { F0 int8; F1 bool; F2 uint8 }
type GenS4 struct { F0 uint8; F1 bool; F2 uint8 }
type GenS5 struct { F0 string }
func genFn0() uint8 {
    var v6 int8 = -128
    var v7 uint64 = 4294967295
    var v8 int = 128
    v9 := GenS1{F0: 997, F1: 2, F2: 9223372036854775807}
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 236
}
func genFn1(p0 uint16, p1 uint32) uint {
    v10 := []int{250, -152}
    if false {
        fmt.Println(p0, p0)
        for fi11 := 0; fi11 < 14; fi11++ {
            p1 = (p1 * p1)
            _ = fi11
        }
    } else {
        switch "foo bar" {
        case "":
            for fi12 := 0; fi12 < 13; fi12++ {
                fmt.Println(p1)
                var v13 int32 = -1
                _ = fi12
                _ = v13
            }
        default:
            fmt.Println(p1, p1)
        }
    }
    v14 := GenS2{F0: -32768, F1: ("" + "foo bar"), F2: 18446744073709551615}
    _ = p0
    _ = p1
    _ = v10
    _ = v14
    return 165
}

func main() {
    var v1 uint8 = 155
    v2 := make(map[int]string)
    var v3 int32 = 2147483647
    var v4 int = 465
    var v5 float64 = ((3.14159 - float64(float32((0.0 - 3.14159)))) * 1.5)
    v3 = (-(v3))
    for fi15 := 0; fi15 < 2; fi15++ {
        v2[(-(fi15))] = "hello"
        v5 += v5
        switch "a" {
        case "":
            v16 := []int{(v4 << ((87) & 63)), 1, (v4 % (((fi15 >> ((14) & 63))) | 1))}
            _ = v16
        case "\n\t":
            if false {
                v17 := []int{(v4 & (-(fi15))), (fi15 >> ((2) & 63)), (fi15 ^ 9223372036854775807)}
                fmt.Println(v1)
                v18 := GenS2{F0: -32768, F1: "a", F2: 599}
                _ = v17
                _ = v18
            }
        }
        _ = fi15
    }
    var v19 float64 = (float64((0.0 * 1.5)) * (-2.25 - (-2.25 - (float64((1.5 + -2.25)) / (float64(float32(v5)) + v5)))))
    if true {
        if true {
            if false {
                var v20 uint8 = genFn0()
                _ = v20
            }
            var v21 uint = 0
            _ = v21
        } else {
            v1 = (v1 * genFn0())
            var v22 int16 = 1000
            _ = v22
        }
        v23 := GenS2{F0: int16(v1), F1: ("hello" + "\n\t"), F2: 18446744073709551615}
        fmt.Println(v3, v3)
        _ = v23
    }
    v5 += v5
    for fi24 := 0; fi24 < 1; fi24++ {
        if (255 < 42) {
            switch ("foo bar" + ("\n\t" + ((("\n\t" + "foo bar") + ("" + ("foo bar" + (("foo bar" + "") + "foo bar")))) + "\n\t"))) {
            case "hello":
                v25 := GenS2{F0: 1000, F1: "foo bar", F2: 18446744073709551615}
                _ = v25
            default:
                v26 := GenS3{F0: 0, F1: ((1 >= 4294967295) || true), F2: 0}
                _ = v26
            }
        } else {
            v27 := []int{98, fi24, 42, 0}
            for fi28 := 0; fi28 < 7; fi28++ {
                v29 := GenS3{F0: 42, F1: (255 == 0), F2: v1}
                v30 := GenS1{F0: -32768, F1: 65535, F2: (v4 << ((378) & 63))}
                v19 -= v19
                _ = fi28
                _ = v29
                _ = v30
            }
            _ = v27
        }
        v1 -= 255
        v5 += ((v19 - -2.25) / 1.0)
        _ = fi24
    }
    v31 := GenS4{F0: (v1 * (v1 >> ((genFn1(65535, 496)) & 63))), F1: (!(false && true)), F2: (v1 % ((v1) | 1))}
    v32 := []int{(v4 ^ -9223372036854775808)}
    for fi33 := 0; fi33 < 20; fi33++ {
        fmt.Println(v5, v4)
        var v34 uint32 = 344
        v1 += v1
        _ = fi33
        _ = v34
    }
    if true {
        switch v4 {
        case 127:
            for fi35 := 0; fi35 < 20; fi35++ {
                if (259 == 1) {
                    v5 = v19
                    for fi36 := 0; fi36 < 13; fi36++ {
                        if (-1 >= int8(v3)) {
                            v19 += float64(1.5)
                            var v37 int = fi36
                            if (2 < 701) {
                                v38 := GenS3{F0: -1, F1: false, F2: (v1 ^ 2)}
                                _ = v38
                            } else {
                                switch "a" {
                                case "hello":
                                    v2[4294967295] = "\n\t"
                                default:
                                    var v39 int32 = (v3 + -2147483648)
                                    _ = v39
                                }
                                v32 = append(v32, fi36)
                            }
                            var v40 int8 = 42
                            var v41 bool = ((751 < 255) || false)
                            switch ((("foo bar" + ("hello" + "hello")) + "a") + "hello") {
                            case "a":
                                fmt.Println(v37)
                            case "foo bar":
                                var v42 float64 = (v5 / 1e10)
                                _ = v42
                            case "":
                                v43 := GenS5{F0: ""}
                                _ = v43
                            }
                            if v41 {
                                v44 := GenS5{F0: "a"}
                                _ = v44
                            } else {
                                fmt.Println(v5)
                                delete(v2, (v37 - v37))
                            }
                            switch int(v1) {
                            case 9223372036854775807:
                                if (v40 > (v40 + v40)) {
                                    var v45 uint16 = 65535
                                    var v46 uint = 42
                                    _ = v45
                                    _ = v46
                                }
                            case -9223372036854775808:
                                var v47 uint32 = 2147483647
                                _ = v47
                            }
                            _ = v37
                            _ = v40
                            _ = v41
                        }
                        _ = len(v32)
                        if len(v32) >= 2 {
                            v32 = v32[1:len(v32)]
                        }
                        _ = fi36
                    }
                }
                _ = fi35
            }
        default:
            var v48 uint32 = 2147483647
            _ = v48
        }
        v49 := GenS5{F0: "a"}
        _ = v49
    } else {
        var v50 int = v4
        v50 += v50
        var v51 uint8 = v1
        _ = v50
        _ = v51
    }
    switch "foo bar" {
    case "a":
        var v52 float64 = v19
        _ = v52
    case "\n\t":
        if (0 > uint16(v1)) {
            if false {
                for fi53 := 0; fi53 < 6; fi53++ {
                    switch "" {
                    case "\n\t":
                        v54 := GenS2{F0: 1000, F1: "hello", F2: 715}
                        _ = v54
                    }
                    for fi55 := 0; fi55 < 11; fi55++ {
                        if false {
                            v1 = 127
                            v1 -= genFn0()
                            var v56 float64 = ((((v19 / v19) + 1e10) + (v19 / 1.0)) / (v5 / 3.14159))
                            _ = v56
                        }
                        fmt.Println(fi55)
                        _ = len(v2)
                        if (fi55 <= (v4 - (v4 % ((v4) | 1)))) {
                            var v57 int32 = v3
                            if false {
                                v5 -= v5
                                fmt.Println(v3)
                                v58 := GenS4{F0: (^(v1)), F1: false, F2: 128}
                                _ = v58
                            }
                            _ = v57
                        }
                        if (true || true) {
                            v32 = append(v32, fi55)
                        }
                        _ = fi55
                    }
                    fmt.Println(v3)
                    _ = fi53
                }
            }
            for fi59 := 0; fi59 < 8; fi59++ {
                var v60 bool = (!(2 != 2))
                v19 = float64(1e10)
                fmt.Println(v4, v3)
                _ = fi59
                _ = v60
            }
            var v61 int64 = 933
            _ = v61
        }
    case "":
        v5 -= ((v19 + (((float64(((1e10 - (float32(((1e10 - 0.1) - v19)) + 1e10)) * 0.0)) / float64(1e10)) * 0.0) - v19)) * 3.14159)
    }
    v62 := v2[v4]
    fmt.Println(v19)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v62
}
