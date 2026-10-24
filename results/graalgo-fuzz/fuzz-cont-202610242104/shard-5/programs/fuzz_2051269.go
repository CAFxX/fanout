package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 int64 }
type GenS2 struct { F0 int8; F1 uint16 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]int)
    var v3 uint16 = 0
    var v4 bool = (127 > 1)
    delete(v1, "")
    v5 := GenS1{F0: 923, F1: 1, F2: -1}
    v4 = true
    var v6 int = 88
    var v7 int64 = -9223372036854775808
    v8 := GenS1{F0: 0, F1: 2147483648, F2: v7}
    v3 += (v3 >> ((4294967295) & 63))
    fmt.Println(v4)
    fmt.Println(v7, v4, v6)
    if v4 {
        if (921 != 1) {
            v2["a"] = 42
            for fi9 := 0; fi9 < 2; fi9++ {
                v10 := []int{fi9, fi9, v6}
                _ = fi9
                _ = v10
            }
        }
    } else {
        v4 = ((((((1.5 / ((1e10 / (((0.1 * 1e10) / (-2.25 * 0.1)) / (1e10 * 3.14159))) / 3.14159)) * ((3.14159 + float32(float64((0.1 / (-2.25 + 0.0))))) / -2.25)) - 0.0) - 1e10) <= float32((3.14159 / (3.14159 + (-2.25 / 1.0))))) && (498 == 18446744073709551615))
        v11 := []int{v6, v6}
        _ = v11
    }
    var v12 uint64 = 516
    v7 = v7
    if (255 != 4294967295) {
        switch "hello" {
        case "":
            v4 = (!(-209 <= -585))
        case "foo bar":
            v13 := []int{v6}
            _ = v13
        case "a":
            var v14 int16 = -32768
            _ = v14
        default:
            v15 := []int{(v6 - v6), v6}
            _ = v15
        }
        v16 := []int{v6, (v6 % ((v6) | 1))}
        _ = v16
    } else {
        if (!v4) {
            if v4 {
                switch (v6 % ((v6) | 1)) {
                case 32767:
                    var v17 int16 = -1
                    _ = v17
                }
                fmt.Println(v7, v7, v7)
                fmt.Println(v4)
            } else {
                fmt.Println(v7, v4, v3)
            }
            v6 = v6
            v1[(("hello" + ("a" + "foo bar")) + "hello")] = (^(v6))
        }
        v18 := GenS2{F0: -93, F1: v3}
        v6 += -2147483648
        _ = v18
    }
    v19 := []int{v6}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
