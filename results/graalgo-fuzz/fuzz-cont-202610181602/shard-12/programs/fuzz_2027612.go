package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int; F1 int32; F2 uint16 }
func genFn0(p0 string, p1 uint8, p2 int16) uint16 {
    p2 = -588
    for fi7 := 0; fi7 < 5; fi7++ {
        fmt.Println(p1, p1)
        fmt.Println(p1, p0, p1)
        _ = fi7
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := []int{fi8}
        _ = fi8
        _ = v9
    }
    switch p0 {
    case "\n\t":
        for fi10 := 0; fi10 < 6; fi10++ {
            if (!((-1 <= 108) || (!false))) {
                v11 := GenS1{F0: 42, F1: 255, F2: 182}
                fmt.Println(p0, p2)
                v12 := GenS1{F0: -497, F1: 907, F2: 537}
                _ = v11
                _ = v12
            }
            fmt.Println(fi10)
            v13 := []int{}
            _ = fi10
            _ = v13
        }
    case "a":
        var v14 bool = true
        _ = v14
    case "":
        var v15 uint32 = 339
        _ = v15
    default:
        v16 := GenS2{F0: 526}
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    return 2
}

func main() {
    var v1 int32 = 382
    var v2 uint8 = 128
    var v3 int16 = 32767
    var v4 uint16 = uint16(v2)
    var v5 float64 = 0.1
    v6 := make(map[uint32]int)
    if (false || false) {
        for fi17 := 0; fi17 < 11; fi17++ {
            var v18 int8 = 100
            var v19 uint32 = uint32(v2)
            _ = fi17
            _ = v18
            _ = v19
        }
        v2 = (v2 >> ((42) & 63))
    } else {
        v20 := GenS3{F0: (^(v3))}
        _ = v20
    }
    var v21 int16 = (v3 - 1000)
    var v22 string = ("foo bar" + "a")
    v23 := GenS4{F0: 32767, F1: (v1 / (((v1 ^ (v1 / ((v1) | 1)))) | 1)), F2: 65535}
    if true {
        v3 = v21
    }
    v21 += v21
    if (1 <= (^(v1))) {
        v24 := v6[2147483648]
        var v25 int8 = 42
        if false {
            if false {
                v26 := GenS3{F0: (v21 - 983)}
                fmt.Println(v22)
                _ = v26
            } else {
                var v27 int16 = (v3 / ((32767) | 1))
                _ = v27
            }
        }
        _ = v24
        _ = v25
    }
    var v28 string = v22
    v29 := []int{4294967295}
    fmt.Println(v4, v28)
    v1 = (v1 & 738)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
