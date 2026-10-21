package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint }
func genFn0(p0 uint32) uint16 {
    v5 := GenS1{F0: 458, F1: 0, F2: 734}
    var v6 uint = 65535
    fmt.Println(p0, p0)
    _ = p0
    _ = v5
    _ = v6
    return 65535
}

func main() {
    var v1 int = 4294967295
    var v2 uint64 = 0
    v3 := make(map[int]int)
    var v4 int8 = 1
    v3[v1] = v1
    v7 := GenS1{F0: v2, F1: v4, F2: 194}
    var v8 uint64 = 180
    v8 = v2
    switch (v1 & int(v4)) {
    case 128:
        v9 := GenS1{F0: (^(v2)), F1: int8(v2), F2: 0}
        _ = v9
    case 972:
        var v10 uint16 = uint16(v8)
        _ = v10
    }
    var v11 uint64 = 0
    for fi12 := 0; fi12 < 1; fi12++ {
        for fi13 := 0; fi13 < 2; fi13++ {
            v11 += 518
            _ = fi13
        }
        v14 := []int{(fi12 >> ((0) & 63)), fi12}
        _ = fi12
        _ = v14
    }
    fmt.Println(v11)
    if (0.0 >= (((((((3.14159 / (float64((float32((1e10 - 3.14159)) / (1e10 * ((((1.5 * ((0.1 - ((1e10 * (0.0 * 0.1)) + float32((1.5 / 1.0)))) - (3.14159 + (3.14159 + 1.5)))) + (0.1 + (((0.1 * ((-2.25 - float32(-2.25)) * -2.25)) + 0.1) * (0.1 + (1e10 - ((1e10 * (3.14159 - 0.0)) + (float32(1e10) * 0.0))))))) - 0.1) * 1e10)))) / 3.14159)) - (3.14159 + 1.5)) - 0.0) - 3.14159) + (0.1 - (float64(0.0) / 1.0))) - -2.25) * 1.5)) {
        for fi15 := 0; fi15 < 11; fi15++ {
            var v16 string = ""
            v16 = (("a" + ("foo bar" + "foo bar")) + (("" + v16) + v16))
            delete(v3, -1)
            _ = fi15
            _ = v16
        }
        if false {
            var v17 uint = 255
            v18 := []int{(v1 * (v1 & 9223372036854775807))}
            _ = v17
            _ = v18
        }
    }
    fmt.Println(v2, v11)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
