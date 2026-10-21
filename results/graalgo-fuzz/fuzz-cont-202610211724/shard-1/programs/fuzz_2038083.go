package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint16 }
type GenS3 struct { F0 uint; F1 int; F2 uint }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64; F1 int16; F2 uint16 }
type GenS6 struct { F0 string; F1 bool }
type GenS7 struct { F0 int16; F1 uint8 }
func genFn0() int {
    var v5 uint = 2
    fmt.Println(v5)
    v5 = (v5 << ((v5) & 63))
    v5 = (v5 + v5)
    _ = v5
    return 127
}
func genFn1() bool {
    var v6 uint32 = 2147483648
    for fi7 := 0; fi7 < 5; fi7++ {
        for fi8 := 0; fi8 < 9; fi8++ {
            v6 -= 926
            var v9 bool = (9223372036854775807 > 9223372036854775807)
            v9 = (("hello" <= ("a" + "foo bar")) && (8 > uint8(v6)))
            _ = fi8
            _ = v9
        }
        v10 := []int{}
        v11 := GenS1{F0: 127, F1: 2}
        _ = fi7
        _ = v10
        _ = v11
    }
    var v12 uint8 = 1
    v6 += (v6 << ((18446744073709551615) & 63))
    _ = v12
    _ = v6
    return false
}

