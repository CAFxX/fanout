package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 bool }
type GenS2 struct { F0 uint8; F1 int64; F2 uint8 }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 uint16; F1 uint64; F2 uint64 }
func genFn0() int32 {
    var v6 uint8 = 128
    for fi7 := 0; fi7 < 19; fi7++ {
        for fi8 := 0; fi8 < 1; fi8++ {
            v6 -= (v6 * v6)
            _ = fi8
        }
        if (1.5 >= float64(1e10)) {
            if true {
                switch "a" {
                case "a":
                    v6 = (v6 & v6)
                default:
                    v9 := []int{fi7, (-(fi7))}
                    _ = v9
                }
            } else {
                v6 = v6
                switch fi7 {
                case 175:
                    fmt.Println(fi7, v6)
                case -1:
                    v10 := []int{fi7, (fi7 * fi7), fi7}
                    _ = v10
                case 328:
                    fmt.Println(fi7, fi7)
                }
                if ((v6 == v6) || false) {
                    if (-1 == 0) {
                        for fi11 := 0; fi11 < 2; fi11++ {
                            v6 += 42
                            var v12 int8 = -128
                            _ = fi11
                            _ = v12
                        }
                    } else {
                        if false {
                            fmt.Println(fi7, v6)
                            switch "foo bar" {
                            case "":
                                v6 = v6
                            }
                        }
                    }
                }
            }
            var v13 uint8 = v6
            _ = v13
        }
        _ = fi7
    }
    v6 -= (v6 ^ 128)
    _ = v6
    return -2147483648
}
func genFn1(p0 uint64, p1 uint32, p2 int) uint8 {
    v14 := GenS1{F0: -67, F1: 48, F2: false}
    p2 += (^(p2))
    p0 -= 4294967295
    p0 -= 378
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 96
}

func main() {
    var v1 int8 = 100
    v2 := make(map[uint32]int)
    var v3 uint16 = 65535
    var v4 uint32 = 4294967295
    var v5 float32 = -2.25
    var v15 int = 128
    v4 = v4
    fmt.Println(v5)
    v16 := []int{(^(v15))}
    v16 = append(v16, (v15 << ((255) & 63)))
    for fi17 := 0; fi17 < 4; fi17++ {
        if (true && (-1 <= 42)) {
            v18 := GenS2{F0: uint8(v4), F1: 9223372036854775807, F2: genFn1(4294967295, (v4 - v4), -9223372036854775808)}
            var v19 uint8 = 128
            _ = v18
            _ = v19
        } else {
            for fi20 := 0; fi20 < 19; fi20++ {
                var v21 uint16 = 2
                v16 = append(v16, (fi17 ^ -415))
                _ = fi20
                _ = v21
            }
            v1 = int8(v4)
        }
        _ = fi17
    }
    switch "hello" {
    case "hello":
        for fi22 := 0; fi22 < 18; fi22++ {
            v2[2147483648] = (v15 & 127)
            switch (v15 & (fi22 | (v15 << ((657) & 63)))) {
            case 32768:
                _ = len(v16)
            case 127:
                for fi23 := 0; fi23 < 6; fi23++ {
                    v24 := GenS2{F0: 146, F1: -9223372036854775808, F2: 255}
                    _ = fi23
                    _ = v24
                }
            case 128:
                v25 := GenS3{F0: 255, F1: v1}
                _ = v25
            }
            v5 = -2.25
            _ = fi22
        }
    }
    v26 := GenS3{F0: 42, F1: 42}
    var v27 uint32 = 2
    fmt.Println(v4)
    fmt.Println(v4, v15)
    if false {
        fmt.Println(v3, v5)
        v28 := GenS4{F0: 65535, F1: 781, F2: 18446744073709551615}
        _ = v28
    }
    if ((v3 | 2) >= (v3 + uint16(v4))) {
        v29 := GenS3{F0: 42, F1: v1}
        _ = v29
    }
    v30 := GenS2{F0: 128, F1: -1, F2: 1}
    for fi31 := 0; fi31 < 7; fi31++ {
        v32 := GenS2{F0: 145, F1: 0, F2: 127}
        _ = fi31
        _ = v32
    }
    v33 := GenS1{F0: -323, F1: (v1 + (v1 ^ (v1 % ((int8(v15)) | 1)))), F2: ("\n\t" < "foo bar")}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v33
    _ = v4
    _ = v5
}
