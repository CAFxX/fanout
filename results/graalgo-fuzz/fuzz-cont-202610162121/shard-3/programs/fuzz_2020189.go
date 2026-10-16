package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 int8 }
type GenS2 struct { F0 uint64; F1 uint16; F2 int64 }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 int; F1 uint64 }
type GenS5 struct { F0 bool; F1 int; F2 int64 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint32; F1 uint8; F2 int8 }
type GenS9 struct { F0 int16 }

func main() {
    var v1 uint64 = 1
    var v2 int16 = 1000
    v3 := make(map[string]string)
    v4 := GenS1{F0: 252, F1: 42, F2: 1}
    if true {
        v5 := GenS1{F0: 42, F1: 42, F2: 127}
        if (true && ((v2 > v2) || true)) {
            fmt.Println(v1, v2)
            var v6 bool = true
            _ = v6
        } else {
            v2 = (v2 & v2)
            for fi7 := 0; fi7 < 2; fi7++ {
                if true {
                    var v8 int16 = 32767
                    if true {
                        for fi9 := 0; fi9 < 4; fi9++ {
                            v10 := v3[(("a" + "foo bar") + "\n\t")]
                            v11 := []int{}
                            _ = len(v11)
                            v12 := GenS1{F0: 42, F1: 970, F2: 0}
                            var v13 bool = true
                            _ = fi9
                            _ = v10
                            _ = v11
                            _ = v12
                            _ = v13
                        }
                        for fi14 := 0; fi14 < 7; fi14++ {
                            fmt.Println(v8, fi14, fi7)
                            if false {
                                v15 := GenS1{F0: 9, F1: 1, F2: 111}
                                if true {
                                    v16 := GenS1{F0: 127, F1: 0, F2: 91}
                                    v17 := GenS1{F0: 1, F1: 0, F2: 1}
                                    _ = v16
                                    _ = v17
                                } else {
                                    switch (("hello" + ((("" + "foo bar") + ("" + ("hello" + ("" + "a")))) + ("a" + "\n\t"))) + ("a" + ("" + (("foo bar" + (("\n\t" + (("\n\t" + (("a" + "hello") + "hello")) + "\n\t")) + ("hello" + ("a" + (("\n\t" + "\n\t") + "foo bar"))))) + "a")))) {
                                    case "a":
                                        fmt.Println(fi14)
                                    }
                                    v2 = v2
                                    v2 -= (v8 * 32767)
                                }
                                _ = v15
                            } else {
                                fmt.Println(v8)
                                v1 = v1
                                v8 -= v8
                            }
                            fmt.Println(fi14, v8)
                            v18 := []int{fi14, 0, (^(fi14)), fi7}
                            var v19 uint8 = 2
                            v1 = 633
                            _ = fi14
                            _ = v18
                            _ = v19
                        }
                        v20 := GenS1{F0: 226, F1: -2147483648, F2: 0}
                        _ = v20
                    } else {
                        for fi21 := 0; fi21 < 6; fi21++ {
                            var v22 int = 128
                            v23 := GenS1{F0: 2, F1: -669, F2: -92}
                            if (9223372036854775807 == -9223372036854775808) {
                                var v24 float32 = (-2.25 / (1.5 * 1.5))
                                var v25 int32 = 0
                                v26 := []int{(v22 >> ((946) & 63)), (fi7 % ((-171) | 1)), int(v25), (v22 / ((2147483647) | 1))}
                                _ = v24
                                _ = v25
                                _ = v26
                            } else {
                                var v27 string = "foo bar"
                                _ = v27
                            }
                            if (true || false) {
                                v28 := GenS2{F0: (v1 % ((v1) | 1)), F1: 0, F2: -769}
                                var v29 uint32 = 2147483647
                                _ = v28
                                _ = v29
                            }
                            var v30 bool = true
                            v8 = (v8 | (v8 * (v8 * (^(v2)))))
                            v8 = (v2 / (((v2 % ((1000) | 1))) | 1))
                            v1 = (v1 + 80)
                            var v31 int16 = v8
                            v32 := GenS2{F0: v1, F1: 944, F2: 661}
                            v33 := GenS3{F0: "\n\t", F1: v30}
                            fmt.Println(v1)
                            v34 := GenS4{F0: v22, F1: (v1 * (^(v1)))}
                            _ = fi21
                            _ = v22
                            _ = v23
                            _ = v30
                            _ = v31
                            _ = v32
                            _ = v33
                            _ = v34
                        }
                    }
                    v35 := []int{}
                    _ = v35
                    _ = v8
                } else {
                    var v36 bool = false
                    v36 = false
                    if v36 {
                        if (("a" >= "foo bar") || (v36 && v36)) {
                            v2 = (v2 % (((v2 << ((42) & 63))) | 1))
                        } else {
                            _ = len(v3)
                            v37 := GenS3{F0: "a", F1: false}
                            _ = v37
                        }
                        v38 := []int{(fi7 & (fi7 ^ (fi7 % ((-2147483648) | 1))))}
                        if v36 {
                            var v39 int8 = -1
                            _ = v39
                        } else {
                            _ = len(v3)
                            v40 := GenS5{F0: false, F1: (fi7 | 128), F2: 909}
                            var v41 uint16 = 0
                            _ = v40
                            _ = v41
                        }
                        _ = v38
                    }
                    _ = v36
                }
                v2 -= (-(v2))
                _ = fi7
            }
            switch "\n\t" {
            case "":
                v2 += v2
            case "foo bar":
                switch ("\n\t" + "hello") {
                case "foo bar":
                    if false {
                        _ = len(v3)
                        v2 = v2
                        switch 42 {
                        case 9223372036854775807:
                            if false {
                                v2 = (v2 * int16(v1))
                                var v42 int64 = 42
                                switch -662 {
                                case 42:
                                    var v43 int = 9223372036854775807
                                    _ = v43
                                default:
                                    if true {
                                        var v44 int64 = 216
                                        switch -626 {
                                        case 666:
                                            fmt.Println(v1)
                                        case 1:
                                            if (!true) {
                                                v2 = v2
                                                _ = len(v3)
                                            }
                                        default:
                                            switch 795 {
                                            case -1:
                                                v45 := GenS2{F0: v1, F1: 0, F2: v44}
                                                _ = v45
                                            }
                                        }
                                        var v46 uint64 = 18446744073709551615
                                        _ = v44
                                        _ = v46
                                    } else {
                                        var v47 uint = 836
                                        v48 := GenS2{F0: 0, F1: 436, F2: (v42 >> ((255) & 63))}
                                        var v49 int8 = int8(v47)
                                        _ = v47
                                        _ = v48
                                        _ = v49
                                    }
                                }
                                _ = v42
                            } else {
                                for fi50 := 0; fi50 < 10; fi50++ {
                                    var v51 string = "\n\t"
                                    var v52 uint32 = 65535
                                    v52 = v52
                                    _ = fi50
                                    _ = v51
                                    _ = v52
                                }
                                if true {
                                    var v53 int8 = -116
                                    _ = v53
                                }
                                if true {
                                    var v54 int8 = 62
                                    v55 := GenS6{F0: false}
                                    _ = v54
                                    _ = v55
                                }
                            }
                        }
                    }
                case "":
                    var v56 uint16 = 13
                    _ = v56
                default:
                    switch "foo bar" {
                    case "":
                        v1 = v1
                    case "foo bar":
                        v57 := GenS7{F0: 2147483647}
                        _ = v57
                    case "\n\t":
                        switch (("" + "") + (("\n\t" + "") + "")) {
                        case "a":
                            if (!true) {
                                v58 := []int{127, 452}
                                _ = v58
                            }
                        default:
                            var v59 uint32 = 6
                            _ = v59
                        }
                    }
                }
            default:
                switch 2147483647 {
                case 32768:
                    v1 = v1
                case -2147483648:
                    v60 := GenS8{F0: 897, F1: 127, F2: -128}
                    _ = v60
                }
            }
        }
        v1 -= 1
        _ = v5
    }
    v61 := GenS9{F0: (v2 + (v2 | v2))}
    var v62 int64 = 9223372036854775807
    fmt.Println(v2)
    v63 := GenS1{F0: 128, F1: 0, F2: 0}
    v62 = v62
    fmt.Println(v1, v1)
    switch (("a" + "foo bar") + ("foo bar" + ("a" + ""))) {
    case "\n\t":
        var v64 int = 4294967295
        _ = v64
    case "":
        for fi65 := 0; fi65 < 5; fi65++ {
            var v66 uint32 = 1
            _ = len(v3)
            v62 = v62
            _ = fi65
            _ = v66
        }
    default:
        fmt.Println(v1, v2, v62)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v61
    _ = v62
    _ = v63
}
