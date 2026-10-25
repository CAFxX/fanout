package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 int64, p1 uint64, p2 uint8) uint {
    for fi5 := 0; fi5 < 1; fi5++ {
        if (fi5 > fi5) {
            for fi6 := 0; fi6 < 3; fi6++ {
                p2 = 104
                _ = fi6
            }
        }
        switch ("hello" + "hello") {
        case "foo bar":
            var v7 int64 = (p0 * -1)
            _ = v7
        case "hello":
            p1 += 4294967295
        case "\n\t":
            p0 -= (p0 | (p0 / ((p0) | 1)))
        default:
            p2 = (^(p2))
        }
        var v8 uint64 = (p1 / ((p1) | 1))
        _ = fi5
        _ = v8
    }
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    return 6
}

func main() {
    var v1 int64 = -181
    v2 := make(map[int]uint32)
    v3 := []int{9223372036854775807, 127, 2147483647}
    var v4 float32 = 1.5
    var v9 uint = 210
    v10 := GenS1{F0: 950}
    switch "\n\t" {
    case "a":
        v3 = append(v3, int(v1))
    case "":
        if (255 < 32767) {
            v9 -= (v9 % ((v9) | 1))
        }
    case "\n\t":
        var v11 float32 = v4
        _ = v11
    default:
        for fi12 := 0; fi12 < 8; fi12++ {
            fmt.Println(v9)
            _ = fi12
        }
    }
    var v13 bool = (((("foo bar" + ((("a" + "") + "foo bar") + "hello")) >= "foo bar") || true) && (false && false))
    var v14 int32 = -760
    var v15 uint = (^(v9))
    fmt.Println(v4)
    for fi16 := 0; fi16 < 2; fi16++ {
        var v17 int32 = v14
        v3 = append(v3, fi16)
        var v18 int64 = (v1 % (((v1 + 0)) | 1))
        _ = fi16
        _ = v17
        _ = v18
    }
    v15 = (v9 >> ((v9) & 63))
    for fi19 := 0; fi19 < 5; fi19++ {
        var v20 uint16 = 32767
        fmt.Println(v20, v15, v14)
        _ = fi19
        _ = v20
    }
    v15 = genFn0((v1 - (v1 + v1)), 4294967295, 0)
    v4 += 0.0
    v4 = 1.5
    fmt.Println(v15)
    v21 := GenS1{F0: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v9
}
