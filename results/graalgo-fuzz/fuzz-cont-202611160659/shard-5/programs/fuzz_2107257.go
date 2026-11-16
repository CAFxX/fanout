package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 int16 }

func main() {
    v1 := []int{32767, 2147483647, -860, 32767}
    v2 := []string{"a", "foo bar", "", ""}
    v3 := make(map[string]bool)
    var v4 bool = true
    v5 := []int{1, 9223372036854775807, -468, 2147483647}
    var v6 int16 = 528
    v7 := GenS1{F0: 1000, F1: 18446744073709551615, F2: 42}
    var v8 int = -2147483648
    if len(v5) > 0 {
        v5[0] = (v8 + (^(v8)))
    }
    for fi9 := 0; fi9 < 16; fi9++ {
        if ((!(0.1 >= ((-2.25 / (0.1 + ((1.5 - 1.5) * 1e10))) * ((0.1 * 1.5) + -2.25)))) && false) {
            v10 := GenS1{F0: v6, F1: 0, F2: (^(v6))}
            _ = v10
        }
        v11 := GenS1{F0: v6, F1: 4294967295, F2: (^(v6))}
        _ = fi9
        _ = v11
    }
    v6 -= (v6 << ((1) & 63))
    for fi12 := 0; fi12 < 19; fi12++ {
        v6 = v6
        fmt.Println(v8, fi12)
        _ = fi12
    }
    switch 127 {
    case -257:
        switch 9223372036854775807 {
        case -2147483648:
            var v13 float64 = float64((0.1 * 0.1))
            _ = v13
        default:
            switch v8 {
            case -1:
                var v14 float32 = ((((1e10 / (float32(((3.14159 + 0.1) * (3.14159 / (3.14159 + 0.0)))) / float32(-2.25))) + 0.1) - 0.1) * 0.0)
                _ = v14
            default:
                if ((42 > 2147483647) || ((2 == 2) && false)) {
                    var v15 int32 = 565
                    var v16 float32 = (((1.5 / 1.0) + 1.5) - 0.0)
                    v16 += v16
                    _ = v15
                    _ = v16
                }
            }
        }
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        var v18 float64 = 0.0
        for fi19 := 0; fi19 < 16; fi19++ {
            var v20 string = "foo bar"
            v4 = (int8(v6) < 42)
            v21 := GenS1{F0: v6, F1: 119, F2: 408}
            _ = fi19
            _ = v20
            _ = v21
        }
        var v22 uint8 = 2
        _ = fi17
        _ = v18
        _ = v22
    }
    v6 = v6
    fmt.Println(v6)
    for fi23 := 0; fi23 < 2; fi23++ {
        for fi24 := 0; fi24 < 9; fi24++ {
            v25 := GenS1{F0: (^(v6)), F1: 255, F2: v6}
            v26 := GenS1{F0: -454, F1: 1, F2: (v6 << ((568) & 63))}
            v8 = -9223372036854775808
            var v27 uint8 = 127
            _ = fi24
            _ = v25
            _ = v26
            _ = v27
        }
        fmt.Println(v6, fi23)
        _ = fi23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
