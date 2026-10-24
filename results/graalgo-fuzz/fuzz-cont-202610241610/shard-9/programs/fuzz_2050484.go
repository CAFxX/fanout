package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 int; F1 int16; F2 int64 }
type GenS5 struct { F0 uint16; F1 uint8 }
type GenS6 struct { F0 int32; F1 int32 }
func genFn0(p0 uint, p1 uint8, p2 uint64) int32 {
    var v5 uint8 = (p1 * p1)
    fmt.Println(p0, v5)
    var v6 float32 = 3.14159
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 2147483647
}
func genFn1() bool {
    for fi7 := 0; fi7 < 4; fi7++ {
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 string = ""
            _ = fi8
            _ = v9
        }
        var v10 string = ("hello" + "a")
        switch fi7 {
        case 9223372036854775807:
            for fi11 := 0; fi11 < 3; fi11++ {
                v10 = (v10 + v10)
                var v12 bool = true
                v13 := GenS1{F0: 333}
                _ = fi11
                _ = v12
                _ = v13
            }
        default:
            var v14 uint8 = 128
            _ = v14
        }
        _ = fi7
        _ = v10
    }
    var v15 uint64 = 18446744073709551615
    if true {
        var v16 int16 = 1000
        fmt.Println(v16, v16)
        _ = v16
    }
    _ = v15
    return true
}

