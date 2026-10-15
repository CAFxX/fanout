package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint }
type GenS2 struct { F0 uint; F1 int32; F2 uint64 }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint; F1 uint }

func main() {
    var v1 float64 = 1e10
    var v2 int = 4294967295
    v3 := []string{"foo bar", ("" + ((((("" + "foo bar") + "\n\t") + "a") + "foo bar") + "foo bar"))}
    var v4 uint64 = 4294967295
    var v5 int64 = 42
    if true {
        switch (("a" + ("a" + "foo bar")) + "") {
        case "hello":
            v1 = v1
        case "a":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case "\n\t":
            if true {
                for fi6 := 0; fi6 < 8; fi6++ {
                    if (42 >= 1) {
                        v7 := GenS1{F0: false, F1: 1, F2: 65535}
                        for fi8 := 0; fi8 < 15; fi8++ {
                            fmt.Println(v2)
                            _ = fi8
                        }
                        _ = v7
                    }
                    v4 += v4
                    _ = fi6
                }
            }
        default:
            var v9 uint16 = 289
            _ = v9
        }
        if ((true && (18446744073709551615 >= (v4 % ((v4) | 1)))) && (0 != int32(v4))) {
            var v10 int8 = 100
            _ = len(v3)
            _ = v10
        } else {
            var v11 string = (("hello" + "") + ("" + ""))
            _ = v11
        }
        v12 := GenS2{F0: 519, F1: 1, F2: v4}
        _ = v12
    } else {
        v13 := GenS2{F0: 0, F1: 1, F2: v4}
        var v14 uint16 = 42
        _ = v13
        _ = v14
    }
    fmt.Println(v1, v4)
    v1 = -2.25
    if (271 <= 1) {
        var v15 string = "hello"
        _ = v15
    } else {
        v1 -= (0.0 / 1.0)
        for fi16 := 0; fi16 < 8; fi16++ {
            for fi17 := 0; fi17 < 15; fi17++ {
                if false {
                    fmt.Println(v4)
                    var v18 int16 = 791
                    if true {
                        v4 -= v4
                    } else {
                        var v19 uint16 = 2
                        _ = v19
                    }
                    _ = v18
                }
                _ = fi17
            }
            _ = fi16
        }
    }
    v1 = v1
    fmt.Println(v5)
    var v20 float32 = 0.0
    switch "foo bar" {
    case "a":
        if false {
            for fi21 := 0; fi21 < 19; fi21++ {
                v22 := GenS2{F0: 255, F1: -1, F2: v4}
                var v23 int8 = -1
                v2 -= v2
                _ = fi21
                _ = v22
                _ = v23
            }
        }
    }
    v5 += 9223372036854775807
    for fi24 := 0; fi24 < 15; fi24++ {
        for fi25 := 0; fi25 < 3; fi25++ {
            v1 -= (v1 + v1)
            switch -597 {
            case 9223372036854775807:
                var v26 int32 = 931
                _ = v26
            default:
                switch fi25 {
                case -426:
                    v27 := GenS3{F0: ("\n\t" + ("foo bar" + ("foo bar" + "a"))), F1: 42}
                    _ = v27
                default:
                    if (-1 == 1) {
                        fmt.Println(v20, fi25, v1)
                        v28 := GenS3{F0: "foo bar", F1: (v5 & (v5 | (-(v5))))}
                        _ = v28
                    }
                }
            }
            v3 = append(v3, ("" + "hello"))
            _ = fi25
        }
        v29 := GenS2{F0: 255, F1: 1, F2: (v4 & (v4 ^ v4))}
        _ = fi24
        _ = v29
    }
    for fi30 := 0; fi30 < 5; fi30++ {
        v31 := GenS2{F0: 80, F1: 2147483647, F2: (^(v4))}
        if false {
            if true {
                v3 = append(v3, "")
            } else {
                for fi32 := 0; fi32 < 8; fi32++ {
                    switch "foo bar" {
                    case "foo bar":
                        v1 -= (1.5 - v1)
                    case "hello":
                        v33 := GenS4{F0: v5}
                        _ = v33
                    }
                    _ = len(v3)
                    _ = fi32
                }
            }
            if (false && (-998 > (fi30 / ((fi30) | 1)))) {
                v5 += -993
                for fi34 := 0; fi34 < 4; fi34++ {
                    fmt.Println(v4)
                    v35 := GenS2{F0: 306, F1: -610, F2: (v4 * 42)}
                    v5 -= v5
                    _ = fi34
                    _ = v35
                }
                v3 = append(v3, "")
            } else {
                v36 := GenS5{F0: 42}
                var v37 int64 = 1
                _ = v36
                _ = v37
            }
            switch (fi30 + v2) {
            case 2147483647:
                var v38 int16 = 1
                _ = v38
            case -9223372036854775808:
                v4 = v4
            case 866:
                if (!false) {
                    v4 = (v4 - 864)
                    var v39 uint16 = 332
                    _ = v39
                } else {
                    v40 := GenS6{F0: 65535, F1: 255}
                    v41 := GenS5{F0: 2147483647}
                    var v42 string = "\n\t"
                    _ = v40
                    _ = v41
                    _ = v42
                }
            }
        }
        _ = fi30
        _ = v31
    }
    var v43 bool = (true || true)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v43
    _ = v5
}
