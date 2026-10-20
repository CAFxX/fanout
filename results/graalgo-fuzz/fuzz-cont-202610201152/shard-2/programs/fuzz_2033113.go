package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 string, p1 int) int64 {
    var v6 uint16 = 32767
    p1 = (p1 & (p1 + p1))
    p1 += (-(p1))
    var v7 float64 = 1.5
    v7 = v7
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 222
}

func main() {
    var v1 string = "hello"
    v2 := make(map[string]bool)
    v3 := make(map[string]int)
    var v4 int8 = -1
    v5 := make(map[string]int)
    for fi8 := 0; fi8 < 6; fi8++ {
        for fi9 := 0; fi9 < 7; fi9++ {
            if (!true) {
                _ = len(v3)
                fmt.Println(v4, fi9)
                _ = len(v3)
                var v10 uint32 = 2147483647
                _ = v10
            }
            _ = fi9
        }
        _ = fi8
    }
    switch -1 {
    case 2147483647:
        var v11 int64 = genFn0(v1, 9223372036854775807)
        _ = v11
    case 32768:
        v1 = (v1 + v1)
    }
    v1 = (((("\n\t" + v1) + v1) + "\n\t") + v1)
    v4 -= v4
    var v12 uint = 1
    fmt.Println(v1, v12)
    switch -822 {
    case -401:
        v1 = v1
    case 2147483647:
        v12 += v12
    default:
        if false {
            for fi13 := 0; fi13 < 15; fi13++ {
                v4 = (v4 % ((51) | 1))
                var v14 uint = (^(v12))
                _ = fi13
                _ = v14
            }
        }
    }
    v12 = v12
    fmt.Println(v1)
    switch -993 {
    case 127:
        v4 += v4
    case 2147483647:
        v15 := GenS1{F0: 42}
        _ = v15
    case 1:
        if false {
            for fi16 := 0; fi16 < 3; fi16++ {
                var v17 uint32 = 65535
                fmt.Println(v4, v4, v4)
                if (147 == 42) {
                    v1 = v1
                } else {
                    v18 := v3[v1]
                    fmt.Println(fi16, fi16)
                    var v19 string = ""
                    _ = v18
                    _ = v19
                }
                _ = fi16
                _ = v17
            }
            v20 := GenS1{F0: 112}
            _ = v20
        } else {
            v21 := GenS2{F0: 42}
            _ = v21
        }
    default:
        for fi22 := 0; fi22 < 11; fi22++ {
            var v23 int16 = -1
            fmt.Println(v12)
            v4 = v4
            _ = fi22
            _ = v23
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
