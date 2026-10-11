package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 uint; F1 uint8 }
func genFn0(p0 uint8, p1 int64, p2 int) string {
    switch p2 {
    case 42:
        fmt.Println(p1, p2, p1)
    case -215:
        var v4 int16 = -32768
        _ = v4
    case -1:
        for fi5 := 0; fi5 < 10; fi5++ {
            p0 = (p0 * p0)
            _ = fi5
        }
    default:
        fmt.Println(p2)
    }
    switch (p2 >> ((293) & 63)) {
    case 127:
        v6 := []int{p2, p2, (-(p2))}
        _ = v6
    case 4294967295:
        v7 := GenS1{F0: (p1 + p1), F1: 255}
        _ = v7
    default:
        for fi8 := 0; fi8 < 16; fi8++ {
            var v9 uint32 = 297
            _ = fi8
            _ = v9
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return ("\n\t" + "\n\t")
}
func genFn1(p0 uint, p1 int16, p2 int32) uint8 {
    if ((1 == 32767) || ("" < "\n\t")) {
        fmt.Println(p1, p2, p2)
    }
    p0 -= p0
    if false {
        v10 := []int{0, 127, 32768, 4294967295}
        v11 := GenS1{F0: -556, F1: 255}
        _ = v10
        _ = v11
    } else {
        p0 -= p0
        v12 := []int{9223372036854775807, 32767, 32768}
        for fi13 := 0; fi13 < 8; fi13++ {
            var v14 uint8 = 2
            _ = fi13
            _ = v14
        }
        _ = v12
    }
    var v15 bool = true
    if ((-527 != -415) || ("\n\t" <= "foo bar")) {
        if ((!v15) && true) {
            fmt.Println(p2)
            fmt.Println(p2, p1)
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return 255
}

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{}
    v3 := make(map[int]bool)
    if (457 <= 607) {
        var v16 float32 = float32(3.14159)
        switch -9223372036854775808 {
        case 476:
            v17 := GenS1{F0: 9223372036854775807, F1: 42}
            _ = v17
        case -917:
            v18 := GenS1{F0: 0, F1: 2}
            _ = v18
        case 901:
            v16 = v16
        default:
            if (42 > 0) {
                if true {
                    for fi19 := 0; fi19 < 8; fi19++ {
                        switch fi19 {
                        case 44:
                            fmt.Println(v16)
                        default:
                            v16 = 1.5
                        }
                        switch (fi19 << ((377) & 63)) {
                        case -1:
                            var v20 uint64 = 1
                            _ = v20
                        case 128:
                            var v21 int16 = -1
                            _ = v21
                        }
                        _ = fi19
                    }
                } else {
                    if true {
                        v16 += v16
                        _ = len(v3)
                        v22 := GenS2{F0: 255, F1: 42}
                        _ = v22
                    } else {
                        for fi23 := 0; fi23 < 3; fi23++ {
                            if true {
                                if true {
                                    v24 := GenS2{F0: 0, F1: 102}
                                    v16 += ((1.5 / v16) * (1.5 + (((v16 - (v16 + (3.14159 - 1.5))) / ((v16 + 1.5) * (v16 - 1e10))) + (1.5 * 0.0))))
                                    _ = v24
                                }
                                var v25 uint32 = 2147483648
                                var v26 float64 = -2.25
                                _ = v25
                                _ = v26
                            }
                            _ = len(v3)
                            var v27 int32 = -413
                            _ = fi23
                            _ = v27
                        }
                        var v28 uint32 = 1
                        _ = v28
                    }
                    var v29 int32 = -2147483648
                    for fi30 := 0; fi30 < 7; fi30++ {
                        v29 = v29
                        _ = fi30
                    }
                    _ = v29
                }
            }
        }
        _ = v16
    }
    var v31 int8 = 0
    var v32 int32 = 2147483647
    v1["a"] = "\n\t"
    for fi33 := 0; fi33 < 14; fi33++ {
        var v34 bool = false
        for fi35 := 0; fi35 < 3; fi35++ {
            fmt.Println(v31, v34)
            _ = fi35
        }
        switch "foo bar" {
        case "a":
            v31 -= 71
        case "":
            switch "a" {
            case "\n\t":
                var v36 int32 = (v32 >> ((255) & 63))
                _ = v36
            default:
                v34 = false
            }
        case "foo bar":
            for fi37 := 0; fi37 < 19; fi37++ {
                fmt.Println(v34, v34)
                v34 = (977 <= 32767)
                _ = fi37
            }
        default:
            switch fi33 {
            case 127:
                var v38 int32 = v32
                _ = v38
            case 32768:
                fmt.Println(fi33, v34)
            }
        }
        _ = fi33
        _ = v34
    }
    v32 = v32
    var v39 uint32 = 61
    fmt.Println(v32, v32)
    v40 := GenS1{F0: 383, F1: 561}
    var v41 int8 = v31
    fmt.Println(v39)
    v31 -= (^(v31))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v39:", v39)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v39
    _ = v40
    _ = v41
}
