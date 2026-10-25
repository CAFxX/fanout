package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }

func main() {
    v1 := []int{9223372036854775807, 0, -964}
    v2 := []uint32{}
    v3 := make(map[string]string)
    for fi4 := 0; fi4 < 10; fi4++ {
        var v5 float32 = (0.0 - (1.5 - 0.1))
        v5 = 1.5
        var v6 int32 = -1
        _ = fi4
        _ = v5
        _ = v6
    }
    if ((!true) && (!(1 == -1))) {
        var v7 uint = 2
        _ = v7
    }
    if true {
        var v8 uint64 = 395
        _ = v8
    }
    var v9 int8 = 98
    switch -9223372036854775808 {
    case 9223372036854775807:
        v9 = 1
    }
    var v10 int32 = 1
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v9 -= 91
        v11 := GenS1{F0: 0, F1: (v9 >= (v9 / (((-(v9))) | 1)))}
        _ = v11
    }
    v12 := GenS1{F0: 143, F1: ("a" != "foo bar")}
    v2 = append(v2, 2147483648)
    fmt.Println(v9, v10)
    if (!(65535 < 0)) {
        var v13 int32 = 42
        _ = len(v1)
        _ = v13
    }
    if ("foo bar" >= ("\n\t" + "a")) {
        for fi14 := 0; fi14 < 1; fi14++ {
            for fi15 := 0; fi15 < 8; fi15++ {
                if false {
                    fmt.Println(fi15, fi15, fi14)
                }
                _ = fi15
            }
            _ = fi14
        }
        v9 -= (v9 - v9)
    } else {
        if (0.1 >= 1e10) {
            var v16 int16 = -108
            var v17 int32 = 42
            fmt.Println(v17, v10)
            _ = v16
            _ = v17
        }
    }
    var v18 uint = 797
    var v19 int16 = -32768
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v9
}
