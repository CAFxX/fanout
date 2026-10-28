package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 int64 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 int16; F1 uint8; F2 uint64 }
type GenS4 struct { F0 uint8; F1 uint8 }
func genFn0(p0 uint16, p1 uint8) int {
    p1 -= p1
    p0 += 458
    if (true || (2147483647 == 286)) {
        v5 := GenS1{F0: -293, F1: -813, F2: -9223372036854775808}
        var v6 uint32 = 42
        v7 := []int{128}
        _ = v5
        _ = v6
        _ = v7
    }
    p0 = (^(p0))
    _ = p0
    _ = p1
    return 128
}
func genFn1(p0 uint) int8 {
    if false {
        var v8 uint32 = 4294967295
        v8 -= 334
        _ = v8
    }
    v9 := GenS2{F0: (p0 | 603), F1: 21}
    switch 4294967295 {
    case -108:
        var v10 int8 = -19
        _ = v10
    default:
        if true {
            v11 := []int{}
            _ = v11
        }
    }
    _ = p0
    _ = v9
    return 127
}

func main() {
    v1 := make(map[int]uint32)
    var v2 string = "\n\t"
    var v3 uint8 = 42
    v4 := make(map[int]bool)
    var v12 int64 = 42
    for fi13 := 0; fi13 < 16; fi13++ {
        for fi14 := 0; fi14 < 3; fi14++ {
            v15 := GenS3{F0: 1000, F1: (v3 - (v3 ^ (v3 << ((18446744073709551615) & 63)))), F2: 4294967295}
            fmt.Println(v2, v12, v3)
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    v2 = ""
    v16 := []int{127}
    var v17 bool = (!true)
    if v17 {
        var v18 uint64 = 704
        _ = v18
    }
    if false {
        v19 := GenS3{F0: 0, F1: 0, F2: 2}
        v20 := GenS1{F0: 1, F1: 2147483647, F2: v12}
        _ = v19
        _ = v20
    } else {
        if v17 {
            v2 = ("hello" + v2)
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
            var v21 int32 = -1
            _ = v21
        }
    }
    var v22 uint = 751
    var v23 int64 = (^(v12))
    v22 -= (v22 / ((v22) | 1))
    v23 = (v12 & 0)
    var v24 bool = (649 != v12)
    v24 = (v24 && v17)
    v22 += (v22 % (((v22 % ((1) | 1))) | 1))
    v2 = "a"
    switch 1 {
    case 32768:
        fmt.Println(v22)
    case 32767:
        for fi25 := 0; fi25 < 16; fi25++ {
            v1[fi25] = 65535
            v16 = append(v16, fi25)
            _ = fi25
        }
    case -9223372036854775808:
        var v26 uint16 = 2
        _ = v26
    default:
        if false {
            for fi27 := 0; fi27 < 7; fi27++ {
                v23 += (v23 & (v12 ^ v23))
                _ = fi27
            }
        } else {
            v28 := GenS4{F0: (v3 | 68), F1: 42}
            _ = v28
        }
    }
    for fi29 := 0; fi29 < 20; fi29++ {
        var v30 float32 = 0.0
        _ = fi29
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
