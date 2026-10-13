package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 uint; F1 uint64; F2 int }
type GenS3 struct { F0 uint; F1 uint64; F2 uint8 }
type GenS4 struct { F0 uint }
func genFn0(p0 bool) uint32 {
    var v5 float32 = 0.1
    var v6 int64 = -56
    fmt.Println(v5, v5, v6)
    _ = p0
    _ = v5
    _ = v6
    return 2147483647
}
func genFn1() int64 {
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := GenS1{F0: "foo bar", F1: 0}
        _ = fi7
        _ = v8
    }
    switch "a" {
    case "a":
        for fi9 := 0; fi9 < 14; fi9++ {
            var v10 float64 = float64((0.0 * float32((0.1 / 1.0))))
            var v11 uint16 = 42
            var v12 float64 = (float64((((0.0 / 1.0) + (float32(v10) - 1e10)) / 1.5)) / v10)
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
    case "\n\t":
        var v13 int8 = -1
        _ = v13
    case "foo bar":
        v14 := GenS2{F0: 4294967295, F1: 590, F2: -1}
        _ = v14
    default:
        for fi15 := 0; fi15 < 5; fi15++ {
            fmt.Println(fi15)
            switch -520 {
            case -2147483648:
                for fi16 := 0; fi16 < 10; fi16++ {
                    if (0.1 != 1e10) {
                        fmt.Println(fi15, fi15)
                        switch (fi16 * -9223372036854775808) {
                        case 1:
                            var v17 uint = 0
                            _ = v17
                        }
                    }
                    fmt.Println(fi16, fi15)
                    fmt.Println(fi16, fi15)
                    var v18 uint8 = 127
                    _ = fi16
                    _ = v18
                }
            case -122:
                var v19 int16 = 1000
                _ = v19
            case 4294967295:
                if false {
                    fmt.Println(fi15)
                    if true {
                        for fi20 := 0; fi20 < 7; fi20++ {
                            fmt.Println(fi15, fi15)
                            switch (-(fi20)) {
                            case 0:
                                v21 := []int{(^(fi20)), -343, (fi15 | (^(fi15)))}
                                _ = v21
                            case 32767:
                                if false {
                                    switch fi20 {
                                    case 29:
                                        var v22 uint16 = 678
                                        _ = v22
                                    }
                                    switch "a" {
                                    case "\n\t":
                                        var v23 uint32 = 1
                                        v23 += (v23 * (^(v23)))
                                        _ = v23
                                    }
                                } else {
                                    v24 := []int{fi20, 127, (^(fi20))}
                                    var v25 int64 = 42
                                    _ = len(v24)
                                    fmt.Println(v25, v25, v25)
                                    v25 += (^(v25))
                                    v25 += v25
                                    v26 := GenS1{F0: "\n\t", F1: 579}
                                    _ = v24
                                    _ = v25
                                    _ = v26
                                }
                            case -1:
                                if true {
                                    var v27 string = ("a" + ("\n\t" + "foo bar"))
                                    v27 = v27
                                    fmt.Println(fi20, fi20)
                                    var v28 uint8 = 1
                                    _ = v27
                                    _ = v28
                                } else {
                                    v29 := GenS1{F0: "", F1: 0}
                                    if false {
                                        var v30 uint32 = 42
                                        _ = v30
                                    } else {
                                        v31 := GenS3{F0: 1, F1: 0, F2: 0}
                                        _ = v31
                                    }
                                    fmt.Println(fi15, fi20)
                                    var v32 uint8 = 2
                                    _ = v29
                                    _ = v32
                                }
                            }
                            v33 := []int{fi20, (fi15 ^ fi20), (fi15 - (fi15 % ((fi15) | 1))), -467}
                            _ = fi20
                            _ = v33
                        }
                        switch fi15 {
                        case -250:
                            v34 := []int{4294967295, (fi15 >> ((0) & 63)), fi15, 42}
                            _ = v34
                        case -1:
                            v35 := []int{fi15, (fi15 << ((1) & 63)), (fi15 + fi15), (fi15 & 313)}
                            _ = v35
                        }
                    }
                } else {
                    var v36 uint16 = 0
                    _ = v36
                }
            default:
                v37 := GenS4{F0: 65535}
                _ = v37
            }
            if true {
                v38 := []int{}
                var v39 int32 = -1
                var v40 int64 = -704
                _ = v38
                _ = v39
                _ = v40
            }
            _ = fi15
        }
    }
    return 0
}

func main() {
    v1 := make(map[int]int)
    var v2 int64 = -9223372036854775808
    v3 := []int{9223372036854775807, -598, 128, 2147483647}
    v4 := make(map[uint32]string)
    v2 = (v2 << ((255) & 63))
    switch -1 {
    case 969:
        if ("a" > "foo bar") {
            v2 = (v2 | (v2 | -294))
        }
    case 4294967295:
        var v41 int64 = -494
        _ = v41
    }
    fmt.Println(v2)
    var v42 int64 = (v2 - (v2 ^ 9223372036854775807))
    v2 -= genFn1()
    v2 += (v42 / ((v2) | 1))
    v43 := GenS1{F0: "foo bar", F1: 32767}
    if false {
        v42 += (v2 * (v2 - (v42 + genFn1())))
        switch 4294967295 {
        case 128:
            var v44 string = ("hello" + "\n\t")
            _ = v44
        case -2147483648:
            fmt.Println(v42, v2)
        case 2147483647:
            if len(v3) > 0 {
                v3[0] = 564
            }
        default:
            if (65535 <= 255) {
                if false {
                    v42 = (v42 - (v2 * 0))
                }
            } else {
                _ = len(v3)
                v45 := v4[42]
                _ = v45
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
}
