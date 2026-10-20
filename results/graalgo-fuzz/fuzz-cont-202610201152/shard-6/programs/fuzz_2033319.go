package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int8 }
func genFn0(p0 int, p1 int) uint32 {
    fmt.Println(p1, p0)
    for fi5 := 0; fi5 < 18; fi5++ {
        var v6 uint64 = 1
        v7 := GenS1{F0: 72, F1: 128, F2: 42}
        _ = fi5
        _ = v6
        _ = v7
    }
    fmt.Println(p1, p0)
    if true {
        fmt.Println(p0)
    }
    if (false || false) {
        v8 := GenS1{F0: 42, F1: 41, F2: -2147483648}
        v9 := GenS2{F0: p1}
        var v10 int16 = -32768
        _ = v10
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    return 738
}
func genFn1(p0 int16) int32 {
    p0 = (p0 + (p0 >> ((4294967295) & 63)))
    if true {
        p0 += (p0 << ((371) & 63))
    } else {
        switch ("" + (("" + "") + (("foo bar" + "a") + "foo bar"))) {
        case "foo bar":
            switch -2147483648 {
            case 9223372036854775807:
                v11 := GenS2{F0: 1}
                _ = v11
            case -426:
                var v12 uint = 381
                _ = v12
            default:
                v13 := GenS3{F0: (("" + "foo bar") + "\n\t")}
                _ = v13
            }
        }
    }
    p0 += (p0 - 1000)
    _ = p0
    return 16
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []string{(("\n\t" + "") + "\n\t")}
    v3 := make(map[string]bool)
    v4 := []int{-9223372036854775808}
    v1 += v1
    v14 := GenS3{F0: "foo bar"}
    v15 := GenS4{F0: -128}
    v1 -= (v1 - v1)
    fmt.Println(v1)
    var v16 int64 = (v1 - v1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v17 uint64 = 4294967295
    var v18 uint32 = 716
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
