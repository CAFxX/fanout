package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int64 }
func genFn0() int {
    var v6 string = ""
    v6 = v6
    switch -1 {
    case 32768:
        fmt.Println(v6)
    }
    if (!(1 <= 32767)) {
        v6 = v6
        v6 = v6
    } else {
        var v7 uint16 = 32767
        for fi8 := 0; fi8 < 7; fi8++ {
            for fi9 := 0; fi9 < 6; fi9++ {
                v10 := []int{}
                if (0 == -1) {
                    if false {
                        switch ("\n\t" + "\n\t") {
                        case "\n\t":
                            fmt.Println(fi8, fi9)
                        case "a":
                            _ = len(v10)
                        default:
                            fmt.Println(v6, v7)
                        }
                        v7 -= 42
                    }
                }
                var v11 float64 = (((0.0 / 1.0) * 3.14159) - float64(float32(3.14159)))
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = fi8
        }
        fmt.Println(v7, v7)
        _ = v7
    }
    _ = v6
    return 236
}

func main() {
    var v1 uint32 = 206
    var v2 int64 = -1
    v3 := make(map[int]string)
    v4 := make(map[int]int)
    v5 := make(map[int]int)
    for fi12 := 0; fi12 < 1; fi12++ {
        for fi13 := 0; fi13 < 11; fi13++ {
            fmt.Println(v1, fi13, fi12)
            _ = fi13
        }
        v14 := GenS1{F0: 32767, F1: 2, F2: v2}
        var v15 int16 = -32768
        _ = fi12
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        v1 += (v1 | 1)
        v2 = (^(v2))
        fmt.Println(v2, v1, v2)
        _ = fi16
    }
    v1 += 2
    switch 314 {
    case 833:
        var v17 uint8 = 0
        _ = v17
    case 32768:
        for fi18 := 0; fi18 < 15; fi18++ {
            v1 -= (v1 * (v1 >> ((65535) & 63)))
            v2 = v2
            _ = fi18
        }
    case 2147483647:
        v2 += (v2 & (^(v2)))
    }
    if (1 == (v1 ^ (v1 * v1))) {
        v19 := []int{127, 127}
        var v20 int = genFn0()
        _ = v19
        _ = v20
    }
    v1 -= 1
    var v21 int32 = -454
    fmt.Println(v1)
    v2 += v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
