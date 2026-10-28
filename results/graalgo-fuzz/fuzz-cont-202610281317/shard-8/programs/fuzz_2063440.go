package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 uint64, p1 int8, p2 int32) uint8 {
    v5 := GenS1{F0: "a"}
    if true {
        switch 32767 {
        case 9223372036854775807:
            p1 = p1
        default:
            for fi6 := 0; fi6 < 3; fi6++ {
                p0 += (p0 << ((65535) & 63))
                _ = fi6
            }
        }
        var v7 int64 = 42
        v7 = 428
        _ = v7
    }
    v8 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v8
    return 255
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]string)
    var v3 bool = false
    v4 := []int{0, -330}
    fmt.Println(v3)
    v4 = append(v4, -2147483648)
    var v9 int = -253
    v10 := GenS1{F0: ""}
    var v11 uint = 42
    v4 = append(v4, v9)
    var v12 int16 = -1
    v13 := GenS1{F0: "\n\t"}
    var v14 int16 = -32768
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
