package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 int8) int16 {
    switch 1 {
    case 0:
        fmt.Println(p0)
    }
    var v7 uint32 = 2147483648
    v8 := []int{4294967295, 4294967295, -712}
    _ = p0
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint = 65535
    v3 := []string{"a"}
    v4 := make(map[int]bool)
    var v5 uint64 = 720
    var v6 int32 = 821
    switch -1 {
    case 1:
        fmt.Println(v2, v6)
    case -80:
        fmt.Println(v2, v6, v2)
    }
    fmt.Println(v2, v2, v5)
    v2 = (v2 % (((v2 - (^(v2)))) | 1))
    var v9 float32 = (3.14159 / -2.25)
    var v10 int8 = -1
    _ = len(v4)
    switch "hello" {
    case "foo bar":
        v11 := GenS1{F0: v10}
        _ = v11
    case "a":
        v9 = (v9 / v9)
    case "\n\t":
        for fi12 := 0; fi12 < 1; fi12++ {
            switch "foo bar" {
            case "":
                fmt.Println(v10, v9)
            default:
                var v13 uint = (v2 % ((255) | 1))
                _ = v13
            }
            _ = fi12
        }
    }
    var v14 float32 = (v9 - 1e10)
    var v15 int8 = v10
    v6 = v6
    v10 = 1
    var v16 int64 = 9223372036854775807
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
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
