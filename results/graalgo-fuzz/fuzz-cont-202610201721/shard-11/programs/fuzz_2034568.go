package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 bool; F2 uint16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 string; F1 int16 }
type GenS6 struct { F0 uint64; F1 bool }
type GenS7 struct { F0 int16; F1 uint8 }

func main() {
    v1 := []string{}
    var v2 float64 = 1.5
    v3 := make(map[int]bool)
    var v4 uint64 = 0
    var v5 uint = 4294967295
    v6 := []string{"hello"}
    if ((1.5 < (1e10 - 0.0)) && (int64(v5) <= 1)) {
        v7 := GenS1{F0: 623}
        for fi8 := 0; fi8 < 13; fi8++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            fmt.Println(v2, v2, fi8)
            var v9 uint32 = 42
            _ = fi8
            _ = v9
        }
        switch 345 {
        case -2147483648:
            v10 := GenS1{F0: 65535}
            _ = v10
        case -1:
            if true {
                v11 := GenS2{F0: -2147483648}
                v12 := GenS3{F0: 30, F1: (-9223372036854775808 != 1), F2: 255}
                _ = v11
                _ = v12
            }
        case 126:
            var v13 uint = (v5 * (^(v5)))
            _ = v13
        }
        _ = v7
    } else {
        var v14 uint64 = (v4 * v4)
        for fi15 := 0; fi15 < 2; fi15++ {
            v14 = 42
            _ = fi15
        }
        if len(v1) > 0 {
            v1[0] = "hello"
        }
        _ = v14
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        var v17 uint32 = 539
        if (true && ((-2.25 * -2.25) == 0.1)) {
            for fi18 := 0; fi18 < 19; fi18++ {
                if len(v6) > 0 {
                    v6[0] = "\n\t"
                }
                var v19 int64 = -9223372036854775808
                var v20 int8 = 100
                _ = fi18
                _ = v19
                _ = v20
            }
            v17 -= v17
            v21 := GenS4{F0: 765}
            _ = v21
        }
        _ = fi16
        _ = v17
    }
    fmt.Println(v5, v4)
    for fi22 := 0; fi22 < 2; fi22++ {
        if true {
            for fi23 := 0; fi23 < 15; fi23++ {
                var v24 float64 = (1.5 + v2)
                v5 += v5
                if false {
                    var v25 float32 = 3.14159
                    switch "a" {
                    case "\n\t":
                        if (1 == -54) {
                            var v26 int64 = 1
                            v4 += uint64(v5)
                            _ = v26
                        }
                    default:
                        fmt.Println(v4)
                    }
                    _ = v25
                } else {
                    switch (-(fi23)) {
                    case 42:
                        var v27 uint = (v5 - v5)
                        _ = v27
                    case 295:
                        if (("\n\t" + "foo bar") > ("hello" + "foo bar")) {
                            v5 = (v5 * 1)
                            var v28 uint64 = (v4 - v4)
                            _ = v28
                        }
                        v24 -= v2
                        switch "" {
                        case "foo bar":
                            var v29 int32 = -1
                            _ = v29
                        default:
                            switch (fi23 + (fi22 + 580)) {
                            case 32767:
                                if false {
                                    switch "" {
                                    case "a":
                                        v30 := GenS1{F0: 389}
                                        _ = v30
                                    case "foo bar":
                                        v5 = (v5 << (((v5 * (v5 ^ 42))) & 63))
                                    case "\n\t":
                                        v2 -= -2.25
                                    }
                                    if len(v6) > 0 {
                                        v6[0] = "a"
                                    }
                                    fmt.Println(v24)
                                } else {
                                    var v31 int32 = 0
                                    v32 := GenS5{F0: "\n\t", F1: 0}
                                    if (false || ("foo bar" != ("" + "hello"))) {
                                        v31 = v31
                                        var v33 uint16 = 33
                                        _ = v33
                                    }
                                    _ = v31
                                    _ = v32
                                }
                            case 9223372036854775807:
                                var v34 int8 = 70
                                _ = v34
                            case -9223372036854775808:
                                v35 := GenS6{F0: 2, F1: false}
                                _ = v35
                            default:
                                v36 := GenS1{F0: 65535}
                                _ = v36
                            }
                            v6 = append(v6, ("\n\t" + ("" + "foo bar")))
                            fmt.Println(v24, fi22, v24)
                        }
                    }
                }
                v24 += (v2 * 1.5)
                v37 := GenS7{F0: -174, F1: 61}
                _ = fi23
                _ = v24
                _ = v37
            }
        }
        var v38 uint64 = 72
        v39 := v3[(fi22 / ((fi22) | 1))]
        _ = fi22
        _ = v38
        _ = v39
    }
    if false {
        v40 := GenS5{F0: ("" + ("a" + ("hello" + "a"))), F1: 1000}
        _ = v40
    } else {
        var v41 bool = true
        v5 = v5
        _ = v41
    }
    var v42 int32 = 42
    v42 = 42
    v5 = (v5 + v5)
    v42 = v42
    v6 = append(v6, "")
    if (!(65535 == 1)) {
        fmt.Println(v4, v4)
        v6 = append(v6, "foo bar")
    } else {
        if false {
            v43 := GenS5{F0: "\n\t", F1: 1}
            var v44 int32 = (v42 & 219)
            _ = v43
            _ = v44
        }
        if ((v42 << (((v5 ^ 1)) & 63)) >= (v42 | v42)) {
            for fi45 := 0; fi45 < 4; fi45++ {
                var v46 int64 = 0
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                var v47 float64 = (v2 - 0.1)
                _ = fi45
                _ = v46
                _ = v47
            }
            v48 := GenS2{F0: 42}
            _ = v48
        }
        v4 += (v4 | (v4 % ((599) | 1)))
    }
    var v49 uint64 = (^(v4))
    switch 32767 {
    case 874:
        var v50 uint32 = 2147483647
        _ = v50
    case -9223372036854775808:
        v49 += 2
    case -2147483648:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v49
    _ = v5
    _ = v6
}
