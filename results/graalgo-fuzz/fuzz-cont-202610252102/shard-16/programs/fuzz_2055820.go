package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 uint8; F2 uint32 }
func genFn0(p0 int32, p1 int8, p2 uint64) int16 {
    switch "hello" {
    case "foo bar":
        var v7 uint8 = 128
        _ = v7
    }
    var v8 float64 = 1e10
    v9 := []int{9223372036854775807, 0, 32767, 127}
    fmt.Println(p2, v8, p1)
    fmt.Println(p2, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return -1
}

func main() {
    var v1 uint16 = 2
    v2 := make(map[string]int)
    v3 := make(map[string]string)
    var v4 string = "hello"
    var v5 string = ""
    v6 := make(map[int]bool)
    var v10 bool = (true || false)
    var v11 uint = 42
    for fi12 := 0; fi12 < 11; fi12++ {
        fmt.Println(v11)
        _ = fi12
    }
    v4 = v4
    var v13 int = 4294967295
    v13 += int(v11)
    var v14 int64 = 0
    var v15 uint8 = 2
    fmt.Println(v5)
    fmt.Println(v5, v11, v1)
    v5 = "hello"
    for fi16 := 0; fi16 < 10; fi16++ {
        v17 := GenS1{F0: (v11 ^ v11)}
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 1; fi18++ {
        if (v10 || ((-2.25 - 0.0) < ((0.1 + ((1e10 + ((((float32(0.1) + 0.0) - (float32(1e10) / 1.0)) - -2.25) / ((((3.14159 / 1.0) * float32(-2.25)) / (0.1 + 3.14159)) - 3.14159))) - 1.5)) / 1.5))) {
            var v19 uint8 = v15
            _ = v19
        } else {
            var v20 int64 = v14
            delete(v2, (("\n\t" + (("hello" + "foo bar") + v4)) + v4))
            _ = v20
        }
        fmt.Println(v10, v5)
        if v10 {
            fmt.Println(v5, v10, v15)
            var v21 int8 = 11
            _ = v21
        } else {
            var v22 uint64 = 0
            for fi23 := 0; fi23 < 9; fi23++ {
                var v24 uint64 = (v22 << ((v11) & 63))
                _ = fi23
                _ = v24
            }
            var v25 int = int(v14)
            _ = v22
            _ = v25
        }
        _ = fi18
    }
    v26 := GenS2{F0: 0, F1: 164, F2: 2}
    v27 := GenS1{F0: v11}
    v14 -= 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
