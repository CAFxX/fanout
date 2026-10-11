package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 int16; F1 uint32; F2 uint8 }
type GenS3 struct { F0 uint; F1 uint64; F2 uint }
type GenS4 struct { F0 int; F1 int32; F2 int32 }
type GenS5 struct { F0 uint64; F1 int }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int8; F1 int8; F2 uint }
type GenS8 struct { F0 int16 }
func genFn0(p0 uint) bool {
    fmt.Println(p0)
    for fi5 := 0; fi5 < 8; fi5++ {
        switch 4294967295 {
        case 771:
            switch fi5 {
            case 995:
                var v6 uint = p0
                _ = v6
            }
        default:
            var v7 bool = false
            _ = v7
        }
        var v8 int16 = 32767
        v9 := GenS1{F0: 255, F1: false}
        _ = fi5
        _ = v8
        _ = v9
    }
    v10 := []int{127, 315, 0, 32767}
    for fi11 := 0; fi11 < 8; fi11++ {
        for fi12 := 0; fi12 < 18; fi12++ {
            p0 = p0
            _ = fi12
        }
        fmt.Println(p0, p0)
        _ = fi11
    }
    _ = p0
    _ = v10
    return (3 < 133)
}

func main() {
    var v1 uint64 = 42
    var v2 float32 = 3.14159
    v3 := make(map[int]int)
    var v4 int = 4294967295
    v13 := []int{}
    for fi14 := 0; fi14 < 9; fi14++ {
        v15 := GenS2{F0: -32768, F1: 666, F2: 196}
        var v16 float64 = (0.1 + 0.0)
        fmt.Println(v2, v4, v16)
        _ = fi14
        _ = v15
        _ = v16
    }
    v13 = append(v13, 184)
    v2 -= (0.1 / v2)
    v17 := GenS3{F0: 65535, F1: v1, F2: 136}
    v2 = (v2 * (v2 * (float32(-2.25) + float32(0.1))))
    if (!false) {
        switch ("foo bar" + ((("" + (("a" + ("hello" + "")) + "")) + ("" + "foo bar")) + "hello")) {
        case "hello":
            v4 = 711
        case "":
            var v18 float64 = 3.14159
            _ = v18
        default:
            switch (v4 * 639) {
            case 9223372036854775807:
                v3[v4] = v4
            case 128:
                var v19 int16 = 0
                _ = v19
            default:
                delete(v3, (v4 >> ((18446744073709551615) & 63)))
            }
        }
        v1 += (^(v1))
    } else {
        var v20 int = (v4 | (v4 >> ((42) & 63)))
        for fi21 := 0; fi21 < 19; fi21++ {
            v20 = v4
            fmt.Println(v4)
            _ = fi21
        }
        _ = v20
    }
    var v22 int = v4
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    var v23 uint16 = 1
    v1 = 42
    for fi24 := 0; fi24 < 8; fi24++ {
        if true {
            var v25 float64 = 1.5
            _ = v25
        } else {
            switch "" {
            case "a":
                var v26 string = ("hello" + "a")
                _ = v26
            case "":
                v27 := GenS4{F0: fi24, F1: 0, F2: -2147483648}
                _ = v27
            default:
                switch "hello" {
                case "a":
                    var v28 uint8 = 0
                    _ = v28
                default:
                    v2 -= (v2 - (v2 * v2))
                }
            }
        }
        v29 := GenS5{F0: 4294967295, F1: fi24}
        _ = fi24
        _ = v29
    }
    v1 += (v1 ^ v1)
    v23 -= v23
    if false {
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
        for fi30 := 0; fi30 < 12; fi30++ {
            v31 := GenS5{F0: v1, F1: fi30}
            _ = fi30
            _ = v31
        }
        if true {
            fmt.Println(v22, v23)
            switch -930 {
            case -2147483648:
                if (false || (32767 < -32768)) {
                    switch "foo bar" {
                    case "hello":
                        if (1e10 > (v2 * (v2 - (1.5 + 1e10)))) {
                            if len(v13) >= 2 {
                                v13 = v13[1:len(v13)]
                            }
                            if (!(((v2 - (float32(0.1) + 1e10)) / 3.14159) < (v2 - ((-2.25 * (3.14159 + v2)) * float32((3.14159 * 1e10)))))) {
                                if false {
                                    fmt.Println(v2)
                                    _ = len(v13)
                                }
                            } else {
                                var v32 int64 = 9223372036854775807
                                var v33 int = 1
                                _ = v32
                                _ = v33
                            }
                            var v34 int8 = -46
                            _ = v34
                        } else {
                            fmt.Println(v22, v22)
                            if (!true) {
                                var v35 bool = (-1 != (^(v22)))
                                v1 += (^(v1))
                                v35 = (0.1 != ((0.0 - 0.0) * (3.14159 - (0.1 - 3.14159))))
                                _ = v35
                            } else {
                                for fi36 := 0; fi36 < 2; fi36++ {
                                    if false {
                                        delete(v3, (fi36 << ((588) & 63)))
                                    } else {
                                        v3[(^(v22))] = fi36
                                        switch ("a" + "a") {
                                        case "hello":
                                            _ = len(v13)
                                        case "foo bar":
                                            if len(v13) >= 2 {
                                                v13 = v13[1:len(v13)]
                                            }
                                        default:
                                            v37 := GenS1{F0: 104, F1: false}
                                            _ = v37
                                        }
                                        v38 := GenS2{F0: 32767, F1: 87, F2: 255}
                                        _ = v38
                                    }
                                    if true {
                                        var v39 int8 = 32
                                        for fi40 := 0; fi40 < 16; fi40++ {
                                            if false {
                                                fmt.Println(fi40, fi40)
                                            } else {
                                                fmt.Println(v22, v4)
                                            }
                                            if (true || false) {
                                                var v41 int8 = (v39 | v39)
                                                v42 := GenS5{F0: (v1 * v1), F1: v22}
                                                if false {
                                                    if len(v13) > 0 {
                                                        v13[0] = 32768
                                                    }
                                                    v1 -= (v1 | uint64(v23))
                                                    v39 += v41
                                                    switch ("foo bar" + "hello") {
                                                    case "foo bar":
                                                        v23 += v23
                                                    case "":
                                                        var v43 int8 = v39
                                                        _ = v43
                                                    default:
                                                        if false {
                                                            v44 := GenS3{F0: 558, F1: 0, F2: 658}
                                                            v3[(^(fi36))] = fi36
                                                            if false {
                                                                v45 := GenS5{F0: (v1 * (v1 / ((v1) | 1))), F1: 580}
                                                                _ = v45
                                                            } else {
                                                                v46 := GenS6{F0: 65535}
                                                                _ = v46
                                                            }
                                                            v1 = (v1 - 506)
                                                            _ = v44
                                                        } else {
                                                            switch (fi40 % ((fi40) | 1)) {
                                                            case -856:
                                                                v47 := GenS7{F0: v39, F1: v41, F2: 1}
                                                                _ = v47
                                                            case -398:
                                                                v48 := GenS8{F0: -1}
                                                                _ = v48
                                                            case -248:
                                                                v3[fi40] = int(v1)
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    var v49 int32 = 599
                                                    v3[(v22 & -2147483648)] = fi40
                                                    v23 -= 556
                                                    _ = v49
                                                }
                                                _ = v41
                                                _ = v42
                                            } else {
                                                switch "" {
                                                case "foo bar":
                                                    var v50 float64 = float64(v2)
                                                    _ = v50
                                                }
                                            }
                                            var v51 uint = 1
                                            _ = fi40
                                            _ = v51
                                        }
                                        _ = v39
                                    } else {
                                        fmt.Println(v1, v1)
                                        var v52 string = "a"
                                        _ = v52
                                    }
                                    for fi53 := 0; fi53 < 14; fi53++ {
                                        if len(v13) > 0 {
                                            v13[0] = (v22 + (v22 + (fi53 >> ((836) & 63))))
                                        }
                                        _ = fi53
                                    }
                                    _ = fi36
                                }
                                var v54 int32 = 42
                                for fi55 := 0; fi55 < 2; fi55++ {
                                    var v56 int16 = -1
                                    _ = fi55
                                    _ = v56
                                }
                                _ = v54
                            }
                        }
                    }
                    _ = len(v13)
                    v22 += v22
                }
            case 518:
                if (!(int64(v1) > 757)) {
                    var v57 int32 = 928
                    _ = v57
                }
            default:
                v4 = v4
            }
            fmt.Println(v2, v2)
        }
    } else {
        var v58 int8 = -1
        _ = v58
    }
    switch "" {
    case "a":
        switch ("a" + (("hello" + ("\n\t" + "")) + "foo bar")) {
        case "a":
            v13 = append(v13, v4)
        case "\n\t":
            var v59 float64 = 0.1
            _ = v59
        case "foo bar":
            var v60 uint32 = 0
            _ = v60
        default:
            var v61 string = "a"
            _ = v61
        }
    default:
        var v62 int32 = 337
        _ = v62
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
