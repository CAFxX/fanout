package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string; F1 int32; F2 uint16 }

func main() {
    var v1 string = "a"
    var v2 uint = 18446744073709551615
    var v3 int8 = -31
    var v4 int = 32767
    v1 = v1
    v5 := GenS1{F0: true}
    var v6 uint8 = 128
    var v7 uint8 = (v6 / (((^(v6))) | 1))
    var v8 int8 = (^(v3))
    switch ("foo bar" + "") {
    case "a":
        switch v1 {
        case "":
            var v9 int64 = 42
            _ = v9
        default:
            fmt.Println(v7, v4, v4)
        }
    case "\n\t":
        fmt.Println(v2, v4, v1)
    case "":
        v10 := GenS2{F0: (v4 >> (((v2 / (((v2 >> ((v2) & 63))) | 1))) & 63))}
        _ = v10
    }
    var v11 uint8 = v6
    switch (v1 + v1) {
    case "foo bar":
        fmt.Println(v7, v1, v3)
    case "hello":
        var v12 uint16 = 1
        _ = v12
    case "a":
        var v13 bool = true
        _ = v13
    default:
        v2 -= 0
    }
    fmt.Println(v1, v4, v11)
    var v14 uint8 = (v11 + v6)
    if (24 >= 625) {
        v15 := GenS3{F0: false}
        var v16 uint64 = 1
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        fmt.Println(v2, v3)
        var v18 int8 = 0
        v19 := []int{(^(v4)), -463, v4, v4}
        _ = fi17
        _ = v18
        _ = v19
    }
    v20 := GenS3{F0: false}
    switch (v1 + (v1 + v1)) {
    case "a":
        var v21 uint32 = 2
        _ = v21
    case "":
        v22 := GenS4{F0: "hello", F1: 883, F2: 0}
        _ = v22
    case "foo bar":
        v7 -= v14
    }
    v23 := []int{v4, v4}
    fmt.Println(v7, v2, v1)
    for fi24 := 0; fi24 < 13; fi24++ {
        for fi25 := 0; fi25 < 3; fi25++ {
            fmt.Println(v7)
            _ = fi25
        }
        if (("" + "") == v1) {
            var v26 uint = 42
            if len(v23) >= 2 {
                v23 = v23[1:len(v23)]
            }
            _ = v26
        }
        var v27 int16 = -32768
        _ = fi24
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
