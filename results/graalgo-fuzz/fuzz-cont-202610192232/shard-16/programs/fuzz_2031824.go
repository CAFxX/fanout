package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 int32; F1 int16; F2 int8 }
func genFn0(p0 int32, p1 string) bool {
    if false {
        fmt.Println(p1, p0)
    } else {
        switch 2147483647 {
        case -1:
            var v5 uint64 = 1
            _ = v5
        case -9223372036854775808:
            var v6 string = "a"
            _ = v6
        case 127:
            v7 := GenS1{F0: int8(p0)}
            _ = v7
        }
    }
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 uint16 = 255
        switch fi8 {
        case 127:
            v10 := GenS1{F0: 0}
            _ = v10
        case 9223372036854775807:
            p1 = (p1 + p1)
        case -2147483648:
            v11 := GenS2{F0: 151, F1: p0}
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    v12 := GenS1{F0: 93}
    p1 = ""
    var v13 float32 = (3.14159 - 3.14159)
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return (false && ((p0 | 970) == p0))
}

func main() {
    var v1 uint16 = 139
    var v2 int8 = 127
    v3 := make(map[int]int)
    var v4 uint64 = 42
    switch ("foo bar" + "foo bar") {
    case "\n\t":
        v14 := GenS3{F0: 127, F1: 553}
        _ = v14
    case "a":
        switch 32768 {
        case 32767:
            var v15 float32 = 1e10
            _ = v15
        }
    default:
        v16 := []int{-860, 1}
        _ = v16
    }
    v17 := []int{0}
    for fi18 := 0; fi18 < 19; fi18++ {
        fmt.Println(v2, v4)
        _ = fi18
    }
    var v19 int8 = v2
    _ = len(v17)
    var v20 string = "a"
    v2 += 5
    v21 := GenS4{F0: 0, F1: 32767, F2: v19}
    switch 49 {
    case 920:
        v20 = "a"
    default:
        v19 = int8(v4)
    }
    v22 := GenS4{F0: 2147483647, F1: 309, F2: int8(v4)}
    v2 += 1
    if ((false || (!false)) && true) {
        fmt.Println(v4, v20)
        delete(v3, 128)
        v3[-9223372036854775808] = 9223372036854775807
    }
    v23 := GenS3{F0: 128, F1: 32767}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
