package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 uint16 }
func genFn0() int {
    var v5 string = "a"
    var v6 uint = 18446744073709551615
    v7 := GenS1{F0: 9223372036854775807}
    var v8 int16 = 0
    var v9 string = v5
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 int = -216
    var v2 int32 = 2147483647
    v3 := make(map[uint32]int)
    v4 := make(map[int]int)
    v2 += (v2 << ((674) & 63))
    if true {
        switch v1 {
        case 4294967295:
            delete(v4, (v1 | genFn0()))
        case 2147483647:
            for fi10 := 0; fi10 < 4; fi10++ {
                var v11 uint32 = 579
                _ = fi10
                _ = v11
            }
        case 127:
            v12 := GenS1{F0: -1}
            _ = v12
        }
        v13 := GenS2{F0: 65535, F1: 2}
        fmt.Println(v2, v1)
        _ = v13
    }
    var v14 int32 = 115
    fmt.Println(v14)
    switch -650 {
    case 32768:
        switch "hello" {
        case "\n\t":
            var v15 uint8 = 0
            _ = v15
        case "hello":
            if (65535 <= 65535) {
                _ = len(v4)
                var v16 uint16 = 1
                _ = v16
            }
        case "foo bar":
            v2 = v14
        }
    }
    if false {
        v17 := []int{(^(v1)), (v1 * (v1 << ((255) & 63))), (v1 | -275), -2147483648}
        _ = v17
    }
    var v18 int = v1
    v14 -= v2
    v1 = v18
    var v19 uint16 = 620
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
