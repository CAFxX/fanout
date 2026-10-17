package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 int16; F1 int; F2 int8 }
type GenS4 struct { F0 int8; F1 int8; F2 uint8 }
type GenS5 struct { F0 uint16; F1 int8 }
type GenS6 struct { F0 uint8; F1 bool; F2 int64 }
func genFn0(p0 string, p1 int8) uint16 {
    if (1 <= -1) {
        p1 += p1
        v5 := GenS1{F0: (2 >= 255)}
        v6 := []int{-839}
        _ = v5
        _ = v6
    } else {
        for fi7 := 0; fi7 < 1; fi7++ {
            p0 = "\n\t"
            _ = fi7
        }
    }
    v8 := []int{2147483647}
    if (!true) {
        p1 += p1
    } else {
        fmt.Println(p1, p0)
        for fi9 := 0; fi9 < 3; fi9++ {
            v10 := GenS1{F0: true}
            v11 := GenS2{F0: fi9, F1: "foo bar"}
            for fi12 := 0; fi12 < 15; fi12++ {
                _ = len(v8)
                var v13 float64 = 1e10
                _ = len(v8)
                _ = fi12
                _ = v13
            }
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    for fi14 := 0; fi14 < 5; fi14++ {
        fmt.Println(p0, fi14, p0)
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        var v15 uint32 = 2147483647
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 11; fi16++ {
        fmt.Println(p0)
        fmt.Println(p1)
        _ = len(v8)
        _ = fi16
    }
    _ = p0
    _ = p1
    _ = v8
    return 42
}
func genFn1(p0 string) int8 {
    for fi17 := 0; fi17 < 15; fi17++ {
        var v18 float64 = float64((((-2.25 * (3.14159 / 1e10)) - (1e10 + 0.1)) + 1e10))
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 12; fi19++ {
        switch p0 {
        case "":
            for fi20 := 0; fi20 < 4; fi20++ {
                var v21 float64 = -2.25
                var v22 uint32 = 4294967295
                _ = fi20
                _ = v21
                _ = v22
            }
        case "\n\t":
            p0 = ((p0 + ("foo bar" + ((p0 + p0) + ((("\n\t" + (p0 + "a")) + ((p0 + p0) + p0)) + "")))) + ("a" + "foo bar"))
        case "foo bar":
            if false {
                v23 := GenS2{F0: -9223372036854775808, F1: p0}
                if (42 < 134) {
                    var v24 int16 = 42
                    _ = v24
                } else {
                    var v25 int64 = -1
                    _ = v25
                }
                _ = v23
            }
        }
        if false {
            p0 = (p0 + "foo bar")
            v26 := GenS3{F0: 32767, F1: (fi19 & (-(fi19))), F2: int8(fi19)}
            var v27 uint16 = 0
            _ = v26
            _ = v27
        } else {
            if (0 == 1) {
                for fi28 := 0; fi28 < 6; fi28++ {
                    var v29 bool = (!false)
                    _ = fi28
                    _ = v29
                }
                var v30 int16 = -860
                _ = v30
            } else {
                if (!true) {
                    fmt.Println(fi19, p0)
                    p0 = p0
                }
                if false {
                    for fi31 := 0; fi31 < 5; fi31++ {
                        var v32 string = ((p0 + (p0 + p0)) + p0)
                        switch "foo bar" {
                        case "foo bar":
                            var v33 float64 = 1.5
                            _ = v33
                        case "\n\t":
                            v34 := GenS3{F0: 42, F1: (^(fi19)), F2: 1}
                            _ = v34
                        case "a":
                            v32 = (("a" + v32) + (("foo bar" + "foo bar") + (p0 + (p0 + ((p0 + "a") + "\n\t")))))
                        }
                        if true {
                            v35 := GenS4{F0: 61, F1: -36, F2: 1}
                            fmt.Println(fi19)
                            switch (fi19 * fi19) {
                            case 32767:
                                v36 := GenS5{F0: 42, F1: 42}
                                _ = v36
                            case 4294967295:
                                var v37 uint = 4294967295
                                _ = v37
                            }
                            _ = v35
                        } else {
                            v38 := []int{42, 128, (fi19 ^ 9223372036854775807), (fi31 | (fi31 - fi31))}
                            _ = v38
                        }
                        _ = fi31
                        _ = v32
                    }
                    var v39 uint32 = 2147483647
                    v40 := GenS3{F0: 1000, F1: (fi19 << ((0) & 63)), F2: 42}
                    _ = v39
                    _ = v40
                } else {
                    var v41 int = -1
                    _ = v41
                }
            }
            p0 = (p0 + p0)
            v42 := GenS4{F0: 0, F1: -128, F2: 255}
            _ = v42
        }
        _ = fi19
    }
    _ = p0
    return -1
}

func main() {
    v1 := []string{"", ("\n\t" + (("foo bar" + ("foo bar" + "a")) + "foo bar"))}
    var v2 uint32 = 2147483648
    var v3 int16 = 32767
    v4 := make(map[uint32]string)
    if (!((!true) && true)) {
        v2 -= v2
    } else {
        var v43 uint32 = v2
        _ = len(v4)
        for fi44 := 0; fi44 < 18; fi44++ {
            v3 = -550
            var v45 int32 = 0
            var v46 int = (fi44 & 776)
            _ = fi44
            _ = v45
            _ = v46
        }
        _ = v43
    }
    if (!false) {
        if false {
            switch ("a" + "a") {
            case "hello":
                if true {
                    if true {
                        var v47 float64 = 0.1
                        _ = v47
                    }
                    if (!true) {
                        var v48 float32 = ((1.5 + 0.1) + (3.14159 / 0.1))
                        v49 := GenS3{F0: (v3 & (v3 | v3)), F1: 81, F2: 127}
                        _ = v48
                        _ = v49
                    }
                }
            case "\n\t":
                v50 := GenS6{F0: 42, F1: false, F2: -82}
                _ = v50
            case "":
                var v51 float64 = 0.1
                _ = v51
            default:
                var v52 int16 = v3
                _ = v52
            }
            var v53 float32 = float32(0.1)
            if len(v1) > 0 {
                v1[0] = "foo bar"
            }
            _ = v53
        } else {
            var v54 float32 = 0.1
            _ = v54
        }
        var v55 float32 = -2.25
        _ = v55
    }
    for fi56 := 0; fi56 < 10; fi56++ {
        v3 += v3
        _ = fi56
    }
    if false {
        v2 -= v2
    }
    var v57 uint8 = 128
    v58 := v4[(v2 * (v2 << ((4294967295) & 63)))]
    for fi59 := 0; fi59 < 17; fi59++ {
        switch v58 {
        case "a":
            switch v58 {
            case "\n\t":
                v60 := v4[(^(v2))]
                _ = v60
            case "a":
                v61 := GenS1{F0: (false || false)}
                _ = v61
            case "hello":
                var v62 uint8 = v57
                _ = v62
            }
        case "foo bar":
            for fi63 := 0; fi63 < 8; fi63++ {
                var v64 int32 = 42
                _ = fi63
                _ = v64
            }
        default:
            var v65 int8 = 127
            _ = v65
        }
        var v66 string = "a"
        _ = fi59
        _ = v66
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 = (v2 + 42)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v57
    _ = v58
}
