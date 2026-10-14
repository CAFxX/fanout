package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 string; F2 bool }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 int16; F1 int16; F2 int64 }

func main() {
    var v1 int16 = -1
    v2 := make(map[uint32]uint32)
    var v3 float64 = 0.0
    var v4 int8 = 127
    var v5 uint64 = 715
    var v6 uint8 = 178
    var v7 uint64 = v5
    if (503 >= (^(v7))) {
        if (-100 != (v4 / ((v4) | 1))) {
            var v8 float32 = 3.14159
            var v9 bool = (false || false)
            v1 += (v1 >> ((59) & 63))
            _ = v8
            _ = v9
        }
    }
    if false {
        var v10 float64 = 1e10
        if false {
            var v11 int = int(v4)
            _ = v11
        } else {
            switch "\n\t" {
            case "hello":
                var v12 uint16 = 173
                _ = v12
            case "a":
                for fi13 := 0; fi13 < 13; fi13++ {
                    v14 := GenS1{F0: 1}
                    if true {
                        v5 = (^(v7))
                        v7 = v7
                        fmt.Println(v3)
                    }
                    v15 := GenS2{F0: 1000, F1: ("hello" + "\n\t"), F2: false}
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            case "foo bar":
                var v16 int64 = 2
                _ = v16
            }
        }
        _ = v10
    }
    if ((0.0 / 1.0) <= (((v3 / ((((float64(1.5) + 1.5) * v3) - (3.14159 - v3)) - v3)) + v3) + float64(float32(v3)))) {
        v3 -= v3
        var v17 string = ("" + ("foo bar" + ""))
        _ = v17
    }
    var v18 float32 = (1.5 - 1.5)
    v19 := GenS3{F0: 42, F1: (v1 * v1)}
    var v20 uint32 = 0
    for fi21 := 0; fi21 < 6; fi21++ {
        var v22 int32 = int32(v20)
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 18; fi23++ {
        if (-786 < 9223372036854775807) {
            _ = len(v2)
            v24 := GenS4{F0: 1, F1: v1, F2: int64(v4)}
            var v25 uint8 = v6
            _ = v24
            _ = v25
        }
        var v26 string = "a"
        var v27 string = v26
        _ = fi23
        _ = v26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
