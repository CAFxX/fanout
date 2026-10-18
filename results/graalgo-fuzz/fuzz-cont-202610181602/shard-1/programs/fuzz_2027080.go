package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool }

func main() {
    var v1 int16 = -32768
    v2 := make(map[uint32]int)
    var v3 int = 32768
    v4 := make(map[string]bool)
    var v5 int = 0
    fmt.Println(v5)
    if false {
        if true {
            for fi6 := 0; fi6 < 9; fi6++ {
                fmt.Println(v5, v5, v1)
                var v7 uint32 = 1
                _ = fi6
                _ = v7
            }
            fmt.Println(v5, v3, v3)
            for fi8 := 0; fi8 < 3; fi8++ {
                v9 := GenS1{F0: 42}
                _ = fi8
                _ = v9
            }
        }
        for fi10 := 0; fi10 < 11; fi10++ {
            var v11 string = "hello"
            _ = fi10
            _ = v11
        }
    } else {
        var v12 uint64 = 42
        _ = v12
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        v1 -= (v1 - 483)
        for fi14 := 0; fi14 < 4; fi14++ {
            switch (v5 - (v5 & (v5 % ((v3) | 1)))) {
            case -965:
                v15 := []int{v3}
                fmt.Println(v1)
                _ = v15
            }
            var v16 float64 = -2.25
            v3 = (v5 << ((18446744073709551615) & 63))
            _ = fi14
            _ = v16
        }
        _ = fi13
    }
    delete(v2, 2147483648)
    switch "foo bar" {
    case "foo bar":
        for fi17 := 0; fi17 < 16; fi17++ {
            v18 := GenS2{F0: (100 < 1)}
            _ = fi17
            _ = v18
        }
    case "\n\t":
        v5 = v5
    case "hello":
        fmt.Println(v5, v3, v3)
    default:
        v19 := []int{2147483647, 4294967295}
        _ = v19
    }
    fmt.Println(v1, v1)
    fmt.Println(v3, v1, v1)
    v1 -= 42
    v1 -= v1
    for fi20 := 0; fi20 < 12; fi20++ {
        v5 = v3
        v3 -= fi20
        _ = fi20
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
