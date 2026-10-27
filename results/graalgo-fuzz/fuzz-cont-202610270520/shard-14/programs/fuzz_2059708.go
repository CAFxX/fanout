package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 uint64 }
func genFn0(p0 uint32) int16 {
    v6 := GenS1{F0: 2, F1: 62}
    switch -1 {
    case 127:
        var v7 uint = 42
        _ = v7
    default:
        if (1 == 2147483647) {
            fmt.Println(p0)
            for fi8 := 0; fi8 < 14; fi8++ {
                for fi9 := 0; fi9 < 5; fi9++ {
                    v10 := GenS2{F0: p0, F1: "hello"}
                    var v11 int8 = 42
                    v11 = 0
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                v12 := GenS2{F0: (p0 & (p0 & (p0 + (^(p0))))), F1: "a"}
                _ = fi8
                _ = v12
            }
        } else {
            for fi13 := 0; fi13 < 19; fi13++ {
                var v14 bool = false
                for fi15 := 0; fi15 < 5; fi15++ {
                    if ((("a" + "") < "") && v14) {
                        var v16 uint = 0
                        if (2 > 128) {
                            v14 = v14
                            var v17 uint32 = 2147483648
                            _ = v17
                        } else {
                            p0 = (p0 + p0)
                            fmt.Println(v14)
                            fmt.Println(v16)
                        }
                        _ = v16
                    } else {
                        var v18 int8 = 0
                        _ = v18
                    }
                    p0 -= 4294967295
                    _ = fi15
                }
                p0 = p0
                _ = fi13
                _ = v14
            }
            v19 := GenS1{F0: 65535, F1: 216}
            v20 := GenS2{F0: (^(p0)), F1: "a"}
            _ = v19
            _ = v20
        }
    }
    v21 := GenS3{F0: 686}
    p0 -= (p0 + p0)
    _ = p0
    _ = v21
    _ = v6
    return 32767
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[int]int)
    var v3 uint8 = 255
    v4 := make(map[uint32]string)
    var v5 uint16 = 42
    for fi22 := 0; fi22 < 7; fi22++ {
        v5 -= v5
        v23 := []int{}
        v24 := GenS2{F0: 796, F1: "hello"}
        _ = fi22
        _ = v23
        _ = v24
    }
    var v25 int8 = 79
    fmt.Println(v5)
    v26 := GenS2{F0: 817, F1: v1}
    fmt.Println(v5, v25, v25)
    v3 += (v3 >> ((0) & 63))
    v3 += 255
    var v27 float64 = (0.1 / 3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
