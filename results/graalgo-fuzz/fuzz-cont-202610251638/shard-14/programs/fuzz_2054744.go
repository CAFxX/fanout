package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 bool; F2 int }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 string }
func genFn0(p0 int16) string {
    if ("hello" >= "hello") {
        if (true || (65535 == 883)) {
            for fi5 := 0; fi5 < 9; fi5++ {
                for fi6 := 0; fi6 < 20; fi6++ {
                    var v7 int = (^(fi5))
                    v7 = fi5
                    var v8 uint = 0
                    _ = fi6
                    _ = v7
                    _ = v8
                }
                v9 := GenS1{F0: 0, F1: -898, F2: -743}
                v10 := GenS2{F0: 0}
                _ = fi5
                _ = v10
                _ = v9
            }
            v11 := []int{}
            fmt.Println(p0)
            _ = v11
        } else {
            var v12 uint64 = 18446744073709551615
            _ = v12
        }
    }
    var v13 string = ("\n\t" + ((("foo bar" + "foo bar") + "\n\t") + ("" + "\n\t")))
    if (128 == 42) {
        v13 = v13
    }
    v14 := GenS3{F0: ("foo bar" + ((v13 + "\n\t") + v13)), F1: ((-(p0)) != (p0 - p0)), F2: 128}
    _ = p0
    _ = v13
    _ = v14
    return "a"
}

func main() {
    var v1 bool = (!true)
    var v2 int32 = 42
    var v3 int = -2147483648
    var v4 uint32 = 42
    v15 := []int{(^(v3)), (v3 % (((v3 % ((v3) | 1))) | 1)), v3}
    fmt.Println(v3)
    if len(v15) > 0 {
        v15[0] = v3
    }
    var v16 uint64 = 0
    v2 = (-(v2))
    if v1 {
        v15 = append(v15, -9223372036854775808)
        v17 := GenS4{F0: v4}
        for fi18 := 0; fi18 < 17; fi18++ {
            v19 := GenS4{F0: v4}
            _ = fi18
            _ = v19
        }
        _ = v17
    } else {
        if v1 {
            switch v3 {
            case 265:
                v3 += (v3 % (((v3 >> ((531) & 63))) | 1))
            case 128:
                if (v1 && v1) {
                    v3 += -1
                    _ = len(v15)
                } else {
                    var v20 uint32 = (v4 * v4)
                    if (943 == v16) {
                        var v21 uint32 = (v4 | 4294967295)
                        _ = v21
                    }
                    _ = v20
                }
            case 186:
                v3 = v3
            default:
                switch "\n\t" {
                case "\n\t":
                    if (42 < (v16 ^ (v16 * v16))) {
                        var v22 uint16 = 42
                        for fi23 := 0; fi23 < 9; fi23++ {
                            v2 = -730
                            if len(v15) >= 2 {
                                v15 = v15[1:len(v15)]
                            }
                            _ = fi23
                        }
                        _ = v22
                    }
                case "hello":
                    fmt.Println(v4)
                case "foo bar":
                    v24 := GenS1{F0: 1, F1: (v2 << ((255) & 63)), F2: 42}
                    _ = v24
                default:
                    var v25 int = (-(v3))
                    _ = v25
                }
            }
            for fi26 := 0; fi26 < 9; fi26++ {
                v27 := GenS1{F0: 65535, F1: (v2 >> ((4294967295) & 63)), F2: 1000}
                _ = fi26
                _ = v27
            }
            switch (^(v3)) {
            case 127:
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
            }
        }
    }
    for fi28 := 0; fi28 < 17; fi28++ {
        var v29 int64 = int64(fi28)
        _ = fi28
        _ = v29
    }
    for fi30 := 0; fi30 < 9; fi30++ {
        var v31 uint = 18446744073709551615
        if ((v31 >> ((v31) & 63)) != 852) {
            for fi32 := 0; fi32 < 6; fi32++ {
                v33 := GenS3{F0: "a", F1: v1, F2: 127}
                if (342 == 1) {
                    if (2 < 227) {
                        v3 -= fi32
                        var v34 string = "foo bar"
                        _ = v34
                    } else {
                        switch 0 {
                        case 128:
                            var v35 uint = (v31 >> ((2) & 63))
                            _ = v35
                        case 4294967295:
                            if true {
                                v36 := GenS2{F0: 0}
                                var v37 string = "hello"
                                v3 = (fi32 << ((uint(v16)) & 63))
                                _ = v36
                                _ = v37
                            }
                        }
                        v38 := GenS5{F0: ""}
                        _ = v38
                    }
                } else {
                    if v1 {
                        switch fi30 {
                        case -1:
                            var v39 int32 = (v2 / ((v2) | 1))
                            var v40 uint8 = 1
                            _ = v39
                            _ = v40
                        case 4294967295:
                            var v41 int8 = 1
                            _ = v41
                        case 32767:
                            fmt.Println(v2)
                        default:
                            if v1 {
                                fmt.Println(v1, v3, v31)
                                if len(v15) >= 2 {
                                    v15 = v15[1:len(v15)]
                                }
                            }
                        }
                        v42 := GenS4{F0: (v4 / (((^(v4))) | 1))}
                        v16 = v16
                        _ = v42
                    }
                    v43 := GenS2{F0: (v4 << ((uint(v4)) & 63))}
                    v4 = 2147483648
                    var v44 int64 = 9223372036854775807
                    _ = v43
                    _ = v44
                }
                _ = fi32
                _ = v33
            }
            var v45 uint = v31
            _ = v45
        } else {
            v46 := GenS2{F0: (v4 % ((v4) | 1))}
            var v47 int = (fi30 & 32768)
            _ = v46
            _ = v47
        }
        switch "a" {
        case "hello":
            for fi48 := 0; fi48 < 3; fi48++ {
                var v49 string = "hello"
                if false {
                    switch (v3 & 128) {
                    case 42:
                        if true {
                            v2 = -725
                        }
                    default:
                        var v50 int64 = -1
                        _ = v50
                    }
                } else {
                    var v51 float64 = 3.14159
                    _ = v51
                }
                _ = fi48
                _ = v49
            }
        default:
            var v52 uint64 = (v16 & v16)
            _ = v52
        }
        _ = fi30
        _ = v31
    }
    v53 := GenS2{F0: (v4 - v4)}
    v54 := GenS3{F0: "foo bar", F1: v1, F2: v3}
    v15 = append(v15, 277)
    v55 := GenS1{F0: 61, F1: v2, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v53
    _ = v54
    _ = v55
}
