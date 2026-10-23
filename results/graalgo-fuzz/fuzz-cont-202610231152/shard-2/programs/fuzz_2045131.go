package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 string }
type GenS2 struct { F0 bool; F1 uint64; F2 int8 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 uint = 65535
    var v2 uint32 = 553
    v3 := make(map[int]int)
    v4 := make(map[int]uint32)
    v5 := make(map[int]uint32)
    v6 := make(map[int]int)
    v7 := GenS1{F0: v1, F1: -1, F2: "hello"}
    for fi8 := 0; fi8 < 10; fi8++ {
        switch 4294967295 {
        case 1:
            var v9 uint = 42
            _ = v9
        default:
            var v10 bool = (v1 == 191)
            _ = v10
        }
        switch (fi8 << ((uint(v2)) & 63)) {
        case 2147483647:
            v11 := []int{fi8, (-(fi8))}
            _ = v11
        }
        _ = fi8
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 uint16 = 54
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        v15 := GenS2{F0: (true || (!true)), F1: 419, F2: 44}
        _ = fi14
        _ = v15
    }
    v2 = (v2 * v2)
    v1 -= v1
    var v16 int64 = 9223372036854775807
    switch 42 {
    case -753:
        v17 := GenS2{F0: (-1 != 42), F1: 4294967295, F2: 63}
        _ = v17
    }
    if false {
        fmt.Println(v16, v16)
        v18 := GenS3{F0: (v16 >> (((v1 % ((v1) | 1))) & 63))}
        _ = v18
    } else {
        for fi19 := 0; fi19 < 17; fi19++ {
            v20 := []int{fi19}
            _ = fi19
            _ = v20
        }
        v16 -= (-(v16))
        if (true && true) {
            if true {
                var v21 string = "hello"
                _ = v21
            }
            v3[219] = 42
        }
    }
    if true {
        fmt.Println(v1)
        switch 127 {
        case 507:
            if (0 > 255) {
                switch "\n\t" {
                case "":
                    fmt.Println(v16)
                case "a":
                    var v22 uint32 = v2
                    _ = v22
                default:
                    v23 := GenS1{F0: 1, F1: -2147483648, F2: ""}
                    _ = v23
                }
            } else {
                switch "" {
                case "":
                    fmt.Println(v1, v16, v2)
                default:
                    v2 -= 2147483647
                }
            }
        case 655:
            v16 = v16
        }
        v24 := []int{9223372036854775807, 901}
        _ = v24
    } else {
        if (true || false) {
            v25 := []int{4294967295}
            _ = v25
        } else {
            v26 := GenS3{F0: v16}
            _ = v26
        }
        v27 := []int{9223372036854775807, -562}
        var v28 bool = (false && (0.0 < 0.1))
        _ = v27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
