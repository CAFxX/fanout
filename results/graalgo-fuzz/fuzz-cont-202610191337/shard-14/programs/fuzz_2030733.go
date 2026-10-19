package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 int16; F1 uint }
type GenS5 struct { F0 int32; F1 string }
type GenS6 struct { F0 int; F1 int; F2 uint16 }
type GenS7 struct { F0 uint8; F1 string }

func main() {
    var v1 bool = ((!false) && (!false))
    v2 := make(map[string]int)
    v3 := []int{507, -904}
    v4 := []string{"hello", "foo bar", ("" + "hello")}
    v1 = (((1e10 - (-2.25 - 3.14159)) - (1.5 / 1e10)) >= 0.0)
    v1 = (9 > 243)
    if v1 {
        v1 = (1 > 1)
        for fi5 := 0; fi5 < 20; fi5++ {
            v1 = (!v1)
            for fi6 := 0; fi6 < 17; fi6++ {
                v1 = (0.1 < ((3.14159 + (float32(float64((((0.1 * float32(-2.25)) / ((-2.25 + ((-2.25 * ((0.1 * ((float32(0.1) - 1e10) + (0.1 / 1.0))) + 1e10)) / (((0.1 / 1e10) * 0.1) / -2.25))) / 1e10)) * (((0.1 * 0.1) / (((-2.25 * 1.5) + ((1e10 / 0.1) / 3.14159)) - (float32(1e10) * (-2.25 / float32(1.5))))) + -2.25)))) * 3.14159)) * float32(-2.25)))
                var v7 float64 = float64(1e10)
                _ = fi6
                _ = v7
            }
            v4 = append(v4, ("foo bar" + (("" + "foo bar") + ("" + (("a" + (("foo bar" + ("a" + ((("\n\t" + "hello") + "foo bar") + ("" + "hello")))) + ("foo bar" + ""))) + ("foo bar" + "\n\t"))))))
            _ = fi5
        }
        var v8 uint = 95
        _ = v8
    }
    var v9 uint16 = 575
    v9 = (v9 & v9)
    var v10 uint16 = (^(v9))
    var v11 float32 = 0.1
    for fi12 := 0; fi12 < 15; fi12++ {
        if (1 < 1000) {
            var v13 uint8 = 255
            _ = v13
        }
        _ = fi12
    }
    v1 = (255 > 128)
    var v14 int32 = 1
    for fi15 := 0; fi15 < 12; fi15++ {
        var v16 float32 = v11
        switch "a" {
        case "foo bar":
            var v17 uint32 = 4294967295
            _ = v17
        }
        switch -1 {
        case 0:
            v1 = true
        case 192:
            delete(v2, ("" + "foo bar"))
        default:
            v9 = v10
        }
        _ = fi15
        _ = v16
    }
    var v18 int64 = -9223372036854775808
    if v1 {
        switch ("foo bar" + (("" + "foo bar") + "a")) {
        case "\n\t":
            v19 := v2[("" + "hello")]
            _ = v19
        }
    } else {
        for fi20 := 0; fi20 < 13; fi20++ {
            for fi21 := 0; fi21 < 9; fi21++ {
                var v22 int64 = v18
                v23 := GenS1{F0: 1000}
                if v1 {
                    fmt.Println(v9, v22, v14)
                    switch fi21 {
                    case 128:
                        v24 := GenS2{F0: 1, F1: 2}
                        _ = v24
                    default:
                        var v25 float64 = 0.0
                        _ = v25
                    }
                    v26 := GenS3{F0: 4294967295, F1: 42}
                    _ = v26
                } else {
                    switch (fi21 * fi21) {
                    case 362:
                        var v27 uint16 = 42
                        _ = v27
                    case 32768:
                        var v28 int64 = (v22 * 817)
                        v10 = (v9 & (^(v10)))
                        delete(v2, "a")
                        _ = v28
                    case 969:
                        v29 := GenS1{F0: 42}
                        _ = v29
                    default:
                        if (!v1) {
                            v30 := GenS4{F0: 42, F1: 2}
                            _ = v30
                        }
                    }
                    v22 = 1
                    switch 127 {
                    case -45:
                        v31 := GenS5{F0: (v14 ^ v14), F1: (((("a" + "foo bar") + "hello") + ("" + ("a" + "hello"))) + (("\n\t" + ("a" + (("foo bar" + "\n\t") + ("hello" + "hello")))) + "hello"))}
                        v14 -= (v14 << ((uint(v9)) & 63))
                        _ = v31
                    case -647:
                        if v1 {
                            var v32 uint8 = 255
                            _ = v32
                        } else {
                            var v33 int8 = 1
                            var v34 bool = true
                            _ = v33
                            _ = v34
                        }
                        switch "foo bar" {
                        case "hello":
                            var v35 uint8 = 135
                            _ = v35
                        case "foo bar":
                            v36 := GenS3{F0: 0, F1: 1}
                            _ = v36
                        default:
                            if (!v1) {
                                v9 = v9
                                var v37 int8 = -93
                                v38 := GenS2{F0: 0, F1: 42}
                                v39 := GenS6{F0: (^(fi20)), F1: 32767, F2: (v10 % ((v10) | 1))}
                                _ = v37
                                _ = v38
                                _ = v39
                            } else {
                                fmt.Println(v10, fi21)
                                var v40 bool = (int8(fi20) == 71)
                                fmt.Println(v9, v22, v14)
                                _ = v40
                            }
                            fmt.Println(v14, v22)
                        }
                    case -860:
                        fmt.Println(fi21, v9)
                    default:
                        var v41 uint8 = 128
                        _ = v41
                    }
                }
                _ = fi21
                _ = v22
                _ = v23
            }
            var v42 float32 = v11
            switch (fi20 * fi20) {
            case 32767:
                v43 := GenS5{F0: v14, F1: "hello"}
                _ = v43
            case -9223372036854775808:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case 932:
                v10 += (^(v9))
            }
            _ = fi20
            _ = v42
        }
        if ((2 == 0) || (!(v11 != v11))) {
            switch 9223372036854775807 {
            case 128:
                v18 -= (v18 * v18)
            case -1:
                for fi44 := 0; fi44 < 1; fi44++ {
                    v2[("a" + ("\n\t" + "foo bar"))] = fi44
                    v1 = (v18 >= -322)
                    switch "\n\t" {
                    case "hello":
                        var v45 uint32 = 65535
                        _ = v45
                    case "foo bar":
                        if (!(42 == (-(v18)))) {
                            var v46 float64 = 1.5
                            v18 = (v18 >> ((522) & 63))
                            _ = v46
                        } else {
                            var v47 int16 = 32767
                            _ = v47
                        }
                    default:
                        fmt.Println(v14)
                    }
                    _ = fi44
                }
            case -103:
                v18 = v18
            }
        } else {
            fmt.Println(v11, v1, v11)
            for fi48 := 0; fi48 < 3; fi48++ {
                fmt.Println(v11)
                var v49 int8 = 0
                _ = fi48
                _ = v49
            }
            for fi50 := 0; fi50 < 9; fi50++ {
                v51 := GenS7{F0: 154, F1: ("a" + ("a" + ("hello" + ("foo bar" + "\n\t"))))}
                fmt.Println(v11, v14)
                v9 = (v10 - (v9 << ((535) & 63)))
                _ = fi50
                _ = v51
            }
        }
        var v52 float32 = 0.1
        _ = v52
    }
    v18 += v18
    v10 = 2
    v14 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
