package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 int16; F2 int32 }

func main() {
    var v1 int32 = 919
    v2 := make(map[int]int)
    var v3 int8 = 127
    for fi4 := 0; fi4 < 10; fi4++ {
        v1 += -491
        fmt.Println(fi4)
        for fi5 := 0; fi5 < 8; fi5++ {
            if false {
                v6 := GenS1{F0: (fi4 ^ (fi5 & 2147483647)), F1: 1}
                _ = v6
            }
            switch "hello" {
            case "a":
                v7 := GenS2{F0: (fi5 & 128)}
                _ = v7
            case "foo bar":
                if true {
                    v3 -= v3
                    var v8 int64 = 423
                    fmt.Println(fi5)
                    var v9 int64 = v8
                    var v10 uint8 = 204
                    switch "hello" {
                    case "foo bar":
                        if false {
                            fmt.Println(v9, v1)
                        }
                    case "\n\t":
                        v10 = (v10 & (v10 ^ 82))
                        switch fi5 {
                        case 879:
                            v1 = -903
                        case 38:
                            var v11 bool = true
                            _ = v11
                        case -9223372036854775808:
                            var v12 int32 = (^(v1))
                            v8 += v8
                            _ = v12
                        }
                    }
                    _ = v10
                    _ = v8
                    _ = v9
                } else {
                    v13 := GenS3{F0: 255}
                    v14 := GenS4{F0: 763, F1: -32768, F2: v1}
                    _ = v13
                    _ = v14
                }
            }
            _ = fi5
        }
        _ = fi4
    }
    var v15 float32 = -2.25
    v15 -= 1e10
    fmt.Println(v1)
    if false {
        var v16 uint = 0
        v16 -= (v16 % ((18446744073709551615) | 1))
        _ = v16
    }
    var v17 uint16 = 974
    var v18 uint64 = 18446744073709551615
    fmt.Println(v1, v3)
    v17 = v17
    var v19 uint64 = v18
    fmt.Println(v15)
    var v20 int16 = 0
    v18 -= 4294967295
    var v21 float32 = (0.0 / v15)
    delete(v2, 56)
    for fi22 := 0; fi22 < 15; fi22++ {
        var v23 string = ("" + "")
        _ = fi22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
