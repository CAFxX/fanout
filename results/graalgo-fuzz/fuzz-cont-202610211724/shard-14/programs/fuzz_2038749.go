package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int }
func genFn0(p0 uint32, p1 string) int64 {
    for fi4 := 0; fi4 < 8; fi4++ {
        var v5 uint8 = 255
        p0 = 65535
        _ = fi4
        _ = v5
    }
    p0 = p0
    _ = p0
    _ = p1
    return 42
}

func main() {
    v1 := []uint32{}
    v2 := []string{}
    var v3 uint = 750
    fmt.Println(v3)
    _ = len(v2)
    var v6 int32 = -1
    for fi7 := 0; fi7 < 7; fi7++ {
        fmt.Println(v6, v3)
        v6 = 42
        v3 = (v3 & v3)
        _ = fi7
    }
    v3 = (^(v3))
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 int32 = -1
        _ = fi8
        _ = v9
    }
    if false {
        if (123 <= 1) {
            fmt.Println(v6, v3)
            for fi10 := 0; fi10 < 18; fi10++ {
                v11 := GenS1{F0: -32768}
                v6 = v6
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi10
                _ = v11
            }
        }
        var v12 uint16 = 255
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v12
    } else {
        var v13 uint = v3
        v6 = (v6 >> ((0) & 63))
        _ = v13
    }
    var v14 uint16 = 42
    var v15 uint8 = 25
    _ = len(v1)
    v16 := GenS2{F0: 32767}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v6
}
