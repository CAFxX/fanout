package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    v1 := []uint32{2147483647}
    var v2 uint64 = 4294967295
    var v3 float32 = 3.14159
    var v4 uint = 65535
    for fi5 := 0; fi5 < 16; fi5++ {
        v4 += v4
        var v6 uint32 = 711
        _ = fi5
        _ = v6
    }
    var v7 int16 = 32767
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 string = ""
        v2 = (v2 << (((^(v4))) & 63))
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi8
        _ = v9
    }
    if true {
        v4 -= (v4 * (v4 & 4294967295))
    } else {
        var v10 float64 = -2.25
        if false {
            var v11 int = -810
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v11
        }
        _ = v10
    }
    var v12 int16 = (-(v7))
    fmt.Println(v7)
    switch 4294967295 {
    case -9223372036854775808:
        v13 := GenS1{F0: "\n\t"}
        _ = v13
    case 229:
        if len(v1) > 0 {
            v1[0] = 42
        }
    }
    v14 := GenS1{F0: ("hello" + "\n\t")}
    fmt.Println(v12)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
