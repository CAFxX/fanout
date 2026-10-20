package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint32, p1 int16, p2 int) uint8 {
    if false {
        fmt.Println(p2, p1, p1)
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    return 37
}

func main() {
    var v1 int16 = -32768
    var v2 uint16 = 423
    v3 := make(map[uint32]string)
    v4 := GenS1{F0: 42, F1: 255}
    v2 -= 255
    if (false && (!true)) {
        fmt.Println(v1, v2)
    }
    for fi5 := 0; fi5 < 11; fi5++ {
        var v6 string = "\n\t"
        for fi7 := 0; fi7 < 6; fi7++ {
            if (!true) {
                v8 := GenS1{F0: 4294967295, F1: 0}
                v1 -= 1
                var v9 int8 = 100
                v9 = v9
                _ = v8
                _ = v9
            } else {
                _ = len(v3)
            }
            _ = fi7
        }
        var v10 uint = 527
        _ = fi5
        _ = v10
        _ = v6
    }
    var v11 uint16 = (v2 << ((283) & 63))
    for fi12 := 0; fi12 < 16; fi12++ {
        fmt.Println(fi12)
        _ = fi12
    }
    v13 := GenS2{F0: -2147483648}
    var v14 int16 = v1
    delete(v3, 0)
    _ = len(v3)
    fmt.Println(v1, v11)
    var v15 uint = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
