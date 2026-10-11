package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 uint64; F2 int8 }
type GenS4 struct { F0 uint16; F1 string }
type GenS5 struct { F0 uint32; F1 int16; F2 uint64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int16; F1 bool }
func genFn0(p0 int) uint8 {
    v5 := []int{0, (p0 * p0), (p0 + p0), p0}
    if false {
        var v6 int8 = 1
        _ = v6
    }
    for fi7 := 0; fi7 < 10; fi7++ {
        for fi8 := 0; fi8 < 7; fi8++ {
            switch ("\n\t" + ("a" + ("\n\t" + "hello"))) {
            case "":
                var v9 bool = false
                _ = v9
            }
            _ = fi8
        }
        v10 := GenS1{F0: -1}
        _ = fi7
        _ = v10
    }
    p0 = p0
    fmt.Println(p0)
    _ = p0
    _ = v5
    return 115
}

func main() {
    var v1 uint8 = 158
    v2 := make(map[uint32]bool)
    var v3 float32 = 0.1
    var v4 uint16 = 0
    if false {
        v1 = (v1 % (((v1 | (v1 ^ (v1 & (^(v1)))))) | 1))
        if false {
            v1 = uint8(v4)
        } else {
            v11 := GenS2{F0: 42}
            _ = v11
        }
        var v12 float32 = (v3 * v3)
        _ = v12
    }
    v3 = (float32(1.5) / 0.1)
    v13 := GenS3{F0: 0, F1: 174, F2: 42}
    v14 := GenS4{F0: v4, F1: "a"}
    v15 := v2[4294967295]
    if ((v15 && v15) || true) {
        var v16 float64 = 0.1
        _ = v16
    }
    fmt.Println(v15, v1, v1)
    for fi17 := 0; fi17 < 18; fi17++ {
        for fi18 := 0; fi18 < 20; fi18++ {
            v1 -= 128
            _ = fi18
        }
        v19 := GenS5{F0: 2, F1: 0, F2: 0}
        _ = fi17
        _ = v19
    }
    for fi20 := 0; fi20 < 6; fi20++ {
        for fi21 := 0; fi21 < 1; fi21++ {
            v15 = true
            switch "" {
            case "\n\t":
                var v22 int8 = -128
                _ = v22
            case "hello":
                if v15 {
                    v23 := GenS4{F0: (v4 >> ((501) & 63)), F1: "\n\t"}
                    if (-1 >= 1) {
                        v3 = v3
                        var v24 string = "foo bar"
                        var v25 uint16 = (v4 % ((v4) | 1))
                        _ = v24
                        _ = v25
                    } else {
                        switch ("foo bar" + ("foo bar" + (((("hello" + "a") + "\n\t") + (("\n\t" + "a") + "\n\t")) + ("\n\t" + "hello")))) {
                        case "a":
                            delete(v2, 1)
                        case "hello":
                            fmt.Println(fi20)
                        case "\n\t":
                            v1 = 2
                        }
                        v3 -= v3
                    }
                    if false {
                        v26 := []int{(fi20 | fi21), fi20}
                        _ = v26
                    }
                    _ = v23
                } else {
                    if (-255 <= 42) {
                        var v27 int32 = -2147483648
                        _ = v27
                    }
                    if v15 {
                        var v28 uint32 = 42
                        _ = v28
                    } else {
                        v15 = (0.0 <= ((float64(v3) / 1.5) - 3.14159))
                        v29 := []int{(fi21 / ((fi20) | 1))}
                        _ = v29
                    }
                    if v15 {
                        v4 = (^(v4))
                    } else {
                        switch fi21 {
                        case 2147483647:
                            if v15 {
                                fmt.Println(fi20)
                            }
                        case 128:
                            var v30 int64 = 9223372036854775807
                            v3 = 1e10
                            v4 -= (v4 / (((v4 | (v4 - (v4 - 255)))) | 1))
                            var v31 int32 = -2147483648
                            fmt.Println(v30, v31)
                            _ = v30
                            _ = v31
                        }
                        if (1 > 4294967295) {
                            v32 := v2[2147483648]
                            v33 := GenS4{F0: (v4 % ((v4) | 1)), F1: "a"}
                            _ = v32
                            _ = v33
                        }
                    }
                    v15 = (!(328 >= 0))
                    if (v15 && (true && (true && ((-2.25 > ((float64(v3) + 0.0) + float64((v3 * (-2.25 + 3.14159))))) || ((1.5 * -2.25) >= (((((3.14159 * (-2.25 / -2.25)) * (0.1 * 0.0)) / 1.0) / ((3.14159 - 1e10) / 3.14159)) - 0.1)))))) {
                        var v34 int8 = -1
                        v35 := GenS3{F0: -2147483648, F1: 484, F2: (^(v34))}
                        if v15 {
                            if (127 <= v1) {
                                switch ("\n\t" + "hello") {
                                case "foo bar":
                                    v34 = v34
                                case "\n\t":
                                    fmt.Println(fi20, v34)
                                case "hello":
                                    var v36 float64 = -2.25
                                    _ = v36
                                default:
                                    v37 := GenS2{F0: 42}
                                    _ = v37
                                }
                                switch (("\n\t" + ("\n\t" + (("\n\t" + "\n\t") + ""))) + "foo bar") {
                                case "hello":
                                    var v38 string = "a"
                                    _ = v38
                                case "":
                                    var v39 int = (fi20 | fi21)
                                    _ = v39
                                case "foo bar":
                                    if v15 {
                                        v40 := GenS2{F0: 42}
                                        v41 := v2[65535]
                                        var v42 string = ("\n\t" + ((("a" + "\n\t") + "foo bar") + ("\n\t" + ("\n\t" + ""))))
                                        _ = v40
                                        _ = v41
                                        _ = v42
                                    } else {
                                        fmt.Println(v15, v34, fi20)
                                        var v43 int16 = -554
                                        switch "foo bar" {
                                        case "a":
                                            delete(v2, 4294967295)
                                        case "hello":
                                            v44 := GenS6{F0: (v4 << ((2) & 63))}
                                            fmt.Println(v34)
                                            _ = v44
                                        }
                                        _ = v43
                                    }
                                }
                            }
                        } else {
                            var v45 float32 = v3
                            _ = v45
                        }
                        _ = len(v2)
                        _ = v34
                        _ = v35
                    } else {
                        fmt.Println(v4, v3)
                        var v46 uint = 707
                        v46 += (v46 / ((v46) | 1))
                        _ = v46
                    }
                }
            default:
                v1 = 128
            }
            fmt.Println(v3, v1)
            _ = fi21
        }
        _ = fi20
    }
    for fi47 := 0; fi47 < 11; fi47++ {
        switch (fi47 >> ((759) & 63)) {
        case 0:
            for fi48 := 0; fi48 < 16; fi48++ {
                v15 = true
                v15 = (v15 || (v15 || (v15 && (0 < 451))))
                v4 = (v4 | (v4 * (^(v4))))
                _ = fi48
            }
        case -50:
            fmt.Println(v15)
        }
        var v49 uint16 = (v4 / ((uint16(v1)) | 1))
        var v50 int = fi47
        _ = fi47
        _ = v49
        _ = v50
    }
    for fi51 := 0; fi51 < 19; fi51++ {
        v52 := GenS6{F0: v4}
        v53 := GenS2{F0: 42}
        if ((2 > 42) && true) {
            var v54 int64 = -1
            var v55 int16 = 269
            v56 := []int{(fi51 & (fi51 % ((2147483647) | 1))), (fi51 ^ fi51), fi51, fi51}
            _ = v54
            _ = v55
            _ = v56
        }
        _ = fi51
        _ = v52
        _ = v53
    }
    var v57 float64 = 1e10
    fmt.Println(v1, v1)
    if (42 == 1) {
        v58 := GenS1{F0: 42}
        switch -2147483648 {
        case 128:
            fmt.Println(v57, v57)
        case -2147483648:
            v4 += (v4 - (v4 / ((706) | 1)))
        }
        fmt.Println(v15)
        _ = v58
    } else {
        switch "hello" {
        case "":
            if true {
                v15 = (v15 || true)
            } else {
                v15 = (v15 || (false && v15))
            }
        case "a":
            var v59 uint16 = v4
            _ = v59
        default:
            v1 = (v1 + (v1 ^ (^(v1))))
        }
        for fi60 := 0; fi60 < 9; fi60++ {
            var v61 int32 = 42
            _ = fi60
            _ = v61
        }
        for fi62 := 0; fi62 < 5; fi62++ {
            v57 -= 1.5
            _ = fi62
        }
    }
    fmt.Println(v15, v3, v15)
    if v15 {
        fmt.Println(v15, v4, v1)
        v63 := GenS7{F0: 42, F1: false}
        v57 = v57
        _ = v63
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v57
}
