package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 int32) uint8 {
    v6 := GenS1{F0: 0}
    fmt.Println(p0)
    v7 := GenS1{F0: 1}
    _ = p0
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 float64 = 0.1
    var v2 uint32 = 65535
    v3 := []int{-1, 32768, 745, -198}
    var v4 int16 = -994
    v5 := make(map[string]uint32)
    for fi8 := 0; fi8 < 3; fi8++ {
        for fi9 := 0; fi9 < 20; fi9++ {
            var v10 int = fi9
            var v11 float32 = -2.25
            var v12 uint32 = (v2 - (v2 % ((0) | 1)))
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
        if ((-(fi8)) <= (fi8 - 127)) {
            for fi13 := 0; fi13 < 13; fi13++ {
                v14 := GenS1{F0: -9223372036854775808}
                if (((v2 & 4294967295) != 90) && ((((((1.5 * -2.25) + float32(-2.25)) + ((1.5 / 1.0) + float32(1.5))) + 0.0) * 0.0) != 0.1)) {
                    var v15 uint = 871
                    _ = v15
                } else {
                    v16 := v5["foo bar"]
                    _ = v16
                }
                v17 := GenS1{F0: -9223372036854775808}
                _ = fi13
                _ = v14
                _ = v17
            }
        }
        if len(v3) > 0 {
            v3[0] = fi8
        }
        _ = fi8
    }
    v2 = 2147483648
    fmt.Println(v4, v1)
    var v18 int32 = -1
    v5["hello"] = 2147483648
    v18 = (v18 - v18)
    v2 += (^(v2))
    v1 = v1
    for fi19 := 0; fi19 < 16; fi19++ {
        var v20 uint64 = 0
        v5["foo bar"] = (v2 ^ 2)
        _ = fi19
        _ = v20
    }
    v18 = v18
    fmt.Println(v18)
    for fi21 := 0; fi21 < 5; fi21++ {
        v22 := GenS1{F0: -9223372036854775808}
        var v23 int = -987
        _ = fi21
        _ = v22
        _ = v23
    }
    v2 -= v2
    v18 += (^(v18))
    v1 = v1
    for fi24 := 0; fi24 < 6; fi24++ {
        fmt.Println(v2, v4, v1)
        _ = fi24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
