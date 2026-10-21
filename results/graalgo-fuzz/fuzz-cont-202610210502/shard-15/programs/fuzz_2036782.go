package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint8 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint32) int8 {
    var v7 int = 128
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 uint32 = (^(p0))
        v10 := []int{(-(v7)), v7, fi8, (-(fi8))}
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: v7, F1: -1, F2: 1}
    p0 += p0
    p0 = p0
    _ = p0
    _ = v11
    _ = v7
    return 127
}

func main() {
    var v1 int = 9223372036854775807
    var v2 float64 = 1.5
    var v3 int16 = -458
    var v4 int64 = -738
    v5 := make(map[uint32]string)
    v6 := []uint32{2147483647, 42, 651, 0}
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 int16 = v3
        _ = fi12
        _ = v13
    }
    var v14 uint16 = 65535
    if true {
        var v15 float32 = 3.14159
        fmt.Println(v14)
        for fi16 := 0; fi16 < 12; fi16++ {
            for fi17 := 0; fi17 < 7; fi17++ {
                switch "hello" {
                case "a":
                    var v18 uint8 = 106
                    _ = v18
                case "hello":
                    switch ("hello" + "\n\t") {
                    case "hello":
                        v6 = append(v6, 0)
                    case "foo bar":
                        if len(v6) > 0 {
                            v6[0] = 417
                        }
                    default:
                        v4 += (-(v4))
                    }
                }
                var v19 int8 = -1
                _ = fi17
                _ = v19
            }
            var v20 uint16 = (v14 - (^(v14)))
            var v21 float64 = v2
            _ = fi16
            _ = v20
            _ = v21
        }
        _ = v15
    } else {
        v1 = v1
    }
    var v22 int64 = (v4 & v4)
    _ = len(v6)
    fmt.Println(v1, v2, v2)
    v22 -= v4
    var v23 uint32 = 1
    _ = len(v6)
    v3 = -32768
    fmt.Println(v22, v14)
    var v24 int32 = -717
    v25 := GenS1{F0: (v1 ^ (v1 >> ((589) & 63))), F1: 9223372036854775807, F2: 0}
    v23 += v23
    v26 := GenS2{F0: -2147483648, F1: v3}
    for fi27 := 0; fi27 < 10; fi27++ {
        v28 := GenS3{F0: (v24 + -990)}
        v1 -= fi27
        _ = fi27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
