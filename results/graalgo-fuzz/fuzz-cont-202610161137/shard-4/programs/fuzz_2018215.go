package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 bool }
func genFn0(p0 int64) bool {
    fmt.Println(p0)
    var v4 float64 = 0.0
    _ = p0
    _ = v4
    return true
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]bool)
    var v3 bool = (42 <= 913)
    if ((1 != 581) || (57 > 0)) {
        v3 = (2147483648 < 42)
        v3 = false
        v3 = (!(!(!false)))
    } else {
        v5 := []int{127, 132, 32768, 1}
        _ = v5
    }
    fmt.Println(v3)
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := []int{(^(fi6)), fi6, 4294967295, (fi6 + fi6)}
        for fi8 := 0; fi8 < 15; fi8++ {
            var v9 float64 = float64(0.1)
            switch "a" {
            case "\n\t":
                if v3 {
                    v2[65535] = false
                    var v10 uint16 = 156
                    _ = v10
                } else {
                    v9 = 3.14159
                }
            case "hello":
                var v11 uint16 = 1
                _ = v11
            }
            var v12 uint16 = 1
            _ = fi8
            _ = v12
            _ = v9
        }
        if ((-2.25 * ((1e10 * 1.5) + (1e10 * (0.1 - (1.5 - ((3.14159 + -2.25) - 0.1)))))) != 3.14159) {
            v3 = (18446744073709551615 != 255)
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            v3 = v3
        } else {
            fmt.Println(fi6, fi6)
            switch 128 {
            case 341:
                v3 = v3
            case -983:
                v1[""] = v3
            default:
                var v13 int32 = int32(fi6)
                _ = v13
            }
        }
        _ = fi6
        _ = v7
    }
    v14 := GenS1{F0: 2, F1: -806, F2: (true && true)}
    v15 := GenS1{F0: 65535, F1: 2147483647, F2: true}
    v3 = (1 < -128)
    var v16 bool = false
    v17 := []int{140}
    switch -9223372036854775808 {
    case -303:
        for fi18 := 0; fi18 < 20; fi18++ {
            fmt.Println(fi18)
            _ = fi18
        }
    default:
        var v19 float64 = 1.5
        _ = v19
    }
    var v20 uint8 = 168
    if (1e10 < float64(((-2.25 + ((1e10 + float32(0.1)) * -2.25)) * 0.1))) {
        if (v16 || v16) {
            v20 = v20
        }
        var v21 string = "foo bar"
        _ = v21
    }
    v20 = (v20 / ((162) | 1))
    v22 := GenS1{F0: 32767, F1: -375, F2: v16}
    fmt.Println(v3, v20)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v22
    _ = v3
}
