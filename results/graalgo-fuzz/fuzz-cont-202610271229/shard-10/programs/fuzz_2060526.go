package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 bool, p1 int64, p2 int16) uint32 {
    p2 += 32767
    if (true && true) {
        fmt.Println(p0, p0)
        var v4 uint32 = 65535
        _ = v4
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 int64 = 906
    var v2 int32 = 1
    var v3 uint32 = 42
    var v5 int = 467
    switch "\n\t" {
    case "hello":
        fmt.Println(v5)
    case "a":
        for fi6 := 0; fi6 < 9; fi6++ {
            var v7 float64 = 0.1
            _ = fi6
            _ = v7
        }
    case "foo bar":
        for fi8 := 0; fi8 < 9; fi8++ {
            switch fi8 {
            case -2147483648:
                v1 -= v1
            }
            fmt.Println(v5, v1)
            _ = fi8
        }
    default:
        v2 -= (^(v2))
    }
    var v9 uint = 42
    if false {
        v10 := GenS1{F0: v5}
        _ = v10
    } else {
        v2 = v2
        var v11 bool = (-128 > 42)
        v12 := GenS1{F0: (v5 - 1)}
        _ = v11
        _ = v12
    }
    switch v5 {
    case 80:
        fmt.Println(v3, v9)
    case -231:
        v2 = v2
    case -1:
        switch v5 {
        case 857:
            v13 := []int{}
            _ = v13
        case 593:
            for fi14 := 0; fi14 < 5; fi14++ {
                fmt.Println(v3, v2)
                _ = fi14
            }
        }
    }
    var v15 uint8 = 202
    v9 = (v9 << ((v9) & 63))
    var v16 string = "foo bar"
    var v17 int = (v5 << (((^(v9))) & 63))
    for fi18 := 0; fi18 < 3; fi18++ {
        if false {
            fmt.Println(v5)
        }
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v5
    _ = v9
}
