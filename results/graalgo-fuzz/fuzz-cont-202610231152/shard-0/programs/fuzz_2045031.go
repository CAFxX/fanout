package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 int; F2 bool }
type GenS3 struct { F0 int32; F1 int8; F2 uint16 }
type GenS4 struct { F0 uint16; F1 string; F2 uint8 }
type GenS5 struct { F0 bool; F1 uint }
type GenS6 struct { F0 int64; F1 int32 }
type GenS7 struct { F0 int16; F1 int64 }
type GenS8 struct { F0 uint; F1 int8; F2 uint16 }

func main() {
    v1 := []uint32{42}
    var v2 uint16 = 338
    var v3 float32 = -2.25
    for fi4 := 0; fi4 < 4; fi4++ {
        if len(v1) > 0 {
            v1[0] = 2147483648
        }
        v2 = (v2 << ((6) & 63))
        _ = fi4
    }
    var v5 int8 = 123
    if true {
        if true {
            var v6 int32 = 95
            var v7 uint64 = 42
            v8 := GenS1{F0: 502}
            _ = v6
            _ = v7
            _ = v8
        }
    }
    v1 = append(v1, 65535)
    if len(v1) > 0 {
        v1[0] = 42
    }
    var v9 uint16 = v2
    if (((false && false) || false) && true) {
        v3 = 1.5
    }
    if true {
        for fi10 := 0; fi10 < 16; fi10++ {
            for fi11 := 0; fi11 < 7; fi11++ {
                fmt.Println(v9, v9, fi10)
                switch -2147483648 {
                case 9223372036854775807:
                    var v12 int64 = -1
                    _ = v12
                case 4294967295:
                    v5 -= (v5 | 1)
                case -9223372036854775808:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                default:
                    v13 := GenS2{F0: 22, F1: fi10, F2: (!false)}
                    _ = v13
                }
                v5 -= v5
                var v14 uint8 = 74
                if len(v1) > 0 {
                    v1[0] = 685
                }
                _ = fi11
                _ = v14
            }
            var v15 float32 = 1e10
            if (1.5 > -2.25) {
                var v16 uint64 = 0
                if (1000 <= 5) {
                    var v17 uint32 = 2
                    for fi18 := 0; fi18 < 2; fi18++ {
                        v19 := GenS3{F0: 0, F1: (-(v5)), F2: v2}
                        var v20 uint32 = (^(v17))
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                    if false {
                        v9 -= (v2 % (((v9 << ((uint(v17)) & 63))) | 1))
                        switch -9223372036854775808 {
                        case 0:
                            if (((("\n\t" + ("a" + "hello")) + ("a" + "foo bar")) + "\n\t") >= "hello") {
                                if (1e10 != float64(0.1)) {
                                    var v21 uint64 = v16
                                    _ = len(v1)
                                    _ = v21
                                }
                                if false {
                                    fmt.Println(v17, v16, v15)
                                } else {
                                    v17 -= (v17 >> ((2) & 63))
                                }
                            } else {
                                if (!false) {
                                    fmt.Println(v17)
                                    v22 := GenS2{F0: 128, F1: (^(fi10)), F2: false}
                                    _ = v22
                                }
                                switch (fi10 + (fi10 << ((752) & 63))) {
                                case 824:
                                    for fi23 := 0; fi23 < 5; fi23++ {
                                        if false {
                                            v9 = 1
                                            fmt.Println(v9)
                                            switch fi10 {
                                            case -380:
                                                fmt.Println(v15, fi10, v9)
                                            case 32768:
                                                switch "\n\t" {
                                                case "hello":
                                                    v16 = v16
                                                case "\n\t":
                                                    v9 = 32767
                                                case "a":
                                                    var v24 int16 = -993
                                                    _ = v24
                                                }
                                            case 127:
                                                var v25 int = (fi23 & (fi10 / ((int(v5)) | 1)))
                                                _ = v25
                                            }
                                        } else {
                                            if (9223372036854775807 == 9223372036854775807) {
                                                fmt.Println(v17, v16)
                                                if len(v1) > 0 {
                                                    v1[0] = (v17 & (v17 & (v17 | uint32(v16))))
                                                }
                                            } else {
                                                _ = len(v1)
                                            }
                                            if (!false) {
                                                if true {
                                                    if (true || false) {
                                                        switch ("foo bar" + "") {
                                                        case "":
                                                            v2 = (v9 << ((2) & 63))
                                                        default:
                                                            if len(v1) > 0 {
                                                                v1[0] = (v17 % ((750) | 1))
                                                            }
                                                        }
                                                        var v26 bool = (42 == 42)
                                                        _ = v26
                                                    }
                                                    fmt.Println(v17, v5, v2)
                                                } else {
                                                    if len(v1) > 0 {
                                                        v1[0] = v17
                                                    }
                                                    var v27 uint = 255
                                                    _ = v27
                                                }
                                                var v28 int64 = -9223372036854775808
                                                var v29 uint = 4294967295
                                                _ = v28
                                                _ = v29
                                            }
                                            v30 := GenS3{F0: 0, F1: 42, F2: 486}
                                            _ = v30
                                        }
                                        v2 = (v2 * (v2 - v2))
                                        switch 969 {
                                        case 238:
                                            v15 -= v3
                                        case 4294967295:
                                            v2 += v9
                                        }
                                        v15 = v3
                                        _ = fi23
                                    }
                                case -2147483648:
                                    v31 := GenS2{F0: 198, F1: (fi10 / (((fi10 * -342)) | 1)), F2: false}
                                    _ = v31
                                case -9223372036854775808:
                                    if false {
                                        v32 := GenS4{F0: v9, F1: ("hello" + ""), F2: 128}
                                        _ = v32
                                    } else {
                                        for fi33 := 0; fi33 < 19; fi33++ {
                                            v34 := GenS5{F0: (-1 == 1), F1: 430}
                                            fmt.Println(v17)
                                            _ = fi33
                                            _ = v34
                                        }
                                        switch (fi10 % ((128) | 1)) {
                                        case 730:
                                            v35 := GenS4{F0: v9, F1: "hello", F2: 1}
                                            _ = v35
                                        case 32768:
                                            var v36 uint16 = (v9 * v2)
                                            _ = v36
                                        }
                                    }
                                }
                            }
                        }
                        if (42 >= -2147483648) {
                            var v37 uint64 = v16
                            for fi38 := 0; fi38 < 11; fi38++ {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                switch (fi38 + (fi38 % ((fi10) | 1))) {
                                case -243:
                                    var v39 int = fi38
                                    _ = v39
                                case 32768:
                                    var v40 int32 = 2147483647
                                    _ = v40
                                default:
                                    fmt.Println(fi10, v15, v5)
                                }
                                _ = fi38
                            }
                            _ = v37
                        } else {
                            var v41 int64 = -139
                            var v42 int = int(v17)
                            if len(v1) > 0 {
                                v1[0] = 4294967295
                            }
                            _ = v41
                            _ = v42
                        }
                    }
                    _ = v17
                } else {
                    _ = len(v1)
                    _ = len(v1)
                }
                _ = v16
            } else {
                switch -960 {
                case 2147483647:
                    var v43 uint64 = 18446744073709551615
                    _ = v43
                case 127:
                    v5 = v5
                default:
                    var v44 string = ("\n\t" + "foo bar")
                    _ = v44
                }
            }
            _ = fi10
            _ = v15
        }
        v1 = append(v1, 2147483647)
        v45 := GenS6{F0: -9223372036854775808, F1: 0}
        _ = v45
    }
    if false {
        v1 = append(v1, 0)
        switch 62 {
        case 32767:
            for fi46 := 0; fi46 < 20; fi46++ {
                v47 := GenS3{F0: 0, F1: (^(v5)), F2: 746}
                if (0 <= 69) {
                    v2 = 278
                    var v48 int64 = -918
                    _ = v48
                }
                if (true || false) {
                    if true {
                        v49 := GenS7{F0: 810, F1: -1}
                        var v50 uint32 = 2147483647
                        for fi51 := 0; fi51 < 6; fi51++ {
                            if (165 < -1) {
                                var v52 int8 = (v5 / ((v5) | 1))
                                _ = v52
                            }
                            _ = len(v1)
                            v50 = (v50 * v50)
                            fmt.Println(fi51, v2, v3)
                            _ = fi51
                        }
                        _ = v49
                        _ = v50
                    }
                }
                _ = fi46
                _ = v47
            }
        case 2147483647:
            v9 = (v2 | (v2 << ((810) & 63)))
        case 162:
            fmt.Println(v9, v5, v5)
        }
        if false {
            v53 := GenS8{F0: 18446744073709551615, F1: v5, F2: 2}
            _ = v53
        } else {
            v3 = float32(1e10)
        }
    }
    for fi54 := 0; fi54 < 3; fi54++ {
        switch 42 {
        case -519:
            var v55 uint64 = 433
            _ = v55
        default:
            v2 = v2
        }
        v56 := GenS4{F0: v9, F1: ("" + ("" + "hello")), F2: 128}
        _ = fi54
        _ = v56
    }
    var v57 uint8 = 206
    var v58 uint = 122
    v57 += 128
    v59 := GenS6{F0: 0, F1: 0}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v57
    _ = v58
    _ = v59
    _ = v9
}
