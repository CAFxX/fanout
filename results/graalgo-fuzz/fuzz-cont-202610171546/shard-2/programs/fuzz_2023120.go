package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 uint64 }
type GenS4 struct { F0 bool; F1 uint }

func main() {
    var v1 uint16 = 0
    v2 := make(map[string]uint32)
    var v3 int32 = 0
    for fi4 := 0; fi4 < 8; fi4++ {
        v5 := []int{1, fi4, (fi4 >> ((744) & 63)), 0}
        for fi6 := 0; fi6 < 14; fi6++ {
            var v7 int = (fi6 >> ((725) & 63))
            _ = fi6
            _ = v7
        }
        v3 -= v3
        _ = fi4
        _ = v5
    }
    var v8 uint8 = 1
    for fi9 := 0; fi9 < 11; fi9++ {
        v10 := GenS1{F0: false, F1: 1}
        _ = fi9
        _ = v10
    }
    v11 := v2["\n\t"]
    v2["\n\t"] = (v11 % (((v11 | (v11 << ((1) & 63)))) | 1))
    if ((-705 >= 1) && true) {
        var v12 uint16 = (v1 | (v1 + (^(v1))))
        for fi13 := 0; fi13 < 7; fi13++ {
            if true {
                var v14 uint32 = v11
                _ = v14
            } else {
                var v15 int64 = 9223372036854775807
                _ = v15
            }
            _ = fi13
        }
        v16 := GenS2{F0: 42}
        _ = v12
        _ = v16
    } else {
        v17 := GenS3{F0: (v3 + v3), F1: 934}
        v1 += v1
        _ = v17
    }
    v18 := GenS2{F0: (v3 * 42)}
    var v19 uint = 65535
    if true {
        v20 := GenS3{F0: -1, F1: 640}
        _ = v20
    }
    for fi21 := 0; fi21 < 18; fi21++ {
        fmt.Println(v11, v11)
        _ = fi21
    }
    switch 298 {
    case 9223372036854775807:
        var v22 int16 = 1000
        _ = v22
    case 32767:
        var v23 int = 32768
        _ = v23
    case 738:
        for fi24 := 0; fi24 < 14; fi24++ {
            v25 := GenS1{F0: (451 > v11), F1: -78}
            _ = fi24
            _ = v25
        }
    default:
        for fi26 := 0; fi26 < 11; fi26++ {
            fmt.Println(fi26, v8, fi26)
            if (1.5 == (1.5 / 1e10)) {
                fmt.Println(v19, v1, v19)
            } else {
                v27 := GenS1{F0: (!false), F1: -1}
                v28 := []int{524, fi26}
                var v29 uint8 = v8
                _ = v27
                _ = v28
                _ = v29
            }
            _ = fi26
        }
    }
    v30 := []int{128}
    for fi31 := 0; fi31 < 2; fi31++ {
        var v32 int64 = 1
        var v33 int16 = 42
        v33 += (-(v33))
        _ = fi31
        _ = v32
        _ = v33
    }
    v34 := GenS4{F0: false, F1: (v19 ^ 65535)}
    var v35 uint64 = uint64(v8)
    if false {
        v35 -= (^(v35))
        v1 = (v1 & (v1 << ((18446744073709551615) & 63)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v35:", v35)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v8
}
