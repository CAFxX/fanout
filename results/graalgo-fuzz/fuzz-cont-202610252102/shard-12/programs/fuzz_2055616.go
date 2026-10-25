package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64; F1 uint16 }
type GenS5 struct { F0 bool; F1 uint32; F2 uint8 }
func genFn0() int {
    var v5 string = "a"
    fmt.Println(v5)
    var v6 int64 = -9223372036854775808
    var v7 uint64 = 4294967295
    _ = v5
    _ = v6
    _ = v7
    return 32767
}

func main() {
    v1 := []int{32767, 9223372036854775807}
    v2 := make(map[uint32]bool)
    v3 := make(map[string]int)
    v4 := make(map[int]string)
    var v8 uint8 = 2
    if true {
        fmt.Println(v8)
        var v9 float32 = (1e10 / 1.5)
        _ = v9
    } else {
        var v10 uint16 = 547
        _ = v10
    }
    fmt.Println(v8)
    v1 = append(v1, 32768)
    if false {
        _ = len(v1)
        v11 := GenS1{F0: 128}
        _ = v11
    } else {
        v12 := GenS2{F0: 111}
        v8 += 128
        v13 := GenS3{F0: 42}
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        v15 := GenS3{F0: 1}
        _ = len(v4)
        _ = fi14
        _ = v15
    }
    var v16 string = "foo bar"
    for fi17 := 0; fi17 < 17; fi17++ {
        v18 := GenS4{F0: 492, F1: 65535}
        _ = fi17
        _ = v18
    }
    var v19 uint8 = (v8 & (v8 & (v8 & v8)))
    v20 := GenS2{F0: -1}
    v8 -= (v8 * (^(v8)))
    if true {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        switch "\n\t" {
        case "hello":
            v21 := GenS5{F0: false, F1: 0, F2: v19}
            _ = v21
        case "\n\t":
            fmt.Println(v8, v8)
        }
        var v22 float32 = 1.5
        _ = v22
    } else {
        for fi23 := 0; fi23 < 9; fi23++ {
            var v24 uint64 = 4294967295
            if (741 >= (v24 - v24)) {
                var v25 uint = 177
                v25 -= (v25 & v25)
                _ = v25
            }
            v8 = uint8(v24)
            _ = fi23
            _ = v24
        }
        v8 = 0
        v8 += (v8 % (((v8 ^ v19)) | 1))
    }
    for fi26 := 0; fi26 < 17; fi26++ {
        var v27 uint = uint(v19)
        _ = fi26
        _ = v27
    }
    var v28 int64 = -220
    var v29 uint = 131
    var v30 uint64 = 42
    v31 := GenS5{F0: (-2.25 > -2.25), F1: uint32(v30), F2: v19}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v8
}
