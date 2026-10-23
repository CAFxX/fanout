package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 uint64; F1 int64; F2 string }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]string)
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]uint32)
    var v5 float64 = (0.0 / 0.1)
    v6 := GenS1{F0: 0, F1: 2}
    v5 += 3.14159
    var v7 int8 = -15
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 uint16 = 1
        _ = fi8
        _ = v9
    }
    if true {
        for fi10 := 0; fi10 < 1; fi10++ {
            v7 -= (v7 * -81)
            fmt.Println(v7, fi10, fi10)
            v7 = v7
            _ = fi10
        }
    }
    var v11 int32 = 42
    if (0 >= 4294967295) {
        var v12 uint32 = 2147483648
        if false {
            if (2147483647 >= 2147483647) {
                var v13 int64 = 42
                _ = v13
            } else {
                fmt.Println(v5, v11)
                var v14 string = ""
                v15 := GenS1{F0: 85, F1: 65535}
                _ = v14
                _ = v15
            }
            var v16 string = (("hello" + "") + ("hello" + "a"))
            _ = v16
        }
        v11 = v11
        _ = v12
    } else {
        v7 = v7
    }
    var v17 string = "foo bar"
    v18 := GenS2{F0: 2, F1: 858, F2: "a"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
