package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 uint16 }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint64 }

func main() {
    var v1 string = "hello"
    v2 := make(map[string]uint32)
    var v3 uint = 0
    fmt.Println(v3, v3)
    switch "hello" {
    case "\n\t":
        v1 = "foo bar"
    default:
        v1 = ""
    }
    var v4 float64 = (1.5 / (-2.25 / -2.25))
    if true {
        var v5 uint16 = 42
        _ = v5
    }
    v6 := v2[(v1 + (v1 + ((("foo bar" + "foo bar") + (v1 + v1)) + "")))]
    v3 += v3
    if true {
        v7 := []int{2147483647, 9223372036854775807}
        v8 := GenS1{F0: -9223372036854775808, F1: 32767, F2: 2}
        _ = v7
        _ = v8
    } else {
        v9 := v2[v1]
        _ = v9
    }
    v10 := GenS1{F0: 2147483647, F1: -108, F2: 255}
    v11 := GenS2{F0: 2, F1: (v6 << ((65535) & 63)), F2: 42}
    for fi12 := 0; fi12 < 6; fi12++ {
        var v13 uint = v3
        _ = fi12
        _ = v13
    }
    var v14 uint8 = uint8(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
