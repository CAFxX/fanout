package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int }

func main() {
    var v1 uint64 = 61
    var v2 int8 = -1
    var v3 float32 = -2.25
    var v4 uint = 18446744073709551615
    v5 := []int{-1, int(v4), -189, 245}
    if ((-2147483648 != -1) || false) {
        var v6 string = "hello"
        _ = v6
    }
    v7 := GenS1{F0: v1}
    v1 = (^(v1))
    for fi8 := 0; fi8 < 1; fi8++ {
        v3 = (float32((3.14159 + (3.14159 / 1.0))) + v3)
        _ = fi8
    }
    v9 := GenS1{F0: 18446744073709551615}
    v2 -= 0
    v10 := GenS1{F0: 854}
    if true {
        v3 += 1e10
    } else {
        fmt.Println(v3, v4)
    }
    v1 = v1
    v4 = v4
    v1 -= (v1 + (v1 << ((v4) & 63)))
    v11 := GenS1{F0: 2}
    for fi12 := 0; fi12 < 16; fi12++ {
        switch (fi12 >> ((255) & 63)) {
        case 1:
            v13 := GenS1{F0: (^(v1))}
            _ = v13
        case -9223372036854775808:
            for fi14 := 0; fi14 < 3; fi14++ {
                v5 = append(v5, 566)
                var v15 string = ("foo bar" + "foo bar")
                v1 = (v1 % ((v1) | 1))
                _ = fi14
                _ = v15
            }
        case 2147483647:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        fmt.Println(v3, v4, v4)
        _ = fi12
    }
    _ = len(v5)
    v4 = (v4 | (v4 & (v4 << ((v4) & 63))))
    for fi16 := 0; fi16 < 3; fi16++ {
        v17 := GenS2{F0: -1}
        _ = fi16
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v9
}
