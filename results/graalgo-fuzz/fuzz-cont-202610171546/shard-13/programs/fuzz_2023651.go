package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 uint64; F1 uint64; F2 int8 }
type GenS3 struct { F0 string; F1 uint16 }
func genFn0() uint {
    v5 := GenS1{F0: 65535, F1: 621}
    if (1 != 95) {
        var v6 int32 = 1
        _ = v6
    }
    _ = v5
    return 942
}

func main() {
    var v1 uint64 = 42
    v2 := []int{32768}
    var v3 int64 = -530
    var v4 float64 = 1e10
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v4 += 1.5
    if false {
        v1 = (v1 + v1)
        fmt.Println(v4)
        var v7 uint64 = (v1 + v1)
        _ = v7
    } else {
        switch "\n\t" {
        case "foo bar":
            fmt.Println(v3, v1, v1)
        case "a":
            if true {
                if (-101 < -128) {
                    switch 9223372036854775807 {
                    case 0:
                        v3 = v3
                    case 32767:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case 4294967295:
                        v3 += (v3 % ((v3) | 1))
                    default:
                        v1 += (v1 & 42)
                    }
                }
                for fi8 := 0; fi8 < 20; fi8++ {
                    v4 -= ((v4 + (3.14159 / -2.25)) - (1.5 * v4))
                    if len(v2) > 0 {
                        v2[0] = (fi8 << ((18446744073709551615) & 63))
                    }
                    if true {
                        var v9 float32 = (1.5 - 1.5)
                        fmt.Println(v3)
                        _ = v9
                    } else {
                        var v10 string = "hello"
                        var v11 int64 = 0
                        var v12 uint64 = 0
                        _ = v10
                        _ = v11
                        _ = v12
                    }
                    _ = fi8
                }
            } else {
                v13 := GenS1{F0: 0, F1: 790}
                switch -2147483648 {
                case 600:
                    if false {
                        switch 1 {
                        case 505:
                            v4 += float64(3.14159)
                        }
                        v4 = 1.5
                    } else {
                        var v14 int32 = 1
                        fmt.Println(v1)
                        _ = v14
                    }
                }
                _ = v13
            }
        }
        v15 := GenS1{F0: 32767, F1: 2}
        fmt.Println(v3)
        _ = v15
    }
    var v16 string = "a"
    v3 = 1
    v2 = append(v2, -1)
    v1 = 808
    v17 := GenS2{F0: v1, F1: v1, F2: 1}
    if (true || (42 >= 2147483647)) {
        v18 := GenS1{F0: 653, F1: 32767}
        _ = v18
    }
    switch v16 {
    case "a":
        for fi19 := 0; fi19 < 19; fi19++ {
            v20 := GenS2{F0: (v1 + (v1 & v1)), F1: 462, F2: 1}
            _ = fi19
            _ = v20
        }
    case "\n\t":
        fmt.Println(v1, v16)
    case "foo bar":
        v21 := GenS3{F0: "hello", F1: uint16(v1)}
        _ = v21
    default:
        var v22 int16 = 785
        _ = v22
    }
    var v23 int64 = (v3 | v3)
    v3 = v3
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
