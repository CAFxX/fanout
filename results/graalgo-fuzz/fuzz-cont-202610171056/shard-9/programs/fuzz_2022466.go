package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 uint8; F1 uint32; F2 uint16 }
func genFn0(p0 int32, p1 int8) bool {
    v7 := []int{9223372036854775807, 127, 32768}
    _ = len(v7)
    var v8 float32 = -2.25
    if ((p0 - 0) > p0) {
        v9 := GenS1{F0: 0}
        _ = v9
    } else {
        switch 776 {
        case 267:
            p0 = p0
        case -786:
            fmt.Println(v8, p1)
        default:
            if false {
                v8 = 1e10
                v10 := GenS2{F0: 764, F1: 44}
                var v11 int32 = (-(p0))
                _ = v10
                _ = v11
            }
        }
    }
    fmt.Println(v8, p1, p0)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return ((false || true) || (true && (42 == 0)))
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]uint32)
    v3 := []string{"a"}
    v4 := []uint32{}
    v5 := make(map[string]uint32)
    v6 := make(map[int]bool)
    if (-2147483648 >= -1) {
        var v12 uint8 = 179
        _ = v12
    }
    v6[-9223372036854775808] = (-2147483648 > 32768)
    for fi13 := 0; fi13 < 13; fi13++ {
        var v14 uint64 = 501
        _ = fi13
        _ = v14
    }
    v4 = append(v4, 42)
    for fi15 := 0; fi15 < 17; fi15++ {
        var v16 uint = 42
        _ = fi15
        _ = v16
    }
    var v17 int8 = 42
    var v18 uint = 178
    var v19 int32 = -1
    for fi20 := 0; fi20 < 14; fi20++ {
        for fi21 := 0; fi21 < 14; fi21++ {
            v19 = -1
            _ = fi21
        }
        for fi22 := 0; fi22 < 4; fi22++ {
            v17 += (v17 << (((v18 ^ v18)) & 63))
            v3 = append(v3, "\n\t")
            _ = fi22
        }
        if true {
            if (!false) {
                v18 -= v18
            }
            if true {
                var v23 uint32 = 736
                v19 = v19
                _ = v23
            }
            v19 -= (v19 >> ((381) & 63))
        } else {
            switch (fi20 ^ 127) {
            case 32767:
                var v24 int64 = int64(v19)
                _ = v24
            case 2147483647:
                v25 := GenS2{F0: 0, F1: -9223372036854775808}
                _ = v25
            case 127:
                v26 := GenS3{F0: 255, F1: 2147483648, F2: 65535}
                _ = v26
            default:
                if len(v4) > 0 {
                    v4[0] = 2147483648
                }
            }
            if (!(0.1 < ((1e10 - 0.0) * 0.1))) {
                if (!true) {
                    if false {
                        var v27 uint16 = 65535
                        _ = v27
                    } else {
                        if len(v4) > 0 {
                            v4[0] = 2
                        }
                    }
                    v28 := v6[fi20]
                    for fi29 := 0; fi29 < 16; fi29++ {
                        var v30 uint64 = 42
                        if (432 > 836) {
                            fmt.Println(v28)
                            v4 = append(v4, 523)
                            v30 = (^(v30))
                        } else {
                            var v31 uint32 = 42
                            var v32 float32 = (1e10 * 0.0)
                            var v33 int16 = -1
                            _ = v31
                            _ = v32
                            _ = v33
                        }
                        _ = fi29
                        _ = v30
                    }
                    _ = v28
                }
            } else {
                v34 := GenS3{F0: 127, F1: 1, F2: 491}
                _ = v34
            }
        }
        _ = fi20
    }
    var v35 float64 = 1.5
    v4 = append(v4, 1)
    fmt.Println(v17)
    var v36 int64 = 580
    v37 := v5["\n\t"]
    for fi38 := 0; fi38 < 1; fi38++ {
        var v39 int8 = v17
        var v40 float32 = (1.5 / float32(v35))
        _ = fi38
        _ = v39
        _ = v40
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
