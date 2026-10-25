package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
func genFn0(p0 string) uint {
    for fi5 := 0; fi5 < 3; fi5++ {
        var v6 float32 = (float32(0.0) / (float32(0.1) - (1.5 / ((1.5 / (((float32(1e10) * float32(-2.25)) + 0.0) - -2.25)) * 3.14159))))
        _ = fi5
        _ = v6
    }
    var v7 int = -778
    if false {
        var v8 int64 = 42
        switch v7 {
        case 0:
            var v9 int8 = 42
            _ = v9
        }
        _ = v8
    } else {
        for fi10 := 0; fi10 < 4; fi10++ {
            for fi11 := 0; fi11 < 16; fi11++ {
                v12 := GenS1{F0: 0, F1: 700}
                _ = fi11
                _ = v12
            }
            _ = fi10
        }
        p0 = p0
    }
    v13 := []int{v7, v7, v7, 1}
    _ = p0
    _ = v13
    _ = v7
    return 1
}

func main() {
    v1 := []int{}
    v2 := make(map[uint32]bool)
    v3 := make(map[string]string)
    var v4 uint8 = 1
    var v14 int16 = -796
    switch "a" {
    case "":
        var v15 int16 = (-(v14))
        _ = v15
    default:
        var v16 int16 = v14
        _ = v16
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        v4 = 1
        _ = fi17
    }
    v4 += v4
    var v18 string = "\n\t"
    v19 := GenS1{F0: 255, F1: 0}
    for fi20 := 0; fi20 < 16; fi20++ {
        for fi21 := 0; fi21 < 8; fi21++ {
            v14 = 32767
            _ = fi21
        }
        var v22 float32 = ((3.14159 - float32(0.1)) * 3.14159)
        _ = fi20
        _ = v22
    }
    var v23 int = 127
    var v24 uint64 = uint64(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
