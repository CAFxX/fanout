package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 int8, p1 int, p2 uint32) int16 {
    if true {
        p2 = (^(p2))
        v7 := []int{p1, p1, 128}
        _ = v7
    }
    var v8 uint16 = 65535
    if (1.5 == -2.25) {
        v9 := []int{(p1 << ((4294967295) & 63))}
        _ = v9
    }
    var v10 uint64 = 2
    var v11 int8 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v8
    return -919
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]string)
    var v3 uint8 = 1
    var v4 int16 = 0
    var v5 uint64 = 42
    var v6 bool = (307 >= 42)
    v12 := GenS1{F0: 2147483647}
    var v13 uint16 = 65535
    for fi14 := 0; fi14 < 12; fi14++ {
        v15 := GenS1{F0: (fi14 % ((fi14) | 1))}
        v16 := GenS1{F0: (fi14 ^ fi14)}
        v6 = (302 > 255)
        _ = fi14
        _ = v15
        _ = v16
    }
    v3 = (v3 % ((v3) | 1))
    switch "hello" {
    case "hello":
        switch ("\n\t" + "foo bar") {
        case "hello":
            v6 = true
        case "":
            var v17 string = (("foo bar" + "") + "foo bar")
            _ = v17
        case "\n\t":
            var v18 int64 = 0
            _ = v18
        }
    }
    if false {
        v13 = (v13 >> ((2) & 63))
        var v19 string = "foo bar"
        _ = v19
    } else {
        fmt.Println(v13, v13, v5)
        var v20 uint8 = (v3 + (^(v3)))
        _ = v20
    }
    v21 := v1[-743]
    var v22 float32 = (((1.5 * 3.14159) + 3.14159) / (((((3.14159 - float32(float64(0.0))) / float32((-2.25 / float64((1e10 - (float32(float64(0.0)) / (0.1 * 3.14159))))))) + -2.25) + 1.5) * (((float32(-2.25) + float32(1.5)) - (0.1 + 1e10)) * 0.1)))
    var v23 float64 = (1e10 * 3.14159)
    v1[2147483647] = (!true)
    v6 = (!v6)
    var v24 int64 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
