package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8; F1 bool }
type GenS5 struct { F0 int8; F1 int64; F2 uint16 }
type GenS6 struct { F0 int32 }
func genFn0(p0 int, p1 int8, p2 int) bool {
    if false {
        switch p0 {
        case 128:
            var v7 int64 = -651
            _ = v7
        default:
            if true {
                for fi8 := 0; fi8 < 10; fi8++ {
                    if false {
                        for fi9 := 0; fi9 < 17; fi9++ {
                            p0 += (fi8 >> ((255) & 63))
                            fmt.Println(fi8, p0, fi9)
                            v10 := GenS1{F0: 517, F1: (p0 / ((p0) | 1))}
                            if false {
                                p0 += (fi8 ^ 32767)
                                var v11 uint = 4294967295
                                _ = v11
                            }
                            var v12 int16 = -1
                            _ = fi9
                            _ = v10
                            _ = v12
                        }
                        var v13 int16 = 32767
                        _ = v13
                    }
                    _ = fi8
                }
                switch -146 {
                case -9223372036854775808:
                    switch (^(p0)) {
                    case -515:
                        var v14 uint16 = 65535
                        _ = v14
                    case 127:
                        if (1.5 <= -2.25) {
                            p1 -= 127
                            v15 := GenS1{F0: 0, F1: p0}
                            _ = v15
                        } else {
                            var v16 int = 42
                            _ = v16
                        }
                    default:
                        v17 := []int{504, p0}
                        _ = v17
                    }
                case 543:
                    fmt.Println(p2, p1, p0)
                default:
                    fmt.Println(p2, p2, p2)
                }
            }
        }
        switch (p0 + (p0 | p2)) {
        case 0:
            v18 := []int{}
            _ = v18
        case 9223372036854775807:
            var v19 uint8 = 253
            _ = v19
        case 42:
            p2 -= (p2 & (-(p0)))
        default:
            v20 := []int{(p0 - 9223372036854775807), (p2 - p0)}
            _ = v20
        }
        for fi21 := 0; fi21 < 12; fi21++ {
            fmt.Println(p0, fi21, p0)
            var v22 uint8 = 127
            v23 := []int{p2, -530, -1}
            _ = fi21
            _ = v22
            _ = v23
        }
    } else {
        switch "a" {
        case "\n\t":
            v24 := []int{p0, p2, 484, (p0 * (p0 + (p0 - 32768)))}
            _ = v24
        case "hello":
            v25 := []int{}
            _ = v25
        }
        for fi26 := 0; fi26 < 7; fi26++ {
            v27 := GenS2{F0: 2147483647}
            _ = fi26
            _ = v27
        }
        var v28 string = ("foo bar" + ("foo bar" + "hello"))
        _ = v28
    }
    var v29 float32 = -2.25
    fmt.Println(p0, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v29
    return (!(!true))
}
func genFn1(p0 int64, p1 string, p2 uint16) uint64 {
    switch "hello" {
    case "foo bar":
        if true {
            p1 = ((p1 + p1) + ((("hello" + ((p1 + ((p1 + (p1 + p1)) + ("foo bar" + ""))) + (p1 + ((((p1 + p1) + ("" + p1)) + (p1 + p1)) + "")))) + p1) + p1))
            if false {
                fmt.Println(p2, p0, p1)
            }
            p1 = p1
        }
    case "a":
        v30 := []int{-1}
        _ = v30
    case "hello":
        v31 := GenS1{F0: 42, F1: -9223372036854775808}
        _ = v31
    default:
        switch p1 {
        case "\n\t":
            var v32 uint64 = uint64(p2)
            _ = v32
        default:
            var v33 bool = true
            _ = v33
        }
    }
    switch 32768 {
    case 32767:
        var v34 uint16 = (p2 | (^(p2)))
        _ = v34
    case -268:
        for fi35 := 0; fi35 < 4; fi35++ {
            p2 += (p2 % ((p2) | 1))
            _ = fi35
        }
    }
    if false {
        if ((1 == 100) || false) {
            for fi36 := 0; fi36 < 6; fi36++ {
                v37 := GenS2{F0: fi36}
                _ = fi36
                _ = v37
            }
        } else {
            if (true && (true || ((p2 ^ (p2 / ((p2) | 1))) > p2))) {
                switch 4294967295 {
                case 42:
                    var v38 uint16 = (p2 << ((18446744073709551615) & 63))
                    _ = v38
                case 32767:
                    p2 = p2
                default:
                    p0 = (^(p0))
                }
            }
            if (2147483647 > int32(p0)) {
                p0 = (p0 / ((p0) | 1))
                var v39 uint = 42
                v39 -= (^(v39))
                _ = v39
            } else {
                fmt.Println(p1, p2)
                v40 := GenS2{F0: 42}
                var v41 int32 = 42
                _ = v40
                _ = v41
            }
            p0 -= (p0 - p0)
        }
        fmt.Println(p0, p0, p2)
    } else {
        fmt.Println(p1)
        p2 = (p2 % ((p2) | 1))
    }
    p1 = ""
    v42 := GenS3{F0: 41}
    _ = p0
    _ = p1
    _ = p2
    _ = v42
    return 4294967295
}

func main() {
    var v1 string = "foo bar"
    var v2 uint32 = 65535
    var v3 uint64 = 0
    var v4 int8 = 127
    v5 := []string{v1, (v1 + "hello"), v1}
    v6 := make(map[string]string)
    var v43 int = 9223372036854775807
    _ = len(v6)
    v4 -= v4
    for fi44 := 0; fi44 < 13; fi44++ {
        if true {
            fmt.Println(v43, fi44)
        }
        if (1e10 >= 0.0) {
            if (42 <= 0) {
                for fi45 := 0; fi45 < 8; fi45++ {
                    delete(v6, ("\n\t" + ""))
                    var v46 uint64 = (v3 % (((v3 ^ v3)) | 1))
                    v3 = v46
                    _ = fi45
                    _ = v46
                }
                var v47 uint64 = (^(v3))
                for fi48 := 0; fi48 < 8; fi48++ {
                    v49 := v6[v1]
                    if len(v5) > 0 {
                        v5[0] = v1
                    }
                    v6[(((v49 + (v1 + (("foo bar" + "") + v49))) + v49) + "a")] = v49
                    _ = fi48
                    _ = v49
                }
                _ = v47
            }
        }
        _ = len(v6)
        _ = fi44
    }
    if (true && (!false)) {
        if false {
            v2 = (^(v2))
            fmt.Println(v1, v2)
        } else {
            switch (v43 ^ 2147483647) {
            case 0:
                v50 := GenS4{F0: v4, F1: false}
                _ = v50
            default:
                switch v1 {
                case "\n\t":
                    var v51 uint32 = v2
                    _ = v51
                default:
                    var v52 uint = 42
                    _ = v52
                }
            }
            var v53 int64 = 1
            if (true || false) {
                for fi54 := 0; fi54 < 20; fi54++ {
                    var v55 uint = 4294967295
                    for fi56 := 0; fi56 < 3; fi56++ {
                        var v57 uint8 = 255
                        _ = fi56
                        _ = v57
                    }
                    _ = fi54
                    _ = v55
                }
                v53 = 1
            } else {
                v58 := GenS5{F0: (v4 & v4), F1: v53, F2: 42}
                if false {
                    fmt.Println(v3, v4, v43)
                    v59 := GenS6{F0: 2147483647}
                    _ = v59
                } else {
                    v4 -= v4
                    var v60 uint = 165
                    for fi61 := 0; fi61 < 15; fi61++ {
                        v5 = append(v5, v1)
                        v62 := GenS4{F0: (v4 >> ((18446744073709551615) & 63)), F1: false}
                        v1 = v1
                        _ = fi61
                        _ = v62
                    }
                    _ = v60
                }
                if (true && (!(!true))) {
                    for fi63 := 0; fi63 < 16; fi63++ {
                        for fi64 := 0; fi64 < 18; fi64++ {
                            var v65 uint64 = v3
                            var v66 uint8 = 42
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = fi64
                            _ = v65
                            _ = v66
                        }
                        var v67 uint16 = 1
                        v68 := GenS3{F0: (v4 / (((v4 - (v4 ^ 96))) | 1))}
                        _ = fi63
                        _ = v67
                        _ = v68
                    }
                    v2 = 978
                    for fi69 := 0; fi69 < 14; fi69++ {
                        v4 += (^(v4))
                        _ = fi69
                    }
                }
                _ = v58
            }
            _ = v53
        }
        if true {
            v43 = (^(v43))
        }
        if false {
            var v70 int64 = 1
            for fi71 := 0; fi71 < 6; fi71++ {
                var v72 uint16 = 1
                var v73 uint8 = 127
                _ = fi71
                _ = v72
                _ = v73
            }
            var v74 uint = 18446744073709551615
            _ = v70
            _ = v74
        }
    }
    v75 := GenS3{F0: v4}
    v5 = append(v5, (v1 + v1))
    v3 = v3
    v1 = "hello"
    if false {
        v43 += v43
        var v76 uint64 = (v3 - (v3 << ((1) & 63)))
        _ = v76
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v75
}
