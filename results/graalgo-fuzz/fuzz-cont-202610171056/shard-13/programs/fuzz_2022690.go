package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint64 = 270
    v3 := make(map[int]string)
    v2 += (v2 * (v2 << ((525) & 63)))
    var v4 uint = 255
    var v5 int8 = 127
    if (1.5 < (3.14159 / 3.14159)) {
        var v6 int = 9223372036854775807
        _ = v6
    } else {
        if true {
            v2 += v2
        } else {
            fmt.Println(v5)
        }
    }
    var v7 float64 = ((3.14159 + -2.25) * 1e10)
    v5 = v5
    if (!false) {
        v4 -= 42
        for fi8 := 0; fi8 < 12; fi8++ {
            v4 -= 42
            if (!((v4 | v4) != v4)) {
                v5 = v5
            } else {
                switch ("a" + "hello") {
                case "hello":
                    fmt.Println(fi8)
                }
                if false {
                    var v9 uint8 = 128
                    v10 := GenS1{F0: uint8(v4)}
                    _ = v10
                    _ = v9
                } else {
                    v11 := []int{810, (fi8 * fi8), 127, (fi8 * (fi8 + 42))}
                    _ = v11
                }
            }
            var v12 string = "a"
            _ = fi8
            _ = v12
        }
        fmt.Println(v4)
    } else {
        var v13 int16 = 1000
        var v14 uint32 = 2147483648
        fmt.Println(v2, v13)
        _ = v13
        _ = v14
    }
    if (-1 != -169) {
        fmt.Println(v2, v7, v2)
        v15 := GenS1{F0: 1}
        _ = v15
    }
    for fi16 := 0; fi16 < 13; fi16++ {
        fmt.Println(v4)
        _ = fi16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
