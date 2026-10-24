package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int; F1 int; F2 int8 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]uint32)
    v4 := GenS1{F0: 0}
    if (0.0 <= 0.1) {
        v5 := GenS2{F0: 335}
        var v6 uint = 4294967295
        _ = v5
        _ = v6
    } else {
        var v7 int32 = -2147483648
        fmt.Println(v7)
        v8 := GenS2{F0: 1}
        _ = v7
        _ = v8
    }
    if false {
        for fi9 := 0; fi9 < 20; fi9++ {
            var v10 uint = 42
            _ = fi9
            _ = v10
        }
    } else {
        if (false || true) {
            v11 := []int{}
            if true {
                if (9223372036854775807 > -2147483648) {
                    var v12 int32 = 1
                    _ = v12
                }
                var v13 string = ""
                _ = v13
            } else {
                var v14 string = ("" + "a")
                var v15 int = 1
                v15 = (v15 << ((904) & 63))
                _ = v14
                _ = v15
            }
            v16 := GenS3{F0: 497}
            _ = v11
            _ = v16
        }
        var v17 uint8 = 2
        if ("foo bar" == "a") {
            fmt.Println(v17)
        } else {
            v18 := GenS3{F0: 0}
            v17 += (v17 + 0)
            _ = v18
        }
        _ = v17
    }
    switch 9223372036854775807 {
    case 127:
        var v19 uint = 321
        _ = v19
    case 4294967295:
        var v20 uint64 = 927
        _ = v20
    }
    v21 := GenS2{F0: 1}
    switch "\n\t" {
    case "foo bar":
        var v22 string = ""
        _ = v22
    case "":
        v23 := GenS4{F0: 2147483647}
        _ = v23
    }
    for fi24 := 0; fi24 < 1; fi24++ {
        fmt.Println(fi24)
        v25 := GenS5{F0: true}
        _ = fi24
        _ = v25
    }
    v26 := GenS6{F0: 9223372036854775807, F1: 32768, F2: -1}
    switch "a" {
    case "a":
        if (0 == -215) {
            var v27 uint32 = 2147483648
            _ = v27
        } else {
            var v28 uint32 = 2147483647
            _ = v28
        }
    default:
        var v29 int = -9223372036854775808
        _ = v29
    }
    v30 := GenS5{F0: ((0 != 9223372036854775807) || false)}
    var v31 int8 = 42
    v32 := GenS3{F0: 4294967295}
    v33 := []int{-2147483648, -2147483648}
    for fi34 := 0; fi34 < 7; fi34++ {
        v35 := GenS4{F0: 42}
        _ = fi34
        _ = v35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v33:", len(v33), v33)
    _ = v1
    _ = v2
    _ = v21
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
