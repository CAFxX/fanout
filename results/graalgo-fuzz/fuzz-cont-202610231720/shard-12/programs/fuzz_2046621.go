package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int8; F1 uint32 }
type GenS5 struct { F0 bool; F1 int; F2 uint32 }
type GenS6 struct { F0 uint32; F1 uint64; F2 uint }
func genFn0(p0 uint64, p1 int32) uint {
    v4 := []int{32768, -2147483648, 2147483647, -413}
    p1 = (p1 << ((334) & 63))
    _ = p0
    _ = p1
    _ = v4
    return uint(p0)
}
func genFn1(p0 uint, p1 int16) int64 {
    switch -1 {
    case -2147483648:
        fmt.Println(p0)
    case -313:
        for fi5 := 0; fi5 < 16; fi5++ {
            for fi6 := 0; fi6 < 4; fi6++ {
                var v7 uint8 = 2
                var v8 int32 = 2147483647
                _ = fi6
                _ = v7
                _ = v8
            }
            var v9 uint8 = 2
            _ = fi5
            _ = v9
        }
    default:
        p0 -= (p0 * p0)
    }
    v10 := []int{904}
    _ = p0
    _ = p1
    _ = v10
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float32 = ((1e10 * ((((3.14159 / 1.5) + (0.1 + 0.0)) * 0.1) / (float32(float64((0.0 / -2.25))) - 3.14159))) + (1e10 + 1.5))
    v3 := make(map[int]int)
    v2 -= 0.0
    v2 = (v2 * 3.14159)
    if false {
        v11 := GenS1{F0: 0}
        for fi12 := 0; fi12 < 13; fi12++ {
            fmt.Println(fi12, v2)
            var v13 int8 = -1
            var v14 int64 = genFn1(0, 830)
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = v11
    }
    v15 := []int{0}
    var v16 uint = 65535
    var v17 string = "a"
    v16 -= 1
    switch v17 {
    case "hello":
        var v18 float64 = 3.14159
        _ = v18
    case "":
        fmt.Println(v16, v16, v2)
    }
    if true {
        var v19 int32 = 442
        for fi20 := 0; fi20 < 12; fi20++ {
            delete(v3, 530)
            if ((796 < 1) && false) {
                v21 := GenS2{F0: fi20}
                _ = v21
            } else {
                v22 := GenS1{F0: 895}
                var v23 int16 = 1000
                if false {
                    v24 := GenS3{F0: "\n\t"}
                    for fi25 := 0; fi25 < 14; fi25++ {
                        v16 += genFn0(2, (v19 / ((-709) | 1)))
                        fmt.Println(fi20, v16)
                        _ = fi25
                    }
                    var v26 bool = ((127 <= 205) && true)
                    _ = v24
                    _ = v26
                }
                _ = v22
                _ = v23
            }
            _ = fi20
        }
        var v27 int64 = 42
        _ = v19
        _ = v27
    } else {
        if (-861 <= 9223372036854775807) {
            if true {
                for fi28 := 0; fi28 < 3; fi28++ {
                    if len(v15) > 0 {
                        v15[0] = (fi28 << (((v16 | v16)) & 63))
                    }
                    v2 += 0.1
                    for fi29 := 0; fi29 < 6; fi29++ {
                        switch v17 {
                        case "\n\t":
                            var v30 string = v17
                            _ = v30
                        case "a":
                            var v31 uint32 = 4294967295
                            _ = v31
                        }
                        if (false || ((("hello" + "hello") + v17) < v17)) {
                            switch "\n\t" {
                            case "\n\t":
                                delete(v3, fi28)
                            default:
                                var v32 uint8 = 249
                                _ = v32
                            }
                            v16 += (v16 % (((v16 % (((v16 | 4294967295)) | 1))) | 1))
                            if true {
                                if ((v17 == v17) && false) {
                                    if len(v15) > 0 {
                                        v15[0] = fi29
                                    }
                                    v33 := GenS3{F0: v17}
                                    _ = v33
                                } else {
                                    var v34 uint64 = 42
                                    fmt.Println(fi29)
                                    v35 := GenS4{F0: 1, F1: 516}
                                    _ = v34
                                    _ = v35
                                }
                                if false {
                                    var v36 uint = v16
                                    v16 = genFn0(18446744073709551615, 637)
                                    _ = v36
                                }
                            }
                            switch (fi29 - fi29) {
                            case -1:
                                fmt.Println(fi29, v2)
                            default:
                                if true {
                                    var v37 uint8 = 127
                                    var v38 bool = (42 != 42)
                                    if len(v15) > 0 {
                                        v15[0] = fi29
                                    }
                                    _ = v37
                                    _ = v38
                                }
                            }
                        } else {
                            v16 = v16
                        }
                        var v39 int16 = -27
                        v40 := GenS5{F0: (588 != 2147483647), F1: (-(fi29)), F2: 2}
                        _ = fi29
                        _ = v39
                        _ = v40
                    }
                    _ = fi28
                }
                if (!true) {
                    for fi41 := 0; fi41 < 18; fi41++ {
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                        _ = fi41
                    }
                    _ = len(v15)
                }
                v42 := GenS6{F0: 152, F1: 0, F2: (^(v16))}
                _ = v42
            } else {
                v2 = 1.5
            }
            v17 = ((v17 + (v17 + v17)) + v17)
        } else {
            v17 = v17
        }
    }
    fmt.Println(v2)
    var v43 uint8 = 10
    v17 = ""
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v43:", v43)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v43
}
