package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 int16 }
type GenS2 struct { F0 int8; F1 uint8; F2 bool }
type GenS3 struct { F0 int; F1 int }
func genFn0() uint32 {
    if (127 <= 2) {
        var v5 int32 = 760
        v5 = (v5 - -2147483648)
        v6 := []int{1, 1, -1, 128}
        _ = v5
        _ = v6
    } else {
        v7 := GenS1{F0: 65535, F1: (1 == 2147483647), F2: 42}
        var v8 bool = true
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        switch "foo bar" {
        case "a":
            v10 := []int{}
            _ = v10
        case "":
            v11 := GenS2{F0: 66, F1: 255, F2: ((1 > 42) || (!(fi9 >= (fi9 + (fi9 ^ (fi9 - 4294967295))))))}
            _ = v11
        case "foo bar":
            var v12 int16 = 0
            _ = v12
        }
        _ = fi9
    }
    switch ("a" + "hello") {
    case "hello":
        var v13 uint8 = 128
        _ = v13
    case "\n\t":
        var v14 int16 = -652
        _ = v14
    case "foo bar":
        var v15 int64 = -9223372036854775808
        _ = v15
    default:
        var v16 float32 = ((float32(((0.0 * (1.5 / float64(((((0.0 - float32((float64(1e10) * 1.5))) / 3.14159) + (0.1 - float32((((((3.14159 + 1e10) / (0.1 - 0.0)) * 3.14159) + 1.5) - 1.5)))) + -2.25)))) * 3.14159)) * (-2.25 / -2.25)) + 3.14159)
        _ = v16
    }
    return 0
}

func main() {
    v1 := make(map[int]int)
    v2 := []int{-9223372036854775808}
    var v3 int32 = 0
    v4 := []int{-2147483648, -808}
    switch 127 {
    case 105:
        if len(v2) > 0 {
            v2[0] = 490
        }
    }
    v3 -= 42
    v3 -= v3
    var v17 bool = false
    var v18 float32 = 1.5
    var v19 bool = false
    if (0.1 > ((float32(0.1) - 1.5) + (v18 * (v18 / (float32(0.0) / (v18 + v18)))))) {
        if false {
            fmt.Println(v18, v19, v18)
            if false {
                v3 = v3
            } else {
                v20 := GenS1{F0: 701, F1: (v18 < 1e10), F2: 32767}
                _ = v20
            }
        }
        v3 = v3
        if v19 {
            v21 := GenS3{F0: -1, F1: 9223372036854775807}
            var v22 uint64 = 0
            v23 := GenS1{F0: 84, F1: true, F2: 1}
            _ = v21
            _ = v22
            _ = v23
        }
    } else {
        v17 = true
    }
    var v24 uint32 = 0
    v17 = (9223372036854775807 < 42)
    switch -9223372036854775808 {
    case 0:
        fmt.Println(v24)
    default:
        var v25 int8 = -96
        _ = v25
    }
    var v26 int8 = 127
    v27 := GenS3{F0: int(v24), F1: -9223372036854775808}
    if v19 {
        switch 0 {
        case 32767:
            v26 += v26
        case 302:
            fmt.Println(v19)
        case 4294967295:
            for fi28 := 0; fi28 < 1; fi28++ {
                _ = len(v1)
                _ = fi28
            }
        }
    } else {
        if v19 {
            var v29 uint64 = 18446744073709551615
            _ = v29
        } else {
            v18 = (v18 / 1.0)
        }
        v19 = v19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
