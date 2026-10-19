package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 uint; F2 int64 }
type GenS3 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 uint8; F1 int; F2 string }
type GenS6 struct { F0 uint }
func genFn0(p0 bool, p1 int16, p2 uint) int {
    v5 := []int{}
    v6 := GenS1{F0: 1}
    switch 249 {
    case -580:
        var v7 uint64 = 18446744073709551615
        _ = v7
    case 1:
        if (!(32767 == p1)) {
            fmt.Println(p1, p2, p0)
            p1 -= (p1 + (p1 * p1))
        } else {
            v8 := GenS1{F0: 0}
            _ = len(v5)
            fmt.Println(p1)
            _ = v8
        }
    case 0:
        var v9 uint8 = 233
        _ = v9
    default:
        var v10 uint32 = 0
        _ = v10
    }
    p0 = p0
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 511
}
func genFn1(p0 int, p1 int64) int64 {
    for fi11 := 0; fi11 < 3; fi11++ {
        for fi12 := 0; fi12 < 8; fi12++ {
            v13 := []int{}
            if false {
                fmt.Println(fi12)
                _ = len(v13)
            }
            p1 += p1
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    if (1 != 1) {
        p1 = (p1 / ((-11) | 1))
    }
    var v14 float64 = 0.0
    var v15 uint32 = 531
    switch (p0 - p0) {
    case -401:
        if false {
            var v16 uint8 = 1
            p0 = (p0 % (((p0 | p0)) | 1))
            fmt.Println(v16)
            _ = v16
        } else {
            v17 := []int{}
            _ = v17
        }
    case -198:
        v15 = v15
    }
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    return -9223372036854775808
}

func main() {
    v1 := []string{("hello" + "a"), (("foo bar" + "a") + (("a" + ("hello" + "")) + "a")), ("\n\t" + "foo bar"), "a"}
    var v2 uint64 = 42
    v3 := make(map[string]bool)
    var v4 int32 = -797
    v4 = v4
    var v18 uint64 = 174
    for fi19 := 0; fi19 < 18; fi19++ {
        v2 += (v2 + (v18 & 2))
        _ = fi19
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v20 := GenS2{F0: genFn0((42 < 2147483648), 1, 512), F1: 2, F2: 1}
    switch "a" {
    case "\n\t":
        if ((true && (2 > 790)) || false) {
            v1 = append(v1, "\n\t")
        }
    default:
        for fi21 := 0; fi21 < 12; fi21++ {
            if true {
                for fi22 := 0; fi22 < 17; fi22++ {
                    if (true && false) {
                        if (v2 != 735) {
                            v4 += 689
                        }
                        v4 = v4
                        var v23 uint64 = 351
                        fmt.Println(v4, fi22, v4)
                        if false {
                            fmt.Println(fi21, v2)
                        } else {
                            var v24 bool = true
                            fmt.Println(v4, v2, v2)
                            if (!(255 <= 104)) {
                                var v25 uint = 758
                                _ = v25
                            } else {
                                v3[((("a" + "foo bar") + ("foo bar" + "a")) + ("\n\t" + "hello"))] = false
                                v3[""] = (!v24)
                                if v24 {
                                    v26 := GenS3{F0: 292, F1: 42, F2: (v2 - v2)}
                                    _ = v26
                                }
                            }
                            _ = v24
                        }
                        _ = v23
                    }
                    var v27 int32 = 1
                    _ = fi22
                    _ = v27
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            } else {
                _ = len(v1)
                var v28 uint64 = v2
                _ = v28
            }
            if len(v1) > 0 {
                v1[0] = ((("a" + (("" + "\n\t") + ("a" + ("" + "")))) + (("foo bar" + "hello") + "hello")) + "\n\t")
            }
            v2 = v18
            _ = fi21
        }
    }
    fmt.Println(v2)
    for fi29 := 0; fi29 < 10; fi29++ {
        for fi30 := 0; fi30 < 7; fi30++ {
            v2 += (v2 | v2)
            if (-242 == -9223372036854775808) {
                v18 = 2
            } else {
                if (!false) {
                    var v31 uint64 = (^(v18))
                    var v32 int64 = genFn1(genFn0((4294967295 != 314), -32768, 255), genFn1(1, -186))
                    v18 = v18
                    _ = v31
                    _ = v32
                }
            }
            v2 -= 213
            _ = fi30
        }
        _ = fi29
    }
    if ((255 < 650) || false) {
        if true {
            if (!(255 > 42)) {
                for fi33 := 0; fi33 < 1; fi33++ {
                    if (fi33 <= fi33) {
                        v34 := GenS3{F0: 0, F1: 2147483648, F2: v18}
                        for fi35 := 0; fi35 < 13; fi35++ {
                            switch (fi33 | (fi35 + fi35)) {
                            case 4294967295:
                                var v36 uint32 = 4294967295
                                _ = v36
                            case 32768:
                                if (-128 <= 42) {
                                    fmt.Println(fi33)
                                } else {
                                    switch ("hello" + (("" + ("foo bar" + ("\n\t" + "a"))) + "foo bar")) {
                                    case "hello":
                                        fmt.Println(fi33)
                                    case "":
                                        v18 -= (v18 | v2)
                                    }
                                }
                            case 32767:
                                switch (("" + "foo bar") + "") {
                                case "foo bar":
                                    if (226 != 1) {
                                        var v37 int = (fi35 * 0)
                                        _ = v37
                                    }
                                default:
                                    v2 -= v2
                                }
                            }
                            _ = fi35
                        }
                        _ = v34
                    } else {
                        fmt.Println(v18, v2, v2)
                        fmt.Println(fi33, v2)
                        fmt.Println(fi33, v4)
                    }
                    fmt.Println(v4, fi33, v18)
                    _ = fi33
                }
                v38 := GenS4{F0: -1, F1: -128}
                _ = v38
            } else {
                v4 = 114
                v18 = v2
                var v39 uint32 = 629
                _ = v39
            }
            v40 := GenS5{F0: 127, F1: 1, F2: "a"}
            _ = v40
        }
    } else {
        _ = len(v3)
        var v41 bool = true
        if (v41 || v41) {
            for fi42 := 0; fi42 < 11; fi42++ {
                v18 = 431
                for fi43 := 0; fi43 < 18; fi43++ {
                    var v44 float32 = 3.14159
                    v45 := v3[("" + "a")]
                    var v46 uint64 = v18
                    _ = fi43
                    _ = v44
                    _ = v45
                    _ = v46
                }
                _ = fi42
            }
            var v47 string = "a"
            _ = v47
        }
        _ = v41
    }
    v18 = 9
    var v48 uint8 = 128
    v49 := GenS6{F0: 611}
    v2 += (v18 + v18)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v48
    _ = v49
}
