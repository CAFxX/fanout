package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]string)
    var v3 uint8 = 2
    v4 := make(map[int]int)
    var v5 bool = ("hello" == "")
    if true {
        for fi6 := 0; fi6 < 19; fi6++ {
            v5 = (-946 <= -214)
            if false {
                var v7 uint16 = 1
                v5 = ((v3 <= v3) && (uint(v7) == 174))
                var v8 float32 = 0.0
                _ = v7
                _ = v8
            }
            var v9 bool = true
            _ = fi6
            _ = v9
        }
        v5 = false
    } else {
        fmt.Println(v3, v5)
        var v10 int8 = 40
        _ = v10
    }
    v5 = v5
    v3 += (v3 + v3)
    var v11 uint8 = (^(v3))
    if (255 > 23) {
        v5 = true
        _ = len(v4)
    }
    var v12 uint16 = 65535
    var v13 int16 = 32767
    v14 := GenS1{F0: 4294967295, F1: 0}
    if (v5 || (42 <= v12)) {
        fmt.Println(v13, v13)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
