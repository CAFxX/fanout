package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 int8; F1 int; F2 bool }
type GenS4 struct { F0 string; F1 uint }
type GenS5 struct { F0 string; F1 int }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint32 }

func main() {
    v1 := []uint32{2147483648, 2147483647, 2147483648, 2}
    v2 := []string{"", ("hello" + ("a" + "foo bar")), "hello", ""}
    v3 := make(map[uint32]bool)
    var v4 uint = 877
    var v5 uint8 = 239
    var v6 int32 = 1
    v7 := GenS1{F0: -1}
    if false {
        var v8 float64 = 1.5
        for fi9 := 0; fi9 < 16; fi9++ {
            for fi10 := 0; fi10 < 13; fi10++ {
                if (v4 >= 65535) {
                    v11 := GenS1{F0: 1}
                    switch -9223372036854775808 {
                    case -525:
                        switch "foo bar" {
                        case "hello":
                            var v12 float32 = (float32(-2.25) * float32((v8 - v8)))
                            fmt.Println(v12)
                            _ = v12
                        }
                    case -2147483648:
                        var v13 int = (fi9 | (fi9 ^ 128))
                        _ = v13
                    default:
                        switch "\n\t" {
                        case "hello":
                            var v14 uint32 = 144
                            _ = v14
                        }
                    }
                    v8 -= 1.5
                    _ = v11
                }
                v15 := v3[65535]
                if (v15 || ((!v15) && v15)) {
                    var v16 uint = (v4 & v4)
                    if (!(18446744073709551615 == 374)) {
                        fmt.Println(fi10)
                        if (1.5 == -2.25) {
                            v8 = v8
                            delete(v3, 2)
                        } else {
                            if ((v6 | v6) <= -2147483648) {
                                v5 -= (^(v5))
                                switch ("hello" + ("" + "\n\t")) {
                                case "a":
                                    v4 = 1
                                }
                                v5 -= (v5 << ((v16) & 63))
                            } else {
                                v15 = false
                                v5 = 127
                                v15 = false
                                switch "\n\t" {
                                case "a":
                                    if ((v15 || v15) || true) {
                                        var v17 int8 = 127
                                        if ((uint32(v5) < 42) && ((^(v5)) == (^(v5)))) {
                                            v18 := GenS1{F0: (-(v17))}
                                            v19 := GenS2{F0: int32(v17), F1: 0}
                                            fmt.Println(v4, v15, v17)
                                            _ = v18
                                            _ = v19
                                        }
                                        _ = v17
                                    }
                                case "hello":
                                    var v20 int64 = 42
                                    _ = v20
                                }
                                var v21 float64 = v8
                                _ = v21
                            }
                            v22 := GenS1{F0: -128}
                            _ = v22
                        }
                        switch "\n\t" {
                        case "a":
                            var v23 int8 = 42
                            _ = v23
                        case "\n\t":
                            delete(v3, 2)
                        }
                    } else {
                        if true {
                            v6 = -192
                        } else {
                            v4 = (v4 * v4)
                            var v24 uint32 = 92
                            v25 := GenS2{F0: 42, F1: (v24 & (^(v24)))}
                            var v26 int32 = -2147483648
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                    }
                    if false {
                        if false {
                            v1 = append(v1, 828)
                        }
                    } else {
                        fmt.Println(v16, fi9, v15)
                    }
                    _ = v16
                } else {
                    v15 = v15
                    switch fi10 {
                    case 1:
                        v4 = (v4 << ((v4) & 63))
                    case -9223372036854775808:
                        v27 := GenS3{F0: 0, F1: fi9, F2: v15}
                        _ = v27
                    }
                    v15 = false
                }
                _ = fi10
                _ = v15
            }
            var v28 int = (fi9 * (^(fi9)))
            var v29 bool = ("" < ("\n\t" + ("" + ("foo bar" + ("hello" + "hello")))))
            _ = fi9
            _ = v28
            _ = v29
        }
        var v30 int64 = 1
        _ = v30
        _ = v8
    }
    for fi31 := 0; fi31 < 8; fi31++ {
        fmt.Println(v5, v4, fi31)
        if (true || ((^(v5)) < (v5 | v5))) {
            fmt.Println(fi31, fi31, fi31)
            v32 := GenS3{F0: 127, F1: (fi31 - 127), F2: false}
            var v33 int = fi31
            _ = v32
            _ = v33
        }
        _ = fi31
    }
    var v34 int8 = 127
    if true {
        v4 += (v4 - v4)
        v34 = v34
        switch "\n\t" {
        case "foo bar":
            v5 = v5
        case "hello":
            var v35 float32 = 1.5
            _ = v35
        default:
            v36 := GenS4{F0: "hello", F1: (v4 ^ v4)}
            _ = v36
        }
    } else {
        var v37 bool = false
        _ = v37
    }
    if (!(("\n\t" != ("hello" + "")) || true)) {
        v6 += 42
        if false {
            for fi38 := 0; fi38 < 2; fi38++ {
                var v39 bool = true
                _ = fi38
                _ = v39
            }
            var v40 uint = v4
            v3[4294967295] = ((-1 == (v34 << ((107) & 63))) || ("" < ("\n\t" + "hello")))
            _ = v40
        }
        for fi41 := 0; fi41 < 12; fi41++ {
            if (2147483647 == 1) {
                v5 += v5
                switch "a" {
                case "\n\t":
                    var v42 uint16 = 65535
                    _ = v42
                case "foo bar":
                    v3[42] = false
                }
            } else {
                v43 := GenS5{F0: "", F1: 4294967295}
                switch (("a" + "") + ("a" + "a")) {
                case "foo bar":
                    var v44 int64 = 1
                    _ = v44
                case "hello":
                    v4 += v4
                case "\n\t":
                    var v45 float64 = 1e10
                    _ = v45
                }
                _ = v43
            }
            _ = fi41
        }
    } else {
        v46 := GenS6{F0: 348}
        v5 = 1
        _ = v46
    }
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    v34 = (v34 % (((v34 >> (((v4 >> (((v4 | 4294967295)) & 63))) & 63))) | 1))
    for fi47 := 0; fi47 < 2; fi47++ {
        if false {
            fmt.Println(v4, v4)
            var v48 uint64 = 2
            _ = len(v2)
            _ = v48
        } else {
            v5 -= (^(v5))
        }
        _ = fi47
    }
    var v49 int16 = 323
    for fi50 := 0; fi50 < 14; fi50++ {
        fmt.Println(v34, v4, v6)
        for fi51 := 0; fi51 < 10; fi51++ {
            _ = len(v2)
            _ = fi51
        }
        switch "foo bar" {
        case "a":
            v4 += v4
        }
        _ = fi50
    }
    v6 += int32(v34)
    for fi52 := 0; fi52 < 7; fi52++ {
        v34 += -1
        _ = fi52
    }
    if true {
        switch "\n\t" {
        case "":
            switch "a" {
            case "\n\t":
                switch 1 {
                case 499:
                    var v53 int64 = 9223372036854775807
                    _ = v53
                default:
                    fmt.Println(v5)
                }
            }
        default:
            if (v4 < v4) {
                v34 -= v34
                v54 := GenS7{F0: 555}
                _ = v54
            }
        }
        fmt.Println(v34, v34)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v49
    _ = v5
    _ = v6
    _ = v7
}
