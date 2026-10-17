package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }

func main() {
    v1 := []int{}
    v2 := make(map[uint32]uint32)
    v3 := []uint32{}
    var v4 uint16 = 552
    var v5 uint32 = 1
    var v6 int8 = 100
    var v7 int64 = 1
    fmt.Println(v7, v6)
    if (42 <= -1) {
        switch 127 {
        case 127:
            switch "hello" {
            case "foo bar":
                if true {
                    fmt.Println(v4, v6, v6)
                }
            case "hello":
                var v8 int16 = 1
                _ = v8
            case "a":
                v9 := GenS1{F0: 2147483647, F1: -650}
                _ = v9
            default:
                v10 := GenS1{F0: -1, F1: -18}
                _ = v10
            }
        case -751:
            if false {
                var v11 int16 = 357
                var v12 uint32 = v5
                v13 := GenS1{F0: 2147483647, F1: 128}
                _ = v11
                _ = v12
                _ = v13
            } else {
                v14 := GenS1{F0: 1, F1: 448}
                var v15 bool = true
                _ = v14
                _ = v15
            }
        case -766:
            if true {
                var v16 int = 755
                _ = v16
            } else {
                switch -1 {
                case -9223372036854775808:
                    fmt.Println(v4, v6)
                case 300:
                    for fi17 := 0; fi17 < 19; fi17++ {
                        fmt.Println(v7)
                        fmt.Println(v5, fi17)
                        v5 += 1
                        _ = fi17
                    }
                case 271:
                    for fi18 := 0; fi18 < 3; fi18++ {
                        var v19 float32 = 0.1
                        v2[v5] = v5
                        _ = fi18
                        _ = v19
                    }
                default:
                    v4 = 2
                }
                switch 32768 {
                case -1:
                    v2[v5] = (v5 % ((v5) | 1))
                case -154:
                    if ((true && false) && true) {
                        var v20 int = 869
                        delete(v2, v5)
                        v7 += -9223372036854775808
                        _ = v20
                    } else {
                        if false {
                            var v21 int32 = -816
                            _ = v21
                        }
                        fmt.Println(v5, v5)
                        v6 = (v6 % ((v6) | 1))
                    }
                case -112:
                    _ = len(v2)
                }
                var v22 uint8 = 1
                _ = v22
            }
        }
    } else {
        _ = len(v2)
        var v23 int = -9223372036854775808
        _ = v23
    }
    v24 := GenS1{F0: 214, F1: 32768}
    _ = len(v3)
    var v25 uint32 = v5
    v26 := GenS1{F0: 0, F1: -448}
    _ = len(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
