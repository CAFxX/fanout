package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 uint16) int32 {
    p0 += p0
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    return 678
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float64 = (float64((((1.5 * 0.1) - ((1e10 / 1.0) / 1.0)) + (-2.25 - -2.25))) * ((0.0 * float64(((-2.25 - (1e10 + (0.1 / ((3.14159 * float32((0.1 + 0.0))) + (-2.25 / 1.5))))) + (0.0 * float32((float64(float32((((1e10 + ((1.5 + 1e10) + -2.25)) * 1e10) + 1.5))) + float64(3.14159))))))) + 0.0))
    var v3 uint8 = 127
    var v4 string = "\n\t"
    var v5 int = 4294967295
    v6 := make(map[string]string)
    v7 := GenS1{F0: -816}
    v2 += float64(1e10)
    v5 = (v5 - -290)
    var v8 uint8 = (v3 - 255)
    v3 = (v8 >> ((0) & 63))
    fmt.Println(v3)
    var v9 string = v4
    if false {
        if false {
            v2 += v2
            var v10 int32 = 913
            _ = v10
        }
        v5 -= (v5 & v5)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
