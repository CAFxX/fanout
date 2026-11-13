package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint16 }
type GenS2 struct { F0 int32; F1 int32; F2 uint8 }
type GenS3 struct { F0 int; F1 uint16 }

func main() {
    v1 := []string{"", "", "a"}
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]bool)
    var v4 string = "hello"
    v4 = "foo bar"
    var v5 uint16 = 65535
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v6 float64 = 1.5
    var v7 int = 32767
    v8 := GenS1{F0: 2147483647, F1: -2147483648, F2: (v5 << ((4294967295) & 63))}
    for fi9 := 0; fi9 < 10; fi9++ {
        v10 := GenS2{F0: 882, F1: -1, F2: 110}
        var v11 uint16 = 806
        if (v5 == (v5 << ((42) & 63))) {
            switch -9223372036854775808 {
            case -372:
                v12 := GenS3{F0: v7, F1: 255}
                _ = v12
            default:
                var v13 float32 = (1.5 - 0.0)
                _ = v13
            }
            var v14 int = (v7 - v7)
            _ = v14
        } else {
            for fi15 := 0; fi15 < 1; fi15++ {
                var v16 int64 = -1
                _ = fi15
                _ = v16
            }
            if false {
                v17 := v2[2]
                _ = v17
            } else {
                if false {
                    var v18 int = (fi9 - v7)
                    v7 += (fi9 * v7)
                    var v19 int8 = -1
                    _ = v18
                    _ = v19
                }
            }
            _ = len(v2)
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    if len(v1) > 0 {
        v1[0] = v4
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
