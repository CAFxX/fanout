package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 uint; F2 int }

func main() {
    var v1 float32 = -2.25
    v2 := make(map[uint32]string)
    var v3 uint32 = 2
    var v4 uint32 = 2
    switch 127 {
    case -9223372036854775808:
        fmt.Println(v1)
    }
    v5 := GenS1{F0: 1}
    for fi6 := 0; fi6 < 18; fi6++ {
        for fi7 := 0; fi7 < 17; fi7++ {
            v8 := []int{}
            _ = fi7
            _ = v8
        }
        fmt.Println(v3)
        _ = fi6
    }
    var v9 uint = 65535
    if true {
        var v10 uint64 = uint64(v3)
        _ = v10
    } else {
        v11 := []int{690}
        switch 9223372036854775807 {
        case 128:
            switch "hello" {
            case "foo bar":
                v4 += 765
            case "a":
                var v12 int32 = -486
                _ = v12
            case "":
                if false {
                    fmt.Println(v9, v1, v1)
                    var v13 uint16 = 313
                    _ = v13
                } else {
                    v9 -= v9
                    var v14 uint16 = 2
                    var v15 uint64 = 25
                    _ = v14
                    _ = v15
                }
            }
        }
        v16 := GenS2{F0: 1, F1: v9, F2: 0}
        _ = v11
        _ = v16
    }
    if ((!(992 <= 860)) || false) {
        v17 := v2[v4]
        var v18 int32 = 1
        v3 = (v4 ^ 42)
        _ = v17
        _ = v18
    }
    var v19 uint32 = (v3 << (((v9 >> (((v9 / ((255) | 1))) & 63))) & 63))
    v20 := v2[2147483647]
    for fi21 := 0; fi21 < 20; fi21++ {
        var v22 int = 136
        _ = fi21
        _ = v22
    }
    v23 := GenS1{F0: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
