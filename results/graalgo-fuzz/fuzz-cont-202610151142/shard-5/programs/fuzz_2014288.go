package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 uint64; F2 uint32 }
type GenS4 struct { F0 uint16; F1 uint16; F2 uint }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int; F1 uint8 }
type GenS7 struct { F0 uint64; F1 uint64; F2 int }
func genFn0() uint {
    var v6 int32 = -1
    if false {
        var v7 uint64 = 4294967295
        var v8 float64 = float64(0.1)
        _ = v7
        _ = v8
    } else {
        var v9 int8 = 0
        _ = v9
    }
    switch "" {
    case "foo bar":
        var v10 string = "hello"
        _ = v10
    case "a":
        for fi11 := 0; fi11 < 10; fi11++ {
            if (100 >= -99) {
                v12 := GenS1{F0: ("foo bar" + (("foo bar" + ("a" + "hello")) + ("foo bar" + ""))), F1: -667, F2: 0}
                var v13 int16 = 0
                _ = v12
                _ = v13
            } else {
                fmt.Println(v6, fi11)
                v14 := GenS1{F0: "", F1: 0, F2: 0}
                var v15 float64 = 0.1
                _ = v14
                _ = v15
            }
            if (32768 >= (fi11 + 32768)) {
                if (!(255 < 148)) {
                    for fi16 := 0; fi16 < 15; fi16++ {
                        fmt.Println(v6, fi16, v6)
                        v17 := GenS2{F0: 1000}
                        var v18 float32 = 1e10
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                } else {
                    if true {
                        if (!(-812 <= -827)) {
                            v6 -= 42
                            var v19 int16 = -385
                            v20 := GenS3{F0: 451, F1: 42, F2: 4294967295}
                            _ = v19
                            _ = v20
                        } else {
                            v6 -= (^(v6))
                            v21 := []int{128, 32768, (-(fi11)), fi11}
                            var v22 string = "\n\t"
                            _ = v21
                            _ = v22
                        }
                        var v23 int64 = 9223372036854775807
                        _ = v23
                    }
                    v6 -= (v6 * 2147483647)
                    v6 -= -2147483648
                }
            }
            _ = fi11
        }
    default:
        for fi24 := 0; fi24 < 18; fi24++ {
            v25 := GenS4{F0: 1, F1: 42, F2: 65535}
            _ = fi24
            _ = v25
        }
    }
    v26 := []int{2147483647}
    var v27 float64 = float64(1e10)
    _ = v26
    _ = v27
    _ = v6
    return 207
}

func main() {
    var v1 int16 = 1000
    var v2 float32 = (3.14159 / (-2.25 * (0.1 / 1.5)))
    var v3 int = 42
    var v4 int64 = 42
    var v5 float64 = 1e10
    switch "foo bar" {
    case "a":
        var v28 uint64 = 1
        _ = v28
    case "hello":
        if ("\n\t" == "") {
            if (!true) {
                v29 := []int{-2147483648, 32768}
                _ = v29
            }
        }
    default:
        var v30 int16 = v1
        _ = v30
    }
    var v31 int32 = -1
    var v32 float32 = ((v2 + (v2 / 1e10)) * (float32(0.1) / 0.1))
    switch "hello" {
    case "":
        v2 = float32((v5 + ((v5 + 0.0) * v5)))
    case "\n\t":
        var v33 uint16 = 1
        _ = v33
    case "foo bar":
        var v34 string = ""
        _ = v34
    }
    v4 = (v4 | (v4 & (v4 + (v4 << ((4294967295) & 63)))))
    switch ("hello" + "a") {
    case "":
        for fi35 := 0; fi35 < 9; fi35++ {
            fmt.Println(v2, fi35, v3)
            if (true || true) {
                v3 -= fi35
            }
            for fi36 := 0; fi36 < 12; fi36++ {
                var v37 int8 = 0
                _ = fi36
                _ = v37
            }
            _ = fi35
        }
    case "a":
        v38 := GenS5{F0: 32767}
        _ = v38
    case "hello":
        var v39 float64 = v5
        _ = v39
    }
    v40 := []int{v3, 42}
    if (!false) {
        if len(v40) > 0 {
            v40[0] = v3
        }
        var v41 float64 = v5
        _ = v41
    } else {
        if ((-1 >= -38) && false) {
            if ("" == ("\n\t" + "hello")) {
                v42 := GenS3{F0: 1, F1: 18446744073709551615, F2: 65535}
                for fi43 := 0; fi43 < 2; fi43++ {
                    v44 := GenS2{F0: (-(v1))}
                    if (((1 == 18446744073709551615) && true) && (0.1 > (float64(0.1) + v5))) {
                        v45 := GenS5{F0: (v1 >> ((42) & 63))}
                        if (true || false) {
                            _ = len(v40)
                            for fi46 := 0; fi46 < 3; fi46++ {
                                switch "foo bar" {
                                case "a":
                                    var v47 int16 = -32768
                                    _ = v47
                                case "hello":
                                    var v48 int32 = (^(v31))
                                    _ = v48
                                case "\n\t":
                                    switch "" {
                                    case "hello":
                                        switch (v3 >> ((71) & 63)) {
                                        case 1:
                                            switch (fi46 % (((-(fi46))) | 1)) {
                                            case 9223372036854775807:
                                                var v49 uint64 = 527
                                                _ = v49
                                            case -1:
                                                if len(v40) >= 2 {
                                                    v40 = v40[1:len(v40)]
                                                }
                                            default:
                                                var v50 uint64 = 2
                                                _ = v50
                                            }
                                        case 9223372036854775807:
                                            if false {
                                                var v51 float32 = ((0.0 + (-2.25 * 0.1)) / (v32 - v32))
                                                _ = v51
                                            }
                                            fmt.Println(fi46)
                                        case -393:
                                            v52 := GenS6{F0: fi46, F1: 0}
                                            fmt.Println(v4, v2)
                                            _ = v52
                                        default:
                                            v2 -= v2
                                        }
                                    default:
                                        var v53 int16 = (v1 + (^(v1)))
                                        _ = v53
                                    }
                                default:
                                    switch ("a" + ("hello" + "")) {
                                    case "":
                                        fmt.Println(v4)
                                    case "a":
                                        _ = len(v40)
                                    case "\n\t":
                                        if len(v40) >= 2 {
                                            v40 = v40[1:len(v40)]
                                        }
                                    }
                                }
                                switch v3 {
                                case 127:
                                    fmt.Println(fi43, v32, v31)
                                case 128:
                                    v54 := GenS6{F0: (v3 << ((18446744073709551615) & 63)), F1: 1}
                                    _ = v54
                                case 42:
                                    v55 := GenS7{F0: 2, F1: 217, F2: (-(v3))}
                                    _ = v55
                                default:
                                    var v56 bool = (false && true)
                                    _ = v56
                                }
                                var v57 int = fi46
                                _ = fi46
                                _ = v57
                            }
                            if len(v40) >= 2 {
                                v40 = v40[1:len(v40)]
                            }
                        } else {
                            for fi58 := 0; fi58 < 19; fi58++ {
                                v59 := GenS6{F0: (fi43 - fi43), F1: 127}
                                var v60 int = fi43
                                var v61 uint32 = 2
                                _ = fi58
                                _ = v59
                                _ = v60
                                _ = v61
                            }
                        }
                        _ = v45
                    }
                    _ = fi43
                    _ = v44
                }
                switch "foo bar" {
                case "a":
                    fmt.Println(v1, v31, v31)
                case "\n\t":
                    for fi62 := 0; fi62 < 4; fi62++ {
                        switch "a" {
                        case "hello":
                            for fi63 := 0; fi63 < 5; fi63++ {
                                fmt.Println(fi63, v2)
                                _ = fi63
                            }
                        case "\n\t":
                            var v64 int64 = (v4 << ((188) & 63))
                            _ = v64
                        case "":
                            if (1 > 2) {
                                v3 = 127
                                v1 = 32767
                            }
                        }
                        if len(v40) > 0 {
                            v40[0] = (v3 ^ (v3 ^ (fi62 % ((fi62) | 1))))
                        }
                        _ = fi62
                    }
                case "":
                    v1 = v1
                }
                _ = v42
            }
            v1 -= v1
        } else {
            for fi65 := 0; fi65 < 10; fi65++ {
                var v66 uint32 = 4294967295
                _ = fi65
                _ = v66
            }
        }
    }
    var v67 uint32 = 356
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v5:", v5)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v40
    _ = v5
    _ = v67
}
