package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 uint8; F2 uint16 }

func main() {
    var v1 uint8 = 2
    v2 := make(map[string]int)
    var v3 bool = false
    var v4 string = ""
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 string = v4
        v3 = v3
        _ = fi5
        _ = v6
    }
    v4 = ("hello" + v4)
    var v7 int32 = -2147483648
    if (-2.25 == ((float32(-2.25) / 3.14159) + 3.14159)) {
        for fi8 := 0; fi8 < 15; fi8++ {
            v7 = 42
            for fi9 := 0; fi9 < 13; fi9++ {
                var v10 uint32 = 2147483647
                v11 := []int{}
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = fi8
        }
        var v12 int32 = (v7 >> ((42) & 63))
        v7 += (-(v7))
        _ = v12
    }
    v7 = (v7 & v7)
    for fi13 := 0; fi13 < 11; fi13++ {
        v2[""] = (fi13 >> ((255) & 63))
        for fi14 := 0; fi14 < 20; fi14++ {
            v15 := GenS1{F0: 0, F1: 1}
            var v16 uint64 = 4294967295
            _ = fi14
            _ = v15
            _ = v16
        }
        var v17 float32 = (float32(1.5) / (0.0 + ((float32(((((3.14159 - 0.1) - -2.25) * 3.14159) + 1.5)) * -2.25) * 1.5)))
        _ = fi13
        _ = v17
    }
    var v18 uint = 1
    fmt.Println(v7, v1)
    v19 := GenS2{F0: (891 != 516)}
    var v20 int32 = (v7 | v7)
    v21 := GenS3{F0: -124, F1: v1, F2: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v7
}
