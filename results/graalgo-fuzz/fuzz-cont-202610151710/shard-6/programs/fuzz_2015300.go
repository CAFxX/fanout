package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 uint32) uint64 {
    v5 := GenS1{F0: 420}
    v6 := []int{1, 4294967295, -2147483648, 4294967295}
    _ = p0
    _ = v5
    _ = v6
    return 2
}

func main() {
    var v1 int32 = 109
    var v2 int32 = v1
    var v3 uint32 = 2147483648
    var v4 uint64 = 18446744073709551615
    fmt.Println(v1, v4)
    v7 := []int{0, -709, -2147483648, 32767}
    v7 = append(v7, 4294967295)
    v2 = v2
    v3 = (^(v3))
    for fi8 := 0; fi8 < 17; fi8++ {
        for fi9 := 0; fi9 < 5; fi9++ {
            if len(v7) > 0 {
                v7[0] = 623
            }
            var v10 int = fi8
            if false {
                var v11 int32 = (^(v1))
                v3 -= (v3 - v3)
                fmt.Println(v11)
                _ = v11
            } else {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                switch (fi8 << ((65535) & 63)) {
                case 2147483647:
                    v4 = (^(v4))
                default:
                    v4 = 2
                }
                var v12 int8 = 45
                var v13 int8 = 0
                _ = v12
                _ = v13
            }
            _ = fi9
            _ = v10
        }
        for fi14 := 0; fi14 < 11; fi14++ {
            fmt.Println(v3, v1, fi8)
            var v15 float32 = ((((0.1 - (((1e10 / (0.1 / 3.14159)) + 1e10) + 3.14159)) / ((1.5 * (0.1 - (1e10 * 1e10))) / ((1e10 + (-2.25 / (-2.25 + 0.0))) + float32(1.5)))) - 0.0) * 1.5)
            _ = fi14
            _ = v15
        }
        _ = fi8
    }
    v3 -= 42
    fmt.Println(v1)
    v16 := GenS1{F0: 845}
    if len(v7) > 0 {
        v7[0] = 2147483647
    }
    v17 := GenS1{F0: 2}
    var v18 string = ("foo bar" + (((("hello" + "foo bar") + "") + "a") + ""))
    fmt.Println(v18, v2)
    for fi19 := 0; fi19 < 3; fi19++ {
        var v20 float64 = 3.14159
        if (false && false) {
            v3 = 2
            var v21 uint8 = 195
            _ = v21
        }
        _ = fi19
        _ = v20
    }
    v4 = genFn0(65535)
    var v22 string = v18
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    if (0 >= 28) {
        v23 := GenS1{F0: 65535}
        _ = v23
    } else {
        v24 := GenS1{F0: 0}
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v7
}
