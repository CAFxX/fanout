package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32; F1 uint8 }
type GenS4 struct { F0 uint8; F1 int8 }
type GenS5 struct { F0 int16; F1 uint16; F2 int32 }
type GenS6 struct { F0 string; F1 int16; F2 uint64 }
func genFn0() uint {
    v4 := GenS1{F0: 1, F1: 964}
    var v5 float32 = (-2.25 + (3.14159 / 1.5))
    _ = v4
    _ = v5
    return 647
}
func genFn1(p0 int, p1 int32, p2 uint8) int16 {
    v6 := GenS1{F0: 2, F1: 506}
    for fi7 := 0; fi7 < 6; fi7++ {
        if false {
            p0 = -1
        } else {
            for fi8 := 0; fi8 < 16; fi8++ {
                if false {
                    v9 := GenS1{F0: 2, F1: 0}
                    _ = v9
                } else {
                    fmt.Println(fi7, p0, p0)
                    v10 := []int{(^(fi8)), 128, p0, (fi7 | 32767)}
                    v11 := GenS2{F0: (p0 % (((fi7 & -1)) | 1))}
                    _ = v10
                    _ = v11
                }
                _ = fi8
            }
            if true {
                v12 := GenS1{F0: 623, F1: 597}
                var v13 uint16 = 65535
                switch (p0 - 127) {
                case 127:
                    if (true && false) {
                        var v14 uint = 827
                        var v15 float32 = 0.0
                        var v16 float64 = 1.5
                        _ = v14
                        _ = v15
                        _ = v16
                    } else {
                        for fi17 := 0; fi17 < 3; fi17++ {
                            if true {
                                fmt.Println(v13, fi7)
                                var v18 int16 = 42
                                _ = v18
                            }
                            p2 = (p2 + (^(p2)))
                            var v19 float32 = 0.0
                            _ = fi17
                            _ = v19
                        }
                        switch (^(fi7)) {
                        case 327:
                            var v20 uint16 = (v13 >> ((genFn0()) & 63))
                            _ = v20
                        case 2147483647:
                            var v21 uint8 = (p2 * p2)
                            _ = v21
                        case -84:
                            var v22 uint32 = 2
                            _ = v22
                        }
                        switch "\n\t" {
                        case "hello":
                            p2 += 128
                        }
                    }
                default:
                    var v23 int16 = 1000
                    _ = v23
                }
                _ = v12
                _ = v13
            }
        }
        _ = fi7
    }
    var v24 int = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    _ = v6
    return 1
}

func main() {
    var v1 bool = (35 < 1)
    v2 := make(map[int]bool)
    var v3 int16 = 0
    var v25 uint = 1
    for fi26 := 0; fi26 < 2; fi26++ {
        v27 := v2[(fi26 ^ (fi26 >> ((v25) & 63)))]
        fmt.Println(fi26, v25)
        if v27 {
            v28 := []int{fi26}
            _ = v28
        } else {
            v29 := GenS2{F0: fi26}
            _ = v29
        }
        _ = fi26
        _ = v27
    }
    for fi30 := 0; fi30 < 19; fi30++ {
        v3 = -417
        if v1 {
            v1 = ((724 == 890) && (19 < 120))
            v25 = v25
            v3 = genFn1((fi30 * fi30), 42, 118)
        } else {
            fmt.Println(v25)
            if (179 <= 42) {
                v1 = ((3.14159 + -2.25) <= 0.1)
            }
            var v31 int8 = 0
            _ = v31
        }
        _ = fi30
    }
    var v32 string = ""
    v3 += (v3 - (v3 / (((v3 | v3)) | 1)))
    v32 = ("a" + v32)
    switch "\n\t" {
    case "foo bar":
        v33 := GenS3{F0: 2, F1: 42}
        _ = v33
    case "\n\t":
        v1 = v1
    case "":
        v32 = (v32 + v32)
    default:
        switch 128 {
        case -2147483648:
            v34 := GenS2{F0: 1}
            _ = v34
        case -1:
            var v35 int16 = v3
            _ = v35
        }
    }
    if v1 {
        v36 := GenS4{F0: 128, F1: 1}
        v37 := GenS3{F0: 2147483648, F1: 2}
        fmt.Println(v32, v1, v3)
        _ = v36
        _ = v37
    }
    v38 := GenS2{F0: -9223372036854775808}
    if (-1 != 601) {
        v39 := []int{9223372036854775807, 128, 230}
        v39 = append(v39, 128)
        _ = v39
    } else {
        v1 = v1
    }
    fmt.Println(v3, v32)
    v32 = ("\n\t" + (v32 + v32))
    if false {
        fmt.Println(v3, v3)
    }
    fmt.Println(v32, v32, v1)
    for fi40 := 0; fi40 < 5; fi40++ {
        for fi41 := 0; fi41 < 17; fi41++ {
            v2[(^(fi40))] = (!v1)
            _ = fi41
        }
        switch "foo bar" {
        case "foo bar":
            v42 := []int{(^(fi40))}
            _ = v42
        default:
            v1 = false
        }
        _ = fi40
    }
    v25 += 412
    switch 824 {
    case -9223372036854775808:
        v43 := GenS5{F0: genFn1(4294967295, 42, 116), F1: 487, F2: 0}
        _ = v43
    case 4294967295:
        v44 := GenS6{F0: "\n\t", F1: v3, F2: 1}
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v32
    _ = v38
}
