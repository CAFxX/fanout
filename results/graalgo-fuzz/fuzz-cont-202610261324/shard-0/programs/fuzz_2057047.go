package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 uint64; F2 int64 }
func genFn0(p0 bool) uint64 {
    if true {
        p0 = (587 >= 32767)
        fmt.Println(p0)
    }
    var v7 bool = p0
    v8 := GenS1{F0: 42}
    _ = p0
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint16 = 0
    v3 := make(map[string]string)
    var v4 int8 = 42
    v5 := []uint32{2}
    var v6 uint8 = 115
    var v9 uint16 = v2
    v9 = v9
    for fi10 := 0; fi10 < 11; fi10++ {
        if (127 != (v6 & v6)) {
            var v11 int = (-(fi10))
            switch "\n\t" {
            case "hello":
                fmt.Println(v2, v1, v2)
            }
            _ = v11
        } else {
            v12 := GenS1{F0: 4294967295}
            switch "" {
            case "\n\t":
                v13 := GenS1{F0: 18446744073709551615}
                _ = v13
            case "":
                var v14 int8 = (v4 << ((101) & 63))
                _ = v14
            case "foo bar":
                v1 += v1
            }
            _ = v12
        }
        _ = fi10
    }
    var v15 float32 = (-2.25 / 3.14159)
    v16 := GenS2{F0: (v1 | -1), F1: 4294967295, F2: 42}
    fmt.Println(v1)
    var v17 int8 = v4
    v1 -= 1
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
