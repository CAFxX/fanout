package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 string }

func main() {
    v1 := []int{0, 0, 128, 171}
    var v2 string = "\n\t"
    var v3 uint64 = 4294967295
    v4 := []int{-740, -2147483648, -9223372036854775808}
    v5 := make(map[uint32]uint32)
    v6 := []uint32{}
    v3 = (^(v3))
    var v7 uint64 = (v3 >> ((178) & 63))
    v8 := GenS1{F0: 65535}
    v7 += (^(v7))
    v9 := GenS1{F0: 65535}
    v10 := GenS1{F0: 692}
    for fi11 := 0; fi11 < 13; fi11++ {
        v7 = v7
        for fi12 := 0; fi12 < 3; fi12++ {
            var v13 int = (fi11 - fi12)
            v7 = (v7 + v3)
            delete(v5, 2147483648)
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    if false {
        v2 = v2
        var v14 int64 = 1
        if false {
            switch (((v2 + ("\n\t" + v2)) + v2) + ((v2 + "") + (v2 + v2))) {
            case "foo bar":
                v7 += v7
            case "":
                v14 = (^(v14))
            }
        }
        _ = v14
    } else {
        v5[2147483647] = 212
        fmt.Println(v3, v3, v3)
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        var v16 int8 = -1
        var v17 uint64 = 93
        _ = fi15
        _ = v16
        _ = v17
    }
    v18 := GenS2{F0: 32767, F1: ("a" + ("hello" + v2))}
    var v19 float64 = ((3.14159 - (0.1 + 0.0)) + 1e10)
    fmt.Println(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
