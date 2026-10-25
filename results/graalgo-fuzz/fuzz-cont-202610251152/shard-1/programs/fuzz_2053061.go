package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0() uint8 {
    var v7 uint16 = 558
    if false {
        v8 := GenS1{F0: v7}
        var v9 int64 = -9223372036854775808
        _ = v8
        _ = v9
    }
    switch "" {
    case "":
        v7 = (v7 | (v7 ^ 1))
    case "a":
        for fi10 := 0; fi10 < 6; fi10++ {
            if (!(640 != 4294967295)) {
                v11 := GenS1{F0: 0}
                fmt.Println(fi10, fi10)
                _ = v11
            } else {
                v12 := GenS1{F0: 108}
                var v13 uint32 = 4294967295
                _ = v12
                _ = v13
            }
            _ = fi10
        }
    case "hello":
        var v14 float32 = (1.5 - 3.14159)
        _ = v14
    default:
        v7 = v7
    }
    _ = v7
    return 213
}

func main() {
    var v1 uint64 = 2
    v2 := []int{42, 42}
    var v3 string = "hello"
    var v4 int64 = 9223372036854775807
    var v5 uint = 65535
    v6 := make(map[int]int)
    v1 = 934
    if ((!(42 >= -624)) || (0 <= -981)) {
        v15 := GenS1{F0: 42}
        v6[4294967295] = 4294967295
        var v16 int32 = 0
        _ = v15
        _ = v16
    } else {
        var v17 uint64 = (v1 >> ((v5) & 63))
        _ = v17
    }
    fmt.Println(v5)
    var v18 uint8 = 42
    v5 += 344
    fmt.Println(v3, v3)
    var v19 int16 = -32768
    fmt.Println(v3, v5)
    for fi20 := 0; fi20 < 12; fi20++ {
        var v21 uint64 = v1
        for fi22 := 0; fi22 < 18; fi22++ {
            fmt.Println(fi20)
            switch fi22 {
            case 425:
                var v23 int64 = (^(v4))
                _ = v23
            }
            _ = fi22
        }
        var v24 float64 = (1.5 * 1.5)
        _ = fi20
        _ = v21
        _ = v24
    }
    if ((^(v5)) >= 461) {
        fmt.Println(v1)
        _ = len(v6)
        var v25 int32 = 2147483647
        _ = v25
    }
    fmt.Println(v5, v18)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
