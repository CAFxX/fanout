package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 string; F1 uint32; F2 uint16 }
func genFn0(p0 bool) int {
    if p0 {
        v6 := []int{}
        _ = v6
    }
    switch -613 {
    case -1:
        fmt.Println(p0)
    case -780:
        fmt.Println(p0)
    case 128:
        if (1000 < 42) {
            fmt.Println(p0)
            if (!p0) {
                var v7 float64 = (1.5 / float64(3.14159))
                var v8 bool = false
                _ = v7
                _ = v8
            }
            fmt.Println(p0)
        }
    default:
        var v9 uint8 = 255
        _ = v9
    }
    _ = p0
    return -1
}

func main() {
    var v1 uint8 = 2
    var v2 bool = false
    v3 := make(map[uint32]uint32)
    v4 := []int{32767}
    var v5 string = "foo bar"
    v10 := GenS1{F0: 82}
    v11 := GenS1{F0: 696}
    v12 := GenS2{F0: 42, F1: 41}
    var v13 uint8 = (v1 - (v1 | v1))
    var v14 bool = (32768 < 32767)
    v13 += 1
    v5 = ((((v5 + v5) + v5) + v5) + v5)
    switch 32768 {
    case 1:
        switch (v5 + (("\n\t" + "\n\t") + (("a" + v5) + "\n\t"))) {
        case "\n\t":
            switch genFn0(v2) {
            case -9223372036854775808:
                switch 413 {
                case 158:
                    switch ("hello" + (v5 + v5)) {
                    case "":
                        fmt.Println(v5)
                    case "foo bar":
                        v4 = append(v4, 42)
                    default:
                        v2 = (-523 >= 32767)
                    }
                default:
                    switch v5 {
                    case "hello":
                        var v15 uint8 = (^(v1))
                        _ = v15
                    case "a":
                        if v14 {
                            v16 := GenS1{F0: 1}
                            _ = v16
                        }
                    default:
                        v17 := GenS2{F0: -1, F1: 42}
                        _ = v17
                    }
                }
            case 117:
                v2 = v14
            }
        case "":
            v18 := GenS2{F0: 302, F1: 1000}
            _ = v18
        default:
            if ((-2147483648 == 42) || (0.1 != 0.1)) {
                if true {
                    v14 = false
                    if (3.14159 != 1.5) {
                        switch "" {
                        case "a":
                            if (1000 != -73) {
                                v2 = v14
                                fmt.Println(v14)
                            }
                        case "":
                            fmt.Println(v1)
                        case "foo bar":
                            fmt.Println(v14, v1, v13)
                        }
                        if v2 {
                            fmt.Println(v14, v2, v2)
                            if (true && (-1 >= 42)) {
                                var v19 int8 = 1
                                v3[42] = 2147483647
                                _ = v19
                            } else {
                                for fi20 := 0; fi20 < 7; fi20++ {
                                    if len(v4) > 0 {
                                        v4[0] = fi20
                                    }
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    v5 = v5
                                    _ = fi20
                                }
                                v14 = v2
                                fmt.Println(v2, v1)
                            }
                        }
                    } else {
                        if false {
                            if true {
                                for fi21 := 0; fi21 < 1; fi21++ {
                                    v14 = v14
                                    var v22 float64 = -2.25
                                    v13 = v1
                                    _ = fi21
                                    _ = v22
                                }
                            } else {
                                var v23 int32 = 42
                                _ = v23
                            }
                            v2 = (v14 && v14)
                        }
                        switch -111 {
                        case 127:
                            var v24 uint = 42
                            _ = v24
                        case 2147483647:
                            var v25 int8 = -128
                            _ = v25
                        default:
                            for fi26 := 0; fi26 < 1; fi26++ {
                                if v14 {
                                    if true {
                                        var v27 uint8 = v13
                                        if false {
                                            v4 = append(v4, (-(fi26)))
                                            switch (v5 + v5) {
                                            case "\n\t":
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            case "a":
                                                var v28 int8 = -1
                                                _ = v28
                                            case "foo bar":
                                                if (!((v13 * v1) < v27)) {
                                                    v1 += (^(v1))
                                                }
                                            default:
                                                v1 = (v13 % (((v27 % ((v13) | 1))) | 1))
                                            }
                                            var v29 uint32 = 249
                                            _ = v29
                                        }
                                        v27 -= (v13 ^ 128)
                                        _ = v27
                                    }
                                } else {
                                    fmt.Println(v14, v14)
                                    v1 = 2
                                    v30 := GenS1{F0: 2}
                                    _ = v30
                                }
                                v31 := v3[0]
                                var v32 float32 = 1.5
                                _ = fi26
                                _ = v31
                                _ = v32
                            }
                        }
                        fmt.Println(v13, v14, v5)
                    }
                } else {
                    for fi33 := 0; fi33 < 12; fi33++ {
                        var v34 bool = v2
                        _ = fi33
                        _ = v34
                    }
                }
                if v2 {
                    v3[1] = 4294967295
                    var v35 float64 = -2.25
                    _ = v35
                }
            }
        }
    case 0:
        v5 = (v5 + "hello")
    default:
        v4 = append(v4, genFn0(((float32(3.14159) + (3.14159 - 3.14159)) > 0.0)))
    }
    v36 := GenS3{F0: v5, F1: 1, F2: 953}
    v5 = ((v5 + (((v5 + "foo bar") + "") + v5)) + "a")
    if true {
        for fi37 := 0; fi37 < 16; fi37++ {
            for fi38 := 0; fi38 < 16; fi38++ {
                if (1 < 1) {
                    fmt.Println(v14)
                    v39 := GenS2{F0: 2147483647, F1: 1}
                    _ = v39
                }
                if v2 {
                    delete(v3, 42)
                    v40 := v3[999]
                    _ = v40
                }
                if true {
                    v5 = v5
                    v13 = 2
                } else {
                    var v41 int64 = 193
                    v42 := GenS1{F0: 65535}
                    fmt.Println(fi37, v2, v1)
                    _ = v41
                    _ = v42
                }
                delete(v3, 2147483648)
                _ = fi38
            }
            if true {
                for fi43 := 0; fi43 < 18; fi43++ {
                    var v44 uint16 = 32767
                    _ = fi43
                    _ = v44
                }
                delete(v3, 2147483647)
            }
            var v45 float32 = -2.25
            _ = fi37
            _ = v45
        }
        switch "\n\t" {
        case "a":
            if (127 != 9223372036854775807) {
                fmt.Println(v14)
                switch (v5 + v5) {
                case "a":
                    for fi46 := 0; fi46 < 6; fi46++ {
                        for fi47 := 0; fi47 < 14; fi47++ {
                            v48 := GenS2{F0: -727, F1: 1}
                            var v49 int64 = 9223372036854775807
                            var v50 uint8 = v13
                            _ = fi47
                            _ = v48
                            _ = v49
                            _ = v50
                        }
                        _ = fi46
                    }
                case "\n\t":
                    for fi51 := 0; fi51 < 19; fi51++ {
                        if true {
                            for fi52 := 0; fi52 < 17; fi52++ {
                                switch (fi51 / ((32768) | 1)) {
                                case 1:
                                    var v53 uint = 888
                                    _ = v53
                                case 180:
                                    if v2 {
                                        v14 = (true || false)
                                    }
                                case -483:
                                    v1 -= v1
                                default:
                                    v1 = v13
                                }
                                _ = fi52
                            }
                            for fi54 := 0; fi54 < 11; fi54++ {
                                var v55 int64 = 0
                                if v14 {
                                    var v56 float32 = 3.14159
                                    _ = v56
                                } else {
                                    delete(v3, 701)
                                    fmt.Println(v5, v2)
                                }
                                if v2 {
                                    if ((((!((float64(3.14159) == 0.1) && v14)) || v14) && true) || false) {
                                        _ = len(v4)
                                    }
                                } else {
                                    v2 = ((((v5 + v5) + v5) + v5) < v5)
                                    var v57 int8 = int8(fi54)
                                    _ = v57
                                }
                                _ = fi54
                                _ = v55
                            }
                        }
                        _ = fi51
                    }
                case "":
                    switch "a" {
                    case "hello":
                        for fi58 := 0; fi58 < 9; fi58++ {
                            var v59 string = (("\n\t" + v5) + (v5 + v5))
                            if len(v4) > 0 {
                                v4[0] = -2147483648
                            }
                            v60 := GenS1{F0: 362}
                            _ = fi58
                            _ = v59
                            _ = v60
                        }
                    case "":
                        var v61 float32 = 0.0
                        _ = v61
                    case "\n\t":
                        if v14 {
                            if false {
                                switch "" {
                                case "":
                                    v13 += (v1 | v13)
                                }
                                var v62 int64 = -1
                                _ = v62
                            } else {
                                var v63 int8 = 0
                                for fi64 := 0; fi64 < 15; fi64++ {
                                    v1 += 0
                                    var v65 int = fi64
                                    _ = fi64
                                    _ = v65
                                }
                                _ = v63
                            }
                            var v66 uint = 42
                            _ = v66
                        }
                    }
                default:
                    var v67 uint8 = (v1 - 108)
                    _ = v67
                }
            }
        }
    } else {
        fmt.Println(v5, v13)
        var v68 uint64 = 0
        _ = v68
    }
    for fi69 := 0; fi69 < 14; fi69++ {
        v14 = v2
        switch "hello" {
        case "hello":
            if (!(42 > 285)) {
                v5 = ((("a" + ((v5 + ("" + v5)) + v5)) + v5) + ((("hello" + v5) + v5) + v5))
                v5 = v5
            }
        default:
            var v70 uint64 = 18446744073709551615
            _ = v70
        }
        _ = fi69
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
