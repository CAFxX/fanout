package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]int)
    var v2 float64 = 1e10
    v3 := make(map[uint32]int)
    var v4 int8 = -128
    fmt.Println(v4, v2)
    if true {
        delete(v3, 2)
        fmt.Println(v4, v2)
        if true {
            if true {
                if (((v2 / 0.1) - v2) == v2) {
                    _ = len(v1)
                    var v5 float64 = v2
                    _ = v5
                }
                v2 = v2
                if (0.1 == (0.0 * v2)) {
                    fmt.Println(v4)
                    fmt.Println(v2, v2)
                }
            } else {
                fmt.Println(v4, v2)
            }
        } else {
            v4 += v4
            v4 -= v4
        }
    } else {
        v6 := []int{2147483647, 16, -1}
        fmt.Println(v4, v4)
        switch (("foo bar" + (("\n\t" + ("hello" + "hello")) + ("" + ("hello" + (("foo bar" + ("foo bar" + (("hello" + "a") + ("a" + "a")))) + ""))))) + "hello") {
        case "hello":
            v2 -= ((((((1e10 + (v2 + (1e10 + ((((float64(1e10) * v2) * v2) - v2) - (v2 + float64(1e10)))))) - ((((v2 * (v2 - v2)) - float64(1e10)) * float64((float32(((float64(float32((v2 + (0.0 / ((float64((3.14159 - ((3.14159 / 1.5) + 0.0))) / 1.0) - v2))))) - ((v2 + float64((-2.25 / 0.1))) * 3.14159)) / 1e10)) / -2.25))) * v2)) + float64((0.0 / -2.25))) + v2) * 0.1) / (v2 + v2))
        default:
            var v7 int64 = -1
            _ = v7
        }
        _ = v6
    }
    v2 = float64(1.5)
    v4 = v4
    v8 := GenS1{F0: -829}
    v2 += ((v2 * (v2 + 1.5)) - (v2 - v2))
    var v9 int8 = 42
    v10 := GenS1{F0: -635}
    var v11 int = 32768
    fmt.Println(v9, v2, v11)
    var v12 uint32 = 682
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
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
