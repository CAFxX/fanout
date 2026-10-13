package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 int32 }
func genFn0(p0 string) uint8 {
    if ((!(772 <= 1)) && true) {
        fmt.Println(p0)
        for fi5 := 0; fi5 < 9; fi5++ {
            for fi6 := 0; fi6 < 6; fi6++ {
                if true {
                    fmt.Println(p0)
                } else {
                    v7 := []int{-2147483648, fi5}
                    _ = v7
                }
                var v8 bool = true
                v9 := []int{}
                _ = fi6
                _ = v8
                _ = v9
            }
            var v10 int32 = -1
            if (true && (0 >= 319)) {
                fmt.Println(p0, fi5, p0)
                var v11 uint = 726
                var v12 uint32 = 704
                _ = v11
                _ = v12
            }
            _ = fi5
            _ = v10
        }
    } else {
        var v13 string = (p0 + p0)
        _ = v13
    }
    var v14 float32 = (1.5 + (0.1 / 1e10))
    _ = p0
    _ = v14
    return 16
}

func main() {
    var v1 int8 = -1
    var v2 float32 = ((0.1 / 0.1) / 1.0)
    var v3 uint32 = 2147483647
    var v4 int16 = -1
    v15 := []int{-9223372036854775808}
    fmt.Println(v2)
    v2 = v2
    for fi16 := 0; fi16 < 16; fi16++ {
        var v17 string = "\n\t"
        _ = fi16
        _ = v17
    }
    var v18 uint = 18446744073709551615
    v15 = append(v15, 0)
    fmt.Println(v2, v1, v18)
    v4 = (v4 & (v4 + v4))
    switch "foo bar" {
    case "a":
        if len(v15) > 0 {
            v15[0] = 127
        }
    case "\n\t":
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
    case "":
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
    default:
        v3 -= v3
    }
    v2 -= (v2 - 0.0)
    if (((-9223372036854775808 > 955) && true) || false) {
        v18 += 0
        var v19 uint = 4294967295
        _ = v19
    }
    for fi20 := 0; fi20 < 2; fi20++ {
        switch "hello" {
        case "":
            fmt.Println(v18)
        default:
            var v21 int16 = 42
            _ = v21
        }
        for fi22 := 0; fi22 < 20; fi22++ {
            var v23 int64 = 9223372036854775807
            fmt.Println(v23, v4)
            _ = fi22
            _ = v23
        }
        _ = fi20
    }
    for fi24 := 0; fi24 < 17; fi24++ {
        fmt.Println(v3, fi24)
        _ = fi24
    }
    var v25 uint64 = 0
    if (713 == 32767) {
        var v26 uint8 = 1
        _ = v26
    } else {
        v27 := GenS1{F0: (v4 ^ (-(v4))), F1: -926, F2: (v25 * (v25 & (v25 & v25)))}
        var v28 float32 = 3.14159
        switch 2147483647 {
        case -1:
            if len(v15) > 0 {
                v15[0] = 4294967295
            }
        case 127:
            if ("" == "") {
                for fi29 := 0; fi29 < 14; fi29++ {
                    v30 := GenS2{F0: (fi29 - (fi29 >> ((255) & 63)))}
                    _ = fi29
                    _ = v30
                }
                v31 := GenS3{F0: 42, F1: 1}
                fmt.Println(v25, v2, v28)
                _ = v31
            }
        case 32768:
            if len(v15) > 0 {
                v15[0] = 385
            }
        default:
            v4 -= (v4 % ((v4) | 1))
        }
        _ = v27
        _ = v28
    }
    if true {
        fmt.Println(v18, v18, v25)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}
