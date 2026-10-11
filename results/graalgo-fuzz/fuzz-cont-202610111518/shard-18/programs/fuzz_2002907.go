package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }

func main() {
    var v1 int8 = -14
    var v2 uint64 = 42
    v3 := make(map[string]int)
    var v4 float64 = -2.25
    v5 := make(map[uint32]string)
    v6 := make(map[int]uint32)
    var v7 int = -748
    v2 = (v2 << ((2) & 63))
    fmt.Println(v2)
    var v8 float64 = float64((1.5 * 0.1))
    switch v7 {
    case -9223372036854775808:
        for fi9 := 0; fi9 < 1; fi9++ {
            fmt.Println(v1, fi9, v2)
            fmt.Println(fi9)
            if true {
                fmt.Println(v4)
                var v10 int = (fi9 * (fi9 - 127))
                _ = v10
            } else {
                fmt.Println(v4, v2, v7)
                if false {
                    v2 += 18446744073709551615
                    var v11 bool = true
                    if true {
                        for fi12 := 0; fi12 < 16; fi12++ {
                            v3["a"] = 128
                            _ = fi12
                        }
                    } else {
                        for fi13 := 0; fi13 < 2; fi13++ {
                            v2 = v2
                            v8 = 3.14159
                            switch "hello" {
                            case "":
                                v4 -= v4
                            case "hello":
                                var v14 uint16 = 32767
                                _ = v14
                            case "foo bar":
                                v15 := []int{(fi13 / ((42) | 1)), fi9}
                                _ = v15
                            }
                            _ = fi13
                        }
                    }
                    _ = v11
                } else {
                    v7 = v7
                    var v16 int32 = -1
                    _ = v16
                }
            }
            _ = fi9
        }
    case 4294967295:
        fmt.Println(v8, v8)
    }
    fmt.Println(v4, v7)
    v4 -= (v4 * v8)
    v7 = 106
    var v17 float32 = (3.14159 * -2.25)
    if true {
        v18 := v3["a"]
        for fi19 := 0; fi19 < 7; fi19++ {
            v20 := []int{v18}
            var v21 bool = (-294 <= 0)
            _ = fi19
            _ = v20
            _ = v21
        }
        _ = v18
    } else {
        v22 := GenS1{F0: 32767, F1: "foo bar"}
        _ = v22
    }
    var v23 uint32 = 373
    v7 = (v7 / ((v7) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
