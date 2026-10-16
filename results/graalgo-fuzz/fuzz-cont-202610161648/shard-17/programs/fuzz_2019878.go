package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 int16; F2 uint16 }

func main() {
    v1 := []int{550, 32767, 42}
    v2 := []string{("hello" + "hello")}
    v3 := make(map[uint32]string)
    var v4 int32 = -550
    v5 := GenS1{F0: false}
    v4 = v4
    var v6 float64 = float64((float32(float64(3.14159)) - -2.25))
    v6 -= 0.0
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 uint16 = 65535
        _ = fi7
        _ = v8
    }
    switch "" {
    case "hello":
        v6 = float64((float32(1.5) - float32(float64(1e10))))
    case "a":
        v4 = (v4 & (v4 * 1))
    }
    fmt.Println(v4)
    var v9 uint = 65535
    for fi10 := 0; fi10 < 15; fi10++ {
        var v11 uint32 = 658
        for fi12 := 0; fi12 < 6; fi12++ {
            if false {
                var v13 float64 = float64(((1.5 * 1e10) + 1e10))
                _ = v13
            } else {
                if ((1e10 + 1.5) < 3.14159) {
                    _ = len(v3)
                }
            }
            _ = fi12
        }
        v14 := GenS2{F0: 538, F1: 1000, F2: 65535}
        _ = fi10
        _ = v11
        _ = v14
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
