package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int16 }
func genFn0(p0 uint8, p1 int32) int {
    var v5 int32 = p1
    var v6 uint64 = 0
    fmt.Println(v6, v5, v6)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 944
}
func genFn1(p0 int32) string {
    var v7 uint8 = 132
    v8 := GenS1{F0: "", F1: 780, F2: 42}
    v9 := []int{}
    v7 = (v7 * v7)
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return ""
}

func main() {
    var v1 int8 = -128
    v2 := []int{-2147483648, -802, 2147483647, int(v1)}
    v3 := make(map[uint32]bool)
    var v4 uint16 = 65535
    for fi10 := 0; fi10 < 19; fi10++ {
        _ = len(v2)
        v11 := GenS1{F0: "hello", F1: 902, F2: -422}
        fmt.Println(v1, v4, fi10)
        _ = fi10
        _ = v11
    }
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    v12 := GenS1{F0: "foo bar", F1: 1, F2: -32768}
    v1 = v1
    for fi13 := 0; fi13 < 8; fi13++ {
        switch 75 {
        case 9223372036854775807:
            fmt.Println(fi13, v1)
        case -1:
            switch (^(fi13)) {
            case -9223372036854775808:
                if false {
                    var v14 uint64 = 4294967295
                    _ = v14
                } else {
                    var v15 int64 = -1
                    _ = v15
                }
            default:
                var v16 string = ""
                _ = v16
            }
        }
        var v17 int = 32767
        v4 += (v4 - (v4 % ((32767) | 1)))
        _ = fi13
        _ = v17
    }
    var v18 string = "hello"
    v4 -= v4
    var v19 float32 = 1e10
    if (!false) {
        switch v18 {
        case "":
            delete(v3, 16)
        case "a":
            if len(v2) > 0 {
                v2[0] = 32767
            }
        case "\n\t":
            v18 = (v18 + v18)
        }
        var v20 uint8 = 1
        _ = v20
    }
    var v21 string = (v18 + "foo bar")
    switch "\n\t" {
    case "\n\t":
        var v22 uint = 4294967295
        _ = v22
    case "a":
        var v23 int32 = 0
        _ = v23
    case "hello":
        fmt.Println(v4, v21, v4)
    }
    v19 = v19
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
