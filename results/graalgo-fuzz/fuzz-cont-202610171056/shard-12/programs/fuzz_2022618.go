package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 int64 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 int; F1 int; F2 uint }
func genFn0(p0 uint8, p1 uint16) int32 {
    for fi4 := 0; fi4 < 7; fi4++ {
        var v5 float32 = (((((-2.25 - 3.14159) / -2.25) + float32(0.0)) - ((3.14159 - 0.1) + (1e10 + 0.0))) + (0.0 + float32(float64((0.1 / float32((((float64(1e10) + float64((3.14159 - 3.14159))) * 0.0) + ((0.1 + 1.5) * (1.5 * ((float64(0.0) - float64(1e10)) - 3.14159))))))))))
        if true {
            if (1 >= 1) {
                v5 = 1.5
                var v6 uint32 = 65535
                switch "foo bar" {
                case "\n\t":
                    switch fi4 {
                    case 127:
                        p0 = (p0 % (((p0 % ((p0) | 1))) | 1))
                    default:
                        fmt.Println(p1)
                    }
                }
                _ = v6
            } else {
                p0 = (^(p0))
            }
        }
        v7 := GenS1{F0: uint64(p1), F1: (fi4 ^ fi4), F2: 9223372036854775807}
        _ = fi4
        _ = v5
        _ = v7
    }
    p0 = 42
    for fi8 := 0; fi8 < 20; fi8++ {
        if false {
            var v9 int32 = -1
            v9 -= (-(v9))
            _ = v9
        } else {
            v10 := GenS2{F0: 2, F1: -128}
            _ = v10
        }
        for fi11 := 0; fi11 < 5; fi11++ {
            if false {
                var v12 uint16 = p1
                _ = v12
            }
            p1 += (^(p1))
            if true {
                p0 = (p0 << ((482) & 63))
                p1 += 32767
            } else {
                var v13 int = 42
                if true {
                    fmt.Println(v13, fi8)
                    fmt.Println(v13, p1)
                    fmt.Println(v13, fi11)
                }
                var v14 int32 = -1
                fmt.Println(v14)
                _ = v13
                _ = v14
            }
            _ = fi11
        }
        _ = fi8
    }
    p1 = uint16(p0)
    if false {
        if false {
            v15 := []int{-1, 9223372036854775807, 32767}
            _ = v15
        }
        var v16 int64 = -1
        _ = v16
    }
    _ = p0
    _ = p1
    return -1
}

func main() {
    var v1 string = "a"
    var v2 int32 = -86
    v3 := make(map[int]int)
    v2 = (v2 >> ((18446744073709551615) & 63))
    var v17 int64 = -992
    v17 -= v17
    fmt.Println(v1)
    v18 := GenS3{F0: 6, F1: 2147483647, F2: 1}
    switch 0 {
    case -2147483648:
        var v19 int8 = 127
        _ = v19
    case -9223372036854775808:
        if false {
            fmt.Println(v2)
        } else {
            v17 -= -9223372036854775808
            v20 := v3[1]
            _ = v20
        }
    default:
        var v21 float32 = 0.1
        _ = v21
    }
    var v22 uint = 0
    v23 := GenS2{F0: 42, F1: 1}
    v24 := GenS1{F0: uint64(v22), F1: 2147483647, F2: (v17 >> ((531) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
