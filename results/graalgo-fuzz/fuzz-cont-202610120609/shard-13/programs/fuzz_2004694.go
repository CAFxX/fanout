package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint; F1 uint32; F2 uint64 }
type GenS5 struct { F0 int16; F1 int8 }
func genFn0(p0 uint8, p1 int, p2 uint8) int32 {
    v7 := GenS1{F0: 42, F1: 4294967295, F2: p1}
    for fi8 := 0; fi8 < 5; fi8++ {
        if (202 == -396) {
            for fi9 := 0; fi9 < 11; fi9++ {
                p2 = (^(p2))
                var v10 int8 = -128
                if true {
                    v11 := []int{fi9, 708, (fi9 << ((18446744073709551615) & 63))}
                    if ((3.14159 != -2.25) && (42 > uint32(p0))) {
                        switch -9223372036854775808 {
                        case 128:
                            v12 := GenS1{F0: 1, F1: 18446744073709551615, F2: 32768}
                            _ = v12
                        case 2147483647:
                            var v13 int64 = -805
                            _ = v13
                        case -244:
                            fmt.Println(p0, fi9)
                        }
                    } else {
                        v14 := GenS1{F0: 1, F1: 4294967295, F2: int(v10)}
                        var v15 uint32 = 4294967295
                        v11 = append(v11, -1)
                        _ = v14
                        _ = v15
                    }
                    fmt.Println(p0, fi8, v10)
                    if (!false) {
                        var v16 float32 = ((-2.25 - (1e10 - -2.25)) + 0.0)
                        v17 := GenS1{F0: 223, F1: 96, F2: p1}
                        if (466 >= 65535) {
                            if true {
                                if (2 <= 42) {
                                    if false {
                                        v16 = v16
                                    } else {
                                        p0 = (p0 << ((65535) & 63))
                                    }
                                    var v18 int64 = 42
                                    _ = v18
                                }
                            } else {
                                if len(v11) > 0 {
                                    v11[0] = (^(p1))
                                }
                                p1 -= (fi8 >> ((925) & 63))
                                p1 = (^(fi8))
                            }
                        } else {
                            var v19 string = "a"
                            _ = v19
                        }
                        _ = v16
                        _ = v17
                    } else {
                        if len(v11) > 0 {
                            v11[0] = p1
                        }
                    }
                    p2 = (p0 - 114)
                    switch (fi9 << ((65535) & 63)) {
                    case 839:
                        switch p1 {
                        case 2147483647:
                            v20 := GenS1{F0: 0, F1: 4294967295, F2: fi9}
                            _ = v20
                        default:
                            p0 = (p2 & 127)
                        }
                    case -1:
                        var v21 uint8 = p2
                        _ = v21
                    }
                    _ = v11
                }
                _ = fi9
                _ = v10
            }
            v22 := GenS2{F0: (p2 + p2)}
            _ = v22
        } else {
            var v23 int8 = 37
            p1 -= 744
            switch "a" {
            case "":
                v24 := GenS2{F0: p0}
                _ = v24
            }
            _ = v23
        }
        p1 += (fi8 + (p1 << ((42) & 63)))
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 499
}
func genFn1(p0 int64, p1 int32) uint16 {
    for fi25 := 0; fi25 < 15; fi25++ {
        v26 := []int{fi25, (fi25 * -2147483648), fi25}
        switch (fi25 ^ -395) {
        case 78:
            for fi27 := 0; fi27 < 5; fi27++ {
                p0 = p0
                fmt.Println(fi25, fi27, fi27)
                v26 = append(v26, (fi25 % ((32768) | 1)))
                _ = fi27
            }
        case 32768:
            v28 := GenS1{F0: 0, F1: 1, F2: 37}
            _ = v28
        case 128:
            var v29 string = "\n\t"
            _ = v29
        }
        _ = fi25
        _ = v26
    }
    v30 := GenS3{F0: 128}
    p1 = p1
    _ = p0
    _ = p1
    _ = v30
    return 65535
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 float64 = -2.25
    var v3 float32 = 1.5
    var v4 uint64 = (v1 + (v1 ^ v1))
    v5 := make(map[string]int)
    v6 := make(map[uint32]string)
    v3 = float32((1e10 / v2))
    var v31 float64 = (v2 / v2)
    v6[65535] = "hello"
    v5[("a" + "hello")] = -2147483648
    var v32 float32 = ((1e10 + v3) + v3)
    var v33 int64 = 1
    fmt.Println(v1)
    v33 -= -108
    var v34 uint8 = 255
    switch "" {
    case "\n\t":
        v2 = 0.1
    }
    v35 := GenS3{F0: v34}
    if (((1000 > int16(v33)) && false) && true) {
        v36 := GenS4{F0: 302, F1: 42, F2: v4}
        var v37 uint = 1
        _ = v36
        _ = v37
    }
    for fi38 := 0; fi38 < 18; fi38++ {
        v39 := GenS5{F0: -32768, F1: 1}
        var v40 int8 = -1
        v3 -= 3.14159
        _ = fi38
        _ = v39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
