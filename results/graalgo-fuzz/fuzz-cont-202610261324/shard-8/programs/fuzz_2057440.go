package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 int32; F2 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32; F1 uint }
func genFn0() bool {
    if false {
        if false {
            var v5 uint16 = 32767
            _ = v5
        } else {
            v6 := []int{-407, 0}
            _ = v6
        }
        var v7 uint32 = 2
        _ = v7
    } else {
        var v8 uint8 = 128
        _ = v8
    }
    for fi9 := 0; fi9 < 19; fi9++ {
        switch 32768 {
        case 9223372036854775807:
            var v10 int = (fi9 | (fi9 * 2147483647))
            _ = v10
        }
        _ = fi9
    }
    return (4294967295 <= 2)
}

func main() {
    var v1 int16 = -338
    var v2 uint64 = 4294967295
    var v3 uint8 = 255
    var v4 string = "foo bar"
    v11 := GenS1{F0: 4294967295}
    if false {
        fmt.Println(v3, v3)
    }
    var v12 float32 = 0.0
    if true {
        v3 += v3
        fmt.Println(v2)
    }
    var v13 uint64 = v2
    v1 -= v1
    v14 := GenS1{F0: 1}
    if true {
        var v15 string = v4
        var v16 int32 = 0
        _ = v15
        _ = v16
    }
    v17 := GenS2{F0: 202, F1: -2147483648, F2: ""}
    for fi18 := 0; fi18 < 15; fi18++ {
        v19 := []int{}
        if (!false) {
            if true {
                if true {
                    fmt.Println(v13)
                } else {
                    var v20 uint64 = v13
                    v21 := GenS3{F0: -128}
                    if len(v19) >= 2 {
                        v19 = v19[1:len(v19)]
                    }
                    _ = v20
                    _ = v21
                }
                if (true || (false || ((false || false) || (59 <= 1)))) {
                    v13 = v13
                    var v22 uint32 = 2147483647
                    var v23 int8 = 1
                    _ = v22
                    _ = v23
                }
                v24 := GenS4{F0: 2, F1: 587}
                _ = v24
            }
            v2 = 2
        }
        _ = fi18
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
