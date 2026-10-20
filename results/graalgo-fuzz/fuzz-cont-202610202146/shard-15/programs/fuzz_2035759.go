package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 int8 }
func genFn0(p0 int) uint8 {
    p0 = p0
    p0 += (p0 / (((p0 << ((0) & 63))) | 1))
    v4 := []int{(p0 >> ((4294967295) & 63)), -1, 128}
    var v5 int32 = 2147483647
    _ = p0
    _ = v4
    _ = v5
    return 42
}

func main() {
    var v1 uint = 4294967295
    var v2 int32 = -945
    v3 := make(map[string]bool)
    var v6 bool = (true || false)
    if true {
        v2 -= -220
        v7 := []int{42}
        var v8 string = ("hello" + "")
        _ = v7
        _ = v8
    }
    v2 = (v2 + v2)
    var v9 int8 = 87
    var v10 uint16 = 355
    v2 = 2147483647
    if v6 {
        v11 := GenS1{F0: (v1 - (^(v1)))}
        _ = v11
    }
    if (v6 || v6) {
        v12 := GenS2{F0: 0, F1: -9223372036854775808}
        _ = v12
    } else {
        if (-614 != int(v9)) {
            if v6 {
                fmt.Println(v1, v10)
            } else {
                v13 := v3[""]
                if (!(!v6)) {
                    fmt.Println(v6)
                    var v14 uint8 = 105
                    if (42 < 980) {
                        v15 := []int{888, -2147483648, 128, 228}
                        var v16 int8 = -41
                        _ = v15
                        _ = v16
                    }
                    _ = v14
                }
                v17 := GenS3{F0: -109}
                _ = v13
                _ = v17
            }
            switch "foo bar" {
            case "foo bar":
                v18 := GenS1{F0: (^(v1))}
                _ = v18
            default:
                delete(v3, "a")
            }
            v2 = v2
        }
    }
    var v19 int16 = int16(v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v3
    _ = v6
    _ = v9
}
