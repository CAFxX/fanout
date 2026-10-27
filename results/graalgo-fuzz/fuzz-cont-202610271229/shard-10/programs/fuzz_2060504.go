package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }

func main() {
    v1 := make(map[uint32]int)
    var v2 float64 = 0.0
    var v3 int32 = 0
    var v4 uint = 65535
    v5 := []uint32{911}
    v6 := []string{"hello", ("" + "\n\t")}
    var v7 int16 = 0
    if false {
        fmt.Println(v4)
        fmt.Println(v2)
    } else {
        v2 = (v2 - (3.14159 / -2.25))
    }
    v7 = (v7 - v7)
    fmt.Println(v4, v2, v7)
    v8 := GenS1{F0: 42, F1: 0}
    if (127 <= 0) {
        var v9 int64 = 0
        _ = v9
    }
    v5 = append(v5, uint32(v4))
    var v10 bool = ((v7 - v7) >= v7)
    v2 -= 0.0
    v10 = (!(!(false || (("\n\t" + ("" + "a")) <= ((("hello" + "") + "") + ("hello" + ("foo bar" + "\n\t")))))))
    var v11 int8 = 70
    var v12 uint8 = 0
    var v13 string = "hello"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
