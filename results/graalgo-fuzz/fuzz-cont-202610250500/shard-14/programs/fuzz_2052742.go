package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 string; F2 bool }
type GenS3 struct { F0 uint32; F1 int32 }

func main() {
    v1 := make(map[string]string)
    var v2 int = 128
    v3 := make(map[uint32]int)
    v4 := GenS1{F0: 42}
    fmt.Println(v2)
    var v5 int64 = 0
    switch 42 {
    case 4294967295:
        v5 = (v5 | v5)
    case 128:
        v6 := v3[322]
        _ = v6
    default:
        v7 := GenS2{F0: 963, F1: (("" + ("" + ((("" + ("a" + "a")) + ("\n\t" + (("\n\t" + "") + "hello"))) + "\n\t"))) + ""), F2: (!true)}
        _ = v7
    }
    fmt.Println(v2, v2)
    v8 := GenS1{F0: 32767}
    v5 = -9223372036854775808
    var v9 uint32 = 42
    v5 = v5
    var v10 int = (v2 ^ v2)
    v2 += (^(v2))
    for fi11 := 0; fi11 < 16; fi11++ {
        v12 := GenS3{F0: (v9 + (v9 << ((65535) & 63))), F1: 1}
        fmt.Println(v10, v10, v2)
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        v14 := []int{(v10 / ((9223372036854775807) | 1)), fi13}
        v10 = 9223372036854775807
        _ = fi13
        _ = v14
    }
    v2 = v2
    var v15 int32 = 592
    var v16 bool = (3.14159 != (((3.14159 / 3.14159) / 1.0) / 3.14159))
    v9 = v9
    var v17 uint8 = 99
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
