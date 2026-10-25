package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int64 }

func main() {
    var v1 float32 = (0.0 * ((3.14159 + 0.1) * (((1.5 - 1.5) - ((1.5 - 0.1) / 1.0)) - ((1e10 + float32(3.14159)) / 1.5))))
    var v2 uint8 = 23
    var v3 uint = 4294967295
    v4 := make(map[uint32]string)
    v5 := make(map[uint32]string)
    var v6 int64 = -137
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := GenS1{F0: int16(fi7), F1: ("hello" + (("hello" + "a") + "")), F2: (v6 - (v6 + 0))}
        var v9 float32 = v1
        v3 = v3
        _ = fi7
        _ = v8
        _ = v9
    }
    fmt.Println(v3, v6, v6)
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 uint = 255
        _ = fi10
        _ = v11
    }
    var v12 uint16 = 32767
    v3 -= v3
    var v13 int32 = -484
    fmt.Println(v13)
    v2 += v2
    v1 -= float32(0.0)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
