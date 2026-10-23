package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint32 }
func genFn0() uint {
    switch 2147483647 {
    case 1:
        if false {
            if false {
                v6 := []int{128, 1, 128}
                _ = v6
            } else {
                var v7 uint32 = 4294967295
                _ = v7
            }
        }
    default:
        var v8 bool = true
        _ = v8
    }
    switch 2147483647 {
    case 127:
        var v9 float32 = 0.1
        _ = v9
    case 2147483647:
        if false {
            var v10 int32 = 42
            _ = v10
        }
    }
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int32 = 1
    var v3 string = ((("a" + "\n\t") + "") + ("foo bar" + "a"))
    v4 := []uint32{}
    var v5 uint64 = 18446744073709551615
    v3 = ("\n\t" + v3)
    var v11 uint64 = v5
    var v12 int16 = -1
    v13 := GenS1{F0: v3, F1: 662, F2: 1}
    _ = len(v1)
    if len(v4) > 0 {
        v4[0] = 2
    }
    var v14 uint8 = 11
    var v15 uint16 = 32767
    _ = len(v4)
    if false {
        switch ("" + "hello") {
        case "hello":
            var v16 string = v3
            _ = v16
        case "\n\t":
            if false {
                for fi17 := 0; fi17 < 15; fi17++ {
                    v1[""] = 0
                    switch -2147483648 {
                    case -9223372036854775808:
                        v2 = v2
                    default:
                        v12 -= v12
                    }
                    var v18 int8 = int8(v12)
                    _ = fi17
                    _ = v18
                }
                v4 = append(v4, 2147483647)
                var v19 uint32 = 2147483648
                _ = v19
            } else {
                v5 += uint64(v15)
            }
        }
    }
    var v20 float64 = 3.14159
    var v21 uint = genFn0()
    v11 = (v5 | (v11 % ((v11) | 1)))
    v21 -= v21
    v12 -= v12
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
