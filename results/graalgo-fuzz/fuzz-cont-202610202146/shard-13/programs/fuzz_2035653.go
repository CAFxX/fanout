package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 uint8 }
func genFn0() int64 {
    var v4 float32 = float32(0.1)
    fmt.Println(v4)
    v5 := []int{42, -1, -821, -1}
    for fi6 := 0; fi6 < 1; fi6++ {
        v4 = 3.14159
        v4 += (1e10 + (0.1 * ((v4 - v4) - ((0.0 * v4) + v4))))
        var v7 int64 = 9223372036854775807
        _ = fi6
        _ = v7
    }
    _ = v4
    _ = v5
    return 999
}
func genFn1() int16 {
    for fi8 := 0; fi8 < 14; fi8++ {
        var v9 int64 = -573
        v10 := GenS1{F0: true}
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 uint8 = 127
    _ = v11
    return 42
}

func main() {
    var v1 bool = true
    var v2 uint8 = 2
    var v3 bool = v1
    var v12 uint8 = 1
    for fi13 := 0; fi13 < 12; fi13++ {
        var v14 uint32 = 4294967295
        if (!((v14 | (v14 << ((42) & 63))) <= v14)) {
            for fi15 := 0; fi15 < 18; fi15++ {
                var v16 int32 = -1
                switch "hello" {
                case "\n\t":
                    v17 := []int{fi15}
                    _ = v17
                case "foo bar":
                    v18 := GenS2{F0: 1}
                    _ = v18
                case "hello":
                    if (1 >= -120) {
                        var v19 int64 = -795
                        _ = v19
                    } else {
                        v20 := []int{fi13, fi15, (fi15 * (fi15 / (((fi15 | fi15)) | 1))), 32767}
                        v16 = v16
                        if false {
                            v21 := GenS3{F0: 653, F1: v2}
                            var v22 uint64 = 42
                            _ = v21
                            _ = v22
                        }
                        _ = v20
                    }
                default:
                    switch "" {
                    case "":
                        if (255 != 255) {
                            switch (("a" + "foo bar") + "a") {
                            case "hello":
                                switch 1 {
                                case 127:
                                    var v23 int64 = 1
                                    _ = v23
                                case 1:
                                    v14 = 2147483648
                                case -460:
                                    switch fi13 {
                                    case 432:
                                        var v24 float64 = -2.25
                                        v3 = v3
                                        if v1 {
                                            v25 := GenS1{F0: (v3 || (!v1))}
                                            _ = v25
                                        }
                                        _ = v24
                                    }
                                    v26 := []int{(fi15 - -101), (fi13 << ((2) & 63)), fi15, 128}
                                    _ = v26
                                }
                            case "foo bar":
                                var v27 uint64 = 314
                                _ = v27
                            default:
                                v3 = true
                            }
                            v3 = true
                            var v28 float32 = 1e10
                            _ = v28
                        }
                    default:
                        v2 = (v2 * (v12 | v2))
                    }
                }
                v16 += 2147483647
                _ = fi15
                _ = v16
            }
            v29 := []int{fi13}
            fmt.Println(fi13)
            _ = v29
        }
        _ = fi13
        _ = v14
    }
    for fi30 := 0; fi30 < 20; fi30++ {
        var v31 float64 = 0.1
        v32 := []int{(fi30 / ((312) | 1)), (fi30 % ((fi30) | 1))}
        _ = fi30
        _ = v31
        _ = v32
    }
    v3 = (!(!((true && (v2 <= 217)) || (0 == 947))))
    for fi33 := 0; fi33 < 18; fi33++ {
        for fi34 := 0; fi34 < 1; fi34++ {
            var v35 int8 = 102
            v36 := GenS2{F0: 9223372036854775807}
            switch -1 {
            case -1:
                v35 += v35
            default:
                v3 = (fi33 <= -437)
            }
            _ = fi34
            _ = v35
            _ = v36
        }
        _ = fi33
    }
    v2 = (v12 & v2)
    for fi37 := 0; fi37 < 5; fi37++ {
        switch "" {
        case "foo bar":
            var v38 uint32 = 728
            _ = v38
        }
        switch fi37 {
        case 2147483647:
            var v39 int16 = 1000
            _ = v39
        case -884:
            for fi40 := 0; fi40 < 6; fi40++ {
                var v41 int8 = 1
                var v42 bool = (!((!true) || (v3 || (((255 > 112) && (v3 || ("" >= "foo bar"))) && (4294967295 <= 2147483647)))))
                _ = fi40
                _ = v41
                _ = v42
            }
        case 9223372036854775807:
            v43 := GenS1{F0: (!v1)}
            _ = v43
        }
        switch (fi37 ^ -2147483648) {
        case 364:
            v44 := GenS3{F0: -608, F1: (v12 / (((v2 / ((v12) | 1))) | 1))}
            _ = v44
        case -227:
            var v45 bool = v3
            _ = v45
        }
        _ = fi37
    }
    v46 := []int{9223372036854775807, 9223372036854775807, 2147483647}
    var v47 string = (("\n\t" + (("a" + "") + "a")) + "foo bar")
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v46
    _ = v47
}
