package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 int8; F2 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 bool; F1 int32 }

func main() {
    v1 := make(map[string]int)
    var v2 uint = 18446744073709551615
    v3 := make(map[string]uint32)
    var v4 int16 = -673
    var v5 string = "foo bar"
    var v6 int64 = 1
    for fi7 := 0; fi7 < 14; fi7++ {
        v3["hello"] = 1
        _ = fi7
    }
    v8 := GenS1{F0: v6, F1: uint16(v2)}
    switch v5 {
    case "a":
        v9 := GenS1{F0: 9223372036854775807, F1: 2}
        _ = v9
    case "\n\t":
        var v10 string = v5
        _ = v10
    }
    switch v5 {
    case "a":
        v11 := GenS1{F0: v6, F1: 851}
        _ = v11
    case "foo bar":
        v12 := GenS2{F0: 42}
        _ = v12
    case "\n\t":
        v13 := GenS3{F0: v2, F1: -71, F2: v6}
        _ = v13
    }
    var v14 uint8 = 135
    v15 := v3["foo bar"]
    if false {
        if ((false && false) && false) {
            var v16 uint32 = (v15 % ((v15) | 1))
            v17 := []int{2147483647, -565}
            _ = v16
            _ = v17
        }
    } else {
        v18 := GenS4{F0: -9223372036854775808}
        if ((612 <= 32767) && true) {
            v2 += v2
        }
        _ = v18
    }
    fmt.Println(v15, v14, v6)
    var v19 bool = false
    var v20 uint = v2
    v21 := GenS5{F0: (-2.25 != ((float64(0.0) + float64(1.5)) / 1.5)), F1: -2147483648}
    var v22 string = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
