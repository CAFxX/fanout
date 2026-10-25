package main

import (
    "fmt"
)

func genFn0(p0 int16, p1 int16) string {
    var v5 uint = 2
    v6 := []int{652, 271, 32768}
    v5 += (v5 / ((964) | 1))
    p0 = 32767
    p0 = (p0 / (((-(p0))) | 1))
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return ("foo bar" + (("foo bar" + ("foo bar" + "\n\t")) + ("" + "")))
}

func main() {
    v1 := []int{127, -610, 1}
    var v2 uint8 = 255
    var v3 uint32 = 657
    var v4 uint16 = 65535
    var v7 float32 = float32((1e10 - 1e10))
    var v8 float32 = v7
    v2 = v2
    v7 = (v7 + (((1e10 * (v7 * 1e10)) - 1e10) / v7))
    if (!(!true)) {
        v2 = (^(v2))
    }
    if true {
        var v9 bool = true
        var v10 uint32 = v3
        var v11 float64 = float64((float32(((1e10 / 1.5) - ((3.14159 / 1.5) + 0.1))) / float32((-2.25 - 0.1))))
        _ = v10
        _ = v11
        _ = v9
    } else {
        v8 += (0.0 * 1.5)
    }
    fmt.Println(v3, v2, v8)
    fmt.Println(v8, v2, v2)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
