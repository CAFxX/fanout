package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0() int16 {
    v5 := []int{-717, -804, -1, 168}
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 int8 = -128
        _ = fi6
        _ = v7
    }
    var v8 int64 = 42
    v8 = v8
    _ = v5
    _ = v8
    return 32767
}

func main() {
    v1 := []int{32767, 42, -9223372036854775808}
    v2 := make(map[string]string)
    var v3 uint8 = 255
    v4 := []int{-2147483648, 32768, 2147483647}
    var v9 uint64 = 653
    v4 = append(v4, -2147483648)
    fmt.Println(v3)
    var v10 float32 = ((-2.25 * -2.25) * 1.5)
    v9 += (^(v9))
    v10 = v10
    delete(v2, "hello")
    v3 -= 255
    var v11 float64 = (-2.25 + (1e10 + (1e10 - (0.1 - 1e10))))
    if false {
        var v12 bool = (!false)
        for fi13 := 0; fi13 < 10; fi13++ {
            switch -12 {
            case 32768:
                var v14 float64 = 3.14159
                _ = v14
            case 42:
                v3 = v3
            case 0:
                var v15 int = fi13
                _ = v15
            }
            _ = fi13
        }
        _ = v12
    }
    var v16 int32 = 0
    v17 := GenS1{F0: 0}
    switch "hello" {
    case "foo bar":
        v16 -= (v16 % (((v16 / (((v16 >> ((635) & 63))) | 1))) | 1))
    default:
        switch "foo bar" {
        case "":
            fmt.Println(v3, v11)
        case "a":
            fmt.Println(v11)
        case "hello":
            v16 = (v16 % (((v16 >> ((4294967295) & 63))) | 1))
        }
    }
    v11 -= v11
    switch 128 {
    case -1:
        if (2147483648 >= 807) {
            v11 = (((v11 + (float64((v10 + v10)) + 0.1)) + (v11 + 1.5)) - v11)
        }
    case 128:
        v10 = (0.1 + (((0.0 - 0.0) - float32(0.0)) - v10))
    default:
        for fi18 := 0; fi18 < 11; fi18++ {
            var v19 float32 = v10
            v20 := GenS1{F0: fi18}
            _ = fi18
            _ = v19
            _ = v20
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
