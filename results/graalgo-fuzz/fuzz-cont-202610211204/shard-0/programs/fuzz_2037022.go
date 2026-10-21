package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 string = "foo bar"
    v2 := make(map[string]bool)
    var v3 uint8 = 42
    v4 := make(map[int]string)
    var v5 uint = 449
    var v6 int8 = 1
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 float64 = -2.25
        v9 := GenS1{F0: (^(v5))}
        fmt.Println(v6, v1)
        _ = fi7
        _ = v8
        _ = v9
    }
    fmt.Println(v1, v6)
    if ((1 == 0) && true) {
        v3 = 74
    }
    if (!false) {
        var v10 uint64 = 42
        v1 = "a"
        _ = v10
    } else {
        if true {
            v6 += (v6 / ((v6) | 1))
        } else {
            v3 = v3
        }
        v11 := GenS1{F0: (v5 >> ((694) & 63))}
        fmt.Println(v5, v5, v3)
        _ = v11
    }
    delete(v4, 127)
    v3 += (v3 | (v3 % (((v3 & v3)) | 1)))
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 float32 = 0.0
        _ = fi12
        _ = v13
    }
    switch 42 {
    case -9223372036854775808:
        v14 := v4[32767]
        _ = v14
    }
    v5 = (^(v5))
    var v15 float32 = (1.5 / 3.14159)
    if false {
        v15 = v15
        var v16 string = ((v1 + (("\n\t" + "foo bar") + "\n\t")) + (v1 + v1))
        _ = v16
    }
    var v17 int32 = -674
    v18 := []int{-2147483648, 64, 73}
    var v19 string = ((("foo bar" + v1) + "a") + v1)
    v19 = "a"
    v18 = append(v18, 942)
    v18 = append(v18, 127)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
