package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 int64 }
func genFn0(p0 bool, p1 int32, p2 string) int8 {
    var v5 uint = 643
    if false {
        v6 := GenS1{F0: p1, F1: (v5 - v5), F2: -601}
        v7 := []int{2147483647, 786, -1}
        var v8 bool = (1 >= 28)
        _ = v6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -128
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]uint32)
    var v3 int16 = 1
    var v4 int32 = -779
    var v9 int16 = v3
    if (true && ((v4 / ((495) | 1)) > 117)) {
        fmt.Println(v4)
        v1[1] = "\n\t"
    }
    v3 = -895
    fmt.Println(v3, v9, v4)
    v4 += (-(v4))
    fmt.Println(v3, v3, v4)
    var v10 uint16 = 65535
    if ("foo bar" <= "\n\t") {
        var v11 int32 = (v4 + -945)
        _ = v11
    } else {
        fmt.Println(v3, v10, v9)
    }
    var v12 float32 = 3.14159
    fmt.Println(v12, v12, v9)
    if true {
        v13 := GenS1{F0: (v4 & 2147483647), F1: 0, F2: -9223372036854775808}
        v12 -= v12
        _ = v13
    } else {
        if (v3 < 537) {
            fmt.Println(v3, v3, v12)
            v10 -= v10
        }
        v9 -= v9
    }
    v3 = (v3 % (((^(v3))) | 1))
    var v14 bool = false
    delete(v2, 32767)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
