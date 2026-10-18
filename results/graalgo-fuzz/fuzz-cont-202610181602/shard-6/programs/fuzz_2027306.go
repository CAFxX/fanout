package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int16 }

func main() {
    var v1 string = ("foo bar" + "\n\t")
    var v2 int8 = 117
    v3 := []int{-972, -2147483648}
    v4 := make(map[string]int)
    v2 = 100
    v5 := GenS1{F0: 9223372036854775807, F1: 867, F2: 32767}
    if (32768 >= 127) {
        fmt.Println(v2, v2)
    }
    var v6 uint32 = 42
    v7 := GenS1{F0: 0, F1: 42, F2: 1}
    for fi8 := 0; fi8 < 18; fi8++ {
        v2 = v2
        if false {
            v2 -= -1
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            if false {
                fmt.Println(fi8)
            }
        }
        var v9 int8 = v2
        _ = fi8
        _ = v9
    }
    if (9223372036854775807 == 850) {
        v6 = (v6 >> ((65535) & 63))
        fmt.Println(v2, v1)
    } else {
        _ = len(v3)
        var v10 int8 = v2
        _ = v10
    }
    delete(v4, (v1 + v1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
