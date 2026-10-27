package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 string; F2 uint64 }
func genFn0(p0 int8) int64 {
    v6 := GenS1{F0: true, F1: -88}
    if true {
        v7 := GenS1{F0: true, F1: -293}
        for fi8 := 0; fi8 < 7; fi8++ {
            for fi9 := 0; fi9 < 6; fi9++ {
                p0 = p0
                p0 = -1
                v10 := GenS1{F0: (0 <= 42), F1: 507}
                _ = fi9
                _ = v10
            }
            if false {
                var v11 uint = 4294967295
                fmt.Println(fi8)
                _ = v11
            } else {
                p0 += (p0 ^ (p0 & p0))
            }
            _ = fi8
        }
        _ = v7
    }
    _ = p0
    _ = v6
    return 1
}
func genFn1() uint32 {
    v12 := GenS2{F0: true}
    switch -1 {
    case 9223372036854775807:
        var v13 int64 = -1
        _ = v13
    case 32767:
        switch "" {
        case "a":
            var v14 int32 = 489
            _ = v14
        case "hello":
            for fi15 := 0; fi15 < 9; fi15++ {
                for fi16 := 0; fi16 < 5; fi16++ {
                    var v17 int8 = -1
                    var v18 uint16 = 0
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                _ = fi15
            }
        }
    case 4294967295:
        var v19 int32 = -2147483648
        _ = v19
    }
    v20 := GenS3{F0: 2147483648, F1: (("a" + "\n\t") + "\n\t"), F2: 688}
    _ = v12
    _ = v20
    return 14
}

func main() {
    var v1 uint16 = 65535
    var v2 int = 127
    v3 := make(map[uint32]uint32)
    var v4 uint16 = 386
    v5 := make(map[uint32]uint32)
    fmt.Println(v1)
    v21 := []int{v2, (v2 << ((2) & 63)), (v2 / (((^(v2))) | 1)), 490}
    v22 := GenS3{F0: 0, F1: "a", F2: 18446744073709551615}
    switch (v2 << ((65535) & 63)) {
    case 9223372036854775807:
        v3[0] = genFn1()
    case -184:
        var v23 uint8 = 127
        _ = v23
    default:
        v4 = (v1 ^ (v1 + 1))
    }
    for fi24 := 0; fi24 < 20; fi24++ {
        v21 = append(v21, 0)
        _ = fi24
    }
    v25 := GenS3{F0: 976, F1: "\n\t", F2: 859}
    var v26 int = v2
    var v27 bool = false
    for fi28 := 0; fi28 < 20; fi28++ {
        if (0.1 >= float64(3.14159)) {
            fmt.Println(fi28, v27, v26)
            v27 = false
        } else {
            var v29 int = (v2 % (((v2 % (((fi28 * v2)) | 1))) | 1))
            _ = v29
        }
        _ = fi28
    }
    var v30 bool = v27
    v26 -= (^(v2))
    var v31 int = (v26 << ((65535) & 63))
    _ = len(v3)
    for fi32 := 0; fi32 < 15; fi32++ {
        v2 += v26
        var v33 float32 = (0.0 / 1.0)
        _ = fi32
        _ = v33
    }
    v4 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
