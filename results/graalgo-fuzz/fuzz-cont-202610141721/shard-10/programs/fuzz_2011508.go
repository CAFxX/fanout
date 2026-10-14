package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8; F1 int64; F2 int32 }
type GenS4 struct { F0 uint8 }
func genFn0() uint64 {
    v5 := GenS1{F0: 4294967295}
    var v6 uint32 = 2147483648
    var v7 int64 = 138
    _ = v5
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint16 = 65535
    var v3 uint = 205
    var v4 string = "foo bar"
    var v8 int = -344
    fmt.Println(v8)
    v9 := GenS2{F0: v2}
    var v10 int8 = -35
    v2 += (^(v2))
    var v11 int64 = 0
    v12 := GenS3{F0: v10, F1: (-(v11)), F2: 124}
    v4 = ""
    var v13 string = "foo bar"
    var v14 int = 9223372036854775807
    for fi15 := 0; fi15 < 19; fi15++ {
        fmt.Println(v4, v11)
        var v16 string = v4
        _ = fi15
        _ = v16
    }
    fmt.Println(v2, v13)
    v17 := GenS4{F0: 2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
