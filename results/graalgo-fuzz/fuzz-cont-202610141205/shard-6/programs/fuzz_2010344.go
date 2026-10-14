package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 uint; F1 uint16; F2 uint8 }
type GenS5 struct { F0 int16; F1 uint64; F2 uint32 }
type GenS6 struct { F0 int32; F1 uint }
func genFn0() bool {
    v7 := GenS1{F0: 128, F1: 266}
    var v8 int32 = -199
    _ = v7
    _ = v8
    return true
}

func main() {
    var v1 uint16 = 1
    v2 := make(map[string]int)
    var v3 bool = ((2 > 1) || false)
    v4 := make(map[uint32]string)
    var v5 uint32 = 2147483647
    v6 := []uint32{}
    fmt.Println(v3)
    if (1e10 != -2.25) {
        switch 1 {
        case 32:
            v9 := GenS2{F0: 127, F1: -90}
            _ = v9
        case 32767:
            _ = len(v2)
        case 1:
            v10 := GenS3{F0: "a", F1: true}
            _ = v10
        }
        v11 := GenS3{F0: "foo bar", F1: v3}
        _ = v11
    }
    if v3 {
        v12 := GenS1{F0: 9223372036854775807, F1: 2}
        var v13 float64 = 0.0
        if (!(("" + "foo bar") > "")) {
            var v14 int32 = 1
            _ = v14
        } else {
            if (32767 < -1) {
                for fi15 := 0; fi15 < 11; fi15++ {
                    var v16 int64 = 0
                    _ = fi15
                    _ = v16
                }
                for fi17 := 0; fi17 < 3; fi17++ {
                    var v18 uint = 2
                    _ = fi17
                    _ = v18
                }
                for fi19 := 0; fi19 < 6; fi19++ {
                    switch (("foo bar" + ("" + "a")) + "a") {
                    case "\n\t":
                        var v20 int8 = -128
                        _ = v20
                    }
                    v13 -= 3.14159
                    fmt.Println(v5, v3, fi19)
                    _ = fi19
                }
            } else {
                fmt.Println(v1, v5)
                if v3 {
                    for fi21 := 0; fi21 < 13; fi21++ {
                        switch "hello" {
                        case "":
                            v22 := GenS2{F0: 104, F1: 42}
                            _ = v22
                        default:
                            var v23 int16 = 1000
                            _ = v23
                        }
                        var v24 int = 0
                        _ = fi21
                        _ = v24
                    }
                    v6 = append(v6, v5)
                    if v3 {
                        var v25 int16 = 32767
                        v13 -= ((((v13 / ((v13 * 0.1) / 1e10)) + 0.1) - v13) + float64((-2.25 + ((0.0 - (((((3.14159 - (0.0 / 3.14159)) / 1.0) * 0.0) / 1e10) / (float32((3.14159 * 3.14159)) - 0.0))) * -2.25))))
                        _ = v25
                    } else {
                        v2[(("hello" + "a") + ("\n\t" + ("" + "hello")))] = 128
                    }
                }
            }
            v13 += (v13 + float64(3.14159))
            for fi26 := 0; fi26 < 5; fi26++ {
                if (-128 >= 127) {
                    delete(v4, v5)
                    fmt.Println(v13)
                } else {
                    v1 = 42
                }
                _ = fi26
            }
        }
        _ = v12
        _ = v13
    } else {
        v3 = v3
        fmt.Println(v5, v3)
        v5 += v5
    }
    fmt.Println(v3, v1, v3)
    v1 += (v1 / ((0) | 1))
    var v27 uint16 = (^(v1))
    v5 -= v5
    v3 = v3
    v28 := GenS1{F0: 9223372036854775807, F1: v5}
    if v3 {
        v27 = (^(v27))
        v6 = append(v6, 2)
        if true {
            var v29 int64 = 42
            _ = v29
        } else {
            var v30 uint16 = 737
            if false {
                v31 := GenS3{F0: "hello", F1: false}
                fmt.Println(v30, v5)
                v3 = v3
                _ = v31
            }
            for fi32 := 0; fi32 < 13; fi32++ {
                fmt.Println(fi32, v1)
                _ = fi32
            }
            _ = v30
        }
    } else {
        if (127 <= 32767) {
            v1 -= v1
            if (!true) {
                var v33 int8 = 0
                v3 = v3
                v6 = append(v6, (v5 - (v5 % ((v5) | 1))))
                _ = v33
            } else {
                v27 = v1
                if v3 {
                    if ((0.0 - 1.5) > float64((float32((0.1 - 1.5)) / (1.5 * 3.14159)))) {
                        v27 -= 1
                        if (!v3) {
                            var v34 int64 = -702
                            var v35 int = -1
                            v3 = ((v5 ^ 708) >= v5)
                            _ = v34
                            _ = v35
                        } else {
                            v1 += (^(v27))
                            v36 := GenS4{F0: 2, F1: 255, F2: 127}
                            _ = v36
                        }
                        for fi37 := 0; fi37 < 2; fi37++ {
                            fmt.Println(v5, fi37)
                            _ = fi37
                        }
                    }
                    v38 := GenS5{F0: 1, F1: 0, F2: v5}
                    switch (("\n\t" + ("" + ("a" + ""))) + ("\n\t" + "\n\t")) {
                    case "a":
                        switch "\n\t" {
                        case "":
                            if v3 {
                                v27 -= (v1 - (v1 * v27))
                            }
                        default:
                            if v3 {
                                switch -204 {
                                case 713:
                                    if len(v6) > 0 {
                                        v6[0] = v5
                                    }
                                default:
                                    if false {
                                        if true {
                                            v3 = (!v3)
                                            v3 = true
                                        } else {
                                            for fi39 := 0; fi39 < 11; fi39++ {
                                                delete(v4, (v5 & (v5 + v5)))
                                                var v40 string = "\n\t"
                                                _ = fi39
                                                _ = v40
                                            }
                                        }
                                        for fi41 := 0; fi41 < 2; fi41++ {
                                            if (!(!(v5 < (v5 % ((2) | 1))))) {
                                                if v3 {
                                                    var v42 int16 = 1000
                                                    _ = v42
                                                }
                                            }
                                            var v43 uint32 = 42
                                            _ = fi41
                                            _ = v43
                                        }
                                        v6 = append(v6, 2147483648)
                                    }
                                }
                            } else {
                                if v3 {
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                }
                                fmt.Println(v27)
                            }
                        }
                    case "\n\t":
                        var v44 int64 = 1
                        _ = v44
                    case "foo bar":
                        for fi45 := 0; fi45 < 13; fi45++ {
                            var v46 uint8 = 208
                            for fi47 := 0; fi47 < 11; fi47++ {
                                var v48 float32 = ((0.1 + ((((float32((0.0 * -2.25)) + ((0.0 / 0.1) + 0.0)) - 3.14159) - ((0.0 - ((((3.14159 * 0.1) + 0.0) * ((-2.25 * (3.14159 * 1.5)) / 3.14159)) - float32(float64(((3.14159 * ((1e10 * 1.5) - -2.25)) / float32((0.0 + float64(1e10)))))))) + -2.25)) * 1.5)) * float32(0.0))
                                v3 = (100 >= 1)
                                v6 = append(v6, v5)
                                switch (^(fi45)) {
                                case 32768:
                                    v49 := GenS6{F0: 788, F1: 4294967295}
                                    _ = v49
                                case 1:
                                    _ = len(v6)
                                default:
                                    if len(v6) > 0 {
                                        v6[0] = (v5 % ((v5) | 1))
                                    }
                                    switch "" {
                                    case "\n\t":
                                        var v50 int64 = 9223372036854775807
                                        _ = v50
                                    }
                                }
                                _ = fi47
                                _ = v48
                            }
                            v51 := GenS1{F0: fi45, F1: v5}
                            _ = fi45
                            _ = v46
                            _ = v51
                        }
                    default:
                        v52 := v4[108]
                        _ = v52
                    }
                    _ = v38
                }
            }
        } else {
            var v53 uint64 = uint64(v5)
            if (v3 && true) {
                v54 := GenS3{F0: (("\n\t" + "") + "a"), F1: (-9223372036854775808 == -9223372036854775808)}
                var v55 float32 = -2.25
                v3 = (255 > 18446744073709551615)
                _ = v54
                _ = v55
            }
            var v56 uint = uint(v5)
            _ = v53
            _ = v56
        }
    }
    v27 = v1
    for fi57 := 0; fi57 < 1; fi57++ {
        var v58 float64 = 1.5
        _ = fi57
        _ = v58
    }
    fmt.Println(v27, v1, v1)
    var v59 int64 = 945
    switch "\n\t" {
    case "\n\t":
        v1 -= 65535
    }
    var v60 float32 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v6
    _ = v60
}
