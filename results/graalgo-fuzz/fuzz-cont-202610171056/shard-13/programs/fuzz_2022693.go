package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint64; F2 int8 }
type GenS3 struct { F0 int8; F1 int64 }
func genFn0(p0 int32) string {
    p0 += -2147483648
    v7 := GenS1{F0: 0}
    _ = p0
    _ = v7
    return "a"
}

func main() {
    var v1 bool = true
    var v2 uint = 18446744073709551615
    v3 := []uint32{42, 608, 0, 0}
    var v4 uint = (v2 >> (((v2 << ((42) & 63))) & 63))
    v5 := []string{("foo bar" + "foo bar"), ("" + ("a" + ("" + "hello"))), "foo bar", "\n\t"}
    v6 := make(map[int]uint32)
    for fi8 := 0; fi8 < 14; fi8++ {
        if true {
            _ = len(v5)
            for fi9 := 0; fi9 < 1; fi9++ {
                fmt.Println(v4, v4)
                var v10 uint8 = 128
                _ = fi9
                _ = v10
            }
            if v1 {
                var v11 bool = true
                _ = v11
            }
        } else {
            var v12 uint = (v4 + v4)
            _ = v12
        }
        switch "a" {
        case "":
            var v13 uint = v2
            _ = v13
        case "foo bar":
            for fi14 := 0; fi14 < 5; fi14++ {
                switch (fi8 + fi8) {
                case -709:
                    v15 := GenS1{F0: 1}
                    _ = v15
                default:
                    if v1 {
                        if ((false && false) || (!v1)) {
                            switch "\n\t" {
                            case "\n\t":
                                v16 := GenS2{F0: 255, F1: 42, F2: -105}
                                _ = v16
                            case "a":
                                fmt.Println(v2, v4, fi14)
                            case "hello":
                                switch (fi8 ^ fi8) {
                                case -2147483648:
                                    if true {
                                        v6[-595] = 65535
                                    } else {
                                        var v17 uint32 = 909
                                        var v18 uint32 = (^(v17))
                                        _ = v17
                                        _ = v18
                                    }
                                case -744:
                                    v1 = v1
                                default:
                                    fmt.Println(v2, v4, v4)
                                }
                            default:
                                v2 = v4
                            }
                            v2 += (^(v2))
                        } else {
                            v19 := GenS2{F0: 1, F1: 1, F2: -128}
                            _ = v19
                        }
                        if len(v5) > 0 {
                            v5[0] = "\n\t"
                        }
                    } else {
                        var v20 int = (^(fi14))
                        v2 = v4
                        if false {
                            var v21 string = ("" + "\n\t")
                            v1 = true
                            _ = v21
                        } else {
                            if (764 != 4294967295) {
                                var v22 int32 = int32(v20)
                                _ = v22
                            } else {
                                if (1 < uint16(v4)) {
                                    _ = len(v5)
                                    var v23 int16 = -1
                                    switch "" {
                                    case "\n\t":
                                        if v1 {
                                            if (42 < 1) {
                                                if (260 < 42) {
                                                    v20 -= (v20 * fi14)
                                                    var v24 int16 = (v23 | v23)
                                                    _ = v24
                                                }
                                                v5 = append(v5, (("foo bar" + "\n\t") + ((("\n\t" + "hello") + "foo bar") + "\n\t")))
                                                var v25 int16 = -1
                                                _ = v25
                                            } else {
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            }
                                            v1 = (v23 == v23)
                                            fmt.Println(fi14)
                                        } else {
                                            v2 = (^(v4))
                                            switch "" {
                                            case "a":
                                                v1 = v1
                                            default:
                                                if (v1 || v1) {
                                                    v26 := v6[-323]
                                                    if true {
                                                        v27 := GenS1{F0: 255}
                                                        _ = len(v6)
                                                        _ = v27
                                                    } else {
                                                        v5 = append(v5, "foo bar")
                                                        fmt.Println(v26)
                                                    }
                                                    _ = v26
                                                }
                                            }
                                            if (v1 || (287 > 2)) {
                                                _ = len(v6)
                                            }
                                        }
                                    case "hello":
                                        v6[(fi14 * 719)] = 104
                                    case "a":
                                        var v28 float64 = 1.5
                                        _ = v28
                                    }
                                    _ = v23
                                }
                                var v29 float64 = 0.0
                                _ = v29
                            }
                            var v30 string = "a"
                            _ = v30
                        }
                        v1 = (1 < 1)
                        if (634 >= -1) {
                            fmt.Println(fi8, v2)
                        } else {
                            _ = len(v6)
                            if (0 < 42) {
                                delete(v6, v20)
                                v31 := GenS1{F0: 788}
                                if v1 {
                                    var v32 float32 = float32(1e10)
                                    _ = v32
                                }
                                _ = v31
                            } else {
                                v4 -= (v4 | (^(v2)))
                            }
                            v4 += v4
                        }
                        _ = v20
                    }
                }
                var v33 int16 = 1
                _ = fi14
                _ = v33
            }
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi8
    }
    var v34 uint64 = 42
    if (!v1) {
        if (!true) {
            _ = len(v5)
            v34 = v34
            fmt.Println(v34)
        }
    } else {
        if len(v3) > 0 {
            v3[0] = 65535
        }
    }
    if true {
        v35 := GenS1{F0: 996}
        if ((!v1) && ("a" > "a")) {
            if (((878 > 617) && (!(v34 <= v34))) && ((v1 && v1) || ((4294967295 >= (^(v34))) || (true || v1)))) {
                v2 = v2
                _ = len(v5)
                v6[128] = 2
            }
            v36 := GenS3{F0: -94, F1: 42}
            if (float32(1e10) >= ((1e10 / (((1.5 * -2.25) / (0.1 * -2.25)) + 3.14159)) * (float32((0.0 + 0.0)) / 1.0))) {
                var v37 bool = (!false)
                _ = v37
            }
            _ = v36
        }
        switch -315 {
        case -59:
            switch "a" {
            case "":
                switch 128 {
                case 4294967295:
                    v4 -= v4
                case 742:
                    v34 -= v34
                case 1:
                    fmt.Println(v34, v4)
                }
            }
        case 128:
            var v38 int32 = -1
            _ = v38
        }
        _ = v35
    }
    v2 += (v2 << ((v2) & 63))
    var v39 uint32 = 1
    var v40 float64 = 0.1
    for fi41 := 0; fi41 < 9; fi41++ {
        v1 = (fi41 != (^(fi41)))
        _ = fi41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
