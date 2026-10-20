package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []int{548, -2147483648, -2147483648, -2147483648}
    var v3 float64 = 1.5
    switch 127 {
    case 128:
        if len(v2) > 0 {
            v2[0] = 405
        }
    }
    if (1 <= 69) {
        v4 := GenS1{F0: 2, F1: "a", F2: "hello"}
        _ = v4
    } else {
        var v5 int32 = 42
        _ = v5
    }
    v6 := GenS2{F0: 1}
    for fi7 := 0; fi7 < 10; fi7++ {
        v1 -= v1
        var v8 bool = false
        _ = fi7
        _ = v8
    }
    switch "foo bar" {
    case "a":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "hello":
        fmt.Println(v1)
    default:
        var v9 uint8 = 27
        _ = v9
    }
    if (!false) {
        v10 := GenS3{F0: 42}
        v11 := GenS3{F0: -1}
        _ = v10
        _ = v11
    }
    var v12 bool = (true || (255 == 2))
    v13 := GenS1{F0: 18446744073709551615, F1: "foo bar", F2: "foo bar"}
    var v14 uint8 = 33
    for fi15 := 0; fi15 < 10; fi15++ {
        v1 = v1
        _ = fi15
    }
    var v16 bool = (769 < 890)
    v2 = append(v2, 32768)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v6
}
