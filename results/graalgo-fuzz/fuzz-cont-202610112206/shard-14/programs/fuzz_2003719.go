package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 uint32 }
func genFn0() int32 {
    var v5 uint16 = 255
    for fi6 := 0; fi6 < 3; fi6++ {
        switch fi6 {
        case -9223372036854775808:
            if (!(0 != 1)) {
                v5 = (v5 << ((65535) & 63))
                v7 := []int{(^(fi6)), (fi6 | (fi6 - -1)), fi6, fi6}
                _ = v7
            }
        }
        _ = fi6
    }
    var v8 uint32 = 0
    _ = v5
    _ = v8
    return 0
}

func main() {
    var v1 string = "a"
    v2 := []uint32{}
    v3 := make(map[int]int)
    var v4 uint64 = 4294967295
    switch ((v1 + v1) + v1) {
    case "":
        var v9 uint64 = 42
        _ = v9
    case "a":
        v10 := GenS1{F0: 2147483647}
        _ = v10
    }
    if (1 != 468) {
        if (1.5 > (3.14159 * -2.25)) {
            if true {
                fmt.Println(v4, v4)
                v11 := GenS2{F0: -902, F1: 1}
                _ = v11
            } else {
                var v12 uint64 = (v4 | v4)
                var v13 uint64 = v4
                var v14 string = v1
                _ = v12
                _ = v13
                _ = v14
            }
            var v15 int = 42
            _ = v15
        } else {
            if false {
                _ = len(v3)
            }
            v16 := GenS1{F0: genFn0()}
            fmt.Println(v4)
            _ = v16
        }
    }
    v4 = v4
    v1 = v1
    var v17 uint8 = 0
    fmt.Println(v4, v1)
    for fi18 := 0; fi18 < 12; fi18++ {
        v19 := GenS2{F0: -468, F1: 558}
        v20 := GenS1{F0: 42}
        _ = fi18
        _ = v19
        _ = v20
    }
    fmt.Println(v4, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
