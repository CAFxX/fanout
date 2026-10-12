package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
func genFn0(p0 uint16, p1 int32) uint32 {
    v6 := []int{-9223372036854775808}
    p1 = p1
    if len(v6) > 0 {
        v6[0] = 559
    }
    _ = p0
    _ = p1
    _ = v6
    return 0
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]int)
    v3 := make(map[int]string)
    var v4 float32 = 3.14159
    var v5 uint64 = 42
    fmt.Println(v4, v4)
    if (2147483648 < 2147483647) {
        v7 := []int{32767, 4294967295, 0}
        _ = v7
    } else {
        var v8 float64 = 1e10
        switch ("foo bar" + "hello") {
        case "\n\t":
            for fi9 := 0; fi9 < 8; fi9++ {
                v10 := GenS1{F0: "foo bar", F1: true}
                _ = fi9
                _ = v10
            }
        }
        _ = v8
    }
    fmt.Println(v5, v4)
    switch 9223372036854775807 {
    case 127:
        v11 := v1[991]
        _ = v11
    case -343:
        v12 := []int{32767}
        _ = v12
    }
    var v13 uint = 1
    for fi14 := 0; fi14 < 7; fi14++ {
        var v15 int64 = 0
        _ = fi14
        _ = v15
    }
    var v16 int8 = 0
    var v17 uint16 = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
