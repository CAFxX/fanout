package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint32; F1 uint64; F2 string }
type GenS3 struct { F0 uint16; F1 uint8; F2 int64 }
type GenS4 struct { F0 bool; F1 int16; F2 uint8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int8; F1 string }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]int)
    var v3 bool = true
    v4 := make(map[int]bool)
    var v5 int64 = -9223372036854775808
    var v6 uint = 0
    fmt.Println(v5)
    switch 128 {
    case 621:
        v7 := GenS1{F0: false, F1: -671, F2: 42}
        _ = v7
    case 2147483647:
        v3 = false
    default:
        for fi8 := 0; fi8 < 12; fi8++ {
            var v9 bool = false
            v10 := GenS2{F0: 2147483648, F1: 166, F2: "\n\t"}
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    v5 = v5
    v6 = v6
    v11 := GenS3{F0: 219, F1: 67, F2: (v5 >> (((v6 / (((^(v6))) | 1))) & 63))}
    if (!true) {
        switch ("a" + ("foo bar" + "")) {
        case "hello":
            fmt.Println(v3, v6)
        case "\n\t":
            switch -9223372036854775808 {
            case -2147483648:
                v3 = (42 <= 2147483648)
            case 32767:
                var v12 int16 = 1
                _ = v12
            default:
                var v13 float32 = 0.0
                _ = v13
            }
        case "a":
            var v14 float32 = (0.0 + ((((0.1 * (float32(1e10) / 1.0)) - (float32((0.1 + (0.1 + (1.5 + 3.14159)))) - ((0.0 / 0.1) / 1e10))) * float32(0.0)) * ((((0.1 / 0.1) / (-2.25 - (3.14159 / -2.25))) * 1.5) - float32(-2.25))))
            _ = v14
        }
        if v3 {
            fmt.Println(v3)
            v15 := []int{159, -2147483648}
            if (!(!(32767 <= -193))) {
                switch "\n\t" {
                case "\n\t":
                    v5 = (v5 * 1)
                case "":
                    _ = len(v4)
                }
                for fi16 := 0; fi16 < 16; fi16++ {
                    var v17 uint64 = 4294967295
                    if true {
                        v6 -= v6
                    }
                    v18 := GenS4{F0: (v3 && v3), F1: 349, F2: 1}
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                for fi19 := 0; fi19 < 11; fi19++ {
                    if v3 {
                        v5 -= (v5 & v5)
                    }
                    var v20 string = "hello"
                    for fi21 := 0; fi21 < 16; fi21++ {
                        v20 = (v20 + ("foo bar" + ("\n\t" + v20)))
                        _ = fi21
                    }
                    _ = fi19
                    _ = v20
                }
            } else {
                v22 := GenS3{F0: 386, F1: 255, F2: (v5 << ((v6) & 63))}
                v23 := GenS4{F0: (v3 || false), F1: -366, F2: 254}
                fmt.Println(v3, v3)
                _ = v22
                _ = v23
            }
            _ = v15
        } else {
            fmt.Println(v5)
            switch -2147483648 {
            case 128:
                if (9223372036854775807 > 2147483647) {
                    v24 := GenS2{F0: 4294967295, F1: 2, F2: ""}
                    switch 42 {
                    case 32768:
                        switch "a" {
                        case "hello":
                            if ((-1 < 2147483647) || (!(1 != 66))) {
                                v5 = v5
                                for fi25 := 0; fi25 < 2; fi25++ {
                                    for fi26 := 0; fi26 < 14; fi26++ {
                                        v27 := GenS3{F0: 1, F1: 255, F2: 1}
                                        fmt.Println(v5)
                                        if ((3.14159 / 1.0) == (1.5 * 1e10)) {
                                            switch "\n\t" {
                                            case "a":
                                                var v28 uint32 = 0
                                                _ = v28
                                            case "hello":
                                                v29 := GenS1{F0: false, F1: 279, F2: 4294967295}
                                                _ = v29
                                            case "":
                                                switch ((("hello" + ("\n\t" + ("hello" + "a"))) + "\n\t") + "foo bar") {
                                                case "\n\t":
                                                    v5 += -1
                                                case "foo bar":
                                                    if v3 {
                                                        v30 := v1[65535]
                                                        v3 = false
                                                        var v31 uint64 = 18446744073709551615
                                                        if v3 {
                                                            if v3 {
                                                                var v32 uint64 = (v31 << ((269) & 63))
                                                                if true {
                                                                    var v33 uint = 1
                                                                    _ = v33
                                                                } else {
                                                                    v30 += (fi25 + v30)
                                                                    v34 := []int{32768, (v30 / (((fi25 % (((v30 - (v30 << ((v6) & 63)))) | 1))) | 1)), v30, 42}
                                                                    _ = v34
                                                                }
                                                                var v35 int64 = -829
                                                                fmt.Println(v3)
                                                                _ = v32
                                                                _ = v35
                                                            }
                                                            fmt.Println(v3, v30, v6)
                                                            var v36 uint8 = 128
                                                            _ = v36
                                                        }
                                                        _ = v30
                                                        _ = v31
                                                    }
                                                default:
                                                    var v37 bool = v3
                                                    _ = v37
                                                }
                                            }
                                            v5 -= (v5 ^ (v5 << ((v6) & 63)))
                                        } else {
                                            v38 := v2[1]
                                            v39 := GenS1{F0: v3, F1: -1, F2: 65535}
                                            _ = v38
                                            _ = v39
                                        }
                                        _ = fi26
                                        _ = v27
                                    }
                                    v40 := GenS3{F0: 301, F1: 2, F2: v5}
                                    switch fi25 {
                                    case -431:
                                        if (!v3) {
                                            v5 += v5
                                            var v41 uint16 = 1
                                            var v42 uint16 = v41
                                            _ = v41
                                            _ = v42
                                        }
                                    }
                                    _ = fi25
                                    _ = v40
                                }
                            }
                        case "a":
                            for fi43 := 0; fi43 < 4; fi43++ {
                                for fi44 := 0; fi44 < 15; fi44++ {
                                    if ((0.1 + 3.14159) != (float32(1.5) / (1.5 + 0.0))) {
                                        v45 := GenS5{F0: 1}
                                        if false {
                                            v6 -= 44
                                            var v46 int8 = 1
                                            _ = v46
                                        }
                                        fmt.Println(fi44, v5, v3)
                                        _ = v45
                                    } else {
                                        fmt.Println(v3, v5, v3)
                                    }
                                    v6 = v6
                                    v6 = v6
                                    _ = fi44
                                }
                                v47 := GenS2{F0: 0, F1: 4294967295, F2: ("\n\t" + ("foo bar" + ""))}
                                _ = fi43
                                _ = v47
                            }
                        default:
                            if v3 {
                                var v48 float64 = 0.1
                                for fi49 := 0; fi49 < 7; fi49++ {
                                    fmt.Println(v5, v48)
                                    var v50 uint16 = 255
                                    _ = fi49
                                    _ = v50
                                }
                                delete(v1, uint32(v6))
                                _ = v48
                            } else {
                                switch ("foo bar" + "") {
                                case "foo bar":
                                    delete(v2, 512)
                                case "":
                                    switch 2147483647 {
                                    case 229:
                                        fmt.Println(v6)
                                    case -734:
                                        switch "" {
                                        case "":
                                            v6 += 65535
                                        case "\n\t":
                                            if (!true) {
                                                var v51 string = "a"
                                                _ = v51
                                            }
                                        }
                                    default:
                                        v3 = v3
                                    }
                                case "\n\t":
                                    fmt.Println(v3, v6, v5)
                                default:
                                    _ = len(v4)
                                }
                                for fi52 := 0; fi52 < 1; fi52++ {
                                    var v53 int64 = (^(v5))
                                    v53 = v53
                                    _ = fi52
                                    _ = v53
                                }
                                var v54 int64 = -772
                                _ = v54
                            }
                        }
                    case 42:
                        var v55 int32 = -2147483648
                        _ = v55
                    case -190:
                        switch "" {
                        case "a":
                            if ((!v3) && true) {
                                _ = len(v2)
                            }
                        case "hello":
                            if v3 {
                                if true {
                                    v6 = v6
                                    v56 := GenS5{F0: 42}
                                    _ = v56
                                }
                            } else {
                                v5 -= v5
                                v6 = v6
                                v6 -= v6
                            }
                        case "":
                            fmt.Println(v5, v6, v6)
                        default:
                            fmt.Println(v6, v3, v6)
                        }
                    }
                    _ = v24
                } else {
                    v57 := GenS6{F0: 127, F1: (("a" + "\n\t") + "\n\t")}
                    _ = v57
                }
            case -2147483648:
                v6 = (v6 + 1)
            case 0:
                v58 := GenS5{F0: 1}
                _ = v58
            default:
                for fi59 := 0; fi59 < 20; fi59++ {
                    fmt.Println(fi59, v5)
                    v5 += v5
                    var v60 int8 = -128
                    _ = fi59
                    _ = v60
                }
            }
            var v61 uint8 = 2
            _ = v61
        }
    }
    var v62 bool = v3
    fmt.Println(v5, v3)
    var v63 float64 = 3.14159
    if (32767 >= -1) {
        for fi64 := 0; fi64 < 16; fi64++ {
            v65 := GenS6{F0: -1, F1: ("a" + "hello")}
            for fi66 := 0; fi66 < 17; fi66++ {
                v63 -= float64((-2.25 / -2.25))
                _ = fi66
            }
            _ = fi64
            _ = v65
        }
        v3 = (("foo bar" > "") || (32767 == 65535))
    }
    delete(v1, 2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v62
    _ = v63
}
