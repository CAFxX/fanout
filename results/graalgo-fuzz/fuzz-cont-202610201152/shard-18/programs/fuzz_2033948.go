package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 string, p1 uint16, p2 uint8) uint16 {
    v7 := []int{409, 506, -442}
    if (!((!(float64((3.14159 + -2.25)) > (-2.25 / -2.25))) && (!(-9223372036854775808 > 1)))) {
        for fi8 := 0; fi8 < 16; fi8++ {
            switch p0 {
            case "foo bar":
                p0 = "a"
            }
            _ = fi8
        }
    } else {
        var v9 uint64 = 1
        fmt.Println(p2)
        _ = v9
    }
    p1 = p1
    var v10 float32 = 1.5
    p1 = 255
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    return (p1 << ((42) & 63))
}

func main() {
    var v1 uint8 = 2
    v2 := make(map[int]uint32)
    var v3 bool = true
    var v4 int32 = 2147483647
    var v5 uint16 = 0
    var v6 int16 = -32768
    if v3 {
        var v11 uint16 = genFn0("foo bar", (^(v5)), v1)
        _ = v11
    } else {
        switch "a" {
        case "\n\t":
            for fi12 := 0; fi12 < 10; fi12++ {
                v13 := GenS1{F0: -413}
                if (int8(fi12) >= -128) {
                    var v14 uint8 = v1
                    var v15 int64 = -593
                    v15 = (v15 - 1)
                    _ = v14
                    _ = v15
                }
                _ = fi12
                _ = v13
            }
        }
    }
    fmt.Println(v3)
    v6 = -32768
    v4 = (v4 ^ v4)
    var v16 float64 = 3.14159
    v17 := v2[629]
    fmt.Println(v3, v4)
    v18 := []int{-2147483648, int(v4)}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
