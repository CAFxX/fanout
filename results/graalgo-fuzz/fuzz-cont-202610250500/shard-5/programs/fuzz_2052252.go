package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 uint64; F1 uint64 }
type GenS5 struct { F0 uint16; F1 uint16 }
type GenS6 struct { F0 int16; F1 int; F2 bool }
type GenS7 struct { F0 uint16; F1 uint16 }
func genFn0(p0 int16, p1 uint16) uint {
    for fi4 := 0; fi4 < 12; fi4++ {
        v5 := []int{}
        v6 := GenS1{F0: 1}
        if (2147483648 == 2) {
            if (2 > 2) {
                switch (fi4 + (-(fi4))) {
                case -9223372036854775808:
                    v5 = append(v5, (fi4 % ((fi4) | 1)))
                default:
                    var v7 string = "hello"
                    _ = v7
                }
                if (42 != 610) {
                    fmt.Println(p0, p0)
                    if (!false) {
                        var v8 float64 = float64((0.1 / float32(1.5)))
                        _ = v8
                    } else {
                        if (18446744073709551615 == 1) {
                            for fi9 := 0; fi9 < 20; fi9++ {
                                if ((p1 & (^(p1))) >= 65535) {
                                    p0 = p0
                                    p1 = p1
                                    v10 := GenS2{F0: 963}
                                    _ = v10
                                }
                                _ = fi9
                            }
                            if (true || false) {
                                v11 := GenS3{F0: 0, F1: 2147483648}
                                _ = v11
                            }
                            for fi12 := 0; fi12 < 14; fi12++ {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                p0 = (^(p0))
                                p1 -= p1
                                _ = fi12
                            }
                        }
                    }
                }
            }
            var v13 int64 = 9223372036854775807
            _ = v13
        }
        _ = fi4
        _ = v5
        _ = v6
    }
    p1 -= 2
    var v14 uint = 42
    for fi15 := 0; fi15 < 12; fi15++ {
        if (!false) {
            for fi16 := 0; fi16 < 6; fi16++ {
                fmt.Println(fi15)
                if (!(false && true)) {
                    p0 -= (^(p0))
                    if (p1 != (p1 >> (((^(v14))) & 63))) {
                        p1 += (p1 & 32767)
                        p1 = (p1 * p1)
                        v17 := GenS1{F0: 255}
                        _ = v17
                    }
                }
                _ = fi16
            }
            p0 += (-(p0))
            if (false && ((p1 - p1) != p1)) {
                for fi18 := 0; fi18 < 10; fi18++ {
                    var v19 string = "\n\t"
                    v14 -= v14
                    var v20 int64 = 0
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                var v21 int32 = int32(p1)
                for fi22 := 0; fi22 < 16; fi22++ {
                    fmt.Println(p1, p1, v14)
                    v14 = v14
                    _ = fi22
                }
                _ = v21
            } else {
                var v23 uint32 = 2147483647
                v24 := []int{32768, -734}
                p1 = 2
                _ = v23
                _ = v24
            }
        } else {
            v25 := []int{fi15, (-(fi15)), (fi15 % ((128) | 1)), 128}
            v14 -= v14
            var v26 uint = uint(p1)
            _ = v25
            _ = v26
        }
        for fi27 := 0; fi27 < 17; fi27++ {
            var v28 int8 = -90
            v14 = 364
            _ = fi27
            _ = v28
        }
        _ = fi15
    }
    v14 = v14
    _ = p0
    _ = p1
    _ = v14
    return (^(v14))
}
func genFn1(p0 int16) uint {
    v29 := []int{-1, -9223372036854775808}
    v30 := GenS4{F0: 400, F1: 2}
    _ = p0
    _ = v29
    _ = v30
    return 65535
}

func main() {
    var v1 int64 = 42
    var v2 uint32 = 2147483647
    v3 := make(map[uint32]uint32)
    var v31 uint64 = 42
    for fi32 := 0; fi32 < 10; fi32++ {
        if (genFn0(-717, 42) <= genFn0(-32768, 42)) {
            if (128 >= 183) {
                v2 = v2
            } else {
                v1 = (v1 & (v1 ^ (v1 * v1)))
                v33 := GenS4{F0: 0, F1: v31}
                for fi34 := 0; fi34 < 3; fi34++ {
                    v35 := GenS2{F0: v31}
                    switch 2147483647 {
                    case -399:
                        switch fi34 {
                        case -910:
                            if false {
                                v1 += int64(fi32)
                                v36 := GenS5{F0: 850, F1: 252}
                                v31 = (v31 ^ v31)
                                v37 := []int{fi32, (fi32 << ((42) & 63))}
                                _ = v36
                                _ = v37
                            }
                        }
                    case -2147483648:
                        v38 := GenS6{F0: -1, F1: 125, F2: ((563 == (v1 << ((42) & 63))) || (v2 >= v2))}
                        _ = v38
                    }
                    _ = fi34
                    _ = v35
                }
                _ = v33
            }
            var v39 int64 = v1
            _ = v39
        }
        if true {
            var v40 float32 = float32((1.5 * 0.1))
            _ = v40
        }
        _ = fi32
    }
    v41 := GenS7{F0: 0, F1: 32767}
    fmt.Println(v31)
    v1 += 9223372036854775807
    fmt.Println(v31, v2, v2)
    v42 := GenS5{F0: 784, F1: 65535}
    v1 -= (v1 + 9223372036854775807)
    for fi43 := 0; fi43 < 3; fi43++ {
        v31 = 159
        switch "" {
        case "hello":
            var v44 uint = 4294967295
            _ = v44
        default:
            v45 := GenS1{F0: 42}
            _ = v45
        }
        fmt.Println(v2)
        _ = fi43
    }
    var v46 int64 = v1
    var v47 float32 = -2.25
    var v48 string = ""
    v1 = (v1 / (((v46 ^ (v46 / ((v46) | 1)))) | 1))
    var v49 int = 9223372036854775807
    var v50 float64 = ((float64(((1e10 / 1.5) / 1.5)) / 1.5) + (0.1 + (float64((1e10 * 1.5)) - 3.14159)))
    v2 = (v2 >> ((378) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v41
    _ = v42
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v50
}
