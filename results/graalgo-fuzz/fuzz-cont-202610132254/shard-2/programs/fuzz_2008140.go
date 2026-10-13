package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 uint; F1 bool; F2 int }
type GenS3 struct { F0 int32; F1 bool; F2 uint16 }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int; F1 bool; F2 int64 }

func main() {
    var v1 uint8 = 128
    var v2 int16 = int16(v1)
    var v3 int64 = 9223372036854775807
    var v4 int8 = 42
    v2 -= -1
    if (!false) {
        for fi5 := 0; fi5 < 3; fi5++ {
            for fi6 := 0; fi6 < 17; fi6++ {
                var v7 uint = 0
                _ = fi6
                _ = v7
            }
            if (((("a" + "\n\t") + ("\n\t" + "")) >= (("a" + ("\n\t" + "\n\t")) + "\n\t")) && false) {
                if false {
                    var v8 uint = 88
                    if false {
                        if (fi5 == -2147483648) {
                            var v9 int = (fi5 - (fi5 * (fi5 % ((fi5) | 1))))
                            for fi10 := 0; fi10 < 8; fi10++ {
                                v3 = (-(v3))
                                var v11 uint8 = v1
                                _ = fi10
                                _ = v11
                            }
                            fmt.Println(v4)
                            _ = v9
                        } else {
                            if ((float64(((1.5 * 1e10) * 0.0)) / (float64((0.1 + 0.1)) + 1.5)) > (-2.25 + -2.25)) {
                                switch "hello" {
                                case "foo bar":
                                    fmt.Println(v8, v8, v3)
                                default:
                                    v4 = v4
                                }
                                for fi12 := 0; fi12 < 8; fi12++ {
                                    var v13 int64 = (v3 << (((v8 - v8)) & 63))
                                    _ = fi12
                                    _ = v13
                                }
                                v3 += v3
                            }
                        }
                        for fi14 := 0; fi14 < 4; fi14++ {
                            v15 := GenS1{F0: (v2 ^ (v2 * v2)), F1: 2}
                            var v16 uint16 = 880
                            fmt.Println(fi14)
                            v17 := GenS1{F0: v2, F1: 2}
                            v18 := []int{(fi5 % ((fi5) | 1)), (^(fi5)), 1}
                            if (v8 >= uint(v16)) {
                                switch ((("" + "a") + "\n\t") + "a") {
                                case "\n\t":
                                    if (((3.14159 + 3.14159) >= 0.0) && false) {
                                        var v19 string = "a"
                                        v2 = v2
                                        if false {
                                            fmt.Println(fi14, v4, v16)
                                        }
                                        _ = v19
                                    } else {
                                        var v20 uint16 = (v16 + (v16 / ((v16) | 1)))
                                        _ = v20
                                    }
                                }
                            } else {
                                if false {
                                    var v21 int8 = (-(v4))
                                    if false {
                                        if len(v18) >= 2 {
                                            v18 = v18[1:len(v18)]
                                        }
                                        fmt.Println(v3)
                                        fmt.Println(v21, fi5, v2)
                                    } else {
                                        var v22 int16 = 1000
                                        if len(v18) >= 2 {
                                            v18 = v18[1:len(v18)]
                                        }
                                        _ = v22
                                    }
                                    if true {
                                        var v23 int32 = 1
                                        fmt.Println(v16, v21, v3)
                                        _ = v23
                                    }
                                    _ = v21
                                } else {
                                    var v24 int16 = v2
                                    _ = v24
                                }
                            }
                            v3 += 0
                            var v25 string = "\n\t"
                            _ = fi14
                            _ = v15
                            _ = v16
                            _ = v17
                            _ = v18
                            _ = v25
                        }
                    }
                    if (true || (!true)) {
                        if (1 <= 629) {
                            if false {
                                var v26 int8 = (v4 >> (((v8 * 1)) & 63))
                                var v27 uint = v8
                                var v28 string = "foo bar"
                                _ = v26
                                _ = v27
                                _ = v28
                            }
                        } else {
                            v29 := GenS2{F0: v8, F1: true, F2: int(v1)}
                            _ = v29
                        }
                    } else {
                        fmt.Println(v3, v1, v2)
                    }
                    _ = v8
                } else {
                    v2 += -32768
                    v30 := GenS3{F0: -1, F1: (!((true && ((v4 << ((2) & 63)) == (v4 << ((0) & 63)))) && true)), F2: 2}
                    _ = v30
                }
                var v31 uint8 = 255
                var v32 int32 = 224
                _ = v31
                _ = v32
            } else {
                if false {
                    var v33 int64 = (v3 << ((102) & 63))
                    fmt.Println(v4, v33)
                    if (v4 <= v4) {
                        var v34 int = fi5
                        _ = v34
                    } else {
                        if true {
                            var v35 float32 = (1.5 * 1e10)
                            v36 := GenS4{F0: 197, F1: 0}
                            v37 := []int{fi5, 9223372036854775807, fi5, (fi5 % ((-718) | 1))}
                            _ = v35
                            _ = v36
                            _ = v37
                        }
                        fmt.Println(v1, v2)
                    }
                    _ = v33
                }
                var v38 uint = 660
                var v39 string = "a"
                _ = v38
                _ = v39
            }
            switch ("" + ("hello" + "\n\t")) {
            case "\n\t":
                v4 += (v4 << ((0) & 63))
            case "foo bar":
                v40 := GenS3{F0: 1, F1: (453 == 65535), F2: 42}
                _ = v40
            default:
                v41 := GenS4{F0: (v2 ^ (v2 ^ v2)), F1: v3}
                _ = v41
            }
            _ = fi5
        }
    }
    var v42 int64 = 1
    for fi43 := 0; fi43 < 18; fi43++ {
        v1 -= (v1 + (v1 << ((1) & 63)))
        v44 := GenS1{F0: (v2 % ((v2) | 1)), F1: 606}
        switch fi43 {
        case 32768:
            v45 := []int{-1, 4294967295, fi43}
            _ = v45
        case 42:
            fmt.Println(v42, v4)
        }
        _ = fi43
        _ = v44
    }
    switch ("\n\t" + "a") {
    case "a":
        switch "a" {
        case "hello":
            for fi46 := 0; fi46 < 11; fi46++ {
                for fi47 := 0; fi47 < 2; fi47++ {
                    fmt.Println(v3)
                    _ = fi47
                }
                v48 := GenS5{F0: ("" + ("foo bar" + "hello"))}
                fmt.Println(v42)
                _ = fi46
                _ = v48
            }
        case "\n\t":
            var v49 uint8 = (v1 << ((255) & 63))
            _ = v49
        default:
            var v50 float64 = float64(1e10)
            _ = v50
        }
    case "foo bar":
        fmt.Println(v4)
    case "":
        var v51 uint = 483
        _ = v51
    default:
        if false {
            var v52 string = ("a" + ("foo bar" + "foo bar"))
            _ = v52
        } else {
            var v53 int16 = (v2 + (-(v2)))
            var v54 float32 = float32(1e10)
            if (!true) {
                var v55 uint64 = 878
                _ = v55
            } else {
                if false {
                    if true {
                        var v56 uint32 = 756
                        for fi57 := 0; fi57 < 18; fi57++ {
                            v42 = -1
                            _ = fi57
                        }
                        _ = v56
                    }
                    v58 := GenS1{F0: -378, F1: 42}
                    _ = v58
                }
            }
            _ = v53
            _ = v54
        }
    }
    switch "foo bar" {
    case "":
        switch -2147483648 {
        case -521:
            v59 := []int{-114}
            _ = v59
        case 4294967295:
            v60 := []int{127}
            _ = v60
        default:
            for fi61 := 0; fi61 < 14; fi61++ {
                v4 = (^(v4))
                _ = fi61
            }
        }
    case "\n\t":
        v1 += v1
    }
    var v62 uint = 65535
    var v63 int8 = (v4 >> ((v62) & 63))
    fmt.Println(v1)
    var v64 uint16 = 302
    if false {
        if (42 < v62) {
            for fi65 := 0; fi65 < 11; fi65++ {
                v66 := GenS6{F0: fi65, F1: (false && false), F2: v3}
                _ = fi65
                _ = v66
            }
            v67 := []int{0, 127, 9223372036854775807, -182}
            _ = v67
        } else {
            if true {
                if true {
                    if (!false) {
                        fmt.Println(v1, v3)
                    }
                    var v68 int16 = v2
                    _ = v68
                } else {
                    var v69 int8 = (v63 >> ((513) & 63))
                    _ = v69
                }
            } else {
                switch 9223372036854775807 {
                case -846:
                    if (!(0 >= uint64(v1))) {
                        var v70 uint8 = (v1 + 2)
                        v71 := []int{-239, -9223372036854775808, -866}
                        _ = v70
                        _ = v71
                    }
                case 0:
                    v72 := GenS6{F0: -640, F1: true, F2: (v42 & (v42 | (v42 - v3)))}
                    _ = v72
                case -398:
                    var v73 uint = v62
                    _ = v73
                }
            }
            var v74 uint64 = 0
            _ = v74
        }
    } else {
        v75 := GenS1{F0: (v2 ^ 0), F1: 2}
        _ = v75
    }
    v2 += v2
    var v76 uint8 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v62
    _ = v63
    _ = v64
    _ = v76
}
