package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint }
func genFn0(p0 int) uint8 {
    switch "\n\t" {
    case "\n\t":
        if (42 >= 1000) {
            fmt.Println(p0)
        }
    }
    fmt.Println(p0)
    p0 = (p0 & (p0 / ((p0) | 1)))
    _ = p0
    return 128
}
func genFn1(p0 int8, p1 uint64, p2 uint32) uint8 {
    var v4 uint16 = 1
    var v5 uint32 = (p2 << ((875) & 63))
    switch ((("hello" + "hello") + "hello") + "a") {
    case "a":
        var v6 int = 0
        _ = v6
    case "\n\t":
        fmt.Println(v4, v4)
    default:
        switch 413 {
        case -2147483648:
            v7 := []int{128, 657, -855, 32767}
            _ = v7
        default:
            v8 := GenS1{F0: 662}
            _ = v8
        }
    }
    var v9 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v9
    return 128
}

func main() {
    var v1 int8 = -51
    var v2 int = 0
    var v3 uint16 = 255
    for fi10 := 0; fi10 < 17; fi10++ {
        v11 := []int{2147483647}
        for fi12 := 0; fi12 < 2; fi12++ {
            v1 += (v1 + -30)
            var v13 int32 = 0
            _ = fi12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    var v14 uint8 = 2
    var v15 uint16 = (v3 << ((255) & 63))
    v3 = v15
    v15 = 2
    v16 := []int{(-(v2)), v2}
    var v17 int8 = (v1 + -1)
    if len(v16) > 0 {
        v16[0] = v2
    }
    v18 := GenS2{F0: 2}
    fmt.Println(v3, v17)
    switch (("foo bar" + (("foo bar" + "hello") + "\n\t")) + (("a" + "") + "\n\t")) {
    case "hello":
        var v19 int16 = 1
        _ = v19
    case "":
        fmt.Println(v15)
    case "\n\t":
        var v20 int64 = -9223372036854775808
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
