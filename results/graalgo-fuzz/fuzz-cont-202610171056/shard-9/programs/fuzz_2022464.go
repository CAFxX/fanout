package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }

func main() {
    var v1 uint8 = 128
    var v2 int = 2147483647
    var v3 uint64 = 4294967295
    var v4 uint64 = v3
    v2 -= v2
    for fi5 := 0; fi5 < 2; fi5++ {
        if true {
            for fi6 := 0; fi6 < 9; fi6++ {
                v7 := []int{v2, fi6, (fi6 & (fi6 & fi5))}
                var v8 uint16 = 255
                _ = fi6
                _ = v7
                _ = v8
            }
            var v9 uint32 = 694
            _ = v9
        } else {
            for fi10 := 0; fi10 < 8; fi10++ {
                v1 -= (v1 + 218)
                v11 := GenS1{F0: 42, F1: (^(v4))}
                _ = fi10
                _ = v11
            }
            v12 := []int{}
            if true {
                v3 -= 42
                var v13 string = "foo bar"
                _ = v13
            }
            _ = v12
        }
        _ = fi5
    }
    v4 = v4
    v2 = v2
    fmt.Println(v3, v4, v4)
    for fi14 := 0; fi14 < 11; fi14++ {
        v15 := []int{}
        _ = len(v15)
        _ = fi14
        _ = v15
    }
    v3 = v4
    v16 := []int{v2}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
