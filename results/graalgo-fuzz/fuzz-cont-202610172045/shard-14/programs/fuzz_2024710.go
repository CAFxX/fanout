package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 string }
func genFn0(p0 int32) uint {
    p0 = (^(p0))
    fmt.Println(p0)
    var v6 string = "\n\t"
    v6 = v6
    _ = p0
    _ = v6
    return 42
}

func main() {
    var v1 int = 128
    v2 := make(map[string]int)
    var v3 uint8 = 216
    var v4 int64 = 0
    var v5 float64 = 3.14159
    var v7 int = (-(v1))
    v8 := GenS1{F0: 0}
    v9 := GenS2{F0: (!false), F1: "\n\t"}
    for fi10 := 0; fi10 < 2; fi10++ {
        fmt.Println(fi10, v4)
        _ = fi10
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        var v12 int32 = -1
        var v13 string = "\n\t"
        _ = fi11
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 20; fi14++ {
        var v15 float64 = (1e10 * v5)
        if (v15 == (1.5 * float64(1e10))) {
            switch fi14 {
            case 32768:
                v16 := GenS1{F0: 73}
                _ = v16
            case 9223372036854775807:
                for fi17 := 0; fi17 < 3; fi17++ {
                    var v18 int8 = -88
                    v18 += v18
                    _ = fi17
                    _ = v18
                }
            default:
                var v19 int64 = v4
                _ = v19
            }
            var v20 int16 = 1
            v21 := GenS2{F0: false, F1: "hello"}
            _ = v20
            _ = v21
        }
        _ = fi14
        _ = v15
    }
    v5 += v5
    _ = len(v2)
    v22 := []int{v1, (v1 >> ((2) & 63)), 1, 0}
    var v23 uint8 = 1
    v2["\n\t"] = -1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
