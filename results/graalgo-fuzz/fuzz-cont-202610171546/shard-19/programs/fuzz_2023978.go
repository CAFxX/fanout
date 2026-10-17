package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 uint16 }
func genFn0(p0 int16, p1 int8, p2 int64) int {
    switch int(p1) {
    case 0:
        v4 := []int{-1, -9223372036854775808, 292, 180}
        _ = v4
    case 128:
        fmt.Println(p0)
    }
    v5 := GenS1{F0: (^(p1))}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 0
}

func main() {
    var v1 string = "\n\t"
    var v2 float32 = (1.5 - (float32((3.14159 + 0.1)) / 1.5))
    v3 := []int{32767, 4294967295}
    v3 = append(v3, 0)
    for fi6 := 0; fi6 < 13; fi6++ {
        v7 := GenS1{F0: -45}
        _ = fi6
        _ = v7
    }
    fmt.Println(v2)
    if false {
        if (true && (18446744073709551615 != 932)) {
            v2 -= v2
        } else {
            _ = len(v3)
            v3 = append(v3, 9223372036854775807)
        }
        v2 += v2
        var v8 uint8 = 2
        _ = v8
    }
    switch v1 {
    case "foo bar":
        if (1 != 292) {
            var v9 float32 = float32(0.0)
            if true {
                v9 -= ((v9 / 1.5) * (v2 * float32((-2.25 - -2.25))))
            }
            var v10 int8 = 1
            _ = v10
            _ = v9
        }
    case "a":
        var v11 uint = 793
        _ = v11
    }
    var v12 float32 = float32(0.0)
    var v13 int8 = 42
    v13 -= (v13 + (v13 + v13))
    v14 := GenS1{F0: v13}
    for fi15 := 0; fi15 < 20; fi15++ {
        var v16 uint8 = 128
        v13 = (v13 >> ((4294967295) & 63))
        fmt.Println(v2)
        _ = fi15
        _ = v16
    }
    var v17 uint32 = 2147483647
    for fi18 := 0; fi18 < 15; fi18++ {
        for fi19 := 0; fi19 < 13; fi19++ {
            v3 = append(v3, fi18)
            if (((0.0 <= v2) || true) || (1.5 <= (float64(((float32(float64((v12 * v2))) + v2) / v12)) * (-2.25 * (0.0 / 3.14159))))) {
                fmt.Println(fi18, v17)
            } else {
                fmt.Println(fi19, v1, v13)
                v20 := GenS2{F0: genFn0(42, (v13 >> ((65535) & 63)), 0), F1: 65535}
                switch v1 {
                case "foo bar":
                    v21 := GenS2{F0: (fi19 % ((fi19) | 1)), F1: 409}
                    _ = v21
                case "a":
                    v12 -= -2.25
                }
                _ = v20
            }
            _ = fi19
        }
        _ = fi18
    }
    var v22 uint = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v22
    _ = v3
}
