package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
func genFn0(p0 uint32, p1 uint16, p2 string) int32 {
    if false {
        v7 := GenS1{F0: -769, F1: 689}
        p0 += p0
        _ = v7
    }
    p0 -= 0
    v8 := GenS1{F0: -2147483648, F1: 1}
    var v9 bool = (false && true)
    v10 := GenS1{F0: 42, F1: 2}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return -767
}
func genFn1(p0 bool, p1 uint) bool {
    fmt.Println(p1, p0)
    v11 := []int{9223372036854775807, -1, 0, 32767}
    if ((1 <= -9223372036854775808) || p0) {
        var v12 uint16 = 42
        var v13 float32 = 3.14159
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: 42, F1: 940}
    _ = p0
    _ = p1
    _ = v11
    _ = v14
    return (42 > 0)
}

func main() {
    var v1 string = ("\n\t" + "foo bar")
    var v2 string = ""
    var v3 uint16 = 65535
    var v4 bool = false
    var v5 uint64 = 1
    var v6 int8 = -128
    v5 = (v5 & (v5 / ((v5) | 1)))
    var v15 int8 = v6
    fmt.Println(v6, v4)
    var v16 uint32 = 1
    v16 -= 42
    switch "a" {
    case "foo bar":
        fmt.Println(v16, v15)
    case "hello":
        if v4 {
            switch 32767 {
            case 158:
                v17 := []int{}
                _ = v17
            case -9223372036854775808:
                v4 = false
            }
            for fi18 := 0; fi18 < 13; fi18++ {
                fmt.Println(v2)
                _ = fi18
            }
        }
    }
    v19 := GenS1{F0: -806, F1: (v5 + (v5 + uint64(v3)))}
    for fi20 := 0; fi20 < 20; fi20++ {
        if v4 {
            fmt.Println(v3, v2, v5)
        }
        _ = fi20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
