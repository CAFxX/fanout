package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 string; F1 uint16; F2 uint }
type GenS3 struct { F0 int16; F1 uint; F2 int64 }
type GenS4 struct { F0 int; F1 int8; F2 int32 }
type GenS5 struct { F0 int64; F1 uint8 }
type GenS6 struct { F0 int32; F1 int32; F2 string }
type GenS7 struct { F0 uint32; F1 uint64 }
type GenS8 struct { F0 int64; F1 uint32 }
type GenS9 struct { F0 bool; F1 int64; F2 uint }

func main() {
    v1 := []int{32768, 32768, 32768}
    var v2 int64 = 260
    v3 := make(map[string]bool)
    fmt.Println(v2)
    v2 = v2
    for fi4 := 0; fi4 < 8; fi4++ {
        if true {
            v5 := v3["hello"]
            for fi6 := 0; fi6 < 10; fi6++ {
                v2 = v2
                var v7 uint64 = 1
                _ = fi6
                _ = v7
            }
            _ = v5
        } else {
            switch (fi4 % (((fi4 / ((fi4) | 1))) | 1)) {
            case 0:
                fmt.Println(fi4, fi4)
            case 8:
                for fi8 := 0; fi8 < 19; fi8++ {
                    if (0 > 1) {
                        fmt.Println(fi4, fi4)
                        var v9 uint = 443
                        fmt.Println(fi4, fi4)
                        v3[""] = true
                        v10 := GenS1{F0: -1, F1: 255}
                        _ = v10
                        _ = v9
                    }
                    var v11 uint8 = 127
                    if (1000 != 42) {
                        var v12 int = fi8
                        switch "a" {
                        case "":
                            v13 := GenS2{F0: ("foo bar" + (("foo bar" + (("\n\t" + "") + "foo bar")) + "")), F1: 2, F2: 633}
                            _ = v13
                        default:
                            var v14 uint8 = (v11 / ((v11) | 1))
                            _ = v14
                        }
                        var v15 int64 = v2
                        if (!((0.1 * (1.5 - float64(0.1))) < 0.1)) {
                            switch (fi4 * fi4) {
                            case 2147483647:
                                v11 += v11
                            case 42:
                                var v16 uint8 = (v11 - 34)
                                _ = v16
                            case 4294967295:
                                var v17 uint = 42
                                _ = v17
                            }
                            var v18 int8 = 127
                            _ = v18
                        }
                        v3[""] = ("a" < (("hello" + ("a" + "a")) + "foo bar"))
                        _ = v12
                        _ = v15
                    } else {
                        if false {
                            delete(v3, ("hello" + ("hello" + "a")))
                            if true {
                                v11 = (v11 ^ (v11 * v11))
                                v11 = v11
                            }
                            v2 = v2
                        } else {
                            if true {
                                switch -2147483648 {
                                case 32768:
                                    if (true && (-594 > 1)) {
                                        v19 := GenS2{F0: "a", F1: 299, F2: 255}
                                        v20 := GenS3{F0: -32, F1: 255, F2: int64(fi8)}
                                        var v21 uint32 = 65535
                                        _ = v19
                                        _ = v20
                                        _ = v21
                                    }
                                case 4294967295:
                                    var v22 uint = 65535
                                    _ = v22
                                case 32767:
                                    var v23 int16 = -34
                                    _ = v23
                                default:
                                    var v24 bool = (true && false)
                                    _ = v24
                                }
                                v25 := GenS3{F0: 1000, F1: 546, F2: v2}
                                v26 := v3[""]
                                _ = v25
                                _ = v26
                            } else {
                                v2 = v2
                            }
                        }
                        var v27 uint16 = 675
                        v28 := GenS4{F0: (-(fi8)), F1: 100, F2: int32(v27)}
                        _ = v27
                        _ = v28
                    }
                    _ = fi8
                    _ = v11
                }
            case 32767:
                for fi29 := 0; fi29 < 7; fi29++ {
                    v2 += v2
                    if false {
                        var v30 int16 = -139
                        v2 = 395
                        v31 := GenS5{F0: v2, F1: 42}
                        _ = v30
                        _ = v31
                    } else {
                        fmt.Println(fi29, fi29)
                        var v32 int32 = 1
                        _ = v32
                    }
                    v2 = (v2 - 0)
                    v33 := GenS6{F0: -2147483648, F1: 0, F2: ("a" + ("hello" + ("foo bar" + "\n\t")))}
                    _ = fi29
                    _ = v33
                }
            }
            v2 = (v2 - (v2 >> ((1) & 63)))
            for fi34 := 0; fi34 < 13; fi34++ {
                fmt.Println(v2, fi4, fi34)
                switch 128 {
                case 860:
                    if ((fi34 / (((^(fi34))) | 1)) == fi34) {
                        v35 := GenS6{F0: 42, F1: 1, F2: ((("\n\t" + "hello") + "a") + "a")}
                        v2 += (v2 >> ((2) & 63))
                        _ = v35
                    } else {
                        var v36 float32 = 1e10
                        var v37 int32 = -860
                        _ = v36
                        _ = v37
                    }
                default:
                    v2 += (^(v2))
                }
                _ = fi34
            }
        }
        _ = fi4
    }
    var v38 bool = (535 < 679)
    var v39 uint16 = 32767
    for fi40 := 0; fi40 < 18; fi40++ {
        var v41 uint32 = 842
        var v42 int32 = 448
        _ = fi40
        _ = v41
        _ = v42
    }
    switch ("\n\t" + "foo bar") {
    case "":
        v39 = 635
    case "\n\t":
        var v43 float32 = ((((0.0 + 0.0) + 0.0) - float32(3.14159)) / -2.25)
        _ = v43
    }
    v2 -= (v2 + (v2 | 778))
    fmt.Println(v39, v38)
    if (float32(((1e10 - float64(0.1)) + 1.5)) >= 0.0) {
        var v44 float32 = 1e10
        v38 = v38
        var v45 int32 = 678
        _ = v44
        _ = v45
    }
    var v46 int64 = (v2 / ((v2) | 1))
    switch "" {
    case "\n\t":
        v47 := GenS7{F0: 479, F1: 357}
        _ = v47
    case "hello":
        v2 -= v46
    case "foo bar":
        for fi48 := 0; fi48 < 17; fi48++ {
            v49 := GenS7{F0: 2, F1: 2}
            var v50 uint = 18446744073709551615
            _ = fi48
            _ = v49
            _ = v50
        }
    }
    v51 := GenS8{F0: v46, F1: 65535}
    if ((v39 ^ 617) >= (^(v39))) {
        var v52 float32 = -2.25
        v46 = v2
        if v38 {
            if (-88 != -109) {
                v53 := GenS6{F0: -2147483648, F1: -2147483648, F2: "\n\t"}
                _ = v53
            } else {
                fmt.Println(v52)
                v52 += (1.5 + 0.0)
                switch (("a" + "hello") + "foo bar") {
                case "a":
                    v54 := GenS3{F0: 42, F1: 4294967295, F2: v46}
                    _ = v54
                }
            }
        } else {
            switch 2147483647 {
            case 2147483647:
                v2 -= 9223372036854775807
            case 128:
                var v55 int64 = v2
                _ = v55
            }
            v56 := GenS6{F0: -778, F1: -558, F2: "\n\t"}
            _ = v56
        }
        _ = v52
    } else {
        if ((^(v39)) >= (v39 | v39)) {
            v2 += v46
            switch ("foo bar" + "\n\t") {
            case "":
                v38 = (!(!v38))
            case "\n\t":
                v57 := v3[("\n\t" + ("\n\t" + "a"))]
                _ = v57
            default:
                v58 := GenS9{F0: (!(v38 && v38)), F1: 622, F2: 2}
                _ = v58
            }
        }
    }
    var v59 int32 = -422
    v59 += v59
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v46:", v46)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v46
    _ = v51
    _ = v59
}
