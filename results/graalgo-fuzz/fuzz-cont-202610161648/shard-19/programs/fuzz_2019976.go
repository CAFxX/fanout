package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 int }
func genFn0(p0 int32, p1 uint64, p2 int32) uint32 {
    p0 = 661
    p0 = p2
    var v4 int64 = 9223372036854775807
    v5 := GenS1{F0: v4}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 2147483648
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := []int{9223372036854775807}
    var v3 int8 = -128
    if false {
        for fi6 := 0; fi6 < 12; fi6++ {
            v3 = (v3 >> ((v1) & 63))
            v3 = (^(v3))
            if false {
                for fi7 := 0; fi7 < 3; fi7++ {
                    fmt.Println(v1, fi6, v1)
                    v1 = (v1 / ((v1) | 1))
                    _ = fi7
                }
                var v8 int32 = 0
                _ = v8
            }
            _ = fi6
        }
        switch 32768 {
        case 174:
            switch "hello" {
            case "foo bar":
                for fi9 := 0; fi9 < 7; fi9++ {
                    var v10 int64 = 42
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v11 int8 = 78
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            case "hello":
                v12 := GenS2{F0: 2147483647}
                _ = v12
            case "":
                var v13 float64 = -2.25
                _ = v13
            default:
                var v14 float32 = -2.25
                _ = v14
            }
        }
        v3 += v3
    }
    for fi15 := 0; fi15 < 6; fi15++ {
        v16 := GenS3{F0: (v3 | v3), F1: 2147483647}
        fmt.Println(v1)
        v17 := GenS2{F0: fi15}
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 bool = (42 <= -1)
    for fi19 := 0; fi19 < 14; fi19++ {
        var v20 uint8 = 116
        if (!true) {
            fmt.Println(v1, v20)
            v3 += v3
        }
        _ = fi19
        _ = v20
    }
    v2 = append(v2, 128)
    for fi21 := 0; fi21 < 4; fi21++ {
        if len(v2) > 0 {
            v2[0] = fi21
        }
        _ = fi21
    }
    fmt.Println(v3, v18, v18)
    for fi22 := 0; fi22 < 4; fi22++ {
        v23 := GenS1{F0: -9223372036854775808}
        var v24 string = "a"
        _ = fi22
        _ = v23
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
}
