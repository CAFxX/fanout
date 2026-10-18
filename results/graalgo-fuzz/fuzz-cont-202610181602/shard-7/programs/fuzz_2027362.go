package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 string }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 int }

func main() {
    v1 := []int{}
    var v2 int = 32768
    var v3 int = v2
    v4 := make(map[uint32]int)
    switch v3 {
    case 2147483647:
        v5 := GenS1{F0: -1, F1: "foo bar", F2: ("a" + "hello")}
        _ = v5
    case -9223372036854775808:
        var v6 uint = 42
        _ = v6
    default:
        switch v2 {
        case -426:
            var v7 bool = false
            _ = v7
        case 1:
            v8 := GenS1{F0: v2, F1: "hello", F2: "\n\t"}
            _ = v8
        case 32768:
            v9 := GenS2{F0: 896, F1: 0}
            _ = v9
        }
    }
    _ = len(v1)
    var v10 uint64 = 2
    var v11 float32 = (0.0 * (0.0 + 1e10))
    v12 := v4[1]
    if false {
        if false {
            v3 = (v3 & v2)
            v3 = -1
            fmt.Println(v2)
        } else {
            var v13 uint = 426
            _ = v13
        }
        v11 -= ((float32(-2.25) - v11) + ((1e10 - ((v11 / (v11 * float32((1e10 / 1e10)))) + float32(-2.25))) / float32(((0.1 * (1e10 * 0.1)) * (((0.1 - 1.5) * float64((v11 * 0.1))) + ((0.1 - 3.14159) - (float64(float32(-2.25)) + 1.5)))))))
    }
    var v14 int32 = -1
    switch "a" {
    case "foo bar":
        if true {
            var v15 uint16 = 665
            _ = v15
        }
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v16 := GenS2{F0: 42, F1: 200}
    if ((!(-32768 == 42)) || (368 < 65535)) {
        var v17 float64 = 1.5
        _ = v17
    }
    var v18 string = "hello"
    if (2 <= 178) {
        v18 = (v18 + ("foo bar" + v18))
        v2 += v12
    } else {
        v14 -= v14
    }
    for fi19 := 0; fi19 < 2; fi19++ {
        v20 := GenS2{F0: 4294967295, F1: 255}
        _ = fi19
        _ = v20
    }
    v21 := GenS3{F0: 9223372036854775807}
    if len(v1) > 0 {
        v1[0] = (-(v3))
    }
    fmt.Println(v11, v12, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v16
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
