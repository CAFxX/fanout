package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }

func main() {
    v1 := make(map[int]bool)
    var v2 float64 = (((float64(1.5) / -2.25) + float64(1e10)) * 0.1)
    v3 := make(map[uint32]int)
    var v4 uint64 = 49
    switch (((("a" + "hello") + "") + ("hello" + "")) + "a") {
    case "a":
        fmt.Println(v4, v4)
    case "":
        fmt.Println(v2, v4)
    default:
        switch 32768 {
        case 4294967295:
            if (0.0 >= (float32(v2) / 3.14159)) {
                var v5 uint32 = 0
                _ = v5
            }
        case 0:
            v6 := GenS1{F0: "hello", F1: (v4 * (^(v4)))}
            _ = v6
        case 1:
            v7 := GenS1{F0: ("" + "foo bar"), F1: v4}
            _ = v7
        default:
            var v8 uint16 = 255
            _ = v8
        }
    }
    var v9 bool = false
    v2 += (1e10 / v2)
    var v10 uint64 = 1
    if v9 {
        var v11 uint8 = 255
        fmt.Println(v4)
        fmt.Println(v10, v9)
        _ = v11
    } else {
        var v12 string = ("\n\t" + "hello")
        _ = v12
    }
    var v13 float32 = (-2.25 / -2.25)
    for fi14 := 0; fi14 < 13; fi14++ {
        v4 += 4294967295
        v15 := GenS1{F0: "", F1: (^(v4))}
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        v3[58] = (fi16 + fi16)
        _ = fi16
    }
    v17 := GenS1{F0: "hello", F1: (v10 * v10)}
    for fi18 := 0; fi18 < 8; fi18++ {
        v9 = (881 >= 42)
        _ = fi18
    }
    v10 = (v4 % (((v10 >> ((135) & 63))) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
