package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint16 }
func genFn0() bool {
    v6 := []int{128, -9223372036854775808}
    var v7 int64 = 42
    v8 := GenS1{F0: v7, F1: 717, F2: 42}
    var v9 int32 = -689
    var v10 bool = false
    _ = v10
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return (v10 && (42 == 65535))
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[int]bool)
    var v3 uint = 4294967295
    var v4 int8 = -128
    v5 := []string{"a", "a", "a"}
    fmt.Println(v3, v3)
    v5 = append(v5, ((((("a" + "\n\t") + "hello") + "a") + ("hello" + (((("" + "hello") + "a") + ("a" + "foo bar")) + (("" + "hello") + "\n\t")))) + ("hello" + ("" + ""))))
    v4 = -70
    if false {
        if true {
            v1 -= (v1 * v1)
            var v11 int32 = -810
            var v12 uint32 = 0
            _ = v11
            _ = v12
        }
        if (false || false) {
            v1 += (v1 >> ((65535) & 63))
        } else {
            v13 := GenS1{F0: -1, F1: 9223372036854775807, F2: 255}
            _ = len(v5)
            if (3.14159 <= ((float64((1e10 - 0.1)) / 1.0) + 3.14159)) {
                fmt.Println(v3, v4, v4)
            } else {
                v14 := GenS1{F0: 42, F1: -876, F2: 255}
                var v15 int = 0
                var v16 int64 = -503
                _ = v14
                _ = v15
                _ = v16
            }
            _ = v13
        }
    } else {
        for fi17 := 0; fi17 < 8; fi17++ {
            v1 = (v1 * 20)
            _ = fi17
        }
    }
    fmt.Println(v3, v4)
    switch "foo bar" {
    case "\n\t":
        v3 = v3
    case "foo bar":
        v2[42] = true
    }
    var v18 bool = ((-9223372036854775808 == 9223372036854775807) || false)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v4, v1, v3)
    v1 = (v1 - v1)
    var v19 string = ("a" + "hello")
    var v20 uint32 = 38
    if false {
        v21 := GenS1{F0: 0, F1: 2147483647, F2: 255}
        var v22 float64 = float64(float32(1e10))
        _ = v21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
