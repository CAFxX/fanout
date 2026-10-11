package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 string }
type GenS2 struct { F0 uint64; F1 int32 }
func genFn0(p0 int16, p1 uint8, p2 bool) string {
    fmt.Println(p0, p2, p0)
    p2 = p2
    if (2 > 192) {
        var v5 uint64 = 1
        fmt.Println(v5)
        var v6 uint8 = uint8(v5)
        _ = v5
        _ = v6
    }
    if p2 {
        p2 = (0.0 < (((((0.0 * float64(0.1)) * 1.5) + (float64(((-2.25 - 0.1) * ((float32(0.1) / 3.14159) / 1.5))) + float64((3.14159 + (1e10 + -2.25))))) * 0.1) / (3.14159 + (-2.25 + (0.1 / (((0.1 / 1.5) * -2.25) * 0.1))))))
    } else {
        var v7 int8 = 12
        _ = v7
    }
    v8 := GenS1{F0: (p1 - p1), F1: 1, F2: "a"}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return ""
}

func main() {
    v1 := []int{42}
    var v2 int16 = -641
    var v3 float32 = 0.1
    var v4 float32 = v3
    var v9 string = "\n\t"
    v4 += v3
    if len(v1) > 0 {
        v1[0] = 607
    }
    if true {
        if (-2147483648 < 4294967295) {
            v10 := GenS2{F0: 4294967295, F1: 1}
            if ((42 >= 1) && (!true)) {
                v1 = append(v1, -9223372036854775808)
            } else {
                v4 = v4
                fmt.Println(v3)
                var v11 float64 = 3.14159
                _ = v11
            }
            _ = v10
        }
        v12 := GenS1{F0: 42, F1: -884, F2: v9}
        if (v2 >= v2) {
            var v13 int16 = (v2 ^ v2)
            switch ((v9 + (v9 + (((v9 + "\n\t") + v9) + "hello"))) + v9) {
            case "hello":
                v14 := GenS1{F0: 71, F1: 0, F2: "\n\t"}
                _ = v14
            case "\n\t":
                fmt.Println(v4, v9)
            case "a":
                v15 := GenS1{F0: 51, F1: -998, F2: v9}
                _ = v15
            }
            v2 = (v13 / ((32767) | 1))
            _ = v13
        }
        _ = v12
    }
    var v16 int16 = 580
    var v17 float32 = v3
    var v18 uint = 255
    fmt.Println(v3, v16, v9)
    _ = len(v1)
    var v19 uint8 = 2
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