func main() {
    var v1 uint32 = 4294967295
    var v2 uint8 = 1
    var v3 uint8 = v2
    var v4 float32 = float32(float64((3.14159 / 1.5)))
    v1 -= v1
    switch 9223372036854775807 {
    case -1:
        v13 := []int{736, 0}
        _ = v13
    }
    var v14 int64 = -973
    v15 := []int{0, genFn0(), 599, -1}
    v2 += 246
    v15 = append(v15, 127)
    switch "foo bar" {
    case "\n\t":
        v2 = 255
    default:
        v16 := GenS2{F0: 65535, F1: v1, F2: uint16(v3)}
        _ = v16
    }
    v17 := GenS2{F0: (^(v1)), F1: (v1 | (v1 | (v1 - v1))), F2: 255}
    if (v1 < 42) {
        if false {
            for fi18 := 0; fi18 < 11; fi18++ {
                if (!true) {
                    switch fi18 {
                    case -9223372036854775808:
                        var v19 int = 898
                        _ = v19
                    case -476:
                        var v20 uint64 = 2
                        _ = v20
                    }
                    if (true || false) {
                        v21 := GenS3{F0: 42, F1: 127, F2: 65535}
                        fmt.Println(v4, v3, v3)
                        _ = v21
                    }
                } else {
                    if (!(42 == -1)) {
                        v22 := GenS2{F0: 453, F1: 0, F2: 132}
                        if true {
                            var v23 uint64 = 131
                            var v24 string = "foo bar"
                            _ = v23
                            _ = v24
                        }
                        _ = v22
                    } else {
                        if true {
                            v25 := GenS2{F0: uint32(v2), F1: v1, F2: 696}
                            if (18446744073709551615 > 586) {
                                v14 = v14
                                switch genFn0() {
                                case -1:
                                    if len(v15) > 0 {
                                        v15[0] = (fi18 + genFn0())
                                    }
                                default:
                                    var v26 uint64 = 0
                                    _ = v26
                                }
                                v4 = float32(3.14159)
                            }
                            _ = v25
                        } else {
                            v1 = v1
                        }
                        v14 = v14
                        var v27 uint64 = 4294967295
                        _ = v27
                    }
                    var v28 bool = false
                    _ = v28
                }
                if len(v15) > 0 {
                    v15[0] = fi18
                }
                _ = fi18
            }
            if (!(32767 < 163)) {
                for fi29 := 0; fi29 < 20; fi29++ {
                    var v30 uint = 255
                    if len(v15) >= 2 {
                        v15 = v15[1:len(v15)]
                    }
                    v2 = 127
                    _ = fi29
                    _ = v30
                }
            } else {
                v31 := GenS4{F0: ""}
                var v32 int8 = 0
                var v33 uint64 = 42
                _ = v31
                _ = v32
                _ = v33
            }
        }
    } else {
        fmt.Println(v3)
        v4 = (v4 * (0.0 / v4))
    }
    v1 += (v1 + v1)
    if (!false) {
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
    }
    fmt.Println(v4, v1, v1)
    var v34 int = 128
    v35 := GenS1{F0: (v3 & v2), F1: v3}
    var v36 int16 = -32768
    switch (v34 << ((312) & 63)) {
    case -1:
        if true {
            if true {
                if false {
                    v4 = v4
                    v15 = append(v15, genFn0())
                    v2 = 128
                }
            } else {
                fmt.Println(v1, v14, v3)
                v4 = (3.14159 * ((((float32(3.14159) + 1.5) / v4) / 1.0) - v4))
                var v37 bool = ("\n\t" <= ("" + ""))
                _ = v37
            }
        }
    case 828:
        v14 = 883
    case 127:
        v15 = append(v15, 0)
    default:
        if false {
            v15 = append(v15, (^(v34)))
            switch v34 {
            case 7:
                for fi38 := 0; fi38 < 9; fi38++ {
                    for fi39 := 0; fi39 < 14; fi39++ {
                        switch ("foo bar" + "a") {
                        case "\n\t":
                            v14 -= (-(v14))
                        case "hello":
                            var v40 int8 = -128
                            _ = v40
                        }
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                        switch (-(v34)) {
                        case -2147483648:
                            switch "hello" {
                            case "":
                                if (!false) {
                                    switch "a" {
                                    case "hello":
                                        var v41 string = ("\n\t" + "a")
                                        _ = v41
                                    case "foo bar":
                                        v42 := GenS5{F0: uint64(v1), F1: (v36 / (((v36 >> ((0) & 63))) | 1)), F2: 457}
                                        _ = v42
                                    case "\n\t":
                                        _ = len(v15)
                                    default:
                                        v36 = 0
                                    }
                                } else {
                                    if (0 >= -367) {
                                        if false {
                                            if (1e10 >= 3.14159) {
                                                v43 := GenS1{F0: 198, F1: v3}
                                                v14 = (v14 | 9223372036854775807)
                                                var v44 uint64 = 639
                                                var v45 int16 = v36
                                                _ = v43
                                                _ = v44
                                                _ = v45
                                            } else {
                                                v1 -= v1
                                            }
                                            if len(v15) >= 2 {
                                                v15 = v15[1:len(v15)]
                                            }
                                            fmt.Println(v2)
                                        }
                                    }
                                    if (2 > 42) {
                                        fmt.Println(v1)
                                        switch ("" + "a") {
                                        case "foo bar":
                                            v46 := GenS3{F0: 1, F1: 136, F2: 2}
                                            _ = v46
                                        case "a":
                                            v47 := GenS6{F0: ("\n\t" + "foo bar"), F1: (1 != 0)}
                                            _ = v47
                                        }
                                        v48 := GenS5{F0: uint64(v3), F1: v36, F2: 0}
                                        _ = v48
                                    } else {
                                        v2 = (^(v3))
                                        var v49 bool = true
                                        var v50 int8 = 127
                                        _ = v49
                                        _ = v50
                                    }
                                }
                            default:
                                fmt.Println(v4, v1, v36)
                            }
                        case 2147483647:
                            fmt.Println(fi39)
                        }
                        v51 := GenS6{F0: ("foo bar" + "\n\t"), F1: false}
                        v2 += v3
                        v52 := GenS7{F0: (^(v36)), F1: (v2 << ((42) & 63))}
                        _ = fi39
                        _ = v51
                        _ = v52
                    }
                    if true {
                        var v53 string = "hello"
                        v54 := GenS2{F0: (v1 + v1), F1: 4294967295, F2: 275}
                        switch (fi38 + fi38) {
                        case 204:
                            v53 = v53
                        case -1:
                            if len(v15) >= 2 {
                                v15 = v15[1:len(v15)]
                            }
                        }
                        _ = v53
                        _ = v54
                    } else {
                        var v55 bool = ((0 > 0) || false)
                        v14 -= v14
                        _ = v55
                    }
                    _ = fi38
                }
            case -2147483648:
                switch (v34 - genFn0()) {
                case -698:
                    for fi56 := 0; fi56 < 19; fi56++ {
                        if (18446744073709551615 != 1) {
                            if (!(!true)) {
                                if len(v15) >= 2 {
                                    v15 = v15[1:len(v15)]
                                }
                                v4 = 0.1
                            }
                            if false {
                                v3 += v2
                            }
                        }
                        _ = fi56
                    }
                }
            case 0:
                fmt.Println(v4)
            default:
                fmt.Println(v14, v3, v3)
            }
            v15 = append(v15, (v34 + v34))
        }
    }
    v3 -= v2
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
