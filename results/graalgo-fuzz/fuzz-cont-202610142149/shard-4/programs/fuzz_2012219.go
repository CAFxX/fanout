package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint64 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 uint; F1 uint; F2 int32 }
type GenS5 struct { F0 string; F1 int8 }
type GenS6 struct { F0 int8 }
func genFn0(p0 int32) string {
    p0 = (p0 & (-(p0)))
    v6 := []int{}
    _ = p0
    _ = v6
    return "foo bar"
}
func genFn1(p0 string) uint64 {
    v7 := []int{128, -940}
    var v8 int16 = -538
    _ = p0
    _ = v7
    _ = v8
    return 18446744073709551615
}

func main() {
    var v1 int32 = -663
    var v2 int32 = (v1 % ((v1) | 1))
    v3 := make(map[int]bool)
    v4 := make(map[int]bool)
    v5 := []uint32{144, 315, 0, 65535}
    if true {
        var v9 bool = (1 <= -1)
        _ = v9
    }
    v10 := GenS1{F0: 4294967295, F1: "", F2: 1}
    var v11 uint8 = 255
    if (392 > -292) {
        switch "hello" {
        case "\n\t":
            var v12 int32 = (v2 ^ -1)
            _ = v12
        case "hello":
            if (!(432 >= 65535)) {
                fmt.Println(v11, v11, v2)
            }
        case "":
            var v13 uint8 = 1
            _ = v13
        }
    } else {
        for fi14 := 0; fi14 < 6; fi14++ {
            v15 := GenS2{F0: (v11 - v11), F1: int32(fi14)}
            if len(v5) > 0 {
                v5[0] = 233
            }
            _ = fi14
            _ = v15
        }
        v11 += (v11 - 127)
        if (true && false) {
            v2 = v1
            v1 += (^(v2))
            switch 311 {
            case 2147483647:
                v11 = v11
            case 32767:
                switch 128 {
                case -175:
                    if (9223372036854775807 != 0) {
                        v1 = v2
                        if (float32(0.1) != 0.1) {
                            switch 0 {
                            case 42:
                                switch "hello" {
                                case "a":
                                    v1 = -2147483648
                                }
                            case 4294967295:
                                for fi16 := 0; fi16 < 10; fi16++ {
                                    fmt.Println(v1, v11, v11)
                                    v2 -= v1
                                    _ = fi16
                                }
                            case 32767:
                                v11 += 0
                            }
                            switch 0 {
                            case 32768:
                                var v17 string = "\n\t"
                                _ = v17
                            case 128:
                                fmt.Println(v11)
                            case 192:
                                for fi18 := 0; fi18 < 9; fi18++ {
                                    var v19 uint = 4294967295
                                    _ = fi18
                                    _ = v19
                                }
                            default:
                                v20 := GenS2{F0: 13, F1: v2}
                                _ = v20
                            }
                            if (!((v2 / ((v1) | 1)) == v2)) {
                                var v21 float64 = 3.14159
                                _ = v21
                            }
                        } else {
                            fmt.Println(v1, v11)
                            fmt.Println(v1, v2, v2)
                            for fi22 := 0; fi22 < 20; fi22++ {
                                fmt.Println(v11, v1, v2)
                                delete(v3, 128)
                                var v23 string = ((("foo bar" + ("foo bar" + "foo bar")) + "foo bar") + "foo bar")
                                _ = fi22
                                _ = v23
                            }
                        }
                        for fi24 := 0; fi24 < 7; fi24++ {
                            for fi25 := 0; fi25 < 6; fi25++ {
                                v11 -= 2
                                _ = fi25
                            }
                            _ = fi24
                        }
                    }
                case 42:
                    switch "\n\t" {
                    case "hello":
                        _ = len(v4)
                    }
                default:
                    v11 = 0
                }
            }
        } else {
            switch ("foo bar" + "") {
            case "hello":
                for fi26 := 0; fi26 < 20; fi26++ {
                    if (13 == v11) {
                        switch "hello" {
                        case "foo bar":
                            v27 := GenS3{F0: -855, F1: (v1 << ((384) & 63))}
                            _ = v27
                        }
                    } else {
                        for fi28 := 0; fi28 < 15; fi28++ {
                            if (56 == 42) {
                                var v29 int64 = -932
                                _ = len(v5)
                                switch (fi28 >> ((65535) & 63)) {
                                case 2147483647:
                                    switch "" {
                                    case "hello":
                                        v2 += 1
                                    default:
                                        fmt.Println(v11, v29)
                                    }
                                }
                                _ = v29
                            }
                            if (true || (((fi28 != fi28) || false) && true)) {
                                v30 := v4[fi28]
                                var v31 uint = 255
                                var v32 uint64 = 393
                                _ = v30
                                _ = v31
                                _ = v32
                            }
                            v33 := GenS4{F0: 115, F1: 4294967295, F2: (v1 - v2)}
                            _ = fi28
                            _ = v33
                        }
                        var v34 int8 = 1
                        _ = v34
                    }
                    v2 = 822
                    _ = fi26
                }
            case "a":
                var v35 int8 = 42
                _ = v35
            case "":
                _ = len(v3)
            default:
                v11 = (v11 ^ 42)
            }
            var v36 int32 = v1
            v4[4294967295] = (!((false || true) || (18446744073709551615 < genFn1("foo bar"))))
            _ = v36
        }
    }
    v37 := GenS3{F0: 4, F1: 1}
    v38 := v3[-2147483648]
    v11 = 125
    if v38 {
        if false {
            _ = len(v4)
            v39 := GenS3{F0: -96, F1: (v2 & -1)}
            _ = v39
        } else {
            fmt.Println(v1, v11)
        }
        v40 := GenS5{F0: "foo bar", F1: -128}
        _ = v40
    }
    var v41 int = -504
    v2 = v1
    var v42 uint32 = 2147483648
    v2 = -496
    if len(v5) > 0 {
        v5[0] = (^(v42))
    }
    for fi43 := 0; fi43 < 2; fi43++ {
        fmt.Println(v42)
        if false {
            v44 := GenS2{F0: v11, F1: (v1 | (v2 / ((v2) | 1)))}
            _ = v44
        } else {
            var v45 bool = ((float64((1.5 - 0.0)) + 1e10) >= 0.1)
            _ = v45
        }
        for fi46 := 0; fi46 < 15; fi46++ {
            v42 -= (^(v42))
            var v47 uint = uint(v11)
            _ = fi46
            _ = v47
        }
        _ = fi43
    }
    v48 := GenS4{F0: 4294967295, F1: 2, F2: -1}
    if false {
        for fi49 := 0; fi49 < 19; fi49++ {
            v50 := GenS3{F0: -472, F1: (v1 ^ (^(v1)))}
            v51 := GenS2{F0: v11, F1: 42}
            for fi52 := 0; fi52 < 10; fi52++ {
                switch fi49 {
                case 128:
                    var v53 uint64 = 18446744073709551615
                    _ = v53
                case 4294967295:
                    var v54 uint16 = 103
                    v2 = 0
                    _ = v54
                case 42:
                    if len(v5) > 0 {
                        v5[0] = (v42 % ((246) | 1))
                    }
                default:
                    v55 := GenS4{F0: 4294967295, F1: 42, F2: 734}
                    _ = v55
                }
                v56 := GenS4{F0: 45, F1: 255, F2: v2}
                v57 := GenS3{F0: 346, F1: (^(v1))}
                var v58 uint64 = 2
                switch v41 {
                case 47:
                    var v59 int8 = 100
                    _ = v59
                case 9223372036854775807:
                    fmt.Println(v58)
                }
                _ = fi52
                _ = v56
                _ = v57
                _ = v58
            }
            _ = fi49
            _ = v50
            _ = v51
        }
        if (18446744073709551615 > 65535) {
            if len(v5) > 0 {
                v5[0] = (v42 | v42)
            }
            if ("foo bar" != ("foo bar" + (("\n\t" + "\n\t") + "\n\t"))) {
                v60 := GenS6{F0: 33}
                _ = v60
            } else {
                for fi61 := 0; fi61 < 13; fi61++ {
                    var v62 uint16 = 2
                    switch "a" {
                    case "\n\t":
                        if (v1 < (v2 ^ 482)) {
                            v2 = v2
                            var v63 float64 = 0.0
                            var v64 bool = (!(genFn1(((("a" + "a") + ("foo bar" + (("hello" + "hello") + ""))) + (("hello" + "a") + ""))) != 109))
                            _ = v63
                            _ = v64
                        }
                    default:
                        var v65 int16 = 42
                        _ = v65
                    }
                    _ = fi61
                    _ = v62
                }
                v4[v41] = false
            }
            v38 = v38
        }
    }
    var v66 uint = 2
    var v67 uint8 = uint8(v66)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v41
    _ = v42
    _ = v48
    _ = v5
    _ = v66
    _ = v67
}
