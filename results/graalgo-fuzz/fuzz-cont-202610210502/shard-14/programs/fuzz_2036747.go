package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 uint }

func main() {
    v1 := []string{"\n\t", ""}
    v2 := make(map[int]uint32)
    v3 := make(map[int]bool)
    var v4 int64 = 0
    var v5 float64 = 1e10
    var v6 uint8 = 128
    v7 := GenS1{F0: (v4 ^ (v4 << ((960) & 63)))}
    v8 := GenS2{F0: 93}
    var v9 float32 = 0.0
    fmt.Println(v6, v4, v4)
    _ = len(v1)
    var v10 uint = 65535
    v6 = (v6 << ((v10) & 63))
    for fi11 := 0; fi11 < 10; fi11++ {
        if false {
            var v12 float64 = -2.25
            fmt.Println(v6)
            if (-85 >= 0) {
                for fi13 := 0; fi13 < 7; fi13++ {
                    v6 = v6
                    var v14 int32 = 1
                    _ = fi13
                    _ = v14
                }
                fmt.Println(v10, v12)
                v12 += v5
            } else {
                for fi15 := 0; fi15 < 1; fi15++ {
                    v16 := GenS3{F0: (v10 << (((v10 | v10)) & 63)), F1: v10}
                    _ = fi15
                    _ = v16
                }
                if len(v1) > 0 {
                    v1[0] = "\n\t"
                }
            }
            _ = v12
        }
        _ = len(v1)
        for fi17 := 0; fi17 < 13; fi17++ {
            if false {
                v6 += (^(v6))
                switch "foo bar" {
                case "foo bar":
                    var v18 uint32 = 65535
                    _ = v18
                case "hello":
                    v19 := v3[848]
                    _ = v19
                default:
                    var v20 uint64 = 2
                    _ = v20
                }
                fmt.Println(v9, v10)
                fmt.Println(v9, v6, fi11)
            }
            v9 += v9
            v21 := v3[fi11]
            _ = fi17
            _ = v21
        }
        _ = fi11
    }
    switch 9223372036854775807 {
    case -9223372036854775808:
        var v22 uint = 0
        _ = v22
    default:
        var v23 uint = (v10 * (v10 / ((183) | 1)))
        _ = v23
    }
    v24 := GenS1{F0: v4}
    var v25 uint64 = 2
    var v26 float32 = v9
    v25 = v25
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
