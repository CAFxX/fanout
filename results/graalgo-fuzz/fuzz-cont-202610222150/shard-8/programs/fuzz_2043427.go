package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 int8; F1 uint32; F2 uint16 }
type GenS4 struct { F0 int64; F1 int }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint64; F1 int16 }
type GenS7 struct { F0 uint64; F1 uint64; F2 int64 }
type GenS8 struct { F0 uint64 }
func genFn0(p0 uint64, p1 int16, p2 int64) int16 {
    var v6 int = -322
    v7 := GenS1{F0: 1, F1: 10}
    if (!((440 <= 757) || false)) {
        var v8 float64 = 0.1
        _ = v8
    } else {
        v9 := GenS2{F0: (v6 + v6), F1: 4294967295}
        for fi10 := 0; fi10 < 18; fi10++ {
            var v11 bool = true
            p2 += (p2 / ((p2) | 1))
            fmt.Println(v6, p2)
            _ = fi10
            _ = v11
        }
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (p1 << ((2) & 63))
}
func genFn1() int8 {
    if true {
        var v12 uint16 = 2
        _ = v12
    } else {
        v13 := []int{}
        switch -383 {
        case -558:
            var v14 int32 = 0
            _ = v14
        case 127:
            var v15 uint32 = 2147483647
            _ = v15
        default:
            var v16 float64 = 1.5
            _ = v16
        }
        _ = v13
    }
    var v17 uint = 873
    v18 := GenS3{F0: -92, F1: 869, F2: 32767}
    _ = v17
    _ = v18
    return 55
}

func main() {
    var v1 int = 9223372036854775807
    v2 := make(map[string]int)
    var v3 int16 = -32768
    var v4 uint = 42
    v5 := []int{1, (v1 % ((-142) | 1)), (v1 << ((4294967295) & 63)), v1}
    for fi19 := 0; fi19 < 20; fi19++ {
        v20 := GenS2{F0: -613, F1: 4294967295}
        v21 := GenS2{F0: fi19, F1: 0}
        v22 := GenS2{F0: 4294967295, F1: 42}
        _ = fi19
        _ = v20
        _ = v21
        _ = v22
    }
    var v23 int16 = genFn0(4294967295, (v3 - (v3 - (v3 - v3))), 218)
    for fi24 := 0; fi24 < 13; fi24++ {
        v25 := GenS3{F0: genFn1(), F1: 0, F2: 2}
        v4 = v4
        for fi26 := 0; fi26 < 8; fi26++ {
            var v27 float32 = (float32(0.0) - 1.5)
            var v28 uint64 = 42
            fmt.Println(v23)
            if ((("foo bar" + "hello") + "\n\t") > "hello") {
                v27 = (v27 / float32(float64(v27)))
                if false {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v29 := v2[("\n\t" + "foo bar")]
                    var v30 uint32 = 718
                    _ = v29
                    _ = v30
                }
                switch "hello" {
                case "\n\t":
                    var v31 bool = true
                    _ = v31
                case "a":
                    switch fi24 {
                    case 1:
                        switch "\n\t" {
                        case "\n\t":
                            v27 += (v27 + (0.0 * v27))
                        case "foo bar":
                            v2["foo bar"] = v1
                        case "":
                            fmt.Println(v4, v4, fi24)
                        default:
                            if true {
                                var v32 uint16 = 42
                                if (!(true && true)) {
                                    v33 := GenS4{F0: 42, F1: v1}
                                    v2[("" + ("" + "hello"))] = v1
                                    _ = v33
                                } else {
                                    v32 = 710
                                }
                                _ = v32
                            } else {
                                var v34 int8 = -128
                                _ = v34
                            }
                        }
                    }
                default:
                    fmt.Println(fi26, v4, v23)
                }
            }
            _ = fi26
            _ = v27
            _ = v28
        }
        _ = fi24
        _ = v25
    }
    v35 := GenS3{F0: 126, F1: 1, F2: 343}
    v36 := GenS4{F0: 502, F1: v1}
    v37 := GenS4{F0: 42, F1: v1}
    for fi38 := 0; fi38 < 18; fi38++ {
        v39 := v2["hello"]
        for fi40 := 0; fi40 < 1; fi40++ {
            fmt.Println(v23, v23, v4)
            switch "hello" {
            case "foo bar":
                v41 := GenS3{F0: 100, F1: 42, F2: 113}
                _ = v41
            case "a":
                if len(v5) > 0 {
                    v5[0] = fi38
                }
            default:
                if false {
                    v42 := GenS5{F0: 48}
                    v23 = v23
                    v4 -= (v4 / ((v4) | 1))
                    if true {
                        _ = len(v2)
                    }
                    if (842 <= 4294967295) {
                        v43 := GenS6{F0: 716, F1: (v23 | (v3 / ((genFn0(964, genFn0(131, v3, 9223372036854775807), 789)) | 1)))}
                        v1 += fi40
                        _ = len(v5)
                        _ = v43
                    } else {
                        v1 = -1
                        if len(v5) > 0 {
                            v5[0] = (fi38 + -2147483648)
                        }
                        if true {
                            var v44 int8 = -1
                            v3 -= (v23 ^ 32767)
                            _ = v44
                        }
                    }
                    v1 += (v39 * (fi38 % ((4294967295) | 1)))
                    fmt.Println(v4)
                    _ = v42
                } else {
                    fmt.Println(fi40, fi40, fi38)
                    if false {
                        if true {
                            v5 = append(v5, (fi40 % (((v1 | (fi38 | 32767))) | 1)))
                            if (0.1 > ((3.14159 + 0.1) - 0.1)) {
                                fmt.Println(v4, v3, fi40)
                            } else {
                                delete(v2, ("hello" + "foo bar"))
                            }
                            if (1 > 255) {
                                if false {
                                    v3 = (-(v3))
                                    var v45 int8 = genFn1()
                                    _ = v45
                                }
                                v46 := GenS6{F0: 1, F1: (^(v23))}
                                var v47 int = (fi40 ^ (^(fi40)))
                                _ = v46
                                _ = v47
                            } else {
                                v3 += v3
                                v3 += (v3 + v3)
                                switch "a" {
                                case "a":
                                    delete(v2, "\n\t")
                                }
                                switch "a" {
                                case "a":
                                    fmt.Println(fi40, v23)
                                case "hello":
                                    v48 := GenS7{F0: 930, F1: 407, F2: 253}
                                    _ = v48
                                default:
                                    var v49 int32 = 2147483647
                                    _ = v49
                                }
                            }
                        } else {
                            if true {
                                v3 -= (v23 | -17)
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                        }
                    }
                }
            }
            var v50 uint32 = 65535
            _ = fi40
            _ = v50
        }
        var v51 int8 = -106
        _ = fi38
        _ = v39
        _ = v51
    }
    var v52 uint = v4
    var v53 bool = true
    fmt.Println(v23)
    if v53 {
        var v54 float64 = (0.0 * 1.5)
        _ = v54
    }
    v53 = ((!(!(!true))) && (-2.25 <= (float64(((0.0 + 1e10) - -2.25)) / (1e10 + 3.14159))))
    if v53 {
        switch (^(v1)) {
        case -1:
            switch "" {
            case "a":
                if true {
                    var v55 uint = v4
                    fmt.Println(v55, v53)
                    _ = v55
                } else {
                    switch "hello" {
                    case "a":
                        var v56 float32 = 0.1
                        _ = v56
                    }
                    var v57 uint32 = 4294967295
                    _ = v57
                }
            case "":
                v1 = (v1 + (v1 / ((2147483647) | 1)))
            }
        }
        fmt.Println(v4, v53)
    }
    for fi58 := 0; fi58 < 19; fi58++ {
        for fi59 := 0; fi59 < 14; fi59++ {
            fmt.Println(v4)
            fmt.Println(v52)
            if (true || v53) {
                var v60 uint64 = 1
                v1 = (fi59 * 127)
                _ = v60
            }
            _ = fi59
        }
        if len(v5) > 0 {
            v5[0] = (fi58 >> ((v4) & 63))
        }
        _ = fi58
    }
    switch v1 {
    case -9223372036854775808:
        v61 := GenS8{F0: 18446744073709551615}
        _ = v61
    default:
        var v62 int32 = 42
        _ = v62
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v52
    _ = v53
}
