package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    v1 := []uint32{0, 1}
    var v2 int8 = 52
    var v3 uint16 = 65535
    var v4 float64 = (((0.1 / float64(3.14159)) + 0.1) - 0.0)
    var v5 uint32 = 165
    var v6 bool = false
    _ = len(v1)
    var v7 bool = (v6 && v6)
    var v8 float64 = (v4 - v4)
    fmt.Println(v7, v7, v7)
    var v9 uint16 = (v3 & v3)
    fmt.Println(v8, v7, v3)
    switch 1 {
    case 127:
        if false {
            if true {
                fmt.Println(v5, v7, v8)
                var v10 uint32 = (v5 % (((v5 - (v5 >> ((18446744073709551615) & 63)))) | 1))
                _ = v10
            } else {
                for fi11 := 0; fi11 < 3; fi11++ {
                    var v12 int64 = 42
                    _ = fi11
                    _ = v12
                }
                for fi13 := 0; fi13 < 5; fi13++ {
                    v5 -= (^(v5))
                    var v14 uint8 = 1
                    _ = fi13
                    _ = v14
                }
                if (65535 < uint(v5)) {
                    var v15 int64 = 1
                    var v16 float64 = 3.14159
                    if false {
                        _ = len(v1)
                    } else {
                        var v17 string = ""
                        var v18 float32 = ((1.5 + ((1.5 * (float32(v16) * float32(v8))) / (0.1 + ((0.1 + 3.14159) / (3.14159 + 1.5))))) / -2.25)
                        var v19 uint64 = 42
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                    _ = v15
                    _ = v16
                }
            }
        } else {
            v20 := GenS1{F0: "foo bar"}
            v1 = append(v1, v5)
            _ = v20
        }
    case -1:
        v21 := GenS1{F0: "\n\t"}
        _ = v21
    default:
        if v6 {
            v7 = true
        }
    }
    v22 := GenS1{F0: "a"}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
