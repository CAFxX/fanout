package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 uint32; F2 uint8 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{128, 9223372036854775807}
    var v3 uint16 = 32767
    fmt.Println(v3)
    var v4 uint32 = 2
    for fi5 := 0; fi5 < 12; fi5++ {
        if false {
            _ = len(v2)
            for fi6 := 0; fi6 < 18; fi6++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi6
            }
        }
        var v7 float32 = (-2.25 - 3.14159)
        v8 := GenS1{F0: (v4 | (^(v4)))}
        _ = fi5
        _ = v7
        _ = v8
    }
    switch "foo bar" {
    case "hello":
        v9 := GenS1{F0: 150}
        _ = v9
    default:
        v2 = append(v2, 2147483647)
    }
    fmt.Println(v4, v3)
    var v10 int32 = 42
    _ = len(v1)
    v10 = 309
    var v11 float64 = 1e10
    v12 := GenS2{F0: 32768}
    for fi13 := 0; fi13 < 15; fi13++ {
        v2 = append(v2, (fi13 + 42))
        for fi14 := 0; fi14 < 18; fi14++ {
            v1[(^(v4))] = "a"
            v15 := GenS2{F0: fi13}
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    v16 := GenS3{F0: 42, F1: v4, F2: 53}
    fmt.Println(v10, v4, v4)
    for fi17 := 0; fi17 < 5; fi17++ {
        v1[65535] = ("foo bar" + ("foo bar" + "hello"))
        _ = fi17
    }
    for fi18 := 0; fi18 < 9; fi18++ {
        var v19 int32 = -1
        v19 -= int32(v4)
        fmt.Println(v3, v10, v19)
        _ = fi18
        _ = v19
    }
    var v20 int8 = int8(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
