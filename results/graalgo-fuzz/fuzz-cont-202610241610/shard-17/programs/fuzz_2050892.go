package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64; F1 string }
type GenS5 struct { F0 bool }
func genFn0(p0 uint64, p1 int) int8 {
    if (((("hello" + "a") + "") + (("a" + ("\n\t" + "foo bar")) + "foo bar")) > "a") {
        var v4 int16 = 32767
        fmt.Println(v4, p1)
        v4 -= v4
        _ = v4
    } else {
        var v5 float64 = 0.1
        p0 += (^(p0))
        p0 = p0
        _ = v5
    }
    for fi6 := 0; fi6 < 2; fi6++ {
        switch -692 {
        case 4294967295:
            var v7 uint64 = p0
            _ = v7
        default:
            p0 -= (p0 | 42)
        }
        _ = fi6
    }
    p0 -= 4294967295
    _ = p0
    _ = p1
    return -117
}
func genFn1(p0 uint16, p1 string, p2 uint64) int16 {
    fmt.Println(p0, p0, p1)
    for fi8 := 0; fi8 < 19; fi8++ {
        v9 := GenS1{F0: 0, F1: 2147483647}
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return 188
}

func main() {
    var v1 int16 = 42
    var v2 string = ("a" + "foo bar")
    var v3 uint32 = 4294967295
    if true {
        if (491 <= 0) {
            v3 = (v3 | v3)
            for fi10 := 0; fi10 < 6; fi10++ {
                v11 := []int{-2147483648, 32767, fi10}
                var v12 string = (v2 + v2)
                _ = fi10
                _ = v11
                _ = v12
            }
        }
    } else {
        switch v2 {
        case "":
            switch 42 {
            case -535:
                v13 := GenS1{F0: -729, F1: (v3 & 65535)}
                _ = v13
            case 9223372036854775807:
                switch 1 {
                case 32767:
                    for fi14 := 0; fi14 < 3; fi14++ {
                        for fi15 := 0; fi15 < 6; fi15++ {
                            if false {
                                if (!(!false)) {
                                    var v16 string = ("foo bar" + v2)
                                    _ = v16
                                } else {
                                    if (255 <= 2) {
                                        v17 := GenS2{F0: 11}
                                        _ = v17
                                    }
                                    if false {
                                        fmt.Println(v3, v1)
                                        if false {
                                            fmt.Println(v3, v2)
                                        }
                                        v18 := []int{32767, fi14, (-(fi14))}
                                        _ = v18
                                    }
                                }
                                v3 -= v3
                                v19 := []int{761, fi15, 128, fi14}
                                _ = v19
                            }
                            _ = fi15
                        }
                        fmt.Println(v3, v3, v3)
                        v3 = v3
                        _ = fi14
                    }
                case 629:
                    var v20 uint16 = 65535
                    _ = v20
                }
            default:
                v2 = (v2 + v2)
            }
        case "foo bar":
            if false {
                var v21 int64 = 1
                v22 := []int{}
                var v23 string = v2
                _ = v21
                _ = v22
                _ = v23
            }
        }
    }
    switch v2 {
    case "hello":
        v24 := GenS1{F0: genFn1(2, (v2 + v2), 1), F1: (v3 + (v3 * (^(v3))))}
        _ = v24
    case "":
        if true {
            v25 := GenS2{F0: 42}
            var v26 int32 = 1
            _ = v25
            _ = v26
        } else {
            var v27 uint16 = 42
            _ = v27
        }
    case "\n\t":
        if (true && (65535 >= 0)) {
            v28 := []int{306}
            for fi29 := 0; fi29 < 16; fi29++ {
                switch ((v2 + v2) + (v2 + "hello")) {
                case "\n\t":
                    v30 := GenS3{F0: 617}
                    _ = v30
                case "":
                    _ = len(v28)
                case "a":
                    switch -460 {
                    case 9223372036854775807:
                        if (!false) {
                            var v31 uint32 = v3
                            var v32 float64 = 1e10
                            _ = v31
                            _ = v32
                        }
                    default:
                        var v33 uint64 = 1
                        _ = v33
                    }
                default:
                    var v34 uint16 = 255
                    _ = v34
                }
                _ = fi29
            }
            _ = v28
        } else {
            fmt.Println(v1)
            for fi35 := 0; fi35 < 1; fi35++ {
                if true {
                    v36 := GenS1{F0: 32767, F1: 42}
                    v37 := GenS2{F0: 255}
                    _ = v36
                    _ = v37
                }
                _ = fi35
            }
        }
    }
    var v38 uint64 = uint64(v3)
    v39 := GenS2{F0: 94}
    var v40 string = (v2 + "foo bar")
    switch "foo bar" {
    case "hello":
        var v41 int32 = -366
        _ = v41
    case "a":
        switch 127 {
        case 2147483647:
            if (genFn0(v38, 32768) <= 56) {
                var v42 int16 = (-(v1))
                switch (("" + "\n\t") + ("" + v2)) {
                case "foo bar":
                    v43 := GenS1{F0: v1, F1: (^(v3))}
                    _ = v43
                }
                var v44 uint32 = (v3 / ((v3) | 1))
                _ = v42
                _ = v44
            }
        case 9223372036854775807:
            v45 := GenS4{F0: (v38 / ((v38) | 1)), F1: "\n\t"}
            _ = v45
        case 127:
            var v46 uint64 = v38
            _ = v46
        }
    }
    switch 961 {
    case 637:
        switch 9223372036854775807 {
        case 42:
            fmt.Println(v40, v40)
        case -2147483648:
            var v47 uint8 = uint8(v38)
            _ = v47
        default:
            var v48 int8 = 99
            _ = v48
        }
    }
    var v49 uint8 = 7
    var v50 string = "hello"
    if true {
        var v51 int8 = 0
        for fi52 := 0; fi52 < 19; fi52++ {
            v1 = v1
            if (820 > 0) {
                v53 := GenS1{F0: -1, F1: v3}
                _ = v53
            } else {
                fmt.Println(v2)
            }
            _ = fi52
        }
        if true {
            for fi54 := 0; fi54 < 6; fi54++ {
                var v55 uint = 255
                _ = fi54
                _ = v55
            }
        }
        _ = v51
    }
    if false {
        if true {
            v56 := []int{-470, 2147483647, -9223372036854775808, 21}
            _ = v56
        }
        for fi57 := 0; fi57 < 3; fi57++ {
            var v58 int8 = -1
            v59 := GenS4{F0: (v38 | (v38 << ((735) & 63))), F1: ("foo bar" + v50)}
            switch -2147483648 {
            case 32767:
                fmt.Println(v3, v1)
            case -9223372036854775808:
                if (v58 != -1) {
                    var v60 int64 = 486
                    fmt.Println(v1, v1, v1)
                    _ = v60
                }
            default:
                var v61 uint8 = uint8(v38)
                _ = v61
            }
            _ = fi57
            _ = v58
            _ = v59
        }
        var v62 int16 = (v1 % ((genFn1(65535, (v50 + (v40 + (v50 + "foo bar"))), v38)) | 1))
        _ = v62
    } else {
        switch 9223372036854775807 {
        case 42:
            switch 4294967295 {
            case 4294967295:
                var v63 int = 32767
                _ = v63
            case 0:
                for fi64 := 0; fi64 < 8; fi64++ {
                    var v65 bool = false
                    fmt.Println(fi64, v2, v40)
                    _ = fi64
                    _ = v65
                }
            default:
                switch 9223372036854775807 {
                case 9223372036854775807:
                    fmt.Println(v1, v1, v50)
                case 0:
                    fmt.Println(v3, v2)
                case 42:
                    var v66 uint16 = 32767
                    _ = v66
                default:
                    var v67 uint = 283
                    _ = v67
                }
            }
        }
        switch 42 {
        case -481:
            var v68 uint8 = (v49 + v49)
            _ = v68
        default:
            if true {
                var v69 bool = (!(42 < int64(v3)))
                v70 := GenS5{F0: true}
                _ = v69
                _ = v70
            }
        }
        v38 -= (v38 >> ((943) & 63))
    }
    var v71 float64 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v40:", v40)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v40
    _ = v49
    _ = v50
    _ = v71
}
