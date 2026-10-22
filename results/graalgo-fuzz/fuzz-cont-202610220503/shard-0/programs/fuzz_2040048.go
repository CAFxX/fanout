package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 uint16 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 uint64 }
func genFn0(p0 uint64) int8 {
    switch 32768 {
    case 1:
        var v7 int16 = -660
        _ = v7
    }
    for fi8 := 0; fi8 < 2; fi8++ {
        if false {
            fmt.Println(fi8, fi8)
        } else {
            switch (fi8 & fi8) {
            case -900:
                if true {
                    fmt.Println(fi8, fi8)
                } else {
                    p0 = (p0 % ((p0) | 1))
                }
            case -1:
                var v9 int16 = 1
                _ = v9
            }
            var v10 float32 = 0.0
            for fi11 := 0; fi11 < 5; fi11++ {
                var v12 int16 = -32768
                v13 := GenS1{F0: 36, F1: -1, F2: 32767}
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = v10
        }
        if false {
            fmt.Println(p0)
            var v14 int32 = -2147483648
            _ = v14
        } else {
            switch fi8 {
            case 4294967295:
                var v15 uint16 = 65535
                _ = v15
            case 0:
                for fi16 := 0; fi16 < 12; fi16++ {
                    var v17 int16 = 42
                    if (127 <= uint8(p0)) {
                        fmt.Println(fi8)
                        p0 += (p0 >> ((109) & 63))
                        v18 := []int{fi8, fi8, (fi8 << ((2) & 63))}
                        _ = v18
                    } else {
                        if (0 == 148) {
                            if true {
                                v17 = 1
                                if false {
                                    v19 := GenS1{F0: 0, F1: 42, F2: 827}
                                    v20 := GenS1{F0: -77, F1: -674, F2: 898}
                                    v21 := []int{}
                                    _ = v19
                                    _ = v20
                                    _ = v21
                                } else {
                                    fmt.Println(fi8, v17, v17)
                                    var v22 int16 = v17
                                    _ = v22
                                }
                                var v23 float64 = 0.1
                                _ = v23
                            }
                            v17 += v17
                            fmt.Println(v17)
                        } else {
                            v17 += v17
                        }
                        var v24 int16 = (v17 - v17)
                        _ = v24
                    }
                    if false {
                        fmt.Println(fi16, fi16)
                        switch "foo bar" {
                        case "\n\t":
                            v25 := []int{fi16, 128, fi8}
                            _ = v25
                        case "hello":
                            if ((!true) && true) {
                                p0 -= (p0 & p0)
                            }
                        default:
                            p0 += (p0 * (^(p0)))
                        }
                        v26 := []int{fi16, (-(fi16)), (fi8 * 0), (fi8 + (fi8 >> ((18446744073709551615) & 63)))}
                        _ = v26
                    }
                    _ = fi16
                    _ = v17
                }
            case -2147483648:
                switch (^(fi8)) {
                case -1:
                    switch fi8 {
                    case 874:
                        if (!false) {
                            p0 = p0
                            if false {
                                var v27 int16 = 0
                                v27 -= v27
                                _ = v27
                            } else {
                                p0 = (p0 % (((p0 << ((371) & 63))) | 1))
                                v28 := GenS1{F0: 1, F1: 2147483647, F2: uint16(p0)}
                                fmt.Println(fi8, p0)
                                _ = v28
                            }
                        }
                    case 9223372036854775807:
                        v29 := []int{fi8, 32768}
                        _ = v29
                    default:
                        var v30 uint64 = 1
                        _ = v30
                    }
                case -9223372036854775808:
                    var v31 uint = 506
                    _ = v31
                }
            default:
                v32 := GenS1{F0: -83, F1: -35, F2: 42}
                _ = v32
            }
        }
        v33 := GenS1{F0: 100, F1: 2147483647, F2: 65535}
        _ = fi8
        _ = v33
    }
    p0 = (p0 / (((p0 + 1)) | 1))
    _ = p0
    return 1
}

func main() {
    var v1 uint64 = 1
    v2 := []int{-452, 0, -345}
    var v3 uint64 = 2
    v4 := make(map[string]uint32)
    var v5 int64 = 9223372036854775807
    var v6 uint64 = 519
    v1 = v3
    if true {
        for fi34 := 0; fi34 < 17; fi34++ {
            fmt.Println(v3, v6, v6)
            v5 = (v5 << ((uint(v6)) & 63))
            _ = fi34
        }
    }
    v35 := GenS1{F0: -64, F1: 42, F2: 42}
    fmt.Println(v3, v1)
    if true {
        var v36 int = -494
        v5 -= -1
        switch "\n\t" {
        case "\n\t":
            v37 := GenS2{F0: 32767, F1: ("foo bar" + "\n\t")}
            _ = v37
        default:
            var v38 int32 = -1
            _ = v38
        }
        _ = v36
    }
    var v39 float64 = (-2.25 / ((-2.25 + -2.25) - 0.1))
    var v40 uint32 = 390
    var v41 int = 32767
    v6 += v1
    var v42 int = (^(v41))
    var v43 int8 = 127
    v3 = 911
    v44 := GenS3{F0: v6}
    if len(v2) > 0 {
        v2[0] = (v42 & (-(v41)))
    }
    v45 := GenS2{F0: 835, F1: ""}
    fmt.Println(v43)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
