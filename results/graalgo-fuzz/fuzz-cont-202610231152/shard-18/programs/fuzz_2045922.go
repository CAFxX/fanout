package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 uint16, p1 bool, p2 int) uint64 {
    switch p2 {
    case -823:
        v5 := []int{(p2 | (p2 / ((p2) | 1))), p2, 9223372036854775807, 1}
        _ = v5
    case -83:
        for fi6 := 0; fi6 < 14; fi6++ {
            var v7 uint8 = 42
            _ = fi6
            _ = v7
        }
    case 128:
        var v8 int = -856
        _ = v8
    }
    var v9 float64 = 1.5
    p2 = -415
    var v10 int = (-(p2))
    for fi11 := 0; fi11 < 16; fi11++ {
        var v12 uint64 = 109
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 932
}

func main() {
    v1 := []int{32768}
    var v2 float32 = 1e10
    v3 := []uint32{42, 0, 2147483648}
    v4 := make(map[string]bool)
    var v13 int8 = 100
    var v14 float32 = 3.14159
    fmt.Println(v2, v2)
    var v15 string = "foo bar"
    for fi16 := 0; fi16 < 9; fi16++ {
        switch v15 {
        case "hello":
            for fi17 := 0; fi17 < 13; fi17++ {
                var v18 string = "hello"
                _ = fi17
                _ = v18
            }
        default:
            if (true || true) {
                var v19 int64 = -836
                _ = v19
            }
        }
        v20 := GenS1{F0: 2}
        _ = fi16
        _ = v20
    }
    v21 := GenS1{F0: 403}
    v1 = append(v1, 0)
    var v22 uint8 = 1
    switch "a" {
    case "":
        v23 := v4["hello"]
        _ = v23
    case "hello":
        fmt.Println(v2)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
