package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := []string{}
    var v2 uint64 = 906
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]int)
    var v5 int = 4294967295
    var v6 int16 = 87
    v2 = (v2 + v2)
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 int8 = 42
        _ = fi7
        _ = v8
    }
    var v9 uint8 = 127
    var v10 uint16 = 65535
    var v11 int16 = -1
    if true {
        if (!true) {
            v11 -= (v11 ^ (^(v11)))
            switch "hello" {
            case "\n\t":
                if false {
                    _ = len(v3)
                    var v12 string = "hello"
                    _ = v12
                } else {
                    var v13 int64 = -1
                    v10 = (^(v10))
                    _ = v13
                }
            }
        }
        v14 := GenS1{F0: 2}
        _ = v14
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        for fi16 := 0; fi16 < 8; fi16++ {
            var v17 int64 = -9223372036854775808
            var v18 uint64 = 42
            _ = fi16
            _ = v17
            _ = v18
        }
        var v19 int16 = v11
        fmt.Println(fi15)
        _ = fi15
        _ = v19
    }
    _ = len(v3)
    v11 -= v11
    v5 = 27
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
