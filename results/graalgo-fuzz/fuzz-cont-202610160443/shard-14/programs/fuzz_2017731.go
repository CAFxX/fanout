package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 uint8 = 0
    v2 := make(map[int]int)
    var v3 uint16 = 42
    v4 := make(map[uint32]bool)
    var v5 int32 = 1
    var v6 int = 4294967295
    var v7 uint8 = v1
    if false {
        if false {
            for fi8 := 0; fi8 < 4; fi8++ {
                if (!(!true)) {
                    if true {
                        var v9 float64 = 0.0
                        v9 = 0.0
                        v1 += v1
                        _ = v9
                    }
                    for fi10 := 0; fi10 < 4; fi10++ {
                        switch "foo bar" {
                        case "a":
                            var v11 float32 = -2.25
                            _ = v11
                        case "foo bar":
                            v7 -= 44
                        case "\n\t":
                            v12 := []int{(fi10 * (fi8 << ((65535) & 63)))}
                            _ = v12
                        default:
                            var v13 int32 = (v5 + (v5 << ((867) & 63)))
                            _ = v13
                        }
                        var v14 int = (^(fi10))
                        _ = fi10
                        _ = v14
                    }
                }
                _ = fi8
            }
            var v15 int16 = -32768
            fmt.Println(v6, v1, v6)
            _ = v15
        } else {
            v6 -= (v6 * 9223372036854775807)
            v1 = 214
        }
    }
    v5 -= (v5 - v5)
    switch v6 {
    case -2147483648:
        v5 = 454
    }
    var v16 int8 = 0
    var v17 int8 = (v16 << ((42) & 63))
    if false {
        v18 := GenS1{F0: 42}
        _ = v18
    } else {
        v19 := GenS1{F0: 2}
        v20 := []int{}
        v5 = v5
        _ = v19
        _ = v20
    }
    v21 := []int{32767, v6}
    var v22 uint16 = 0
    v1 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
