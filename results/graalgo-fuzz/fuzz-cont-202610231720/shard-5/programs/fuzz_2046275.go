package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 int64; F1 int64; F2 int64 }
func genFn0(p0 int) bool {
    fmt.Println(p0)
    var v7 uint32 = 1
    _ = p0
    _ = v7
    return true
}
func genFn1() uint32 {
    if false {
        v8 := []int{-338, 744, -842, 32768}
        if true {
            switch 32768 {
            case -723:
                var v9 int32 = 1
                _ = v9
            case -840:
                var v10 float32 = float32((float64((0.1 * -2.25)) / 1.0))
                _ = v10
            }
            var v11 uint16 = 499
            _ = v11
        }
        _ = v8
    }
    var v12 uint16 = 255
    var v13 uint64 = 4294967295
    var v14 int16 = 32767
    v13 = v13
    _ = v12
    _ = v13
    _ = v14
    return 0
}

func main() {
    var v1 int = -1
    var v2 uint64 = 4294967295
    v3 := make(map[int]bool)
    var v4 float32 = 0.1
    v5 := make(map[uint32]uint32)
    v6 := []string{"a", "\n\t", "hello", ((("hello" + ("a" + (("" + ((("" + "") + "foo bar") + "a")) + "\n\t"))) + ("\n\t" + "hello")) + "foo bar")}
    v2 = (v2 ^ v2)
    var v15 uint = 217
    var v16 int8 = 42
    var v17 uint64 = v2
    if true {
        v16 = v16
        if true {
            if (true && (true && false)) {
                var v18 uint64 = v2
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = v18
            }
            v17 = (v2 & v2)
            var v19 int16 = -32768
            _ = v19
        }
        if ((v15 + v15) < v15) {
            fmt.Println(v17, v16, v15)
            v1 = v1
            v20 := GenS1{F0: ("\n\t" + "a"), F1: v15, F2: 0}
            _ = v20
        }
    } else {
        delete(v5, genFn1())
        if (false || (false || true)) {
            if (2 == 127) {
                v21 := GenS2{F0: "\n\t"}
                _ = v21
            } else {
                var v22 bool = (float64(1.5) == (1e10 + (1.5 + 0.0)))
                var v23 bool = (v4 >= -2.25)
                _ = v22
                _ = v23
            }
            v16 -= (v16 * (v16 ^ 100))
            v24 := v3[v1]
            _ = v24
        } else {
            if len(v6) > 0 {
                v6[0] = "hello"
            }
            for fi25 := 0; fi25 < 2; fi25++ {
                v26 := GenS2{F0: ""}
                v16 = v16
                _ = fi25
                _ = v26
            }
            for fi27 := 0; fi27 < 10; fi27++ {
                switch "foo bar" {
                case "hello":
                    var v28 uint8 = 2
                    _ = v28
                case "":
                    for fi29 := 0; fi29 < 15; fi29++ {
                        _ = len(v3)
                        var v30 int32 = 950
                        var v31 string = ("a" + (("hello" + "hello") + ""))
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                case "\n\t":
                    for fi32 := 0; fi32 < 1; fi32++ {
                        v17 += 0
                        _ = fi32
                    }
                default:
                    for fi33 := 0; fi33 < 8; fi33++ {
                        v17 -= 529
                        _ = fi33
                    }
                }
                _ = len(v5)
                if (true || (false && ((float64(v4) < 1e10) && (!false)))) {
                    v16 = v16
                } else {
                    for fi34 := 0; fi34 < 19; fi34++ {
                        v16 = (v16 * (v16 ^ v16))
                        switch "" {
                        case "foo bar":
                            v2 += v17
                        default:
                            switch "a" {
                            case "a":
                                var v35 int16 = -32768
                                _ = v35
                            case "foo bar":
                                if false {
                                    switch "a" {
                                    case "hello":
                                        fmt.Println(v16)
                                    case "a":
                                        fmt.Println(v1, v1)
                                    }
                                }
                            }
                        }
                        _ = fi34
                    }
                    switch ("a" + "a") {
                    case "a":
                        if (0 == (v1 + v1)) {
                            v36 := GenS2{F0: (("foo bar" + "hello") + "hello")}
                            for fi37 := 0; fi37 < 20; fi37++ {
                                switch fi37 {
                                case -279:
                                    if true {
                                        switch v1 {
                                        case 758:
                                            fmt.Println(v2)
                                        default:
                                            var v38 uint16 = 42
                                            _ = v38
                                        }
                                        v17 += (v17 >> (((^(v15))) & 63))
                                        var v39 int8 = (^(v16))
                                        _ = v39
                                    }
                                case -174:
                                    switch "a" {
                                    case "a":
                                        switch (fi27 + (v1 - (v1 - (fi27 - (-(fi37)))))) {
                                        case 32767:
                                            v16 -= 0
                                        }
                                    }
                                    if (v15 > v15) {
                                        if false {
                                            var v40 int32 = 42
                                            _ = v40
                                        } else {
                                            v2 += (v2 << ((18446744073709551615) & 63))
                                        }
                                    } else {
                                        if ((v16 < 105) && false) {
                                            switch (v1 >> ((v15) & 63)) {
                                            case -2147483648:
                                                v41 := GenS1{F0: "a", F1: v15, F2: 25}
                                                _ = v41
                                            case 42:
                                                if len(v6) > 0 {
                                                    v6[0] = "\n\t"
                                                }
                                            case 0:
                                                fmt.Println(fi37)
                                            default:
                                                v42 := GenS3{F0: 79, F1: 838}
                                                _ = v42
                                            }
                                            v1 -= v1
                                        } else {
                                            v43 := GenS4{F0: -9223372036854775808, F1: -9223372036854775808, F2: 1}
                                            fmt.Println(v15)
                                            fmt.Println(v2, v16, v1)
                                            _ = v43
                                        }
                                    }
                                    v4 = (float32(1.5) / ((v4 / (v4 - 1.5)) - 1.5))
                                    switch (fi27 << ((v15) & 63)) {
                                    case -1:
                                        delete(v3, fi27)
                                    case 121:
                                        if true {
                                            if (false || true) {
                                                if (9223372036854775807 > 1) {
                                                    fmt.Println(v15, fi27, fi37)
                                                }
                                            }
                                            switch (("a" + "\n\t") + "") {
                                            case "":
                                                var v44 int32 = -1
                                                _ = v44
                                            }
                                        } else {
                                            v45 := v5[2147483648]
                                            if len(v6) > 0 {
                                                v6[0] = "\n\t"
                                            }
                                            fmt.Println(v4)
                                            _ = v45
                                        }
                                        fmt.Println(v4, v2)
                                    case 128:
                                        var v46 float64 = 1.5
                                        _ = v46
                                    default:
                                        var v47 uint16 = 42
                                        _ = v47
                                    }
                                default:
                                    var v48 uint32 = 0
                                    v2 += v17
                                    fmt.Println(v2)
                                    _ = v48
                                }
                                _ = fi37
                            }
                            _ = v36
                        }
                    }
                }
                _ = fi27
            }
        }
        v17 = (v17 | (^(v2)))
    }
    v17 = (v17 % ((862) | 1))
    if false {
        v17 += (^(v2))
        var v49 int16 = 397
        _ = v49
    }
    var v50 bool = (873 > 790)
    v51 := GenS1{F0: "\n\t", F1: uint(v2), F2: -631}
    for fi52 := 0; fi52 < 15; fi52++ {
        v53 := GenS2{F0: "foo bar"}
        if false {
            for fi54 := 0; fi54 < 5; fi54++ {
                if (42 == -9223372036854775808) {
                    var v55 float64 = 1.5
                    v16 = v16
                    _ = v55
                } else {
                    var v56 uint8 = 211
                    _ = len(v6)
                    v2 = 551
                    _ = v56
                }
                v1 = fi54
                _ = fi54
            }
            for fi57 := 0; fi57 < 4; fi57++ {
                var v58 float32 = v4
                v59 := GenS3{F0: 224, F1: 255}
                _ = len(v6)
                _ = fi57
                _ = v58
                _ = v59
            }
        } else {
            v4 += v4
            var v60 int16 = 1000
            var v61 uint32 = 42
            _ = v60
            _ = v61
        }
        _ = fi52
        _ = v53
    }
    v16 = v16
    v15 += v15
    fmt.Println(v16, v16)
    var v62 uint16 = 1
    var v63 uint = 786
    fmt.Println(v1)
    v64 := GenS4{F0: 1, F1: 1, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v6
    _ = v62
    _ = v63
    _ = v64
}
