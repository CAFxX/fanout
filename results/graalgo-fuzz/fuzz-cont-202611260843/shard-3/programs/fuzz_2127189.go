package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int }
type GenS2 struct { F0 int64; F1 uint; F2 int }
type GenS3 struct { F0 uint32; F1 int; F2 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8; F1 int64 }
type GenS6 struct { F0 int; F1 int16; F2 uint16 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint16; F1 int8 }
type GenS9 struct { F0 bool; F1 uint }

func main() {
    var v1 string = ("" + "foo bar")
    var v2 uint = 4294967295
    v3 := make(map[string]bool)
    v4 := []uint32{2147483648, 4294967295}
    var v5 int64 = -512
    if (3.14159 > float32(-2.25)) {
        var v6 uint = (v2 >> (((^(v2))) & 63))
        _ = v6
    } else {
        _ = len(v3)
        v7 := GenS1{F0: 2147483648, F1: 0, F2: 128}
        for fi8 := 0; fi8 < 9; fi8++ {
            switch "" {
            case "foo bar":
                if false {
                    var v9 float32 = 1e10
                    _ = v9
                } else {
                    var v10 int64 = -1
                    _ = v10
                }
            }
            for fi11 := 0; fi11 < 19; fi11++ {
                v12 := GenS1{F0: 31, F1: 850, F2: (fi8 << (((v2 & 979)) & 63))}
                var v13 bool = false
                var v14 uint16 = 375
                if len(v4) > 0 {
                    v4[0] = 1
                }
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
            _ = fi8
        }
        _ = v7
    }
    var v15 uint = v2
    for fi16 := 0; fi16 < 14; fi16++ {
        for fi17 := 0; fi17 < 7; fi17++ {
            var v18 string = ((("\n\t" + v1) + v1) + v1)
            v19 := GenS1{F0: 2147483647, F1: 1000, F2: 32768}
            v20 := GenS2{F0: -1, F1: v15, F2: (-(fi17))}
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
        }
        fmt.Println(fi16, v5)
        _ = fi16
    }
    if (v2 >= (v2 & 4294967295)) {
        if false {
            switch v1 {
            case "":
                switch 9223372036854775807 {
                case -2147483648:
                    if (42 < 42) {
                        v1 = ((v1 + v1) + v1)
                    } else {
                        for fi21 := 0; fi21 < 18; fi21++ {
                            var v22 uint16 = 1
                            if (171 >= 42) {
                                fmt.Println(v22, v5)
                                var v23 float32 = float32(1.5)
                                v15 = (^(v2))
                                _ = v23
                            }
                            _ = fi21
                            _ = v22
                        }
                    }
                default:
                    var v24 bool = false
                    _ = v24
                }
            default:
                v25 := GenS3{F0: 2147483647, F1: 2147483647, F2: -469}
                _ = v25
            }
        }
    } else {
        v26 := GenS4{F0: 2}
        _ = v26
    }
    if (-6 != 375) {
        switch 535 {
        case 903:
            var v27 float64 = 0.1
            _ = v27
        }
        var v28 uint16 = 0
        _ = len(v3)
        _ = v28
    }
    if (((554 > 1000) && true) || false) {
        v29 := GenS5{F0: 26, F1: 1}
        _ = v29
    }
    v30 := GenS6{F0: 459, F1: -32768, F2: 32767}
    for fi31 := 0; fi31 < 7; fi31++ {
        v15 = (v2 * 65535)
        for fi32 := 0; fi32 < 10; fi32++ {
            if false {
                var v33 int16 = -1
                var v34 uint64 = 42
                v1 = v1
                var v35 int32 = 42
                v15 -= (v2 + v15)
                _ = v33
                _ = v34
                _ = v35
            }
            if false {
                v2 -= v15
                fmt.Println(v15)
                v5 = v5
            }
            var v36 int32 = -128
            _ = fi32
            _ = v36
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi31
    }
    var v37 string = v1
    fmt.Println(v1, v2, v37)
    v37 = "a"
    if (true || false) {
        if false {
            v2 -= 0
            if (true && (!(false && true))) {
                if true {
                    switch -2147483648 {
                    case 420:
                        for fi38 := 0; fi38 < 13; fi38++ {
                            fmt.Println(v2, v37, fi38)
                            _ = fi38
                        }
                    case -9223372036854775808:
                        v5 = (v5 & (v5 >> (((v2 % (((v2 | v2)) | 1))) & 63)))
                    }
                    var v39 float32 = 1.5
                    _ = v39
                } else {
                    if (4294967295 < 42) {
                        v2 -= (v2 ^ (v15 % ((65535) | 1)))
                    } else {
                        v15 += (v15 - v2)
                    }
                    if false {
                        v37 = (v1 + v1)
                        v40 := GenS7{F0: -32768}
                        fmt.Println(v2)
                        _ = v40
                    } else {
                        switch (v37 + v1) {
                        case "":
                            for fi41 := 0; fi41 < 9; fi41++ {
                                fmt.Println(fi41)
                                _ = fi41
                            }
                        case "hello":
                            var v42 int16 = 32767
                            _ = v42
                        }
                    }
                }
                for fi43 := 0; fi43 < 1; fi43++ {
                    for fi44 := 0; fi44 < 17; fi44++ {
                        switch (fi43 - (fi43 / ((fi44) | 1))) {
                        case 2147483647:
                            var v45 bool = false
                            _ = v45
                        case 357:
                            v37 = v37
                        }
                        _ = fi44
                    }
                    _ = fi43
                }
                var v46 int = 32768
                _ = v46
            } else {
                fmt.Println(v2)
                var v47 int16 = 1000
                _ = v47
            }
            v48 := GenS8{F0: 2, F1: 127}
            _ = v48
        } else {
            v5 = (v5 | v5)
        }
        fmt.Println(v15, v15, v1)
        if false {
            v3[(v37 + v37)] = true
        }
    } else {
        for fi49 := 0; fi49 < 17; fi49++ {
            v5 -= -1
            _ = fi49
        }
    }
    v2 = (v2 ^ v2)
    var v50 float64 = 1e10
    v51 := v3[(("foo bar" + ("" + v1)) + v1)]
    switch v37 {
    case "":
        var v52 int16 = -1
        _ = v52
    case "\n\t":
        if ((((65535 > 442) || v51) && (1 != 42)) && true) {
            if v51 {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            v15 = (v15 / (((v2 << ((v15) & 63))) | 1))
        }
    case "a":
        var v53 string = v37
        _ = v53
    }
    _ = len(v4)
    for fi54 := 0; fi54 < 17; fi54++ {
        _ = len(v3)
        if v51 {
            var v55 uint16 = 612
            v56 := GenS2{F0: v5, F1: (v2 ^ (v15 ^ (v2 & v15))), F2: (fi54 / ((-778) | 1))}
            if true {
                v57 := GenS3{F0: 630, F1: fi54, F2: 42}
                v4 = append(v4, 2147483648)
                if (1.5 != 1.5) {
                    var v58 uint32 = 4294967295
                    v59 := GenS9{F0: (332 >= 0), F1: 1}
                    _ = v58
                    _ = v59
                }
                _ = v57
            } else {
                if true {
                    for fi60 := 0; fi60 < 8; fi60++ {
                        v61 := GenS8{F0: v55, F1: -93}
                        v2 -= 2
                        var v62 int16 = 42
                        _ = fi60
                        _ = v61
                        _ = v62
                    }
                    if (0.0 < (float32((v50 / float64(-2.25))) - 1.5)) {
                        if (!false) {
                            if false {
                                switch v1 {
                                case "\n\t":
                                    fmt.Println(v2)
                                case "foo bar":
                                    v1 = (v37 + v37)
                                }
                            }
                            v1 = "\n\t"
                        } else {
                            v37 = ""
                            if true {
                                var v63 float32 = 1e10
                                v1 = v37
                                _ = v63
                            } else {
                                var v64 uint8 = 67
                                var v65 float32 = float32((v50 + v50))
                                var v66 string = "hello"
                                _ = v64
                                _ = v65
                                _ = v66
                            }
                            var v67 int64 = (-(v5))
                            _ = v67
                        }
                        var v68 uint8 = 128
                        v2 -= 2
                        _ = v68
                    }
                }
            }
            _ = v55
            _ = v56
        }
        var v69 float32 = 0.1
        _ = fi54
        _ = v69
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v37
    _ = v4
    _ = v5
    _ = v50
    _ = v51
}
