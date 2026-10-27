package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }

func main() {
    v1 := []uint32{2, 630, 1, 2147483648}
    var v2 string = "\n\t"
    var v3 uint8 = 114
    var v4 int16 = 1000
    v5 := make(map[uint32]bool)
    var v6 int = -251
    v4 = (^(v4))
    v6 = (-(v6))
    var v7 uint8 = v3
    if (!false) {
        var v8 uint16 = 32767
        delete(v5, 894)
        _ = v8
    } else {
        v9 := GenS1{F0: false, F1: -128}
        _ = v9
    }
    switch (((v2 + v2) + "hello") + (((v2 + v2) + ("hello" + ((("\n\t" + v2) + (("a" + ((((v2 + v2) + ((v2 + ((v2 + "a") + (v2 + v2))) + v2)) + v2) + v2)) + "foo bar")) + v2))) + "\n\t")) {
    case "hello":
        var v10 float64 = (0.0 + (3.14159 + (3.14159 * float64(float32(float64(0.1))))))
        _ = v10
    }
    fmt.Println(v6, v2, v7)
    var v11 int = 399
    v3 = 127
    v4 += -926
    v5[915] = (!false)
    v12 := GenS1{F0: (0 <= (v11 | 356)), F1: -97}
    if len(v1) > 0 {
        v1[0] = 581
    }
    var v13 uint = 42
    for fi14 := 0; fi14 < 1; fi14++ {
        fmt.Println(v13, v4, v6)
        _ = fi14
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
