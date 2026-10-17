package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
func genFn0(p0 int, p1 uint16) uint {
    switch p0 {
    case 9223372036854775807:
        var v5 float32 = 1.5
        _ = v5
    }
    var v6 int = (p0 / (((p0 - p0)) | 1))
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := GenS1{F0: 127, F1: 42}
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v6
    return 780
}

func main() {
    var v1 uint = 65535
    var v2 uint16 = 2
    var v3 uint16 = v2
    v4 := make(map[string]bool)
    v9 := GenS1{F0: 2, F1: 86}
    var v10 float32 = -2.25
    if true {
        var v11 uint64 = 526
        _ = v11
    } else {
        var v12 uint64 = 410
        _ = v12
    }
    var v13 bool = false
    switch "hello" {
    case "hello":
        v1 = v1
    case "foo bar":
        if v13 {
            v14 := GenS1{F0: 128, F1: 2147483647}
            var v15 uint = genFn0(-1, 2)
            _ = v14
            _ = v15
        }
    case "":
        var v16 uint = v1
        _ = v16
    default:
        v3 += v3
    }
    var v17 int64 = -566
    switch 1 {
    case 42:
        v2 = (v3 % (((v3 + v3)) | 1))
    default:
        fmt.Println(v17)
    }
    var v18 int8 = 0
    if v13 {
        v18 += int8(v2)
        var v19 bool = v13
        _ = v19
    } else {
        v1 += 65535
        v20 := GenS1{F0: 1, F1: 2147483647}
        v21 := []int{0}
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        v23 := []int{4294967295, (fi22 % ((fi22) | 1)), (^(fi22)), (fi22 + (-(fi22)))}
        var v24 int16 = -1
        v18 += v18
        _ = fi22
        _ = v23
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
