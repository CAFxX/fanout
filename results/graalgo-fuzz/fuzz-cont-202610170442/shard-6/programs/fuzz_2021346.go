package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 string }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint32 = 4294967295
    var v3 int = -2147483648
    v4 := GenS1{F0: 101}
    fmt.Println(v2)
    fmt.Println(v2)
    if true {
        v5 := GenS2{F0: 255, F1: 1}
        _ = v5
    } else {
        switch -2147483648 {
        case 4294967295:
            fmt.Println(v2, v2)
        }
    }
    fmt.Println(v2, v2)
    v6 := GenS3{F0: 42, F1: v2}
    switch v3 {
    case 128:
        var v7 uint32 = (^(v2))
        _ = v7
    case 4294967295:
        fmt.Println(v2, v3)
    case 224:
        v8 := []int{}
        _ = v8
    }
    v2 = (v2 * v2)
    v9 := []int{}
    var v10 uint32 = 65535
    v10 = (v10 ^ v2)
    if (-1 >= -308) {
        v11 := v1[1]
        if (-1 == -1) {
            v2 += (v10 - (v2 >> ((199) & 63)))
            v11 = v11
            var v12 uint16 = 513
            _ = v12
        } else {
            switch "a" {
            case "foo bar":
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
            case "a":
                if len(v9) > 0 {
                    v9[0] = v11
                }
            case "hello":
                var v13 float64 = 0.1
                _ = v13
            }
            v14 := GenS3{F0: 324, F1: (v2 / ((2147483647) | 1))}
            _ = v14
        }
        v11 = 248
        _ = v11
    }
    for fi15 := 0; fi15 < 18; fi15++ {
        if true {
            v16 := GenS2{F0: 42, F1: 1}
            fmt.Println(v2, v10, v2)
            fmt.Println(v10, v2, v2)
            _ = v16
        } else {
            fmt.Println(v2, fi15)
            if false {
                v17 := GenS4{F0: ("" + ("hello" + "\n\t"))}
                _ = v17
            }
            var v18 uint64 = 0
            _ = v18
        }
        var v19 int8 = 1
        _ = fi15
        _ = v19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v9
}
