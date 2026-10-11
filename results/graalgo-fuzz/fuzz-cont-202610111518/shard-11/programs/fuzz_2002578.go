package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 uint16; F1 int8 }
type GenS5 struct { F0 int; F1 int8 }
type GenS6 struct { F0 int; F1 uint; F2 int32 }
type GenS7 struct { F0 uint16 }
func genFn0(p0 uint32) uint32 {
    v7 := []int{}
    v7 = append(v7, 1)
    v8 := GenS1{F0: 1000, F1: 32768}
    _ = p0
    _ = v7
    _ = v8
    return 995
}

func main() {
    var v1 uint = 65535
    var v2 uint64 = 4294967295
    var v3 uint16 = 65535
    v4 := make(map[int]string)
    var v5 uint = v1
    v6 := []string{}
    if (-199 >= 796) {
        var v9 int8 = 42
        delete(v4, 9223372036854775807)
        _ = v9
    } else {
        fmt.Println(v5, v2)
        var v10 float64 = (((-2.25 / (0.1 + (float64(3.14159) + 1e10))) + float64(((0.1 - 1.5) * -2.25))) - -2.25)
        _ = v10
    }
    v5 += (v5 / ((446) | 1))
    for fi11 := 0; fi11 < 5; fi11++ {
        _ = len(v4)
        v1 = uint(v2)
        fmt.Println(v2)
        _ = fi11
    }
    switch 0 {
    case 892:
        fmt.Println(v2, v1, v3)
    case -2147483648:
        if false {
            delete(v4, -2147483648)
            if false {
                var v12 uint64 = 4294967295
                var v13 float32 = -2.25
                v12 = 1
                _ = v12
                _ = v13
            } else {
                var v14 float64 = (1e10 + -2.25)
                for fi15 := 0; fi15 < 10; fi15++ {
                    _ = len(v4)
                    for fi16 := 0; fi16 < 6; fi16++ {
                        _ = len(v4)
                        _ = fi16
                    }
                    v17 := GenS1{F0: -1, F1: fi15}
                    _ = fi15
                    _ = v17
                }
                _ = v14
            }
            var v18 float32 = 3.14159
            _ = v18
        } else {
            if (0 == -406) {
                v19 := GenS2{F0: 150, F1: 127}
                _ = v19
            } else {
                if true {
                    var v20 int32 = 726
                    _ = v20
                } else {
                    switch "hello" {
                    case "hello":
                        v1 = v1
                    default:
                        v21 := GenS2{F0: 1, F1: -85}
                        _ = v21
                    }
                }
            }
            if true {
                v22 := GenS1{F0: -32768, F1: 127}
                if (-128 < 0) {
                    if false {
                        if true {
                            for fi23 := 0; fi23 < 19; fi23++ {
                                v3 += (^(v3))
                                var v24 uint32 = 268
                                _ = fi23
                                _ = v24
                            }
                        }
                        switch "" {
                        case "foo bar":
                            v5 += (^(v1))
                        default:
                            v25 := GenS2{F0: 593, F1: 42}
                            _ = v25
                        }
                    } else {
                        fmt.Println(v2, v1)
                    }
                }
                if false {
                    if len(v6) > 0 {
                        v6[0] = "foo bar"
                    }
                    if len(v6) > 0 {
                        v6[0] = "foo bar"
                    }
                    var v26 float32 = (float32(0.0) + 0.1)
                    _ = v26
                }
                _ = v22
            } else {
                v27 := GenS3{F0: 0, F1: (v1 << (((v5 ^ v5)) & 63))}
                v5 -= (v5 << ((2) & 63))
                _ = v27
            }
            _ = len(v4)
        }
    }
    var v28 uint64 = (v2 | (v2 - 18446744073709551615))
    var v29 int = 186
    fmt.Println(v28, v1)
    fmt.Println(v28, v3, v1)
    v30 := GenS2{F0: 0, F1: 1}
    v3 = v3
    if false {
        for fi31 := 0; fi31 < 16; fi31++ {
            v3 -= (v3 % (((^(v3))) | 1))
            fmt.Println(v1)
            if true {
                var v32 int64 = 1
                _ = v32
            }
            _ = fi31
        }
        if true {
            var v33 uint = (v1 << ((0) & 63))
            v3 += (v3 % ((0) | 1))
            v6 = append(v6, (("\n\t" + "") + "hello"))
            _ = v33
        }
        var v34 int8 = 38
        _ = v34
    } else {
        switch ("a" + "\n\t") {
        case "hello":
            for fi35 := 0; fi35 < 3; fi35++ {
                var v36 uint = v5
                v37 := GenS4{F0: 42, F1: 100}
                var v38 uint64 = v28
                _ = fi35
                _ = v36
                _ = v37
                _ = v38
            }
        default:
            var v39 int = v29
            _ = v39
        }
        if false {
            v4[1] = ("a" + "a")
            fmt.Println(v2)
        } else {
            v1 = v1
            v40 := GenS5{F0: (v29 + (v29 >> ((65535) & 63))), F1: -17}
            _ = v40
        }
    }
    fmt.Println(v3, v2)
    var v41 bool = (!false)
    if (true || v41) {
        if (1e10 < (0.0 * 1e10)) {
            v42 := GenS6{F0: (v29 & (v29 | v29)), F1: (v1 * uint(v28)), F2: 2147483647}
            fmt.Println(v2, v3, v41)
            for fi43 := 0; fi43 < 12; fi43++ {
                switch ("a" + "\n\t") {
                case "\n\t":
                    for fi44 := 0; fi44 < 4; fi44++ {
                        v45 := GenS2{F0: genFn0(519), F1: 22}
                        v28 = (v2 - (v2 | 18446744073709551615))
                        _ = fi44
                        _ = v45
                    }
                case "foo bar":
                    if len(v6) > 0 {
                        v6[0] = ("foo bar" + "hello")
                    }
                case "":
                    for fi46 := 0; fi46 < 18; fi46++ {
                        v47 := GenS7{F0: v3}
                        switch fi46 {
                        case 1:
                            if (((("" + ("" + "hello")) + "a") + "") < "foo bar") {
                                var v48 string = "a"
                                _ = v48
                            }
                        default:
                            v29 += (fi43 << (((v1 & 2)) & 63))
                        }
                        _ = fi46
                        _ = v47
                    }
                }
                var v49 uint64 = v28
                v6 = append(v6, (("\n\t" + "") + ("foo bar" + "a")))
                _ = fi43
                _ = v49
            }
            _ = v42
        }
        v6 = append(v6, "\n\t")
    }
    switch (v29 | v29) {
    case 32768:
        v5 = (v5 % ((v1) | 1))
    case -2147483648:
        var v50 uint16 = (v3 % (((v3 * v3)) | 1))
        _ = v50
    }
    switch (v29 / ((-1) | 1)) {
    case 42:
        for fi51 := 0; fi51 < 15; fi51++ {
            var v52 uint32 = 42
            v3 += uint16(v1)
            fmt.Println(v3, v29)
            _ = fi51
            _ = v52
        }
    case 128:
        fmt.Println(v3, v2)
    default:
        var v53 uint = (v1 * v1)
        _ = v53
    }
    var v54 uint32 = 979
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v41
    _ = v5
    _ = v54
    _ = v6
}
