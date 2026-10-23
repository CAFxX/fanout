package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 uint, p1 int16, p2 uint16) int8 {
    fmt.Println(p1)
    var v7 bool = ((1 == 0) || true)
    v8 := GenS1{F0: p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -1
}

func main() {
    var v1 int8 = -128
    var v2 uint16 = 1
    var v3 int8 = (v1 / ((v1) | 1))
    v4 := make(map[uint32]bool)
    var v5 uint16 = 32767
    var v6 float32 = 1.5
    switch ("hello" + ("foo bar" + "a")) {
    case "foo bar":
        v1 -= v1
    case "\n\t":
        if true {
            var v9 int32 = -449
            _ = v9
        } else {
            var v10 uint8 = 128
            v2 = (^(v5))
            _ = v10
        }
    case "a":
        var v11 int64 = 0
        _ = v11
    }
    var v12 int16 = -32768
    fmt.Println(v3, v5, v6)
    switch "\n\t" {
    case "hello":
        if (true || ("" > "foo bar")) {
            v13 := []int{547, int(v5), -9223372036854775808, 127}
            _ = v13
        }
    case "\n\t":
        fmt.Println(v3, v1, v5)
    default:
        fmt.Println(v6)
    }
    v5 = v5
    delete(v4, 518)
    switch "a" {
    case "":
        fmt.Println(v2)
    }
    v5 = 42
    switch -9223372036854775808 {
    case -1:
        var v14 uint = 722
        _ = v14
    case 42:
        var v15 int64 = 201
        _ = v15
    }
    var v16 int16 = v12
    var v17 int8 = v3
    v4[0] = false
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
