package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int64; F1 string; F2 bool }
type GenS6 struct { F0 uint8; F1 int64; F2 uint8 }

func main() {
    var v1 string = ""
    v2 := []string{(((v1 + v1) + ((((v1 + v1) + v1) + ("foo bar" + (v1 + (v1 + ((((v1 + "\n\t") + "foo bar") + v1) + (v1 + ("foo bar" + v1))))))) + (v1 + "foo bar"))) + v1), "hello", v1}
    var v3 uint = 371
    var v4 int64 = 42
    fmt.Println(v1, v3)
    fmt.Println(v1, v1, v3)
    switch 128 {
    case -2147483648:
        v2 = append(v2, v1)
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v5 uint = v3
    if len(v2) > 0 {
        v2[0] = (((v1 + v1) + v1) + (v1 + v1))
    }
    v4 = -9223372036854775808
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 uint64 = 18446744073709551615
        _ = fi6
        _ = v7
    }
    switch -430 {
    case -402:
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 uint64 = 1
            if ((true || (77 == uint8(v9))) && (!(0 >= 2))) {
                var v10 uint16 = 460
                if (42 < 579) {
                    for fi11 := 0; fi11 < 7; fi11++ {
                        var v12 float64 = -2.25
                        _ = fi11
                        _ = v12
                    }
                } else {
                    fmt.Println(v9)
                    v9 -= 0
                }
                _ = v10
            }
            if false {
                if len(v2) > 0 {
                    v2[0] = v1
                }
                switch -9223372036854775808 {
                case 32768:
                    if true {
                        fmt.Println(fi8, v9)
                        var v13 int32 = 1
                        _ = v13
                    }
                case 1:
                    for fi14 := 0; fi14 < 7; fi14++ {
                        v5 -= v3
                        var v15 uint = v3
                        fmt.Println(fi8)
                        if (42 < 1) {
                            if (!true) {
                                switch (fi8 - fi8) {
                                case -9223372036854775808:
                                    var v16 uint = 813
                                    _ = v16
                                case 2147483647:
                                    var v17 string = v1
                                    _ = v17
                                case 4294967295:
                                    var v18 uint32 = uint32(v15)
                                    _ = v18
                                default:
                                    v19 := GenS1{F0: 255}
                                    _ = v19
                                }
                                v20 := GenS2{F0: 32767, F1: v4}
                                var v21 int16 = -601
                                v22 := GenS2{F0: 32767, F1: v4}
                                _ = v20
                                _ = v21
                                _ = v22
                            } else {
                                var v23 uint8 = 42
                                v24 := GenS3{F0: -1, F1: v4}
                                v23 = (v23 >> ((v15) & 63))
                                v5 -= (v15 - (^(v15)))
                                _ = v23
                                _ = v24
                            }
                        } else {
                            if true {
                                switch 0 {
                                case 1:
                                    if (!(!(2147483647 != -599))) {
                                        v25 := GenS3{F0: -1, F1: (v4 ^ (v4 % ((-9223372036854775808) | 1)))}
                                        v26 := GenS4{F0: 0}
                                        v27 := GenS3{F0: 32767, F1: -9223372036854775808}
                                        _ = v25
                                        _ = v26
                                        _ = v27
                                    } else {
                                        switch (v1 + (("a" + v1) + v1)) {
                                        case "hello":
                                            v15 = (v5 / ((v15) | 1))
                                        case "\n\t":
                                            v28 := GenS5{F0: (v4 >> (((v3 + v5)) & 63)), F1: "\n\t", F2: true}
                                            _ = v28
                                        }
                                    }
                                case 865:
                                    if (float64(0.1) < ((1e10 + -2.25) - 3.14159)) {
                                        v29 := GenS6{F0: 1, F1: (v4 | (v4 & -665)), F2: 198}
                                        if len(v2) > 0 {
                                            v2[0] = ((("hello" + "foo bar") + "\n\t") + v1)
                                        }
                                        _ = v29
                                    }
                                }
                                switch v1 {
                                case "hello":
                                    v3 -= v3
                                case "":
                                    var v30 uint32 = 218
                                    _ = v30
                                case "foo bar":
                                    fmt.Println(fi14, fi14, v4)
                                default:
                                    if (true && false) {
                                        if (!true) {
                                            var v31 int = -1
                                            _ = v31
                                        } else {
                                            switch v1 {
                                            case "hello":
                                                var v32 float64 = (1.5 / float64(0.1))
                                                _ = v32
                                            case "":
                                                v9 = v9
                                            }
                                            var v33 int8 = -1
                                            _ = v33
                                        }
                                        v4 += (^(v4))
                                    }
                                }
                                v1 = ("a" + v1)
                                switch v1 {
                                case "hello":
                                    switch v1 {
                                    case "foo bar":
                                        if true {
                                            var v34 string = ("a" + v1)
                                            _ = v34
                                        } else {
                                            var v35 uint8 = 130
                                            if true {
                                                fmt.Println(v35)
                                            }
                                            _ = v35
                                        }
                                    case "hello":
                                        if false {
                                            v15 -= v15
                                        }
                                    default:
                                        var v36 int16 = 42
                                        _ = v36
                                    }
                                default:
                                    var v37 int = 642
                                    _ = v37
                                }
                            }
                            fmt.Println(v15, v4)
                            if false {
                                switch (fi8 / ((628) | 1)) {
                                case -1:
                                    v38 := GenS5{F0: (v4 | v4), F1: v1, F2: false}
                                    _ = v38
                                default:
                                    _ = len(v2)
                                }
                                _ = len(v2)
                            }
                        }
                        v4 += 9223372036854775807
                        if true {
                            v3 = v15
                        }
                        var v39 string = "a"
                        v1 = v39
                        fmt.Println(v39, v1, v4)
                        _ = fi14
                        _ = v15
                        _ = v39
                    }
                case 128:
                    v5 += v3
                }
                var v40 int32 = -518
                _ = v40
            } else {
                v9 += v9
            }
            _ = fi8
            _ = v9
        }
    case 4294967295:
        for fi41 := 0; fi41 < 3; fi41++ {
            v5 += v5
            _ = fi41
        }
    case -2147483648:
        if (!(false || true)) {
            if true {
                if false {
                    v4 += (v4 * (v4 << (((v5 / (((^(v5))) | 1))) & 63)))
                }
            } else {
                if true {
                    v1 = (v1 + v1)
                }
                var v42 float32 = 1e10
                for fi43 := 0; fi43 < 14; fi43++ {
                    var v44 uint32 = 1
                    _ = fi43
                    _ = v44
                }
                _ = v42
            }
        }
    }
    for fi45 := 0; fi45 < 14; fi45++ {
        var v46 uint64 = 4294967295
        _ = fi45
        _ = v46
    }
    var v47 uint32 = 1
    var v48 int = 827
    var v49 int16 = 32767
    if len(v2) > 0 {
        v2[0] = ((("foo bar" + v1) + ("\n\t" + "\n\t")) + v1)
    }
    if (!true) {
        v3 = v3
        var v50 uint64 = 18446744073709551615
        _ = v50
    }
    v51 := GenS3{F0: 32767, F1: (v4 + 42)}
    if false {
        v5 = (v5 | (v5 ^ (v3 - v5)))
        for fi52 := 0; fi52 < 4; fi52++ {
            var v53 uint32 = (v47 | 42)
            var v54 int16 = v49
            _ = fi52
            _ = v53
            _ = v54
        }
        if true {
            v2 = append(v2, "a")
            switch v1 {
            case "\n\t":
                v1 = v1
            }
        } else {
            v5 = (v5 << ((v5) & 63))
        }
    }
    var v55 string = ""
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v51
    _ = v55
}
