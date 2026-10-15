package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 int32 }
func genFn0(p0 string) string {
    var v6 string = ((p0 + p0) + "a")
    v7 := []int{0}
    fmt.Println(p0)
    if false {
        v6 = (("foo bar" + ((p0 + v6) + ("a" + "\n\t"))) + v6)
        v6 = "hello"
        if len(v7) > 0 {
            v7[0] = 128
        }
    } else {
        if true {
            p0 = v6
            var v8 float32 = ((3.14159 - (0.1 * (((1.5 / float32(((0.0 - 1.5) * float64((0.1 * 1.5))))) / 3.14159) + ((-2.25 - 1.5) + (0.0 - ((0.1 * (3.14159 - 0.0)) * 1e10)))))) / float32(1e10))
            var v9 int16 = 341
            _ = v8
            _ = v9
        } else {
            fmt.Println(p0)
        }
    }
    _ = p0
    _ = v6
    _ = v7
    return v6
}

func main() {
    var v1 float32 = 1.5
    var v2 float32 = 3.14159
    var v3 uint16 = 0
    var v4 float32 = ((0.1 / v2) * -2.25)
    v5 := []string{}
    if (false && true) {
        v10 := GenS1{F0: 32767}
        _ = v10
    }
    v11 := GenS2{F0: 394}
    if true {
        v12 := GenS3{F0: 50, F1: -670}
        var v13 float64 = (1.5 - 0.1)
        var v14 uint8 = 127
        _ = v12
        _ = v13
        _ = v14
    }
    v5 = append(v5, (("a" + "a") + ("" + "hello")))
    fmt.Println(v2)
    v5 = append(v5, "\n\t")
    if ((!false) || false) {
        v15 := GenS4{F0: -2147483648}
        v3 -= v3
        var v16 float64 = (1.5 * (0.0 / 0.1))
        _ = v15
        _ = v16
    } else {
        if len(v5) > 0 {
            v5[0] = "\n\t"
        }
    }
    if len(v5) > 0 {
        v5[0] = ""
    }
    var v17 float32 = v4
    fmt.Println(v17, v3, v1)
    fmt.Println(v17, v4)
    var v18 int16 = -79
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
