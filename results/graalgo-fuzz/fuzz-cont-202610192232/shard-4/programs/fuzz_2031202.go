package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool }

func main() {
    var v1 int16 = -32768
    var v2 int32 = -2147483648
    v3 := []uint32{459, 2147483648, 4294967295, 4294967295}
    var v4 int8 = 127
    v5 := []int{4294967295, 128, 32767}
    fmt.Println(v1, v4)
    if true {
        v2 += (v2 << ((4294967295) & 63))
        var v6 int = 42
        _ = v6
    } else {
        var v7 float32 = (-2.25 - (0.0 / 3.14159))
        var v8 float32 = (v7 - 1e10)
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 18446744073709551615}
    v3 = append(v3, 0)
    if false {
        fmt.Println(v2, v2)
    }
    v1 = (v1 ^ (v1 ^ v1))
    var v10 int32 = v2
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v2, v10, v10)
    switch "foo bar" {
    case "hello":
        v11 := GenS2{F0: true}
        _ = v11
    case "foo bar":
        v12 := GenS1{F0: 412}
        _ = v12
    case "a":
        var v13 uint8 = 42
        _ = v13
    default:
        if false {
            v5 = append(v5, 127)
            fmt.Println(v2, v1)
        } else {
            for fi14 := 0; fi14 < 1; fi14++ {
                var v15 bool = false
                for fi16 := 0; fi16 < 3; fi16++ {
                    v15 = (1 == 4294967295)
                    _ = fi16
                }
                _ = fi14
                _ = v15
            }
            v10 -= (v10 & v10)
            v1 = v1
        }
    }
    for fi17 := 0; fi17 < 7; fi17++ {
        var v18 uint64 = 587
        _ = fi17
        _ = v18
    }
    v19 := GenS1{F0: 63}
    if true {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
