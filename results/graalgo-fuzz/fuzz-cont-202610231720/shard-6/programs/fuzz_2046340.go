package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string }

func main() {
    var v1 int = 4294967295
    var v2 uint16 = 333
    var v3 uint32 = 65535
    var v4 int16 = 32767
    switch "\n\t" {
    case "foo bar":
        v4 = 56
    case "hello":
        fmt.Println(v1)
    case "\n\t":
        v1 -= v1
    default:
        for fi5 := 0; fi5 < 8; fi5++ {
            var v6 int32 = -1
            var v7 float32 = (((0.0 + (((float32(0.0) * 1.5) / 1.0) / float32(1e10))) - float32(1.5)) * 0.0)
            _ = fi5
            _ = v6
            _ = v7
        }
    }
    fmt.Println(v2, v3)
    v8 := []int{v1, int(v4)}
    v4 -= (v4 / (((v4 & v4)) | 1))
    var v9 int8 = 9
    for fi10 := 0; fi10 < 7; fi10++ {
        var v11 bool = (v3 <= v3)
        var v12 int64 = 1
        _ = fi10
        _ = v11
        _ = v12
    }
    v3 -= v3
    if true {
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    }
    var v13 bool = ((28 != 136) && true)
    v14 := GenS1{F0: 0}
    v4 = (v4 + (v4 / (((v4 + (v4 ^ -1))) | 1)))
    v3 = 24
    v15 := GenS2{F0: "hello"}
    if (!((v1 % (((v1 << ((466) & 63))) | 1)) >= v1)) {
        v2 -= (v2 + v2)
    } else {
        var v16 uint64 = 928
        _ = v16
    }
    var v17 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
