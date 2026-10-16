package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool }

func main() {
    v1 := []uint32{2147483648, 291, 430, 2147483647}
    var v2 int = -9223372036854775808
    v3 := make(map[int]bool)
    v4 := make(map[string]string)
    var v5 int32 = -1
    var v6 uint32 = 2147483648
    v7 := GenS1{F0: (^(v5)), F1: v6}
    v5 -= v5
    if (!true) {
        if false {
            v2 -= v2
            if false {
                for fi8 := 0; fi8 < 9; fi8++ {
                    v5 = v5
                    fmt.Println(fi8, v2)
                    _ = fi8
                }
                v9 := GenS2{F0: (v5 | (^(v5)))}
                _ = len(v3)
                _ = v9
            } else {
                for fi10 := 0; fi10 < 19; fi10++ {
                    fmt.Println(v2, fi10, fi10)
                    if false {
                        for fi11 := 0; fi11 < 17; fi11++ {
                            v2 -= fi11
                            var v12 float32 = 3.14159
                            _ = fi11
                            _ = v12
                        }
                        v2 = -81
                        v2 -= v2
                    } else {
                        for fi13 := 0; fi13 < 5; fi13++ {
                            var v14 int = (fi10 | fi10)
                            _ = fi13
                            _ = v14
                        }
                        var v15 uint = 18446744073709551615
                        _ = v15
                    }
                    _ = fi10
                }
            }
            var v16 string = ("hello" + "a")
            _ = v16
        }
    } else {
        fmt.Println(v2, v6, v2)
    }
    v17 := GenS3{F0: false}
    switch v2 {
    case -855:
        if len(v1) > 0 {
            v1[0] = (v6 / ((v6) | 1))
        }
    }
    v6 = 2
    fmt.Println(v2)
    for fi18 := 0; fi18 < 3; fi18++ {
        v19 := v4["a"]
        v4["foo bar"] = ("" + (v19 + v19))
        _ = fi18
        _ = v19
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
