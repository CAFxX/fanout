package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int; F1 bool }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint32; F1 uint8 }
type GenS8 struct { F0 int64; F1 bool; F2 int64 }
func genFn0(p0 uint, p1 bool) int16 {
    p0 += 255
    var v6 bool = p1
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 int64 = 0
        var v9 uint8 = 2
        v6 = v6
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    return 1000
}

func main() {
    var v1 uint16 = 889
    var v2 int64 = 9223372036854775807
    v3 := make(map[string]uint32)
    var v4 uint32 = 688
    v5 := []string{"a"}
    for fi10 := 0; fi10 < 4; fi10++ {
        if ("foo bar" == "foo bar") {
            var v11 string = "foo bar"
            for fi12 := 0; fi12 < 15; fi12++ {
                switch "hello" {
                case "foo bar":
                    v4 -= 511
                case "hello":
                    v2 = (v2 ^ v2)
                default:
                    var v13 int = 32768
                    v14 := GenS1{F0: 2, F1: genFn0(476, (0.0 >= 1.5))}
                    _ = v13
                    _ = v14
                }
                _ = fi12
            }
            if (!(0 >= 352)) {
                var v15 uint64 = 2
                fmt.Println(v2)
                _ = v15
            } else {
                v16 := GenS1{F0: 956, F1: 0}
                if false {
                    switch (v11 + v11) {
                    case "a":
                        fmt.Println(v4, v1, v1)
                    case "":
                        var v17 int32 = -573
                        _ = v17
                    case "hello":
                        var v18 int8 = 82
                        _ = v18
                    }
                    v19 := GenS2{F0: v2, F1: 0}
                    _ = v19
                } else {
                    for fi20 := 0; fi20 < 11; fi20++ {
                        if (true && (!true)) {
                            if true {
                                v21 := GenS2{F0: (v2 % (((v2 | (v2 << ((uint(v1)) & 63)))) | 1)), F1: -85}
                                switch (((v11 + ((v11 + ((v11 + v11) + "hello")) + v11)) + "hello") + v11) {
                                case "":
                                    var v22 string = v11
                                    _ = v22
                                case "\n\t":
                                    v23 := GenS2{F0: (v2 | -1), F1: int8(v4)}
                                    _ = len(v3)
                                    _ = v23
                                default:
                                    if len(v5) > 0 {
                                        v5[0] = "foo bar"
                                    }
                                }
                                _ = v21
                            }
                            v2 -= v2
                            if false {
                                fmt.Println(fi20, v4, fi10)
                                var v24 uint8 = 128
                                v25 := GenS1{F0: 0, F1: -361}
                                _ = v24
                                _ = v25
                            } else {
                                switch fi10 {
                                case 128:
                                    fmt.Println(v4, v2)
                                case 127:
                                    fmt.Println(v1, v2)
                                }
                            }
                        }
                        v26 := v3[v11]
                        v4 = (v4 - (^(v4)))
                        fmt.Println(fi20, v11)
                        v27 := GenS1{F0: 139, F1: genFn0(368, false)}
                        if (fi20 > int(v2)) {
                            v28 := GenS1{F0: 878, F1: -1}
                            var v29 int8 = -44
                            v1 += (^(v1))
                            _ = v28
                            _ = v29
                        } else {
                            var v30 uint32 = v26
                            fmt.Println(v1, v4)
                            v31 := GenS3{F0: (v26 + v26)}
                            _ = v30
                            _ = v31
                        }
                        v32 := GenS4{F0: (fi10 - (fi10 ^ (fi10 << ((0) & 63))))}
                        var v33 uint16 = v1
                        _ = fi20
                        _ = v26
                        _ = v27
                        _ = v32
                        _ = v33
                    }
                }
                v34 := GenS3{F0: (v4 % ((v4) | 1))}
                _ = v16
                _ = v34
            }
            _ = v11
        }
        _ = fi10
    }
    for fi35 := 0; fi35 < 8; fi35++ {
        for fi36 := 0; fi36 < 14; fi36++ {
            switch ("" + "") {
            case "":
                if true {
                    v1 += 32767
                } else {
                    delete(v3, "\n\t")
                    v3[("\n\t" + ("hello" + ("hello" + ("\n\t" + "foo bar"))))] = v4
                }
            case "\n\t":
                if len(v5) > 0 {
                    v5[0] = ("a" + "a")
                }
            }
            var v37 int32 = -1
            _ = fi36
            _ = v37
        }
        _ = fi35
    }
    var v38 string = "foo bar"
    switch "foo bar" {
    case "a":
        switch v38 {
        case "":
            _ = len(v5)
        case "\n\t":
            v39 := GenS5{F0: 32768, F1: true}
            _ = v39
        default:
            fmt.Println(v38)
        }
    case "hello":
        var v40 int16 = -32768
        _ = v40
    }
    switch "" {
    case "a":
        v2 = v2
    }
    switch 32767 {
    case 32768:
        v4 = (^(v4))
    case 2147483647:
        if ((32767 != v1) && (0.1 > 1.5)) {
            for fi41 := 0; fi41 < 9; fi41++ {
                var v42 float64 = (0.1 + (1e10 / 1.0))
                _ = fi41
                _ = v42
            }
            _ = len(v5)
            for fi43 := 0; fi43 < 17; fi43++ {
                if (42 != 108) {
                    if ((v1 + v1) != v1) {
                        switch ("hello" + v38) {
                        case "hello":
                            v38 = v38
                        case "":
                            fmt.Println(v1)
                        case "a":
                            fmt.Println(v38, v4)
                        }
                        v38 = ("foo bar" + v38)
                        v1 += v1
                    } else {
                        for fi44 := 0; fi44 < 10; fi44++ {
                            switch (fi44 | (fi43 << ((18446744073709551615) & 63))) {
                            case -745:
                                v45 := GenS4{F0: fi43}
                                _ = v45
                            case -823:
                                v4 = (^(v4))
                            case 32767:
                                var v46 uint = 497
                                _ = v46
                            }
                            if len(v5) > 0 {
                                v5[0] = v38
                            }
                            v1 = v1
                            _ = fi44
                        }
                        if (false || (1.5 >= 1e10)) {
                            var v47 int = (-(fi43))
                            var v48 uint = 50
                            v47 = 1
                            _ = v47
                            _ = v48
                        }
                        for fi49 := 0; fi49 < 13; fi49++ {
                            var v50 uint16 = (v1 >> ((0) & 63))
                            _ = fi49
                            _ = v50
                        }
                    }
                    var v51 int = fi43
                    var v52 int8 = 127
                    _ = v51
                    _ = v52
                } else {
                    v1 = v1
                    var v53 uint16 = v1
                    for fi54 := 0; fi54 < 16; fi54++ {
                        v55 := v3[v38]
                        _ = fi54
                        _ = v55
                    }
                    _ = v53
                }
                switch "hello" {
                case "a":
                    var v56 uint8 = 0
                    _ = v56
                case "\n\t":
                    fmt.Println(v4, v4, v4)
                case "":
                    v1 -= v1
                }
                _ = fi43
            }
        }
    }
    v57 := GenS6{F0: -128}
    v58 := GenS7{F0: v4, F1: 2}
    _ = len(v5)
    fmt.Println(v38, v4, v1)
    var v59 uint16 = uint16(v4)
    v5 = append(v5, v38)
    v60 := GenS8{F0: int64(v4), F1: (true && false), F2: (v2 >> ((0) & 63))}
    v61 := v3[((v38 + "") + ("hello" + "foo bar"))]
    for fi62 := 0; fi62 < 19; fi62++ {
        v1 += v59
        _ = fi62
    }
    v61 += 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v59:", v59)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v57
    _ = v58
    _ = v59
    _ = v60
    _ = v61
}
