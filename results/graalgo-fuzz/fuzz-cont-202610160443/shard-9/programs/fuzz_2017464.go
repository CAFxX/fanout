package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 uint32 }
func genFn0(p0 int32, p1 uint32) int16 {
    p0 = (p0 << ((42) & 63))
    if ((871 != 65535) || (1e10 <= 0.1)) {
        v7 := GenS1{F0: 39, F1: 1, F2: 42}
        _ = v7
    } else {
        var v8 uint32 = p1
        _ = v8
    }
    v9 := GenS1{F0: 545, F1: 2, F2: 2}
    if (4294967295 >= p1) {
        p1 = 193
        switch -313 {
        case 285:
            p0 = -2147483648
        case 32767:
            v10 := GenS1{F0: 126, F1: 128, F2: 4294967295}
            _ = v10
        default:
            fmt.Println(p0, p0)
        }
    } else {
        switch 1 {
        case -559:
            v11 := GenS2{F0: p0}
            _ = v11
        case -20:
            p0 += (p0 | (p0 & p0))
        case -535:
            p1 += p1
        default:
            p1 -= 2147483648
        }
    }
    _ = p0
    _ = p1
    _ = v9
    return 0
}

func main() {
    v1 := make(map[string]string)
    var v2 float64 = 1.5
    v3 := make(map[string]string)
    var v4 uint = 649
    var v5 float32 = 3.14159
    v6 := []int{127, 128, 32767, 350}
    v4 += 964
    var v12 int8 = 100
    _ = len(v1)
    v13 := GenS3{F0: 32767, F1: 0}
    v14 := GenS1{F0: 1, F1: 2, F2: (v4 - (v4 ^ (v4 - (v4 & (v4 ^ v4)))))}
    v1["foo bar"] = ((("foo bar" + (("hello" + "hello") + (("foo bar" + "a") + "foo bar"))) + "hello") + "hello")
    for fi15 := 0; fi15 < 15; fi15++ {
        var v16 uint = (v4 % ((1) | 1))
        fmt.Println(v5)
        _ = fi15
        _ = v16
    }
    fmt.Println(v5)
    var v17 uint32 = 65535
    if (true || (33 >= 42)) {
        v4 = (^(v4))
    }
    if false {
        var v18 int8 = v12
        var v19 int32 = 0
        _ = v18
        _ = v19
    } else {
        v4 -= v4
    }
    v1["foo bar"] = "hello"
    var v20 uint16 = 256
    switch "hello" {
    case "\n\t":
        for fi21 := 0; fi21 < 19; fi21++ {
            var v22 uint64 = 908
            _ = fi21
            _ = v22
        }
    case "a":
        v2 = (v2 - ((v2 * 1e10) / 0.1))
    }
    v4 = 255
    _ = len(v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
