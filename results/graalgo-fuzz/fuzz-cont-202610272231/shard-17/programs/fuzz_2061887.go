package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint; F1 int64 }
type GenS5 struct { F0 uint; F1 int16 }
type GenS6 struct { F0 uint16; F1 int16; F2 uint }
func genFn0(p0 int, p1 uint16) uint16 {
    var v4 uint8 = 1
    var v5 uint64 = 1
    switch p0 {
    case 1:
        fmt.Println(p0)
    case 2147483647:
        switch (p0 + 4294967295) {
        case 128:
            v4 = 2
        default:
            var v6 string = ("foo bar" + "\n\t")
            _ = v6
        }
    default:
        for fi7 := 0; fi7 < 4; fi7++ {
            v8 := GenS1{F0: -267}
            var v9 uint8 = v4
            switch ("hello" + "a") {
            case "\n\t":
                fmt.Println(p0)
            case "foo bar":
                v4 += (v9 - (v9 - 127))
            default:
                var v10 float32 = 1e10
                _ = v10
            }
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    var v11 uint16 = p1
    _ = p0
    _ = p1
    _ = v11
    _ = v4
    _ = v5
    return (p1 << ((255) & 63))
}
func genFn1(p0 uint8) uint8 {
    var v12 int = 128
    var v13 float32 = (((-2.25 / 0.1) * 1.5) * 1.5)
    var v14 int64 = 174
    switch "\n\t" {
    case "foo bar":
        var v15 float32 = 0.1
        _ = v15
    case "a":
        v16 := []int{v12}
        _ = v16
    case "":
        var v17 uint = 65535
        _ = v17
    default:
        var v18 int16 = 127
        _ = v18
    }
    v13 = 3.14159
    _ = p0
    _ = v12
    _ = v13
    _ = v14
    return (^(p0))
}

func main() {
    v1 := []string{}
    var v2 int16 = -32768
    v3 := []string{"foo bar", ("hello" + "foo bar")}
    if true {
        var v19 int64 = 0
        if len(v1) > 0 {
            v1[0] = ("hello" + ("\n\t" + ("\n\t" + ("foo bar" + "hello"))))
        }
        _ = v19
    } else {
        var v20 bool = ("foo bar" > "\n\t")
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v21 int = 32768
        _ = v20
        _ = v21
    }
    fmt.Println(v2)
    for fi22 := 0; fi22 < 11; fi22++ {
        v23 := GenS2{F0: "a", F1: 255}
        _ = fi22
        _ = v23
    }
    v24 := GenS2{F0: "hello", F1: 9223372036854775807}
    if (255 != 42) {
        for fi25 := 0; fi25 < 4; fi25++ {
            _ = len(v3)
            v26 := GenS1{F0: -2147483648}
            _ = fi25
            _ = v26
        }
        v27 := GenS3{F0: 18446744073709551615}
        if (4294967295 >= 2147483648) {
            if false {
                v2 -= v2
                v1 = append(v1, "")
            } else {
                v28 := GenS3{F0: 4294967295}
                for fi29 := 0; fi29 < 6; fi29++ {
                    v2 -= v2
                    _ = fi29
                }
                _ = v28
            }
            if true {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v2 -= (v2 >> ((13) & 63))
                if true {
                    if true {
                        v2 = 576
                    }
                    var v30 int32 = -901
                    for fi31 := 0; fi31 < 19; fi31++ {
                        switch "\n\t" {
                        case "\n\t":
                            var v32 uint = 584
                            _ = v32
                        default:
                            switch fi31 {
                            case -816:
                                var v33 uint8 = 128
                                _ = v33
                            case 42:
                                fmt.Println(v2, v30)
                            case -1:
                                fmt.Println(v30, v2)
                            default:
                                if true {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    var v34 int64 = -9223372036854775808
                                    _ = v34
                                }
                            }
                        }
                        switch (fi31 * (fi31 - fi31)) {
                        case -9223372036854775808:
                            var v35 string = "hello"
                            _ = v35
                        case 127:
                            if false {
                                fmt.Println(v2, fi31)
                                v36 := GenS4{F0: 0, F1: 1}
                                _ = v36
                            } else {
                                for fi37 := 0; fi37 < 5; fi37++ {
                                    v2 = v2
                                    v3 = append(v3, "\n\t")
                                    _ = fi37
                                }
                                v1 = append(v1, "")
                                for fi38 := 0; fi38 < 5; fi38++ {
                                    var v39 uint16 = 0
                                    v2 = v2
                                    v39 = (v39 * (v39 - genFn0((-(fi31)), v39)))
                                    fmt.Println(v2)
                                    v40 := GenS5{F0: 2, F1: (v2 % ((172) | 1))}
                                    _ = fi38
                                    _ = v39
                                    _ = v40
                                }
                            }
                        case 661:
                            v1 = append(v1, "\n\t")
                        default:
                            if false {
                                var v41 float32 = 3.14159
                                v1 = append(v1, (("a" + "\n\t") + ((("a" + "\n\t") + "foo bar") + "foo bar")))
                                if len(v1) > 0 {
                                    v1[0] = "a"
                                }
                                _ = v41
                            }
                        }
                        var v42 uint8 = 128
                        _ = fi31
                        _ = v42
                    }
                    _ = v30
                } else {
                    fmt.Println(v2)
                }
            } else {
                v43 := GenS6{F0: 42, F1: (-(v2)), F2: 255}
                v2 = (v2 & (-(v2)))
                _ = v43
            }
        } else {
            switch ("hello" + ("" + "foo bar")) {
            case "\n\t":
                for fi44 := 0; fi44 < 20; fi44++ {
                    var v45 int8 = -128
                    _ = fi44
                    _ = v45
                }
            case "":
                v2 += (v2 + 1000)
            case "foo bar":
                switch 143 {
                case -2147483648:
                    for fi46 := 0; fi46 < 2; fi46++ {
                        fmt.Println(fi46, fi46)
                        v2 = -567
                        fmt.Println(fi46, fi46)
                        _ = fi46
                    }
                }
            }
            for fi47 := 0; fi47 < 4; fi47++ {
                var v48 int16 = (v2 % ((v2) | 1))
                if len(v1) > 0 {
                    v1[0] = "\n\t"
                }
                for fi49 := 0; fi49 < 16; fi49++ {
                    v2 = v2
                    _ = fi49
                }
                _ = fi47
                _ = v48
            }
        }
        _ = v27
    }
    if (2 >= 190) {
        var v50 uint8 = 128
        _ = v50
    }
    if (false || (1.5 <= float64(((((((-2.25 * float32(-2.25)) - (((1e10 * ((-2.25 + ((1.5 + ((0.0 * 0.1) / 1.0)) / -2.25)) * (0.0 + 1.5))) + 3.14159) - ((-2.25 + 1.5) / 0.1))) / 3.14159) * 1.5) / ((1.5 / 1.5) - 0.0)) - (-2.25 * 0.0))))) {
        fmt.Println(v2)
    }
    v2 -= v2
    _ = len(v3)
    if (!false) {
        v51 := GenS4{F0: 196, F1: 0}
        v2 -= (v2 + (v2 % ((v2) | 1)))
        _ = v51
    } else {
        fmt.Println(v2)
        if (false && (!(-352 >= -357))) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v52 int = 42
            _ = v52
        } else {
            fmt.Println(v2)
        }
    }
    v2 -= (v2 & (v2 % ((v2) | 1)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
}
