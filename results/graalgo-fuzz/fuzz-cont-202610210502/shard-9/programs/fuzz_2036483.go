package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int32; F1 uint32; F2 uint64 }
type GenS5 struct { F0 uint16; F1 uint16 }
func genFn0(p0 uint64, p1 bool) uint8 {
    var v5 uint16 = 812
    p0 = p0
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 int64 = 1
        v7 += v7
        var v8 int8 = -58
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v5
    return uint8(p0)
}

func main() {
    var v1 float32 = 0.1
    v2 := []uint32{65535, 2, 791, 2}
    var v3 uint = 602
    v4 := make(map[uint32]int)
    fmt.Println(v1, v3)
    v9 := GenS1{F0: "hello"}
    fmt.Println(v1, v3)
    v10 := GenS1{F0: "\n\t"}
    v11 := GenS1{F0: ((((("" + "foo bar") + ("" + "hello")) + "hello") + "hello") + "a")}
    v3 = v3
    var v12 int8 = -1
    if false {
        var v13 float64 = 3.14159
        _ = len(v2)
        for fi14 := 0; fi14 < 20; fi14++ {
            v2 = append(v2, 42)
            _ = len(v4)
            _ = fi14
        }
        _ = v13
    } else {
        switch -2147483648 {
        case 557:
            v12 = v12
        case 4294967295:
            if true {
                switch "" {
                case "\n\t":
                    var v15 bool = false
                    _ = v15
                case "hello":
                    var v16 uint64 = 42
                    _ = v16
                case "foo bar":
                    var v17 int = 755
                    _ = v17
                default:
                    var v18 uint64 = 10
                    _ = v18
                }
                var v19 int16 = 0
                if (179 <= 105) {
                    switch "a" {
                    case "":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case "\n\t":
                        v12 -= v12
                    case "hello":
                        for fi20 := 0; fi20 < 4; fi20++ {
                            switch fi20 {
                            case 951:
                                v12 = v12
                            case 2147483647:
                                var v21 int8 = v12
                                _ = v21
                            case -1:
                                for fi22 := 0; fi22 < 3; fi22++ {
                                    var v23 uint32 = 584
                                    v1 = v1
                                    v12 -= v12
                                    _ = fi22
                                    _ = v23
                                }
                            }
                            if (!(!true)) {
                                switch (("" + "foo bar") + ((("foo bar" + "hello") + ("a" + "hello")) + "\n\t")) {
                                case "hello":
                                    for fi24 := 0; fi24 < 14; fi24++ {
                                        var v25 uint32 = 2
                                        var v26 int32 = 42
                                        _ = fi24
                                        _ = v25
                                        _ = v26
                                    }
                                default:
                                    var v27 float64 = 1.5
                                    _ = v27
                                }
                            }
                            for fi28 := 0; fi28 < 3; fi28++ {
                                v3 += v3
                                _ = fi28
                            }
                            _ = fi20
                        }
                    default:
                        fmt.Println(v1)
                    }
                    if len(v2) > 0 {
                        v2[0] = 4294967295
                    }
                } else {
                    for fi29 := 0; fi29 < 4; fi29++ {
                        v30 := GenS1{F0: "hello"}
                        for fi31 := 0; fi31 < 8; fi31++ {
                            if true {
                                switch ((("" + (("" + "a") + "foo bar")) + "\n\t") + (((("" + "") + "foo bar") + "foo bar") + ("\n\t" + "\n\t"))) {
                                case "foo bar":
                                    v19 -= v19
                                case "\n\t":
                                    fmt.Println(v19)
                                case "hello":
                                    v1 += (v1 + 0.1)
                                default:
                                    var v32 string = ("a" + "foo bar")
                                    _ = v32
                                }
                                v19 = 32767
                                fmt.Println(v12, v3, v12)
                            } else {
                                v4[2] = fi29
                            }
                            _ = fi31
                        }
                        for fi33 := 0; fi33 < 7; fi33++ {
                            if true {
                                var v34 int8 = 0
                                if true {
                                    if false {
                                        var v35 float64 = 0.1
                                        var v36 uint32 = 2147483648
                                        var v37 uint16 = 555
                                        _ = v35
                                        _ = v36
                                        _ = v37
                                    } else {
                                        switch fi33 {
                                        case 141:
                                            v3 = v3
                                        }
                                        v38 := GenS1{F0: "a"}
                                        fmt.Println(v1)
                                        _ = v38
                                    }
                                    v12 += (^(v12))
                                } else {
                                    var v39 int32 = -1
                                    v2 = append(v2, 2)
                                    _ = v39
                                }
                                _ = v34
                            }
                            fmt.Println(fi29, v19, v12)
                            switch (fi29 >> ((255) & 63)) {
                            case -863:
                                switch (fi33 / ((128) | 1)) {
                                case 128:
                                    if ((^(v3)) <= (v3 << ((v3) & 63))) {
                                        var v40 uint32 = 304
                                        v41 := GenS2{F0: -748}
                                        if true {
                                            fmt.Println(v19, v1, fi29)
                                        } else {
                                            var v42 bool = (!false)
                                            v3 = 65535
                                            _ = v42
                                        }
                                        var v43 int64 = -9223372036854775808
                                        v44 := GenS3{F0: 0}
                                        _ = v40
                                        _ = v41
                                        _ = v43
                                        _ = v44
                                    }
                                case 32767:
                                    var v45 int8 = v12
                                    _ = v45
                                case 9223372036854775807:
                                    if false {
                                        if (42 < 1) {
                                            v1 -= 1.5
                                        }
                                        v3 -= (^(v3))
                                        if len(v2) > 0 {
                                            v2[0] = 365
                                        }
                                    } else {
                                        if (v12 < (v12 * v12)) {
                                            var v46 int8 = 0
                                            v46 = 100
                                            _ = len(v2)
                                            _ = v46
                                        }
                                        var v47 int64 = 846
                                        _ = v47
                                    }
                                default:
                                    switch 128 {
                                    case 2147483647:
                                        v3 += v3
                                    case 616:
                                        v12 += (v12 + v12)
                                    }
                                }
                            case 4294967295:
                                if len(v2) > 0 {
                                    v2[0] = 42
                                }
                            default:
                                v4[919] = fi33
                            }
                            _ = fi33
                        }
                        _ = fi29
                        _ = v30
                    }
                }
                _ = v19
            } else {
                v48 := GenS1{F0: "foo bar"}
                v49 := GenS4{F0: 2147483647, F1: 0, F2: 4294967295}
                _ = v48
                _ = v49
            }
        case 0:
            v3 = (v3 / ((v3) | 1))
        default:
            if (127 < 12) {
                var v50 string = "a"
                v3 -= v3
                if (!false) {
                    v51 := GenS3{F0: 4294967295}
                    _ = v51
                }
                _ = v50
            } else {
                if true {
                    for fi52 := 0; fi52 < 9; fi52++ {
                        fmt.Println(v12, v12, v3)
                        _ = fi52
                    }
                    if (2147483647 <= 517) {
                        fmt.Println(v3, v1, v1)
                    }
                } else {
                    v2 = append(v2, 2147483647)
                    var v53 uint = (v3 / ((v3) | 1))
                    _ = v53
                }
            }
        }
    }
    v54 := v4[42]
    if false {
        _ = len(v2)
        fmt.Println(v1, v1, v54)
        fmt.Println(v3, v1)
    } else {
        v55 := GenS5{F0: 255, F1: 517}
        _ = v55
    }
    for fi56 := 0; fi56 < 10; fi56++ {
        v57 := GenS3{F0: 42}
        var v58 int16 = -271
        var v59 uint8 = 121
        _ = fi56
        _ = v57
        _ = v58
        _ = v59
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v54:", v54)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v54
    _ = v9
}
