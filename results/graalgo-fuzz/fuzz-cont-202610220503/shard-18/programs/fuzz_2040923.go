package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int16 }
func genFn0(p0 int8) int64 {
    v4 := []int{32767, -2147483648, 32768}
    v4 = append(v4, 1)
    var v5 string = "\n\t"
    var v6 bool = true
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 string, p1 uint16) string {
    p1 = p1
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return "foo bar"
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int8 = 120
    var v3 uint8 = 2
    v1[2147483647] = 2147483647
    for fi7 := 0; fi7 < 1; fi7++ {
        if ((float32((0.1 - 0.0)) / (1e10 * 3.14159)) == 0.1) {
            var v8 float64 = 0.0
            _ = v8
        }
        v9 := v1[2147483648]
        _ = fi7
        _ = v9
    }
    var v10 uint16 = 471
    v11 := GenS1{F0: v10, F1: "hello", F2: int16(v10)}
    v12 := GenS1{F0: v10, F1: (("foo bar" + ("foo bar" + "foo bar")) + ("" + "hello")), F2: 198}
    v2 -= -5
    var v13 float32 = 0.0
    v10 += v10
    v13 = (1.5 + (3.14159 - v13))
    var v14 string = "hello"
    if (false || false) {
        var v15 uint8 = (v3 - 1)
        var v16 uint16 = (v10 ^ 2)
        v17 := GenS1{F0: (^(v10)), F1: v14, F2: -32768}
        _ = v15
        _ = v16
        _ = v17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
}
