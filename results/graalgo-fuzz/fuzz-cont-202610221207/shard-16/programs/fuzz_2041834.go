package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 uint }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 uint64; F1 int64; F2 uint16 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 int32, p1 int64) bool {
    v7 := GenS1{F0: 128, F1: 1, F2: 212}
    p1 = p1
    _ = p0
    _ = p1
    _ = v7
    return false
}

func main() {
    v1 := make(map[int]int)
    var v2 int16 = -1
    var v3 int32 = 0
    var v4 uint32 = 42
    v5 := make(map[string]int)
    v6 := make(map[string]int)
    if (-128 <= -1) {
        var v8 int32 = (v3 >> ((6) & 63))
        v9 := GenS2{F0: 229, F1: 2147483647}
        _ = v8
        _ = v9
    } else {
        var v10 int = 2147483647
        v3 += v3
        v10 -= v10
        _ = v10
    }
    v3 = (^(v3))
    var v11 uint = 0
    var v12 float32 = 0.1
    v13 := []int{1, 42, -478}
    v11 = v11
    switch int(v2) {
    case -9223372036854775808:
        fmt.Println(v12, v11, v11)
    default:
        v2 -= (v2 & (v2 | (v2 << ((18446744073709551615) & 63))))
    }
    v5[("\n\t" + ((("hello" + "hello") + "\n\t") + "hello"))] = 32768
    if true {
        var v14 string = "a"
        fmt.Println(v12, v12, v4)
        _ = v14
    } else {
        v15 := GenS3{F0: 42, F1: 9223372036854775807, F2: 2}
        var v16 uint8 = 1
        _ = len(v13)
        _ = v15
        _ = v16
    }
    switch -9223372036854775808 {
    case -387:
        v11 = (v11 % ((18446744073709551615) | 1))
    default:
        switch -349 {
        case -1:
            var v17 uint8 = 128
            _ = v17
        case 32768:
            fmt.Println(v2)
        case 1:
            v18 := GenS1{F0: 128, F1: 287, F2: v11}
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 20; fi19++ {
        var v20 uint8 = 0
        _ = fi19
        _ = v20
    }
    if true {
        v6[""] = 178
        var v21 bool = (477 >= 637)
        v3 -= 1
        _ = v21
    } else {
        var v22 string = ("foo bar" + "\n\t")
        var v23 int16 = v2
        if (42 == 32767) {
            var v24 int64 = 9223372036854775807
            for fi25 := 0; fi25 < 3; fi25++ {
                var v26 uint = v11
                _ = len(v1)
                v4 += (^(v4))
                _ = fi25
                _ = v26
            }
            delete(v5, (("\n\t" + v22) + ("" + v22)))
            _ = v24
        } else {
            v27 := GenS4{F0: 1}
            v12 = 1e10
            v11 += v11
            _ = v27
        }
        _ = v22
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
