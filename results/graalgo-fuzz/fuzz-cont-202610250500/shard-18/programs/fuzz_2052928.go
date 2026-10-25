package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = 1
    var v3 uint32 = 2147483648
    v4 := make(map[string]uint32)
    var v5 bool = (127 > -9223372036854775808)
    v3 -= (v3 << ((42) & 63))
    v6 := GenS1{F0: 0}
    var v7 float32 = 0.0
    v2 += 42
    if ("foo bar" == ("" + "\n\t")) {
        if v5 {
            v8 := []int{0, 2147483647, 32767}
            _ = v8
        } else {
            v9 := GenS1{F0: -1}
            _ = v9
        }
        v7 = ((v7 - float32(((1.5 / (1e10 * 1e10)) * 1.5))) - 3.14159)
    }
    var v10 uint8 = 52
    for fi11 := 0; fi11 < 1; fi11++ {
        v12 := []int{(-(fi11))}
        fmt.Println(v2)
        v2 = v2
        _ = fi11
        _ = v12
    }
    switch "hello" {
    case "":
        switch ("" + (("foo bar" + "\n\t") + ("" + ("foo bar" + "foo bar")))) {
        case "hello":
            var v13 uint16 = 0
            _ = v13
        case "foo bar":
            v14 := []int{}
            _ = v14
        case "":
            var v15 uint8 = v10
            _ = v15
        default:
            fmt.Println(v2, v7, v3)
        }
    case "foo bar":
        v2 = v2
    case "a":
        v5 = false
    }
    var v16 int8 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
