package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 int16; F2 int16 }
func genFn0(p0 int32, p1 uint8, p2 uint) uint32 {
    fmt.Println(p0, p0)
    var v4 uint32 = 623
    var v5 uint8 = (p1 >> ((2) & 63))
    v5 -= (p1 & 248)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 2147483647
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 42
    var v3 uint = 18446744073709551615
    switch "foo bar" {
    case "foo bar":
        var v6 int64 = 14
        _ = v6
    case "a":
        _ = len(v1)
    case "hello":
        v7 := []int{}
        _ = v7
    }
    v8 := GenS1{F0: "a", F1: 255, F2: v3}
    var v9 uint16 = (v2 >> ((v3) & 63))
    v3 = (v3 & (v3 >> ((0) & 63)))
    v10 := []int{729}
    var v11 uint16 = (^(v2))
    switch "a" {
    case "a":
        v12 := GenS1{F0: "", F1: 255, F2: (v3 << ((v3) & 63))}
        _ = v12
    }
    v9 += v11
    for fi13 := 0; fi13 < 7; fi13++ {
        switch ((("a" + "a") + (("" + "a") + "")) + ("hello" + "")) {
        case "hello":
            switch -54 {
            case 32767:
                var v14 uint16 = (v11 | v9)
                _ = v14
            case 32768:
                v15 := GenS2{F0: 42}
                _ = v15
            case 0:
                v3 -= 718
            default:
                if len(v10) > 0 {
                    v10[0] = (fi13 % ((fi13) | 1))
                }
            }
        }
        delete(v1, 4294967295)
        _ = fi13
    }
    if (!(v9 <= 1)) {
        fmt.Println(v11, v11)
        v2 = (^(v2))
    } else {
        v3 = (v3 % ((v3) | 1))
        if (false || true) {
            fmt.Println(v9, v3)
            fmt.Println(v9, v2, v2)
        } else {
            v16 := GenS1{F0: "\n\t", F1: 1, F2: (v3 ^ (v3 << ((168) & 63)))}
            _ = v16
        }
        var v17 uint8 = 2
        _ = v17
    }
    if (((!true) || false) || true) {
        switch ("" + "\n\t") {
        case "\n\t":
            var v18 bool = false
            _ = v18
        case "hello":
            v3 = v3
        }
        v3 = v3
        v3 = 255
    }
    if true {
        v19 := GenS1{F0: "foo bar", F1: 71, F2: (v3 & v3)}
        v9 -= (v2 - (v2 / ((2) | 1)))
        _ = v19
    } else {
        v20 := GenS3{F0: 42, F1: -618, F2: 1000}
        _ = v20
    }
    for fi21 := 0; fi21 < 5; fi21++ {
        var v22 int32 = -1
        var v23 uint = (v3 / ((65535) | 1))
        _ = fi21
        _ = v22
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
