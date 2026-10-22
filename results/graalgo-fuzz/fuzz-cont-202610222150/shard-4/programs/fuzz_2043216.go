package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 uint) int16 {
    fmt.Println(p0)
    p0 += p0
    var v7 string = ("" + "\n\t")
    _ = p0
    _ = v7
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 1
    var v3 uint32 = 706
    var v4 float64 = 0.0
    var v5 float32 = 1e10
    v6 := []uint32{}
    switch -2147483648 {
    case -1:
        for fi8 := 0; fi8 < 18; fi8++ {
            v4 = 1e10
            _ = fi8
        }
    case 9223372036854775807:
        if false {
            if false {
                v4 -= (float64(1e10) - v4)
                var v9 float64 = v4
                delete(v1, (v3 % ((42) | 1)))
                _ = v9
            }
        } else {
            v6 = append(v6, 2)
            delete(v1, (^(v3)))
        }
    case 128:
        var v10 uint32 = v3
        _ = v10
    default:
        var v11 uint16 = 981
        _ = v11
    }
    v4 += 0.0
    var v12 int8 = -59
    v3 = v3
    v4 = 0.1
    if (-1 <= -465) {
        switch "\n\t" {
        case "a":
            fmt.Println(v2)
        case "foo bar":
            switch ("hello" + "hello") {
            case "hello":
                v3 = v3
            case "\n\t":
                switch 128 {
                case -9223372036854775808:
                    v12 = -1
                case 9223372036854775807:
                    _ = len(v6)
                }
            case "foo bar":
                if false {
                    var v13 bool = (!true)
                    _ = v13
                } else {
                    v5 = v5
                    v14 := GenS1{F0: false}
                    _ = v14
                }
            default:
                _ = len(v1)
            }
        default:
            var v15 float32 = v5
            _ = v15
        }
        if false {
            var v16 float64 = v4
            _ = v16
        }
    } else {
        v2 = 0
    }
    v5 = (float32((float64(v5) - (0.0 - -2.25))) * (v5 * v5))
    v17 := v1[(v3 & v3)]
    v1[(v3 - (v3 >> ((1) & 63)))] = (!((42 == 42) && false))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
