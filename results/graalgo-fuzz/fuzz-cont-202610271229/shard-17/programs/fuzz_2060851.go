package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 uint; F1 string; F2 uint64 }
type GenS5 struct { F0 uint64; F1 bool; F2 uint16 }
type GenS6 struct { F0 int32; F1 int64; F2 int8 }
type GenS7 struct { F0 int8 }

func main() {
    v1 := []int{}
    var v2 uint8 = 128
    var v3 float32 = (3.14159 + 1.5)
    var v4 uint = 18446744073709551615
    var v5 int32 = 2147483647
    var v6 int32 = v5
    fmt.Println(v6)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v7 uint = (v4 * 1)
    for fi8 := 0; fi8 < 12; fi8++ {
        v1 = append(v1, (fi8 | fi8))
        _ = fi8
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v9 int = 42
    if (127 != (v2 - 255)) {
        if true {
            fmt.Println(v3, v6)
            var v10 float32 = (0.1 / 1.0)
            _ = v10
        } else {
            if false {
                var v11 bool = (!(!false))
                for fi12 := 0; fi12 < 15; fi12++ {
                    v6 = (-(v6))
                    v9 = (v9 - (fi12 << (((^(v4))) & 63)))
                    _ = fi12
                }
                v4 -= 1
                _ = v11
            } else {
                v1 = append(v1, (v9 & (^(v9))))
            }
            v13 := GenS1{F0: "", F1: v7}
            if true {
                for fi14 := 0; fi14 < 12; fi14++ {
                    v15 := GenS2{F0: 1}
                    _ = fi14
                    _ = v15
                }
                _ = len(v1)
            } else {
                switch "\n\t" {
                case "a":
                    var v16 uint = 42
                    _ = v16
                case "":
                    v3 -= v3
                }
                v2 = v2
            }
            _ = v13
        }
        for fi17 := 0; fi17 < 2; fi17++ {
            switch "hello" {
            case "\n\t":
                if false {
                    var v18 int32 = v5
                    if (false || true) {
                        var v19 string = (("" + ((("hello" + "hello") + ("a" + "hello")) + (("foo bar" + "foo bar") + (("hello" + "") + (("a" + "\n\t") + "\n\t"))))) + ("hello" + ""))
                        _ = v19
                    } else {
                        if true {
                            if ((!(0.0 < 1.5)) && (false || (42 >= 65535))) {
                                var v20 uint32 = 634
                                var v21 uint64 = 1
                                var v22 uint = v7
                                _ = v20
                                _ = v21
                                _ = v22
                            } else {
                                switch "a" {
                                case "foo bar":
                                    fmt.Println(v7, v9)
                                }
                                if len(v1) > 0 {
                                    v1[0] = -9223372036854775808
                                }
                                v1 = append(v1, 926)
                            }
                            var v23 uint32 = 918
                            if false {
                                v1 = append(v1, fi17)
                                var v24 int64 = 42
                                _ = v24
                            }
                            _ = v23
                        }
                    }
                    _ = v18
                } else {
                    var v25 int = (^(fi17))
                    _ = v25
                }
            default:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            v26 := GenS2{F0: 42}
            _ = fi17
            _ = v26
        }
        v27 := GenS2{F0: -400}
        _ = v27
    } else {
        v3 += 3.14159
        v6 += v6
        v3 = float32((1e10 * 3.14159))
    }
    v28 := GenS2{F0: -1}
    v29 := GenS3{F0: 9223372036854775807, F1: -472}
    switch (v9 % ((v9) | 1)) {
    case 0:
        var v30 string = ("\n\t" + "hello")
        _ = v30
    case 129:
        var v31 uint64 = 456
        _ = v31
    case 694:
        var v32 bool = (-803 != 1)
        _ = v32
    }
    v3 -= (3.14159 / float32(3.14159))
    switch (("a" + "") + "foo bar") {
    case "a":
        var v33 int32 = -326
        _ = v33
    case "hello":
        switch "foo bar" {
        case "foo bar":
            v1 = append(v1, -951)
        case "\n\t":
            if len(v1) > 0 {
                v1[0] = 32768
            }
        default:
            for fi34 := 0; fi34 < 14; fi34++ {
                var v35 uint = v4
                _ = fi34
                _ = v35
            }
        }
    case "\n\t":
        v6 -= int32(v4)
    }
    if true {
        var v36 uint32 = 4294967295
        var v37 uint = (v4 ^ 736)
        fmt.Println(v5, v6)
        _ = v36
        _ = v37
    } else {
        v4 = (^(v4))
        if ((v6 * v6) == -1) {
            if ((!true) && (!true)) {
                if (1e10 >= (0.1 * 1.5)) {
                    v6 = (v5 + (v6 & v6))
                    var v38 uint16 = 1
                    _ = v38
                } else {
                    if (v5 != (v5 & v6)) {
                        if (((-2.25 / v3) - v3) <= ((v3 * (v3 / (1e10 + (0.1 - 1.5)))) + ((float32((((1.5 + 1e10) * (((0.0 - ((0.0 + 0.0) * 0.0)) / ((float64(float32(float64((0.1 / ((-2.25 + 0.1) + v3))))) / 1e10) * 1e10)) + 0.0)) * -2.25)) + (v3 * (3.14159 + (float32(0.1) * ((float32(float64(float32(((3.14159 / 0.1) - (0.1 - 1e10))))) / (0.0 + ((v3 / (0.1 * v3)) - ((0.1 / ((v3 + -2.25) - v3)) - v3)))) - 0.1))))) - v3))) {
                            for fi39 := 0; fi39 < 17; fi39++ {
                                v40 := GenS4{F0: (v4 ^ (v4 * 4294967295)), F1: "hello", F2: 4294967295}
                                _ = fi39
                                _ = v40
                            }
                            var v41 int8 = 111
                            switch (v9 >> (((^(v4))) & 63)) {
                            case 466:
                                v42 := GenS5{F0: 558, F1: true, F2: 42}
                                _ = v42
                            case -302:
                                fmt.Println(v4)
                            case 32768:
                                if (-2.25 >= v3) {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    v1 = append(v1, 32767)
                                }
                            default:
                                v43 := GenS5{F0: 42, F1: true, F2: 742}
                                _ = v43
                            }
                            _ = v41
                        }
                        var v44 uint32 = 619
                        _ = len(v1)
                        _ = v44
                    } else {
                        _ = len(v1)
                        var v45 int = v9
                        v46 := GenS6{F0: v5, F1: 68, F2: -73}
                        _ = v45
                        _ = v46
                    }
                }
                if ((-78 <= 1) && ((986 <= -32768) && (528 < 303))) {
                    v3 = 0.0
                } else {
                    fmt.Println(v3, v3, v9)
                    _ = len(v1)
                }
                v47 := GenS5{F0: 474, F1: true, F2: 2}
                _ = v47
            } else {
                v9 += (v9 - (^(v9)))
                var v48 int = v9
                v4 = v4
                _ = v48
            }
            v6 = 42
        }
    }
    v49 := GenS7{F0: -110}
    if len(v1) > 0 {
        v1[0] = 32767
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
