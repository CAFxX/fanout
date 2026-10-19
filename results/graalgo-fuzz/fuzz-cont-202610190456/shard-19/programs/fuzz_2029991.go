package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 int16; F2 uint8 }
type GenS3 struct { F0 uint64; F1 uint32; F2 string }
type GenS4 struct { F0 string; F1 int32 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int8; F1 uint64 }
func genFn0() bool {
    v6 := []int{338, 127, 32768, 32767}
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 int8 = 42
        if (!true) {
            for fi9 := 0; fi9 < 14; fi9++ {
                switch 0 {
                case 39:
                    var v10 float32 = 3.14159
                    var v11 bool = true
                    _ = v10
                    _ = v11
                case -221:
                    if true {
                        switch -713 {
                        case 664:
                            var v12 bool = (("hello" + "\n\t") == "foo bar")
                            if (false || (!(100 >= (v8 >> ((4294967295) & 63))))) {
                                fmt.Println(fi7, fi7)
                                if false {
                                    v13 := GenS1{F0: 1}
                                    v14 := GenS2{F0: (fi7 - 1), F1: -762, F2: 0}
                                    _ = v13
                                    _ = v14
                                } else {
                                    v15 := GenS2{F0: 9223372036854775807, F1: 1, F2: 0}
                                    var v16 int16 = 32767
                                    var v17 float64 = -2.25
                                    _ = v15
                                    _ = v16
                                    _ = v17
                                }
                                switch "" {
                                case "hello":
                                    v8 += v8
                                case "":
                                    v8 = v8
                                case "foo bar":
                                    if len(v6) > 0 {
                                        v6[0] = -261
                                    }
                                default:
                                    v18 := GenS3{F0: 0, F1: 2147483648, F2: "foo bar"}
                                    v19 := GenS1{F0: 20}
                                    fmt.Println(fi7)
                                    _ = len(v6)
                                    _ = v18
                                    _ = v19
                                }
                                switch ("foo bar" + "\n\t") {
                                case "\n\t":
                                    switch (fi7 / (((-(fi9))) | 1)) {
                                    case -1:
                                        switch ("foo bar" + "\n\t") {
                                        case "a":
                                            v8 = (v8 % ((-77) | 1))
                                        default:
                                            _ = len(v6)
                                        }
                                    }
                                case "foo bar":
                                    if (fi7 == int(v8)) {
                                        var v20 int8 = (^(v8))
                                        v21 := GenS4{F0: "hello", F1: -62}
                                        _ = v20
                                        _ = v21
                                    }
                                }
                            }
                            var v22 int = fi9
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            var v23 uint16 = 32767
                            v12 = v12
                            _ = len(v6)
                            _ = v12
                            _ = v22
                            _ = v23
                        }
                        if false {
                            switch "hello" {
                            case "a":
                                switch "\n\t" {
                                case "\n\t":
                                    switch "a" {
                                    case "hello":
                                        var v24 uint8 = 0
                                        _ = v24
                                    }
                                default:
                                    v8 = int8(fi7)
                                }
                            }
                        } else {
                            var v25 uint64 = 4294967295
                            _ = v25
                        }
                    } else {
                        v26 := GenS5{F0: 2147483647}
                        v8 = (v8 % (((v8 - (^(v8)))) | 1))
                        _ = v26
                    }
                }
                _ = fi9
            }
        }
        _ = fi7
        _ = v8
    }
    var v27 bool = false
    switch "\n\t" {
    case "foo bar":
        if len(v6) > 0 {
            v6[0] = -9223372036854775808
        }
    case "a":
        var v28 int16 = -32768
        _ = v28
    default:
        v29 := GenS1{F0: 0}
        _ = v29
    }
    _ = v27
    _ = v6
    return (v27 || (-128 > -38))
}
func genFn1(p0 int32, p1 int32) uint16 {
    switch ("a" + "") {
    case "hello":
        var v30 float64 = (0.1 * 0.1)
        _ = v30
    case "":
        switch "\n\t" {
        case "\n\t":
            if true {
                if false {
                    if (false || (0 != 32767)) {
                        if (0.1 > (0.1 / 1.5)) {
                            for fi31 := 0; fi31 < 1; fi31++ {
                                p1 = 0
                                v32 := GenS5{F0: 4294967295}
                                _ = fi31
                                _ = v32
                            }
                        }
                        var v33 float32 = (1e10 / (0.0 - 3.14159))
                        _ = v33
                    } else {
                        p0 = (^(p1))
                    }
                } else {
                    v34 := GenS3{F0: 42, F1: 377, F2: "foo bar"}
                    _ = v34
                }
            }
        case "a":
            p0 = p0
        }
    default:
        for fi35 := 0; fi35 < 3; fi35++ {
            p1 -= p0
            p1 -= -33
            for fi36 := 0; fi36 < 1; fi36++ {
                v37 := []int{989, fi35, 1}
                switch ("" + "\n\t") {
                case "a":
                    v38 := GenS6{F0: 1, F1: 18446744073709551615}
                    _ = v38
                default:
                    v39 := GenS2{F0: fi35, F1: 1, F2: 128}
                    _ = v39
                }
                fmt.Println(p0, p1, p0)
                _ = fi36
                _ = v37
            }
            _ = fi35
        }
    }
    p1 = 799
    _ = p0
    _ = p1
    return 32767
}

func main() {
    var v1 bool = (1 > 0)
    v2 := make(map[int]bool)
    var v3 float32 = 0.0
    var v4 float64 = -2.25
    var v5 int = 42
    var v40 int = v5
    var v41 bool = true
    var v42 uint8 = 255
    v4 = (-2.25 / -2.25)
    switch ("" + "") {
    case "":
        switch "" {
        case "hello":
            delete(v2, -588)
        }
    default:
        var v43 float64 = 0.1
        _ = v43
    }
    for fi44 := 0; fi44 < 8; fi44++ {
        v45 := GenS4{F0: ("a" + "foo bar"), F1: 42}
        v46 := []int{0}
        var v47 int = (v40 * fi44)
        _ = fi44
        _ = v45
        _ = v46
        _ = v47
    }
    var v48 uint = 1
    for fi49 := 0; fi49 < 20; fi49++ {
        if false {
            var v50 bool = true
            fmt.Println(v3, v50, v48)
            _ = v50
        } else {
            var v51 int = fi49
            _ = v51
        }
        _ = fi49
    }
    for fi52 := 0; fi52 < 5; fi52++ {
        v48 += (^(v48))
        v41 = v1
        _ = fi52
    }
    var v53 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v48
    _ = v5
    _ = v53
}
