package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 uint8 = 128
    v2 := make(map[string]string)
    var v3 bool = (-153 > 9223372036854775807)
    var v4 string = "hello"
    var v5 int32 = 872
    v6 := []int{-593, 4294967295, -9223372036854775808}
    if (v3 || v3) {
        var v7 uint8 = v1
        _ = v7
    }
    var v8 string = v4
    v9 := GenS1{F0: v1}
    var v10 int8 = -50
    var v11 int32 = v5
    fmt.Println(v4, v5)
    var v12 float32 = 0.1
    if (-1 <= -9223372036854775808) {
        if (617 <= 571) {
            delete(v2, v4)
        } else {
            v6 = append(v6, 32767)
            v10 = int8(v1)
            v5 = (-(v5))
        }
        var v13 float64 = (0.1 / (1e10 - 1.5))
        _ = v13
    }
    fmt.Println(v1, v4)
    for fi14 := 0; fi14 < 11; fi14++ {
        v15 := GenS2{F0: 4294967295}
        _ = fi14
        _ = v15
    }
    if len(v6) > 0 {
        v6[0] = -2147483648
    }
    delete(v2, v4)
    if (737 > 2) {
        var v16 int32 = v5
        v2[("" + "hello")] = "foo bar"
        if (-998 == -137) {
            var v17 bool = v3
            _ = len(v2)
            fmt.Println(v12, v8)
            _ = v17
        }
        _ = v16
    } else {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    v12 = v12
    for fi18 := 0; fi18 < 9; fi18++ {
        v19 := GenS1{F0: (v1 >> ((2) & 63))}
        if (((false || true) || v3) && v3) {
            if (9223372036854775807 > -1) {
                var v20 uint = 516
                var v21 uint8 = v1
                _ = v20
                _ = v21
            } else {
                v5 += v5
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
            if (!(!(!((0.1 * 1.5) != (1.5 + ((1e10 / 3.14159) / -2.25)))))) {
                var v22 uint8 = 127
                _ = v22
            }
        }
        _ = fi18
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
