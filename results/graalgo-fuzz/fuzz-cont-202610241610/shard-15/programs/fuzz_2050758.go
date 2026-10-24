package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 int64; F2 uint16 }
type GenS3 struct { F0 uint8 }
func genFn0() string {
    var v4 int = 454
    var v5 int32 = -550
    var v6 bool = true
    _ = v4
    _ = v5
    _ = v6
    return "hello"
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]int)
    var v3 int8 = -128
    var v7 float32 = float32(((3.14159 * (0.1 + -2.25)) - ((3.14159 * -2.25) - 0.1)))
    v3 -= (-(v3))
    fmt.Println(v3)
    if (((0 == 2) && (!(true || true))) || (!(true || (true || true)))) {
        v2[32767] = 4294967295
        var v8 int16 = 71
        fmt.Println(v3)
        _ = v8
    } else {
        switch "a" {
        case "foo bar":
            for fi9 := 0; fi9 < 3; fi9++ {
                var v10 uint = 4294967295
                if ((("foo bar" + "hello") <= "hello") || false) {
                    v11 := GenS1{F0: uint32(v10)}
                    _ = v11
                }
                _ = fi9
                _ = v10
            }
        }
        v12 := GenS2{F0: 1, F1: -9223372036854775808, F2: 255}
        v13 := GenS3{F0: 118}
        _ = v12
        _ = v13
    }
    var v14 int8 = (^(v3))
    fmt.Println(v14, v3)
    fmt.Println(v14, v3)
    var v15 int8 = (v3 >> ((217) & 63))
    var v16 float32 = 1e10
    v7 = 0.0
    var v17 float32 = 1.5
    var v18 uint8 = 94
    v18 += v18
    for fi19 := 0; fi19 < 1; fi19++ {
        v14 = (^(v3))
        _ = fi19
    }
    v7 += v17
    var v20 uint16 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v7
}
