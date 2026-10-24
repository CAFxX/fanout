package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 int8; F1 string }
func genFn0(p0 int64, p1 int8) int16 {
    p1 += p1
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 int64 = 1
    var v2 float32 = ((float32((((0.0 + 1.5) * (1.5 * (1e10 + -2.25))) + (1e10 - (0.1 / 1.0)))) * 3.14159) * 0.0)
    v3 := make(map[int]uint32)
    v4 := make(map[int]int)
    v5 := make(map[uint32]uint32)
    if (128 >= 2) {
        switch ((("\n\t" + "foo bar") + "foo bar") + "foo bar") {
        case "a":
            fmt.Println(v1)
        case "foo bar":
            v6 := GenS1{F0: 255}
            _ = v6
        case "\n\t":
            v7 := []int{-2147483648, 354}
            _ = v7
        default:
            v8 := GenS2{F0: true, F1: 1}
            _ = v8
        }
    } else {
        v1 = v1
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 uint = 406
            switch "hello" {
            case "foo bar":
                v2 = -2.25
            case "":
                fmt.Println(v1)
            default:
                fmt.Println(v2, v2, v10)
            }
            var v11 int64 = (v1 * v1)
            _ = fi9
            _ = v10
            _ = v11
        }
        var v12 bool = false
        _ = v12
    }
    v1 = (v1 << ((286) & 63))
    v1 += (^(v1))
    switch 42 {
    case -265:
        v1 -= v1
    case 32767:
        if false {
            v13 := []int{}
            if false {
                fmt.Println(v1, v1)
                var v14 uint16 = 32767
                _ = len(v3)
                _ = v14
            }
            v15 := GenS3{F0: 42, F1: "a"}
            _ = v13
            _ = v15
        }
    case 2147483647:
        v16 := GenS2{F0: (525 < -1), F1: 0}
        _ = v16
    default:
        var v17 uint32 = 2147483648
        _ = v17
    }
    v2 = v2
    switch "hello" {
    case "":
        v18 := GenS4{F0: 99, F1: (("\n\t" + "a") + "hello")}
        _ = v18
    }
    v1 = -1
    v1 += (v1 + -9223372036854775808)
    fmt.Println(v2, v2)
    v1 += (v1 - (^(v1)))
    switch 1 {
    case 287:
        var v19 uint = 0
        _ = v19
    case -1:
        fmt.Println(v2, v2)
    case 42:
        var v20 uint64 = 202
        _ = v20
    default:
        var v21 string = ("hello" + "foo bar")
        _ = v21
    }
    fmt.Println(v1, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
