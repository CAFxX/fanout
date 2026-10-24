package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }

func main() {
    var v1 int64 = 0
    var v2 int = 32768
    v3 := make(map[int]uint32)
    var v4 int32 = -994
    v5 := make(map[int]bool)
    switch ("" + "hello") {
    case "\n\t":
        v6 := GenS1{F0: "\n\t", F1: true}
        _ = v6
    default:
        v7 := GenS1{F0: "", F1: true}
        _ = v7
    }
    var v8 float32 = (0.0 * (3.14159 - (-2.25 * float32((0.0 * (1e10 / ((float64(0.1) - 1.5) / 1.0)))))))
    fmt.Println(v8, v2, v2)
    v9 := GenS1{F0: "", F1: (v2 > int(v1))}
    var v10 bool = true
    for fi11 := 0; fi11 < 4; fi11++ {
        fmt.Println(v1)
        _ = fi11
    }
    switch ("\n\t" + ("" + "")) {
    case "a":
        var v12 uint16 = 32767
        _ = v12
    default:
        var v13 uint32 = 65535
        _ = v13
    }
    var v14 float32 = v8
    for fi15 := 0; fi15 < 6; fi15++ {
        var v16 int8 = 37
        v17 := []int{fi15, (v2 << ((65535) & 63))}
        _ = fi15
        _ = v16
        _ = v17
    }
    v18 := []int{703, (^(v2)), v2}
    v4 += 1
    v19 := GenS1{F0: ((("hello" + ("foo bar" + "foo bar")) + (("" + ("" + "a")) + ((("hello" + "hello") + "a") + ("a" + "hello")))) + ((("a" + "a") + "foo bar") + ("\n\t" + (("" + "a") + "")))), F1: v10}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
