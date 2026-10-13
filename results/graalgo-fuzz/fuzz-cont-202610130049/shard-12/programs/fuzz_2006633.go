package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 uint8 }
type GenS2 struct { F0 string; F1 int16; F2 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16; F1 int8; F2 int32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int32; F1 uint16 }
type GenS8 struct { F0 uint16 }
func genFn0(p0 int) uint {
    switch p0 {
    case 42:
        p0 -= p0
    }
    v4 := GenS1{F0: (0 <= 549), F1: (0.1 == float64(-2.25)), F2: 76}
    _ = p0
    _ = v4
    return 65535
}
func genFn1(p0 int32) int32 {
    var v5 int16 = 984
    fmt.Println(p0)
    switch (("\n\t" + "foo bar") + "") {
    case "hello":
        v6 := GenS1{F0: (true || true), F1: true, F2: 255}
        _ = v6
    case "foo bar":
        p0 = p0
    case "\n\t":
        switch 127 {
        case 42:
            v5 = -263
        case -484:
            var v7 uint8 = 128
            _ = v7
        case 489:
            var v8 bool = ((0 == 255) && ((42 != 18446744073709551615) && (!(false && (true || true)))))
            _ = v8
        }
    default:
        for fi9 := 0; fi9 < 13; fi9++ {
            var v10 uint = 4294967295
            v10 -= v10
            if true {
                fmt.Println(fi9, v10)
                if true {
                    v11 := GenS2{F0: "a", F1: -32768, F2: (false && true)}
                    var v12 int64 = -1
                    v5 += (v5 >> (((v10 | (v10 / ((v10) | 1)))) & 63))
                    _ = v11
                    _ = v12
                } else {
                    var v13 int8 = 0
                    v14 := GenS2{F0: "foo bar", F1: v5, F2: false}
                    if (-904 != 0) {
                        for fi15 := 0; fi15 < 1; fi15++ {
                            fmt.Println(v13, p0, fi9)
                            _ = fi15
                        }
                    }
                    _ = v13
                    _ = v14
                }
                for fi16 := 0; fi16 < 17; fi16++ {
                    if (!(-32768 >= v5)) {
                        v10 = 798
                        var v17 int = fi16
                        _ = v17
                    }
                    fmt.Println(fi16, v5, v5)
                    v18 := GenS3{F0: v10}
                    v19 := []int{-681}
                    _ = fi16
                    _ = v18
                    _ = v19
                }
            } else {
                var v20 int16 = (v5 << ((genFn0(32767)) & 63))
                v21 := GenS4{F0: uint16(v10), F1: -20, F2: (p0 ^ p0)}
                _ = v20
                _ = v21
            }
            _ = fi9
            _ = v10
        }
    }
    for fi22 := 0; fi22 < 15; fi22++ {
        v23 := []int{-2147483648}
        var v24 uint = 18446744073709551615
        _ = fi22
        _ = v23
        _ = v24
    }
    _ = p0
    _ = v5
    return -359
}

