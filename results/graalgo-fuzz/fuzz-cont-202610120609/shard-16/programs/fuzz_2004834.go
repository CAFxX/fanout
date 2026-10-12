package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 int64; F2 int8 }
type GenS3 struct { F0 bool; F1 uint }
func genFn0() int32 {
    if (42 >= 127) {
        var v7 int64 = 1
        for fi8 := 0; fi8 < 18; fi8++ {
            fmt.Println(fi8, v7)
            for fi9 := 0; fi9 < 3; fi9++ {
                if false {
                    switch fi9 {
                    case -9223372036854775808:
                        var v10 uint16 = 115
                        _ = v10
                    }
                    switch "hello" {
                    case "foo bar":
                        v11 := GenS1{F0: 187}
                        _ = v11
                    }
                    v7 += v7
                } else {
                    var v12 int64 = v7
                    v13 := []int{(fi8 / ((fi8) | 1)), (fi8 / (((-(fi9))) | 1))}
                    if false {
                        var v14 uint32 = 156
                        v7 = (v12 >> ((65535) & 63))
                        v13 = append(v13, (fi8 ^ (fi8 & 77)))
                        var v15 bool = (117 == 1)
                        var v16 uint16 = 607
                        _ = v14
                        _ = v15
                        _ = v16
                    } else {
                        var v17 int16 = -32768
                        if len(v13) > 0 {
                            v13[0] = (fi8 | (fi9 | -2147483648))
                        }
                        var v18 string = "\n\t"
                        _ = v17
                        _ = v18
                    }
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                    _ = v12
                    _ = v13
                }
                v7 -= (v7 >> ((4294967295) & 63))
                v7 += 42
                _ = fi9
            }
            _ = fi8
        }
        if (false || true) {
            var v19 int8 = -128
            var v20 int16 = 634
            if (v20 > v20) {
                v21 := GenS2{F0: 925, F1: 114, F2: -32}
                _ = v21
            }
            _ = v19
            _ = v20
        }
        _ = v7
    }
    if (295 >= 1) {
        v22 := []int{-1, 497, 32767}
        _ = v22
    }
    for fi23 := 0; fi23 < 5; fi23++ {
        if true {
            v24 := []int{}
            _ = v24
        }
        _ = fi23
    }
    switch "" {
    case "\n\t":
        var v25 float64 = 0.0
        _ = v25
    case "":
        v26 := []int{}
        _ = v26
    default:
        var v27 float64 = (((float64((float32(-2.25) / -2.25)) - 0.0) * 1.5) / 1.5)
        _ = v27
    }
    return 42
}
func genFn1(p0 int64, p1 int8, p2 uint) int32 {
    fmt.Println(p2)
    v28 := GenS2{F0: 0, F1: (p0 ^ p0), F2: (p1 | -1)}
    for fi29 := 0; fi29 < 6; fi29++ {
        v30 := []int{}
        if ((p2 % ((p2) | 1)) > p2) {
            p1 -= p1
            p2 = p2
            v31 := GenS3{F0: false, F1: p2}
            _ = v31
        } else {
            v32 := GenS1{F0: 255}
            var v33 int32 = 2147483647
            fmt.Println(p0, p0, p2)
            _ = v32
            _ = v33
        }
        fmt.Println(p1)
        _ = fi29
        _ = v30
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v28
    return -113
}

func main() {
    var v1 uint = 1
    var v2 uint64 = 256
    v3 := []uint32{1, 920, 787, 42}
    var v4 uint = (v1 >> ((4294967295) & 63))
    v5 := make(map[int]string)
    var v6 float64 = 1.5
    switch 0 {
    case -466:
        switch "hello" {
        case "":
            var v34 float64 = v6
            _ = v34
        default:
            delete(v5, 0)
        }
    case 32768:
        var v35 uint = v4
        _ = v35
    }
    var v36 uint32 = 2147483648
    var v37 uint = v4
    v5[32767] = ""
    var v38 float32 = -2.25
    var v39 int64 = 42
    v40 := GenS3{F0: false, F1: 785}
    for fi41 := 0; fi41 < 10; fi41++ {
        v37 += v1
        var v42 int16 = 32767
        fmt.Println(v1)
        _ = fi41
        _ = v42
    }
    if true {
        var v43 uint64 = v2
        var v44 uint8 = 255
        _ = v43
        _ = v44
    } else {
        for fi45 := 0; fi45 < 17; fi45++ {
            v37 = (v1 / ((v1) | 1))
            v3 = append(v3, v36)
            _ = fi45
        }
    }
    for fi46 := 0; fi46 < 1; fi46++ {
        var v47 float32 = v38
        for fi48 := 0; fi48 < 9; fi48++ {
            v3 = append(v3, (v36 & (^(v36))))
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v49 uint16 = 471
            fmt.Println(v37, v49)
            var v50 int16 = 32767
            fmt.Println(v39, v6, v49)
            var v51 uint8 = 127
            v49 += 1
            if ((-2147483648 <= 2147483647) || true) {
                v49 += (v49 / ((uint16(v37)) | 1))
            } else {
                fmt.Println(v6, v36)
                var v52 int8 = 0
                var v53 string = "foo bar"
                _ = v52
                _ = v53
            }
            _ = fi48
            _ = v49
            _ = v50
            _ = v51
        }
        if true {
            fmt.Println(v36)
            var v54 int = fi46
            v5[9223372036854775807] = "hello"
            _ = v54
        } else {
            delete(v5, (fi46 >> ((v1) & 63)))
        }
        _ = fi46
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
