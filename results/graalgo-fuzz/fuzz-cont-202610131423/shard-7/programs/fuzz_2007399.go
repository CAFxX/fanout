package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 int16; F2 int32 }
type GenS3 struct { F0 string }

func main() {
    v1 := []uint32{697, 787}
    var v2 string = "foo bar"
    var v3 uint32 = 65535
    v3 = (v3 / ((v3) | 1))
    v2 = (v2 + v2)
    v2 = v2
    for fi4 := 0; fi4 < 18; fi4++ {
        v5 := GenS1{F0: true}
        if true {
            for fi6 := 0; fi6 < 18; fi6++ {
                var v7 int16 = 1
                if true {
                    v8 := GenS1{F0: false}
                    _ = v8
                }
                _ = fi6
                _ = v7
            }
            var v9 int8 = -128
            if false {
                if len(v1) > 0 {
                    v1[0] = 65535
                }
            } else {
                v2 = ((v2 + "foo bar") + (v2 + (v2 + (v2 + v2))))
                if false {
                    v10 := GenS2{F0: -865, F1: 42, F2: 42}
                    for fi11 := 0; fi11 < 7; fi11++ {
                        var v12 float64 = 0.0
                        var v13 int32 = -193
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                    _ = v10
                } else {
                    v14 := GenS1{F0: true}
                    var v15 uint64 = 404
                    v1 = append(v1, v3)
                    _ = v14
                    _ = v15
                }
            }
            _ = v9
        } else {
            var v16 int64 = -9223372036854775808
            v2 = "a"
            fmt.Println(fi4)
            _ = v16
        }
        switch (^(fi4)) {
        case 32768:
            var v17 int32 = -318
            _ = v17
        case -1:
            v18 := GenS1{F0: (((v2 + (v2 + v2)) + (v2 + v2)) == v2)}
            _ = v18
        default:
            v3 += (v3 - (v3 * 461))
        }
        _ = fi4
        _ = v5
    }
    v2 = "\n\t"
    v19 := GenS3{F0: v2}
    var v20 float64 = 3.14159
    var v21 string = v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
