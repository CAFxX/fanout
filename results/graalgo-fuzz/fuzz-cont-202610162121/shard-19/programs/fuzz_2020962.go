package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint }

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint8 = 1
    var v3 int8 = -1
    v4 := make(map[int]uint32)
    var v5 uint = 178
    var v6 uint16 = 2
    var v7 int64 = 0
    v8 := []int{42}
    var v9 uint32 = 319
    fmt.Println(v9, v5)
    if true {
        if true {
            if (true || (0.0 < 0.0)) {
                v3 = 0
                fmt.Println(v7, v6)
            }
            delete(v4, -971)
        }
        v2 -= (v2 + v2)
        v1 = 4294967295
    } else {
        v1 -= 986
        if false {
            v3 = int8(v2)
            v5 += uint(v2)
            v5 += uint(v1)
        }
        v10 := GenS1{F0: (v5 ^ (^(v5))), F1: -353}
        _ = v10
    }
    switch 2147483647 {
    case -704:
        v2 = (^(v2))
    case 113:
        fmt.Println(v9, v2, v9)
    case 32768:
        var v11 int = 42
        _ = v11
    }
    var v12 int16 = -32768
    if (v6 != 65535) {
        if ((!true) && ((v7 * v7) <= v7)) {
            for fi13 := 0; fi13 < 11; fi13++ {
                var v14 uint64 = (v1 * (v1 + 4294967295))
                if false {
                    var v15 bool = false
                    _ = v15
                } else {
                    if true {
                        v16 := GenS2{F0: v5}
                        _ = len(v8)
                        _ = v16
                    }
                }
                _ = fi13
                _ = v14
            }
        } else {
            var v17 int8 = -52
            _ = v17
        }
    } else {
        for fi18 := 0; fi18 < 1; fi18++ {
            var v19 uint64 = 42
            if false {
                fmt.Println(v7)
                var v20 uint16 = v6
                _ = v20
            }
            var v21 bool = (v7 < v7)
            _ = fi18
            _ = v19
            _ = v21
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
