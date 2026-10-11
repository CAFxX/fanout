package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64; F1 int16 }
type GenS5 struct { F0 int8; F1 uint8 }

func main() {
    v1 := []int{960, 332, 4294967295, -9223372036854775808}
    var v2 int32 = -1
    var v3 bool = true
    v4 := make(map[uint32]string)
    var v5 float32 = (1e10 / 3.14159)
    var v6 int64 = 185
    v6 += (v6 & (v6 >> ((2) & 63)))
    var v7 float64 = 1e10
    v3 = (42 < (-(v2)))
    switch "a" {
    case "":
        for fi8 := 0; fi8 < 3; fi8++ {
            for fi9 := 0; fi9 < 8; fi9++ {
                switch "a" {
                case "foo bar":
                    fmt.Println(v5)
                }
                var v10 int = 32767
                var v11 int64 = (^(v6))
                v12 := GenS1{F0: 33, F1: 18446744073709551615}
                if false {
                    v13 := GenS1{F0: 42, F1: 885}
                    fmt.Println(fi9)
                    switch (fi8 / (((-(fi8))) | 1)) {
                    case 1:
                        fmt.Println(v6, fi9)
                    }
                    _ = v13
                } else {
                    if (!true) {
                        delete(v4, 65535)
                        var v14 uint = 0
                        _ = v14
                    } else {
                        switch "hello" {
                        case "a":
                            fmt.Println(v6)
                        case "foo bar":
                            v15 := GenS2{F0: ("\n\t" + (("\n\t" + "") + "hello")), F1: (fi8 % ((234) | 1))}
                            _ = v15
                        case "\n\t":
                            var v16 uint32 = 390
                            _ = v16
                        }
                        v10 -= (fi9 & (fi9 - int(v2)))
                        if (2147483648 == 2147483647) {
                            v4[2] = "foo bar"
                            v1 = append(v1, 1)
                        } else {
                            v7 -= -2.25
                            var v17 uint = 400
                            _ = v17
                        }
                    }
                }
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
            if len(v1) > 0 {
                v1[0] = -1
            }
            var v18 uint16 = 42
            _ = fi8
            _ = v18
        }
    case "a":
        fmt.Println(v2, v5, v7)
    case "\n\t":
        var v19 int = 128
        _ = v19
    }
    if false {
        for fi20 := 0; fi20 < 18; fi20++ {
            if v3 {
                if false {
                    v7 -= v7
                    v21 := GenS3{F0: -32768}
                    _ = v21
                }
                v3 = true
            } else {
                v22 := GenS3{F0: 1}
                if len(v1) > 0 {
                    v1[0] = fi20
                }
                var v23 string = ("a" + (((("a" + ("\n\t" + ("a" + "foo bar"))) + "foo bar") + "\n\t") + "hello"))
                _ = v22
                _ = v23
            }
            _ = fi20
        }
    } else {
        var v24 int64 = v6
        _ = v24
    }
    var v25 uint = 255
    switch 129 {
    case -564:
        v26 := v4[2]
        _ = v26
    case 32767:
        for fi27 := 0; fi27 < 2; fi27++ {
            for fi28 := 0; fi28 < 16; fi28++ {
                var v29 int16 = 1
                var v30 uint64 = 18446744073709551615
                if (fi27 > (fi28 & fi28)) {
                    switch "foo bar" {
                    case "hello":
                        var v31 float64 = 1.5
                        _ = v31
                    case "foo bar":
                        if (453 < 558) {
                            _ = len(v4)
                            v4[2147483648] = "a"
                        } else {
                            var v32 float64 = (((1.5 - (0.1 + float64((v5 * 1.5)))) * v7) / v7)
                            v6 -= v6
                            _ = v32
                        }
                    case "":
                        if (!((719 < 613) || v3)) {
                            if (127 <= 123) {
                                v30 -= 1
                                if v3 {
                                    v3 = (false || false)
                                    var v33 int = -2
                                    _ = len(v4)
                                    _ = v33
                                } else {
                                    if v3 {
                                        v6 = v6
                                        switch 489 {
                                        case 566:
                                            v5 = v5
                                        case 9223372036854775807:
                                            if (!(!v3)) {
                                                v6 = (v6 - 1)
                                                var v34 int64 = (v6 << (((v25 << ((254) & 63))) & 63))
                                                _ = v34
                                            }
                                        default:
                                            v35 := GenS4{F0: (v30 - 564), F1: (-(v29))}
                                            _ = v35
                                        }
                                        switch int(v30) {
                                        case 325:
                                            v6 -= 9223372036854775807
                                        case -332:
                                            v36 := GenS1{F0: -1, F1: 2}
                                            _ = v36
                                        }
                                    }
                                    v1 = append(v1, fi28)
                                    var v37 float64 = 0.0
                                    _ = v37
                                }
                            } else {
                                var v38 bool = ((1 < 42) || true)
                                if v3 {
                                    v3 = v38
                                    v38 = false
                                    var v39 uint64 = (v30 * 25)
                                    v6 += -529
                                    _ = v39
                                }
                                _ = v38
                            }
                        }
                    }
                    var v40 int16 = (v29 % ((42) | 1))
                    v2 = v2
                    _ = v40
                }
                _ = fi28
                _ = v29
                _ = v30
            }
            v3 = v3
            _ = fi27
        }
    default:
        switch "\n\t" {
        case "hello":
            v5 = (v5 + v5)
        case "\n\t":
            for fi41 := 0; fi41 < 9; fi41++ {
                for fi42 := 0; fi42 < 11; fi42++ {
                    v43 := GenS4{F0: 42, F1: -32768}
                    var v44 bool = true
                    switch -859 {
                    case -230:
                        if ((v6 != (^(v6))) && (1.5 >= v5)) {
                            var v45 uint64 = 0
                            var v46 uint32 = 832
                            fmt.Println(v45, v2, v45)
                            _ = v45
                            _ = v46
                        }
                    case 128:
                        var v47 uint64 = 0
                        _ = v47
                    default:
                        if v3 {
                            v7 = 0.1
                        }
                    }
                    _ = fi42
                    _ = v43
                    _ = v44
                }
                var v48 bool = (1000 != 32767)
                _ = fi41
                _ = v48
            }
        case "a":
            var v49 float32 = v5
            _ = v49
        default:
            fmt.Println(v25, v25, v5)
        }
    }
    fmt.Println(v3, v2, v6)
    switch 32768 {
    case 949:
        v7 += ((-2.25 - v7) - 3.14159)
    }
    for fi50 := 0; fi50 < 1; fi50++ {
        v51 := v4[2]
        switch (fi50 | fi50) {
        case 901:
            for fi52 := 0; fi52 < 11; fi52++ {
                if v3 {
                    _ = len(v4)
                    fmt.Println(v51, fi50, v51)
                    if v3 {
                        var v53 int16 = 1
                        v4[880] = "\n\t"
                        var v54 int64 = v6
                        var v55 float32 = 3.14159
                        var v56 uint64 = 18446744073709551615
                        _ = v53
                        _ = v54
                        _ = v55
                        _ = v56
                    } else {
                        fmt.Println(v5, v7, v6)
                    }
                } else {
                    v5 = 0.1
                }
                _ = len(v4)
                var v57 string = ("\n\t" + ("hello" + ("\n\t" + ((("foo bar" + "\n\t") + "foo bar") + ((((v51 + "") + v51) + ("foo bar" + v51)) + ("\n\t" + ("" + v51)))))))
                v57 = v57
                var v58 int16 = 0
                _ = fi52
                _ = v57
                _ = v58
            }
        default:
            if len(v1) > 0 {
                v1[0] = fi50
            }
        }
        for fi59 := 0; fi59 < 11; fi59++ {
            switch (fi59 * fi50) {
            case 9223372036854775807:
                v2 = 2147483647
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v6 = (v6 << ((v25) & 63))
            _ = fi59
        }
        _ = fi50
        _ = v51
    }
    v5 = 0.1
    v5 += v5
    switch -101 {
    case 642:
        if (0 == 18446744073709551615) {
            v60 := GenS5{F0: 0, F1: 2}
            _ = v60
        } else {
            v3 = false
        }
    default:
        var v61 uint32 = 2
        _ = v61
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
