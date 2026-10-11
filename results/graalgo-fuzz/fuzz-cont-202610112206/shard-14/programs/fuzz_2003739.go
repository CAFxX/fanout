package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 int }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 int16; F1 bool; F2 int8 }
type GenS5 struct { F0 uint16; F1 int32; F2 uint32 }
type GenS6 struct { F0 int64; F1 uint64; F2 int16 }
type GenS7 struct { F0 uint64; F1 int64; F2 int64 }

func main() {
    var v1 uint16 = 355
    var v2 uint8 = 127
    var v3 uint = 80
    var v4 float32 = 0.0
    var v5 int16 = 32767
    v6 := make(map[uint32]string)
    v2 -= (v2 >> ((v3) & 63))
    v7 := []int{-359, 128, 32768}
    fmt.Println(v3, v1)
    for fi8 := 0; fi8 < 20; fi8++ {
        v1 += 32767
        var v9 float32 = ((1.5 * (0.1 + 0.1)) * v4)
        var v10 float64 = ((-2.25 + (float64((v4 - 3.14159)) * (0.1 * 0.1))) + 0.0)
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: 65, F1: 770, F2: 4294967295}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    v12 := GenS1{F0: uint32(v1), F1: -2147483648, F2: 9223372036854775807}
    v5 += (v5 * v5)
    _ = len(v6)
    fmt.Println(v2)
    v3 += v3
    switch 1 {
    case -921:
        if false {
            for fi13 := 0; fi13 < 5; fi13++ {
                for fi14 := 0; fi14 < 9; fi14++ {
                    v5 = -763
                    v15 := GenS1{F0: 911, F1: fi13, F2: 1}
                    _ = fi14
                    _ = v15
                }
                var v16 uint16 = (v1 + v1)
                var v17 int64 = -884
                _ = fi13
                _ = v16
                _ = v17
            }
            for fi18 := 0; fi18 < 8; fi18++ {
                v7 = append(v7, (^(fi18)))
                v19 := GenS2{F0: (v2 & 71), F1: (("hello" + "hello") + "\n\t")}
                for fi20 := 0; fi20 < 5; fi20++ {
                    var v21 uint = 789
                    if (!false) {
                        var v22 float32 = float32(3.14159)
                        if false {
                            if false {
                                if ((!true) && (!(true && true))) {
                                    var v23 int32 = int32(v3)
                                    v5 += (v5 % ((v5) | 1))
                                    if true {
                                        _ = len(v7)
                                        fmt.Println(v4)
                                        fmt.Println(v2)
                                        fmt.Println(v4)
                                    } else {
                                        var v24 int16 = (-(v5))
                                        v25 := GenS3{F0: 32768, F1: 2147483648}
                                        v2 = (v2 % (((v2 | v2)) | 1))
                                        if true {
                                            v6[4294967295] = "hello"
                                            v22 = float32(0.1)
                                        } else {
                                            switch 1 {
                                            case 255:
                                                switch "foo bar" {
                                                case "foo bar":
                                                    v26 := GenS2{F0: (v2 & uint8(v1)), F1: "a"}
                                                    _ = v26
                                                case "":
                                                    var v27 uint8 = (v2 << ((v3) & 63))
                                                    _ = v27
                                                }
                                            default:
                                                var v28 float64 = float64(float32(float64((-2.25 * (v4 * ((1e10 - float32(1.5)) - (((1.5 / float32(1.5)) * ((v22 / -2.25) / ((-2.25 - v4) * (v4 + 1e10)))) - v22)))))))
                                                _ = v28
                                            }
                                            var v29 int16 = v5
                                            if true {
                                                v3 = 2
                                            } else {
                                                v5 = (v29 >> ((v3) & 63))
                                                var v30 uint16 = 32767
                                                _ = v30
                                            }
                                            _ = v29
                                        }
                                        v3 = v21
                                        _ = v24
                                        _ = v25
                                    }
                                    _ = v23
                                } else {
                                    var v31 int16 = v5
                                    _ = v31
                                }
                                var v32 float32 = float32((0.1 + (1.5 / 1e10)))
                                fmt.Println(v3)
                                _ = v32
                            }
                            v4 -= v4
                            if (v2 != v2) {
                                fmt.Println(v4, v21, v22)
                                switch ("\n\t" + "hello") {
                                case "\n\t":
                                    _ = len(v6)
                                case "hello":
                                    v33 := GenS4{F0: 726, F1: false, F2: 49}
                                    _ = v33
                                case "":
                                    switch "foo bar" {
                                    case "a":
                                        v34 := GenS3{F0: 752, F1: 2}
                                        if (1 == -147) {
                                            v35 := GenS5{F0: 32767, F1: 1, F2: 2147483648}
                                            v36 := GenS1{F0: 667, F1: (^(fi20)), F2: 32768}
                                            _ = v35
                                            _ = v36
                                        }
                                        _ = v34
                                    case "\n\t":
                                        v5 = (v5 >> (((v3 << ((v3) & 63))) & 63))
                                    default:
                                        var v37 int = fi20
                                        _ = v37
                                    }
                                default:
                                    v22 = ((v4 * v4) / v22)
                                }
                            } else {
                                if true {
                                    fmt.Println(fi20, fi20, v3)
                                    var v38 bool = ((true && (true || (false && (-9223372036854775808 > 920)))) && false)
                                    switch (("foo bar" + ("\n\t" + "hello")) + "\n\t") {
                                    case "a":
                                        if false {
                                            v21 += (^(v21))
                                        } else {
                                            if v38 {
                                                var v39 uint32 = 2147483648
                                                _ = v39
                                            }
                                            v40 := GenS5{F0: uint16(v21), F1: -2147483648, F2: 42}
                                            switch "hello" {
                                            case "hello":
                                                v38 = (((v2 ^ v2) != (v2 & v2)) || (-128 >= -128))
                                            }
                                            _ = v40
                                        }
                                    case "\n\t":
                                        fmt.Println(fi20, v22)
                                    case "":
                                        v41 := GenS6{F0: 42, F1: 42, F2: (v5 & v5)}
                                        _ = v41
                                    default:
                                        fmt.Println(v21, v4)
                                    }
                                    _ = v38
                                } else {
                                    if true {
                                        var v42 uint32 = 839
                                        var v43 float32 = float32(0.0)
                                        _ = v42
                                        _ = v43
                                    }
                                }
                            }
                        } else {
                            var v44 uint32 = uint32(v2)
                            if len(v7) > 0 {
                                v7[0] = 127
                            }
                            _ = v44
                        }
                        if (false && true) {
                            var v45 string = "\n\t"
                            v46 := v6[65535]
                            _ = v45
                            _ = v46
                        }
                        if true {
                            fmt.Println(v3, v1, v2)
                            fmt.Println(v2)
                        } else {
                            switch ("foo bar" + ("" + ("" + (((((("" + ("a" + "\n\t")) + ((("a" + ("hello" + "\n\t")) + "\n\t") + ("hello" + ""))) + "") + "\n\t") + ((((((("" + (("\n\t" + (("hello" + ("hello" + ("\n\t" + ""))) + ("\n\t" + "hello"))) + ("foo bar" + "a"))) + "hello") + "a") + ((((("" + "") + "\n\t") + ("a" + ("" + ("\n\t" + "foo bar")))) + "\n\t") + ("a" + ("hello" + "hello")))) + "foo bar") + "hello") + ("" + "a"))) + "hello")))) {
                            case "a":
                                switch "hello" {
                                case "":
                                    if false {
                                        var v47 float64 = -2.25
                                        _ = v47
                                    }
                                default:
                                    fmt.Println(fi20, v4, v22)
                                }
                            case "\n\t":
                                var v48 float32 = (3.14159 / (v22 / v4))
                                _ = v48
                            default:
                                var v49 float64 = (0.0 + 0.0)
                                _ = v49
                            }
                        }
                        if true {
                            v1 = v1
                            v6[2147483648] = "a"
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                        } else {
                            if (1 < v2) {
                                v22 -= (3.14159 * (3.14159 * ((float32(0.1) * 3.14159) / float32(1.5))))
                            }
                            v6[2147483648] = ""
                            var v50 int16 = (v5 - v5)
                            _ = v50
                        }
                        if (1 == -960) {
                            fmt.Println(v5)
                            switch fi20 {
                            case 319:
                                if (127 <= -53) {
                                    v51 := GenS3{F0: int(v1), F1: 2147483647}
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                    var v52 uint16 = v1
                                    v53 := GenS7{F0: 2, F1: 0, F2: -9223372036854775808}
                                    _ = v51
                                    _ = v52
                                    _ = v53
                                } else {
                                    var v54 int = 225
                                    v55 := GenS7{F0: 68, F1: -1, F2: 9223372036854775807}
                                    var v56 uint = v21
                                    _ = v54
                                    _ = v55
                                    _ = v56
                                }
                            case -1:
                                var v57 uint8 = (v2 * 234)
                                _ = v57
                            case -2147483648:
                                fmt.Println(v2, fi20)
                            }
                            v58 := v6[47]
                            _ = v58
                        } else {
                            var v59 uint16 = (v1 * 2)
                            if (!false) {
                                v2 = v2
                                _ = len(v6)
                            }
                            var v60 uint = v21
                            _ = v59
                            _ = v60
                        }
                        _ = v22
                    } else {
                        var v61 int64 = 0
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                        fmt.Println(v5, v3)
                        _ = v61
                    }
                    if ("foo bar" < ("foo bar" + ("a" + "a"))) {
                        var v62 float32 = v4
                        v62 = v62
                        _ = v62
                    }
                    v63 := v6[42]
                    _ = fi20
                    _ = v21
                    _ = v63
                }
                _ = fi18
                _ = v19
            }
            for fi64 := 0; fi64 < 10; fi64++ {
                v3 += v3
                _ = fi64
            }
        }
    case 0:
        v6[2147483648] = "\n\t"
    default:
        v5 -= v5
    }
    v2 = v2
    v5 += v5
    v65 := GenS1{F0: 0, F1: 32767, F2: -369}
    switch "hello" {
    case "foo bar":
        var v66 int = -180
        _ = v66
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v65
    _ = v7
}
