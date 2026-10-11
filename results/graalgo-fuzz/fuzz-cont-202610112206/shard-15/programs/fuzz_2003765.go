package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string }
func genFn0(p0 uint8, p1 uint) uint64 {
    if false {
        switch "foo bar" {
        case "\n\t":
            var v6 uint32 = 4294967295
            _ = v6
        default:
            p1 = p1
        }
        p1 -= (p1 >> ((73) & 63))
    }
    switch 18 {
    case 128:
        v7 := GenS1{F0: 9223372036854775807, F1: (200 <= 42)}
        _ = v7
    case 240:
        if ((1 <= 1) || (2 > 65535)) {
            switch ((("" + "hello") + "a") + "a") {
            case "a":
                for fi8 := 0; fi8 < 20; fi8++ {
                    p0 -= p0
                    _ = fi8
                }
            case "":
                var v9 uint64 = 18446744073709551615
                _ = v9
            default:
                v10 := GenS2{F0: -9223372036854775808}
                _ = v10
            }
            fmt.Println(p0, p0)
            switch 2147483647 {
            case -864:
                fmt.Println(p1, p1)
            case 1:
                p1 -= p1
            case 2147483647:
                var v11 uint16 = 2
                _ = v11
            }
        } else {
            fmt.Println(p0)
            p0 -= p0
        }
    }
    _ = p0
    _ = p1
    return 412
}

func main() {
    var v1 float32 = 1e10
    var v2 uint64 = 284
    var v3 float32 = 1e10
    var v4 int64 = 0
    v5 := make(map[int]string)
    v2 = genFn0(255, 376)
    var v12 int = 190
    var v13 float32 = ((v3 / 1.5) * (0.0 - 1e10))
    var v14 uint16 = 333
    _ = len(v5)
    fmt.Println(v14, v13, v13)
    v15 := []int{829, (v12 % (((v12 * 127)) | 1)), 4294967295}
    switch (-(v12)) {
    case 128:
        fmt.Println(v14, v12)
    default:
        _ = len(v5)
    }
    v12 += v12
    var v16 uint = 65535
    for fi17 := 0; fi17 < 20; fi17++ {
        var v18 int64 = v4
        _ = fi17
        _ = v18
    }
    fmt.Println(v12, v16, v2)
    v2 = (v2 >> ((v16) & 63))
    var v19 uint8 = 127
    switch 127 {
    case -9223372036854775808:
        v4 = v4
    }
    v20 := GenS3{F0: "\n\t"}
    fmt.Println(v3, v16)
    var v21 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
