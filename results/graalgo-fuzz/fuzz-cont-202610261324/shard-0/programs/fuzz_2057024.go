package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 int }

func main() {
    v1 := []uint32{2, 42}
    v2 := make(map[uint32]bool)
    var v3 uint16 = 370
    v4 := GenS1{F0: 42, F1: -828, F2: 65535}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v5 := GenS1{F0: 42, F1: 0, F2: 44}
    var v6 uint8 = 42
    v3 = (^(v3))
    v6 = v6
    for fi7 := 0; fi7 < 9; fi7++ {
        if (-1 <= 84) {
            v8 := GenS1{F0: 1, F1: 722, F2: 4}
            var v9 uint16 = 162
            _ = v8
            _ = v9
        }
        var v10 uint8 = v6
        var v11 int64 = 9223372036854775807
        _ = fi7
        _ = v10
        _ = v11
    }
    var v12 string = "\n\t"
    v1 = append(v1, 1)
    for fi13 := 0; fi13 < 20; fi13++ {
        if len(v1) > 0 {
            v1[0] = 2
        }
        v14 := GenS2{F0: 42}
        v15 := GenS3{F0: v12, F1: fi13}
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 string = ""
    if true {
        var v17 uint64 = 1
        _ = v17
    }
    switch v12 {
    case "hello":
        if (v16 > v12) {
            if (!(2147483647 <= int32(v6))) {
                var v18 int32 = 1
                _ = v18
            }
            v16 = (("\n\t" + v16) + ("a" + "foo bar"))
            var v19 bool = true
            _ = v19
        } else {
            _ = len(v1)
            if true {
                var v20 float32 = (1e10 - 0.1)
                _ = v20
            } else {
                var v21 bool = (59 > 1)
                v21 = (!v21)
                _ = v21
            }
        }
    case "":
        _ = len(v1)
    case "a":
        v6 += v6
    }
    if false {
        fmt.Println(v3, v3)
    }
    for fi22 := 0; fi22 < 5; fi22++ {
        v2[2147483647] = false
        var v23 uint32 = 42
        _ = fi22
        _ = v23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
