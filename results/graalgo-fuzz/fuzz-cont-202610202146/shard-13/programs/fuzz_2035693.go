package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 float32 = (float32(3.14159) - 0.0)
    var v2 float32 = v1
    var v3 int8 = -1
    var v4 uint8 = 127
    if true {
        v5 := GenS1{F0: 32767}
        v6 := []int{32768, -9223372036854775808, -542}
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: -8}
    v3 += -128
    v8 := GenS1{F0: 1000}
    v9 := []int{int(v3), -553, -9223372036854775808}
    switch "foo bar" {
    case "foo bar":
        if (true || true) {
            var v10 bool = false
            v1 = 0.0
            _ = v10
        }
    case "\n\t":
        if true {
            v2 -= ((v2 + 1e10) - 1.5)
            var v11 float64 = -2.25
            _ = v11
        }
    }
    var v12 string = "a"
    v12 = (((v12 + v12) + v12) + ("a" + ("foo bar" + ("a" + ((v12 + "a") + (("" + "a") + v12))))))
    if true {
        _ = len(v9)
    } else {
        v13 := GenS1{F0: -1}
        _ = v13
    }
    v3 = v3
    for fi14 := 0; fi14 < 4; fi14++ {
        fmt.Println(fi14, v1)
        _ = fi14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
