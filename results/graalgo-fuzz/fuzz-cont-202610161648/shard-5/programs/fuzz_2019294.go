package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 int64) int16 {
    fmt.Println(p0)
    v6 := GenS1{F0: 18446744073709551615}
    _ = p0
    _ = v6
    return 1
}
func genFn1(p0 uint16) uint64 {
    for fi7 := 0; fi7 < 19; fi7++ {
        v8 := GenS1{F0: 0}
        _ = fi7
        _ = v8
    }
    fmt.Println(p0)
    _ = p0
    return 0
}

func main() {
    v1 := []uint32{}
    var v2 float64 = (-2.25 + (((3.14159 + (((((-2.25 / (1e10 / (((0.1 * 1.5) * float64(((0.0 + ((1.5 + (0.1 * (-2.25 + (-2.25 + 0.1)))) - 1.5)) + 3.14159))) + (float64(0.0) / (float64(1.5) / 1.0))))) / 1.5) + (0.1 - -2.25)) + ((1e10 + ((((float64(0.1) + (1e10 / (0.0 + 1.5))) * 1e10) - 0.0) / 0.1)) / (float64(0.1) + ((1e10 / float64((1e10 * 1.5))) - (1e10 - (1.5 / 3.14159)))))) + (((3.14159 * 0.0) / 3.14159) / 1.0))) + -2.25) / (0.1 / ((3.14159 * (0.0 + float64(3.14159))) * 1e10))))
    v3 := []int{124, 9223372036854775807}
    var v4 int64 = 0
    var v5 uint8 = 50
    for fi9 := 0; fi9 < 6; fi9++ {
        v5 = v5
        v4 += (v4 & -9223372036854775808)
        if (713 < 255) {
            var v10 uint8 = 98
            if (65535 > 42) {
                for fi11 := 0; fi11 < 20; fi11++ {
                    var v12 uint8 = (v10 + (^(v10)))
                    var v13 int8 = 127
                    if len(v1) > 0 {
                        v1[0] = 2147483647
                    }
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                v4 += int64(fi9)
            } else {
                v2 -= (((v2 + v2) + (0.1 - -2.25)) - 1.5)
            }
            fmt.Println(v4)
            _ = v10
        }
        _ = fi9
    }
    v2 = 0.0
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v14 float64 = float64((3.14159 * float32(v2)))
    v5 -= 0
    var v15 int = 894
    fmt.Println(v14, v5, v5)
    v15 -= (v15 >> ((255) & 63))
    if len(v3) > 0 {
        v3[0] = -9223372036854775808
    }
    var v16 uint8 = (^(v5))
    if (1 <= 211) {
        if false {
            v16 = 32
            var v17 uint16 = 32767
            _ = v17
        }
        for fi18 := 0; fi18 < 3; fi18++ {
            var v19 bool = ("\n\t" > "foo bar")
            _ = fi18
            _ = v19
        }
    }
    v16 = (v16 << ((785) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
