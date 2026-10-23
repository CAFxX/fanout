package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 uint32; F2 int }
func genFn0(p0 string) uint16 {
    fmt.Println(p0)
    var v4 int8 = -48
    if true {
        var v5 bool = (true && (42 < 65535))
        for fi6 := 0; fi6 < 7; fi6++ {
            if v5 {
                p0 = "a"
                v4 += v4
                switch (fi6 - (fi6 % (((fi6 % (((fi6 & 827)) | 1))) | 1))) {
                case 4294967295:
                    v7 := []int{int(v4), (^(fi6))}
                    _ = v7
                case -922:
                    var v8 int8 = (^(v4))
                    _ = v8
                case 127:
                    var v9 uint8 = 127
                    _ = v9
                default:
                    p0 = (p0 + "foo bar")
                }
            } else {
                var v10 int32 = 0
                _ = v10
            }
            v11 := []int{fi6, (fi6 | fi6), (fi6 / (((fi6 % ((fi6) | 1))) | 1))}
            _ = fi6
            _ = v11
        }
        _ = v5
    } else {
        var v12 int64 = -9223372036854775808
        if (((v12 - 42) <= v12) && (!false)) {
            v13 := []int{}
            v12 += (v12 * v12)
            fmt.Println(v4, v4, v12)
            _ = v13
        }
        switch 127 {
        case 127:
            switch p0 {
            case "a":
                var v14 int32 = 489
                _ = v14
            }
        default:
            fmt.Println(v4, p0, v4)
        }
        _ = v12
    }
    switch 397 {
    case 0:
        v15 := GenS1{F0: 2}
        _ = v15
    case 127:
        p0 = "\n\t"
    case 128:
        var v16 uint16 = 2
        _ = v16
    default:
        var v17 int8 = -111
        _ = v17
    }
    _ = p0
    _ = v4
    return 0
}

func main() {
    var v1 bool = true
    var v2 uint = 42
    v3 := make(map[int]bool)
    v18 := GenS2{F0: 926}
    if (-1 >= 0) {
        for fi19 := 0; fi19 < 11; fi19++ {
            for fi20 := 0; fi20 < 14; fi20++ {
                var v21 string = (("a" + "a") + "a")
                _ = fi20
                _ = v21
            }
            if (false && v1) {
                v22 := GenS3{F0: "a", F1: 2147483648, F2: (^(fi19))}
                var v23 string = ""
                v3[(fi19 ^ fi19)] = v1
                _ = v22
                _ = v23
            } else {
                var v24 int8 = 0
                var v25 uint8 = 101
                _ = v24
                _ = v25
            }
            v2 -= 520
            _ = fi19
        }
    } else {
        var v26 int16 = 1
        _ = v26
    }
    var v27 uint = v2
    var v28 uint64 = 2
    v28 -= v28
    v29 := []int{int(v27), -9223372036854775808}
    v28 -= v28
    v27 = (v2 * v2)
    v30 := v3[9223372036854775807]
    v2 -= (^(v27))
    fmt.Println(v27)
    switch "\n\t" {
    case "\n\t":
        v31 := GenS1{F0: 725}
        _ = v31
    }
    var v32 bool = v30
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v18
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v32
}
