package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 int64; F2 uint8 }
type GenS3 struct { F0 uint32; F1 uint32; F2 uint32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint }
func genFn0(p0 uint) int32 {
    p0 += p0
    v6 := GenS1{F0: -96}
    switch "" {
    case "":
        v7 := GenS2{F0: "", F1: 9223372036854775807, F2: 127}
        _ = v7
    case "\n\t":
        if (!false) {
            var v8 int16 = -757
            if true {
                v9 := GenS3{F0: 42, F1: 0, F2: 4294967295}
                v10 := GenS1{F0: 829}
                if false {
                    var v11 uint32 = 65535
                    p0 -= uint(v11)
                    v8 += -682
                    _ = v11
                }
                _ = v10
                _ = v9
            }
            _ = v8
        }
    }
    p0 -= 183
    _ = p0
    _ = v6
    return 2147483647
}
func genFn1(p0 uint16, p1 uint8) int64 {
    if (419 != 902) {
        v12 := []int{}
        fmt.Println(p0)
        var v13 uint16 = uint16(p1)
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: 0}
    var v15 float64 = (-2.25 / -2.25)
    for fi16 := 0; fi16 < 17; fi16++ {
        var v17 uint8 = 255
        var v18 float32 = 0.0
        p1 += (^(v17))
        _ = fi16
        _ = v17
        _ = v18
    }
    p1 += (p1 >> ((255) & 63))
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    return 346
}

func main() {
    var v1 int16 = -32768
    var v2 int32 = 2147483647
    var v3 int64 = 0
    v4 := make(map[int]int)
    v5 := []int{42, -9223372036854775808, 1, -2147483648}
    fmt.Println(v3, v2)
    switch "" {
    case "hello":
        switch 32768 {
        case 1:
            v19 := GenS4{F0: v3}
            _ = v19
        case 42:
            v1 = v1
        case 9223372036854775807:
            v2 = 0
        }
    }
    v3 = (v3 + (v3 << ((42) & 63)))
    var v20 int16 = int16(v3)
    var v21 uint = 246
    var v22 int32 = 42
    for fi23 := 0; fi23 < 11; fi23++ {
        switch fi23 {
        case -714:
            v24 := GenS5{F0: 65535}
            _ = v24
        }
        v22 = genFn0((v21 >> ((841) & 63)))
        _ = fi23
    }
    var v25 int = 262
    switch (^(v25)) {
    case -9223372036854775808:
        if true {
            v26 := GenS6{F0: (^(v21))}
            v20 -= v20
            if false {
                fmt.Println(v2, v1)
                var v27 int8 = 100
                _ = v27
            }
            _ = v26
        }
    case 127:
        v21 -= 18446744073709551615
    default:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
