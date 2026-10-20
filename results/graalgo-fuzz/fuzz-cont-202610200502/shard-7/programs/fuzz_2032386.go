package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint32 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 string; F1 string; F2 int }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int32; F1 uint64 }

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[int]bool)
    var v3 uint16 = 42
    var v4 float32 = (1e10 - 0.0)
    var v5 uint = 1
    v6 := make(map[uint32]int)
    var v7 string = ("foo bar" + "hello")
    var v8 int8 = 100
    v6[0] = 0
    var v9 uint8 = 127
    v4 += (v4 + v4)
    v10 := []int{128, 9223372036854775807, 42}
    var v11 int8 = (v8 / ((v8) | 1))
    var v12 uint8 = v9
    v13 := GenS1{F0: 2, F1: 4294967295, F2: 237}
    for fi14 := 0; fi14 < 14; fi14++ {
        v5 = (v5 % ((v5) | 1))
        _ = fi14
    }
    v15 := GenS1{F0: 783, F1: 294, F2: 42}
    fmt.Println(v1, v9, v5)
    v16 := GenS1{F0: uint32(v9), F1: 558, F2: 4294967295}
    v5 -= 65535
    var v17 int64 = 425
    for fi18 := 0; fi18 < 8; fi18++ {
        var v19 uint8 = 127
        if true {
            fmt.Println(v11)
        } else {
            v1 += 42
            switch (v7 + v7) {
            case "":
                v10 = append(v10, (fi18 - -9223372036854775808))
            case "foo bar":
                v2[fi18] = (fi18 > fi18)
            case "\n\t":
                var v20 int8 = (v8 + v8)
                _ = v20
            default:
                if ((0.0 * v4) < 0.0) {
                    switch fi18 {
                    case 127:
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                    case -9:
                        var v21 uint32 = 543
                        _ = v21
                    case 42:
                        switch fi18 {
                        case 32767:
                            fmt.Println(v11, v7)
                        case 15:
                            switch (fi18 & 32767) {
                            case 1:
                                if (true || (1.5 < 0.0)) {
                                    v19 = 88
                                    for fi22 := 0; fi22 < 1; fi22++ {
                                        if (false && (-211 == 532)) {
                                            v23 := GenS2{F0: (v1 / (((v1 * v1)) | 1)), F1: -32768}
                                            var v24 bool = ((fi22 + fi18) < 128)
                                            if len(v10) > 0 {
                                                v10[0] = -2147483648
                                            }
                                            _ = v23
                                            _ = v24
                                        } else {
                                            var v25 int = 9223372036854775807
                                            _ = v25
                                        }
                                        _ = fi22
                                    }
                                }
                            default:
                                for fi26 := 0; fi26 < 14; fi26++ {
                                    v3 = v3
                                    if ((v5 >= v5) && true) {
                                        v8 += v11
                                        v2[(fi26 & (fi18 | (-(fi26))))] = (true && (fi18 > 2147483647))
                                        v27 := GenS3{F0: ("\n\t" + v7), F1: v7, F2: (fi26 | int(v3))}
                                        switch v7 {
                                        case "\n\t":
                                            var v28 int64 = (-(v17))
                                            _ = v28
                                        case "a":
                                            if true {
                                                if len(v10) > 0 {
                                                    v10[0] = fi18
                                                }
                                                if true {
                                                    v8 += -128
                                                }
                                            } else {
                                                var v29 int = 4294967295
                                                switch int(v3) {
                                                case -514:
                                                    var v30 bool = (!false)
                                                    _ = v30
                                                case -259:
                                                    if false {
                                                        var v31 uint32 = 4294967295
                                                        var v32 int16 = 0
                                                        _ = v31
                                                        _ = v32
                                                    }
                                                }
                                                v33 := GenS4{F0: ((v17 << ((902) & 63)) >= (v17 % ((-786) | 1)))}
                                                _ = v29
                                                _ = v33
                                            }
                                        case "":
                                            var v34 int = (^(fi26))
                                            _ = v34
                                        }
                                        if false {
                                            if false {
                                                v11 += (v11 % (((v11 | v11)) | 1))
                                                v12 += 127
                                            } else {
                                                v3 += v3
                                            }
                                        } else {
                                            switch (fi26 * fi26) {
                                            case -737:
                                                if (!(!false)) {
                                                    v35 := GenS1{F0: 771, F1: 2, F2: 117}
                                                    var v36 uint16 = (v3 >> ((v5) & 63))
                                                    var v37 uint = v5
                                                    _ = v35
                                                    _ = v36
                                                    _ = v37
                                                }
                                            default:
                                                v7 = v7
                                            }
                                            fmt.Println(v11, v3)
                                        }
                                        _ = v27
                                    }
                                    fmt.Println(v4, fi18)
                                    _ = fi26
                                }
                            }
                        case -797:
                            switch "\n\t" {
                            case "foo bar":
                                v19 -= v9
                            case "a":
                                var v38 float32 = (3.14159 + ((v4 + (v4 * (v4 + 0.0))) - -2.25))
                                _ = v38
                            }
                        }
                    }
                    for fi39 := 0; fi39 < 1; fi39++ {
                        var v40 uint32 = 349
                        var v41 int = -1
                        v42 := GenS2{F0: v1, F1: int16(v17)}
                        _ = fi39
                        _ = v40
                        _ = v41
                        _ = v42
                    }
                    if false {
                        switch (fi18 << ((278) & 63)) {
                        case -680:
                            var v43 uint8 = v9
                            _ = v43
                        default:
                            var v44 uint16 = v3
                            _ = v44
                        }
                        v45 := GenS5{F0: 42, F1: (v1 - uint64(v5))}
                        for fi46 := 0; fi46 < 17; fi46++ {
                            if ((-(v11)) > -128) {
                                fmt.Println(fi18, fi46, fi46)
                                var v47 int16 = 32767
                                var v48 uint32 = 2
                                _ = v47
                                _ = v48
                            } else {
                                if false {
                                    if ((^(v9)) == (v12 * v19)) {
                                        if ((v3 * v3) > 255) {
                                            v3 = v3
                                            v49 := GenS2{F0: 473, F1: 1000}
                                            _ = v49
                                        }
                                    }
                                }
                            }
                            var v50 int32 = 1
                            v11 -= -1
                            _ = fi46
                            _ = v50
                        }
                        _ = v45
                    } else {
                        var v51 string = v7
                        v51 = v7
                        _ = v51
                    }
                }
            }
        }
        _ = fi18
        _ = v19
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
