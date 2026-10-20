package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 int64 }

func main() {
    v1 := make(map[int]int)
    var v2 uint16 = 1
    v3 := make(map[string]string)
    var v4 int = -2147483648
    v2 += (v2 / ((32767) | 1))
    v5 := []int{v4, v4}
    v2 = (v2 & (v2 % ((v2) | 1)))
    v2 -= v2
    fmt.Println(v4, v4)
    var v6 int16 = 42
    fmt.Println(v4, v4, v6)
    v6 = (v6 + (v6 ^ v6))
    _ = len(v1)
    var v7 string = "a"
    for fi8 := 0; fi8 < 9; fi8++ {
        switch (v7 + (v7 + "hello")) {
        case "foo bar":
            v9 := GenS1{F0: (127 >= 0)}
            _ = v9
        case "a":
            v4 -= v4
        case "hello":
            for fi10 := 0; fi10 < 10; fi10++ {
                var v11 uint = 725
                switch fi10 {
                case 4294967295:
                    v12 := GenS2{F0: "\n\t", F1: -933}
                    _ = v12
                case 740:
                    switch v4 {
                    case 0:
                        v11 = 253
                    case 32767:
                        v1[128] = v4
                    default:
                        switch v7 {
                        case "\n\t":
                            var v13 int = v4
                            _ = v13
                        }
                    }
                    var v14 bool = (0.0 == 0.1)
                    _ = len(v5)
                    _ = v14
                case 2147483647:
                    fmt.Println(v6, v4, v4)
                default:
                    var v15 uint64 = 4294967295
                    _ = v15
                }
                v16 := v3["a"]
                _ = fi10
                _ = v11
                _ = v16
            }
        }
        var v17 int = 781
        _ = fi8
        _ = v17
    }
    var v18 int64 = 0
    if (100 != -60) {
        var v19 float32 = 1.5
        if false {
            for fi20 := 0; fi20 < 19; fi20++ {
                var v21 int32 = -2147483648
                v22 := v1[int(v2)]
                _ = fi20
                _ = v21
                _ = v22
            }
            delete(v1, v4)
            v19 = v19
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
