package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 int16 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 int16 }
func genFn0(p0 int8, p1 int16) bool {
    switch 487 {
    case 2147483647:
        var v4 int8 = p0
        _ = v4
    case -703:
        v5 := GenS1{F0: "hello", F1: 42, F2: p1}
        _ = v5
    default:
        var v6 int8 = p0
        _ = v6
    }
    var v7 int8 = (-(p0))
    _ = p0
    _ = p1
    _ = v7
    return (42 != 32768)
}

func main() {
    var v1 int32 = 819
    var v2 int32 = (^(v1))
    v3 := make(map[uint32]uint32)
    switch 4294967295 {
    case 127:
        v8 := []int{0, -255, 32768, 656}
        _ = v8
    case 4294967295:
        fmt.Println(v2, v2)
    case 2147483647:
        fmt.Println(v1, v2)
    }
    fmt.Println(v2, v2)
    fmt.Println(v2)
    v9 := GenS2{F0: 42, F1: 2}
    v1 = (v1 - (v1 >> ((4294967295) & 63)))
    for fi10 := 0; fi10 < 6; fi10++ {
        var v11 int64 = -1
        v2 = v1
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        var v13 int32 = (v2 | v2)
        switch (fi12 >> ((255) & 63)) {
        case 4294967295:
            v13 -= v1
        case -9223372036854775808:
            var v14 int32 = v13
            _ = v14
        case 545:
            fmt.Println(fi12, v1)
        }
        fmt.Println(v13)
        _ = fi12
        _ = v13
    }
    var v15 uint32 = 0
    switch (("\n\t" + ("\n\t" + "\n\t")) + "") {
    case "\n\t":
        v16 := GenS3{F0: 871}
        _ = v16
    case "foo bar":
        v17 := []int{4294967295, 1, -9223372036854775808}
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v9
}
