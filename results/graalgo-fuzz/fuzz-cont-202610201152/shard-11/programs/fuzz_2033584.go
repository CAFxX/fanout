package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 bool = false
    var v2 uint64 = 2
    v3 := []int{}
    switch 128 {
    case 0:
        v4 := GenS1{F0: 779, F1: 32767}
        _ = v4
    case 42:
        for fi5 := 0; fi5 < 7; fi5++ {
            var v6 float32 = float32(3.14159)
            var v7 int = fi5
            _ = fi5
            _ = v6
            _ = v7
        }
    case -1:
        v2 += (v2 % (((v2 & (v2 & v2))) | 1))
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch "hello" {
    case "hello":
        var v8 int64 = 0
        _ = v8
    }
    if len(v3) > 0 {
        v3[0] = 4294967295
    }
    var v9 bool = v1
    v10 := GenS2{F0: 240, F1: (v2 + (v2 - v2))}
    var v11 int16 = 924
    v2 -= v2
    if false {
        if v9 {
            for fi12 := 0; fi12 < 1; fi12++ {
                _ = len(v3)
                for fi13 := 0; fi13 < 9; fi13++ {
                    fmt.Println(v11, v9)
                    _ = fi13
                }
                var v14 bool = ((!v1) || v9)
                _ = fi12
                _ = v14
            }
        }
        if true {
            var v15 uint32 = 1
            _ = v15
        }
    } else {
        v11 += (v11 - v11)
        v16 := GenS3{F0: 42}
        _ = v16
    }
    var v17 uint = 2
    v17 -= 65535
    for fi18 := 0; fi18 < 1; fi18++ {
        if len(v3) > 0 {
            v3[0] = -1
        }
        _ = fi18
    }
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v9
}
