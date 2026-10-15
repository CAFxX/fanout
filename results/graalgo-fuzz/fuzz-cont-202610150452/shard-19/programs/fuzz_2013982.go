package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 uint8; F1 uint32; F2 int8 }
type GenS4 struct { F0 int32; F1 string }

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = (127 > 116)
    v3 := make(map[string]int)
    var v4 int8 = 127
    v5 := []int{32768, -1}
    v6 := GenS1{F0: v2, F1: 0}
    v7 := GenS1{F0: false, F1: -1}
    for fi8 := 0; fi8 < 2; fi8++ {
        if v2 {
            var v9 float64 = 0.0
            _ = v9
        } else {
            if v2 {
                v10 := v3[""]
                _ = v10
            }
            delete(v3, ((("" + "hello") + "hello") + ("\n\t" + ((("\n\t" + "foo bar") + "hello") + ("foo bar" + "")))))
        }
        v11 := GenS2{F0: 2147483647, F1: 1000}
        v4 = 0
        _ = fi8
        _ = v11
    }
    var v12 bool = ((v4 + (^(v4))) < 100)
    if v12 {
        switch 827 {
        case 121:
            v13 := GenS2{F0: -858, F1: 0}
            _ = v13
        case 1:
            v4 = (v4 / ((127) | 1))
        case -217:
            v5 = append(v5, 1)
        default:
            var v14 float32 = float32(0.0)
            _ = v14
        }
        switch 32767 {
        case 42:
            v4 += v4
        }
    } else {
        v5 = append(v5, -2147483648)
        v4 = v4
        for fi15 := 0; fi15 < 10; fi15++ {
            v16 := v3["hello"]
            v4 = v4
            var v17 uint32 = 2147483647
            _ = fi15
            _ = v16
            _ = v17
        }
    }
    v4 -= (v4 * v4)
    v5 = append(v5, 9223372036854775807)
    v18 := GenS3{F0: 58, F1: 2147483648, F2: v4}
    for fi19 := 0; fi19 < 2; fi19++ {
        if false {
            v20 := GenS2{F0: -902, F1: -32768}
            _ = v20
        } else {
            v12 = (v12 || (v12 || ((1.5 + 1e10) >= 0.1)))
            v21 := GenS2{F0: 42, F1: -32768}
            for fi22 := 0; fi22 < 3; fi22++ {
                v23 := GenS4{F0: 1, F1: "hello"}
                var v24 uint = 276
                v12 = v12
                _ = fi22
                _ = v23
                _ = v24
            }
            _ = v21
        }
        _ = fi19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
