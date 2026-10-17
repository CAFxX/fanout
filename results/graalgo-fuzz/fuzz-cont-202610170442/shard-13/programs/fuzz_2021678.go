package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 int = 42
    var v2 int64 = -141
    var v3 float64 = (0.1 + float64(0.0))
    v4 := make(map[int]int)
    var v5 uint64 = 0
    switch 32767 {
    case 1:
        v6 := []int{-360, (v1 >> ((792) & 63)), v1}
        _ = v6
    case 9223372036854775807:
        var v7 bool = (0 <= 42)
        _ = v7
    default:
        var v8 int = v1
        _ = v8
    }
    v9 := []int{(v1 + (v1 >> ((51) & 63))), (^(v1))}
    v10 := GenS1{F0: v1}
    v11 := GenS1{F0: (v1 % (((^(v1))) | 1))}
    for fi12 := 0; fi12 < 6; fi12++ {
        if (!false) {
            var v13 uint64 = 1
            var v14 uint32 = 65535
            _ = v13
            _ = v14
        } else {
            var v15 int = v1
            if len(v9) > 0 {
                v9[0] = v15
            }
            _ = v15
        }
        if (!(!(true && (!(127 >= 255))))) {
            switch v1 {
            case -9223372036854775808:
                v16 := GenS1{F0: fi12}
                _ = v16
            }
            var v17 bool = false
            _ = v17
        } else {
            v18 := v4[(v1 & (fi12 << ((1) & 63)))]
            _ = v18
        }
        var v19 uint64 = 855
        _ = fi12
        _ = v19
    }
    if true {
        fmt.Println(v3, v3)
    } else {
        if false {
            switch "" {
            case "foo bar":
                var v20 uint = 4294967295
                _ = v20
            case "":
                fmt.Println(v1)
            case "a":
                v21 := GenS1{F0: (v1 / ((v1) | 1))}
                _ = v21
            default:
                if false {
                    for fi22 := 0; fi22 < 15; fi22++ {
                        var v23 uint8 = 127
                        _ = fi22
                        _ = v23
                    }
                }
            }
            v24 := GenS1{F0: v1}
            _ = v24
        } else {
            v3 -= (((v3 / (v3 - v3)) + (v3 / v3)) * (v3 * 1.5))
            switch "\n\t" {
            case "foo bar":
                v2 -= 446
            }
        }
    }
    delete(v4, (v1 >> ((65535) & 63)))
    var v25 int16 = int16(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
