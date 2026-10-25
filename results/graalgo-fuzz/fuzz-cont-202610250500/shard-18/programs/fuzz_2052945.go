package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 int; F1 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int; F1 string; F2 string }
type GenS6 struct { F0 uint; F1 int8; F2 uint64 }

func main() {
    var v1 int16 = 42
    var v2 int = 344
    v3 := make(map[int]int)
    v4 := make(map[int]uint32)
    var v5 bool = false
    v6 := []uint32{2147483647, 2147483648, 997}
    v1 = (^(v1))
    v7 := GenS1{F0: ("a" + ""), F1: 0}
    fmt.Println(v1)
    fmt.Println(v5)
    fmt.Println(v1, v1, v1)
    switch (v2 & v2) {
    case -880:
        if v5 {
            v8 := GenS1{F0: "foo bar", F1: 1}
            switch "hello" {
            case "foo bar":
                var v9 string = ""
                _ = v9
            case "a":
                var v10 uint64 = 0
                _ = v10
            }
            _ = v8
        } else {
            for fi11 := 0; fi11 < 4; fi11++ {
                var v12 uint8 = 128
                _ = fi11
                _ = v12
            }
            for fi13 := 0; fi13 < 8; fi13++ {
                var v14 uint = 65535
                _ = fi13
                _ = v14
            }
            if v5 {
                v5 = true
            } else {
                if len(v6) > 0 {
                    v6[0] = 483
                }
                if v5 {
                    if false {
                        v2 -= (v2 | -335)
                        for fi15 := 0; fi15 < 10; fi15++ {
                            switch "a" {
                            case "foo bar":
                                if (v1 < (v1 << ((18446744073709551615) & 63))) {
                                    if v5 {
                                        v2 = v2
                                        for fi16 := 0; fi16 < 1; fi16++ {
                                            v2 += fi16
                                            v17 := GenS2{F0: v1, F1: 0}
                                            _ = fi16
                                            _ = v17
                                        }
                                    } else {
                                        v18 := GenS1{F0: "foo bar", F1: 0}
                                        var v19 float32 = -2.25
                                        delete(v4, v2)
                                        _ = v18
                                        _ = v19
                                    }
                                    switch (v2 - -9223372036854775808) {
                                    case 662:
                                        delete(v3, fi15)
                                    }
                                    if (-2.25 != 0.0) {
                                        if false {
                                            var v20 uint32 = 65535
                                            switch 613 {
                                            case 0:
                                                v2 -= fi15
                                            case -2147483648:
                                                if v5 {
                                                    v21 := GenS3{F0: v2, F1: 42}
                                                    v22 := GenS1{F0: ("foo bar" + "\n\t"), F1: 197}
                                                    var v23 uint64 = 4294967295
                                                    _ = v21
                                                    _ = v22
                                                    _ = v23
                                                }
                                            case 32767:
                                                if ("" != "a") {
                                                    if v5 {
                                                        v3[(v2 % ((42) | 1))] = 1
                                                        var v24 int64 = 42
                                                        _ = v24
                                                    }
                                                    if len(v6) > 0 {
                                                        v6[0] = (v20 + (v20 % ((v20) | 1)))
                                                    }
                                                } else {
                                                    v25 := GenS1{F0: "\n\t", F1: 128}
                                                    _ = v25
                                                }
                                            default:
                                                if true {
                                                    fmt.Println(v20, v5, v5)
                                                }
                                            }
                                            v20 = 188
                                            _ = v20
                                        }
                                        if len(v6) > 0 {
                                            v6[0] = 914
                                        }
                                    } else {
                                        var v26 int64 = 1
                                        var v27 uint64 = 751
                                        delete(v3, (v2 & v2))
                                        _ = v26
                                        _ = v27
                                    }
                                }
                            default:
                                switch (("hello" + "") + "foo bar") {
                                case "":
                                    v28 := GenS4{F0: 630}
                                    _ = v28
                                case "hello":
                                    _ = len(v6)
                                case "a":
                                    fmt.Println(v2, v5)
                                default:
                                    if v5 {
                                        for fi29 := 0; fi29 < 8; fi29++ {
                                            v30 := GenS1{F0: "\n\t", F1: 15}
                                            if (!false) {
                                                switch fi29 {
                                                case 1:
                                                    switch "hello" {
                                                    case "\n\t":
                                                        v31 := GenS2{F0: v1, F1: 9223372036854775807}
                                                        _ = v31
                                                    case "a":
                                                        var v32 float64 = (0.0 + -2.25)
                                                        _ = v32
                                                    }
                                                }
                                            }
                                            fmt.Println(v2, fi15)
                                            _ = fi29
                                            _ = v30
                                        }
                                    } else {
                                        v2 += v2
                                        var v33 uint = 18446744073709551615
                                        _ = v33
                                    }
                                }
                            }
                            switch "" {
                            case "foo bar":
                                fmt.Println(v1, fi15, v2)
                            case "a":
                                switch v2 {
                                case 681:
                                    fmt.Println(v5, fi15, v5)
                                case -9223372036854775808:
                                    v4[(fi15 >> ((0) & 63))] = 2147483648
                                }
                            default:
                                fmt.Println(v2)
                            }
                            _ = fi15
                        }
                        v1 = 1
                    }
                } else {
                    var v34 string = "hello"
                    v35 := GenS2{F0: v1, F1: 329}
                    _ = v34
                    _ = v35
                }
            }
        }
    }
    var v36 int64 = -418
    v36 = (v36 >> ((18446744073709551615) & 63))
    v2 = v2
    if len(v6) > 0 {
        v6[0] = 748
    }
    v37 := GenS4{F0: 290}
    var v38 int8 = 42
    if (false && (((true && v5) && v5) && true)) {
        for fi39 := 0; fi39 < 18; fi39++ {
            switch (v2 - (^(fi39))) {
            case 128:
                v2 += (fi39 & (-(fi39)))
            case -2147483648:
                var v40 float32 = float32(((0.1 - (float64(0.0) * 0.1)) + 1e10))
                _ = v40
            }
            _ = fi39
        }
        v2 += 128
    } else {
        if (-2.25 < 0.1) {
            if (-212 >= 0) {
                v5 = ((v36 * v36) >= (v36 << ((190) & 63)))
                fmt.Println(v5, v38)
            } else {
                for fi41 := 0; fi41 < 15; fi41++ {
                    if (fi41 <= (-(fi41))) {
                        for fi42 := 0; fi42 < 20; fi42++ {
                            v38 = -128
                            v38 += v38
                            fmt.Println(fi41)
                            _ = fi42
                        }
                        v1 = v1
                    } else {
                        if v5 {
                            var v43 int32 = 702
                            v38 = (-(v38))
                            _ = v43
                        }
                    }
                    _ = fi41
                }
            }
            v44 := GenS1{F0: ("hello" + (("" + "foo bar") + "foo bar")), F1: 1}
            _ = v44
        }
        v5 = false
    }
    v45 := GenS5{F0: (^(v2)), F1: ("a" + "\n\t"), F2: ("\n\t" + ("hello" + "a"))}
    if (!(42 <= 563)) {
        v46 := GenS6{F0: 65535, F1: v38, F2: 1}
        var v47 int16 = v1
        _ = v46
        _ = v47
    } else {
        v48 := v3[(v2 >> ((4294967295) & 63))]
        v36 -= 0
        var v49 int8 = v38
        _ = v48
        _ = v49
    }
    v2 = (-(v2))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v45
    _ = v5
    _ = v6
    _ = v7
}
