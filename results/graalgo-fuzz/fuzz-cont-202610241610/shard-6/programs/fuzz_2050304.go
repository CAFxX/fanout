package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 bool; F2 int32 }
type GenS3 struct { F0 int16; F1 int16; F2 int }
type GenS4 struct { F0 int64; F1 int16 }
func genFn0() int16 {
    if (false && (0.1 < 1e10)) {
        if ((!true) || (461 == -1)) {
            var v7 float32 = 0.0
            _ = v7
        }
        v8 := GenS1{F0: 1}
        var v9 string = "hello"
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        fmt.Println(fi10)
        _ = fi10
    }
    if false {
        var v11 int = 32767
        _ = v11
    }
    v12 := GenS1{F0: -1}
    v13 := GenS1{F0: 0}
    _ = v12
    _ = v13
    return 42
}
func genFn1(p0 string, p1 bool) int16 {
    v14 := GenS1{F0: 0}
    v15 := []int{32768, -338, 627, 32767}
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    return genFn0()
}

func main() {
    var v1 int32 = 2147483647
    var v2 uint = 0
    var v3 uint = 1
    var v4 int64 = 673
    var v5 int = 4294967295
    v6 := make(map[int]string)
    v16 := v6[(^(v5))]
    v17 := []int{(-(v5)), v5, int(v2), v5}
    if (v4 == v4) {
        v1 = v1
    } else {
        v4 += (-(v4))
        var v18 int64 = (v4 ^ (v4 + 42))
        var v19 uint64 = 1
        _ = v18
        _ = v19
    }
    switch "a" {
    case "hello":
        _ = len(v6)
    case "":
        v3 -= v2
    case "foo bar":
        v20 := GenS2{F0: (v1 / ((v1) | 1)), F1: false, F2: int32(v3)}
        _ = v20
    default:
        v21 := GenS3{F0: -32768, F1: 1, F2: 506}
        _ = v21
    }
    var v22 uint8 = 1
    v17 = append(v17, v5)
    v4 = (-(v4))
    v2 = 18446744073709551615
    v3 = v3
    var v23 uint = (^(v3))
    v3 = v2
    switch 1 {
    case 4294967295:
        v3 = v2
    case -9223372036854775808:
        if false {
            if (v3 < 6) {
                v24 := GenS4{F0: (v4 ^ int64(v22)), F1: -1}
                _ = v24
            }
            var v25 int32 = v1
            var v26 float64 = (3.14159 + ((3.14159 * float64(((0.1 + (1.5 / 1.0)) - -2.25))) * -2.25))
            _ = v25
            _ = v26
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
