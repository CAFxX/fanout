package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 uint32; F1 uint8; F2 int32 }
func genFn0(p0 uint, p1 int) uint {
    p1 = 705
    if (!(2 >= 65535)) {
        v7 := []int{p1, p1, p1, -469}
        _ = v7
    }
    var v8 float32 = 1.5
    p1 += p1
    _ = p0
    _ = p1
    _ = v8
    return 18446744073709551615
}
func genFn1(p0 int, p1 uint16, p2 int) int64 {
    p1 -= p1
    switch "hello" {
    case "":
        for fi9 := 0; fi9 < 8; fi9++ {
            p0 = int(p1)
            v10 := []int{p0, p2, fi9}
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            _ = fi9
            _ = v10
        }
    case "foo bar":
        v11 := []int{(^(p2)), p0, (p0 + (p2 - -36)), (p0 ^ (p0 ^ p0))}
        _ = v11
    case "a":
        p0 += int(p1)
    default:
        p1 = (p1 % ((32767) | 1))
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 float32 = 1.5
    var v2 uint32 = 42
    v3 := make(map[uint32]bool)
    v4 := make(map[int]int)
    var v5 int8 = 70
    var v6 uint = 2
    if (1 >= 983) {
        v12 := []int{9223372036854775807, 389}
        _ = v12
    }
    v1 = (v1 / (v1 - 3.14159))
    v2 += 65535
    var v13 uint8 = 197
    switch -2147483648 {
    case -2147483648:
        var v14 string = "hello"
        _ = v14
    default:
        fmt.Println(v6, v6, v1)
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        if false {
            fmt.Println(v2, fi15, v13)
            var v16 int16 = -387
            _ = v16
        }
        _ = fi15
    }
    fmt.Println(v6, v13, v1)
    var v17 uint16 = 419
    switch -77 {
    case -838:
        v13 = 1
    default:
        switch -2147483648 {
        case -2147483648:
            if ((!(-482 < 2147483647)) && false) {
                if false {
                    for fi18 := 0; fi18 < 1; fi18++ {
                        v19 := GenS1{F0: (fi18 >= (^(fi18))), F1: (fi18 > (fi18 / (((fi18 << (((^(v6))) & 63))) | 1)))}
                        _ = fi18
                        _ = v19
                    }
                } else {
                    fmt.Println(v5)
                }
                delete(v4, -9223372036854775808)
            }
        }
    }
    var v20 uint32 = (v2 << ((42) & 63))
    for fi21 := 0; fi21 < 14; fi21++ {
        v13 = v13
        _ = fi21
    }
    var v22 float64 = 1.5
    for fi23 := 0; fi23 < 16; fi23++ {
        fmt.Println(v2, v5, v22)
        v24 := []int{fi23}
        var v25 uint32 = 2147483647
        _ = fi23
        _ = v24
        _ = v25
    }
    for fi26 := 0; fi26 < 12; fi26++ {
        var v27 uint16 = (^(v17))
        fmt.Println(v2)
        switch "\n\t" {
        case "hello":
            v28 := GenS1{F0: false, F1: false}
            _ = v28
        default:
            v29 := GenS2{F0: (v20 - uint32(v6)), F1: (^(v13)), F2: 2147483647}
            _ = v29
        }
        _ = fi26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
