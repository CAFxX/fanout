package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8 }
func genFn0(p0 int16, p1 int8, p2 string) int {
    if (!(("a" + p2) != ("a" + "a"))) {
        switch "" {
        case "":
            p1 = (p1 | (p1 >> ((42) & 63)))
        case "hello":
            v6 := []int{0, -852, -1}
            _ = v6
        }
        p0 = int16(p1)
        v7 := []int{}
        _ = v7
    }
    p0 = (p0 << ((467) & 63))
    if (-999 <= -1) {
        fmt.Println(p0)
    }
    if ((255 != 738) && (!false)) {
        p2 = p2
    } else {
        var v8 bool = true
        switch 4294967295 {
        case 128:
            v9 := GenS1{F0: -850, F1: 91}
            _ = v9
        }
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    return 128
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := []string{("foo bar" + "hello"), ""}
    v3 := []string{}
    v4 := []int{549, int(v1), 128}
    v5 := []int{1, 128}
    var v10 uint = (v1 * (v1 >> ((v1) & 63)))
    v1 -= v1
    if true {
        if (383 == (^(v10))) {
            v2 = append(v2, "a")
        }
        switch 2147483647 {
        case -9223372036854775808:
            var v11 uint32 = 65535
            _ = v11
        case 423:
            _ = len(v3)
        case 0:
            v12 := GenS2{F0: (v10 - v1)}
            _ = v12
        default:
            v13 := GenS2{F0: 18446744073709551615}
            _ = v13
        }
        fmt.Println(v1)
    } else {
        fmt.Println(v1, v10)
        var v14 int64 = 42
        for fi15 := 0; fi15 < 7; fi15++ {
            v14 = 1
            _ = fi15
        }
        _ = v14
    }
    if true {
        if ((true || (32767 < 625)) || (2 >= 4294967295)) {
            v16 := GenS2{F0: v1}
            _ = v16
        } else {
            v17 := GenS1{F0: 0, F1: 58}
            fmt.Println(v1)
            _ = v17
        }
        var v18 uint = (v10 - (^(v1)))
        v19 := GenS2{F0: (v10 >> ((255) & 63))}
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 15; fi20++ {
        fmt.Println(v1)
        v10 += (v1 - v1)
        _ = fi20
    }
    v10 -= (^(v1))
    _ = len(v2)
    v21 := GenS3{F0: -29}
    v1 = v10
    v1 -= 65535
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
