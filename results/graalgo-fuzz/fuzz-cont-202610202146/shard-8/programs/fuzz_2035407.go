package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int; F1 int8; F2 uint16 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint32; F1 uint }

func main() {
    var v1 string = "\n\t"
    var v2 uint16 = 42
    v3 := []string{(v1 + (("a" + "a") + ("\n\t" + v1))), v1, "\n\t"}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v2)
    switch (((v1 + v1) + "hello") + v1) {
    case "hello":
        v4 := GenS1{F0: 42, F1: -2147483648}
        _ = v4
    default:
        var v5 int64 = -641
        _ = v5
    }
    v6 := GenS2{F0: -32768, F1: 819}
    if true {
        fmt.Println(v2, v1)
        for fi7 := 0; fi7 < 8; fi7++ {
            fmt.Println(fi7, v2)
            _ = fi7
        }
        _ = len(v3)
    } else {
        switch 0 {
        case -656:
            if false {
                var v8 int8 = 39
                _ = v8
            } else {
                fmt.Println(v2)
                switch v1 {
                case "hello":
                    for fi9 := 0; fi9 < 2; fi9++ {
                        v3 = append(v3, (((v1 + "a") + ("" + ("hello" + v1))) + v1))
                        for fi10 := 0; fi10 < 16; fi10++ {
                            if (255 < 255) {
                                v11 := GenS2{F0: 42, F1: 18446744073709551615}
                                v2 += (v2 ^ 2)
                                _ = v11
                            } else {
                                var v12 string = v1
                                var v13 int64 = -761
                                if (false && (0.0 > (0.1 / 3.14159))) {
                                    if true {
                                        var v14 string = "hello"
                                        var v15 int8 = int8(fi9)
                                        _ = v14
                                        _ = v15
                                    } else {
                                        var v16 int32 = 875
                                        var v17 int32 = 2147483647
                                        _ = v16
                                        _ = v17
                                    }
                                }
                                _ = v12
                                _ = v13
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            var v18 int32 = 0
                            _ = fi10
                            _ = v18
                        }
                        for fi19 := 0; fi19 < 18; fi19++ {
                            switch ("\n\t" + v1) {
                            case "a":
                                switch v1 {
                                case "a":
                                    if false {
                                        v20 := GenS3{F0: 859}
                                        var v21 int32 = 42
                                        _ = v20
                                        _ = v21
                                    } else {
                                        if (false && false) {
                                            v1 = (v1 + (v1 + v1))
                                            v22 := GenS4{F0: -816}
                                            _ = v22
                                        } else {
                                            v23 := GenS3{F0: 65535}
                                            if ((127 == 39) && false) {
                                                v2 += v2
                                                var v24 int64 = 9223372036854775807
                                                _ = v24
                                            }
                                            if len(v3) > 0 {
                                                v3[0] = ("foo bar" + (("foo bar" + v1) + ""))
                                            }
                                            _ = v23
                                        }
                                        v25 := GenS2{F0: -1, F1: 1}
                                        _ = v25
                                    }
                                case "":
                                    var v26 int8 = 42
                                    var v27 bool = true
                                    v28 := GenS2{F0: 1, F1: 255}
                                    if false {
                                        var v29 string = ("foo bar" + (((v1 + v1) + "hello") + "hello"))
                                        _ = v29
                                    }
                                    switch (fi9 | 717) {
                                    case 9223372036854775807:
                                        var v30 uint32 = 2
                                        _ = v30
                                    case 42:
                                        var v31 int64 = -9223372036854775808
                                        _ = v31
                                    default:
                                        v3 = append(v3, ("" + v1))
                                    }
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                case "\n\t":
                                    switch v1 {
                                    case "":
                                        if true {
                                            var v32 uint16 = 42
                                            if len(v3) > 0 {
                                                v3[0] = v1
                                            }
                                            _ = v32
                                        } else {
                                            fmt.Println(fi19)
                                            switch (v1 + "foo bar") {
                                            case "hello":
                                                v2 = 42
                                            case "a":
                                                var v33 uint16 = (v2 & v2)
                                                _ = v33
                                            case "":
                                                if true {
                                                    switch ("" + "foo bar") {
                                                    case "hello":
                                                        var v34 uint8 = 255
                                                        _ = v34
                                                    default:
                                                        switch v1 {
                                                        case "":
                                                            v2 -= v2
                                                        default:
                                                            if true {
                                                                var v35 float64 = float64(1e10)
                                                                if false {
                                                                    var v36 int32 = 2147483647
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                    _ = v36
                                                                } else {
                                                                    v37 := GenS5{F0: 42}
                                                                    _ = v37
                                                                }
                                                                var v38 float32 = 3.14159
                                                                _ = v35
                                                                _ = v38
                                                            }
                                                        }
                                                    }
                                                    fmt.Println(v1, fi9, v1)
                                                }
                                            }
                                            switch v1 {
                                            case "a":
                                                var v39 float64 = 3.14159
                                                _ = v39
                                            }
                                        }
                                    case "a":
                                        if (true || false) {
                                            v2 = v2
                                            _ = len(v3)
                                            var v40 int64 = -848
                                            _ = v40
                                        }
                                    default:
                                        switch fi9 {
                                        case -9223372036854775808:
                                            fmt.Println(fi19, v2, fi9)
                                        default:
                                            if false {
                                                v1 = v1
                                                if (2 == 474) {
                                                    var v41 uint = 1
                                                    _ = v41
                                                }
                                                var v42 float64 = 0.1
                                                _ = v42
                                            }
                                        }
                                    }
                                default:
                                    if (v1 < "hello") {
                                        v43 := GenS2{F0: 32767, F1: 773}
                                        _ = v43
                                    } else {
                                        v1 = v1
                                        if true {
                                            var v44 uint16 = 2
                                            fmt.Println(v44, fi19)
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            _ = v44
                                        } else {
                                            v1 = ("foo bar" + (v1 + ("hello" + v1)))
                                            v45 := GenS6{F0: -1, F1: 42, F2: v2}
                                            fmt.Println(fi19)
                                            _ = v45
                                        }
                                        v3 = append(v3, "a")
                                    }
                                }
                            case "\n\t":
                                v1 = v1
                            case "foo bar":
                                switch -889 {
                                case -1:
                                    fmt.Println(fi9)
                                    v46 := GenS5{F0: 127}
                                    fmt.Println(fi9, v1, v1)
                                    _ = v46
                                case 4294967295:
                                    if (317 == 102) {
                                        switch 1 {
                                        case 127:
                                            if false {
                                                if (2 > 42) {
                                                    fmt.Println(v1)
                                                }
                                                var v47 int32 = 42
                                                switch fi9 {
                                                case -817:
                                                    var v48 uint8 = 76
                                                    _ = v48
                                                case 4294967295:
                                                    _ = len(v3)
                                                    var v49 uint = 4294967295
                                                    v1 = ("a" + v1)
                                                    _ = v49
                                                }
                                                _ = v47
                                            } else {
                                                var v50 uint64 = 392
                                                _ = v50
                                            }
                                        case 925:
                                            v2 = v2
                                        case 675:
                                            var v51 uint16 = 255
                                            _ = v51
                                        }
                                    } else {
                                        fmt.Println(v2, v1, fi19)
                                    }
                                case 368:
                                    v52 := GenS2{F0: 42, F1: 543}
                                    _ = v52
                                default:
                                    v53 := GenS2{F0: 30, F1: 255}
                                    _ = v53
                                }
                            }
                            fmt.Println(fi19)
                            var v54 float64 = 1.5
                            _ = fi19
                            _ = v54
                        }
                        _ = fi9
                    }
                case "\n\t":
                    v1 = "foo bar"
                case "foo bar":
                    switch ((v1 + v1) + v1) {
                    case "":
                        if (128 != 127) {
                            var v55 float32 = 1e10
                            fmt.Println(v1, v1, v1)
                            _ = v55
                        } else {
                            if (4294967295 != 9223372036854775807) {
                                v1 = v1
                            } else {
                                if (true || (!true)) {
                                    var v56 int8 = 1
                                    _ = v56
                                }
                                if len(v3) > 0 {
                                    v3[0] = "\n\t"
                                }
                            }
                            if false {
                                _ = len(v3)
                            }
                            var v57 uint64 = 206
                            _ = v57
                        }
                    }
                default:
                    v3 = append(v3, ("a" + v1))
                }
            }
        case 530:
            switch -2147483648 {
            case 774:
                switch (v1 + "hello") {
                case "":
                    for fi58 := 0; fi58 < 4; fi58++ {
                        var v59 int64 = 42
                        _ = fi58
                        _ = v59
                    }
                case "hello":
                    fmt.Println(v2)
                case "foo bar":
                    var v60 string = ("hello" + ("hello" + (((((v1 + ("hello" + ("\n\t" + "a"))) + ("" + ("a" + "hello"))) + ("hello" + ("hello" + (v1 + ("a" + ""))))) + "\n\t") + (v1 + ("a" + (v1 + "\n\t"))))))
                    _ = v60
                default:
                    var v61 uint32 = 4294967295
                    _ = v61
                }
            default:
                if true {
                    for fi62 := 0; fi62 < 6; fi62++ {
                        var v63 int16 = 0
                        for fi64 := 0; fi64 < 16; fi64++ {
                            v65 := GenS7{F0: (v1 + (v1 + "a"))}
                            _ = fi64
                            _ = v65
                        }
                        for fi66 := 0; fi66 < 1; fi66++ {
                            v63 += v63
                            switch "\n\t" {
                            case "":
                                var v67 uint32 = 0
                                _ = v67
                            default:
                                v2 += v2
                            }
                            switch ("" + (v1 + (("hello" + v1) + "a"))) {
                            case "hello":
                                var v68 uint32 = 2
                                var v69 int = (fi62 ^ (-(fi66)))
                                _ = v68
                                _ = v69
                            case "a":
                                v63 += 935
                                v63 = (-(v63))
                            case "foo bar":
                                var v70 int64 = -1
                                _ = v70
                            default:
                                v3 = append(v3, "hello")
                            }
                            var v71 uint32 = 0
                            var v72 float64 = 3.14159
                            var v73 string = v1
                            v2 -= (v2 >> ((18446744073709551615) & 63))
                            v73 = v1
                            _ = fi66
                            _ = v71
                            _ = v72
                            _ = v73
                        }
                        _ = fi62
                        _ = v63
                    }
                } else {
                    fmt.Println(v2, v2)
                }
            }
        case 9223372036854775807:
            if false {
                v74 := GenS3{F0: 17}
                var v75 uint64 = 42
                switch v1 {
                case "\n\t":
                    fmt.Println(v1, v1, v75)
                case "foo bar":
                    fmt.Println(v75, v75, v75)
                case "a":
                    var v76 int = 9223372036854775807
                    _ = v76
                default:
                    for fi77 := 0; fi77 < 3; fi77++ {
                        switch (^(fi77)) {
                        case 385:
                            var v78 string = (v1 + ("" + "a"))
                            _ = v78
                        case 2147483647:
                            switch (fi77 % (((fi77 >> ((1) & 63))) | 1)) {
                            case 32767:
                                v75 -= v75
                            case 1:
                                if true {
                                    var v79 float32 = 3.14159
                                    _ = v79
                                }
                            default:
                                v3 = append(v3, v1)
                            }
                        default:
                            if (!((!(!true)) || false)) {
                                v80 := GenS8{F0: 499, F1: 65535}
                                v1 = "a"
                                _ = v80
                            } else {
                                var v81 int16 = 0
                                v2 += (^(v2))
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = v81
                            }
                        }
                        for fi82 := 0; fi82 < 2; fi82++ {
                            v83 := GenS6{F0: fi77, F1: 42, F2: (^(v2))}
                            var v84 string = v1
                            _ = fi82
                            _ = v83
                            _ = v84
                        }
                        _ = fi77
                    }
                }
                _ = v74
                _ = v75
            } else {
                v2 -= v2
            }
        }
    }
    var v85 int8 = -27
    if len(v3) > 0 {
        v3[0] = v1
    }
    if len(v3) > 0 {
        v3[0] = ((v1 + v1) + v1)
    }
    v2 = v2
    fmt.Println(v2, v2, v85)
    for fi86 := 0; fi86 < 8; fi86++ {
        var v87 uint = 65535
        _ = fi86
        _ = v87
    }
    var v88 bool = true
    v88 = (true && (v85 > (v85 % ((v85) | 1))))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v85:", v85)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v85
    _ = v88
}
