package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int16; F1 uint64 }

func main() {
    var v1 uint32 = 65535
    var v2 uint64 = 665
    var v3 string = "hello"
    switch "a" {
    case "foo bar":
        fmt.Println(v1, v1, v3)
    }
    var v4 float32 = -2.25
    for fi5 := 0; fi5 < 11; fi5++ {
        fmt.Println(v2, v1)
        for fi6 := 0; fi6 < 1; fi6++ {
            var v7 uint = 42
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    v8 := GenS1{F0: 0}
    v9 := GenS2{F0: v3}
    v3 = "\n\t"
    var v10 float32 = v4
    for fi11 := 0; fi11 < 5; fi11++ {
        var v12 int = fi11
        switch "a" {
        case "foo bar":
            if (v4 <= (v10 - (v4 + float32(1e10)))) {
                v13 := []int{128, (fi11 + v12), fi11, (v12 * v12)}
                var v14 int32 = -35
                _ = v13
                _ = v14
            }
        case "hello":
            if false {
                var v15 string = ""
                v15 = "hello"
                switch v15 {
                case "":
                    var v16 int32 = 1
                    _ = v16
                case "\n\t":
                    v17 := GenS3{F0: 518, F1: v15}
                    _ = v17
                default:
                    v4 = v4
                }
                _ = v15
            }
        }
        _ = fi11
        _ = v12
    }
    if false {
        v18 := GenS4{F0: 127}
        _ = v18
    } else {
        v4 = v4
        for fi19 := 0; fi19 < 19; fi19++ {
            v4 = (float32(-2.25) - float32(3.14159))
            _ = fi19
        }
    }
    v20 := []int{-2147483648}
    var v21 uint32 = v1
    var v22 float32 = v10
    if false {
        _ = len(v20)
        for fi23 := 0; fi23 < 15; fi23++ {
            var v24 float32 = v22
            _ = fi23
            _ = v24
        }
        fmt.Println(v22, v21, v3)
    }
    var v25 int64 = 42
    v26 := GenS5{F0: 0, F1: (v2 >> ((461) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
