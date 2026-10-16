package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64; F1 int64; F2 uint64 }
type GenS4 struct { F0 int32; F1 int64; F2 uint64 }
type GenS5 struct { F0 uint8; F1 uint16 }
func genFn0(p0 bool, p1 int8) int8 {
    if (206 <= 18446744073709551615) {
        for fi4 := 0; fi4 < 7; fi4++ {
            var v5 float64 = 3.14159
            fmt.Println(v5, p1, fi4)
            _ = fi4
            _ = v5
        }
    } else {
        for fi6 := 0; fi6 < 17; fi6++ {
            p0 = false
            for fi7 := 0; fi7 < 11; fi7++ {
                p0 = (0 != 255)
                if p0 {
                    v8 := GenS1{F0: p1, F1: 255, F2: -743}
                    var v9 int = (fi6 - 0)
                    switch fi6 {
                    case -259:
                        var v10 float32 = 1e10
                        _ = v10
                    case 4294967295:
                        if p0 {
                            v9 -= 2147483647
                            if (!p0) {
                                switch (fi6 ^ fi6) {
                                case 127:
                                    v11 := []int{(fi7 - (-(fi7)))}
                                    _ = v11
                                default:
                                    v12 := GenS1{F0: p1, F1: 255, F2: -1}
                                    v13 := GenS1{F0: p1, F1: 4294967295, F2: 42}
                                    _ = v12
                                    _ = v13
                                }
                                switch 42 {
                                case 32768:
                                    if true {
                                        var v14 bool = p0
                                        _ = v14
                                    }
                                case 499:
                                    if false {
                                        var v15 int = (fi6 - fi7)
                                        _ = v15
                                    }
                                }
                                var v16 uint64 = 2
                                v17 := []int{}
                                v9 += fi7
                                if (!((float64(0.1) != 3.14159) || (!p0))) {
                                    if p0 {
                                        v18 := GenS2{F0: "hello"}
                                        var v19 int32 = 0
                                        var v20 uint32 = 65535
                                        switch int(v20) {
                                        case 881:
                                            v21 := GenS3{F0: -9223372036854775808, F1: -9223372036854775808, F2: v16}
                                            _ = v21
                                        case 32767:
                                            v22 := GenS4{F0: v19, F1: 0, F2: v16}
                                            _ = v22
                                        default:
                                            v9 = fi7
                                        }
                                        _ = v18
                                        _ = v19
                                        _ = v20
                                    }
                                    var v23 uint = 514
                                    var v24 int64 = -299
                                    _ = v23
                                    _ = v24
                                }
                                fmt.Println(v9, fi6, v16)
                                _ = v16
                                _ = v17
                            } else {
                                fmt.Println(fi6, p0, fi7)
                                fmt.Println(p1)
                            }
                            var v25 bool = p0
                            _ = v25
                        } else {
                            p0 = p0
                            if (!true) {
                                if p0 {
                                    fmt.Println(fi6, fi6, fi6)
                                    v26 := []int{42, (v9 * (fi7 + v9)), fi7, 32768}
                                    var v27 uint32 = 0
                                    v9 = v9
                                    _ = v26
                                    _ = v27
                                }
                                v28 := []int{-880}
                                _ = v28
                            }
                        }
                    case 42:
                        fmt.Println(v9, fi6, p1)
                    default:
                        v9 -= (fi6 & (fi7 % ((v9) | 1)))
                    }
                    _ = v8
                    _ = v9
                }
                p0 = p0
                var v29 uint8 = 127
                v30 := []int{fi6, fi7}
                var v31 bool = (32767 == 0)
                v32 := GenS4{F0: 0, F1: 1, F2: 756}
                _ = fi7
                _ = v29
                _ = v30
                _ = v31
                _ = v32
            }
            if true {
                p0 = (2147483647 == 1)
            }
            _ = fi6
        }
    }
    v33 := []int{}
    _ = p0
    _ = p1
    _ = v33
    return p1
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]bool)
    v3 := make(map[int]string)
    var v34 uint8 = 128
    v34 -= (v34 << ((4294967295) & 63))
    fmt.Println(v34)
    fmt.Println(v34)
    var v35 uint32 = 2147483647
    if true {
        v34 = (v34 % (((v34 | (v34 >> ((625) & 63)))) | 1))
    }
    for fi36 := 0; fi36 < 9; fi36++ {
        v37 := GenS2{F0: "a"}
        _ = fi36
        _ = v37
    }
    if false {
        var v38 int16 = 0
        v39 := GenS2{F0: "a"}
        _ = v38
        _ = v39
    } else {
        v40 := GenS1{F0: 0, F1: 255, F2: -1}
        if false {
            var v41 int16 = 42
            _ = v41
        }
        _ = v40
    }
    v35 = (v35 & v35)
    v42 := GenS5{F0: v34, F1: 32767}
    var v43 uint8 = v34
    var v44 uint = 112
    var v45 uint8 = v43
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v42
    _ = v43
    _ = v44
    _ = v45
}
