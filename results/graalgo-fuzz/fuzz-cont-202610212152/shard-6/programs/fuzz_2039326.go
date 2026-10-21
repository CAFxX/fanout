package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 uint16; F1 int }

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = 112
    var v3 int8 = v2
    v4 := []int{32767}
    v5 := make(map[string]uint32)
    v6 := make(map[uint32]int)
    switch "" {
    case "foo bar":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    fmt.Println(v3)
    v7 := GenS1{F0: 578}
    fmt.Println(v2)
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 float64 = (0.1 * 3.14159)
        for fi10 := 0; fi10 < 19; fi10++ {
            v11 := GenS2{F0: 42, F1: 0}
            v12 := GenS2{F0: 2147483647, F1: 26}
            _ = fi10
            _ = v11
            _ = v12
        }
        v13 := v6[65535]
        _ = fi8
        _ = v13
        _ = v9
    }
    v14 := GenS3{F0: (1e10 < 0.0), F1: -1}
    var v15 uint = 4294967295
    for fi16 := 0; fi16 < 20; fi16++ {
        if (!true) {
            v15 += (^(v15))
            if (1 < 179) {
                v17 := GenS4{F0: 255, F1: fi16}
                v3 = (v2 | (v2 * (v3 - 42)))
                v15 = v15
                _ = v17
            } else {
                var v18 string = "a"
                var v19 uint = (^(v15))
                _ = v18
                _ = v19
            }
            fmt.Println(v2, v3, v15)
        } else {
            fmt.Println(v2, v3, fi16)
            switch "hello" {
            case "hello":
                var v20 uint8 = 184
                _ = v20
            case "foo bar":
                if len(v4) > 0 {
                    v4[0] = (fi16 / (((fi16 | -561)) | 1))
                }
            case "":
                delete(v6, 0)
            }
            fmt.Println(v3)
        }
        _ = fi16
    }
    for fi21 := 0; fi21 < 3; fi21++ {
        if (!(true && (((1.5 - 0.1) + 3.14159) > 0.1))) {
            var v22 int = fi21
            _ = v22
        }
        _ = fi21
    }
    fmt.Println(v3, v15, v15)
    v23 := GenS4{F0: 361, F1: 839}
    v3 = (v3 * v3)
    v15 -= v15
    v15 -= v15
    v24 := GenS1{F0: 42}
    v4 = append(v4, 480)
    var v25 int = -21
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
