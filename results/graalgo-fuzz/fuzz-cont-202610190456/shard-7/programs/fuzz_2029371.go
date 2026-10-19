package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int8 }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 int, p1 string) string {
    switch p0 {
    case 2147483647:
        fmt.Println(p1, p1)
    case 32767:
        if false {
            v7 := GenS1{F0: 128, F1: (p0 ^ 127), F2: -1}
            v8 := []int{}
            p0 = p0
            _ = v7
            _ = v8
        }
    default:
        if (!(124 > 42)) {
            if (-363 >= (p0 + -9223372036854775808)) {
                fmt.Println(p1, p1)
            }
        } else {
            var v9 float64 = ((0.1 + float64(1.5)) + 0.1)
            var v10 bool = false
            _ = v10
            _ = v9
        }
    }
    fmt.Println(p1, p1)
    v11 := []int{128, (p0 ^ p0), (p0 / (((p0 | (p0 | p0))) | 1))}
    _ = p0
    _ = p1
    _ = v11
    return p1
}

func main() {
    var v1 uint = 2
    v2 := make(map[uint32]uint32)
    var v3 int = 1
    var v4 int16 = 32767
    var v5 uint64 = 578
    v6 := []int{-834, (v3 ^ v3), (v3 & v3)}
    switch v3 {
    case -1:
        v12 := GenS1{F0: 128, F1: (v3 | (v3 + v3)), F2: -128}
        _ = v12
    case 32768:
        var v13 uint32 = 234
        _ = v13
    case 127:
        var v14 uint32 = 42
        _ = v14
    default:
        var v15 uint64 = (^(v5))
        _ = v15
    }
    v3 += (v3 | (-(v3)))
    v16 := GenS1{F0: 42, F1: v3, F2: 1}
    v17 := GenS2{F0: -128, F1: -319}
    for fi18 := 0; fi18 < 10; fi18++ {
        v5 = (^(v5))
        switch -701 {
        case 531:
            var v19 float64 = 1.5
            _ = v19
        case -1:
            var v20 int8 = -128
            _ = v20
        }
        v21 := GenS3{F0: 18446744073709551615}
        _ = fi18
        _ = v21
    }
    var v22 int16 = -1
    var v23 int = v3
    v24 := GenS3{F0: v5}
    fmt.Println(v5)
    if false {
        v25 := GenS3{F0: (v5 >> ((18446744073709551615) & 63))}
        _ = v25
    }
    v5 -= (v5 >> (((v1 % ((4294967295) | 1))) & 63))
    var v26 int = v3
    switch (v3 & 103) {
    case -1:
        if false {
            switch v26 {
            case 364:
                for fi27 := 0; fi27 < 14; fi27++ {
                    var v28 float64 = 3.14159
                    var v29 uint64 = 633
                    _ = fi27
                    _ = v28
                    _ = v29
                }
            }
        }
    case -9223372036854775808:
        _ = len(v6)
    }
    var v30 string = ("\n\t" + "foo bar")
    v5 += 0
    var v31 int32 = -1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
