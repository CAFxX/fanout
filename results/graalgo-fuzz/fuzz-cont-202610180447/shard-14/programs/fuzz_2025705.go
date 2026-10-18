package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 string; F1 uint32 }
type GenS6 struct { F0 string; F1 string }
type GenS7 struct { F0 int32; F1 int16 }
type GenS8 struct { F0 uint32 }

func main() {
    var v1 float64 = 1.5
    v2 := []uint32{65535, 550, 42, 2147483648}
    v3 := make(map[uint32]string)
    var v4 uint64 = 42
    var v5 bool = ((true || true) && (-128 == -128))
    var v6 uint32 = 2
    v5 = true
    for fi7 := 0; fi7 < 6; fi7++ {
        if false {
            v8 := GenS1{F0: "", F1: 15, F2: 42}
            if v5 {
                v9 := GenS1{F0: "", F1: 127, F2: 255}
                for fi10 := 0; fi10 < 18; fi10++ {
                    v1 += v1
                    switch (fi7 >> ((0) & 63)) {
                    case 42:
                        var v11 uint64 = (v4 - v4)
                        _ = v11
                    case -9223372036854775808:
                        if (false && true) {
                            v6 = v6
                            var v12 uint = 255
                            v5 = (!v5)
                            _ = v12
                        }
                    case -483:
                        if false {
                            if ((fi10 * fi7) <= fi10) {
                                if (-2147483648 <= 0) {
                                    var v13 uint64 = (v4 / ((563) | 1))
                                    if false {
                                        switch fi10 {
                                        case 2147483647:
                                            var v14 string = ""
                                            _ = v14
                                        case 127:
                                            var v15 uint = 4294967295
                                            _ = v15
                                        case 32768:
                                            v6 += v6
                                        default:
                                            v5 = (255 < 127)
                                        }
                                        v1 = (0.1 - (0.1 * v1))
                                    } else {
                                        switch "a" {
                                        case "hello":
                                            v5 = v5
                                            var v16 int = fi10
                                            switch "foo bar" {
                                            case "":
                                                if len(v2) > 0 {
                                                    v2[0] = v6
                                                }
                                            }
                                            _ = v16
                                        case "\n\t":
                                            v1 = v1
                                        default:
                                            v1 += (float64((3.14159 * -2.25)) - float64((((0.0 * (1.5 - float32(v1))) + 0.1) - 1e10)))
                                            if (v5 && v5) {
                                                var v17 float64 = (((v1 * (v1 - 1e10)) / v1) / float64((1e10 + 1e10)))
                                                _ = v17
                                            }
                                            var v18 float32 = (3.14159 + (0.0 / (float32((v1 / (v1 * v1))) / -2.25)))
                                            _ = v18
                                        }
                                        if true {
                                            var v19 uint16 = 255
                                            switch (^(fi10)) {
                                            case 127:
                                                if len(v2) > 0 {
                                                    v2[0] = (^(v6))
                                                }
                                            case 128:
                                                var v20 uint = 18446744073709551615
                                                _ = v20
                                            case -88:
                                                var v21 uint8 = 127
                                                var v22 int8 = 1
                                                fmt.Println(v1)
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                var v23 float64 = -2.25
                                                _ = v21
                                                _ = v22
                                                _ = v23
                                            default:
                                                switch fi7 {
                                                case 393:
                                                    if true {
                                                        var v24 uint16 = 65535
                                                        if v5 {
                                                            v4 += v4
                                                            v24 -= (^(v19))
                                                            switch fi10 {
                                                            case 32767:
                                                                delete(v3, (^(v6)))
                                                            default:
                                                                v5 = v5
                                                            }
                                                        } else {
                                                            v24 -= 0
                                                            delete(v3, (v6 | (v6 << ((68) & 63))))
                                                            var v25 string = ""
                                                            _ = v25
                                                        }
                                                        v26 := GenS1{F0: ("" + "\n\t"), F1: 0, F2: 42}
                                                        _ = v24
                                                        _ = v26
                                                    } else {
                                                        var v27 int = fi7
                                                        _ = v27
                                                    }
                                                default:
                                                    switch "" {
                                                    case "hello":
                                                        var v28 uint = 42
                                                        _ = v28
                                                    case "":
                                                        v6 -= 65535
                                                    default:
                                                        switch "\n\t" {
                                                        case "foo bar":
                                                            v4 += (v4 & (^(v13)))
                                                        }
                                                        switch "\n\t" {
                                                        case "":
                                                            if (-794 != 1) {
                                                                switch ((((("foo bar" + (("hello" + "a") + ("foo bar" + "hello"))) + "") + ("a" + "\n\t")) + ("a" + "")) + ("hello" + "hello")) {
                                                                case "a":
                                                                    v5 = ("a" > "hello")
                                                                }
                                                                var v29 float32 = -2.25
                                                                _ = v29
                                                            } else {
                                                                switch "a" {
                                                                case "a":
                                                                    var v30 uint16 = 728
                                                                    _ = v30
                                                                case "foo bar":
                                                                    delete(v3, v6)
                                                                case "\n\t":
                                                                    v31 := v3[1]
                                                                    _ = v31
                                                                }
                                                            }
                                                        case "a":
                                                            var v32 uint16 = v19
                                                            _ = v32
                                                        default:
                                                            if false {
                                                                fmt.Println(v1, v6)
                                                                v33 := GenS1{F0: ("foo bar" + (("\n\t" + ("\n\t" + "foo bar")) + "a")), F1: 1, F2: 3}
                                                                fmt.Println(v1)
                                                                _ = v33
                                                            }
                                                        }
                                                    }
                                                }
                                            }
                                            _ = v19
                                        } else {
                                            v34 := GenS2{F0: 255}
                                            var v35 uint8 = 1
                                            _ = v34
                                            _ = v35
                                        }
                                    }
                                    var v36 int64 = 42
                                    _ = v13
                                    _ = v36
                                }
                                var v37 uint8 = 127
                                _ = v37
                            }
                            switch (fi7 / ((fi7) | 1)) {
                            case 128:
                                v38 := GenS3{F0: -837, F1: (v4 - v4)}
                                _ = v38
                            case -28:
                                v39 := GenS4{F0: v4}
                                _ = v39
                            }
                            switch fi7 {
                            case 4294967295:
                                if v5 {
                                    if v5 {
                                        switch "hello" {
                                        case "a":
                                            fmt.Println(v4, v1)
                                        case "hello":
                                            var v40 uint8 = 255
                                            _ = v40
                                        case "\n\t":
                                            if v5 {
                                                v6 += (v6 >> ((42) & 63))
                                                v4 = (^(v4))
                                            } else {
                                                v41 := GenS5{F0: "hello", F1: v6}
                                                _ = v41
                                            }
                                        }
                                    }
                                }
                            case 1:
                                if (32767 > 0) {
                                    v42 := GenS6{F0: "hello", F1: "foo bar"}
                                    fmt.Println(v4, v4)
                                    if v5 {
                                        var v43 string = ""
                                        switch (v43 + "foo bar") {
                                        case "foo bar":
                                            v44 := GenS5{F0: "foo bar", F1: v6}
                                            _ = v44
                                        case "a":
                                            v5 = false
                                        case "\n\t":
                                            if v5 {
                                                v1 -= float64((float32(v1) / ((3.14159 / (1e10 * 0.1)) - -2.25)))
                                                v3[v6] = "hello"
                                                v3[v6] = ""
                                                var v45 float32 = -2.25
                                                _ = v45
                                            } else {
                                                fmt.Println(v1)
                                                _ = len(v3)
                                            }
                                        }
                                        v46 := GenS3{F0: -1, F1: v4}
                                        _ = v43
                                        _ = v46
                                    } else {
                                        fmt.Println(v1, v6)
                                    }
                                    _ = v42
                                } else {
                                    switch "foo bar" {
                                    case "":
                                        fmt.Println(v1)
                                        _ = len(v2)
                                        v4 = (v4 - (v4 / ((v4) | 1)))
                                    case "foo bar":
                                        v4 -= (v4 | 42)
                                    default:
                                        var v47 uint32 = (v6 + 952)
                                        _ = v47
                                    }
                                    v6 += 65535
                                    v1 += v1
                                }
                            case -9223372036854775808:
                                var v48 uint32 = (^(v6))
                                _ = v48
                            default:
                                var v49 int16 = 505
                                _ = v49
                            }
                            v5 = (458 > 42)
                            v1 -= 3.14159
                        } else {
                            v50 := GenS3{F0: -1, F1: 126}
                            var v51 uint8 = 0
                            fmt.Println(v51, v4)
                            _ = v50
                            _ = v51
                        }
                    default:
                        v52 := v3[723]
                        v6 -= v6
                        _ = v52
                    }
                    fmt.Println(v1)
                    _ = fi10
                }
                _ = v9
            } else {
                var v53 float32 = -2.25
                var v54 float32 = 1.5
                var v55 float32 = float32(1e10)
                _ = v53
                _ = v54
                _ = v55
            }
            _ = v8
        } else {
            var v56 float64 = 1.5
            v57 := GenS7{F0: 2147483647, F1: -265}
            v4 = 1
            _ = v56
            _ = v57
        }
        switch fi7 {
        case 419:
            if (3.14159 < v1) {
                var v58 int8 = int8(v4)
                _ = v58
            } else {
                var v59 uint8 = 255
                _ = v59
            }
        case 9223372036854775807:
            v60 := GenS5{F0: "\n\t", F1: (v6 & v6)}
            _ = v60
        default:
            for fi61 := 0; fi61 < 14; fi61++ {
                switch "\n\t" {
                case "a":
                    if len(v2) > 0 {
                        v2[0] = (v6 << ((255) & 63))
                    }
                }
                v5 = ((65 >= 2) && (true || v5))
                switch ("hello" + ("a" + "")) {
                case "hello":
                    fmt.Println(fi61)
                case "a":
                    v62 := v3[(^(v6))]
                    _ = v62
                case "\n\t":
                    var v63 uint64 = 590
                    _ = v63
                default:
                    v64 := GenS2{F0: 65}
                    _ = v64
                }
                _ = fi61
            }
        }
        _ = fi7
    }
    var v65 uint64 = 140
    if (!(1000 <= 1)) {
        switch "hello" {
        case "a":
            v1 += 0.0
        case "hello":
            for fi66 := 0; fi66 < 12; fi66++ {
                var v67 float32 = (1e10 / 1.0)
                _ = fi66
                _ = v67
            }
        case "":
            delete(v3, v6)
        }
    } else {
        var v68 uint64 = v4
        _ = v68
    }
    var v69 bool = true
    v70 := GenS8{F0: (^(v6))}
    v71 := v3[uint32(v65)]
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v65:", v65)
    fmt.Println("v69:", v69)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v65
    _ = v69
    _ = v70
    _ = v71
}
