package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 uint32; F1 uint8; F2 bool }
type GenS3 struct { F0 uint; F1 uint32 }

func main() {
    var v1 int = 127
    var v2 int8 = 42
    v3 := make(map[string]string)
    var v4 int64 = 9223372036854775807
    fmt.Println(v1)
    var v5 int32 = -899
    v5 = (v5 | (v5 - 1))
    if (214 > 589) {
        for fi6 := 0; fi6 < 18; fi6++ {
            v7 := []int{(v1 >> ((1) & 63))}
            _ = fi6
            _ = v7
        }
        switch "a" {
        case "":
            var v8 uint64 = 2
            _ = v8
        default:
            switch "hello" {
            case "\n\t":
                for fi9 := 0; fi9 < 7; fi9++ {
                    v5 = (v5 ^ (^(v5)))
                    v10 := []int{(^(v1)), fi9, 32768, (v1 + (fi9 ^ 32767))}
                    switch "a" {
                    case "a":
                        for fi11 := 0; fi11 < 9; fi11++ {
                            v2 = (v2 << ((653) & 63))
                            _ = fi11
                        }
                    case "\n\t":
                        var v12 int16 = 235
                        _ = v12
                    }
                    _ = fi9
                    _ = v10
                }
            default:
                for fi13 := 0; fi13 < 12; fi13++ {
                    v14 := GenS1{F0: v4, F1: v4}
                    v5 -= (v5 % ((v5) | 1))
                    _ = fi13
                    _ = v14
                }
            }
        }
    } else {
        v15 := GenS2{F0: 2147483647, F1: 253, F2: (32767 < 32767)}
        v1 += v1
        _ = v15
    }
    fmt.Println(v2)
    var v16 bool = true
    v4 = v4
    v17 := v3["foo bar"]
    v18 := GenS2{F0: 2147483648, F1: 127, F2: true}
    fmt.Println(v4)
    v5 = -2147483648
    v19 := GenS2{F0: 4294967295, F1: 128, F2: false}
    switch v17 {
    case "a":
        v20 := []int{(v1 % (((-(v1))) | 1)), (^(v1))}
        _ = v20
    case "foo bar":
        v21 := []int{(v1 ^ v1), (-(v1)), 127}
        _ = v21
    }
    if v16 {
        var v22 bool = v16
        if v16 {
            fmt.Println(v4, v22, v2)
            switch "hello" {
            case "hello":
                v16 = v22
            case "":
                v2 = int8(v1)
            default:
                var v23 uint64 = 2
                _ = v23
            }
            fmt.Println(v2, v4)
        }
        delete(v3, (v17 + v17))
        _ = v22
    } else {
        var v24 int32 = v5
        _ = v24
    }
    v25 := GenS3{F0: 417, F1: 4294967295}
    v26 := GenS2{F0: 234, F1: 81, F2: (v4 <= 42)}
    for fi27 := 0; fi27 < 13; fi27++ {
        fmt.Println(fi27, v5)
        delete(v3, v17)
        delete(v3, v17)
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
