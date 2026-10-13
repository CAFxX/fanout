package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 int; F2 int8 }
type GenS3 struct { F0 uint; F1 bool }
func genFn0(p0 uint64, p1 int32) int32 {
    p0 = 1
    p1 = p1
    switch 9223372036854775807 {
    case 128:
        p1 += (p1 - (p1 * -1))
    default:
        var v6 string = "hello"
        _ = v6
    }
    _ = p0
    _ = p1
    return p1
}

func main() {
    var v1 bool = true
    v2 := make(map[uint32]string)
    var v3 int64 = 0
    var v4 int64 = v3
    var v5 uint = 2
    v7 := []int{-1, -351, -2147483648}
    fmt.Println(v3, v1)
    _ = len(v2)
    if (320 <= 2147483648) {
        var v8 string = "hello"
        for fi9 := 0; fi9 < 11; fi9++ {
            v4 += v4
            var v10 int32 = -1
            _ = fi9
            _ = v10
        }
        _ = v8
    } else {
        var v11 string = "a"
        if v1 {
            var v12 uint64 = 0
            _ = v12
        }
        _ = v11
    }
    var v13 int32 = 948
    v7 = append(v7, 4294967295)
    switch -840 {
    case -957:
        v3 -= v4
    }
    fmt.Println(v3, v3)
    _ = len(v2)
    if v1 {
        var v14 string = "\n\t"
        _ = v14
    } else {
        for fi15 := 0; fi15 < 15; fi15++ {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            v13 += v13
            _ = fi15
        }
        fmt.Println(v13, v13)
        v16 := GenS1{F0: v4}
        _ = v16
    }
    v17 := GenS2{F0: (v5 ^ 0), F1: 32767, F2: -50}
    for fi18 := 0; fi18 < 5; fi18++ {
        switch -9223372036854775808 {
        case -2147483648:
            switch (("" + ("foo bar" + "hello")) + "\n\t") {
            case "foo bar":
                v19 := v2[4294967295]
                _ = v19
            case "":
                v13 += 1
            case "hello":
                v20 := GenS3{F0: v5, F1: v1}
                _ = v20
            }
        case 888:
            v4 = (v4 << ((1) & 63))
        case 2147483647:
            for fi21 := 0; fi21 < 13; fi21++ {
                fmt.Println(fi18, fi21, v13)
                var v22 uint16 = 42
                v2[2147483648] = "hello"
                v2[42] = "foo bar"
                v4 = v3
                _ = fi21
                _ = v22
            }
        default:
            for fi23 := 0; fi23 < 2; fi23++ {
                v1 = v1
                _ = fi23
            }
        }
        _ = fi18
    }
    var v24 uint32 = 448
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
