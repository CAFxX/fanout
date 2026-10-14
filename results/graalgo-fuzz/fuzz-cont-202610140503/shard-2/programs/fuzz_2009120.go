package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 int16; F1 int; F2 string }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint16; F1 string; F2 bool }
type GenS7 struct { F0 bool; F1 int8; F2 uint32 }
type GenS8 struct { F0 uint }
func genFn0(p0 string) uint64 {
    v7 := []int{}
    var v8 uint32 = 977
    v8 = v8
    v9 := GenS1{F0: 2}
    v10 := GenS2{F0: -31, F1: -98}
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 18446744073709551615
}
func genFn1(p0 uint32) int {
    var v11 int = -9223372036854775808
    if true {
        fmt.Println(v11, p0)
    } else {
        var v12 uint16 = 32767
        if true {
            v13 := []int{(v11 >> ((1) & 63)), 1}
            var v14 uint = 0
            if false {
                v15 := GenS2{F0: -128, F1: 17}
                _ = v15
            } else {
                for fi16 := 0; fi16 < 6; fi16++ {
                    v17 := GenS1{F0: genFn0("hello")}
                    var v18 int32 = 42
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                if false {
                    fmt.Println(p0)
                    switch (v11 - 2147483647) {
                    case 2147483647:
                        v12 = 256
                    case 0:
                        v11 = (v11 >> ((v14) & 63))
                    case 32767:
                        var v19 uint = 18446744073709551615
                        _ = v19
                    default:
                        fmt.Println(v14, p0)
                    }
                    v20 := GenS3{F0: 1000, F1: (v11 % ((-2147483648) | 1)), F2: "foo bar"}
                    _ = v20
                } else {
                    var v21 float64 = 0.1
                    switch "\n\t" {
                    case "\n\t":
                        v22 := GenS4{F0: 255}
                        _ = v22
                    case "a":
                        var v23 int64 = 1
                        _ = v23
                    case "hello":
                        v14 = v14
                    }
                    _ = v21
                }
                switch v11 {
                case 127:
                    if (!true) {
                        v24 := GenS1{F0: 42}
                        for fi25 := 0; fi25 < 8; fi25++ {
                            switch (fi25 >> (((v14 >> (((v14 % ((v14) | 1))) & 63))) & 63)) {
                            case 188:
                                p0 = p0
                            case 2147483647:
                                switch "" {
                                case "a":
                                    v26 := GenS5{F0: 96}
                                    _ = v26
                                default:
                                    v14 += 255
                                }
                            case 0:
                                if true {
                                    var v27 uint64 = 670
                                    v28 := GenS5{F0: 55}
                                    v11 += (v11 - 507)
                                    _ = v27
                                    _ = v28
                                } else {
                                    _ = len(v13)
                                    var v29 int8 = 100
                                    _ = v29
                                }
                            default:
                                for fi30 := 0; fi30 < 15; fi30++ {
                                    if len(v13) > 0 {
                                        v13[0] = (fi30 << ((v14) & 63))
                                    }
                                    v31 := GenS4{F0: (v12 << (((v14 * v14)) & 63))}
                                    var v32 uint64 = 4294967295
                                    var v33 int8 = -128
                                    var v34 bool = true
                                    _ = fi30
                                    _ = v31
                                    _ = v32
                                    _ = v33
                                    _ = v34
                                }
                            }
                            if true {
                                for fi35 := 0; fi35 < 19; fi35++ {
                                    fmt.Println(v11, p0)
                                    var v36 int64 = 0
                                    fmt.Println(v11)
                                    _ = fi35
                                    _ = v36
                                }
                                fmt.Println(p0)
                            }
                            var v37 int64 = 0
                            _ = fi25
                            _ = v37
                        }
                        for fi38 := 0; fi38 < 10; fi38++ {
                            switch (fi38 & (fi38 | (-(fi38)))) {
                            case 127:
                                if (0 == 42) {
                                    fmt.Println(fi38)
                                    for fi39 := 0; fi39 < 17; fi39++ {
                                        if false {
                                            var v40 int16 = 1000
                                            if false {
                                                v41 := GenS4{F0: (v12 | v12)}
                                                _ = v41
                                            } else {
                                                var v42 float64 = 3.14159
                                                _ = v42
                                            }
                                            if len(v13) > 0 {
                                                v13[0] = (fi38 ^ 42)
                                            }
                                            _ = v40
                                        } else {
                                            _ = len(v13)
                                            var v43 float64 = (0.1 - (0.0 - 3.14159))
                                            _ = v43
                                        }
                                        _ = fi39
                                    }
                                }
                            case 42:
                                var v44 string = ("\n\t" + (("hello" + ("hello" + "foo bar")) + "foo bar"))
                                _ = v44
                            }
                            _ = fi38
                        }
                        _ = v24
                    }
                case -9223372036854775808:
                    v14 += (v14 << (((v14 | v14)) & 63))
                case -24:
                    var v45 int16 = -32768
                    _ = v45
                }
            }
            _ = v13
            _ = v14
        }
        _ = v12
    }
    for fi46 := 0; fi46 < 10; fi46++ {
        if (true && ((42 <= 0) || ((("" + "") + "foo bar") > "foo bar"))) {
            v47 := []int{}
            var v48 uint = 42
            _ = v47
            _ = v48
        }
        var v49 uint64 = 568
        _ = fi46
        _ = v49
    }
    fmt.Println(v11)
    if (93 <= p0) {
        var v50 float64 = (1.5 / 1.0)
        _ = v50
    }
    _ = p0
    _ = v11
    return (v11 & 9223372036854775807)
}

func main() {
    var v1 uint64 = 172
    var v2 float32 = 1.5
    var v3 int32 = -1
    v4 := make(map[string]bool)
    v5 := make(map[int]int)
    var v6 uint32 = 2147483647
    v1 = v1
    fmt.Println(v6, v3)
    for fi51 := 0; fi51 < 17; fi51++ {
        for fi52 := 0; fi52 < 14; fi52++ {
            v1 = (v1 ^ 4294967295)
            _ = fi52
        }
        var v53 uint64 = (v1 & 915)
        v54 := GenS5{F0: 255}
        _ = fi51
        _ = v53
        _ = v54
    }
    var v55 string = ""
    if false {
        v1 = v1
    }
    for fi56 := 0; fi56 < 3; fi56++ {
        if ((!true) && true) {
            if true {
                switch v55 {
                case "foo bar":
                    v2 -= 0.0
                }
                v6 = (v6 ^ v6)
            } else {
                if (!true) {
                    fmt.Println(v55)
                    var v57 int64 = -55
                    _ = v57
                } else {
                    v2 += ((float32(((0.1 * ((-2.25 / (((0.1 - 3.14159) * 3.14159) / 1.0)) - ((-2.25 + 1.5) - 0.0))) / 1.0)) + (((v2 - v2) + v2) + 1.5)) + 1e10)
                    var v58 uint16 = 32767
                    _ = v58
                }
                v2 += (float32((1.5 / 1e10)) + v2)
                _ = len(v5)
            }
            switch (fi56 & (fi56 ^ (fi56 | genFn1((v6 / (((^(v6))) | 1)))))) {
            case 4294967295:
                var v59 string = (((v55 + v55) + "") + (("\n\t" + v55) + ""))
                _ = v59
            case -642:
                var v60 int64 = 645
                _ = v60
            case 32767:
                if false {
                    delete(v5, 128)
                    for fi61 := 0; fi61 < 1; fi61++ {
                        var v62 uint32 = v6
                        _ = fi61
                        _ = v62
                    }
                    if true {
                        v63 := []int{}
                        _ = v63
                    }
                } else {
                    v64 := GenS5{F0: 2}
                    _ = v64
                }
            default:
                switch (v55 + ((((v55 + v55) + v55) + v55) + v55)) {
                case "hello":
                    fmt.Println(v2)
                case "a":
                    v65 := []int{fi56}
                    _ = v65
                default:
                    for fi66 := 0; fi66 < 5; fi66++ {
                        var v67 uint64 = 18446744073709551615
                        v68 := GenS6{F0: 1, F1: v55, F2: (!(-450 == 988))}
                        _ = fi66
                        _ = v67
                        _ = v68
                    }
                }
            }
        }
        var v69 string = "a"
        _ = fi56
        _ = v69
    }
    v70 := GenS1{F0: (^(v1))}
    var v71 int8 = 42
    switch 128 {
    case -9223372036854775808:
        v71 -= 42
    case 9223372036854775807:
        v72 := GenS7{F0: (!false), F1: 127, F2: (v6 & v6)}
        _ = v72
    }
    var v73 int32 = v3
    if (float32(((float64(v2) / 1.5) * (float64(v2) * 0.1))) < 3.14159) {
        v55 = v55
    }
    if (!(!true)) {
        for fi74 := 0; fi74 < 1; fi74++ {
            v75 := GenS8{F0: 4294967295}
            _ = fi74
            _ = v75
        }
    }
    v71 -= v71
    v1 += (v1 ^ v1)
    switch 164 {
    case -46:
        v55 = v55
    }
    v6 = v6
    v76 := []int{}
    var v77 int16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    fmt.Println("v71:", v71)
    fmt.Println("v73:", v73)
    fmt.Println("slicelen v76:", len(v76), v76)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v6
    _ = v70
    _ = v71
    _ = v73
    _ = v76
    _ = v77
}
