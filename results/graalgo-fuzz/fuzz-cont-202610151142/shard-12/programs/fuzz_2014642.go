package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32 }
func genFn0(p0 int) uint8 {
    v6 := GenS1{F0: 2}
    p0 += p0
    p0 = p0
    fmt.Println(p0)
    p0 += p0
    _ = p0
    _ = v6
    return 81
}

func main() {
    var v1 int = -97
    var v2 int8 = 10
    var v3 int64 = 115
    v4 := make(map[int]int)
    var v5 float64 = float64((3.14159 / 1e10))
    if (937 == -57) {
        var v7 int8 = (v2 & v2)
        var v8 uint32 = 0
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 42}
    if ((v1 << ((98) & 63)) < v1) {
        _ = len(v4)
    } else {
        if false {
            v10 := GenS2{F0: 42}
            if false {
                var v11 uint64 = 18446744073709551615
                switch v1 {
                case 4294967295:
                    for fi12 := 0; fi12 < 16; fi12++ {
                        var v13 float32 = 0.0
                        _ = fi12
                        _ = v13
                    }
                case 451:
                    v1 += v1
                default:
                    v1 = 1
                }
                _ = v11
            }
            _ = v10
        }
        v14 := GenS2{F0: 2147483647}
        _ = v14
    }
    for fi15 := 0; fi15 < 13; fi15++ {
        v1 = 42
        _ = fi15
    }
    switch -9223372036854775808 {
    case -330:
        v16 := GenS2{F0: 0}
        _ = v16
    case -846:
        fmt.Println(v1, v3)
    default:
        var v17 int = 568
        _ = v17
    }
    if false {
        var v18 uint16 = 65535
        v4[v1] = 0
        if false {
            v5 += (v5 + ((v5 + (v5 / -2.25)) + (((1e10 / v5) + (-2.25 - (3.14159 + (v5 + float64(-2.25))))) - 0.1)))
            delete(v4, (^(v1)))
        }
        _ = v18
    }
    fmt.Println(v5, v2, v2)
    v4[v1] = -203
    v1 = (v1 % ((291) | 1))
    v2 -= (v2 << ((255) & 63))
    fmt.Println(v2, v1, v1)
    if (275 < 655) {
        for fi19 := 0; fi19 < 11; fi19++ {
            switch (fi19 >> ((2) & 63)) {
            case 0:
                fmt.Println(v1)
            case 9223372036854775807:
                var v20 int8 = (v2 - (v2 * 0))
                _ = v20
            }
            var v21 int16 = 0
            v22 := []int{v1, (^(v1))}
            _ = fi19
            _ = v21
            _ = v22
        }
    }
    switch (((("a" + "\n\t") + "hello") + ("\n\t" + ("" + ("a" + ("foo bar" + (("a" + ("" + "\n\t")) + (("a" + "a") + ""))))))) + "\n\t") {
    case "\n\t":
        v5 -= 3.14159
    case "hello":
        var v23 int16 = -406
        _ = v23
    case "":
        v2 -= 100
    }
    v24 := GenS2{F0: -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
