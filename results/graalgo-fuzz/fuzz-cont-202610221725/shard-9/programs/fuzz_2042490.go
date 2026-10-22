package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 uint64 }
type GenS2 struct { F0 int }
func genFn0(p0 string, p1 int8, p2 int32) int32 {
    var v7 int64 = int64(p1)
    for fi8 := 0; fi8 < 1; fi8++ {
        v9 := GenS1{F0: 4294967295, F1: 32767, F2: 762}
        v10 := []int{(fi8 & 42), (fi8 + 344), (fi8 * fi8), (fi8 & fi8)}
        if true {
            v11 := GenS1{F0: 42, F1: 1, F2: 295}
            _ = v11
        } else {
            var v12 uint64 = 2
            _ = v12
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -1
}

func main() {
    var v1 uint = 874
    var v2 int16 = 32767
    var v3 int = 321
    v4 := []int{v3, (v3 ^ (v3 % ((v3) | 1))), (v3 / (((v3 | 9223372036854775807)) | 1)), v3}
    var v5 uint64 = 18446744073709551615
    v6 := []int{318}
    var v13 int16 = (v2 << ((42) & 63))
    for fi14 := 0; fi14 < 13; fi14++ {
        if true {
            switch ("a" + "") {
            case "hello":
                var v15 int64 = 9223372036854775807
                _ = v15
            }
            v16 := GenS2{F0: fi14}
            v17 := GenS1{F0: (v1 & 42), F1: 863, F2: (v5 * 0)}
            _ = v16
            _ = v17
        }
        _ = fi14
    }
    fmt.Println(v2)
    v3 = (^(v3))
    var v18 uint = v1
    var v19 uint = 1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi20 := 0; fi20 < 3; fi20++ {
        switch "\n\t" {
        case "a":
            for fi21 := 0; fi21 < 10; fi21++ {
                if false {
                    if ((v5 * (v5 << ((v1) & 63))) <= 18446744073709551615) {
                        v22 := GenS1{F0: (v1 - 650), F1: 834, F2: (v5 & v5)}
                        _ = v22
                    }
                    v23 := GenS2{F0: (v3 << ((v18) & 63))}
                    _ = v23
                } else {
                    _ = len(v4)
                    if true {
                        var v24 int32 = genFn0("", 0, 2147483647)
                        _ = v24
                    } else {
                        if (-2.25 > 3.14159) {
                            v25 := GenS1{F0: v19, F1: 0, F2: v5}
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            fmt.Println(v13, fi21)
                            _ = v25
                        } else {
                            fmt.Println(v19, v5)
                        }
                    }
                    var v26 int8 = 55
                    var v27 uint16 = 32767
                    _ = v26
                    _ = v27
                }
                _ = fi21
            }
        default:
            v1 = (v1 | uint(v5))
        }
        _ = fi20
    }
    if true {
        var v28 uint16 = 703
        fmt.Println(v2)
        for fi29 := 0; fi29 < 2; fi29++ {
            if len(v6) > 0 {
                v6[0] = (v3 - fi29)
            }
            _ = fi29
        }
        _ = v28
    }
    for fi30 := 0; fi30 < 11; fi30++ {
        for fi31 := 0; fi31 < 8; fi31++ {
            if len(v4) > 0 {
                v4[0] = fi31
            }
            _ = fi31
        }
        _ = fi30
    }
    if (1 < 785) {
        if len(v4) > 0 {
            v4[0] = v3
        }
        v19 -= 18446744073709551615
    } else {
        switch "\n\t" {
        case "\n\t":
            v32 := GenS2{F0: v3}
            _ = v32
        case "a":
            if len(v4) > 0 {
                v4[0] = v3
            }
        case "":
            fmt.Println(v19, v13, v13)
        default:
            var v33 uint64 = v5
            _ = v33
        }
    }
    v2 -= v13
    v1 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
