package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 bool; F2 int }
type GenS3 struct { F0 bool; F1 int; F2 uint32 }
func genFn0(p0 int64) uint8 {
    p0 = p0
    if (!(((!(32767 == 1)) || false) && true)) {
        v7 := GenS1{F0: 42}
        _ = v7
    }
    fmt.Println(p0)
    for fi8 := 0; fi8 < 4; fi8++ {
        if (-2147483648 > -2147483648) {
            switch (^(fi8)) {
            case 232:
                p0 = (p0 & (p0 / (((p0 - (p0 | 9223372036854775807))) | 1)))
            case 454:
                var v9 string = ("" + "foo bar")
                _ = v9
            default:
                if true {
                    switch (^(fi8)) {
                    case 9223372036854775807:
                        for fi10 := 0; fi10 < 14; fi10++ {
                            var v11 bool = (!(!true))
                            _ = fi10
                            _ = v11
                        }
                    case 128:
                        v12 := []int{(fi8 / (((fi8 % ((9223372036854775807) | 1))) | 1)), (fi8 & (fi8 & fi8))}
                        _ = v12
                    case 32768:
                        switch "a" {
                        case "\n\t":
                            if false {
                                var v13 uint = 158
                                _ = v13
                            } else {
                                switch "\n\t" {
                                case "":
                                    var v14 int32 = -1
                                    _ = v14
                                case "foo bar":
                                    for fi15 := 0; fi15 < 15; fi15++ {
                                        p0 -= p0
                                        var v16 float64 = 0.0
                                        p0 = (p0 * p0)
                                        _ = fi15
                                        _ = v16
                                    }
                                default:
                                    if true {
                                        v17 := []int{601, fi8}
                                        v17 = append(v17, fi8)
                                        _ = v17
                                    } else {
                                        switch "hello" {
                                        case "\n\t":
                                            v18 := []int{693, fi8, fi8}
                                            _ = v18
                                        default:
                                            var v19 int16 = 42
                                            _ = v19
                                        }
                                        p0 -= (p0 | -1)
                                    }
                                }
                            }
                        case "foo bar":
                            var v20 uint8 = 1
                            _ = v20
                        default:
                            v21 := GenS1{F0: 32767}
                            _ = v21
                        }
                    }
                    var v22 bool = false
                    var v23 uint64 = 18446744073709551615
                    _ = v22
                    _ = v23
                }
            }
        } else {
            fmt.Println(fi8)
        }
        p0 = (p0 << ((2) & 63))
        for fi24 := 0; fi24 < 5; fi24++ {
            p0 += p0
            v25 := []int{}
            _ = fi24
            _ = v25
        }
        _ = fi8
    }
    _ = p0
    return 254
}

func main() {
    var v1 bool = (-195 >= -32768)
    var v2 int = -463
    v3 := make(map[int]string)
    var v4 string = (("hello" + ("a" + ("a" + "\n\t"))) + "")
    v5 := []uint32{2147483647, 4294967295, 793, 935}
    var v6 uint32 = 148
    v5 = append(v5, 2147483648)
    v6 = (^(v6))
    var v26 int = (v2 >> ((458) & 63))
    v2 -= v2
    var v27 float32 = 0.0
    if ((v4 != v4) && v1) {
        for fi28 := 0; fi28 < 6; fi28++ {
            for fi29 := 0; fi29 < 19; fi29++ {
                var v30 int = (v26 % ((128) | 1))
                _ = fi29
                _ = v30
            }
            for fi31 := 0; fi31 < 6; fi31++ {
                var v32 int64 = 188
                v33 := GenS1{F0: fi31}
                _ = fi31
                _ = v32
                _ = v33
            }
            _ = fi28
        }
    }
    if true {
        switch (v26 / (((v2 | v2)) | 1)) {
        case 9223372036854775807:
            v2 -= 9223372036854775807
        case 37:
            v27 += (1.5 - ((v27 + ((1.5 / (float32(0.0) / v27)) - ((0.0 + v27) * float32((0.1 * (-2.25 * 0.1)))))) * v27))
        case -9223372036854775808:
            v27 = v27
        }
        v1 = v1
        v1 = (!(18446744073709551615 < 42))
    } else {
        var v34 int8 = 21
        _ = v34
    }
    v35 := GenS1{F0: -762}
    for fi36 := 0; fi36 < 17; fi36++ {
        var v37 uint64 = 42
        _ = fi36
        _ = v37
    }
    fmt.Println(v6, v1, v6)
    v2 = v2
    if (42 <= 756) {
        v6 = 2147483648
        v2 -= 575
        _ = len(v3)
    }
    var v38 int = 0
    if v1 {
        _ = len(v5)
    }
    var v39 uint16 = 586
    var v40 float64 = 1.5
    fmt.Println(v39, v40, v38)
    if v1 {
        for fi41 := 0; fi41 < 4; fi41++ {
            v42 := GenS2{F0: 65535, F1: v1, F2: 2147483647}
            _ = fi41
            _ = v42
        }
        v43 := GenS3{F0: (!(float64(v27) == 1e10)), F1: v38, F2: v6}
        v44 := GenS3{F0: v1, F1: 4294967295, F2: (v6 >> ((0) & 63))}
        _ = v43
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v35
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
