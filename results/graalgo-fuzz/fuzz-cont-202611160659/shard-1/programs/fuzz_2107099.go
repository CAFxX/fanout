package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 uint64; F1 int8 }
func genFn0(p0 int64, p1 uint16, p2 int8) int {
    switch -490 {
    case 329:
        v5 := GenS1{F0: 42, F1: -6, F2: 911}
        _ = v5
    case 1:
        if true {
            switch 42 {
            case -48:
                for fi6 := 0; fi6 < 7; fi6++ {
                    var v7 uint8 = 127
                    _ = fi6
                    _ = v7
                }
            }
        } else {
            for fi8 := 0; fi8 < 14; fi8++ {
                var v9 int = 32767
                v10 := GenS2{F0: p0}
                _ = fi8
                _ = v10
                _ = v9
            }
            for fi11 := 0; fi11 < 12; fi11++ {
                var v12 uint64 = 4294967295
                var v13 float32 = 3.14159
                for fi14 := 0; fi14 < 8; fi14++ {
                    fmt.Println(v13, p2)
                    if false {
                        var v15 float32 = v13
                        v12 -= v12
                        _ = v15
                    } else {
                        var v16 string = (("" + "a") + "")
                        fmt.Println(p0, v16, p2)
                        switch ((("hello" + "a") + v16) + v16) {
                        case "\n\t":
                            v17 := GenS1{F0: uint8(v12), F1: p2, F2: 540}
                            _ = v17
                        case "a":
                            p2 += p2
                        case "hello":
                            if true {
                                if (!(!((true || true) || true))) {
                                    v18 := GenS3{F0: ("\n\t" + v16), F1: (fi11 * (fi11 + fi11))}
                                    switch fi11 {
                                    case -9223372036854775808:
                                        v19 := GenS1{F0: 61, F1: p2, F2: -32768}
                                        _ = v19
                                    case 60:
                                        var v20 uint64 = v12
                                        _ = v20
                                    default:
                                        var v21 int64 = (p0 % ((-427) | 1))
                                        v22 := GenS4{F0: 1, F1: (^(p2))}
                                        _ = v21
                                        _ = v22
                                    }
                                    var v23 uint8 = 0
                                    _ = v18
                                    _ = v23
                                }
                                switch ((v16 + ("" + v16)) + "hello") {
                                case "foo bar":
                                    var v24 bool = true
                                    _ = v24
                                }
                                v13 = (v13 - ((((3.14159 / v13) + float32(1e10)) + -2.25) + 3.14159))
                            } else {
                                var v25 int16 = 0
                                fmt.Println(v13)
                                _ = v25
                            }
                        }
                        _ = v16
                    }
                    _ = fi14
                }
                _ = fi11
                _ = v12
                _ = v13
            }
        }
    }
    var v26 int = 2147483647
    var v27 uint = 393
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    _ = v27
    return (v26 << ((0) & 63))
}

func main() {
    v1 := []int{42, 1, 1, -743}
    var v2 int64 = 9223372036854775807
    var v3 uint8 = 128
    var v4 bool = (true && false)
    for fi28 := 0; fi28 < 20; fi28++ {
        for fi29 := 0; fi29 < 6; fi29++ {
            if true {
                fmt.Println(v3, fi29, v2)
            } else {
                if (18446744073709551615 > 18446744073709551615) {
                    var v30 int32 = 1
                    var v31 int32 = v30
                    _ = v30
                    _ = v31
                }
                switch fi29 {
                case 1:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case 32767:
                    fmt.Println(fi28, fi28)
                }
                var v32 int = 127
                _ = v32
            }
            fmt.Println(v2, fi28, v2)
            _ = fi29
        }
        _ = fi28
    }
    _ = len(v1)
    _ = len(v1)
    fmt.Println(v2, v3)
    v33 := GenS2{F0: v2}
    v2 = (v2 & (v2 & 716))
    for fi34 := 0; fi34 < 1; fi34++ {
        v35 := GenS2{F0: v2}
        var v36 float32 = -2.25
        switch "" {
        case "foo bar":
            v4 = v4
        case "\n\t":
            for fi37 := 0; fi37 < 14; fi37++ {
                v2 = v2
                _ = fi37
            }
        case "hello":
            var v38 float64 = 3.14159
            _ = v38
        default:
            v36 = 1.5
        }
        _ = fi34
        _ = v35
        _ = v36
    }
    if ((-1 == 1) || false) {
        v4 = (genFn0(0, 32767, -1) < 9223372036854775807)
        var v39 uint = uint(v3)
        _ = v39
    } else {
        if (v3 != v3) {
            v40 := GenS1{F0: v3, F1: -128, F2: 32767}
            v41 := GenS2{F0: (v2 | -639)}
            _ = v40
            _ = v41
        }
    }
    for fi42 := 0; fi42 < 13; fi42++ {
        fmt.Println(v4, v3)
        v4 = (1 != 765)
        _ = fi42
    }
    var v43 int64 = v2
    v44 := GenS1{F0: (v3 / ((255) | 1)), F1: -128, F2: 32767}
    v4 = v4
    for fi45 := 0; fi45 < 3; fi45++ {
        var v46 uint32 = 42
        _ = fi45
        _ = v46
    }
    v1 = append(v1, 1)
    v3 -= v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v43
    _ = v44
}
