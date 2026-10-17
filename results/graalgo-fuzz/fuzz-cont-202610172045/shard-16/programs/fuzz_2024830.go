package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 string }
type GenS2 struct { F0 uint8; F1 uint64; F2 uint8 }
type GenS3 struct { F0 int16; F1 int8; F2 uint64 }
type GenS4 struct { F0 uint16; F1 uint32 }

func main() {
    v1 := []uint32{}
    var v2 int32 = 93
    var v3 float64 = 3.14159
    var v4 float64 = ((((3.14159 * (((1.5 / v3) + 1e10) / v3)) / v3) - v3) + (v3 + v3))
    v5 := make(map[string]string)
    v6 := make(map[string]bool)
    v4 += (v3 / (3.14159 / (float64(0.1) + v4)))
    if ((4294967295 >= 2) || false) {
        v7 := GenS1{F0: 303, F1: 65535, F2: ("a" + "hello")}
        var v8 float32 = 1.5
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        v3 -= ((-2.25 + float64((-2.25 - -2.25))) - 0.1)
        _ = fi9
    }
    v4 = float64(1.5)
    v10 := GenS2{F0: 0, F1: 637, F2: 128}
    if false {
        var v11 uint32 = 2
        for fi12 := 0; fi12 < 11; fi12++ {
            delete(v5, (("a" + "") + "foo bar"))
            v13 := GenS2{F0: 0, F1: 0, F2: 2}
            _ = fi12
            _ = v13
        }
        var v14 int32 = v2
        _ = v11
        _ = v14
    } else {
        switch -878 {
        case 2147483647:
            var v15 int8 = 0
            _ = v15
        case -9223372036854775808:
            if ((true || ((v2 / ((2147483647) | 1)) < (v2 * -1))) || false) {
                v3 += 1.5
            }
        }
        v16 := GenS3{F0: 32767, F1: 0, F2: 1}
        _ = v16
    }
    var v17 uint32 = 846
    if (true && false) {
        for fi18 := 0; fi18 < 10; fi18++ {
            v19 := GenS1{F0: 1000, F1: 2, F2: (((("\n\t" + ("foo bar" + "a")) + "\n\t") + "hello") + "a")}
            _ = fi18
            _ = v19
        }
        v2 -= v2
    } else {
        var v20 int32 = -1
        _ = v20
    }
    switch "" {
    case "foo bar":
        for fi21 := 0; fi21 < 2; fi21++ {
            var v22 int64 = 1
            for fi23 := 0; fi23 < 1; fi23++ {
                if (v17 > (v17 + v17)) {
                    var v24 uint64 = uint64(v17)
                    switch "foo bar" {
                    case "hello":
                        var v25 int64 = (-(v22))
                        if len(v1) > 0 {
                            v1[0] = v17
                        }
                        _ = len(v5)
                        _ = v25
                    case "a":
                        var v26 uint32 = (v17 % (((v17 % ((v17) | 1))) | 1))
                        _ = v26
                    case "\n\t":
                        switch "foo bar" {
                        case "a":
                            v27 := GenS4{F0: 255, F1: (v17 | (^(v17)))}
                            _ = v27
                        case "hello":
                            if ((v17 + v17) != (v17 / (((^(v17))) | 1))) {
                                v17 += 1
                            } else {
                                if false {
                                    v22 += v22
                                    var v28 float64 = -2.25
                                    _ = v28
                                }
                                v22 = v22
                                v4 += v3
                            }
                        case "\n\t":
                            if false {
                                v29 := GenS1{F0: -32768, F1: 253, F2: "foo bar"}
                                v30 := GenS1{F0: 65, F1: 1, F2: (("" + (("\n\t" + ("foo bar" + "\n\t")) + "foo bar")) + ("\n\t" + ""))}
                                _ = v29
                                _ = v30
                            }
                        }
                    }
                    delete(v6, "hello")
                    _ = v24
                } else {
                    var v31 uint16 = 65535
                    delete(v6, "a")
                    delete(v5, "hello")
                    _ = v31
                }
                switch (^(fi21)) {
                case 2147483647:
                    switch "\n\t" {
                    case "a":
                        var v32 float64 = v4
                        _ = v32
                    case "hello":
                        if false {
                            switch fi23 {
                            case 42:
                                var v33 uint64 = 530
                                var v34 float64 = v3
                                _ = v33
                                _ = v34
                            default:
                                if len(v1) > 0 {
                                    v1[0] = v17
                                }
                            }
                            switch fi21 {
                            case 9223372036854775807:
                                if ((false && false) || true) {
                                    v22 += 9223372036854775807
                                } else {
                                    v3 = -2.25
                                    var v35 int32 = v2
                                    var v36 uint8 = 127
                                    _ = v35
                                    _ = v36
                                }
                            case 32768:
                                fmt.Println(fi21, v3)
                            case 1:
                                v4 -= v3
                            }
                            var v37 int64 = v22
                            _ = v37
                        }
                    default:
                        _ = len(v1)
                    }
                case 9223372036854775807:
                    var v38 uint = 103
                    _ = v38
                }
                var v39 int16 = -528
                _ = fi23
                _ = v39
            }
            switch fi21 {
            case 2147483647:
                if false {
                    v17 += 2147483648
                    v5["hello"] = "foo bar"
                    var v40 string = (("\n\t" + "a") + "")
                    _ = v40
                }
            case 128:
                switch "hello" {
                case "foo bar":
                    if (v4 <= v3) {
                        var v41 int8 = -1
                        _ = v41
                    } else {
                        v3 += v4
                    }
                case "\n\t":
                    v4 = v3
                }
            default:
                for fi42 := 0; fi42 < 5; fi42++ {
                    v4 -= 0.1
                    _ = fi42
                }
            }
            _ = fi21
            _ = v22
        }
    case "hello":
        var v43 uint64 = 1
        _ = v43
    case "\n\t":
        if false {
            fmt.Println(v3)
            var v44 int16 = -32768
            for fi45 := 0; fi45 < 9; fi45++ {
                switch "hello" {
                case "a":
                    v4 = v4
                case "":
                    v44 = v44
                case "\n\t":
                    _ = len(v6)
                }
                fmt.Println(v17, fi45, v2)
                _ = fi45
            }
            _ = v44
        }
    }
    var v46 float32 = float32(1.5)
    var v47 bool = false
    if ((false && (v17 != (v17 / ((42) | 1)))) || true) {
        if v47 {
            v5[""] = "\n\t"
            if (!true) {
                if ((!(9223372036854775807 > 9223372036854775807)) && true) {
                    switch 128 {
                    case -589:
                        switch "" {
                        case "":
                            var v48 int8 = 127
                            _ = v48
                        case "\n\t":
                            var v49 int8 = 22
                            _ = v49
                        }
                    case 2147483647:
                        v50 := GenS1{F0: -32768, F1: 255, F2: "foo bar"}
                        _ = v50
                    case 957:
                        if (1000 > -32768) {
                            var v51 int64 = -729
                            _ = v51
                        } else {
                            for fi52 := 0; fi52 < 2; fi52++ {
                                v47 = ((fi52 / ((-9223372036854775808) | 1)) >= (fi52 & 603))
                                _ = fi52
                            }
                        }
                    }
                    for fi53 := 0; fi53 < 5; fi53++ {
                        switch ("" + "") {
                        case "foo bar":
                            v1 = append(v1, 4294967295)
                        case "\n\t":
                            for fi54 := 0; fi54 < 5; fi54++ {
                                switch "" {
                                case "hello":
                                    v55 := v5["\n\t"]
                                    _ = v55
                                case "foo bar":
                                    var v56 int32 = (-(v2))
                                    v47 = (127 < 84)
                                    _ = v56
                                case "a":
                                    v3 = (0.1 * v4)
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = fi54
                            }
                        case "hello":
                            for fi57 := 0; fi57 < 14; fi57++ {
                                v58 := v5[""]
                                _ = fi57
                                _ = v58
                            }
                        default:
                            fmt.Println(v46)
                        }
                        _ = fi53
                    }
                }
            }
            v17 = v17
        }
        v46 = 1e10
    } else {
        var v59 float64 = v4
        var v60 uint16 = 85
        for fi61 := 0; fi61 < 2; fi61++ {
            v62 := GenS3{F0: -32768, F1: 42, F2: 42}
            var v63 bool = (v60 >= (^(v60)))
            switch "hello" {
            case "a":
                v64 := GenS3{F0: 42, F1: 127, F2: 265}
                _ = v64
            }
            _ = fi61
            _ = v62
            _ = v63
        }
        _ = v59
        _ = v60
    }
    v1 = append(v1, v17)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
