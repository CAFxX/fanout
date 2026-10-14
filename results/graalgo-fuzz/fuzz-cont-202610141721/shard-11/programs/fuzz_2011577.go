package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 string; F2 uint }
type GenS4 struct { F0 int64; F1 int; F2 int8 }
func genFn0(p0 uint32, p1 uint16) uint {
    var v6 uint16 = (^(p1))
    for fi7 := 0; fi7 < 6; fi7++ {
        v8 := GenS1{F0: (p0 / (((^(p0))) | 1)), F1: 42, F2: -1}
        v9 := GenS2{F0: (4294967295 < 255)}
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := []int{144}
    for fi11 := 0; fi11 < 15; fi11++ {
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        v12 := GenS3{F0: 42, F1: "foo bar", F2: 0}
        _ = fi11
        _ = v12
    }
    switch ("foo bar" + ("hello" + "hello")) {
    case "":
        v13 := GenS1{F0: (p0 ^ (p0 >> ((42) & 63))), F1: 260, F2: 1}
        _ = v13
    case "foo bar":
        if len(v10) > 0 {
            v10[0] = 9223372036854775807
        }
    case "a":
        v6 = (p1 << ((18446744073709551615) & 63))
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v6
    return 808
}

func main() {
    var v1 int8 = 127
    var v2 int8 = v1
    var v3 uint16 = 66
    v4 := []uint32{}
    v5 := []uint32{2147483647, 151, 2147483647}
    v1 = v2
    _ = len(v4)
    fmt.Println(v2, v2)
    var v14 int = 4294967295
    v15 := GenS4{F0: -1, F1: (-(v14)), F2: (v1 >> ((921) & 63))}
    fmt.Println(v14, v14)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v16 float64 = 1.5
    _ = len(v5)
    v16 = v16
    v4 = append(v4, 42)
    switch "" {
    case "foo bar":
        v1 += (v2 / (((v1 ^ v1)) | 1))
    }
    var v17 string = ("foo bar" + ("foo bar" + "\n\t"))
    v14 += (v14 >> ((2) & 63))
    fmt.Println(v14, v2)
    fmt.Println(v1, v2)
    v18 := GenS4{F0: 837, F1: v14, F2: (v2 / ((v2) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
