package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    v1 := make(map[int]int)
    var v2 int8 = 100
    v3 := make(map[uint32]uint32)
    var v4 int8 = v2
    v5 := make(map[int]string)
    v2 = (^(v2))
    v6 := []int{-841}
    v2 = v2
    var v7 int16 = 32767
    v2 = (-(v4))
    for fi8 := 0; fi8 < 14; fi8++ {
        v4 = (^(v4))
        fmt.Println(v4, fi8)
        _ = fi8
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        if ((fi9 << ((148) & 63)) != fi9) {
            v10 := GenS1{F0: -2147483648}
            _ = v10
        } else {
            v7 -= -1
        }
        if false {
            var v11 float64 = 3.14159
            var v12 int64 = int64(v2)
            for fi13 := 0; fi13 < 18; fi13++ {
                v12 = 1
                var v14 uint64 = 0
                v12 += v12
                _ = fi13
                _ = v14
            }
            _ = v11
            _ = v12
        } else {
            if false {
                var v15 int8 = v4
                v6 = append(v6, 825)
                _ = v15
            }
        }
        _ = fi9
    }
    if false {
        switch "a" {
        case "\n\t":
            var v16 uint32 = 42
            _ = v16
        case "hello":
            fmt.Println(v4, v4)
        case "":
            if (true || (true && false)) {
                fmt.Println(v4, v7)
                if (4294967295 < 325) {
                    var v17 int8 = (v2 << ((55) & 63))
                    fmt.Println(v2)
                    v6 = append(v6, -369)
                    _ = v17
                } else {
                    var v18 uint = 255
                    _ = v18
                }
                fmt.Println(v4, v7)
            }
        }
        fmt.Println(v2, v4)
        v4 += (v2 >> ((255) & 63))
    } else {
        v19 := GenS1{F0: -1}
        if (2 > 996) {
            v2 = -128
            v2 = (v2 - (v2 % ((v4) | 1)))
        } else {
            switch ("a" + ("a" + "")) {
            case "foo bar":
                v7 -= v7
            }
            fmt.Println(v4, v4, v4)
            v7 += (v7 ^ -982)
        }
        var v20 uint32 = 2147483648
        _ = v19
        _ = v20
    }
    fmt.Println(v4, v7, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
