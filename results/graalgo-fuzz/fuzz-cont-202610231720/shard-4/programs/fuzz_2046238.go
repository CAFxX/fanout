package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 uint32 }
type GenS2 struct { F0 uint32; F1 uint16; F2 int }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 uint32; F1 string; F2 int64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int16; F1 int }
func genFn0() uint16 {
    v5 := GenS1{F0: 32767, F1: 1, F2: 0}
    var v6 int64 = -1
    if (!false) {
        var v7 int64 = (v6 - v6)
        if false {
            v8 := []int{}
            v7 -= 0
            v6 += v6
            _ = v8
        } else {
            if (!((!true) || true)) {
                for fi9 := 0; fi9 < 2; fi9++ {
                    if false {
                        v6 += v7
                        if (!(32767 > 1000)) {
                            v6 = v6
                            v10 := GenS1{F0: 497, F1: 220, F2: 42}
                            _ = v10
                        } else {
                            if (-2.25 > -2.25) {
                                v11 := GenS1{F0: 2, F1: 42, F2: 65535}
                                _ = v11
                            } else {
                                var v12 int8 = 42
                                v12 += v12
                                v13 := GenS2{F0: 1, F1: 284, F2: 2147483647}
                                _ = v12
                                _ = v13
                            }
                        }
                    }
                    fmt.Println(v7, v7)
                    if (255 > 127) {
                        var v14 uint64 = 18446744073709551615
                        v15 := GenS2{F0: 0, F1: 65535, F2: 0}
                        fmt.Println(v7, v7)
                        _ = v14
                        _ = v15
                    } else {
                        for fi16 := 0; fi16 < 11; fi16++ {
                            v17 := []int{(fi16 | fi9), fi16}
                            _ = fi16
                            _ = v17
                        }
                        for fi18 := 0; fi18 < 13; fi18++ {
                            v19 := GenS3{F0: v7, F1: 42}
                            _ = fi18
                            _ = v19
                        }
                    }
                    _ = fi9
                }
                if (42 >= 484) {
                    for fi20 := 0; fi20 < 20; fi20++ {
                        var v21 int = fi20
                        v22 := []int{(^(v21))}
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                    if true {
                        v23 := GenS3{F0: v6, F1: 42}
                        _ = v23
                    }
                } else {
                    if false {
                        fmt.Println(v7)
                        switch "" {
                        case "foo bar":
                            var v24 uint32 = 56
                            _ = v24
                        case "a":
                            var v25 uint16 = 68
                            _ = v25
                        case "\n\t":
                            var v26 float32 = (0.0 - 3.14159)
                            _ = v26
                        }
                        var v27 int8 = 0
                        _ = v27
                    }
                }
            }
            v7 = (v7 & -1)
            var v28 int32 = 42
            _ = v28
        }
        _ = v7
    } else {
        v6 = 1
        if (!(-1 >= v6)) {
            v29 := []int{-1}
            var v30 int = 9223372036854775807
            _ = v29
            _ = v30
        }
        var v31 float64 = 1e10
        _ = v31
    }
    v6 -= (-(v6))
    _ = v5
    _ = v6
    return 65535
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]uint32)
    var v4 bool = true
    var v32 uint16 = 141
    var v33 int16 = 1000
    v4 = false
    v34 := GenS4{F0: 42, F1: ("\n\t" + "foo bar"), F2: 9223372036854775807}
    var v35 int64 = 1
    v3[590] = 2147483648
    delete(v1, 2)
    fmt.Println(v33, v35, v4)
    if v4 {
        if true {
            v35 += (v35 | v35)
            switch -669 {
            case 42:
                v32 = 248
            case 2147483647:
                v36 := GenS5{F0: v33}
                _ = v36
            case 0:
                switch "hello" {
                case "\n\t":
                    v37 := GenS4{F0: uint32(v32), F1: "", F2: (v35 * v35)}
                    _ = v37
                }
            default:
                delete(v3, 4294967295)
            }
            v38 := []int{9223372036854775807, -81, 2147483647}
            _ = v38
        } else {
            if ((-(v35)) == (v35 * -9223372036854775808)) {
                v39 := []int{128, 32768}
                if len(v39) > 0 {
                    v39[0] = 1
                }
                _ = v39
            }
            for fi40 := 0; fi40 < 6; fi40++ {
                var v41 string = ""
                if false {
                    v35 = v35
                    var v42 uint16 = v32
                    _ = v42
                }
                v43 := GenS2{F0: 2147483648, F1: v32, F2: fi40}
                _ = fi40
                _ = v41
                _ = v43
            }
        }
        delete(v2, 0)
    }
    _ = len(v3)
    v44 := []int{}
    for fi45 := 0; fi45 < 14; fi45++ {
        v32 = (^(v32))
        _ = fi45
    }
    for fi46 := 0; fi46 < 4; fi46++ {
        switch (fi46 ^ (fi46 >> ((0) & 63))) {
        case 859:
            v4 = v4
        case 42:
            fmt.Println(fi46)
        case 9:
            var v47 float64 = float64(1e10)
            _ = v47
        default:
            v33 -= (v33 & v33)
        }
        var v48 int8 = 22
        _ = fi46
        _ = v48
    }
    fmt.Println(v35)
    if true {
        v32 += (^(v32))
    } else {
        v32 += genFn0()
        for fi49 := 0; fi49 < 1; fi49++ {
            var v50 float32 = 0.0
            for fi51 := 0; fi51 < 18; fi51++ {
                v52 := GenS2{F0: 1, F1: (v32 & 65535), F2: int(v32)}
                fmt.Println(fi51, v35, fi51)
                _ = fi51
                _ = v52
            }
            var v53 uint8 = 159
            _ = fi49
            _ = v50
            _ = v53
        }
        var v54 int = -209
        _ = v54
    }
    v55 := GenS3{F0: 42, F1: 65535}
    v56 := GenS6{F0: 443, F1: -819}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v44:", len(v44), v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v44
    _ = v55
    _ = v56
}
