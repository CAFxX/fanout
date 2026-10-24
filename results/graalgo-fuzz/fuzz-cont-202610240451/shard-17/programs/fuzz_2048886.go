package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 uint64 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]uint32)
    var v3 float64 = (0.1 / (1e10 * ((float64(3.14159) * 1.5) / float64(1.5))))
    if ((false || (1 == 933)) && (65535 == 65535)) {
        switch -676 {
        case -9223372036854775808:
            var v4 uint = 0
            _ = v4
        case 9223372036854775807:
            v3 = v3
        case 42:
            v5 := v1[332]
            _ = v5
        default:
            v6 := GenS1{F0: 100, F1: 864}
            _ = v6
        }
    } else {
        fmt.Println(v3)
        v7 := GenS1{F0: 100, F1: 1}
        _ = v7
    }
    v8 := GenS2{F0: 32767}
    v3 = 0.1
    var v9 int = 32768
    fmt.Println(v9, v9)
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 uint8 = 181
        fmt.Println(v11, v3, v9)
        _ = fi10
        _ = v11
    }
    var v12 uint = 0
    switch "a" {
    case "foo bar":
        var v13 uint16 = 32767
        _ = v13
    case "":
        v14 := []int{v9, 32767, v9}
        _ = v14
    }
    var v15 bool = (("\n\t" + "") > "")
    var v16 uint16 = 1
    for fi17 := 0; fi17 < 9; fi17++ {
        var v18 int64 = 866
        var v19 uint8 = 2
        var v20 float32 = ((1.5 - 3.14159) + 1.5)
        _ = fi17
        _ = v18
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 20; fi21++ {
        switch "a" {
        case "":
            v22 := GenS3{F0: 42, F1: 852}
            _ = v22
        default:
            v23 := []int{(^(v9)), fi21, fi21}
            _ = v23
        }
        switch (fi21 >> (((v12 / ((v12) | 1))) & 63)) {
        case 32768:
            v24 := []int{v9, (v9 << ((548) & 63)), fi21}
            _ = v24
        }
        _ = fi21
    }
    fmt.Println(v9, v3)
    v25 := v1[v9]
    v26 := GenS1{F0: 42, F1: 201}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v8
    _ = v9
}
