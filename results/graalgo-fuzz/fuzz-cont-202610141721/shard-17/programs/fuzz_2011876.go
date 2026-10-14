package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0() bool {
    v7 := []int{182}
    if (!(42 < 18446744073709551615)) {
        for fi8 := 0; fi8 < 3; fi8++ {
            switch ("hello" + "hello") {
            case "":
                fmt.Println(fi8)
            default:
                var v9 uint = 42
                _ = v9
            }
            fmt.Println(fi8)
            _ = fi8
        }
    }
    if ("a" == "a") {
        var v10 float64 = 1.5
        _ = v10
    } else {
        v11 := GenS1{F0: 556}
        _ = v11
    }
    _ = v7
    return (0.1 < (3.14159 * 3.14159))
}

func main() {
    var v1 int16 = 1
    v2 := []uint32{1, 0, 206}
    var v3 uint16 = 65535
    v4 := make(map[uint32]string)
    v5 := make(map[int]uint32)
    v6 := make(map[string]uint32)
    for fi12 := 0; fi12 < 16; fi12++ {
        v3 = (v3 >> ((65535) & 63))
        _ = fi12
    }
    v13 := GenS1{F0: 240}
    var v14 uint = 305
    var v15 uint = v14
    var v16 uint8 = 1
    var v17 float32 = 1e10
    var v18 int = 42
    var v19 uint16 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
