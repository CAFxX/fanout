package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int64 }
type GenS2 struct { F0 int16; F1 int16; F2 int16 }
type GenS3 struct { F0 int32; F1 int32 }
type GenS4 struct { F0 uint32 }
func genFn0() string {
    var v5 int16 = -32768
    switch "" {
    case "hello":
        v6 := []int{128}
        _ = v6
    case "\n\t":
        if (!false) {
            v7 := []int{-9223372036854775808, 32767, 127, 1}
            _ = v7
        }
    default:
        fmt.Println(v5)
    }
    _ = v5
    return ("a" + "")
}

func main() {
    v1 := make(map[string]string)
    var v2 int8 = 1
    var v3 float32 = (1.5 * 1.5)
    var v4 int32 = 2147483647
    v8 := GenS1{F0: -32768, F1: -1, F2: 9223372036854775807}
    var v9 uint = 247
    var v10 uint32 = 0
    v2 = (^(v2))
    v9 -= v9
    v11 := GenS2{F0: 1000, F1: -480, F2: -502}
    var v12 string = "a"
    v9 += (v9 << ((v9) & 63))
    v3 = -2.25
    switch v12 {
    case "a":
        if false {
            var v13 int64 = 9223372036854775807
            v9 += v9
            _ = v13
        }
    case "":
        if false {
            switch ("" + ("hello" + v12)) {
            case "a":
                if (v4 < (v4 / (((v4 * (v4 * int32(v9)))) | 1))) {
                    v10 = 0
                    v3 -= v3
                } else {
                    fmt.Println(v3, v4, v12)
                }
            case "foo bar":
                fmt.Println(v4, v10)
            }
            fmt.Println(v3, v4)
            v14 := []int{9223372036854775807, 1, -9223372036854775808}
            _ = v14
        } else {
            if ((1 >= 0) && false) {
                v15 := GenS3{F0: v4, F1: 42}
                v9 = v9
                if true {
                    switch -2147483648 {
                    case 957:
                        v16 := []int{9223372036854775807, -2147483648}
                        _ = v16
                    }
                    v3 = (((float32((0.0 + 1.5)) * (v3 + v3)) * v3) - ((((v3 / v3) - v3) + (v3 / (v3 * v3))) * (v3 + (v3 - v3))))
                }
                _ = v15
            }
            var v17 uint16 = 42
            v18 := GenS2{F0: -32768, F1: 1000, F2: 818}
            _ = v17
            _ = v18
        }
    default:
        if (v12 <= "hello") {
            v19 := []int{}
            v20 := GenS4{F0: (v10 + v10)}
            _ = v19
            _ = v20
        } else {
            v21 := []int{-9223372036854775808, -768, -856}
            v22 := v1[""]
            _ = v21
            _ = v22
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
