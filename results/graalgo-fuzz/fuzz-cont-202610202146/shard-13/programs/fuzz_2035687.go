package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 bool, p1 uint16, p2 int) uint16 {
    p0 = true
    v6 := GenS1{F0: p0}
    p1 = 328
    if (1 != 0) {
        var v7 float64 = 1e10
        for fi8 := 0; fi8 < 8; fi8++ {
            var v9 int16 = -643
            switch "a" {
            case "\n\t":
                switch (fi8 * (p2 % ((fi8) | 1))) {
                case 32768:
                    var v10 int8 = 127
                    _ = v10
                case 70:
                    v11 := GenS1{F0: p0}
                    _ = v11
                }
            }
            _ = fi8
            _ = v9
        }
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (p1 * p1)
}

func main() {
    var v1 bool = (32768 >= 0)
    var v2 float32 = 1e10
    var v3 int = 9223372036854775807
    v4 := make(map[uint32]bool)
    var v5 string = "a"
    fmt.Println(v2, v2, v3)
    v12 := GenS1{F0: (18446744073709551615 != 4294967295)}
    fmt.Println(v3, v1, v3)
    var v13 uint64 = 2
    var v14 uint16 = 0
    for fi15 := 0; fi15 < 20; fi15++ {
        switch v5 {
        case "foo bar":
            var v16 uint = uint(v13)
            _ = v16
        }
        var v17 int = (v3 | v3)
        v1 = ((-32768 > 1000) || (9223372036854775807 < 610))
        _ = fi15
        _ = v17
    }
    v3 = (-(v3))
    v1 = v1
    v2 -= (v2 / v2)
    v14 += 42
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
