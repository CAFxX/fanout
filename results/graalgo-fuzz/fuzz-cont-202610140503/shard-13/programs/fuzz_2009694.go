package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 uint8; F2 uint64 }

func main() {
    var v1 int16 = 42
    v2 := make(map[string]string)
    var v3 uint16 = 1
    var v4 uint = 18446744073709551615
    var v5 string = "\n\t"
    fmt.Println(v4)
    v6 := v2[v5]
    v6 = (v5 + (v6 + (v6 + "\n\t")))
    v7 := GenS1{F0: false}
    for fi8 := 0; fi8 < 10; fi8++ {
        for fi9 := 0; fi9 < 7; fi9++ {
            v5 = v6
            delete(v2, v6)
            _ = fi9
        }
        _ = fi8
    }
    fmt.Println(v5, v3, v6)
    if false {
        v10 := GenS2{F0: (-(v1))}
        _ = v10
    } else {
        switch 967 {
        case -40:
            v11 := []int{}
            _ = v11
        case -492:
            v12 := GenS3{F0: 1, F1: 0, F2: 4294967295}
            _ = v12
        default:
            for fi13 := 0; fi13 < 8; fi13++ {
                switch fi13 {
                case -877:
                    _ = len(v2)
                case 0:
                    v4 -= (^(v4))
                case 235:
                    if false {
                        v2[v6] = v6
                        var v14 bool = false
                        for fi15 := 0; fi15 < 5; fi15++ {
                            fmt.Println(v6, fi13, v6)
                            if (-2.25 >= 0.1) {
                                v4 -= (v4 * v4)
                            }
                            v16 := []int{2147483647, fi13}
                            fmt.Println(v3)
                            _ = fi15
                            _ = v16
                        }
                        _ = v14
                    }
                }
                _ = fi13
            }
        }
        var v17 uint64 = 18446744073709551615
        v1 = 1
        _ = v17
    }
    var v18 uint8 = 128
    for fi19 := 0; fi19 < 2; fi19++ {
        var v20 int = fi19
        var v21 int32 = 1
        _ = fi19
        _ = v20
        _ = v21
    }
    v5 = "hello"
    v18 = (^(v18))
    var v22 bool = (-2.25 == 3.14159)
    var v23 uint8 = v18
    fmt.Println(v3, v4, v22)
    var v24 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
