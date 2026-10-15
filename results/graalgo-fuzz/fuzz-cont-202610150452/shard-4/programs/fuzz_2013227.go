package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
func genFn0(p0 int8, p1 uint16) string {
    fmt.Println(p0)
    v5 := GenS1{F0: 65535, F1: -185}
    v6 := []int{1, 1}
    if true {
        var v7 uint32 = 42
        var v8 float64 = 0.1
        if ((false || (true || true)) && false) {
            if (!((!(false && false)) && (0 < v7))) {
                p1 += (^(p1))
                if (-8 > -9223372036854775808) {
                    if true {
                        v9 := GenS1{F0: 42, F1: 2147483647}
                        _ = v9
                    } else {
                        var v10 string = "\n\t"
                        _ = v10
                    }
                    var v11 float32 = (-2.25 / (((float32(3.14159) + -2.25) + -2.25) - -2.25))
                    _ = v11
                }
            }
            for fi12 := 0; fi12 < 13; fi12++ {
                p0 = (p0 / (((p0 >> ((4294967295) & 63))) | 1))
                _ = fi12
            }
        } else {
            fmt.Println(v7)
            var v13 bool = (0 != 128)
            _ = v13
        }
        _ = v7
        _ = v8
    } else {
        v14 := GenS1{F0: 255, F1: 4294967295}
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return ("hello" + "foo bar")
}

func main() {
    v1 := []uint32{215, 2147483647}
    var v2 uint = 18446744073709551615
    v3 := make(map[uint32]uint32)
    var v4 bool = ((1e10 <= 0.0) && ((("" + "a") == "hello") && false))
    var v15 uint16 = 255
    v3[412] = 968
    v2 += (v2 ^ 0)
    var v16 bool = (uint64(v2) > 42)
    fmt.Println(v16)
    var v17 uint8 = 127
    v17 -= (^(v17))
    fmt.Println(v17, v4, v4)
    v3[65535] = 42
    var v18 uint8 = (v17 / ((uint8(v15)) | 1))
    var v19 int64 = 0
    v20 := GenS1{F0: 42, F1: 0}
    var v21 bool = (127 == 127)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