func main() {
    var v1 uint32 = 65535
    var v2 float32 = 0.0
    var v3 float32 = 3.14159
    for fi25 := 0; fi25 < 13; fi25++ {
        for fi26 := 0; fi26 < 16; fi26++ {
            v2 = 1e10
            _ = fi26
        }
        if ((-566 != -200) && false) {
            v3 += v3
        } else {
            var v27 int = fi25
            for fi28 := 0; fi28 < 18; fi28++ {
                v29 := []int{(v27 ^ -867), 307, fi28}
                v1 += 65535
                switch (fi25 << ((0) & 63)) {
                case 333:
                    v1 = (^(v1))
                case 32767:
                    if len(v29) >= 2 {
                        v29 = v29[1:len(v29)]
                    }
                }
                _ = fi28
                _ = v29
            }
            v2 = v2
            _ = v27
        }
        v1 = (v1 + v1)
        _ = fi25
    }
    v30 := GenS4{F0: 88, F1: 1, F2: genFn1(2147483647)}
    switch 590 {
    case 9223372036854775807:
        v2 = v2
    case 0:
        fmt.Println(v2, v1)
    case -772:
        v31 := []int{32767}
        _ = v31
    }
    v32 := []int{4294967295, -108, 9223372036854775807, 607}
    var v33 uint32 = v1
    if ((!true) || (false || (((!false) || (1.5 != -2.25)) || false))) {
        if (287 != 0) {
            fmt.Println(v1, v2)
            v33 -= v33
            if (65535 == 740) {
                fmt.Println(v33)
                var v34 int8 = 1
                _ = v34
            } else {
                var v35 uint = 4294967295
                v36 := GenS5{F0: (v35 | (v35 % (((v35 - genFn0(128))) | 1)))}
                _ = v35
                _ = v36
            }
        }
        fmt.Println(v33, v1, v1)
    } else {
        if (690 <= 5) {
            switch -9223372036854775808 {
            case 4294967295:
                switch 32768 {
                case 4294967295:
                    v2 += 1.5
                case -2147483648:
                    for fi37 := 0; fi37 < 15; fi37++ {
                        v38 := GenS6{F0: (v33 + (^(v33)))}
                        fmt.Println(v3)
                        v2 = 0.1
                        _ = fi37
                        _ = v38
                    }
                default:
                    if (-2.25 == (((0.0 + 0.0) / 1.0) + 1.5)) {
                        v32 = append(v32, 32767)
                    } else {
                        var v39 uint8 = 1
                        v3 = 1.5
                        _ = v39
                    }
                }
            case 241:
                switch (("a" + "\n\t") + "\n\t") {
                case "hello":
                    var v40 uint = 255
                    _ = v40
                case "":
                    fmt.Println(v2, v3, v1)
                case "a":
                    v1 += (v33 & 1)
                }
            case 42:
                for fi41 := 0; fi41 < 14; fi41++ {
                    if (false && false) {
                        for fi42 := 0; fi42 < 10; fi42++ {
                            var v43 uint16 = 32767
                            _ = fi42
                            _ = v43
                        }
                        v44 := GenS7{F0: 0, F1: 387}
                        _ = v44
                    }
                    _ = fi41
                }
            default:
                if ((v33 << ((0) & 63)) >= (v33 << ((255) & 63))) {
                    v1 -= v1
                } else {
                    var v45 int8 = 42
                    var v46 int = 42
                    v47 := GenS1{F0: (127 < 0), F1: ((v45 - v45) > (v45 * (v45 + -39))), F2: 42}
                    _ = v45
                    _ = v46
                    _ = v47
                }
            }
            var v48 int32 = 104
            _ = len(v32)
            _ = v48
        } else {
            fmt.Println(v33, v33, v3)
        }
    }
    for fi49 := 0; fi49 < 2; fi49++ {
        var v50 int64 = 1
        _ = fi49
        _ = v50
    }
    v3 = ((v2 + 0.0) * 3.14159)
    switch -509 {
    case 127:
        var v51 float32 = -2.25
        _ = v51
    case -2147483648:
        for fi52 := 0; fi52 < 9; fi52++ {
            if false {
                if false {
                    var v53 float32 = 0.0
                    if (!("foo bar" == "foo bar")) {
                        _ = len(v32)
                    } else {
                        for fi54 := 0; fi54 < 1; fi54++ {
                            if true {
                                v2 = (v2 * v2)
                                v3 = v2
                            } else {
                                v55 := GenS6{F0: v33}
                                if (502 != 0) {
                                    v56 := GenS8{F0: 144}
                                    _ = v56
                                }
                                _ = v55
                            }
                            _ = fi54
                        }
                    }
                    _ = v53
                } else {
                    for fi57 := 0; fi57 < 10; fi57++ {
                        if false {
                            var v58 int64 = -40
                            v32 = append(v32, fi57)
                            var v59 float64 = 3.14159
                            var v60 int32 = 42
                            _ = v58
                            _ = v59
                            _ = v60
                        }
                        switch "\n\t" {
                        case "":
                            var v61 float32 = (v3 / ((v3 * 1e10) / (v2 * 1.5)))
                            if len(v32) >= 2 {
                                v32 = v32[1:len(v32)]
                            }
                            v32 = append(v32, -2147483648)
                            _ = v61
                        case "hello":
                            if (!true) {
                                v2 -= float32(((float64(v2) - (1e10 + (1.5 / 1.0))) / 1.0))
                            }
                        }
                        _ = fi57
                    }
                    for fi62 := 0; fi62 < 2; fi62++ {
                        var v63 string = ("foo bar" + "a")
                        fmt.Println(v33, v3)
                        _ = fi62
                        _ = v63
                    }
                    fmt.Println(v1, fi52, v33)
                }
                _ = len(v32)
            }
            _ = fi52
        }
    case -1:
        fmt.Println(v3, v33)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v32
    _ = v33
}
