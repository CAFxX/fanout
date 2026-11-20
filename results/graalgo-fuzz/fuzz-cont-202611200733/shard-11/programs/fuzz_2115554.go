package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int; F1 uint64; F2 int64 }
type GenS3 struct { F0 uint32; F1 uint64; F2 bool }
type GenS4 struct { F0 uint32 }
func genFn0(p0 int, p1 uint16) uint64 {
    fmt.Println(p1)
    for fi4 := 0; fi4 < 11; fi4++ {
        p1 += p1
        switch fi4 {
        case -875:
            var v5 int32 = 1
            _ = v5
        case 9223372036854775807:
            for fi6 := 0; fi6 < 11; fi6++ {
                fmt.Println(fi4, p1, p1)
                if false {
                    var v7 int16 = 1000
                    var v8 uint8 = 62
                    _ = v7
                    _ = v8
                } else {
                    var v9 uint64 = 2
                    v10 := []int{-559, (fi6 % (((fi6 % ((p0) | 1))) | 1)), 42, (p0 * p0)}
                    _ = v10
                    _ = v9
                }
                _ = fi6
            }
        case 345:
            var v11 int = (^(fi4))
            _ = v11
        }
        p1 += (p1 | 0)
        _ = fi4
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        var v13 int = (p0 << ((2) & 63))
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int32 = -2147483648
    var v3 int64 = -9223372036854775808
    fmt.Println(v1, v2)
    v3 -= -759
    v14 := GenS1{F0: 42, F1: 2, F2: (v1 / ((v1) | 1))}
    v3 = (v3 * v3)
    v15 := GenS2{F0: 127, F1: (^(v1)), F2: v3}
    var v16 bool = (!true)
    var v17 int = -772
    var v18 int8 = 7
    if ((v3 + v3) == v3) {
        switch "hello" {
        case "hello":
            var v19 uint8 = 42
            _ = v19
        case "\n\t":
            var v20 uint64 = genFn0(v17, 794)
            _ = v20
        }
        v21 := []int{(v17 * (v17 * v17)), (v17 << ((18446744073709551615) & 63)), (v17 >> ((2) & 63))}
        _ = v21
    } else {
        v22 := []int{-953}
        switch v17 {
        case 817:
            var v23 uint = 255
            _ = v23
        case 128:
            switch "a" {
            case "\n\t":
                v24 := GenS2{F0: 42, F1: genFn0(v17, 179), F2: (v3 + (v3 | v3))}
                _ = v24
            case "a":
                if len(v22) >= 2 {
                    v22 = v22[1:len(v22)]
                }
            default:
                if len(v22) > 0 {
                    v22[0] = (v17 & (v17 & (v17 * v17)))
                }
            }
        default:
            fmt.Println(v2, v1)
        }
        if (!(v17 > (v17 ^ (v17 >> ((42) & 63))))) {
            if v16 {
                if v16 {
                    v16 = true
                    v25 := GenS3{F0: 565, F1: genFn0(v17, 255), F2: v16}
                    _ = v25
                } else {
                    for fi26 := 0; fi26 < 1; fi26++ {
                        if v16 {
                            if v16 {
                                _ = len(v22)
                                v27 := GenS2{F0: (fi26 | (^(fi26))), F1: (v1 ^ (v1 ^ v1)), F2: v3}
                                _ = v27
                            }
                        }
                        v16 = ((-9223372036854775808 == (^(v3))) || (int64(v1) <= (v3 << ((2) & 63))))
                        _ = fi26
                    }
                    var v28 bool = false
                    fmt.Println(v28)
                    _ = v28
                }
                v29 := GenS2{F0: (v17 - v17), F1: (v1 & 2), F2: v3}
                v22 = append(v22, (v17 - (v17 << ((255) & 63))))
                _ = v29
            } else {
                for fi30 := 0; fi30 < 11; fi30++ {
                    fmt.Println(fi30, v2)
                    _ = fi30
                }
            }
            var v31 bool = true
            _ = v31
        } else {
            v17 += v17
            v32 := GenS3{F0: 739, F1: 18446744073709551615, F2: true}
            _ = v32
        }
        _ = v22
    }
    v33 := []int{}
    if (v18 <= v18) {
        var v34 uint64 = v1
        v35 := GenS3{F0: 231, F1: 42, F2: false}
        _ = v34
        _ = v35
    } else {
        if (-1 >= 2147483647) {
            v18 += v18
        }
    }
    for fi36 := 0; fi36 < 14; fi36++ {
        v37 := GenS2{F0: v17, F1: v1, F2: (v3 * (-(v3)))}
        for fi38 := 0; fi38 < 10; fi38++ {
            v39 := GenS2{F0: 32768, F1: 384, F2: (v3 >> ((42) & 63))}
            v17 = (fi36 * (^(v17)))
            if (861 <= 2) {
                if v16 {
                    if len(v33) >= 2 {
                        v33 = v33[1:len(v33)]
                    }
                }
                v40 := GenS4{F0: 653}
                v41 := GenS3{F0: 1, F1: (v1 * (^(v1))), F2: v16}
                fmt.Println(fi36, v18)
                fmt.Println(v1)
                fmt.Println(v2, v2)
                switch (("foo bar" + "a") + "foo bar") {
                case "hello":
                    if (false && v16) {
                        v2 = (v2 ^ v2)
                        fmt.Println(v18, v1, v16)
                        var v42 uint64 = v1
                        _ = v42
                    } else {
                        if len(v33) > 0 {
                            v33[0] = 256
                        }
                        v33 = append(v33, fi36)
                    }
                case "foo bar":
                    v18 = int8(fi36)
                default:
                    v43 := GenS4{F0: 2147483647}
                    _ = v43
                }
                v33 = append(v33, fi36)
                _ = v40
                _ = v41
            }
            _ = fi38
            _ = v39
        }
        _ = fi36
        _ = v37
    }
    var v44 int = 32767
    fmt.Println(v3)
    var v45 uint16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v33
    _ = v44
    _ = v45
}
