package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int16 }
func genFn0(p0 int) int {
    var v4 string = (("" + "a") + "\n\t")
    p0 = p0
    p0 = 32767
    for fi5 := 0; fi5 < 8; fi5++ {
        for fi6 := 0; fi6 < 16; fi6++ {
            v7 := GenS1{F0: 2147483647, F1: 2, F2: -1}
            v8 := GenS1{F0: 2147483647, F1: 1, F2: 1}
            _ = fi6
            _ = v7
            _ = v8
        }
        v4 = ((v4 + ("foo bar" + v4)) + "hello")
        _ = fi5
    }
    switch 32767 {
    case 908:
        fmt.Println(p0)
    }
    _ = p0
    _ = v4
    return p0
}

func main() {
    var v1 int64 = -292
    v2 := make(map[string]bool)
    var v3 string = ""
    v1 = v1
    if false {
        fmt.Println(v3)
    }
    var v9 int64 = 661
    var v10 bool = false
    fmt.Println(v9)
    fmt.Println(v3, v1)
    fmt.Println(v10, v10)
    for fi11 := 0; fi11 < 17; fi11++ {
        _ = len(v2)
        _ = fi11
    }
    fmt.Println(v10)
    for fi12 := 0; fi12 < 13; fi12++ {
        v1 = (v1 % (((v9 - v1)) | 1))
        _ = fi12
    }
    v9 += (^(v9))
    for fi13 := 0; fi13 < 14; fi13++ {
        v1 = (v9 >> ((2) & 63))
        _ = fi13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v9
}
