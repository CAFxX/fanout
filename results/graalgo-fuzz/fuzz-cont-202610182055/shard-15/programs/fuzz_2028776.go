package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint; F1 string; F2 uint32 }
type GenS5 struct { F0 uint32; F1 uint8 }
type GenS6 struct { F0 int }
func genFn0(p0 string) uint16 {
    v7 := GenS1{F0: 9223372036854775807, F1: 442}
    if true {
        fmt.Println(p0)
    }
    _ = p0
    _ = v7
    return 1
}
func genFn1(p0 bool, p1 int16, p2 int64) int32 {
    p1 = p1
    p1 += (^(p1))
    if (!p0) {
        v8 := GenS1{F0: 32767, F1: 0}
        v9 := []int{1, -770, 32768}
        _ = v8
        _ = v9
    } else {
        v10 := []int{861, 9223372036854775807, 1}
        v11 := GenS1{F0: -909, F1: 65535}
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return 247
}

func main() {
    v1 := []uint32{4294967295, 2147483647, 4294967295, 2147483648}
    var v2 uint = 336
    v3 := []string{"a", "foo bar", "foo bar"}
    v4 := []uint32{155, 1, 697, 65535}
    var v5 uint8 = uint8(v2)
    var v6 int = 32767
    for fi12 := 0; fi12 < 6; fi12++ {
        switch ("hello" + ("" + "")) {
        case "foo bar":
            for fi13 := 0; fi13 < 15; fi13++ {
                switch fi12 {
                case -979:
                    v2 += (^(v2))
                default:
                    switch "foo bar" {
                    case "a":
                        v14 := GenS2{F0: 603}
                        _ = v14
                    case "hello":
                        var v15 float64 = (1.5 + 0.0)
                        _ = v15
                    default:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                }
                if (-613 == 511) {
                    v16 := GenS3{F0: 100}
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v16
                }
                v17 := GenS1{F0: v6, F1: 711}
                switch (-(fi12)) {
                case -589:
                    v6 -= 4294967295
                case 32767:
                    v5 -= 2
                case -2147483648:
                    v18 := GenS1{F0: (^(fi12)), F1: 2}
                    _ = v18
                default:
                    v6 -= (^(fi13))
                }
                v2 -= v2
                switch "\n\t" {
                case "\n\t":
                    var v19 int = fi13
                    _ = v19
                case "foo bar":
                    v20 := GenS4{F0: (^(v2)), F1: ("" + "\n\t"), F2: 65535}
                    _ = v20
                case "":
                    v2 -= (^(v2))
                }
                _ = fi13
                _ = v17
            }
        case "a":
            if (!(!true)) {
                v4 = append(v4, 1)
                v3 = append(v3, ("a" + "\n\t"))
                switch "foo bar" {
                case "foo bar":
                    if true {
                        var v21 float32 = (1e10 - 1.5)
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        for fi22 := 0; fi22 < 14; fi22++ {
                            if (-9223372036854775808 != 627) {
                                fmt.Println(v21, v2, v6)
                                v6 += fi22
                                if false {
                                    v1 = append(v1, 65535)
                                    if (v21 != 0.0) {
                                        v21 -= 1e10
                                        _ = len(v3)
                                        if ((("hello" + ((("a" + "a") + "foo bar") + ("foo bar" + "hello"))) + "\n\t") != ("\n\t" + ("\n\t" + ("a" + "a")))) {
                                            if true {
                                                v21 -= v21
                                                v23 := GenS3{F0: -1}
                                                _ = v23
                                            }
                                        } else {
                                            v24 := GenS3{F0: 42}
                                            _ = v24
                                        }
                                        var v25 uint16 = 32767
                                        _ = v25
                                    } else {
                                        _ = len(v4)
                                        v3 = append(v3, ("\n\t" + "hello"))
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    }
                                    v4 = append(v4, 65535)
                                    var v26 string = ""
                                    var v27 bool = (false || (790 >= -742))
                                    _ = v26
                                    _ = v27
                                }
                            } else {
                                switch fi12 {
                                case 9223372036854775807:
                                    var v28 int16 = -1
                                    _ = v28
                                }
                                _ = len(v4)
                            }
                            v29 := GenS4{F0: 65535, F1: "\n\t", F2: 2147483648}
                            v30 := GenS5{F0: 42, F1: (^(v5))}
                            _ = fi22
                            _ = v29
                            _ = v30
                        }
                        _ = v21
                    }
                case "\n\t":
                    fmt.Println(fi12)
                case "":
                    if false {
                        for fi31 := 0; fi31 < 15; fi31++ {
                            var v32 float64 = float64(1.5)
                            if ((0 != -542) && ((false && false) || false)) {
                                var v33 uint16 = 355
                                fmt.Println(fi31)
                                _ = v33
                            } else {
                                var v34 uint32 = 939
                                v34 = 42
                                v34 += 998
                                v35 := GenS4{F0: 2, F1: (("foo bar" + "") + (("hello" + ((("hello" + "") + "") + ("foo bar" + "\n\t"))) + "a")), F2: (v34 * 2147483647)}
                                v2 += (v2 + 4294967295)
                                _ = v34
                                _ = v35
                            }
                            _ = fi31
                            _ = v32
                        }
                        var v36 int32 = -215
                        _ = v36
                    } else {
                        switch "" {
                        case "\n\t":
                            switch "\n\t" {
                            case "a":
                                switch "\n\t" {
                                case "\n\t":
                                    fmt.Println(fi12, fi12, v2)
                                case "foo bar":
                                    fmt.Println(v6, v6, v2)
                                }
                            }
                        case "a":
                            if (0 > -418) {
                                v6 += 42
                                if false {
                                    switch fi12 {
                                    case -904:
                                        fmt.Println(fi12, v5)
                                    case 9223372036854775807:
                                        _ = len(v3)
                                    case -1:
                                        var v37 string = ("hello" + "a")
                                        _ = v37
                                    }
                                    var v38 bool = false
                                    v5 = v5
                                    _ = v38
                                }
                            }
                        }
                    }
                }
            } else {
                fmt.Println(v6, fi12)
                var v39 uint16 = 32767
                _ = v39
            }
        case "":
            v40 := GenS5{F0: 0, F1: (^(v5))}
            _ = v40
        default:
            v41 := GenS2{F0: fi12}
            _ = v41
        }
        for fi42 := 0; fi42 < 17; fi42++ {
            v43 := GenS4{F0: 435, F1: "a", F2: 1}
            var v44 string = (("foo bar" + "foo bar") + ("a" + "foo bar"))
            fmt.Println(fi42, fi12, v6)
            v45 := GenS6{F0: fi42}
            _ = fi42
            _ = v43
            _ = v44
            _ = v45
        }
        var v46 int64 = 42
        _ = fi12
        _ = v46
    }
    fmt.Println(v6, v6, v6)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v47 int8 = -128
    _ = len(v3)
    fmt.Println(v47, v2)
    v47 -= int8(v2)
    fmt.Println(v6)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}
