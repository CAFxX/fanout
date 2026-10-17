package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string; F1 int64; F2 int32 }
type GenS6 struct { F0 int8; F1 uint }

func main() {
    var v1 string = "\n\t"
    var v2 string = ("\n\t" + "a")
    v3 := make(map[int]string)
    switch v1 {
    case "":
        v4 := GenS1{F0: (0 > -262), F1: 734}
        _ = v4
    case "\n\t":
        var v5 float64 = (float64(0.1) / float64(3.14159))
        _ = v5
    }
    v2 = (v2 + v2)
    switch v1 {
    case "hello":
        v2 = "hello"
    case "\n\t":
        for fi6 := 0; fi6 < 20; fi6++ {
            var v7 uint16 = 0
            _ = fi6
            _ = v7
        }
    case "foo bar":
        fmt.Println(v2, v1)
    }
    var v8 float32 = float32((1.5 / ((1.5 / ((3.14159 / float64(-2.25)) - 1e10)) - -2.25)))
    for fi9 := 0; fi9 < 6; fi9++ {
        v2 = "a"
        var v10 uint16 = 195
        _ = fi9
        _ = v10
    }
    fmt.Println(v1)
    fmt.Println(v1, v2, v8)
    var v11 uint8 = 2
    v12 := GenS1{F0: false, F1: 42}
    var v13 int = -19
    v2 = v1
    var v14 uint64 = 42
    if false {
        var v15 int = v13
        _ = v15
    } else {
        switch (v2 + (("a" + v2) + "a")) {
        case "\n\t":
            for fi16 := 0; fi16 < 5; fi16++ {
                v17 := GenS2{F0: (v11 & (v11 | (v11 >> ((255) & 63))))}
                _ = fi16
                _ = v17
            }
        case "a":
            if true {
                v18 := GenS3{F0: int16(v13), F1: 127}
                fmt.Println(v2, v11, v8)
                v14 = v14
                _ = v18
            }
        default:
            v11 -= (v11 | 255)
        }
        _ = len(v3)
    }
    switch (v1 + (v2 + "a")) {
    case "hello":
        var v19 float64 = -2.25
        _ = v19
    case "":
        for fi20 := 0; fi20 < 7; fi20++ {
            if ((-2.25 + 3.14159) >= 1e10) {
                var v21 int64 = -874
                if (670 < 523) {
                    v22 := []int{}
                    _ = v22
                } else {
                    var v23 int64 = (v21 << ((65535) & 63))
                    var v24 uint32 = 4294967295
                    var v25 uint = 0
                    _ = v23
                    _ = v24
                    _ = v25
                }
                v11 += v11
                _ = v21
            } else {
                var v26 string = "hello"
                if true {
                    fmt.Println(v26, v1)
                    if (!(177 != -32768)) {
                        v27 := v3[fi20]
                        _ = v27
                    }
                }
                _ = v26
            }
            v28 := []int{}
            _ = fi20
            _ = v28
        }
    case "\n\t":
        for fi29 := 0; fi29 < 13; fi29++ {
            v11 = 128
            v14 = 0
            _ = fi29
        }
    }
    switch -11 {
    case -1:
        v30 := []int{}
        _ = v30
    default:
        v8 = v8
    }
    switch "a" {
    case "\n\t":
        var v31 uint8 = 127
        _ = v31
    case "foo bar":
        _ = len(v3)
    default:
        fmt.Println(v14, v1)
    }
    for fi32 := 0; fi32 < 14; fi32++ {
        delete(v3, fi32)
        if true {
            var v33 int8 = 42
            _ = v33
        } else {
            for fi34 := 0; fi34 < 18; fi34++ {
                var v35 bool = (!(!(v14 < 42)))
                v36 := GenS2{F0: uint8(v14)}
                _ = fi34
                _ = v35
                _ = v36
            }
            if true {
                switch v2 {
                case "foo bar":
                    switch v1 {
                    case "a":
                        fmt.Println(v11)
                    case "hello":
                        if false {
                            for fi37 := 0; fi37 < 15; fi37++ {
                                v38 := GenS4{F0: (v11 % ((127) | 1))}
                                fmt.Println(v2, v11)
                                v39 := GenS5{F0: v2, F1: 0, F2: -841}
                                if false {
                                    var v40 int16 = 0
                                    fmt.Println(fi32, v2)
                                    var v41 int64 = 0
                                    _ = v40
                                    _ = v41
                                } else {
                                    switch "foo bar" {
                                    case "a":
                                        switch 127 {
                                        case 42:
                                            v14 = (v14 ^ 0)
                                            var v42 uint64 = v14
                                            _ = v42
                                        case -790:
                                            v2 = v2
                                        case 128:
                                            fmt.Println(v1, v11, fi37)
                                            if (552 == 1) {
                                                var v43 uint32 = 2147483648
                                                _ = len(v3)
                                                _ = v43
                                            }
                                            if (42 >= -78) {
                                                v14 -= (v14 / ((v14) | 1))
                                            } else {
                                                v11 -= v11
                                                v44 := GenS3{F0: 16, F1: 127}
                                                switch v1 {
                                                case "a":
                                                    v45 := GenS3{F0: 1, F1: int8(v11)}
                                                    _ = v45
                                                case "\n\t":
                                                    v46 := GenS3{F0: -1, F1: 35}
                                                    _ = v46
                                                case "":
                                                    var v47 int64 = -98
                                                    v48 := GenS2{F0: v11}
                                                    _ = v47
                                                    _ = v48
                                                }
                                                _ = v44
                                            }
                                        default:
                                            fmt.Println(v2, v13, fi37)
                                            fmt.Println(v1, fi32)
                                        }
                                    case "foo bar":
                                        var v49 int32 = 1
                                        _ = v49
                                    case "\n\t":
                                        v1 = "foo bar"
                                    }
                                    if false {
                                        delete(v3, 973)
                                    }
                                    if (160 == 409) {
                                        v50 := v3[(v13 << ((391) & 63))]
                                        if (!(!false)) {
                                            var v51 uint8 = v11
                                            switch v50 {
                                            case "\n\t":
                                                switch 32767 {
                                                case 0:
                                                    var v52 float32 = (v8 * -2.25)
                                                    _ = v52
                                                }
                                            case "hello":
                                                var v53 string = ("" + v1)
                                                _ = v53
                                            }
                                            v54 := v3[fi37]
                                            var v55 float64 = (float64((v8 - float32(float64(v8)))) * 0.1)
                                            _ = v51
                                            _ = v54
                                            _ = v55
                                        }
                                        _ = v50
                                    }
                                    var v56 int = fi32
                                    var v57 int32 = -41
                                    v56 = v13
                                    _ = v56
                                    _ = v57
                                }
                                _ = fi37
                                _ = v38
                                _ = v39
                            }
                        }
                    }
                }
                var v58 string = "foo bar"
                _ = v58
            } else {
                fmt.Println(v2, v8, v2)
            }
            if (!false) {
                var v59 float32 = (v8 + -2.25)
                for fi60 := 0; fi60 < 6; fi60++ {
                    v2 = v2
                    _ = fi60
                }
                fmt.Println(v8)
                _ = v59
            } else {
                var v61 uint64 = v14
                v62 := GenS3{F0: 362, F1: 127}
                _ = v61
                _ = v62
            }
        }
        v63 := GenS6{F0: 42, F1: 201}
        _ = fi32
        _ = v63
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v8
}
