package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 string) uint {
    p0 = p0
    p0 = "\n\t"
    _ = p0
    return 42
}
func genFn1() int32 {
    var v7 bool = (false || false)
    if v7 {
        for fi8 := 0; fi8 < 20; fi8++ {
            v9 := GenS1{F0: genFn0("\n\t")}
            var v10 float32 = 1.5
            _ = fi8
            _ = v10
            _ = v9
        }
    } else {
        var v11 int16 = -268
        v11 = (-(v11))
        _ = v11
    }
    _ = v7
    return 1
}

func main() {
    var v1 int16 = -32768
    var v2 int16 = (v1 << ((740) & 63))
    var v3 int32 = 822
    var v4 uint32 = 231
    var v5 uint16 = 2
    var v6 int = 2147483647
    v12 := GenS1{F0: uint(v5)}
    v5 -= (v5 ^ (v5 | (^(v5))))
    v5 = v5
    fmt.Println(v3)
    v13 := GenS1{F0: 18446744073709551615}
    var v14 uint8 = 169
    v15 := []int{42}
    v3 = (v3 - (v3 | (v3 & (v3 << ((0) & 63)))))
    var v16 float64 = ((float64(((3.14159 / 1.5) + float32(-2.25))) + ((-2.25 + ((3.14159 * -2.25) * (3.14159 / 1e10))) / (0.1 + 1e10))) * (-2.25 - 0.1))
    var v17 bool = true
    var v18 bool = (!true)
    v15 = append(v15, (v6 >> ((2) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
