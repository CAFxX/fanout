package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint32; F1 uint16 }
type GenS6 struct { F0 int; F1 uint16; F2 bool }
type GenS7 struct { F0 uint8; F1 uint8 }
type GenS8 struct { F0 uint; F1 int64 }
func genFn0(p0 int16) int32 {
    v5 := []int{128, 0}
    if false {
        switch 128 {
        case -647:
            v6 := GenS1{F0: ("" + ("\n\t" + "foo bar"))}
            _ = v6
        case -426:
            switch (((("hello" + "\n\t") + "\n\t") + "") + ("\n\t" + "")) {
            case "foo bar":
                switch "\n\t" {
                case "\n\t":
                    for fi7 := 0; fi7 < 9; fi7++ {
                        for fi8 := 0; fi8 < 2; fi8++ {
                            v9 := GenS1{F0: "hello"}
                            switch (fi8 >> ((837) & 63)) {
                            case 1:
                                var v10 string = "foo bar"
                                _ = v10
                            }
                            _ = fi8
                            _ = v9
                        }
                        _ = fi7
                    }
                case "a":
                    var v11 uint16 = 42
                    _ = v11
                case "foo bar":
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
            case "hello":
                v12 := GenS1{F0: "foo bar"}
                _ = v12
            default:
                if (0.1 < 1.5) {
                    if (0.1 != ((-2.25 / ((1.5 + 0.0) * 1.5)) / 3.14159)) {
                        p0 = p0
                        var v13 bool = (2 <= 255)
                        _ = v13
                    } else {
                        var v14 bool = (float32(1.5) < (0.1 / (-2.25 - 1.5)))
                        fmt.Println(p0, p0)
                        for fi15 := 0; fi15 < 18; fi15++ {
                            fmt.Println(v14, p0, p0)
                            _ = fi15
                        }
                        _ = v14
                    }
                    var v16 float64 = 0.0
                    if len(v5) > 0 {
                        v5[0] = 127
                    }
                    _ = v16
                } else {
                    var v17 float32 = 0.0
                    for fi18 := 0; fi18 < 15; fi18++ {
                        var v19 float64 = 1.5
                        _ = fi18
                        _ = v19
                    }
                    var v20 float32 = (1e10 - (((v17 - v17) / v17) / 1.5))
                    _ = v17
                    _ = v20
                }
            }
        }
        if false {
            var v21 float64 = 0.0
            _ = v21
        }
        for fi22 := 0; fi22 < 18; fi22++ {
            v23 := GenS1{F0: "hello"}
            _ = fi22
            _ = v23
        }
    } else {
        p0 = p0
        v5 = append(v5, -187)
        if true {
            var v24 uint16 = 255
            if (((-(p0)) == p0) || false) {
                fmt.Println(p0, v24)
                v24 = v24
                var v25 int64 = -280
                _ = v25
            } else {
                for fi26 := 0; fi26 < 20; fi26++ {
                    _ = len(v5)
                    v27 := GenS2{F0: (p0 << ((255) & 63))}
                    _ = fi26
                    _ = v27
                }
                var v28 int16 = p0
                fmt.Println(v28)
                _ = v28
            }
            if true {
                var v29 bool = true
                v30 := GenS3{F0: -719}
                if v29 {
                    var v31 uint32 = 1
                    for fi32 := 0; fi32 < 12; fi32++ {
                        var v33 int64 = 42
                        var v34 float32 = 3.14159
                        if v29 {
                            v33 -= v33
                            v34 += (v34 + ((0.1 + ((float32((3.14159 - 1e10)) - (v34 - ((((-2.25 - ((1e10 * (v34 * 3.14159)) / v34)) * v34) * (v34 - float32(-2.25))) + 3.14159))) * 0.1)) / float32((0.1 - ((0.0 * (0.0 * float64(((v34 / 1.5) + (float32(3.14159) * v34))))) / float64(3.14159))))))
                            for fi35 := 0; fi35 < 10; fi35++ {
                                var v36 int = fi32
                                var v37 int8 = -34
                                _ = fi35
                                _ = v36
                                _ = v37
                            }
                        }
                        _ = fi32
                        _ = v33
                        _ = v34
                    }
                    v38 := GenS4{F0: 100}
                    _ = v31
                    _ = v38
                }
                _ = v29
                _ = v30
            }
            _ = v24
        } else {
            var v39 int16 = 1000
            if (true || ("foo bar" > ("" + (((("a" + (("a" + "foo bar") + "\n\t")) + "hello") + "hello") + "")))) {
                switch "foo bar" {
                case "foo bar":
                    fmt.Println(p0, v39)
                case "\n\t":
                    if true {
                        var v40 float32 = (-2.25 + (0.0 / float32(-2.25)))
                        switch 966 {
                        case 32767:
                            _ = len(v5)
                        case -222:
                            fmt.Println(p0, v40, v40)
                        default:
                            fmt.Println(p0, p0)
                        }
                        _ = v40
                    } else {
                        v41 := GenS5{F0: 2, F1: 0}
                        _ = v41
                    }
                case "hello":
                    var v42 uint64 = 446
                    _ = v42
                }
                var v43 uint8 = 255
                var v44 uint32 = 42
                _ = v43
                _ = v44
            }
            _ = v39
        }
    }
    fmt.Println(p0)
    fmt.Println(p0)
    var v45 uint = 0
    _ = p0
    _ = v45
    _ = v5
    return 42
}

func main() {
    var v1 uint = 807
    var v2 uint64 = 1
    v3 := []int{127, 128, -490, 0}
    var v4 uint64 = 2
    switch 42 {
    case -2147483648:
        if (42 < 127) {
            v46 := GenS6{F0: -2147483648, F1: 65535, F2: (128 < 0)}
            _ = v46
        } else {
            var v47 uint = v1
            if ((1.5 + 3.14159) >= 0.0) {
                if len(v3) > 0 {
                    v3[0] = 128
                }
                v4 = (v2 | (v2 + v2))
                v2 -= 1
            }
            _ = v47
        }
    default:
        switch 0 {
        case 762:
            if (58 >= 1) {
                if len(v3) > 0 {
                    v3[0] = 127
                }
                _ = len(v3)
            }
        case 32767:
            var v48 float32 = 3.14159
            _ = v48
        case -2147483648:
            v49 := GenS2{F0: 48}
            _ = v49
        default:
            var v50 int16 = 402
            _ = v50
        }
    }
    v2 = (v4 + (^(v2)))
    var v51 float64 = float64(3.14159)
    fmt.Println(v1, v4)
    var v52 int16 = 1000
    v52 = v52
    var v53 bool = false
    var v54 int32 = 1
    if len(v3) > 0 {
        v3[0] = 9223372036854775807
    }
    var v55 float64 = v51
    var v56 int32 = genFn0((v52 - (^(v52))))
    fmt.Println(v4)
    v1 = (v1 ^ v1)
    v51 += (float64((0.0 / -2.25)) / (v55 / (v51 / float64((3.14159 / -2.25)))))
    var v57 float64 = (v55 - (float64((3.14159 + (3.14159 - -2.25))) + v51))
    for fi58 := 0; fi58 < 20; fi58++ {
        for fi59 := 0; fi59 < 15; fi59++ {
            v60 := GenS7{F0: uint8(v1), F1: 19}
            v61 := GenS8{F0: 18446744073709551615, F1: int64(v56)}
            _ = fi59
            _ = v60
            _ = v61
        }
        _ = fi58
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v55
    _ = v56
    _ = v57
}
