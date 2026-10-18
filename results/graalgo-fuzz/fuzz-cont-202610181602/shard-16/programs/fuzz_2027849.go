package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 uint32 }

func main() {
    var v1 int32 = 0
    v2 := make(map[int]string)
    var v3 uint32 = 711
    v4 := make(map[string]bool)
    v1 = (v1 + (v1 ^ v1))
    var v5 int = 952
    if false {
        v5 = v5
    } else {
        var v6 int16 = 1
        _ = v6
    }
    v7 := GenS1{F0: 1, F1: 105, F2: 46}
    v5 = (v5 / ((v5) | 1))
    fmt.Println(v1, v5)
    v1 -= (v1 % ((v1) | 1))
    if (-1 != -9223372036854775808) {
        v8 := []int{v5, v5, 32767}
        if true {
            for fi9 := 0; fi9 < 3; fi9++ {
                switch -51 {
                case 388:
                    if ((-2.25 * 3.14159) != 1e10) {
                        v5 += fi9
                        switch -1 {
                        case 9223372036854775807:
                            for fi10 := 0; fi10 < 11; fi10++ {
                                v5 -= v5
                                v5 = v5
                                _ = fi10
                            }
                        case 924:
                            fmt.Println(v3, v5)
                        case 128:
                            v11 := v4[("\n\t" + ("hello" + "hello"))]
                            _ = v11
                        }
                        _ = len(v4)
                    } else {
                        fmt.Println(v3, v1, v1)
                        v12 := GenS2{F0: 2147483647}
                        fmt.Println(v5, v5, v5)
                        _ = v12
                    }
                case -9223372036854775808:
                    for fi13 := 0; fi13 < 12; fi13++ {
                        v1 += v1
                        _ = fi13
                    }
                }
                v1 = v1
                var v14 bool = false
                _ = fi9
                _ = v14
            }
            var v15 uint = 18446744073709551615
            _ = v15
        }
        _ = v8
    } else {
        var v16 float32 = 1.5
        for fi17 := 0; fi17 < 6; fi17++ {
            v18 := GenS3{F0: 42, F1: (v3 >> ((255) & 63))}
            for fi19 := 0; fi19 < 4; fi19++ {
                v20 := []int{fi17, 930, fi19, fi17}
                _ = fi19
                _ = v20
            }
            _ = fi17
            _ = v18
        }
        _ = v16
    }
    fmt.Println(v3, v1, v1)
    v2[(v5 ^ 9223372036854775807)] = "\n\t"
    v21 := []int{v5, v5, v5}
    fmt.Println(v3, v1, v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
