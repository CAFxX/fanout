package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 int64) int32 {
    fmt.Println(p0)
    v5 := GenS1{F0: 4294967295}
    _ = p0
    _ = v5
    return -1
}

func main() {
    var v1 uint16 = 21
    var v2 int = -603
    var v3 uint32 = 42
    v4 := []string{("hello" + ("a" + ("a" + ""))), "a"}
    _ = len(v4)
    var v6 uint = 4294967295
    v6 = (v6 >> (((v6 & v6)) & 63))
    v3 = (^(v3))
    v6 -= 1
    switch 32767 {
    case -420:
        v6 = (v6 | (v6 ^ v6))
    }
    var v7 string = "foo bar"
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 uint32 = v3
        var v10 int = v2
        _ = fi8
        _ = v10
        _ = v9
    }
    v3 -= (^(v3))
    for fi11 := 0; fi11 < 8; fi11++ {
        if false {
            v1 = 515
            if len(v4) > 0 {
                v4[0] = "\n\t"
            }
        }
        _ = fi11
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
