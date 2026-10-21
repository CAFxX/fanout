package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 uint8 = 42
    var v2 int16 = 984
    v3 := make(map[string]uint32)
    v4 := make(map[int]bool)
    var v5 string = "foo bar"
    v6 := GenS1{F0: 192, F1: 42, F2: 1}
    var v7 uint8 = v1
    v8 := GenS2{F0: (v1 % ((v7) | 1)), F1: 127}
    switch v5 {
    case "hello":
        if (!(!((-2147483648 <= -1) || (585 < 1)))) {
            v1 = 215
            for fi9 := 0; fi9 < 16; fi9++ {
                _ = len(v4)
                v7 = v7
                v1 -= v1
                _ = fi9
            }
            v7 = v7
        }
    case "a":
        v7 += v7
    case "":
        v10 := []int{864, 4294967295, -9223372036854775808}
        _ = v10
    }
    v11 := GenS1{F0: 42, F1: 2, F2: 406}
    var v12 uint32 = 1
    v1 = (v7 << ((451) & 63))
    for fi13 := 0; fi13 < 1; fi13++ {
        v14 := GenS3{F0: v7}
        _ = fi13
        _ = v14
    }
    v12 = (v12 & v12)
    var v15 bool = false
    fmt.Println(v5, v5)
    var v16 uint = 1
    var v17 string = (("a" + v5) + "foo bar")
    for fi18 := 0; fi18 < 1; fi18++ {
        var v19 float64 = (0.1 * 1.5)
        for fi20 := 0; fi20 < 5; fi20++ {
            if v15 {
                if false {
                    fmt.Println(v17, v19)
                    v19 = v19
                    var v21 uint8 = 0
                    _ = v21
                }
            } else {
                v3["foo bar"] = v12
                v15 = v15
                v7 = v1
            }
            _ = fi20
        }
        _ = fi18
        _ = v19
    }
    var v22 int64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
