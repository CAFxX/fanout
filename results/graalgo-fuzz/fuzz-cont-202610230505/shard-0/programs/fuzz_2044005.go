package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 string; F1 string; F2 uint }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 string; F1 string }
type GenS5 struct { F0 int8; F1 uint8 }
type GenS6 struct { F0 int32; F1 uint8; F2 int16 }
func genFn0(p0 uint32, p1 uint64, p2 uint64) int8 {
    v4 := []int{}
    p1 -= p2
    _ = len(v4)
    var v5 float64 = float64((3.14159 / (3.14159 / 1.5)))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 117
}

func main() {
    var v1 int = -1
    var v2 uint = 42
    var v3 bool = (790 <= 65535)
    v6 := GenS1{F0: 42, F1: 65535}
    v7 := []int{v1, (v1 >> (((v2 >> (((v2 << ((v2) & 63))) & 63))) & 63)), v1}
    fmt.Println(v3, v2)
    if ((!v3) && (42 <= 2)) {
        v8 := GenS2{F0: "foo bar", F1: "hello", F2: (v2 & 878)}
        v9 := GenS2{F0: "\n\t", F1: "foo bar", F2: v2}
        _ = v8
        _ = v9
    } else {
        var v10 int32 = -2147483648
        var v11 bool = v3
        var v12 bool = ((("\n\t" + "hello") < (("\n\t" + "\n\t") + "")) && v3)
        _ = v10
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        switch v1 {
        case 1:
            for fi14 := 0; fi14 < 5; fi14++ {
                var v15 string = (("\n\t" + "") + ("hello" + "foo bar"))
                v16 := GenS2{F0: "\n\t", F1: (v15 + v15), F2: v2}
                _ = fi14
                _ = v15
                _ = v16
            }
        case 909:
            var v17 string = ("foo bar" + ("" + ("foo bar" + "foo bar")))
            _ = v17
        case -2147483648:
            var v18 int8 = 127
            _ = v18
        }
        v1 = -9223372036854775808
        v19 := GenS3{F0: 127, F1: true}
        _ = fi13
        _ = v19
    }
    if v3 {
        v2 -= (v2 ^ (^(v2)))
    } else {
        v2 += (v2 << (((v2 - (v2 >> (((^(v2))) & 63)))) & 63))
    }
    switch v1 {
    case 199:
        switch (v1 / ((v1) | 1)) {
        case -9223372036854775808:
            v20 := GenS1{F0: -955, F1: 42}
            _ = v20
        default:
            v21 := GenS1{F0: -691, F1: 80}
            _ = v21
        }
    }
    v22 := GenS2{F0: "hello", F1: "\n\t", F2: v2}
    for fi23 := 0; fi23 < 6; fi23++ {
        var v24 uint8 = 1
        for fi25 := 0; fi25 < 13; fi25++ {
            if false {
                v3 = (2147483647 > 42)
                switch fi23 {
                case -282:
                    switch "hello" {
                    case "hello":
                        fmt.Println(v3, fi23)
                    case "":
                        v1 = -2147483648
                        if (251 > 42) {
                            v7 = append(v7, (-(v1)))
                            v26 := GenS1{F0: 42, F1: uint16(v24)}
                            _ = v26
                        }
                    case "foo bar":
                        fmt.Println(v2)
                    default:
                        var v27 int16 = -32768
                        _ = v27
                    }
                case 523:
                    v28 := GenS4{F0: "a", F1: ("" + "hello")}
                    _ = v28
                }
            } else {
                switch (v1 | (fi25 << (((v2 / ((v2) | 1))) & 63))) {
                case 128:
                    fmt.Println(v1)
                default:
                    _ = len(v7)
                }
                switch ("foo bar" + "\n\t") {
                case "hello":
                    if v3 {
                        fmt.Println(fi23, v24, v2)
                        if v3 {
                            _ = len(v7)
                            switch fi23 {
                            case -576:
                                var v29 int16 = -1
                                switch (fi25 - -9223372036854775808) {
                                case -9223372036854775808:
                                    fmt.Println(v2, v24, v29)
                                case -1:
                                    var v30 uint16 = 139
                                    _ = v30
                                default:
                                    if v3 {
                                        fmt.Println(v29, v3, fi25)
                                        var v31 string = ""
                                        v32 := GenS5{F0: -128, F1: (v24 ^ v24)}
                                        _ = v31
                                        _ = v32
                                    }
                                }
                                var v33 uint16 = 899
                                _ = v29
                                _ = v33
                            case 32768:
                                if (!true) {
                                    var v34 int16 = 32767
                                    _ = v34
                                } else {
                                    v35 := GenS1{F0: -390, F1: 835}
                                    if ((!(v3 || (1 == 0))) || v3) {
                                        if ((!((v3 || v3) || (!v3))) && true) {
                                            var v36 string = (((("a" + ("\n\t" + (((("" + "hello") + "a") + ("a" + ("" + "a"))) + "foo bar"))) + ("hello" + ("foo bar" + ("hello" + ("a" + "hello"))))) + "foo bar") + "hello")
                                            _ = v36
                                        }
                                    } else {
                                        v2 += v2
                                        var v37 int64 = 0
                                        _ = v37
                                    }
                                    if v3 {
                                        if (-1 == int32(v24)) {
                                            _ = len(v7)
                                        }
                                        v7 = append(v7, 9223372036854775807)
                                        fmt.Println(v3, v1, v3)
                                    } else {
                                        fmt.Println(fi25)
                                        fmt.Println(v3)
                                    }
                                    _ = v35
                                }
                            case -283:
                                v7 = append(v7, 127)
                                _ = len(v7)
                                v38 := GenS2{F0: ("a" + ""), F1: ("" + "foo bar"), F2: v2}
                                if false {
                                    v3 = v3
                                    var v39 uint = (v2 ^ (v2 | (v2 >> ((v2) & 63))))
                                    v2 += v39
                                    _ = v39
                                }
                                v3 = false
                                switch "" {
                                case "hello":
                                    v7 = append(v7, (-(v1)))
                                case "":
                                    var v40 string = ("\n\t" + (((("" + "a") + ("hello" + "hello")) + "a") + "a"))
                                    _ = v40
                                }
                                v1 += 32768
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                                if true {
                                    if v3 {
                                        if v3 {
                                            var v41 int64 = 0
                                            _ = v41
                                        }
                                    }
                                    var v42 string = "a"
                                    if len(v7) > 0 {
                                        v7[0] = fi25
                                    }
                                    _ = v42
                                }
                                _ = v38
                            }
                            v24 += v24
                        } else {
                            v43 := GenS4{F0: "a", F1: "a"}
                            v1 = fi25
                            if true {
                                if len(v7) > 0 {
                                    v7[0] = 32768
                                }
                            } else {
                                var v44 uint64 = 162
                                _ = v44
                            }
                            v45 := GenS6{F0: -2147483648, F1: v24, F2: 286}
                            _ = v43
                            _ = v45
                        }
                        fmt.Println(v24)
                    }
                default:
                    v2 -= 65535
                }
                v46 := GenS2{F0: "", F1: "\n\t", F2: (v2 & (v2 / (((v2 % (((^(v2))) | 1))) | 1)))}
                v24 = (^(v24))
                fmt.Println(v2, v3)
                var v47 uint8 = 2
                if (!v3) {
                    fmt.Println(fi25)
                    if v3 {
                        v1 = (fi23 - (v1 % ((fi23) | 1)))
                    }
                }
                _ = v46
                _ = v47
            }
            switch ("" + ("foo bar" + ("a" + ("\n\t" + "a")))) {
            case "":
                v48 := GenS1{F0: int32(v2), F1: 60}
                _ = v48
            case "\n\t":
                fmt.Println(v2, v3, v24)
            default:
                v49 := GenS4{F0: "\n\t", F1: "a"}
                _ = v49
            }
            _ = fi25
        }
        v24 -= 255
        _ = fi23
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v6
    _ = v7
}
