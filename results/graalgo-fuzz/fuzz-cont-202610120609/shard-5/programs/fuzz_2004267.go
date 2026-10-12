package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 bool; F1 uint8; F2 int64 }
type GenS3 struct { F0 uint8; F1 int8; F2 int }

func main() {
    v1 := []uint32{65535, 2147483647, 4294967295}
    v2 := make(map[int]uint32)
    var v3 uint16 = 2
    var v4 string = ""
    v5 := make(map[uint32]string)
    v2[127] = 961
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v6 := GenS1{F0: 0, F1: 86}
    v3 += v3
    if true {
        if false {
            v4 = v4
            var v7 uint16 = (^(v3))
            _ = v7
        }
        v3 -= v3
        switch ("foo bar" + ("hello" + v4)) {
        case "\n\t":
            var v8 float64 = ((1.5 - (float64(((0.1 - 1.5) - float32(1e10))) * (-2.25 * 3.14159))) + (1.5 + 0.0))
            _ = v8
        case "a":
            if true {
                v9 := GenS1{F0: -478, F1: 255}
                if true {
                    _ = len(v5)
                    var v10 uint32 = 0
                    switch 128 {
                    case 32768:
                        v10 = v10
                    }
                    _ = v10
                }
                if ((128 < 0) || ((!(v3 < v3)) && ((145 > 205) || true))) {
                    var v11 uint8 = 1
                    _ = v11
                } else {
                    if false {
                        fmt.Println(v3)
                        if false {
                            var v12 uint8 = 127
                            _ = v12
                        }
                        if (v4 == "a") {
                            if false {
                                var v13 uint16 = 42
                                v3 = (v13 | v3)
                                _ = v13
                            } else {
                                var v14 int16 = 1000
                                for fi15 := 0; fi15 < 1; fi15++ {
                                    v14 = v14
                                    v16 := GenS1{F0: (fi15 << ((48) & 63)), F1: 48}
                                    _ = fi15
                                    _ = v16
                                }
                                v3 = 1
                                _ = v14
                            }
                            v4 = "hello"
                            fmt.Println(v4, v4)
                        } else {
                            v17 := GenS2{F0: ((v3 == (^(v3))) && false), F1: 2, F2: -1}
                            var v18 int16 = -32768
                            var v19 int16 = 1
                            _ = v17
                            _ = v18
                            _ = v19
                        }
                    } else {
                        v4 = v4
                    }
                    var v20 string = (v4 + v4)
                    _ = v20
                }
                _ = v9
            } else {
                if (9223372036854775807 > 970) {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    switch "a" {
                    case "\n\t":
                        delete(v5, 0)
                    default:
                        v3 = v3
                    }
                    switch ("" + v4) {
                    case "":
                        var v21 uint = 967
                        _ = v21
                    case "\n\t":
                        if len(v1) > 0 {
                            v1[0] = uint32(v3)
                        }
                    case "hello":
                        v4 = v4
                    default:
                        var v22 float32 = -2.25
                        _ = v22
                    }
                    v23 := GenS1{F0: int(v3), F1: 127}
                    _ = v23
                }
            }
        }
    }
    for fi24 := 0; fi24 < 2; fi24++ {
        v25 := GenS2{F0: false, F1: uint8(v3), F2: 1}
        v4 = "hello"
        var v26 int64 = -9223372036854775808
        _ = fi24
        _ = v25
        _ = v26
    }
    if true {
        for fi27 := 0; fi27 < 6; fi27++ {
            v3 += (v3 | 255)
            _ = fi27
        }
        for fi28 := 0; fi28 < 11; fi28++ {
            v29 := GenS1{F0: (fi28 % ((fi28) | 1)), F1: 10}
            _ = fi28
            _ = v29
        }
    } else {
        v3 = 65535
        delete(v2, 0)
    }
    switch "\n\t" {
    case "\n\t":
        var v30 uint32 = 2
        _ = v30
    case "a":
        var v31 float64 = (1.5 - float64(1e10))
        _ = v31
    case "":
        for fi32 := 0; fi32 < 10; fi32++ {
            for fi33 := 0; fi33 < 20; fi33++ {
                fmt.Println(v4)
                var v34 int64 = -337
                v35 := GenS3{F0: 151, F1: -1, F2: (fi32 / (((fi32 / ((588) | 1))) | 1))}
                _ = fi33
                _ = v34
                _ = v35
            }
            if (18446744073709551615 >= 695) {
                for fi36 := 0; fi36 < 16; fi36++ {
                    var v37 int = 1
                    fmt.Println(v37, fi36, v4)
                    v38 := GenS3{F0: 255, F1: 108, F2: fi32}
                    _ = fi36
                    _ = v37
                    _ = v38
                }
                var v39 uint32 = 26
                _ = v39
            }
            v3 = (v3 + 32767)
            _ = fi32
        }
    default:
        v3 = v3
    }
    for fi40 := 0; fi40 < 12; fi40++ {
        v41 := GenS3{F0: 134, F1: 100, F2: (fi40 / ((fi40) | 1))}
        v42 := GenS1{F0: (fi40 * fi40), F1: 174}
        v3 -= (v3 % ((171) | 1))
        _ = fi40
        _ = v41
        _ = v42
    }
    for fi43 := 0; fi43 < 18; fi43++ {
        for fi44 := 0; fi44 < 5; fi44++ {
            fmt.Println(v3, v4)
            fmt.Println(v4, v3)
            var v45 uint = 272
            _ = fi44
            _ = v45
        }
        if false {
            fmt.Println(v4, fi43)
            switch -1 {
            case 128:
                if true {
                    var v46 int = fi43
                    switch v46 {
                    case 9223372036854775807:
                        if (780 >= 4294967295) {
                            var v47 int16 = -635
                            var v48 bool = false
                            _ = v47
                            _ = v48
                        }
                    default:
                        delete(v2, (fi43 * (v46 << ((1) & 63))))
                    }
                    _ = v46
                } else {
                    var v49 int = 89
                    if ((-2.25 > 0.0) || false) {
                        var v50 float32 = 3.14159
                        _ = v50
                    } else {
                        var v51 int8 = 1
                        var v52 float64 = 3.14159
                        _ = v51
                        _ = v52
                    }
                    v3 += (v3 & 326)
                    _ = v49
                }
            case 468:
                var v53 int = 21
                _ = v53
            case 4294967295:
                for fi54 := 0; fi54 < 9; fi54++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v55 := GenS3{F0: 86, F1: 0, F2: (fi43 / ((fi54) | 1))}
                    _ = fi54
                    _ = v55
                }
            default:
                v4 = "\n\t"
            }
        }
        _ = fi43
    }
    var v56 bool = (42 == -9223372036854775808)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v6
}
