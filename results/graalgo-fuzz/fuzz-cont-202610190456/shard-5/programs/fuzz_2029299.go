package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 uint; F2 int32 }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 uint64; F1 int64 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 int) int64 {
    var v5 uint64 = 1
    v5 += 4294967295
    _ = p0
    _ = v5
    return -9223372036854775808
}

func main() {
    var v1 int64 = -59
    var v2 int64 = (^(v1))
    v3 := make(map[int]int)
    var v4 float32 = 0.0
    fmt.Println(v4, v2, v1)
    if (2 < 42) {
        var v6 uint32 = 2147483647
        switch ("hello" + ("a" + "a")) {
        case "\n\t":
            v7 := GenS1{F0: (v1 << ((0) & 63))}
            _ = v7
        }
        if (true || (!((242 < 0) || true))) {
            v6 = (v6 + 381)
        }
        _ = v6
    } else {
        switch -935 {
        case -1:
            v3[-504] = 128
        case 2147483647:
            v1 = (v1 >> ((40) & 63))
        default:
            if false {
                v1 = (v1 ^ v2)
                switch -122 {
                case 523:
                    for fi8 := 0; fi8 < 15; fi8++ {
                        var v9 uint16 = 2
                        fmt.Println(v9)
                        v10 := GenS1{F0: v1}
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                }
            }
        }
    }
    fmt.Println(v4, v4)
    switch -950 {
    case -2147483648:
        if false {
            v1 = 241
            v11 := []int{0}
            _ = v11
        }
    case 32767:
        if false {
            v12 := GenS2{F0: 42, F1: 1, F2: 35}
            if false {
                v1 = (v1 + 9223372036854775807)
                switch 42 {
                case 830:
                    if (false || (3.14159 <= 3.14159)) {
                        v1 = (^(v2))
                    }
                case 128:
                    v13 := GenS2{F0: 2, F1: 2, F2: -2147483648}
                    _ = v13
                case -741:
                    switch "hello" {
                    case "foo bar":
                        if true {
                            switch ("hello" + "hello") {
                            case "\n\t":
                                var v14 int32 = 1
                                _ = v14
                            case "a":
                                v2 += (v2 >> ((42) & 63))
                            default:
                                v15 := GenS3{F0: 2, F1: 42}
                                _ = v15
                            }
                            v4 = v4
                        }
                    }
                }
            }
            v16 := v3[-2147483648]
            _ = v12
            _ = v16
        } else {
            v4 += (1e10 - v4)
        }
    default:
        var v17 int64 = (v2 + v1)
        _ = v17
    }
    var v18 float32 = (v4 * (v4 / 1.0))
    v19 := GenS4{F0: 0, F1: (v1 >> ((255) & 63))}
    v4 += v18
    v20 := GenS5{F0: 127}
    v21 := []int{-9223372036854775808, 9223372036854775807, 870}
    var v22 uint8 = 191
    fmt.Println(v1, v4)
    v2 -= genFn0(9223372036854775807)
    v23 := GenS3{F0: 686, F1: 42}
    if true {
        var v24 float64 = (float64(3.14159) / -2.25)
        _ = v24
    } else {
        v2 += (v1 % ((-9223372036854775808) | 1))
    }
    delete(v3, 128)
    for fi25 := 0; fi25 < 8; fi25++ {
        switch (-(fi25)) {
        case -399:
            v2 = 1
        case 910:
            var v26 uint16 = 0
            _ = v26
        default:
            fmt.Println(v4)
        }
        _ = fi25
    }
    fmt.Println(v18, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
