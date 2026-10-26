package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint }
type GenS2 struct { F0 int; F1 int8; F2 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string }
func genFn0() uint {
    var v4 int = 32768
    var v5 uint8 = 2
    _ = v4
    _ = v5
    return 590
}
func genFn1(p0 uint8) uint64 {
    p0 = p0
    v6 := GenS1{F0: -1, F1: -1, F2: 1}
    var v7 int = 0
    var v8 int16 = 42
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 996
}

func main() {
    v1 := []uint32{}
    var v2 bool = (-290 >= 0)
    v3 := []string{"\n\t", "\n\t", "\n\t", ("\n\t" + "a")}
    v9 := GenS1{F0: 974, F1: 2147483647, F2: 18446744073709551615}
    var v10 uint = genFn0()
    _ = len(v3)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi11 := 0; fi11 < 3; fi11++ {
        if v2 {
            var v12 int64 = 0
            _ = v12
        } else {
            for fi13 := 0; fi13 < 17; fi13++ {
                switch "\n\t" {
                case "hello":
                    v14 := GenS1{F0: -32768, F1: fi11, F2: (v10 - (v10 * genFn0()))}
                    _ = v14
                default:
                    if v2 {
                        var v15 int8 = -84
                        var v16 uint8 = 16
                        _ = v15
                        _ = v16
                    } else {
                        v10 = genFn0()
                        switch ("foo bar" + "hello") {
                        case "hello":
                            v2 = (((0.0 * -2.25) > 0.1) || v2)
                        default:
                            if v2 {
                                v17 := GenS2{F0: fi13, F1: 1, F2: uint32(v10)}
                                if (2147483647 != 1) {
                                    var v18 int16 = 42
                                    switch "foo bar" {
                                    case "":
                                        var v19 uint16 = 42
                                        _ = v19
                                    case "hello":
                                        if (!(0 == 3)) {
                                            if (!(127 <= 42)) {
                                                v20 := GenS1{F0: v18, F1: fi13, F2: 2}
                                                _ = v20
                                            }
                                            v3 = append(v3, "foo bar")
                                            v3 = append(v3, "\n\t")
                                        }
                                    }
                                    var v21 uint = v10
                                    _ = v18
                                    _ = v21
                                }
                                var v22 uint = v10
                                _ = v17
                                _ = v22
                            } else {
                                fmt.Println(v2, v10)
                                if (!false) {
                                    v23 := GenS1{F0: -68, F1: (fi13 - -468), F2: 18446744073709551615}
                                    _ = v23
                                } else {
                                    if (true && (v2 || true)) {
                                        v24 := GenS1{F0: 846, F1: (fi13 & fi13), F2: (^(v10))}
                                        if len(v1) > 0 {
                                            v1[0] = 65535
                                        }
                                        v10 += (v10 & v10)
                                        _ = v24
                                    }
                                    v2 = true
                                    fmt.Println(v10)
                                }
                            }
                        }
                    }
                }
                _ = fi13
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            switch fi11 {
            case -2147483648:
                _ = len(v1)
            case 42:
                v25 := GenS2{F0: fi11, F1: 127, F2: 2147483648}
                _ = v25
            default:
                if true {
                    fmt.Println(fi11, fi11, fi11)
                    if (1 != 42) {
                        for fi26 := 0; fi26 < 14; fi26++ {
                            var v27 float32 = 0.1
                            if len(v1) > 0 {
                                v1[0] = 2147483647
                            }
                            var v28 string = "hello"
                            v2 = v2
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                        for fi29 := 0; fi29 < 15; fi29++ {
                            if len(v3) > 0 {
                                v3[0] = (("\n\t" + "hello") + ("foo bar" + "foo bar"))
                            }
                            fmt.Println(fi29)
                            _ = fi29
                        }
                    }
                } else {
                    switch 1 {
                    case 127:
                        switch "" {
                        case "\n\t":
                            v3 = append(v3, "hello")
                        case "foo bar":
                            v30 := GenS3{F0: v2}
                            _ = v30
                        default:
                            if ("" >= ("a" + "hello")) {
                                v2 = (-410 >= 42)
                                _ = len(v3)
                            } else {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v31 int32 = 374
                                switch "" {
                                case "\n\t":
                                    if (v2 && v2) {
                                        v2 = (genFn0() != (^(v10)))
                                    } else {
                                        v32 := GenS1{F0: 0, F1: (fi11 & fi11), F2: v10}
                                        _ = v32
                                    }
                                case "":
                                    var v33 uint32 = 65535
                                    _ = v33
                                case "hello":
                                    for fi34 := 0; fi34 < 6; fi34++ {
                                        v31 = 42
                                        _ = fi34
                                    }
                                default:
                                    if (genFn1(255) >= 131) {
                                        _ = len(v3)
                                    } else {
                                        switch "foo bar" {
                                        case "hello":
                                            switch (fi11 >> ((2) & 63)) {
                                            case 4294967295:
                                                if len(v3) > 0 {
                                                    v3[0] = ("" + "foo bar")
                                                }
                                            case -9223372036854775808:
                                                v2 = (!v2)
                                            case 398:
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            }
                                        case "\n\t":
                                            fmt.Println(v10, v31, fi11)
                                        case "foo bar":
                                            for fi35 := 0; fi35 < 6; fi35++ {
                                                var v36 int32 = (v31 << ((genFn0()) & 63))
                                                _ = fi35
                                                _ = v36
                                            }
                                        default:
                                            var v37 string = "hello"
                                            _ = v37
                                        }
                                        var v38 float64 = 0.1
                                        _ = v38
                                    }
                                }
                                _ = v31
                            }
                        }
                    default:
                        switch ("" + ("" + ((("\n\t" + "a") + "hello") + "foo bar"))) {
                        case "hello":
                            v10 = v10
                        case "\n\t":
                            for fi39 := 0; fi39 < 16; fi39++ {
                                fmt.Println(fi39)
                                _ = fi39
                            }
                        case "a":
                            for fi40 := 0; fi40 < 16; fi40++ {
                                var v41 uint32 = 532
                                v10 += 415
                                _ = fi40
                                _ = v41
                            }
                        }
                    }
                }
            }
        }
        v3 = append(v3, "a")
        _ = fi11
    }
    var v42 uint64 = 369
    if v2 {
        var v43 int8 = 0
        _ = v43
    } else {
        switch "\n\t" {
        case "\n\t":
            var v44 bool = false
            _ = v44
        case "hello":
            fmt.Println(v10, v42)
        case "foo bar":
            v2 = (1 <= 0)
        }
    }
    switch ("a" + "foo bar") {
    case "hello":
        v42 = v42
    case "foo bar":
        var v45 int8 = 1
        _ = v45
    }
    for fi46 := 0; fi46 < 2; fi46++ {
        var v47 int = 282
        var v48 float32 = (1.5 + (0.1 * -2.25))
        fmt.Println(v47)
        _ = fi46
        _ = v47
        _ = v48
    }
    v49 := GenS4{F0: ("a" + ("a" + "a"))}
    if len(v1) > 0 {
        v1[0] = 65535
    }
    v10 = v10
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v42
    _ = v49
    _ = v9
}
