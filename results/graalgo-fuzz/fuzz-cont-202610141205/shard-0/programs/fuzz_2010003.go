package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 uint16 = 619
    var v2 uint8 = 128
    v3 := make(map[int]string)
    var v4 uint8 = 255
    var v5 uint64 = 1
    var v6 uint64 = 0
    var v7 int = 1
    var v8 uint32 = 176
    var v9 uint8 = v2
    v9 -= (v4 + (v4 / ((uint8(v8)) | 1)))
    _ = len(v3)
    for fi10 := 0; fi10 < 9; fi10++ {
        v11 := []int{fi10, (fi10 ^ (-(v7))), -2147483648, fi10}
        var v12 uint8 = v4
        _ = fi10
        _ = v11
        _ = v12
    }
    v3[(v7 & v7)] = "foo bar"
    for fi13 := 0; fi13 < 18; fi13++ {
        v14 := GenS1{F0: 65535}
        switch "" {
        case "hello":
            for fi15 := 0; fi15 < 19; fi15++ {
                fmt.Println(v6, v1, fi15)
                v3[880] = "a"
                v16 := []int{fi13, 1, (v7 | -9223372036854775808)}
                _ = fi15
                _ = v16
            }
        case "":
            v17 := []int{fi13}
            _ = v17
        }
        v3[fi13] = ("" + "")
        _ = fi13
        _ = v14
    }
    var v18 int16 = 42
    fmt.Println(v7, v4, v2)
    v19 := GenS1{F0: 767}
    var v20 uint64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
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
    _ = v8
    _ = v9
}
