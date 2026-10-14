package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 uint8, p1 int8, p2 string) int16 {
    for fi7 := 0; fi7 < 2; fi7++ {
        switch fi7 {
        case 127:
            p0 = (^(p0))
        case -906:
            v8 := []int{32767, -2147483648}
            _ = v8
        default:
            fmt.Println(p0, fi7, p2)
        }
        p2 = p2
        _ = fi7
    }
    p0 -= (^(p0))
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 int16 = 0
    v2 := []int{}
    var v3 uint16 = 2
    var v4 int16 = 1
    var v5 int64 = 1
    v6 := make(map[string]uint32)
    _ = len(v2)
    v9 := GenS1{F0: 42}
    _ = len(v6)
    delete(v6, ("a" + "foo bar"))
    fmt.Println(v1, v5)
    switch "" {
    case "a":
        var v10 int32 = -743
        _ = v10
    case "hello":
        switch -2147483648 {
        case 1:
            v11 := v6["hello"]
            _ = v11
        }
    case "\n\t":
        switch "foo bar" {
        case "\n\t":
            v3 = (v3 >> ((0) & 63))
        case "hello":
            fmt.Println(v3, v1)
        case "":
            v12 := GenS1{F0: 844}
            _ = v12
        default:
            v13 := GenS1{F0: 42}
            _ = v13
        }
    }
    var v14 uint8 = 127
    var v15 string = ""
    var v16 uint64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
