package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
func genFn0() uint32 {
    var v6 uint64 = 186
    v6 = (v6 | 0)
    _ = v6
    return 65535
}

func main() {
    var v1 bool = false
    var v2 uint = 255
    v3 := make(map[string]uint32)
    var v4 int32 = 958
    var v5 float64 = 1e10
    fmt.Println(v4, v5)
    v7 := []int{32767, 2147483647, -9223372036854775808}
    if (128 <= -9223372036854775808) {
        var v8 string = "a"
        _ = v8
    }
    var v9 int32 = (v4 | (v4 << ((v2) & 63)))
    v7 = append(v7, 2147483647)
    switch -472 {
    case 32768:
        v10 := GenS1{F0: "a", F1: -49}
        _ = v10
    case 758:
        var v11 float64 = 1.5
        _ = v11
    case 12:
        v12 := v3[("hello" + "")]
        _ = v12
    default:
        for fi13 := 0; fi13 < 12; fi13++ {
            fmt.Println(v1, v2)
            var v14 int32 = 179
            fmt.Println(v14)
            _ = fi13
            _ = v14
        }
    }
    if (1 == 0) {
        var v15 float32 = float32(v5)
        _ = v15
    }
    var v16 int32 = 2147483647
    if (3.14159 != float32(v5)) {
        v2 = (v2 ^ (^(v2)))
        for fi17 := 0; fi17 < 15; fi17++ {
            if v1 {
                for fi18 := 0; fi18 < 15; fi18++ {
                    if len(v7) > 0 {
                        v7[0] = (fi18 / ((2147483647) | 1))
                    }
                    switch "a" {
                    case "":
                        v5 = v5
                    }
                    _ = fi18
                }
            }
            _ = fi17
        }
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        v16 = (v4 | int32(fi19))
        if (!v1) {
            v4 += v4
            var v20 uint64 = 42
            var v21 int32 = 165
            _ = v20
            _ = v21
        } else {
            var v22 float32 = (3.14159 + 1e10)
            _ = v22
        }
        _ = fi19
    }
    _ = len(v3)
    switch 32768 {
    case 2147483647:
        fmt.Println(v9, v16)
    case -9223372036854775808:
        var v23 bool = (0 >= 42)
        _ = v23
    case 4294967295:
        fmt.Println(v16, v2)
    }
    var v24 uint32 = uint32(v2)
    for fi25 := 0; fi25 < 20; fi25++ {
        var v26 float32 = 3.14159
        var v27 uint = v2
        _ = fi25
        _ = v26
        _ = v27
    }
    var v28 uint16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v24
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v9
}
