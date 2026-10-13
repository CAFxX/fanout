package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int8 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 uint8 = 202
    var v3 float64 = 1.5
    var v4 bool = (false || (0.0 > 1.5))
    for fi5 := 0; fi5 < 17; fi5++ {
        v4 = (18446744073709551615 <= 18446744073709551615)
        v6 := []int{(fi5 >> ((1) & 63))}
        for fi7 := 0; fi7 < 18; fi7++ {
            if (!(666 > 255)) {
                var v8 uint8 = 42
                v4 = (!(1000 >= 783))
                _ = v8
            }
            var v9 int16 = 32767
            _ = fi7
            _ = v9
        }
        _ = fi5
        _ = v6
    }
    fmt.Println(v1, v3, v3)
    var v10 uint16 = 490
    v1 += (v1 % ((203) | 1))
    v11 := []int{0, 32768, 2147483647}
    for fi12 := 0; fi12 < 9; fi12++ {
        var v13 bool = false
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        var v15 uint64 = 42
        for fi16 := 0; fi16 < 15; fi16++ {
            if (((v3 * v3) / float64(3.14159)) < (v3 / ((3.14159 / v3) + -2.25))) {
                switch "foo bar" {
                case "":
                    if len(v11) > 0 {
                        v11[0] = fi14
                    }
                case "foo bar":
                    v10 = (v10 ^ 0)
                    if ((!v4) || v4) {
                        v15 = 42
                    }
                    var v17 int = (fi14 & (^(fi16)))
                    _ = v17
                default:
                    var v18 uint8 = (^(v2))
                    _ = v18
                }
                var v19 string = "hello"
                var v20 int64 = v1
                if v4 {
                    v3 = v3
                    if (v4 || ((-289 <= int64(v2)) || ((0.1 == 1e10) || false))) {
                        v21 := GenS1{F0: 1, F1: 1, F2: 42}
                        _ = len(v11)
                        var v22 int16 = 447
                        v20 = (v20 * v1)
                        _ = v21
                        _ = v22
                    } else {
                        fmt.Println(v10)
                        v1 += (v20 * v1)
                        v20 += v20
                        var v23 int = (fi14 - (^(fi16)))
                        _ = v23
                    }
                    var v24 int8 = 0
                    _ = v24
                } else {
                    v15 = v15
                    v25 := GenS1{F0: 0, F1: 0, F2: 42}
                    v1 -= 42
                    if len(v11) > 0 {
                        v11[0] = fi14
                    }
                    _ = v25
                }
                _ = v19
                _ = v20
            }
            if (!v4) {
                if (1 < uint(v15)) {
                    v1 -= 121
                } else {
                    if len(v11) > 0 {
                        v11[0] = (fi14 - (^(fi16)))
                    }
                    v26 := GenS1{F0: 1000, F1: 2147483647, F2: 0}
                    _ = v26
                }
                var v27 bool = ((-2.25 - 1.5) < 3.14159)
                if len(v11) > 0 {
                    v11[0] = 42
                }
                _ = v27
            } else {
                v10 = (v10 * (v10 | (v10 ^ 884)))
                _ = len(v11)
                if (255 <= 1) {
                    if true {
                        if (false || v4) {
                            if len(v11) > 0 {
                                v11[0] = 128
                            }
                            v2 += (v2 & (v2 % ((166) | 1)))
                            v28 := GenS2{F0: 65535, F1: (v1 * v1)}
                            if false {
                                switch fi16 {
                                case -842:
                                    v29 := GenS3{F0: (v10 << ((0) & 63))}
                                    _ = v29
                                case 4294967295:
                                    if len(v11) > 0 {
                                        v11[0] = (fi16 & fi14)
                                    }
                                case -2147483648:
                                    fmt.Println(v10, v4, v4)
                                default:
                                    fmt.Println(v10, v3, v2)
                                }
                                v10 -= (v10 >> ((0) & 63))
                            }
                            _ = v28
                        }
                    }
                    switch int(v15) {
                    case 32768:
                        if true {
                            if len(v11) > 0 {
                                v11[0] = (fi14 & fi14)
                            }
                        }
                    case -633:
                        var v30 bool = v4
                        _ = v30
                    case 28:
                        v10 -= v10
                        var v31 int = (^(fi16))
                        v32 := GenS1{F0: -32768, F1: 2147483647, F2: 127}
                        _ = v31
                        _ = v32
                    default:
                        if true {
                            var v33 int32 = 42
                            _ = v33
                        }
                    }
                    if ((0.0 + ((float32(-2.25) / 1.0) * -2.25)) <= 1.5) {
                        if (true || v4) {
                            var v34 int8 = 100
                            var v35 int = (fi16 / ((fi16) | 1))
                            var v36 uint16 = v10
                            v36 -= (^(v36))
                            _ = v34
                            _ = v35
                            _ = v36
                        }
                        switch (fi16 ^ fi14) {
                        case 9223372036854775807:
                            v37 := GenS1{F0: 0, F1: 0, F2: 11}
                            _ = v37
                        case 120:
                            var v38 int8 = 42
                            _ = v38
                        case 618:
                            switch "a" {
                            case "foo bar":
                                if (!v4) {
                                    var v39 uint = 4294967295
                                    _ = v39
                                }
                            case "\n\t":
                                var v40 int16 = -279
                                _ = v40
                            case "hello":
                                if (-72 != fi16) {
                                    v3 = 3.14159
                                } else {
                                    v41 := GenS3{F0: 1}
                                    v3 = (v3 + v3)
                                    _ = len(v11)
                                    _ = v41
                                }
                            }
                        }
                        v11 = append(v11, (fi14 >> ((65535) & 63)))
                    }
                    var v42 int8 = 127
                    _ = v42
                } else {
                    v10 += v10
                    fmt.Println(v15, v1, v2)
                    v43 := GenS3{F0: v10}
                    _ = v43
                }
            }
            _ = fi16
        }
        _ = fi14
        _ = v15
    }
    if len(v11) > 0 {
        v11[0] = 535
    }
    for fi44 := 0; fi44 < 14; fi44++ {
        var v45 uint64 = 304
        if false {
            v3 -= v3
            _ = len(v11)
        }
        _ = fi44
        _ = v45
    }
    if false {
        var v46 uint32 = 0
        if v4 {
            fmt.Println(v4, v4)
            v4 = ((v10 / ((820) | 1)) > 42)
        } else {
            switch ("foo bar" + (((("a" + "a") + "hello") + (((("hello" + "foo bar") + ("\n\t" + "\n\t")) + "hello") + "")) + ((("foo bar" + "\n\t") + "hello") + "\n\t"))) {
            case "hello":
                v1 -= (v1 << ((141) & 63))
            }
            v2 = v2
            fmt.Println(v2)
        }
        _ = v46
    } else {
        _ = len(v11)
        fmt.Println(v2, v3, v2)
        var v47 float32 = (-2.25 * 3.14159)
        _ = v47
    }
    var v48 int = 127
    v49 := GenS2{F0: 18446744073709551615, F1: (v1 | (v1 | 42))}
    var v50 float64 = (-2.25 + (3.14159 * float64((((1.5 + (-2.25 * float32(v3))) - (0.0 * 1.5)) / (3.14159 / 3.14159)))))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v50
}
