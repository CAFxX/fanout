package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 bool; F1 uint; F2 uint64 }
type GenS7 struct { F0 uint8 }
func genFn0(p0 int64, p1 uint64, p2 int8) uint32 {
    if true {
        for fi6 := 0; fi6 < 7; fi6++ {
            for fi7 := 0; fi7 < 19; fi7++ {
                fmt.Println(fi6, p1, fi6)
                v8 := []int{(fi6 - fi6), 127, (fi6 - 195)}
                p2 += p2
                _ = fi7
                _ = v8
            }
            for fi9 := 0; fi9 < 3; fi9++ {
                switch (fi6 ^ fi9) {
                case 128:
                    var v10 float32 = 0.0
                    _ = v10
                case -2147483648:
                    var v11 uint16 = 0
                    _ = v11
                default:
                    if true {
                        if (8 != 253) {
                            p1 -= 49
                            v12 := GenS1{F0: 566, F1: 2, F2: 566}
                            _ = v12
                        } else {
                            fmt.Println(fi6, fi9)
                            var v13 bool = true
                            if ((-(p2)) != p2) {
                                var v14 uint8 = 127
                                _ = v14
                            } else {
                                v13 = ((p2 & p2) >= 1)
                                v15 := GenS1{F0: 65535, F1: 199, F2: 65535}
                                var v16 int64 = 42
                                _ = v15
                                _ = v16
                            }
                            _ = v13
                        }
                        var v17 uint = 551
                        if false {
                            p2 = (p2 % ((int8(v17)) | 1))
                            v17 -= v17
                            v18 := []int{fi9}
                            v19 := GenS1{F0: (^(v17)), F1: 42, F2: 2}
                            _ = v18
                            _ = v19
                        } else {
                            var v20 int8 = p2
                            _ = v20
                        }
                        _ = v17
                    }
                    var v21 int8 = p2
                    if (fi9 != fi9) {
                        v22 := GenS2{F0: int32(p1)}
                        var v23 int64 = p0
                        _ = v22
                        _ = v23
                    }
                    _ = v21
                }
                p2 = (p2 - int8(p1))
                var v24 int32 = -2147483648
                _ = fi9
                _ = v24
            }
            _ = fi6
        }
    }
    var v25 int64 = (p0 % ((p0) | 1))
    v26 := []int{-1, 726, 315, 9223372036854775807}
    var v27 string = "hello"
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    _ = v26
    _ = v27
    return 295
}
func genFn1() int16 {
    var v28 int32 = 42
    v28 += v28
    v29 := []int{32768, 907}
    switch ("" + "") {
    case "a":
        v30 := GenS3{F0: -466, F1: 9223372036854775807}
        _ = v30
    case "hello":
        fmt.Println(v28)
    }
    for fi31 := 0; fi31 < 4; fi31++ {
        if false {
            for fi32 := 0; fi32 < 7; fi32++ {
                var v33 string = "\n\t"
                var v34 bool = false
                _ = fi32
                _ = v33
                _ = v34
            }
            _ = len(v29)
        }
        _ = fi31
    }
    _ = v28
    _ = v29
    return 225
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 128
    var v3 string = "hello"
    v4 := []uint32{}
    var v5 int64 = 9223372036854775807
    for fi35 := 0; fi35 < 6; fi35++ {
        fmt.Println(v3, v3, v5)
        v2 = (v2 / ((v2) | 1))
        _ = fi35
    }
    switch (v3 + v3) {
    case "":
        if (!(false && ((1 != 127) && false))) {
            switch v3 {
            case "\n\t":
                var v36 uint64 = 343
                _ = v36
            case "a":
                v3 = v3
            case "":
                v1[4294967295] = ("" + (v3 + v3))
            }
        } else {
            switch 0 {
            case 1:
                fmt.Println(v2, v5, v2)
            case 128:
                v2 = v2
            case -2147483648:
                if true {
                    var v37 uint8 = (v2 + (v2 | (v2 + 127)))
                    v1[65535] = ("\n\t" + v3)
                    _ = v37
                } else {
                    v38 := GenS1{F0: 2, F1: 1, F2: genFn0(-1, 0, -125)}
                    _ = v38
                }
            }
            fmt.Println(v5, v2, v2)
        }
    }
    switch "\n\t" {
    case "\n\t":
        delete(v1, 536)
    case "":
        var v39 int16 = -613
        _ = v39
    case "a":
        var v40 int32 = 249
        _ = v40
    }
    v41 := GenS1{F0: 0, F1: 0, F2: 2147483647}
    fmt.Println(v5, v2)
    v42 := GenS1{F0: uint(v2), F1: 1, F2: 4294967295}
    v3 = "a"
    var v43 float64 = (0.0 * 1e10)
    v44 := GenS4{F0: v2}
    if (false && (42 > 0)) {
        if (0 != 773) {
            fmt.Println(v3, v2)
            v45 := GenS5{F0: genFn1()}
            _ = v45
        }
        fmt.Println(v43, v2)
        v46 := GenS5{F0: 1000}
        _ = v46
    }
    var v47 float32 = -2.25
    v5 -= v5
    if len(v4) > 0 {
        v4[0] = 65535
    }
    for fi48 := 0; fi48 < 18; fi48++ {
        v49 := v1[725]
        v50 := GenS5{F0: genFn1()}
        v5 = (v5 << ((42) & 63))
        _ = fi48
        _ = v49
        _ = v50
    }
    if true {
        v51 := GenS6{F0: false, F1: 1, F2: 42}
        v43 += (v43 - (v43 * -2.25))
        _ = v51
    } else {
        v1[2] = "\n\t"
        v2 = 2
        v5 -= (-(v5))
    }
    v52 := GenS7{F0: v2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v47
    _ = v5
    _ = v52
}
