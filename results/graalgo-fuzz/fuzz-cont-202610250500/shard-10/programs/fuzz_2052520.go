package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 int32; F2 uint32 }

func main() {
    var v1 uint16 = 118
    v2 := make(map[int]uint32)
    v3 := make(map[int]int)
    var v4 int = 2147483647
    var v5 int8 = 30
    v6 := GenS1{F0: (!false)}
    v7 := GenS1{F0: true}
    v8 := []int{(v4 / ((819) | 1)), 360}
    switch "hello" {
    case "foo bar":
        if (1 >= 1) {
            for fi9 := 0; fi9 < 10; fi9++ {
                for fi10 := 0; fi10 < 6; fi10++ {
                    var v11 int8 = (v5 / (((v5 | int8(v1))) | 1))
                    v1 = 42
                    _ = fi10
                    _ = v11
                }
                _ = fi9
            }
        } else {
            var v12 int = (v4 - v4)
            var v13 uint16 = (v1 | v1)
            switch "hello" {
            case "\n\t":
                var v14 uint32 = 2147483647
                _ = v14
            case "a":
                v15 := GenS1{F0: true}
                _ = v15
            }
            _ = v12
            _ = v13
        }
    }
    v4 = v4
    var v16 string = ""
    v4 = (v4 + (v4 << ((755) & 63)))
    if (1 >= 0) {
        v17 := GenS2{F0: -209}
        fmt.Println(v1)
        v18 := GenS1{F0: true}
        _ = v17
        _ = v18
    } else {
        for fi19 := 0; fi19 < 20; fi19++ {
            var v20 string = v16
            var v21 uint16 = v1
            _ = fi19
            _ = v20
            _ = v21
        }
        v1 = v1
    }
    _ = len(v8)
    for fi22 := 0; fi22 < 20; fi22++ {
        v23 := GenS1{F0: false}
        _ = fi22
        _ = v23
    }
    v4 = v4
    if (-9223372036854775808 != 1) {
        v24 := GenS3{F0: 18446744073709551615, F1: 42, F2: 664}
        _ = v24
    } else {
        var v25 uint32 = 8
        var v26 uint = 2
        _ = v25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
