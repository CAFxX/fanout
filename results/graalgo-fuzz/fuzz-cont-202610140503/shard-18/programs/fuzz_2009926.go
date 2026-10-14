package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int }

func main() {
    v1 := []int{1, 2147483647, 1, 4294967295}
    v2 := []uint32{2, 763}
    v3 := make(map[string]int)
    v4 := make(map[uint32]uint32)
    v5 := make(map[string]bool)
    var v6 uint8 = 178
    switch 32768 {
    case -129:
        for fi7 := 0; fi7 < 16; fi7++ {
            var v8 uint64 = 0
            for fi9 := 0; fi9 < 8; fi9++ {
                var v10 string = ("hello" + "")
                _ = fi9
                _ = v10
            }
            var v11 uint = 536
            _ = fi7
            _ = v11
            _ = v8
        }
    }
    var v12 uint32 = 42
    var v13 int = -636
    var v14 uint16 = 42
    for fi15 := 0; fi15 < 3; fi15++ {
        v14 = v14
        var v16 int8 = 0
        fmt.Println(v12)
        _ = fi15
        _ = v16
    }
    v12 = uint32(v6)
    switch (v13 >> ((255) & 63)) {
    case 4294967295:
        v17 := GenS1{F0: -9223372036854775808, F1: 1, F2: false}
        _ = v17
    case 725:
        v12 = v12
    case 2147483647:
        if len(v1) > 0 {
            v1[0] = 4294967295
        }
    default:
        var v18 int = (v13 % ((v13) | 1))
        _ = v18
    }
    var v19 uint16 = 42
    _ = len(v4)
    var v20 bool = (!(!false))
    var v21 int8 = 84
    switch "\n\t" {
    case "a":
        switch 32768 {
        case 889:
            v22 := GenS2{F0: "hello"}
            _ = v22
        case 9223372036854775807:
            fmt.Println(v21, v6)
        case 4294967295:
            v6 -= (v6 & v6)
        }
    case "foo bar":
        _ = len(v5)
    default:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    switch v13 {
    case 2147483647:
        v14 = v19
    case 42:
        var v23 int8 = v21
        _ = v23
    case -9223372036854775808:
        v20 = false
    }
    if (v19 >= (^(v14))) {
        var v24 uint32 = v12
        var v25 uint32 = uint32(v6)
        var v26 uint = 4294967295
        _ = v24
        _ = v25
        _ = v26
    } else {
        if true {
            var v27 int = -806
            v28 := v4[(v12 + v12)]
            _ = v27
            _ = v28
        } else {
            var v29 uint64 = 877
            fmt.Println(v6, v6)
            for fi30 := 0; fi30 < 19; fi30++ {
                v31 := GenS2{F0: ""}
                v32 := GenS3{F0: (^(v13))}
                v19 += (v14 / ((v14) | 1))
                _ = fi30
                _ = v31
                _ = v32
            }
            _ = v29
        }
        for fi33 := 0; fi33 < 13; fi33++ {
            fmt.Println(v12, v20, v20)
            _ = fi33
        }
        v6 = 1
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
