package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0() int {
    var v5 int = 587
    var v6 int = (^(v5))
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := GenS1{F0: 32767}
        _ = fi7
        _ = v8
    }
    var v9 uint = 255
    _ = v5
    _ = v6
    _ = v9
    return 0
}
func genFn1() string {
    v10 := GenS1{F0: 32767}
    var v11 uint32 = 2
    v12 := []int{128, 564}
    v11 = v11
    var v13 float32 = (((0.1 + (3.14159 + 3.14159)) / float32((1.5 - 3.14159))) * 3.14159)
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    return "\n\t"
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int = 28
    var v3 uint8 = 255
    v4 := make(map[string]bool)
    var v14 string = ""
    v14 = (((v14 + "\n\t") + v14) + "foo bar")
    if (!((1.5 * 0.0) > float64(((float32(1e10) * (0.1 / 1.5)) * 3.14159)))) {
        v15 := GenS1{F0: 816}
        _ = v15
    } else {
        v14 = v14
    }
    v16 := []int{}
    for fi17 := 0; fi17 < 16; fi17++ {
        v3 -= (v3 * 127)
        var v18 uint16 = 255
        v18 = v18
        _ = fi17
        _ = v18
    }
    var v19 int16 = 42
    var v20 float64 = ((-2.25 / ((((float64((-2.25 + 1e10)) / 1.5) - (1e10 + -2.25)) * (-2.25 / (0.1 - (float64(0.1) * (((-2.25 + (3.14159 - (((1e10 + ((float64(0.0) * 0.0) - 0.0)) / 0.1) - 1.5))) + (((float64(3.14159) - (((3.14159 + (0.0 + 0.0)) / 1.5) + (1.5 + (1e10 + float64(float32(1.5)))))) / -2.25) * (1.5 / (-2.25 - 3.14159)))) - (float64((-2.25 / 1.0)) + 0.1)))))) / 1.5)) - 0.0)
    var v21 uint16 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
