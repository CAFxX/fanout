package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 string }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]int)
    switch ("hello" + ("" + "")) {
    case "\n\t":
        var v4 int = 1
        _ = v4
    case "hello":
        delete(v3, 65535)
    }
    var v5 uint32 = 2147483648
    var v6 float32 = float32(0.0)
    var v7 int = 1
    for fi8 := 0; fi8 < 19; fi8++ {
        fmt.Println(v5, v7, v6)
        v5 += 484
        _ = fi8
    }
    var v9 bool = (!(65535 >= 801))
    v10 := []int{v7}
    var v11 uint8 = 2
    if len(v10) > 0 {
        v10[0] = v7
    }
    fmt.Println(v7, v5, v9)
    for fi12 := 0; fi12 < 5; fi12++ {
        v7 = fi12
        var v13 int32 = 990
        _ = fi12
        _ = v13
    }
    v14 := GenS1{F0: -9223372036854775808, F1: "hello", F2: ("" + "")}
    v15 := GenS1{F0: -983, F1: "foo bar", F2: (("foo bar" + "\n\t") + "a")}
    for fi16 := 0; fi16 < 14; fi16++ {
        var v17 uint = 643
        _ = fi16
        _ = v17
    }
    v6 = (((v6 / v6) + (((v6 * float32((0.1 * -2.25))) + (float32(3.14159) + -2.25)) + 0.0)) / (((v6 * -2.25) * (((v6 / (v6 + 3.14159)) / (float32(-2.25) + 0.0)) - float32(-2.25))) * 1.5))
    v11 = v11
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
