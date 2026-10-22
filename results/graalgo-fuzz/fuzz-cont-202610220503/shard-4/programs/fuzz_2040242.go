package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 uint32 = 1
    v2 := []string{""}
    var v3 uint32 = 2147483647
    var v4 int32 = 2147483647
    var v5 int16 = 0
    v6 := make(map[uint32]int)
    var v7 bool = true
    v4 = (v4 | v4)
    v8 := GenS1{F0: 32767}
    delete(v6, v3)
    fmt.Println(v3)
    fmt.Println(v4)
    var v9 uint8 = 0
    fmt.Println(v9, v4, v4)
    switch 0 {
    case 0:
        for fi10 := 0; fi10 < 4; fi10++ {
            v11 := GenS1{F0: (fi10 + (fi10 << ((65535) & 63)))}
            if (fi10 > -9223372036854775808) {
                v6[v3] = fi10
                if false {
                    var v12 int16 = (v5 | v5)
                    if v7 {
                        if v7 {
                            v9 += (^(v9))
                        }
                    }
                    _ = v12
                } else {
                    v1 -= 65535
                    v13 := GenS1{F0: fi10}
                    _ = v13
                }
                var v14 int32 = (v4 + (^(v4)))
                _ = v14
            }
            _ = fi10
            _ = v11
        }
    default:
        fmt.Println(v4, v4)
    }
    v15 := GenS1{F0: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
