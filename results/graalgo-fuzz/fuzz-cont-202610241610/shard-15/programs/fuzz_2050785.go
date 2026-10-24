package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint }
type GenS2 struct { F0 uint64; F1 int64; F2 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16; F1 uint8; F2 string }
func genFn0(p0 uint16, p1 int16) int64 {
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 int32 = 1
        v8 = v8
        _ = fi7
        _ = v8
    }
    p0 = 32767
    v9 := GenS1{F0: -17, F1: 598, F2: 18446744073709551615}
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = v9
    return 1
}
func genFn1(p0 int16, p1 int16, p2 uint) int {
    switch "a" {
    case "\n\t":
        for fi10 := 0; fi10 < 18; fi10++ {
            for fi11 := 0; fi11 < 12; fi11++ {
                var v12 uint64 = 18446744073709551615
                if (!true) {
                    var v13 uint64 = v12
                    _ = v13
                }
                if true {
                    v14 := []int{}
                    p0 += (p1 << ((p2) & 63))
                    _ = v14
                } else {
                    if false {
                        v15 := GenS1{F0: -1, F1: 0, F2: 225}
                        _ = v15
                    } else {
                        var v16 int = 620
                        fmt.Println(v16, p1, p0)
                        fmt.Println(v16, p0)
                        _ = v16
                    }
                }
                _ = fi11
                _ = v12
            }
            v17 := GenS2{F0: 0, F1: -9223372036854775808, F2: 791}
            v18 := GenS3{F0: 2147483647}
            _ = fi10
            _ = v17
            _ = v18
        }
    default:
        var v19 bool = (-477 == 0)
        _ = v19
    }
    p2 = p2
    p1 += (p0 + (p0 << ((p2) & 63)))
    _ = p0
    _ = p1
    _ = p2
    return 32768
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]uint32)
    var v3 uint = 18446744073709551615
    var v4 int32 = 2147483647
    v5 := make(map[int]bool)
    v6 := []int{-2147483648, 127}
    fmt.Println(v3, v4)
    var v20 uint32 = 42
    var v21 string = "\n\t"
    v6 = append(v6, 50)
    _ = len(v5)
    v22 := GenS3{F0: uint32(v3)}
    switch (v21 + "\n\t") {
    case "hello":
        for fi23 := 0; fi23 < 16; fi23++ {
            for fi24 := 0; fi24 < 20; fi24++ {
                var v25 int64 = -65
                fmt.Println(v4)
                _ = fi24
                _ = v25
            }
            var v26 uint8 = 127
            _ = fi23
            _ = v26
        }
    }
    var v27 int32 = v4
    for fi28 := 0; fi28 < 12; fi28++ {
        v29 := GenS4{F0: 0, F1: 1, F2: v21}
        _ = fi28
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
