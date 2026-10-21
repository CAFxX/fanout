package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int16 }
func genFn0(p0 uint32) bool {
    var v4 bool = false
    var v5 int16 = 0
    _ = p0
    _ = v4
    _ = v5
    return v4
}
func genFn1(p0 int, p1 int16, p2 string) int {
    for fi6 := 0; fi6 < 1; fi6++ {
        v7 := []int{32768, fi6, fi6}
        _ = fi6
        _ = v7
    }
    p1 = 1000
    _ = p0
    _ = p1
    _ = p2
    return p0
}

func main() {
    v1 := make(map[int]int)
    var v2 float64 = 3.14159
    v3 := []uint32{2147483647, 2}
    var v8 bool = (42 == 2147483647)
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 uint32 = 2
        var v11 int64 = 15
        switch "" {
        case "a":
            for fi12 := 0; fi12 < 11; fi12++ {
                var v13 float64 = (((0.0 + float64(1e10)) / v2) - float64((1.5 - 1.5)))
                var v14 uint64 = 42
                switch "hello" {
                case "foo bar":
                    switch fi9 {
                    case 32767:
                        if v8 {
                            if ((v14 > (v14 ^ 18446744073709551615)) && v8) {
                                v15 := GenS1{F0: 42, F1: -1, F2: -32768}
                                var v16 int8 = 35
                                _ = v15
                                _ = v16
                            } else {
                                var v17 uint8 = 31
                                _ = v17
                            }
                            _ = len(v3)
                            var v18 string = "a"
                            _ = v18
                        }
                    default:
                        var v19 uint32 = (v10 & 1)
                        _ = v19
                    }
                }
                _ = fi12
                _ = v13
                _ = v14
            }
        default:
            v11 -= (v11 & int64(v10))
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    v8 = (!v8)
    var v20 uint64 = 986
    _ = len(v3)
    fmt.Println(v2)
    var v21 int16 = -360
    _ = len(v3)
    v22 := GenS1{F0: 2, F1: -24, F2: 1}
    fmt.Println(v8, v8)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v8
}
