package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int32 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 int16; F1 int8; F2 int8 }
func genFn0(p0 string, p1 uint) int16 {
    var v7 int64 = 478
    if (0 > 2) {
        v8 := GenS1{F0: 2147483647, F1: 2, F2: int32(v7)}
        _ = v8
    }
    var v9 int8 = -126
    _ = p0
    _ = p1
    _ = v7
    _ = v9
    return 32767
}
func genFn1(p0 uint, p1 int) uint64 {
    v10 := []int{p1, p1}
    if false {
        if (0 < 0) {
            if (false && (!(2 >= 1))) {
                var v11 uint64 = 2
                var v12 int = (p1 >> (((p0 / ((uint(v11)) | 1))) & 63))
                switch v12 {
                case 2147483647:
                    var v13 int8 = 0
                    _ = v13
                case -9223372036854775808:
                    for fi14 := 0; fi14 < 15; fi14++ {
                        var v15 uint16 = 563
                        for fi16 := 0; fi16 < 8; fi16++ {
                            switch "hello" {
                            case "\n\t":
                                p0 += (^(p0))
                            }
                            if (!true) {
                                var v17 float64 = 1e10
                                _ = v17
                            } else {
                                switch (fi16 >> ((p0) & 63)) {
                                case -9223372036854775808:
                                    if (2147483647 <= 2147483647) {
                                        v15 += v15
                                        v18 := GenS1{F0: (fi16 >> ((p0) & 63)), F1: v15, F2: 560}
                                        if false {
                                            if ((uint8(v11) >= 255) || (242 > 127)) {
                                                fmt.Println(v15, fi16)
                                                v12 = (p1 % (((fi14 / ((fi16) | 1))) | 1))
                                                switch "" {
                                                case "foo bar":
                                                    fmt.Println(p0, p1, p1)
                                                    var v19 uint16 = (v15 >> ((p0) & 63))
                                                    v12 -= (v12 ^ fi14)
                                                    _ = v19
                                                case "":
                                                    v12 += (fi16 - (p1 * 127))
                                                case "\n\t":
                                                    switch -2147483648 {
                                                    case 32768:
                                                        var v20 int32 = 2147483647
                                                        _ = v20
                                                    }
                                                }
                                            }
                                            var v21 uint64 = 42
                                            _ = v21
                                        } else {
                                            var v22 string = "foo bar"
                                            var v23 float64 = 0.0
                                            _ = v22
                                            _ = v23
                                        }
                                        _ = v18
                                    } else {
                                        switch "\n\t" {
                                        case "":
                                            switch "foo bar" {
                                            case "foo bar":
                                                if (uint8(v11) != 128) {
                                                    fmt.Println(v12, fi14)
                                                }
                                            }
                                        case "hello":
                                            var v24 int16 = 1
                                            _ = v24
                                        case "a":
                                            fmt.Println(fi16)
                                        default:
                                            var v25 uint16 = v15
                                            _ = v25
                                        }
                                    }
                                default:
                                    switch (fi14 - (fi14 | -2147483648)) {
                                    case 207:
                                        v26 := GenS1{F0: int(p0), F1: 858, F2: -2147483648}
                                        _ = v26
                                    case 0:
                                        var v27 uint = 65535
                                        _ = v27
                                    case 30:
                                        if len(v10) >= 2 {
                                            v10 = v10[1:len(v10)]
                                        }
                                    }
                                }
                            }
                            p1 = (fi16 % ((v12) | 1))
                            v28 := GenS1{F0: v12, F1: (^(v15)), F2: 0}
                            if len(v10) > 0 {
                                v10[0] = 127
                            }
                            var v29 uint32 = 13
                            _ = fi16
                            _ = v28
                            _ = v29
                        }
                        _ = fi14
                        _ = v15
                    }
                }
                _ = v11
                _ = v12
            } else {
                _ = len(v10)
                fmt.Println(p0, p1)
                _ = len(v10)
            }
            var v30 int = (p1 + (^(p1)))
            _ = v30
        }
        p0 -= (^(p0))
    }
    _ = p0
    _ = p1
    _ = v10
    return 137
}

func main() {
    var v1 uint8 = 255
    v2 := make(map[int]int)
    v3 := make(map[uint32]bool)
    var v4 int32 = 42
    v5 := make(map[string]int)
    var v6 int8 = 57
    fmt.Println(v4, v1)
    v31 := GenS1{F0: 1, F1: 65535, F2: -588}
    v6 -= v6
    var v32 uint16 = 0
    for fi33 := 0; fi33 < 6; fi33++ {
        v1 += v1
        _ = fi33
    }
    fmt.Println(v1, v6)
    delete(v5, "a")
    if (v6 > (-(v6))) {
        var v34 uint64 = 4294967295
        v35 := []int{2147483647, 330, -413, -9223372036854775808}
        fmt.Println(v6)
        _ = v34
        _ = v35
    } else {
        var v36 int64 = -525
        v37 := GenS2{F0: -72, F1: 1}
        v38 := GenS1{F0: -298, F1: 32767, F2: (v4 ^ v4)}
        _ = v36
        _ = v37
        _ = v38
    }
    for fi39 := 0; fi39 < 8; fi39++ {
        v40 := GenS1{F0: fi39, F1: v32, F2: (v4 / ((v4) | 1))}
        var v41 uint16 = v32
        v42 := GenS3{F0: genFn0((("\n\t" + "") + "foo bar"), 1), F1: v6, F2: v6}
        _ = fi39
        _ = v40
        _ = v41
        _ = v42
    }
    var v43 int16 = -612
    var v44 int64 = -102
    v45 := []int{-374, 0, -414, int(v44)}
    switch "" {
    case "":
        var v46 uint8 = v1
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
