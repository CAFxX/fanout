package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 int32, p1 uint32, p2 uint8) uint8 {
    v5 := GenS1{F0: p0}
    p2 -= p2
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p2
}
func genFn1(p0 int) bool {
    if ((true || true) && false) {
        p0 = -761
        for fi6 := 0; fi6 < 1; fi6++ {
            fmt.Println(fi6, p0)
            var v7 uint64 = 543
            _ = fi6
            _ = v7
        }
    } else {
        fmt.Println(p0)
    }
    v8 := GenS1{F0: -2147483648}
    v9 := []int{(p0 ^ -22), p0, (p0 - (p0 - p0)), (p0 & (p0 / ((p0) | 1)))}
    _ = p0
    _ = v8
    _ = v9
    return true
}

func main() {
    v1 := make(map[int]int)
    v2 := []string{}
    var v3 uint32 = 65535
    v4 := make(map[uint32]string)
    if true {
        var v10 uint16 = 854
        var v11 uint = 18446744073709551615
        _ = v10
        _ = v11
    }
    v3 = (^(v3))
    var v12 uint8 = 2
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if true {
        v12 -= v12
    }
    var v13 uint64 = uint64(v12)
    switch "a" {
    case "a":
        v3 -= (v3 - uint32(v13))
    case "":
        for fi14 := 0; fi14 < 18; fi14++ {
            if len(v2) > 0 {
                v2[0] = ("hello" + "a")
            }
            _ = len(v1)
            for fi15 := 0; fi15 < 14; fi15++ {
                v12 = v12
                var v16 uint = 18446744073709551615
                var v17 uint64 = (v13 * v13)
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = fi14
        }
    case "\n\t":
        fmt.Println(v13, v3)
    default:
        v2 = append(v2, ((("foo bar" + "foo bar") + ((("hello" + "a") + "hello") + "")) + "foo bar"))
    }
    v18 := GenS1{F0: 207}
    for fi19 := 0; fi19 < 1; fi19++ {
        v3 = (v3 | (^(v3)))
        _ = fi19
    }
    v12 += 0
    fmt.Println(v13)
    switch 664 {
    case -9223372036854775808:
        v20 := v1[42]
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
