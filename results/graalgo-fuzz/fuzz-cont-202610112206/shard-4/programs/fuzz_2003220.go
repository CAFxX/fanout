package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint }
type GenS2 struct { F0 uint32; F1 int32; F2 uint16 }
func genFn0(p0 uint64) uint16 {
    var v4 string = ("" + "foo bar")
    for fi5 := 0; fi5 < 6; fi5++ {
        if true {
            for fi6 := 0; fi6 < 13; fi6++ {
                var v7 uint16 = 42
                if ((1e10 * 0.1) != 1e10) {
                    var v8 float64 = (float64(3.14159) / 3.14159)
                    v9 := GenS1{F0: v4, F1: 32767, F2: 262}
                    v10 := []int{fi5, 127}
                    var v11 uint = 42
                    _ = v10
                    _ = v11
                    _ = v8
                    _ = v9
                } else {
                    var v12 float32 = 0.0
                    fmt.Println(p0, v12, fi5)
                    _ = v12
                }
                _ = fi6
                _ = v7
            }
        }
        _ = fi5
    }
    _ = p0
    _ = v4
    return 313
}

func main() {
    var v1 uint = 1
    var v2 uint8 = 128
    var v3 int8 = 126
    fmt.Println(v2, v3, v2)
    if (("a" + "a") >= "hello") {
        var v13 int8 = (-(v3))
        _ = v13
    }
    var v14 string = "hello"
    var v15 uint32 = 42
    v16 := []int{4294967295}
    var v17 uint = 1
    var v18 uint16 = 255
    switch -9223372036854775808 {
    case 128:
        v19 := GenS2{F0: (v15 ^ v15), F1: 0, F2: v18}
        _ = v19
    default:
        v16 = append(v16, 128)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
