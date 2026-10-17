package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }

func main() {
    var v1 int64 = 1
    var v2 uint = 255
    v3 := make(map[int]string)
    v4 := make(map[int]string)
    var v5 uint32 = uint32(v2)
    for fi6 := 0; fi6 < 8; fi6++ {
        fmt.Println(v2)
        v7 := []int{(fi6 | fi6)}
        _ = fi6
        _ = v7
    }
    v5 = 1
    v2 = v2
    var v8 int8 = -114
    if (42 > 42) {
        fmt.Println(v2, v1, v1)
        v9 := GenS1{F0: -2147483648, F1: (v8 << ((v2) & 63))}
        _ = v9
    } else {
        var v10 uint64 = 4294967295
        switch 2147483647 {
        case -1:
            var v11 float64 = 1e10
            _ = v11
        case -985:
            var v12 int = 32768
            _ = v12
        case 831:
            fmt.Println(v5)
        default:
            var v13 float64 = 1e10
            _ = v13
        }
        _ = v10
    }
    delete(v3, 630)
    if (1 <= (v5 * 79)) {
        v3[128] = "hello"
    }
    if ((3.14159 <= 1e10) || true) {
        v14 := []int{2147483647, 0, 4294967295, int(v8)}
        _ = v14
    } else {
        switch -147 {
        case 0:
            switch 127 {
            case 1:
                var v15 bool = (false || ((v1 ^ v1) != v1))
                _ = v15
            }
        default:
            switch ("" + "") {
            case "":
                var v16 uint16 = 1
                _ = v16
            case "a":
                var v17 uint32 = 4294967295
                _ = v17
            default:
                var v18 string = ""
                _ = v18
            }
        }
        v19 := []int{42, 1, -748, -2147483648}
        var v20 string = ""
        _ = v19
        _ = v20
    }
    var v21 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
