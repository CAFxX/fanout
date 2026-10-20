package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 int32; F1 int8 }
func genFn0() int32 {
    var v5 float64 = 3.14159
    v6 := GenS1{F0: 32767, F1: false}
    var v7 uint = 708
    _ = v5
    _ = v6
    _ = v7
    return 594
}

func main() {
    v1 := []string{}
    var v2 bool = false
    var v3 uint8 = 62
    v4 := make(map[string]int)
    if ((true && (1.5 < (float32(((1e10 - float64((3.14159 - float32(0.0)))) / 1.0)) * 0.1))) && (0 < -75)) {
        fmt.Println(v3)
        for fi8 := 0; fi8 < 11; fi8++ {
            var v9 uint8 = (v3 ^ (^(v3)))
            for fi10 := 0; fi10 < 9; fi10++ {
                if true {
                    _ = len(v4)
                    delete(v4, "foo bar")
                    var v11 float64 = 0.0
                    _ = v11
                } else {
                    fmt.Println(v2, fi8, fi10)
                    switch (("foo bar" + ("" + ("foo bar" + ""))) + ("" + ("a" + "foo bar"))) {
                    case "\n\t":
                        v12 := GenS1{F0: -357, F1: true}
                        _ = v12
                    default:
                        if v2 {
                            v13 := GenS1{F0: -1, F1: (-22 > 8)}
                            var v14 bool = true
                            _ = v13
                            _ = v14
                        } else {
                            fmt.Println(v3)
                            fmt.Println(fi10, v3)
                        }
                    }
                }
                var v15 bool = (fi8 > fi10)
                _ = fi10
                _ = v15
            }
            v1 = append(v1, ("" + "a"))
            _ = fi8
            _ = v9
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v16 int16 = -32768
    v17 := GenS2{F0: -1, F1: (v3 & (v3 * (v3 - 57)))}
    fmt.Println(v16, v3, v16)
    v2 = v2
    _ = len(v4)
    v18 := GenS3{F0: -1, F1: -128}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
