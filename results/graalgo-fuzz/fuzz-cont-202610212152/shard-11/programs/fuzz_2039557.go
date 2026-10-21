package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 uint8 = 104
    var v2 int64 = 0
    var v3 int16 = -593
    if true {
        v3 += -1
        fmt.Println(v1, v1)
        fmt.Println(v2, v3, v1)
    }
    fmt.Println(v3, v3)
    var v4 float64 = 1.5
    for fi5 := 0; fi5 < 10; fi5++ {
        switch int(v1) {
        case -298:
            fmt.Println(v2, v1)
        }
        v1 -= v1
        v6 := GenS1{F0: 985}
        _ = fi5
        _ = v6
    }
    v7 := []int{520}
    v3 = v3
    v8 := GenS1{F0: 563}
    var v9 int = 9223372036854775807
    v10 := GenS1{F0: 15}
    _ = len(v7)
    _ = len(v7)
    switch (("foo bar" + "hello") + "foo bar") {
    case "a":
        for fi11 := 0; fi11 < 9; fi11++ {
            var v12 float32 = 3.14159
            _ = fi11
            _ = v12
        }
    case "foo bar":
        v2 += (v2 ^ -550)
    default:
        fmt.Println(v4)
    }
    v9 = (v9 >> ((159) & 63))
    var v13 float32 = 3.14159
    for fi14 := 0; fi14 < 4; fi14++ {
        var v15 uint64 = 4294967295
        _ = fi14
        _ = v15
    }
    v1 = (v1 * (v1 * v1))
    v1 -= (v1 | (v1 << ((18446744073709551615) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
