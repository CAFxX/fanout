package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint8, p1 string) string {
    p1 = ("" + ("hello" + "hello"))
    var v6 int16 = -1
    _ = p0
    _ = p1
    _ = v6
    return p1
}

func main() {
    v1 := []uint32{2147483648, 2147483648}
    var v2 uint8 = 1
    var v3 uint8 = v2
    v4 := []uint32{853, 70, 446, 973}
    var v5 int32 = -1
    var v7 uint16 = 1
    var v8 bool = true
    var v9 uint8 = (v3 << ((54) & 63))
    v10 := GenS1{F0: 699}
    var v11 int8 = 127
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if true {
        fmt.Println(v2)
        if true {
            for fi12 := 0; fi12 < 13; fi12++ {
                if len(v1) > 0 {
                    v1[0] = 4294967295
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v13 int64 = 42
                _ = fi12
                _ = v13
            }
            fmt.Println(v8, v11, v9)
            switch "foo bar" {
            case "":
                v3 = (v9 + 127)
            }
        }
        v8 = v8
    }
    var v14 uint64 = 1
    for fi15 := 0; fi15 < 8; fi15++ {
        v16 := GenS2{F0: 100}
        _ = fi15
        _ = v16
    }
    v17 := GenS1{F0: 0}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
