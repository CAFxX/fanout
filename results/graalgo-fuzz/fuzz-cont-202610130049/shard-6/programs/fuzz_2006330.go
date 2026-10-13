package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 uint32; F1 int64; F2 uint8 }

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[uint32]string)
    v3 := make(map[string]string)
    var v4 uint64 = 38
    v5 := []int{1, 522, 42, 437}
    var v6 uint32 = 42
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := GenS1{F0: (v6 * 2147483648), F1: 0}
        switch ((("foo bar" + (("hello" + "") + "")) + "hello") + (("foo bar" + (("a" + ("\n\t" + "\n\t")) + ("" + "hello"))) + ("a" + "a"))) {
        case "\n\t":
            v5 = append(v5, fi7)
        case "":
            v5 = append(v5, fi7)
        case "foo bar":
            for fi9 := 0; fi9 < 3; fi9++ {
                var v10 uint64 = (v4 ^ uint64(v6))
                fmt.Println(v10, v4, v10)
                delete(v2, 2)
                _ = fi9
                _ = v10
            }
        }
        _ = fi7
        _ = v8
    }
    switch 42 {
    case 4294967295:
        if (2147483647 >= 0) {
            v1 = v1
        }
    case 938:
        v1 += 1e10
    }
    var v11 bool = false
    var v12 float64 = ((1.5 * -2.25) + -2.25)
    fmt.Println(v11)
    switch ("hello" + "foo bar") {
    case "foo bar":
        switch "hello" {
        case "foo bar":
            var v13 uint = 0
            _ = v13
        case "\n\t":
            for fi14 := 0; fi14 < 3; fi14++ {
                v15 := GenS2{F0: 1000}
                _ = fi14
                _ = v15
            }
        case "":
            v5 = append(v5, 537)
        }
    }
    fmt.Println(v4, v6)
    for fi16 := 0; fi16 < 12; fi16++ {
        v1 += v1
        var v17 bool = v11
        v6 = (^(v6))
        _ = fi16
        _ = v17
    }
    switch -484 {
    case 32768:
        if true {
            fmt.Println(v11)
            for fi18 := 0; fi18 < 6; fi18++ {
                delete(v3, "foo bar")
                var v19 uint32 = 344
                for fi20 := 0; fi20 < 14; fi20++ {
                    v1 -= v1
                    if len(v5) > 0 {
                        v5[0] = (fi20 * fi18)
                    }
                    _ = fi20
                }
                _ = fi18
                _ = v19
            }
            if v11 {
                for fi21 := 0; fi21 < 5; fi21++ {
                    _ = len(v5)
                    var v22 uint32 = (v6 | 4294967295)
                    v6 -= 2147483648
                    _ = fi21
                    _ = v22
                }
                for fi23 := 0; fi23 < 13; fi23++ {
                    v4 = 413
                    v4 = (v4 << ((482) & 63))
                    v24 := GenS2{F0: -726}
                    _ = fi23
                    _ = v24
                }
                var v25 uint8 = 1
                _ = v25
            } else {
                for fi26 := 0; fi26 < 5; fi26++ {
                    switch 740 {
                    case -9223372036854775808:
                        for fi27 := 0; fi27 < 14; fi27++ {
                            v28 := v2[v6]
                            v29 := GenS3{F0: 58, F1: 0}
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                    default:
                        switch fi26 {
                        case -788:
                            for fi30 := 0; fi30 < 14; fi30++ {
                                var v31 string = ("" + (("a" + "foo bar") + ("\n\t" + "")))
                                v11 = v11
                                var v32 uint32 = 953
                                _ = fi30
                                _ = v31
                                _ = v32
                            }
                        case 32767:
                            var v33 float64 = v12
                            _ = v33
                        case 2147483647:
                            var v34 uint = 244
                            _ = v34
                        default:
                            v35 := GenS4{F0: v6, F1: -9223372036854775808, F2: 2}
                            _ = v35
                        }
                    }
                    _ = fi26
                }
                if (v11 && ((((("foo bar" + "\n\t") + "foo bar") + ("" + "hello")) == (("hello" + "a") + "a")) || v11)) {
                    fmt.Println(v12, v12)
                    v3[("foo bar" + ("a" + ""))] = "\n\t"
                    switch 2147483647 {
                    case -9223372036854775808:
                        v6 -= v6
                    default:
                        var v36 uint = 4294967295
                        _ = v36
                    }
                }
            }
        }
    case 32767:
        fmt.Println(v6)
    }
    if v11 {
        v3[""] = "foo bar"
        v37 := GenS4{F0: (^(v6)), F1: 42, F2: 127}
        var v38 uint32 = (v6 / ((v6) | 1))
        _ = v37
        _ = v38
    } else {
        var v39 uint32 = v6
        _ = v39
    }
    switch 767 {
    case 1:
        v11 = (1e10 <= v1)
    case 627:
        v4 += v4
    }
    if v11 {
        v2[(v6 >> ((18446744073709551615) & 63))] = "foo bar"
        fmt.Println(v1)
    } else {
        v4 -= v4
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
