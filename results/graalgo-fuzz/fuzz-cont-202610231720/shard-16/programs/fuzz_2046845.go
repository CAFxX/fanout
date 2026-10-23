package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }

func main() {
    var v1 int8 = 0
    v2 := make(map[string]uint32)
    v3 := make(map[int]int)
    v4 := make(map[uint32]uint32)
    v1 += (v1 << ((42) & 63))
    if true {
        for fi5 := 0; fi5 < 1; fi5++ {
            v1 -= (^(v1))
            _ = fi5
        }
        fmt.Println(v1)
    }
    v1 = (-(v1))
    var v6 int8 = -1
    var v7 int64 = -9223372036854775808
    var v8 int32 = -2147483648
    var v9 int16 = 32767
    fmt.Println(v7, v7, v1)
    v8 -= v8
    for fi10 := 0; fi10 < 1; fi10++ {
        switch "foo bar" {
        case "\n\t":
            v11 := []int{(-(fi10))}
            _ = v11
        case "hello":
            v12 := GenS1{F0: true, F1: 0}
            _ = v12
        }
        _ = fi10
    }
    var v13 int32 = 1
    v6 -= (v6 % (((v6 & (^(v6)))) | 1))
    switch "\n\t" {
    case "a":
        var v14 float32 = 1.5
        _ = v14
    case "":
        v15 := []int{59, 9223372036854775807}
        _ = v15
    case "hello":
        v13 += v13
    }
    if false {
        v16 := GenS1{F0: (65535 == 123), F1: 1}
        _ = v16
    } else {
        switch ("\n\t" + "") {
        case "hello":
            var v17 bool = (!false)
            _ = v17
        }
    }
    fmt.Println(v6, v13)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
