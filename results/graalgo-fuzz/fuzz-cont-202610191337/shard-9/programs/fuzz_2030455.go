package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16; F1 uint16; F2 int16 }
type GenS5 struct { F0 uint16; F1 int16; F2 bool }
type GenS6 struct { F0 int8; F1 int64; F2 string }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int; F1 bool; F2 int64 }

func main() {
    var v1 uint = 315
    var v2 int8 = -1
    var v3 int64 = 937
    fmt.Println(v3, v2, v1)
    v3 = 188
    for fi4 := 0; fi4 < 16; fi4++ {
        v3 = (v3 ^ 42)
        var v5 string = ((("" + "a") + "foo bar") + "hello")
        _ = fi4
        _ = v5
    }
    v6 := []int{}
    fmt.Println(v2, v3)
    var v7 uint = v1
    v1 += (v7 ^ 0)
    var v8 bool = false
    var v9 int16 = 766
    if (42 >= 99) {
        if true {
            if v8 {
                switch ("" + "hello") {
                case "hello":
                    for fi10 := 0; fi10 < 10; fi10++ {
                        fmt.Println(v8, v2)
                        for fi11 := 0; fi11 < 13; fi11++ {
                            if v8 {
                                v2 = (-(v2))
                                switch (-(fi11)) {
                                case -2147483648:
                                    switch ("" + (("hello" + "a") + ("" + "a"))) {
                                    case "\n\t":
                                        switch (fi10 + (fi11 % (((^(fi10))) | 1))) {
                                        case 4294967295:
                                            v2 = (v2 | -85)
                                        default:
                                            v7 -= v1
                                        }
                                    case "a":
                                        if ((v9 % (((-(v9))) | 1)) >= -32768) {
                                            if ((fi10 >> ((563) & 63)) < fi11) {
                                                switch ("" + "foo bar") {
                                                case "foo bar":
                                                    switch ("" + "foo bar") {
                                                    case "a":
                                                        fmt.Println(v1)
                                                    case "\n\t":
                                                        fmt.Println(v9)
                                                    case "foo bar":
                                                        var v12 uint32 = 138
                                                        _ = v12
                                                    default:
                                                        if len(v6) > 0 {
                                                            v6[0] = (fi10 - (^(fi11)))
                                                        }
                                                    }
                                                case "a":
                                                    var v13 uint64 = 127
                                                    _ = v13
                                                default:
                                                    v14 := GenS1{F0: "foo bar", F1: fi11}
                                                    if len(v6) > 0 {
                                                        v6[0] = 2147483647
                                                    }
                                                    _ = v14
                                                }
                                                switch "" {
                                                case "":
                                                    v1 = v7
                                                case "\n\t":
                                                    v9 -= v9
                                                case "a":
                                                    var v15 float32 = 3.14159
                                                    _ = v15
                                                }
                                                v7 -= 2
                                            }
                                            var v16 uint64 = 738
                                            switch 2147483647 {
                                            case -2147483648:
                                                switch -328 {
                                                case -2147483648:
                                                    fmt.Println(v1, fi10)
                                                case 127:
                                                    v17 := GenS2{F0: -1}
                                                    _ = v17
                                                case -396:
                                                    v18 := GenS2{F0: v3}
                                                    _ = v18
                                                default:
                                                    v16 = v16
                                                }
                                                if len(v6) > 0 {
                                                    v6[0] = int(v1)
                                                }
                                            case 128:
                                                v19 := GenS2{F0: (v3 | int64(v2))}
                                                _ = v19
                                            default:
                                                var v20 int16 = -492
                                                _ = v20
                                            }
                                            v3 += (v3 / ((v3) | 1))
                                            _ = v16
                                        }
                                    case "foo bar":
                                        v9 = 125
                                    default:
                                        var v21 bool = (v8 && v8)
                                        _ = v21
                                    }
                                }
                                v3 = (v3 + 9223372036854775807)
                                v22 := GenS2{F0: (v3 / (((^(v3))) | 1))}
                                fmt.Println(v2, v9, v7)
                                _ = v22
                            }
                            var v23 int32 = 0
                            v24 := GenS3{F0: "a"}
                            _ = fi11
                            _ = v23
                            _ = v24
                        }
                        fmt.Println(fi10, v7)
                        _ = fi10
                    }
                case "a":
                    var v25 int32 = 1
                    _ = v25
                case "\n\t":
                    if (v8 && ((v8 || v8) || v8)) {
                        var v26 uint32 = 489
                        if (("" != "hello") || true) {
                            var v27 uint32 = 1
                            _ = v27
                        } else {
                            v28 := GenS3{F0: "hello"}
                            _ = v28
                        }
                        _ = v26
                    } else {
                        for fi29 := 0; fi29 < 20; fi29++ {
                            var v30 int = (fi29 & (fi29 >> (((v1 / (((^(v1))) | 1))) & 63)))
                            _ = fi29
                            _ = v30
                        }
                    }
                default:
                    _ = len(v6)
                }
                v7 = 0
            } else {
                v2 += (v2 - (v2 | (v2 | (v2 | (v2 + -76)))))
                v31 := GenS4{F0: 611, F1: 65535, F2: (v9 / ((v9) | 1))}
                for fi32 := 0; fi32 < 11; fi32++ {
                    switch "foo bar" {
                    case "hello":
                        for fi33 := 0; fi33 < 16; fi33++ {
                            switch "" {
                            case "":
                                v6 = append(v6, (fi33 % ((fi33) | 1)))
                            default:
                                if len(v6) > 0 {
                                    v6[0] = (fi32 % ((904) | 1))
                                }
                            }
                            v34 := GenS5{F0: 935, F1: v9, F2: (!false)}
                            _ = fi33
                            _ = v34
                        }
                    default:
                        v35 := GenS1{F0: "hello", F1: fi32}
                        _ = v35
                    }
                    var v36 int64 = (v3 - (v3 + -76))
                    _ = fi32
                    _ = v36
                }
                _ = v31
            }
            if (!true) {
                v37 := GenS6{F0: (v2 - v2), F1: v3, F2: "foo bar"}
                fmt.Println(v3)
                var v38 uint16 = 0
                _ = v37
                _ = v38
            } else {
                for fi39 := 0; fi39 < 1; fi39++ {
                    switch "" {
                    case "a":
                        v40 := GenS7{F0: 806}
                        _ = v40
                    }
                    if ((1.5 + 1.5) <= 1.5) {
                        var v41 uint64 = 0
                        v1 += v1
                        var v42 bool = ((v8 || true) || v8)
                        _ = v41
                        _ = v42
                    }
                    for fi43 := 0; fi43 < 18; fi43++ {
                        switch (-(fi43)) {
                        case 32768:
                            v1 = v1
                        case 4294967295:
                            v2 += 100
                        }
                        var v44 string = ("\n\t" + "hello")
                        v7 -= v7
                        _ = fi43
                        _ = v44
                    }
                    _ = fi39
                }
                if true {
                    for fi45 := 0; fi45 < 8; fi45++ {
                        v8 = v8
                        if len(v6) > 0 {
                            v6[0] = (fi45 & int(v3))
                        }
                        _ = fi45
                    }
                }
                var v46 uint64 = 1
                _ = v46
            }
            var v47 float64 = 1e10
            _ = v47
        }
        switch -704 {
        case 2147483647:
            var v48 uint32 = 2147483648
            _ = v48
        case 32767:
            var v49 uint16 = 0
            _ = v49
        case 42:
            v9 = -1
        }
    } else {
        switch "hello" {
        case "a":
            _ = len(v6)
        case "hello":
            for fi50 := 0; fi50 < 13; fi50++ {
                v1 += v1
                v8 = true
                fmt.Println(v8, v1)
                _ = fi50
            }
        case "":
            for fi51 := 0; fi51 < 2; fi51++ {
                var v52 uint8 = 127
                for fi53 := 0; fi53 < 10; fi53++ {
                    v2 -= 42
                    var v54 uint32 = 677
                    if (v54 > 42) {
                        v55 := GenS8{F0: (fi51 * fi53), F1: true, F2: -9223372036854775808}
                        var v56 string = "a"
                        _ = v55
                        _ = v56
                    } else {
                        v52 += (^(v52))
                        _ = len(v6)
                    }
                    _ = fi53
                    _ = v54
                }
                _ = fi51
                _ = v52
            }
        }
        var v57 uint64 = 517
        var v58 int64 = (v3 << ((v1) & 63))
        _ = v57
        _ = v58
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
