package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int }
type GenS2 struct { F0 int32; F1 string; F2 uint16 }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 uint16; F1 int64 }
type GenS5 struct { F0 int32 }
func genFn0(p0 int8, p1 uint16, p2 int64) int16 {
    fmt.Println(p1, p2)
    if (((!(992 >= 4294967295)) || false) && ((((float64(1.5) * 3.14159) + 3.14159) == (0.0 + -2.25)) || (((p1 << ((65535) & 63)) != 515) || true))) {
        for fi4 := 0; fi4 < 11; fi4++ {
            p1 += (p1 ^ 32767)
            _ = fi4
        }
        v5 := GenS1{F0: (p1 << ((65535) & 63)), F1: 65535, F2: 42}
        _ = v5
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[int]int)
    v3 := []int{-451}
    _ = len(v2)
    fmt.Println(v1)
    var v6 bool = true
    v1 = v1
    v1 = v1
    var v7 bool = v6
    for fi8 := 0; fi8 < 9; fi8++ {
        if len(v3) > 0 {
            v3[0] = 32767
        }
        for fi9 := 0; fi9 < 15; fi9++ {
            var v10 int32 = -10
            fmt.Println(fi9, fi8)
            if v6 {
                var v11 float32 = float32(3.14159)
                if (true || (false && v7)) {
                    if len(v3) > 0 {
                        v3[0] = fi8
                    }
                    v12 := GenS2{F0: (v10 / (((v10 / (((v10 % (((v10 + (^(v10)))) | 1))) | 1))) | 1)), F1: (v1 + (v1 + "hello")), F2: 42}
                    if false {
                        v11 = (v11 - 1e10)
                        v13 := GenS3{F0: 2, F1: 2147483647}
                        _ = v13
                    }
                    v14 := GenS4{F0: 31, F1: -1}
                    v15 := GenS5{F0: 2147483647}
                    _ = v12
                    _ = v14
                    _ = v15
                } else {
                    v10 = -1
                }
                _ = v11
            } else {
                var v16 int64 = 42
                if v6 {
                    var v17 uint8 = 42
                    _ = v17
                }
                _ = v16
            }
            _ = fi9
            _ = v10
        }
        if (!(42 <= -529)) {
            v18 := GenS4{F0: 42, F1: 9223372036854775807}
            for fi19 := 0; fi19 < 11; fi19++ {
                fmt.Println(v1, fi19)
                v6 = v7
                _ = fi19
            }
            switch "\n\t" {
            case "a":
                v6 = (0.0 <= (((-2.25 + float64((3.14159 - 1.5))) - (3.14159 / 0.1)) * 1e10))
            case "":
                var v20 int64 = 9223372036854775807
                _ = v20
            case "\n\t":
                for fi21 := 0; fi21 < 2; fi21++ {
                    var v22 uint = 255
                    _ = fi21
                    _ = v22
                }
            default:
                var v23 int32 = -2147483648
                _ = v23
            }
            _ = v18
        }
        _ = fi8
    }
    fmt.Println(v1)
    v7 = (v7 || (257 == 1))
    if v6 {
        v3 = append(v3, -552)
        fmt.Println(v6)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
