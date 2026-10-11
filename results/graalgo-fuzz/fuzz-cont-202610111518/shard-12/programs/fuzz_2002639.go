package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint; F1 int8; F2 uint8 }
type GenS5 struct { F0 int; F1 int16; F2 bool }
type GenS6 struct { F0 int32; F1 int64; F2 uint64 }
func genFn0(p0 int8, p1 uint8, p2 int32) int64 {
    switch 127 {
    case 127:
        fmt.Println(p2)
    case 4294967295:
        for fi6 := 0; fi6 < 3; fi6++ {
            p0 += p0
            var v7 string = ""
            _ = fi6
            _ = v7
        }
    }
    v8 := GenS1{F0: 255}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 42
}
func genFn1(p0 uint8, p1 int) uint16 {
    if ((!true) && false) {
        var v9 int = p1
        _ = v9
    } else {
        if (!false) {
            var v10 uint64 = 8
            _ = v10
        } else {
            fmt.Println(p1, p0)
            v11 := []int{(p1 * p1), (p1 % ((p1) | 1)), (p1 % ((4294967295) | 1)), (p1 * p1)}
            var v12 uint = 594
            _ = v11
            _ = v12
        }
        v13 := GenS1{F0: 34}
        for fi14 := 0; fi14 < 17; fi14++ {
            for fi15 := 0; fi15 < 14; fi15++ {
                var v16 int32 = 42
                p0 = p0
                var v17 int32 = -1
                _ = fi15
                _ = v16
                _ = v17
            }
            var v18 uint8 = p0
            _ = fi14
            _ = v18
        }
        _ = v13
    }
    var v19 int32 = 0
    var v20 string = ("" + "\n\t")
    var v21 int = (p1 & (p1 ^ p1))
    v19 += (v19 % (((v19 - -514)) | 1))
    _ = p0
    _ = p1
    _ = v19
    _ = v20
    _ = v21
    return 759
}

func main() {
    var v1 uint8 = 235
    v2 := make(map[int]uint32)
    v3 := make(map[int]uint32)
    var v4 int64 = -979
    v5 := []uint32{}
    fmt.Println(v4, v1)
    fmt.Println(v1)
    switch 1 {
    case -878:
        if true {
            if true {
                v22 := GenS2{F0: 42, F1: -1}
                v23 := GenS3{F0: v4}
                _ = v22
                _ = v23
            } else {
                v4 = genFn0(100, (v1 >> ((1) & 63)), -1)
            }
            delete(v3, 373)
            var v24 bool = true
            _ = v24
        }
    default:
        var v25 int64 = -442
        _ = v25
    }
    var v26 uint32 = 2147483648
    switch ("\n\t" + "") {
    case "hello":
        if (-2147483648 < -74) {
            var v27 int16 = 0
            if (!(!false)) {
                for fi28 := 0; fi28 < 2; fi28++ {
                    var v29 uint8 = (^(v1))
                    v30 := GenS1{F0: uint16(v1)}
                    v4 = v4
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                var v31 uint32 = (v26 / (((v26 * (v26 % ((v26) | 1)))) | 1))
                v26 = v31
                _ = v31
            }
            _ = v27
        }
    }
    v32 := GenS4{F0: 679, F1: 42, F2: (^(v1))}
    if true {
        fmt.Println(v26)
    }
    if false {
        if (!true) {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        var v33 uint32 = (v26 & (v26 << ((0) & 63)))
        v33 -= (v26 * (v33 / (((^(v26))) | 1)))
        _ = v33
    } else {
        var v34 float32 = (1.5 / 1e10)
        for fi35 := 0; fi35 < 17; fi35++ {
            v26 = (v26 - uint32(v1))
            v1 = (v1 / (((v1 + 2)) | 1))
            for fi36 := 0; fi36 < 14; fi36++ {
                if len(v5) > 0 {
                    v5[0] = v26
                }
                v37 := GenS2{F0: (v4 >> ((uint(v26)) & 63)), F1: 42}
                if true {
                    if (42 != 1000) {
                        switch "a" {
                        case "":
                            v26 -= v26
                        }
                        v34 -= (0.1 - float32((0.1 / 3.14159)))
                        v38 := GenS5{F0: (fi36 ^ (^(fi35))), F1: 319, F2: true}
                        _ = v38
                    } else {
                        v5 = append(v5, (v26 & 2))
                        v4 -= v4
                    }
                    var v39 int32 = 118
                    v40 := GenS6{F0: (v39 / ((v39) | 1)), F1: v4, F2: 1}
                    _ = v39
                    _ = v40
                } else {
                    v41 := GenS2{F0: v4, F1: 0}
                    _ = v41
                }
                _ = fi36
                _ = v37
            }
            _ = fi35
        }
        _ = v34
    }
    v42 := GenS2{F0: v4, F1: 1000}
    var v43 string = "foo bar"
    switch v43 {
    case "hello":
        var v44 int16 = 564
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v32
    _ = v4
    _ = v42
    _ = v43
    _ = v5
}
