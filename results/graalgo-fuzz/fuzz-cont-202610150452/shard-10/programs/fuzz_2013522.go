package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 int }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint8 = 127
    v3 := []int{42, 409, -1, -1}
    var v4 string = (("a" + "foo bar") + "\n\t")
    var v5 int32 = -2147483648
    v2 += v2
    v4 = "a"
    var v6 float64 = ((1e10 * (((((0.1 - -2.25) + (1.5 + 3.14159)) - (1e10 * 1e10)) / 1.0) * ((0.1 - (((((float64(float32((1.5 - (float64(3.14159) + (1e10 + 3.14159))))) + 0.0) / (3.14159 + 0.1)) * (-2.25 + -2.25)) / 0.1) - ((0.1 * 0.1) + -2.25))) - (float64(0.0) + 1e10)))) / 0.1)
    var v7 string = (v4 + "")
    var v8 int16 = 83
    for fi9 := 0; fi9 < 18; fi9++ {
        for fi10 := 0; fi10 < 15; fi10++ {
            v5 = v5
            if true {
                if (true && ((((v6 / 3.14159) / ((v6 - v6) + v6)) < -2.25) && false)) {
                    if (int(v8) <= fi10) {
                        if len(v3) > 0 {
                            v3[0] = 9223372036854775807
                        }
                    }
                    v11 := GenS1{F0: 42, F1: -23, F2: fi10}
                    _ = v11
                }
                v2 = (v2 & (v2 % (((v2 & (v2 << ((4294967295) & 63)))) | 1)))
                if (951 == 0) {
                    var v12 int16 = v8
                    v8 += v12
                    var v13 uint8 = v2
                    _ = v12
                    _ = v13
                }
                var v14 uint = 18446744073709551615
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = len(v3)
                _ = v14
            }
            _ = fi10
        }
        var v15 float32 = 3.14159
        var v16 string = v7
        _ = fi9
        _ = v15
        _ = v16
    }
    v17 := GenS1{F0: 4294967295, F1: 100, F2: 32767}
    for fi18 := 0; fi18 < 2; fi18++ {
        v7 = (("\n\t" + ((v7 + v7) + v7)) + "")
        var v19 int8 = 1
        _ = fi18
        _ = v19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
