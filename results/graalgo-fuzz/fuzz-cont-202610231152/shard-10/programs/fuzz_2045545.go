package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 uint }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int64 = -489
    v3 := make(map[uint32]string)
    var v4 uint8 = 127
    v5 := make(map[string]string)
    if false {
        var v6 string = "\n\t"
        v3[201] = v6
        _ = v6
    } else {
        if (!(!false)) {
            fmt.Println(v2)
            v2 = 9223372036854775807
            v4 += v4
        } else {
            fmt.Println(v2, v4)
        }
        var v7 uint64 = 42
        _ = v7
    }
    var v8 int64 = (v2 << ((139) & 63))
    if false {
        v4 = (v4 & 0)
    }
    v4 = (^(v4))
    v4 += v4
    v4 = (v4 << ((183) & 63))
    v9 := v5[("a" + "foo bar")]
    v8 += -1
    if ((((v9 + ((v9 + "") + (v9 + v9))) + ("a" + "hello")) > (v9 + (v9 + (v9 + "hello")))) || (402 < 1)) {
        v10 := GenS1{F0: 0}
        _ = v10
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        fmt.Println(v2, v2)
        var v12 int8 = -28
        var v13 uint32 = 618
        _ = fi11
        _ = v12
        _ = v13
    }
    v4 -= 42
    switch 127 {
    case 128:
        v14 := GenS1{F0: 74}
        _ = v14
    case 4294967295:
        var v15 uint32 = 65535
        _ = v15
    case 32768:
        v16 := GenS2{F0: 42}
        _ = v16
    default:
        if (!(2147483647 < 0)) {
            v17 := []int{9223372036854775807, -9223372036854775808, 42}
            switch -18 {
            case -1:
                v2 = v8
            }
            _ = v17
        }
    }
    v18 := GenS3{F0: v8, F1: 2}
    fmt.Println(v8, v4, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
