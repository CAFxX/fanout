package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
func genFn0(p0 string, p1 int64, p2 int64) int16 {
    fmt.Println(p2)
    var v7 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -1
}
func genFn1(p0 uint64, p1 int) int32 {
    v8 := GenS1{F0: -9223372036854775808, F1: 0}
    v9 := []int{(p1 - 9223372036854775807), p1, -9223372036854775808}
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return -380
}

func main() {
    var v1 string = ""
    var v2 int32 = -819
    var v3 string = ("foo bar" + "")
    v4 := []uint32{2}
    v5 := make(map[uint32]int)
    var v6 string = v3
    var v10 uint8 = 1
    v3 = "a"
    var v11 int16 = 32767
    delete(v5, 2)
    var v12 bool = true
    v3 = "a"
    v5[uint32(v10)] = 1
    switch v3 {
    case "a":
        var v13 uint16 = 824
        _ = v13
    case "foo bar":
        if true {
            switch -1 {
            case 9223372036854775807:
                switch 9223372036854775807 {
                case 1:
                    v6 = v6
                }
            case -54:
                var v14 uint8 = (v10 + v10)
                _ = v14
            }
        }
    default:
        fmt.Println(v1)
    }
    fmt.Println(v10, v11, v2)
    var v15 uint8 = v10
    fmt.Println(v6, v3)
    var v16 uint16 = 0
    var v17 uint64 = 724
    var v18 uint32 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
