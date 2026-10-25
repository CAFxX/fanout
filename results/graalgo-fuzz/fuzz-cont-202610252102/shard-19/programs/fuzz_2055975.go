package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 int32 = 42
    var v2 float32 = float32((1e10 * 0.0))
    var v3 int32 = 42
    v4 := make(map[int]string)
    var v5 uint32 = 1
    v6 := GenS1{F0: -128}
    fmt.Println(v2)
    v5 -= 764
    var v7 uint8 = 127
    v3 = (v3 % ((v3) | 1))
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 float32 = 3.14159
        fmt.Println(v3, v2)
        _ = fi8
        _ = v9
    }
    v7 = 255
    v7 -= v7
    var v10 uint32 = (v5 ^ 2147483648)
    v11 := []int{806, -169, 0}
    v12 := GenS2{F0: -890}
    v13 := GenS1{F0: -1}
    v2 -= float32(-2.25)
    for fi14 := 0; fi14 < 7; fi14++ {
        v15 := GenS1{F0: -94}
        v3 = (v3 & v3)
        _ = fi14
        _ = v15
    }
    var v16 uint = 0
    fmt.Println(v10, v3)
    if true {
        var v17 float32 = v2
        switch "foo bar" {
        case "\n\t":
            for fi18 := 0; fi18 < 15; fi18++ {
                var v19 uint64 = 158
                v3 = (^(v3))
                for fi20 := 0; fi20 < 18; fi20++ {
                    switch "hello" {
                    case "hello":
                        if (!true) {
                            v21 := GenS2{F0: 0}
                            _ = v21
                        } else {
                            v1 += v1
                            if (!(!((3.14159 + ((1e10 - v2) + 0.1)) < float32(float64(-2.25))))) {
                                var v22 bool = (1 < -1)
                                var v23 bool = v22
                                _ = v22
                                _ = v23
                            }
                        }
                        v24 := v4[127]
                        _ = v24
                    case "foo bar":
                        v7 = v7
                    }
                    v11 = append(v11, (fi20 << ((uint(v19)) & 63)))
                    _ = fi20
                }
                _ = fi18
                _ = v19
            }
        case "foo bar":
            _ = len(v4)
        }
        v5 = v10
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
