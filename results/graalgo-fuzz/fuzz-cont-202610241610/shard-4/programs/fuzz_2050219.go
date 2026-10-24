package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int16 }
type GenS2 struct { F0 uint; F1 bool; F2 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64; F1 uint16; F2 uint32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int; F1 int32; F2 int }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int64; F1 int; F2 int }
func genFn0(p0 int) int8 {
    p0 = -1
    var v4 int64 = 842
    var v5 string = ("a" + "foo bar")
    v4 += v4
    _ = p0
    _ = v4
    _ = v5
    return 42
}

func main() {
    var v1 uint16 = 42
    v2 := []int{-2147483648, 127}
    v3 := make(map[int]int)
    fmt.Println(v1)
    if (("hello" + "foo bar") <= "a") {
        v6 := GenS1{F0: 0, F1: (42 >= 100), F2: 1}
        v7 := v3[2147483647]
        if (!(!((!(false || true)) || true))) {
            for fi8 := 0; fi8 < 6; fi8++ {
                var v9 float32 = 0.1
                switch (("foo bar" + ("\n\t" + "a")) + ((("" + "foo bar") + "a") + "\n\t")) {
                case "hello":
                    if false {
                        switch 561 {
                        case 127:
                            v9 = v9
                        case -2147483648:
                            if len(v2) > 0 {
                                v2[0] = (^(v7))
                            }
                        default:
                            fmt.Println(v1)
                        }
                    } else {
                        var v10 int8 = 127
                        for fi11 := 0; fi11 < 3; fi11++ {
                            v3[(-(fi8))] = -2147483648
                            var v12 int32 = -1
                            v7 += fi11
                            if (true || false) {
                                fmt.Println(v1, fi8)
                            }
                            _ = fi11
                            _ = v12
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = v10
                    }
                case "\n\t":
                    fmt.Println(v9, v1, v7)
                case "foo bar":
                    v13 := GenS1{F0: -525, F1: false, F2: -32768}
                    _ = v13
                }
                _ = fi8
                _ = v9
            }
            v1 = (v1 + v1)
            v14 := GenS2{F0: 520, F1: ((!(1 != -847)) && true), F2: -9223372036854775808}
            _ = v14
        }
        _ = v6
        _ = v7
    }
    if false {
        v15 := GenS3{F0: 9223372036854775807}
        _ = len(v3)
        for fi16 := 0; fi16 < 8; fi16++ {
            var v17 uint64 = 22
            for fi18 := 0; fi18 < 19; fi18++ {
                fmt.Println(fi18, fi16)
                if false {
                    v17 = (^(v17))
                    fmt.Println(v1, v1, v1)
                } else {
                    switch "a" {
                    case "\n\t":
                        v19 := GenS4{F0: 928, F1: (v1 & (^(v1))), F2: 0}
                        _ = v19
                    case "a":
                        switch ("" + "hello") {
                        case "hello":
                            fmt.Println(v1, v17)
                            v17 = v17
                        case "a":
                            if (!false) {
                                var v20 int64 = -9223372036854775808
                                if true {
                                    switch "hello" {
                                    case "a":
                                        v21 := GenS2{F0: 98, F1: false, F2: (fi16 - int(v17))}
                                        _ = v21
                                    }
                                }
                                _ = v20
                            } else {
                                v22 := GenS2{F0: 2, F1: (1 >= -288), F2: 128}
                                if (!false) {
                                    v23 := GenS5{F0: v17}
                                    v1 = v1
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    v24 := GenS6{F0: (fi16 | 9223372036854775807), F1: 1, F2: (fi18 & (fi16 | fi18))}
                                    v17 += (v17 / (((v17 | (^(v17)))) | 1))
                                    _ = v23
                                    _ = v24
                                }
                                _ = v22
                            }
                        case "foo bar":
                            v1 = v1
                        default:
                            switch fi18 {
                            case 4294967295:
                                fmt.Println(v1, fi16, fi18)
                            case -2147483648:
                                v17 += 820
                            }
                        }
                    }
                    v17 += 501
                }
                v1 = (v1 / ((897) | 1))
                switch "hello" {
                case "foo bar":
                    var v25 string = "a"
                    _ = v25
                case "":
                    switch (fi16 % (((fi16 - fi18)) | 1)) {
                    case -2147483648:
                        v3[(fi18 % ((32768) | 1))] = -9223372036854775808
                        v26 := GenS5{F0: 18446744073709551615}
                        _ = v26
                    case 128:
                        v27 := GenS5{F0: v17}
                        _ = v27
                    case -315:
                        fmt.Println(v1)
                    default:
                        if ((!(704 <= fi18)) || true) {
                            var v28 int16 = 1000
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = v28
                        }
                    }
                case "hello":
                    v17 -= (v17 | (v17 + v17))
                default:
                    v1 += (v1 - v1)
                }
                v29 := GenS1{F0: 1, F1: (1.5 == (1.5 + 1.5)), F2: 32767}
                _ = fi18
                _ = v29
            }
            if false {
                var v30 bool = (2 < v17)
                _ = v30
            } else {
                var v31 int16 = 1000
                _ = v31
            }
            _ = fi16
            _ = v17
        }
        _ = v15
    }
    v1 = (v1 << ((650) & 63))
    switch 429 {
    case -798:
        for fi32 := 0; fi32 < 18; fi32++ {
            if len(v2) > 0 {
                v2[0] = 42
            }
            _ = fi32
        }
    default:
        v1 = (v1 | v1)
    }
    fmt.Println(v1)
    switch 470 {
    case 2147483647:
        var v33 float32 = ((1e10 * 0.1) / (1.5 - 1e10))
        _ = v33
    case -2147483648:
        for fi34 := 0; fi34 < 20; fi34++ {
            for fi35 := 0; fi35 < 15; fi35++ {
                v36 := GenS5{F0: 42}
                var v37 int = (fi34 & (fi35 % (((-(fi34))) | 1)))
                _ = fi35
                _ = v36
                _ = v37
            }
            _ = fi34
        }
    case 127:
        if false {
            v38 := GenS4{F0: 0, F1: 65535, F2: 394}
            v2 = append(v2, 2147483647)
            if ("" > "foo bar") {
                switch 201 {
                case 42:
                    for fi39 := 0; fi39 < 1; fi39++ {
                        switch -1 {
                        case 299:
                            v40 := GenS7{F0: 42}
                            _ = v40
                        case 862:
                            var v41 uint32 = 110
                            _ = v41
                        }
                        _ = fi39
                    }
                case -407:
                    v42 := GenS8{F0: 18446744073709551615}
                    _ = v42
                }
                v1 += (v1 + (^(v1)))
                for fi43 := 0; fi43 < 7; fi43++ {
                    if ((v1 * v1) >= v1) {
                        fmt.Println(fi43)
                        if len(v2) > 0 {
                            v2[0] = fi43
                        }
                        v1 = v1
                    }
                    _ = fi43
                }
            } else {
                switch "a" {
                case "foo bar":
                    for fi44 := 0; fi44 < 15; fi44++ {
                        for fi45 := 0; fi45 < 16; fi45++ {
                            var v46 uint32 = 2
                            if ((v46 ^ v46) != 2) {
                                if (false || false) {
                                    fmt.Println(v1, v1, v46)
                                    var v47 int8 = 0
                                    _ = v47
                                } else {
                                    if (!(42 == 178)) {
                                        var v48 float32 = ((0.1 + (-2.25 * float32((-2.25 * 0.0)))) - (1e10 - 0.0))
                                        if true {
                                            switch (("foo bar" + ("\n\t" + "a")) + ("\n\t" + ("foo bar" + "foo bar"))) {
                                            case "hello":
                                                v49 := GenS3{F0: -892}
                                                if false {
                                                    var v50 int64 = 9223372036854775807
                                                    var v51 uint32 = 620
                                                    _ = v50
                                                    _ = v51
                                                } else {
                                                    var v52 int64 = 42
                                                    _ = v52
                                                }
                                                fmt.Println(fi44)
                                                _ = v49
                                            case "\n\t":
                                                fmt.Println(v1)
                                            }
                                        } else {
                                            fmt.Println(v1, v46)
                                        }
                                        _ = v48
                                    } else {
                                        var v53 int8 = 127
                                        fmt.Println(fi44, fi44)
                                        _ = v53
                                    }
                                }
                                v1 = v1
                                if ((0 == 1) || true) {
                                    v2 = append(v2, fi44)
                                    v1 -= 32767
                                }
                            }
                            v54 := GenS1{F0: 1, F1: (1 >= 0), F2: 32767}
                            _ = fi45
                            _ = v46
                            _ = v54
                        }
                        var v55 uint64 = 42
                        _ = fi44
                        _ = v55
                    }
                case "a":
                    var v56 uint8 = 178
                    _ = v56
                }
                if true {
                    v1 += 1
                    v1 += (v1 ^ (^(v1)))
                } else {
                    switch 2147483647 {
                    case -9223372036854775808:
                        var v57 uint = 65535
                        _ = v57
                    case 127:
                        var v58 int = 295
                        _ = v58
                    }
                    var v59 int16 = -32768
                    _ = v59
                }
                v60 := GenS9{F0: 42, F1: -1, F2: -1}
                _ = v60
            }
            _ = v38
        }
    }
    if len(v2) > 0 {
        v2[0] = 923
    }
    fmt.Println(v1)
    v3[2147483647] = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v3
}
