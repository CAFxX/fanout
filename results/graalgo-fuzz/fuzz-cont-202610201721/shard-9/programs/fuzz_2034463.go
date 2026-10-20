package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string }
func genFn0(p0 int, p1 uint64, p2 uint16) int8 {
    p1 -= (p1 >> ((4294967295) & 63))
    var v7 int32 = -883
    fmt.Println(p0, p0)
    var v8 int32 = 42
    switch "hello" {
    case "a":
        var v9 uint8 = 128
        _ = v9
    default:
        fmt.Println(v8)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := []uint32{}
    var v2 uint64 = 1
    v3 := make(map[int]uint32)
    var v4 uint = 1
    var v5 int = 127
    v6 := []string{"a", "hello", "foo bar", "hello"}
    delete(v3, 32767)
    v4 -= (^(v4))
    v3[-92] = 2147483648
    if (1e10 < 0.1) {
        v10 := GenS1{F0: v4}
        v4 += v4
        var v11 int8 = 54
        _ = v10
        _ = v11
    } else {
        switch "a" {
        case "\n\t":
            v12 := GenS1{F0: v4}
            _ = v12
        case "foo bar":
            v5 = (v5 << (((v4 * (v4 - (v4 / (((v4 + v4)) | 1))))) & 63))
        case "hello":
            switch v5 {
            case 0:
                var v13 uint = (v4 * (v4 | (v4 << (((v4 & v4)) & 63))))
                _ = v13
            case -9223372036854775808:
                fmt.Println(v5, v2)
            }
        default:
            if (42 < 2147483647) {
                v14 := GenS1{F0: v4}
                _ = v14
            }
        }
    }
    v15 := GenS1{F0: v4}
    var v16 string = "hello"
    var v17 uint16 = 32767
    v18 := GenS2{F0: v2}
    v19 := GenS2{F0: v2}
    fmt.Println(v17, v16)
    var v20 int64 = -1
    var v21 string = "hello"
    for fi22 := 0; fi22 < 9; fi22++ {
        if false {
            v23 := GenS3{F0: (v16 + ("\n\t" + v16))}
            if (false && (!false)) {
                v16 = ((v16 + v21) + v21)
            }
            for fi24 := 0; fi24 < 18; fi24++ {
                _ = len(v3)
                v21 = "foo bar"
                _ = fi24
            }
            _ = v23
        } else {
            fmt.Println(fi22, v2, v16)
        }
        _ = fi22
    }
    v20 -= 622
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
