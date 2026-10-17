package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 uint16; F2 uint32 }
type GenS3 struct { F0 uint64; F1 int; F2 uint16 }
func genFn0(p0 uint16, p1 int, p2 int) uint32 {
    switch "a" {
    case "a":
        for fi6 := 0; fi6 < 7; fi6++ {
            var v7 uint32 = 2147483648
            p2 = (p1 >> ((559) & 63))
            var v8 uint = 4294967295
            _ = fi6
            _ = v7
            _ = v8
        }
    }
    var v9 uint32 = 0
    fmt.Println(v9, v9)
    v10 := GenS1{F0: 591}
    for fi11 := 0; fi11 < 4; fi11++ {
        v12 := GenS2{F0: -9223372036854775808, F1: p0, F2: (v9 >> ((1) & 63))}
        fmt.Println(p1, p0)
        v13 := []int{2147483647, (p1 & p2), (p2 + 32768), (p2 | p1)}
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return (v9 % ((v9) | 1))
}
func genFn1(p0 int8, p1 int32) uint8 {
    fmt.Println(p1, p0)
    v14 := []int{-2147483648, 2147483647, 32768}
    var v15 float32 = 1.5
    v15 = v15
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    return 127
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint8 = 255
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]uint32)
    v5 := []string{}
    if (42 != 9223372036854775807) {
        for fi16 := 0; fi16 < 11; fi16++ {
            for fi17 := 0; fi17 < 2; fi17++ {
                v1 = (v1 * (^(v1)))
                delete(v3, (fi16 % ((fi17) | 1)))
                if false {
                    var v18 bool = (v1 <= v1)
                    v19 := GenS2{F0: -433, F1: 729, F2: genFn0(414, (fi17 / (((-(fi16))) | 1)), 32768)}
                    v18 = (v18 && v18)
                    _ = v18
                    _ = v19
                } else {
                    _ = len(v4)
                    var v20 int16 = -302
                    var v21 int16 = -505
                    _ = v20
                    _ = v21
                }
                _ = fi17
            }
            switch ("a" + "a") {
            case "\n\t":
                fmt.Println(v2, v1, v2)
            case "foo bar":
                var v22 string = "a"
                _ = v22
            }
            switch 32768 {
            case 127:
                fmt.Println(v2)
            }
            _ = fi16
        }
        v23 := GenS3{F0: 310, F1: 1, F2: 1}
        _ = v23
    }
    fmt.Println(v1)
    for fi24 := 0; fi24 < 11; fi24++ {
        v25 := GenS1{F0: 4294967295}
        fmt.Println(v1)
        var v26 int32 = 1
        _ = fi24
        _ = v25
        _ = v26
    }
    v1 = (v1 % ((v1) | 1))
    var v27 string = "\n\t"
    fmt.Println(v2, v1)
    for fi28 := 0; fi28 < 18; fi28++ {
        if ((1.5 + 0.1) < ((0.1 / float64((float32(3.14159) + 3.14159))) / float64(1e10))) {
            fmt.Println(v2)
        }
        var v29 uint = v1
        _ = fi28
        _ = v29
    }
    _ = len(v5)
    var v30 int8 = -4
    var v31 int64 = 0
    v3[32767] = 1
    v32 := GenS1{F0: 970}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
