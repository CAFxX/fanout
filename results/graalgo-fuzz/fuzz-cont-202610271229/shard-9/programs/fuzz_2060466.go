package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 float32 = (((float32(1e10) - (-2.25 * 3.14159)) + 0.0) * float32(1.5))
    v2 := []string{("hello" + "\n\t"), (("\n\t" + (("\n\t" + "hello") + (("\n\t" + "hello") + ("hello" + "\n\t")))) + "")}
    v3 := make(map[int]int)
    v1 = 3.14159
    v1 -= float32(0.0)
    if false {
        v4 := GenS1{F0: 127}
        _ = v4
    }
    v5 := GenS1{F0: 147}
    var v6 string = ("foo bar" + "\n\t")
    if (588 == 4294967295) {
        if (!false) {
            var v7 float64 = 1.5
            v7 = v7
            fmt.Println(v7)
            _ = v7
        }
        var v8 bool = (v6 == v6)
        _ = v8
    }
    var v9 int64 = 0
    var v10 float64 = 1e10
    v11 := GenS1{F0: 255}
    v12 := GenS1{F0: 236}
    v9 = v9
    v3[1] = 4294967295
    if true {
        if false {
            v13 := GenS1{F0: 82}
            var v14 string = ((((v6 + "hello") + v6) + (v6 + v6)) + (("hello" + v6) + ((("hello" + "hello") + v6) + "\n\t")))
            _ = v13
            _ = v14
        }
        fmt.Println(v1, v9, v10)
    } else {
        var v15 int = 9223372036854775807
        var v16 float32 = (1e10 - (v1 / v1))
        _ = v15
        _ = v16
    }
    var v17 string = (v6 + (v6 + ((v6 + v6) + "hello")))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v9
}
