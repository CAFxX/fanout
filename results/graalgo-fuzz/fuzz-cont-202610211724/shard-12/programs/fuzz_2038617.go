package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8; F1 int; F2 int }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int8; F1 int8 }
func genFn0(p0 bool) uint16 {
    for fi7 := 0; fi7 < 12; fi7++ {
        v8 := GenS1{F0: 2, F1: 1, F2: 128}
        v9 := []int{fi7, fi7}
        var v10 string = ("hello" + "hello")
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    v11 := GenS1{F0: 2, F1: 65535, F2: 42}
    switch -623 {
    case 128:
        v12 := []int{}
        _ = v12
    case -352:
        var v13 bool = p0
        _ = v13
    case 0:
        v14 := GenS1{F0: 2, F1: 29, F2: 50}
        _ = v14
    default:
        p0 = (p0 || (4294967295 != 1))
    }
    fmt.Println(p0)
    v15 := []int{32768}
    _ = p0
    _ = v11
    _ = v15
    return 2
}
func genFn1(p0 int64, p1 uint, p2 string) bool {
    var v16 uint8 = 206
    var v17 bool = false
    for fi18 := 0; fi18 < 9; fi18++ {
        v16 = (v16 - (v16 ^ (v16 / (((v16 << ((4294967295) & 63))) | 1))))
        v19 := []int{(fi18 << ((p1) & 63)), 4294967295}
        switch p2 {
        case "a":
            fmt.Println(p1, p1)
        case "foo bar":
            var v20 uint64 = 875
            _ = v20
        }
        _ = fi18
        _ = v19
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    return v17
}

func main() {
    var v1 uint32 = 4294967295
    var v2 int64 = 42
    v3 := []uint32{(v1 | v1), 841, (v1 % ((524) | 1)), v1}
    v4 := []string{(("\n\t" + ("foo bar" + ("foo bar" + "foo bar"))) + "hello"), ("foo bar" + ("foo bar" + "\n\t")), "\n\t", ""}
    var v5 uint = 443
    v6 := make(map[int]bool)
    v5 -= (v5 | v5)
    var v21 uint64 = 4294967295
    if true {
        v22 := GenS1{F0: (v1 >> (((^(v5))) & 63)), F1: 65535, F2: 0}
        var v23 uint16 = 330
        for fi24 := 0; fi24 < 8; fi24++ {
            v6[(fi24 % (((fi24 / ((32768) | 1))) | 1))] = true
            _ = fi24
        }
        _ = v22
        _ = v23
    }
    v25 := GenS2{F0: 116}
    fmt.Println(v21)
    if false {
        var v26 uint16 = 65535
        switch ("hello" + (("hello" + "") + ("a" + ((("\n\t" + "foo bar") + "foo bar") + ("a" + ("" + "hello")))))) {
        case "\n\t":
            v6[-49] = ((904 <= 128) || ((^(v1)) < (v1 * (^(v1)))))
        case "foo bar":
            var v27 uint64 = v21
            _ = v27
        case "":
            if false {
                var v28 string = "hello"
                v2 = 9223372036854775807
                for fi29 := 0; fi29 < 18; fi29++ {
                    if false {
                        var v30 int64 = 9223372036854775807
                        _ = v30
                    } else {
                        v31 := GenS2{F0: uint8(v5)}
                        var v32 int = (fi29 >> ((42) & 63))
                        v26 -= v26
                        _ = v31
                        _ = v32
                    }
                    var v33 uint8 = 156
                    _ = fi29
                    _ = v33
                }
                _ = v28
            }
        default:
            for fi34 := 0; fi34 < 5; fi34++ {
                v5 += 255
                for fi35 := 0; fi35 < 2; fi35++ {
                    v36 := v6[42]
                    _ = fi35
                    _ = v36
                }
                v21 = v21
                _ = fi34
            }
        }
        _ = v26
    } else {
        if true {
            var v37 uint8 = 136
            switch 980 {
            case -988:
                var v38 uint64 = (v21 | v21)
                _ = v38
            }
            v39 := v6[32768]
            _ = v37
            _ = v39
        } else {
            v40 := GenS3{F0: 32767}
            _ = v40
        }
        for fi41 := 0; fi41 < 16; fi41++ {
            v42 := GenS1{F0: v1, F1: v5, F2: 0}
            _ = fi41
            _ = v42
        }
    }
    if true {
        switch 1 {
        case -2147483648:
            var v43 int = 9223372036854775807
            _ = v43
        }
        v44 := GenS4{F0: 255, F1: 0, F2: 9223372036854775807}
        _ = v44
    } else {
        v2 = 540
        v5 += 1
    }
    switch ("hello" + "foo bar") {
    case "a":
        v5 = (v5 & (v5 * (^(v5))))
    case "\n\t":
        var v45 int64 = (v2 | (v2 << (((v5 ^ 18446744073709551615)) & 63)))
        _ = v45
    case "":
        v46 := GenS5{F0: v1}
        _ = v46
    default:
        if true {
            var v47 uint8 = 127
            v5 -= 4294967295
            v4 = append(v4, "\n\t")
            _ = v47
        } else {
            v1 = uint32(v5)
        }
    }
    switch "a" {
    case "foo bar":
        if ((180 <= 42) && (!(v21 == (v21 << (((v5 >> (((v5 / (((v5 / ((v5) | 1))) | 1))) & 63))) & 63))))) {
            v48 := GenS3{F0: uint16(v5)}
            _ = v48
        }
    case "a":
        for fi49 := 0; fi49 < 3; fi49++ {
            if ("" == ("a" + "a")) {
                v21 -= v21
                switch (("foo bar" + "a") + "\n\t") {
                case "":
                    delete(v6, (fi49 * (fi49 << ((1) & 63))))
                case "hello":
                    if (0 <= (v1 - v1)) {
                        v5 += uint(v21)
                        if false {
                            v1 -= 2147483648
                        } else {
                            v21 -= v21
                            switch "foo bar" {
                            case "":
                                v2 -= v2
                            case "\n\t":
                                if false {
                                    v6[fi49] = false
                                    for fi50 := 0; fi50 < 16; fi50++ {
                                        if (float32((1.5 / 0.1)) == 1e10) {
                                            var v51 int32 = -958
                                            delete(v6, fi50)
                                            if true {
                                                v5 += 586
                                            } else {
                                                switch -551 {
                                                case -2147483648:
                                                    v52 := GenS6{F0: (^(v1))}
                                                    _ = v52
                                                case 9223372036854775807:
                                                    if true {
                                                        v53 := GenS5{F0: (^(v1))}
                                                        _ = len(v6)
                                                        var v54 uint32 = v1
                                                        _ = v53
                                                        _ = v54
                                                    } else {
                                                        delete(v6, fi50)
                                                        v1 += 42
                                                        v51 = -1
                                                    }
                                                case 0:
                                                    if (-181 != (fi49 * -1)) {
                                                        v51 -= int32(v2)
                                                    }
                                                }
                                            }
                                            _ = v51
                                        }
                                        var v55 int16 = 31
                                        fmt.Println(v1, v1, v21)
                                        _ = fi50
                                        _ = v55
                                    }
                                    if (42 < uint16(v5)) {
                                        var v56 uint32 = 4294967295
                                        switch (fi49 >> ((1) & 63)) {
                                        case 9223372036854775807:
                                            for fi57 := 0; fi57 < 13; fi57++ {
                                                if false {
                                                    switch (("a" + ((("" + ("foo bar" + ("" + "a"))) + "\n\t") + "")) + "a") {
                                                    case "hello":
                                                        v4 = append(v4, ("foo bar" + "foo bar"))
                                                    }
                                                    v5 -= 4294967295
                                                    v56 -= v56
                                                    v6[(fi57 << ((255) & 63))] = true
                                                    v5 -= (v5 - (^(v5)))
                                                    v5 += (v5 ^ (v5 * v5))
                                                }
                                                v1 = (v56 * 775)
                                                _ = fi57
                                            }
                                        case 187:
                                            switch ("foo bar" + "hello") {
                                            case "foo bar":
                                                v56 -= 2147483648
                                            case "":
                                                var v58 float64 = 1.5
                                                _ = v58
                                            default:
                                                var v59 uint = (v5 >> (((v5 / ((v5) | 1))) & 63))
                                                _ = v59
                                            }
                                        case -2147483648:
                                            var v60 float64 = (1e10 - -2.25)
                                            _ = v60
                                        }
                                        _ = v56
                                    } else {
                                        v2 = 9223372036854775807
                                    }
                                } else {
                                    fmt.Println(fi49, v1, v21)
                                }
                            }
                            v1 += v1
                        }
                    }
                default:
                    v21 = (^(v21))
                }
                v61 := GenS1{F0: (v1 ^ (v1 >> ((42) & 63))), F1: 18446744073709551615, F2: 0}
                _ = v61
            } else {
                var v62 int32 = 2147483647
                _ = v62
            }
            if ("\n\t" >= "\n\t") {
                v63 := GenS4{F0: 3, F1: (fi49 + fi49), F2: (-(fi49))}
                v2 += v2
                fmt.Println(v2, fi49)
                _ = v63
            } else {
                for fi64 := 0; fi64 < 19; fi64++ {
                    var v65 float32 = 3.14159
                    fmt.Println(v21, v21)
                    _ = fi64
                    _ = v65
                }
                v1 += 142
                for fi66 := 0; fi66 < 7; fi66++ {
                    var v67 string = ""
                    var v68 int64 = v2
                    _ = fi66
                    _ = v67
                    _ = v68
                }
            }
            _ = fi49
        }
    case "\n\t":
        var v69 int = -1
        _ = v69
    }
    switch ("\n\t" + "\n\t") {
    case "foo bar":
        var v70 int16 = -1
        _ = v70
    case "\n\t":
        v21 = 18446744073709551615
    case "":
        fmt.Println(v2, v1, v2)
    default:
        switch -694 {
        case 9223372036854775807:
            v5 += v5
        case 322:
            var v71 uint32 = (^(v1))
            _ = v71
        default:
            v72 := GenS6{F0: v1}
            _ = v72
        }
    }
    for fi73 := 0; fi73 < 20; fi73++ {
        delete(v6, (fi73 - (^(fi73))))
        v5 = (v5 ^ 274)
        var v74 float64 = 0.0
        _ = fi73
        _ = v74
    }
    if (false && true) {
        var v75 float32 = 1e10
        v76 := v6[913]
        _ = v75
        _ = v76
    }
    var v77 uint8 = 87
    for fi78 := 0; fi78 < 4; fi78++ {
        var v79 uint8 = (v77 * 128)
        _ = fi78
        _ = v79
    }
    switch 859 {
    case 722:
        if (v2 >= (v2 * v2)) {
            if len(v3) > 0 {
                v3[0] = v1
            }
            v77 += (v77 + (v77 % (((^(v77))) | 1)))
        }
    }
    var v80 uint32 = v1
    v81 := GenS7{F0: 1, F1: 124}
    if len(v4) > 0 {
        v4[0] = "a"
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v77:", v77)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v21
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v77
    _ = v80
    _ = v81
}
