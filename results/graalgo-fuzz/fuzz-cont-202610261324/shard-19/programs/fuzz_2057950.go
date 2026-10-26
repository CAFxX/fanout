package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 uint32 }

func main() {
    var v1 bool = ((true && (2 < 1)) || false)
    v2 := make(map[int]bool)
    var v3 uint8 = 0
    var v4 uint = 65535
    v3 -= v3
    v4 -= (v4 >> (((v4 | 42)) & 63))
    var v5 float64 = 0.0
    v5 += (v5 / 3.14159)
    delete(v2, 19)
    var v6 bool = false
    v7 := GenS1{F0: ((("hello" + ("" + "a")) + "\n\t") + "a")}
    v8 := GenS1{F0: "foo bar"}
    v2[128] = v1
    switch 1 {
    case 0:
        if true {
            var v9 int = int(v4)
            v9 += (v9 * v9)
            _ = v9
        }
    default:
        fmt.Println(v4)
    }
    v1 = v1
    switch ((("" + "a") + "\n\t") + "") {
    case "":
        v10 := []int{183, 32768}
        _ = v10
    }
    switch "a" {
    case "foo bar":
        v1 = (0 >= 1)
    case "":
        for fi11 := 0; fi11 < 11; fi11++ {
            v12 := GenS2{F0: ((!true) || v1), F1: 1}
            _ = fi11
            _ = v12
        }
    case "a":
        v1 = v6
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
