package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 int; F2 uint64 }
func genFn0() int {
    var v4 uint64 = 4294967295
    var v5 string = "hello"
    if true {
        var v6 int = -1
        v7 := GenS1{F0: 42}
        v6 -= 9223372036854775807
        _ = v6
        _ = v7
    } else {
        v8 := []int{-339, -9223372036854775808}
        v8 = append(v8, -2147483648)
        var v9 string = v5
        _ = v8
        _ = v9
    }
    switch v5 {
    case "":
        if true {
            v4 = (v4 + v4)
            v5 = v5
        }
    case "hello":
        fmt.Println(v4, v4)
    }
    var v10 uint16 = 32767
    _ = v10
    _ = v4
    _ = v5
    return 1
}

func main() {
    v1 := []string{}
    var v2 bool = false
    var v3 float32 = 1.5
    var v11 uint32 = 42
    for fi12 := 0; fi12 < 6; fi12++ {
        if (true || (((29 > 100) && (1 == -128)) && true)) {
            var v13 uint16 = 490
            _ = len(v1)
            for fi14 := 0; fi14 < 14; fi14++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v2 = (-472 > 0)
                _ = fi14
            }
            _ = v13
        }
        fmt.Println(v2, v3, v2)
        _ = fi12
    }
    var v15 bool = false
    var v16 float32 = ((0.1 / 1.0) + v3)
    v1 = append(v1, "\n\t")
    var v17 float32 = v16
    v15 = v15
    var v18 int64 = -769
    v19 := GenS1{F0: v18}
    for fi20 := 0; fi20 < 12; fi20++ {
        v21 := GenS2{F0: 891, F1: fi20, F2: 0}
        _ = fi20
        _ = v21
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v22 int = genFn0()
    v23 := GenS2{F0: 18446744073709551615, F1: genFn0(), F2: 893}
    v24 := GenS1{F0: 9223372036854775807}
    v25 := GenS2{F0: 0, F1: 2147483647, F2: 0}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
