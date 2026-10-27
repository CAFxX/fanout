package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 uint64 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 int }
func genFn0(p0 uint, p1 int8, p2 int) int8 {
    fmt.Println(p0, p0, p2)
    v4 := GenS1{F0: -9223372036854775808, F1: 2147483647, F2: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -1
}
func genFn1(p0 bool, p1 int64) uint64 {
    fmt.Println(p0, p1)
    v5 := []int{997, 127, 9223372036854775807}
    for fi6 := 0; fi6 < 16; fi6++ {
        p0 = ((^(fi6)) >= fi6)
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = v5
    return 803
}

func main() {
    v1 := []int{128, 9223372036854775807, 671}
    v2 := []int{-9223372036854775808, 1, -2147483648}
    var v3 float64 = (float64((0.0 * 1.5)) + ((((-2.25 + 0.1) + (1.5 + (0.0 + 3.14159))) / 3.14159) / 0.1))
    var v7 float64 = v3
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v7 = 1e10
    for fi8 := 0; fi8 < 12; fi8++ {
        if false {
            var v9 bool = ("a" > ("foo bar" + ""))
            for fi10 := 0; fi10 < 13; fi10++ {
                var v11 bool = (v9 || (2 >= 42))
                v9 = (993 == 2147483648)
                if (false || (96 <= 65535)) {
                    v12 := GenS1{F0: 9223372036854775807, F1: 0, F2: 18446744073709551615}
                    if (845 <= 1) {
                        v3 += float64((3.14159 * 1.5))
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    } else {
                        v13 := GenS2{F0: 42, F1: 42}
                        var v14 uint64 = 1
                        var v15 float64 = ((v3 * float64(3.14159)) * (0.1 / (float64(0.1) * ((v7 / 1.0) - ((v3 - 3.14159) - 1.5)))))
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                    fmt.Println(v3, v7, fi8)
                    _ = v12
                }
                _ = fi10
                _ = v11
            }
            _ = v9
        }
        var v16 string = "\n\t"
        if len(v1) > 0 {
            v1[0] = fi8
        }
        _ = fi8
        _ = v16
    }
    if false {
        v3 = ((float64(3.14159) - ((v7 + v7) / float64(((-2.25 / (3.14159 + 1e10)) / 1e10)))) + (0.0 - v7))
        v17 := GenS1{F0: -9223372036854775808, F1: 2147483647, F2: 823}
        v18 := GenS1{F0: 303, F1: 1, F2: 227}
        _ = v17
        _ = v18
    }
    var v19 uint64 = 620
    if false {
        fmt.Println(v19)
        v20 := GenS3{F0: -9223372036854775808}
        _ = v20
    }
    if false {
        v3 += (((v7 + v3) + (3.14159 - v3)) * (v7 + (float64((3.14159 - (((1e10 + (1e10 / 1.0)) + float32(1.5)) / 1e10))) + 0.0)))
    }
    var v21 int16 = 37
    v19 = 4294967295
    fmt.Println(v7)
    v22 := GenS2{F0: 1, F1: 18446744073709551615}
    switch 210 {
    case -1:
        var v23 int32 = 42
        _ = v23
    case 9223372036854775807:
        _ = len(v1)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v7
}
