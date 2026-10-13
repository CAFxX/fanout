package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 int16; F2 uint16 }
type GenS3 struct { F0 int8; F1 uint64; F2 int8 }
type GenS4 struct { F0 int; F1 string; F2 int }
type GenS5 struct { F0 int8; F1 int; F2 bool }
type GenS6 struct { F0 string; F1 int16; F2 uint }

func main() {
    var v1 int = -235
    v2 := make(map[uint32]bool)
    var v3 int16 = 32767
    v4 := GenS1{F0: -42}
    v1 = v1
    if true {
        var v5 int16 = 42
        v6 := GenS1{F0: -1}
        for fi7 := 0; fi7 < 13; fi7++ {
            var v8 int = 32768
            _ = fi7
            _ = v8
        }
        _ = v5
        _ = v6
    }
    var v9 uint8 = 128
    if (!(((("foo bar" + ("hello" + "a")) <= "foo bar") || true) || (100 >= -1))) {
        fmt.Println(v1, v9, v3)
        switch ("foo bar" + "hello") {
        case "":
            v3 = (v3 / ((v3) | 1))
        case "hello":
            v2[173] = true
        default:
            var v10 uint16 = 1
            _ = v10
        }
        for fi11 := 0; fi11 < 10; fi11++ {
            var v12 uint8 = v9
            _ = fi11
            _ = v12
        }
    }
    var v13 string = "a"
    var v14 uint64 = 42
    v13 = ((("foo bar" + v13) + "a") + "a")
    if true {
        var v15 int8 = 100
        _ = v15
    }
    switch v13 {
    case "hello":
        switch v1 {
        case 1:
            for fi16 := 0; fi16 < 19; fi16++ {
                switch (v13 + (v13 + v13)) {
                case "foo bar":
                    if (v3 != v3) {
                        v17 := GenS1{F0: 42}
                        switch (v13 + "a") {
                        case "\n\t":
                            if true {
                                for fi18 := 0; fi18 < 13; fi18++ {
                                    v19 := GenS1{F0: 0}
                                    _ = fi18
                                    _ = v19
                                }
                                if false {
                                    if true {
                                        v2[4294967295] = false
                                        var v20 uint = 4294967295
                                        _ = v20
                                    } else {
                                        if (((v3 | 1000) >= 42) || (2 >= 0)) {
                                            v9 += (v9 << ((4294967295) & 63))
                                            for fi21 := 0; fi21 < 14; fi21++ {
                                                v22 := GenS2{F0: 42, F1: (v3 >> ((1) & 63)), F2: 0}
                                                _ = fi21
                                                _ = v22
                                            }
                                            if false {
                                                v23 := GenS1{F0: 9223372036854775807}
                                                v24 := GenS1{F0: 42}
                                                _ = v23
                                                _ = v24
                                            } else {
                                                for fi25 := 0; fi25 < 4; fi25++ {
                                                    v13 = "hello"
                                                    _ = fi25
                                                }
                                            }
                                        }
                                        switch -9223372036854775808 {
                                        case 32767:
                                            for fi26 := 0; fi26 < 9; fi26++ {
                                                v13 = "foo bar"
                                                _ = fi26
                                            }
                                        case 9223372036854775807:
                                            for fi27 := 0; fi27 < 6; fi27++ {
                                                v28 := []int{(v1 + fi16), 2147483647}
                                                _ = fi27
                                                _ = v28
                                            }
                                        case 127:
                                            v2[uint32(v14)] = (!(-102 >= 42))
                                        }
                                    }
                                } else {
                                    if (0 == -3) {
                                        var v29 uint64 = (v14 / (((v14 + (v14 + v14))) | 1))
                                        var v30 int32 = -2147483648
                                        var v31 uint8 = v9
                                        _ = v29
                                        _ = v30
                                        _ = v31
                                    }
                                }
                            }
                        case "a":
                            switch v1 {
                            case 988:
                                v13 = (v13 + v13)
                            case 32768:
                                switch (v13 + "hello") {
                                case "foo bar":
                                    v9 -= 2
                                case "\n\t":
                                    var v32 int16 = (v3 ^ v3)
                                    _ = v32
                                case "":
                                    var v33 int32 = 42
                                    _ = v33
                                default:
                                    v1 = (v1 >> ((500) & 63))
                                }
                            case -9223372036854775808:
                                if true {
                                    v3 += 0
                                    v9 = (v9 ^ v9)
                                }
                            }
                        case "":
                            for fi34 := 0; fi34 < 1; fi34++ {
                                v9 = v9
                                _ = fi34
                            }
                        }
                        _ = v17
                    }
                default:
                    switch v13 {
                    case "a":
                        if ((v9 < 221) && (!(-1 != 100))) {
                            fmt.Println(v9, fi16, v9)
                            var v35 float64 = 0.1
                            var v36 uint32 = 42
                            _ = v35
                            _ = v36
                        } else {
                            v9 += (v9 | (v9 ^ (v9 + (v9 ^ v9))))
                            v2[0] = true
                        }
                    case "\n\t":
                        v37 := GenS1{F0: 307}
                        _ = v37
                    case "hello":
                        var v38 uint64 = 247
                        _ = v38
                    default:
                        if true {
                            v13 = ""
                            v13 = (("\n\t" + ((v13 + "a") + "")) + "")
                            var v39 int = 9223372036854775807
                            _ = v39
                        } else {
                            v40 := []int{fi16, v1, (v1 >> ((65535) & 63)), (v1 << ((52) & 63))}
                            _ = v40
                        }
                    }
                }
                v14 = v14
                fmt.Println(v9, fi16)
                _ = fi16
            }
        case 2147483647:
            switch (v1 ^ 32768) {
            case -2147483648:
                if ((v9 ^ v9) != v9) {
                    v41 := GenS2{F0: 1, F1: (v3 * v3), F2: 32767}
                    v42 := GenS3{F0: 3, F1: (v14 + v14), F2: 0}
                    var v43 float64 = 1e10
                    _ = v41
                    _ = v42
                    _ = v43
                } else {
                    if (32767 != (-(v1))) {
                        for fi44 := 0; fi44 < 20; fi44++ {
                            v45 := GenS2{F0: -1, F1: v3, F2: 3}
                            fmt.Println(v13)
                            var v46 int8 = int8(fi44)
                            _ = fi44
                            _ = v45
                            _ = v46
                        }
                        var v47 int8 = -17
                        v3 -= (v3 % (((^(v3))) | 1))
                        _ = v47
                    } else {
                        if true {
                            v48 := GenS4{F0: (v1 - 1), F1: v13, F2: -2147483648}
                            _ = v48
                        } else {
                            for fi49 := 0; fi49 < 4; fi49++ {
                                var v50 int64 = 0
                                fmt.Println(fi49)
                                _ = fi49
                                _ = v50
                            }
                            if false {
                                v51 := GenS5{F0: 42, F1: v1, F2: false}
                                v1 = -1
                                _ = v51
                            }
                            for fi52 := 0; fi52 < 19; fi52++ {
                                v53 := GenS4{F0: (fi52 + (v1 | (fi52 ^ fi52))), F1: (v13 + v13), F2: 4294967295}
                                v54 := []int{}
                                _ = fi52
                                _ = v53
                                _ = v54
                            }
                        }
                        fmt.Println(v13, v14)
                        v9 = (^(v9))
                    }
                    switch (v13 + v13) {
                    case "foo bar":
                        var v55 uint = 716
                        _ = v55
                    case "":
                        var v56 int16 = 32767
                        _ = v56
                    default:
                        for fi57 := 0; fi57 < 10; fi57++ {
                            var v58 int64 = -9223372036854775808
                            _ = fi57
                            _ = v58
                        }
                    }
                }
            case -23:
                v59 := GenS2{F0: 249, F1: v3, F2: 977}
                _ = v59
            default:
                for fi60 := 0; fi60 < 15; fi60++ {
                    fmt.Println(fi60, v9, fi60)
                    _ = fi60
                }
            }
        case 297:
            v61 := v2[2147483647]
            _ = v61
        }
    case "a":
        switch v13 {
        case "":
            delete(v2, 638)
        case "a":
            v62 := []int{}
            _ = v62
        }
    case "foo bar":
        v63 := []int{v1, (-(v1)), (v1 * v1)}
        _ = v63
    default:
        var v64 uint64 = v14
        _ = v64
    }
    if (true && (!(127 <= 0))) {
        v65 := []int{}
        switch "a" {
        case "a":
            v13 = v13
        }
        var v66 int8 = 100
        _ = v65
        _ = v66
    }
    for fi67 := 0; fi67 < 1; fi67++ {
        if (!(723 != 42)) {
            _ = len(v2)
            var v68 int64 = -872
            _ = v68
        }
        for fi69 := 0; fi69 < 5; fi69++ {
            var v70 float32 = 1e10
            _ = fi69
            _ = v70
        }
        _ = fi67
    }
    if true {
        v71 := GenS6{F0: v13, F1: (v3 - v3), F2: 65535}
        fmt.Println(v9)
        v72 := GenS5{F0: 127, F1: (v1 >> ((927) & 63)), F2: true}
        _ = v71
        _ = v72
    }
    switch v13 {
    case "\n\t":
        var v73 int32 = -1
        _ = v73
    }
    for fi74 := 0; fi74 < 17; fi74++ {
        if false {
            for fi75 := 0; fi75 < 12; fi75++ {
                v76 := GenS6{F0: v13, F1: (v3 * 1), F2: 42}
                _ = fi75
                _ = v76
            }
        }
        fmt.Println(v14)
        v14 = v14
        _ = fi74
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
