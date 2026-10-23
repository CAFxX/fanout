package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 uint; F1 int8 }
func genFn0(p0 uint16, p1 bool) string {
    var v4 int16 = -620
    v5 := GenS1{F0: (^(v4))}
    v6 := []int{247, -517}
    p0 = (^(p0))
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return "\n\t"
}
func genFn1(p0 uint16) string {
    v7 := GenS2{F0: -32768, F1: 42}
    switch 32768 {
    case 0:
        var v8 int32 = 586
        _ = v8
    }
    if ((true || (-1 < -1)) || true) {
        p0 += p0
        switch -9223372036854775808 {
        case 2147483647:
            for fi9 := 0; fi9 < 4; fi9++ {
                fmt.Println(p0, p0)
                _ = fi9
            }
        case -699:
            var v10 float64 = 0.0
            _ = v10
        }
    }
    _ = p0
    _ = v7
    return "foo bar"
}

func main() {
    v1 := make(map[string]bool)
    var v2 float32 = 3.14159
    v3 := []int{42, 0, 32768}
    var v11 int16 = 1000
    fmt.Println(v2, v11)
    _ = len(v3)
    switch ("hello" + "\n\t") {
    case "":
        v11 += (v11 / ((-1) | 1))
    default:
        if (1e10 > (float64(float32(-2.25)) - 0.0)) {
            v12 := GenS3{F0: 504, F1: int8(v11)}
            v11 += v11
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 1; fi13++ {
        var v14 int16 = v11
        _ = fi13
        _ = v14
    }
    var v15 int8 = 127
    var v16 int32 = -2147483648
    v17 := v1["\n\t"]
    fmt.Println(v2, v16, v16)
    v18 := GenS2{F0: (v11 / ((v11) | 1)), F1: 195}
    _ = len(v1)
    var v19 float64 = (-2.25 * 3.14159)
    var v20 uint8 = 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