func main() {
    v1 := make(map[int]bool)
    var v2 int = 32768
    v3 := []uint32{2, 42}
    v4 := []int{v2, (v2 - v2), (v2 + v2), (v2 % (((v2 + 127)) | 1))}
    switch 42 {
    case -2147483648:
        if (2 <= 530) {
            var v17 int8 = 100
            if true {
                var v18 int8 = -128
                v2 = -1
                _ = v18
            }
            _ = v17
        } else {
            fmt.Println(v2)
            fmt.Println(v2)
        }
    case -9223372036854775808:
        v19 := GenS2{F0: 65535, F1: 65535}
        _ = v19
    }
    if false {
        if ((!(!false)) || (!false)) {
            if false {
                var v20 float32 = 1e10
                switch v2 {
                case -2147483648:
                    v20 += (v20 - v20)
                default:
                    delete(v1, v2)
                }
                switch "a" {
                case "a":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                default:
                    if true {
                        for fi21 := 0; fi21 < 10; fi21++ {
                            v22 := GenS2{F0: 0, F1: 2147483647}
                            _ = fi21
                            _ = v22
                        }
                    }
                }
                _ = v20
            } else {
                switch (v2 >> ((255) & 63)) {
                case 2147483647:
                    v2 = (v2 / (((v2 << ((977) & 63))) | 1))
                case 42:
                    if false {
                        if len(v3) > 0 {
                            v3[0] = 65535
                        }
                    }
                case 0:
                    for fi23 := 0; fi23 < 14; fi23++ {
                        for fi24 := 0; fi24 < 5; fi24++ {
                            var v25 uint16 = 1
                            var v26 uint = 332
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                        switch "\n\t" {
                        case "a":
                            for fi27 := 0; fi27 < 12; fi27++ {
                                v2 = fi27
                                _ = fi27
                            }
                        case "hello":
                            for fi28 := 0; fi28 < 18; fi28++ {
                                v2 -= 2147483647
                                v2 -= fi23
                                _ = fi28
                            }
                        }
                        _ = fi23
                    }
                }
                var v29 uint64 = 1
                switch -1 {
                case -2147483648:
                    v30 := GenS1{F0: (v29 ^ (v29 * v29))}
                    _ = v30
                case 4294967295:
                    v31 := GenS1{F0: (^(v29))}
                    _ = v31
                default:
                    if (!(false || (193 < uint(v29)))) {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v2 -= 9223372036854775807
                        for fi32 := 0; fi32 < 13; fi32++ {
                            v33 := v1[v2]
                            _ = fi32
                            _ = v33
                        }
                    } else {
                        v29 -= v29
                        v34 := v1[13]
                        _ = v34
                    }
                }
                _ = v29
            }
        }
        if false {
            var v35 int8 = 100
            for fi36 := 0; fi36 < 12; fi36++ {
                if false {
                    _ = len(v1)
                    if false {
                        fmt.Println(fi36, v2, v35)
                    }
                    var v37 uint16 = 1
                    _ = v37
                }
                _ = fi36
            }
            _ = v35
        } else {
            v38 := GenS3{F0: 1, F1: 96}
            switch v2 {
            case 9223372036854775807:
                var v39 float32 = 3.14159
                _ = v39
            }
            switch (v2 | v2) {
            case -803:
                for fi40 := 0; fi40 < 10; fi40++ {
                    var v41 int64 = 0
                    for fi42 := 0; fi42 < 9; fi42++ {
                        var v43 float32 = 1e10
                        _ = fi42
                        _ = v43
                    }
                    _ = fi40
                    _ = v41
                }
            default:
                var v44 int16 = 32767
                _ = v44
            }
            _ = v38
        }
    }
    v45 := GenS4{F0: (v2 * v2), F1: -1, F2: 574}
    switch ("hello" + "hello") {
    case "":
        v2 -= (v2 - (-(v2)))
    case "hello":
        v46 := GenS5{F0: 65535, F1: 42}
        _ = v46
    default:
        var v47 uint16 = 2
        _ = v47
    }
    v48 := GenS6{F0: 1, F1: 42}
    if len(v3) > 0 {
        v3[0] = 3
    }
    var v49 uint64 = 4294967295
    v1[(v2 + v2)] = false
    v1[v2] = (2147483648 != 748)
    if false {
        fmt.Println(v49)
        switch v2 {
        case 4294967295:
            var v50 uint64 = (v49 / ((4294967295) | 1))
            _ = v50
        case -103:
            for fi51 := 0; fi51 < 19; fi51++ {
                var v52 uint = 292
                fmt.Println(fi51)
                v49 -= v49
                _ = fi51
                _ = v52
            }
        case 240:
            var v53 uint16 = 0
            _ = v53
        default:
            if ((v2 * (v2 & v2)) > int(v49)) {
                for fi54 := 0; fi54 < 1; fi54++ {
                    switch 32768 {
                    case -439:
                        for fi55 := 0; fi55 < 17; fi55++ {
                            var v56 float64 = 1.5
                            fmt.Println(fi54, v56)
                            _ = fi55
                            _ = v56
                        }
                    default:
                        v2 -= v2
                    }
                    var v57 bool = false
                    _ = fi54
                    _ = v57
                }
                v58 := v1[v2]
                _ = len(v1)
                _ = v58
            }
        }
        if ((!(!("hello" >= "foo bar"))) || (-1 != 127)) {
            var v59 uint = 65535
            var v60 int16 = -1
            fmt.Println(v2, v49)
            _ = v59
            _ = v60
        } else {
            var v61 int16 = 42
            switch "\n\t" {
            case "\n\t":
                v49 = (v49 + v49)
            default:
                v1[v2] = true
            }
            switch "foo bar" {
            case "\n\t":
                fmt.Println(v61, v2)
            default:
                if (79 > 190) {
                    var v62 uint8 = 1
                    v62 += v62
                    if (v62 >= v62) {
                        v2 += v2
                    } else {
                        if false {
                            v2 = v2
                        }
                        for fi63 := 0; fi63 < 15; fi63++ {
                            v64 := GenS2{F0: 703, F1: 0}
                            var v65 int64 = 549
                            var v66 bool = (int64(v49) != v65)
                            _ = fi63
                            _ = v64
                            _ = v65
                            _ = v66
                        }
                        v2 -= (v2 << ((763) & 63))
                    }
                    _ = v62
                }
            }
            _ = v61
        }
    }
    v2 -= (v2 ^ (v2 | v2))
    v67 := GenS5{F0: 2, F1: uint8(v49)}
    for fi68 := 0; fi68 < 18; fi68++ {
        v69 := v1[fi68]
        _ = fi68
        _ = v69
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v48
    _ = v49
    _ = v67
}
