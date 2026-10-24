package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 int16 }
func genFn0(p0 int16) uint32 {
    var v4 int8 = 1
    switch ("" + "") {
    case "":
        for fi5 := 0; fi5 < 2; fi5++ {
            if (42 >= -1) {
                var v6 int64 = -1
                var v7 int32 = 708
                _ = v6
                _ = v7
            }
            fmt.Println(fi5)
            _ = fi5
        }
    }
    var v8 int8 = v4
    var v9 uint16 = 255
    _ = p0
    _ = v4
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := make(map[string]string)
    var v2 int64 = 9223372036854775807
    var v3 uint = 255
    fmt.Println(v3, v3)
    var v10 int = 127
    fmt.Println(v2, v3, v2)
    v11 := []int{(v10 << (((v3 << ((255) & 63))) & 63)), (v10 ^ (v10 % ((v10) | 1)))}
    v12 := GenS1{F0: 467}
    for fi13 := 0; fi13 < 15; fi13++ {
        v10 += fi13
        _ = fi13
    }
    v1[("\n\t" + "\n\t")] = ""
    var v14 int32 = 46
    v3 -= 569
    v2 = (v2 ^ (^(v2)))
    if (false && ((false || (0.1 > 1.5)) || true)) {
        v14 -= 0
        v2 = v2
    }
    fmt.Println(v10, v10, v10)
    var v15 int32 = 2147483647
    v3 = (v3 << ((0) & 63))
    switch "a" {
    case "\n\t":
        v16 := GenS2{F0: (v14 >= v14), F1: (v2 << ((581) & 63))}
        _ = v16
    case "a":
        v17 := GenS1{F0: v10}
        _ = v17
    case "foo bar":
        if len(v11) > 0 {
            v11[0] = v10
        }
    }
    if (1 != 1) {
        v14 -= -1
    }
    v3 = (v3 % (((^(v3))) | 1))
    v18 := GenS3{F0: 32767}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
}
