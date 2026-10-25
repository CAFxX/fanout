package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
func genFn0(p0 int64) uint8 {
    var v7 float64 = ((3.14159 - (-2.25 + -2.25)) * 1e10)
    var v8 int8 = 127
    v7 = 1e10
    if (0.0 < (float32(1e10) / (float32(float64(((0.0 / ((float32((v7 * float64(0.0))) - 1e10) - 3.14159)) + (1e10 - float32(v7))))) - 0.1))) {
        v8 += 0
        fmt.Println(p0, v8)
    }
    var v9 int = -215
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 uint8 = 159
    var v2 float64 = 1e10
    var v3 uint8 = 127
    var v4 int32 = -227
    var v5 int16 = -560
    v6 := []int{0, 4294967295}
    _ = len(v6)
    v2 = (0.0 + 3.14159)
    fmt.Println(v4, v3, v2)
    v2 = (v2 * (v2 - 1e10))
    if len(v6) > 0 {
        v6[0] = 2147483647
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        for fi11 := 0; fi11 < 5; fi11++ {
            var v12 float64 = v2
            fmt.Println(v3, fi11, v4)
            var v13 uint32 = 304
            v6 = append(v6, 397)
            var v14 int32 = (^(v4))
            v6 = append(v6, fi10)
            switch (fi11 - fi10) {
            case 280:
                fmt.Println(v5, v13)
                _ = len(v6)
                var v15 uint8 = 136
                _ = v15
            case 4294967295:
                var v16 bool = true
                _ = v16
            }
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        _ = fi10
    }
    v1 -= (v1 << ((792) & 63))
    for fi17 := 0; fi17 < 17; fi17++ {
        for fi18 := 0; fi18 < 6; fi18++ {
            switch ("a" + (((("a" + "hello") + (("\n\t" + "foo bar") + "foo bar")) + "hello") + "a")) {
            case "":
                v19 := GenS1{F0: false, F1: v4}
                _ = v19
            default:
                var v20 int = (fi17 & (fi17 | fi18))
                _ = v20
            }
            _ = fi18
        }
        if true {
            v6 = append(v6, (fi17 ^ 128))
            v21 := GenS1{F0: true, F1: v4}
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = v21
        } else {
            v6 = append(v6, 9223372036854775807)
        }
        _ = fi17
    }
    fmt.Println(v1, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
