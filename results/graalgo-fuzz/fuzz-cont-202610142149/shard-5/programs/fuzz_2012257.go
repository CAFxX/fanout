package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 string }
type GenS2 struct { F0 uint16 }
func genFn0() uint8 {
    v6 := []int{-9223372036854775808, 42, 1, -845}
    v7 := GenS1{F0: 1, F1: 172, F2: "a"}
    var v8 float32 = -2.25
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v8 += v8
    _ = v6
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := []int{-74, 1}
    v2 := make(map[int]int)
    var v3 int = -1
    var v4 int16 = -1
    v5 := []uint32{}
    var v9 int32 = 564
    var v10 uint = 42
    switch ("\n\t" + ("hello" + "")) {
    case "a":
        var v11 string = "hello"
        _ = v11
    case "hello":
        var v12 float32 = 3.14159
        _ = v12
    default:
        v9 += 0
    }
    if true {
        v13 := GenS2{F0: uint16(v10)}
        fmt.Println(v9, v9, v9)
        if false {
            var v14 bool = true
            _ = v14
        }
        _ = v13
    } else {
        if len(v5) > 0 {
            v5[0] = 42
        }
        for fi15 := 0; fi15 < 17; fi15++ {
            var v16 uint = (v10 ^ 255)
            if (!(!true)) {
                v3 = -291
            } else {
                v3 = (-(v3))
            }
            _ = fi15
            _ = v16
        }
    }
    var v17 uint = (^(v10))
    v9 += v9
    if len(v1) > 0 {
        v1[0] = 921
    }
    v18 := GenS2{F0: 356}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
