package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int8; F1 int64 }
type GenS6 struct { F0 uint16; F1 int32 }
func genFn0(p0 int16, p1 string) uint8 {
    fmt.Println(p1, p0)
    p1 = ((p1 + "a") + "")
    v7 := []int{9223372036854775807, 9223372036854775807, 127}
    v8 := GenS1{F0: 759}
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 uint8 = 1
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 uint8 = 207
    v2 := []uint32{}
    v3 := make(map[int]int)
    v4 := make(map[string]uint32)
    var v5 int16 = 42
    v6 := make(map[uint32]uint32)
    switch "a" {
    case "":
        v11 := GenS2{F0: 21}
        _ = v11
    case "a":
        var v12 int8 = 42
        _ = v12
    }
    delete(v4, "\n\t")
    if (!(!true)) {
        v13 := GenS3{F0: 2, F1: 626}
        _ = len(v6)
        _ = v13
    } else {
        if (249 <= 128) {
            v1 = (^(v1))
            var v14 uint32 = 64
            if false {
                var v15 float32 = (0.0 + 0.0)
                _ = v15
            }
            _ = v14
        } else {
            fmt.Println(v5, v5)
            for fi16 := 0; fi16 < 11; fi16++ {
                for fi17 := 0; fi17 < 4; fi17++ {
                    fmt.Println(v5, fi17)
                    v18 := GenS3{F0: 42, F1: -737}
                    _ = fi17
                    _ = v18
                }
                v1 += 127
                v1 = genFn0(v5, "\n\t")
                _ = fi16
            }
        }
    }
    var v19 float32 = (1.5 + -2.25)
    for fi20 := 0; fi20 < 13; fi20++ {
        var v21 float64 = float64(v19)
        var v22 int16 = 0
        if true {
            delete(v3, fi20)
            var v23 uint8 = (^(v1))
            _ = v23
        } else {
            fmt.Println(v1, v1)
            var v24 int8 = 127
            _ = v24
        }
        _ = fi20
        _ = v21
        _ = v22
    }
    if (int32(v5) == -1) {
        fmt.Println(v1)
        v19 += v19
    }
    v19 += 0.1
    switch "" {
    case "hello":
        switch "" {
        case "foo bar":
            v1 = v1
        default:
            for fi25 := 0; fi25 < 1; fi25++ {
                if (-1 < 42) {
                    var v26 int32 = -237
                    fmt.Println(v19)
                    _ = v26
                }
                _ = fi25
            }
        }
    case "a":
        v27 := GenS4{F0: -305}
        _ = v27
    }
    switch "" {
    case "":
        for fi28 := 0; fi28 < 19; fi28++ {
            v29 := v6[692]
            _ = fi28
            _ = v29
        }
    case "\n\t":
        v30 := GenS4{F0: 184}
        _ = v30
    case "a":
        v19 = v19
    default:
        if false {
            for fi31 := 0; fi31 < 14; fi31++ {
                v1 -= (v1 / (((v1 << ((42) & 63))) | 1))
                fmt.Println(v1, fi31)
                for fi32 := 0; fi32 < 18; fi32++ {
                    v1 = genFn0(1, "")
                    _ = fi32
                }
                _ = fi31
            }
            var v33 int64 = 1
            v5 -= (^(v5))
            _ = v33
        } else {
            if (127 < 0) {
                var v34 uint64 = 0
                if false {
                    for fi35 := 0; fi35 < 1; fi35++ {
                        fmt.Println(v34, v19)
                        var v36 int64 = -9223372036854775808
                        _ = fi35
                        _ = v36
                    }
                }
                _ = v34
            } else {
                delete(v3, -373)
                v37 := GenS5{F0: 42, F1: -27}
                if (true || ((!(-526 < 0)) || (int16(v1) >= 1000))) {
                    if (42 != 42) {
                        fmt.Println(v5, v19)
                        switch 1 {
                        case 9223372036854775807:
                            if false {
                                var v38 int = 42
                                var v39 uint8 = v1
                                _ = len(v3)
                                _ = v38
                                _ = v39
                            }
                        case 695:
                            fmt.Println(v5, v1, v1)
                        default:
                            var v40 int8 = 0
                            _ = v40
                        }
                        v41 := GenS4{F0: 9223372036854775807}
                        _ = v41
                    } else {
                        v5 -= (v5 / ((int16(v1)) | 1))
                    }
                }
                _ = v37
            }
            _ = len(v2)
            v42 := GenS6{F0: 146, F1: 348}
            _ = v42
        }
    }
    v4[(("a" + ("hello" + ((("" + ("hello" + "hello")) + "hello") + "\n\t"))) + ("foo bar" + ("hello" + "\n\t")))] = 2
    var v43 string = (("\n\t" + "hello") + "")
    v1 = v1
    var v44 float32 = v19
    v45 := GenS4{F0: -9223372036854775808}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
