package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int, p1 int32) int64 {
    p1 += p1
    p1 += 2147483647
    v7 := GenS1{F0: 225}
    v8 := []int{p0}
    var v9 bool = (!(2 == 1))
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 float64 = 1.5
    var v2 int32 = 0
    var v3 uint32 = 848
    var v4 string = "a"
    var v5 int32 = (v2 & v2)
    v6 := make(map[int]bool)
    v1 -= v1
    if false {
        switch "foo bar" {
        case "a":
            var v10 uint = 256
            _ = v10
        case "hello":
            v11 := []int{32768}
            _ = v11
        case "":
            var v12 uint64 = 2
            _ = v12
        default:
            var v13 int16 = 1
            _ = v13
        }
        var v14 int32 = 2147483647
        switch "" {
        case "\n\t":
            fmt.Println(v1, v14, v4)
        default:
            fmt.Println(v2, v2)
        }
        _ = v14
    }
    v15 := GenS1{F0: 127}
    v16 := GenS1{F0: 1}
    v2 -= (v2 ^ -2147483648)
    var v17 uint8 = 2
    switch 128 {
    case 1:
        v17 = (v17 % (((v17 ^ (^(v17)))) | 1))
    }
    switch v4 {
    case "a":
        for fi18 := 0; fi18 < 12; fi18++ {
            var v19 int = 128
            v20 := GenS1{F0: v17}
            switch (fi18 * int(v2)) {
            case -9223372036854775808:
                v2 += 2147483647
            case 696:
                v21 := GenS1{F0: v17}
                _ = v21
            case -197:
                v4 = "a"
            }
            _ = fi18
            _ = v19
            _ = v20
        }
    case "foo bar":
        for fi22 := 0; fi22 < 8; fi22++ {
            v1 -= v1
            _ = fi22
        }
    default:
        if (!(((v17 >> ((922) & 63)) != (v17 << ((255) & 63))) || (-1 <= -976))) {
            v2 = (v2 & v2)
        } else {
            v23 := GenS2{F0: 42}
            var v24 int = 307
            _ = v23
            _ = v24
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
