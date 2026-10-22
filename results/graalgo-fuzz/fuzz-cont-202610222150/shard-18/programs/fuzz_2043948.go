package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 bool = false
    var v2 uint64 = 1
    var v3 float32 = (float32(0.1) - 0.0)
    v4 := []int{-9223372036854775808, 32767, 0}
    v2 += v2
    v4 = append(v4, 537)
    if len(v4) > 0 {
        v4[0] = 32767
    }
    switch (("\n\t" + "\n\t") + (("a" + "") + "\n\t")) {
    case "hello":
        v1 = (!(!((0 != 401) && false)))
    case "a":
        for fi5 := 0; fi5 < 20; fi5++ {
            v2 = 2
            _ = fi5
        }
    case "":
        v3 -= (float32(-2.25) + v3)
    }
    var v6 int = 92
    if len(v4) > 0 {
        v4[0] = (v6 - 32767)
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 float32 = (0.0 * ((v3 + 0.1) / v3))
        _ = fi7
        _ = v8
    }
    v2 += (v2 % ((v2) | 1))
    var v9 float32 = v3
    if len(v4) > 0 {
        v4[0] = (v6 / ((v6) | 1))
    }
    v3 = ((3.14159 + 1.5) + v9)
    for fi10 := 0; fi10 < 14; fi10++ {
        v11 := GenS1{F0: (fi10 + (v6 << ((435) & 63)))}
        fmt.Println(v3, v1, fi10)
        v12 := GenS1{F0: (fi10 * v6)}
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 float32 = v3
    v1 = ((v6 < (v6 >> ((29) & 63))) || (127 < 108))
    fmt.Println(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v9
}
