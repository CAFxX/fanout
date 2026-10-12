package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 int }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 string; F1 bool; F2 int }
type GenS4 struct { F0 int8; F1 int8; F2 string }

func main() {
    var v1 bool = (true || (0 >= 1))
    var v2 uint8 = 128
    var v3 float32 = 1.5
    for fi4 := 0; fi4 < 13; fi4++ {
        if v1 {
            switch "" {
            case "hello":
                for fi5 := 0; fi5 < 6; fi5++ {
                    v6 := []int{}
                    _ = fi5
                    _ = v6
                }
            case "\n\t":
                var v7 int8 = -98
                _ = v7
            case "foo bar":
                fmt.Println(fi4)
            }
        }
        v8 := []int{fi4, (^(fi4))}
        v9 := GenS1{F0: 9223372036854775807, F1: "\n\t", F2: fi4}
        _ = fi4
        _ = v8
        _ = v9
    }
    switch -575 {
    case 32767:
        var v10 float32 = v3
        _ = v10
    case 32768:
        v3 = 0.0
    }
    v1 = (1 >= 1)
    for fi11 := 0; fi11 < 18; fi11++ {
        if (0.0 >= (0.0 * 1.5)) {
            v1 = (779 > 4294967295)
            var v12 uint8 = 127
            _ = v12
        }
        switch ("foo bar" + "\n\t") {
        case "\n\t":
            var v13 float32 = v3
            _ = v13
        default:
            if (18446744073709551615 == 1) {
                for fi14 := 0; fi14 < 8; fi14++ {
                    fmt.Println(v1, fi14, fi11)
                    fmt.Println(fi11)
                    _ = fi14
                }
                v1 = (!v1)
                for fi15 := 0; fi15 < 14; fi15++ {
                    var v16 float32 = 0.0
                    switch "" {
                    case "\n\t":
                        v17 := GenS1{F0: -1, F1: "", F2: fi15}
                        _ = v17
                    case "":
                        switch "hello" {
                        case "hello":
                            var v18 int = 4294967295
                            _ = v18
                        }
                    default:
                        switch ("foo bar" + "") {
                        case "a":
                            var v19 uint64 = 885
                            _ = v19
                        default:
                            v20 := []int{}
                            _ = v20
                        }
                    }
                    _ = fi15
                    _ = v16
                }
            }
        }
        _ = fi11
    }
    var v21 int = -496
    var v22 int32 = 42
    var v23 int16 = 1
    for fi24 := 0; fi24 < 11; fi24++ {
        v2 -= 2
        _ = fi24
    }
    for fi25 := 0; fi25 < 15; fi25++ {
        v26 := GenS1{F0: -9223372036854775808, F1: "", F2: fi25}
        for fi27 := 0; fi27 < 12; fi27++ {
            var v28 string = ""
            v29 := GenS1{F0: int64(v23), F1: v28, F2: fi27}
            v30 := GenS1{F0: 0, F1: "a", F2: int(v23)}
            _ = fi27
            _ = v28
            _ = v29
            _ = v30
        }
        _ = fi25
        _ = v26
    }
    if (v21 < v21) {
        var v31 uint32 = 2147483648
        switch "foo bar" {
        case "foo bar":
            v22 -= (v22 << ((315) & 63))
        case "":
            if ((18446744073709551615 >= 63) && true) {
                for fi32 := 0; fi32 < 9; fi32++ {
                    v22 = -2147483648
                    v22 += (^(v22))
                    _ = fi32
                }
                v33 := GenS2{F0: true, F1: v21}
                for fi34 := 0; fi34 < 11; fi34++ {
                    switch "" {
                    case "hello":
                        for fi35 := 0; fi35 < 20; fi35++ {
                            switch 42 {
                            case -1:
                                var v36 float64 = (1.5 / 0.1)
                                _ = v36
                            case 4294967295:
                                var v37 float32 = 0.1
                                _ = v37
                            }
                            v38 := []int{}
                            _ = fi35
                            _ = v38
                        }
                    case "foo bar":
                        for fi39 := 0; fi39 < 2; fi39++ {
                            v31 = 532
                            if (v3 > 1e10) {
                                if ((v2 % ((v2) | 1)) > v2) {
                                    v1 = v1
                                }
                                v40 := []int{}
                                _ = v40
                            } else {
                                if v1 {
                                    switch v21 {
                                    case 396:
                                        v41 := []int{(fi39 & fi39), fi39, (^(v21))}
                                        _ = v41
                                    case 128:
                                        v1 = v1
                                    case 2147483647:
                                        fmt.Println(v3, v22, v22)
                                    }
                                    if false {
                                        switch "a" {
                                        case "":
                                            fmt.Println(v23, v31, v31)
                                        case "\n\t":
                                            fmt.Println(v22, v31)
                                        }
                                        v42 := GenS1{F0: -9223372036854775808, F1: ("" + (("" + ("foo bar" + ("foo bar" + ("\n\t" + "\n\t")))) + "a")), F2: (fi34 * -66)}
                                        _ = v42
                                    }
                                } else {
                                    switch fi34 {
                                    case 32767:
                                        v43 := GenS3{F0: "hello", F1: false, F2: 9223372036854775807}
                                        _ = v43
                                    case 1:
                                        v3 = (((0.1 - v3) / float32(1e10)) - v3)
                                        switch (fi34 ^ (v21 >> ((uint(v31)) & 63))) {
                                        case -48:
                                            if false {
                                                v44 := GenS2{F0: (v3 > float32(1.5)), F1: int(v31)}
                                                v3 = v3
                                                var v45 uint16 = 255
                                                _ = v44
                                                _ = v45
                                            } else {
                                                v2 -= 42
                                                var v46 int = (fi39 & 1)
                                                _ = v46
                                            }
                                        default:
                                            v3 = float32((3.14159 / 1.5))
                                        }
                                        v47 := []int{-9223372036854775808, v21}
                                        _ = v47
                                    }
                                    v23 = (v23 % ((32767) | 1))
                                    v2 = 237
                                    v48 := []int{9223372036854775807, 128, -348}
                                    _ = v48
                                }
                                var v49 int64 = -26
                                _ = v49
                            }
                            v1 = v1
                            _ = fi39
                        }
                    default:
                        var v50 bool = ((^(v31)) < v31)
                        _ = v50
                    }
                    switch (fi34 + v21) {
                    case -9223372036854775808:
                        v51 := GenS4{F0: 1, F1: -128, F2: ""}
                        _ = v51
                    }
                    _ = fi34
                }
                _ = v33
            } else {
                if v1 {
                    var v52 int16 = 0
                    _ = v52
                }
                v53 := []int{int(v23), (^(v21)), v21, v21}
                v3 = (((float32(1e10) / -2.25) * (((v3 / -2.25) / v3) / v3)) + v3)
                _ = v53
            }
        default:
            v31 = v31
        }
        fmt.Println(v3, v1, v22)
        _ = v31
    }
    var v54 int64 = -639
    v3 -= v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v54
}
