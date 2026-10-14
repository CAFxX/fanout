package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 float64 = 0.0
    var v2 int32 = 1
    var v3 uint32 = 258
    var v4 uint16 = 1
    var v5 int16 = -466
    var v6 uint8 = 145
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 int32 = (v2 * (v2 + (-(v2))))
        var v9 uint = uint(v6)
        _ = fi7
        _ = v8
        _ = v9
    }
    switch ("\n\t" + (("a" + "a") + "\n\t")) {
    case "a":
        v10 := GenS1{F0: "\n\t"}
        _ = v10
    }
    var v11 int32 = 506
    switch (("hello" + "foo bar") + "") {
    case "\n\t":
        v5 += v5
    default:
        var v12 string = "hello"
        _ = v12
    }
    v13 := []int{9223372036854775807, -432}
    fmt.Println(v3, v4)
    switch 2147483647 {
    case 4294967295:
        for fi14 := 0; fi14 < 9; fi14++ {
            v3 -= v3
            v6 += (v6 >> ((255) & 63))
            _ = fi14
        }
    case -2147483648:
        v5 = 593
    }
    fmt.Println(v1, v4, v6)
    fmt.Println(v3)
    var v15 float64 = v1
    var v16 int32 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
