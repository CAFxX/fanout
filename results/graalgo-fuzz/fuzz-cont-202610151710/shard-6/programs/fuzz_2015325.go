package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 uint32 }
func genFn0() int {
    if false {
        var v7 int32 = 42
        _ = v7
    } else {
        v8 := []int{}
        _ = v8
    }
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 uint = 18446744073709551615
        _ = fi9
        _ = v10
    }
    switch -677 {
    case 2147483647:
        var v11 int8 = -1
        _ = v11
    case 32768:
        var v12 uint64 = 654
        _ = v12
    case -1:
        if (2 <= 127) {
            v13 := GenS1{F0: 18446744073709551615, F1: -32768}
            var v14 uint32 = 42
            _ = v13
            _ = v14
        } else {
            var v15 uint64 = 886
            _ = v15
        }
    default:
        var v16 uint32 = 2147483647
        _ = v16
    }
    return 0
}

func main() {
    var v1 int8 = 100
    var v2 bool = true
    var v3 bool = v2
    var v4 int16 = 0
    v5 := make(map[uint32]bool)
    v6 := []int{4294967295, 619, -84}
    v3 = (1 <= 0)
    v6 = append(v6, 943)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    switch "" {
    case "\n\t":
        var v17 int = 261
        _ = v17
    case "foo bar":
        fmt.Println(v4, v1)
    default:
        var v18 int32 = 1
        _ = v18
    }
    v19 := GenS1{F0: 4294967295, F1: (v4 + (v4 << ((65) & 63)))}
    var v20 float64 = (((((float64((0.1 - 0.0)) - (0.1 * -2.25)) * 0.0) - (0.0 + 3.14159)) / 3.14159) * ((1.5 - float64(float32(float64(-2.25)))) * float64(((1.5 + 1.5) + 0.1))))
    for fi21 := 0; fi21 < 17; fi21++ {
        if (v1 <= (v1 >> ((42) & 63))) {
            v22 := GenS2{F0: -786}
            _ = v22
        }
        v23 := GenS1{F0: 18446744073709551615, F1: (v4 & v4)}
        fmt.Println(fi21, v2, v20)
        _ = fi21
        _ = v23
    }
    v20 = 1.5
    v4 += (v4 >> ((42) & 63))
    switch -923 {
    case -2147483648:
        v24 := GenS2{F0: 0}
        _ = v24
    case 2147483647:
        v25 := GenS3{F0: "foo bar", F1: 108}
        _ = v25
    case -9223372036854775808:
        var v26 uint = 18446744073709551615
        _ = v26
    }
    v2 = (1 != v1)
    for fi27 := 0; fi27 < 16; fi27++ {
        var v28 int64 = -473
        _ = fi27
        _ = v28
    }
    fmt.Println(v2, v1, v3)
    for fi29 := 0; fi29 < 7; fi29++ {
        var v30 int16 = v4
        _ = fi29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
