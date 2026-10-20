package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 uint16 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]int)
    var v3 int32 = 0
    var v4 float64 = 3.14159
    v3 = (v3 * 0)
    v4 = float64(3.14159)
    var v5 uint32 = 2147483647
    var v6 string = "foo bar"
    var v7 bool = (v5 <= 4294967295)
    v7 = v7
    for fi8 := 0; fi8 < 14; fi8++ {
        v4 -= (1e10 - (0.1 * v4))
        _ = fi8
    }
    v7 = (32767 < 0)
    fmt.Println(v3, v6)
    for fi9 := 0; fi9 < 1; fi9++ {
        for fi10 := 0; fi10 < 16; fi10++ {
            v11 := GenS1{F0: 18446744073709551615}
            fmt.Println(fi10, fi9, fi10)
            _ = fi10
            _ = v11
        }
        if v7 {
            for fi12 := 0; fi12 < 17; fi12++ {
                v13 := GenS1{F0: 4294967295}
                v14 := []int{(^(fi12)), -26}
                _ = fi12
                _ = v13
                _ = v14
            }
        } else {
            if v7 {
                v15 := GenS1{F0: 857}
                for fi16 := 0; fi16 < 15; fi16++ {
                    v17 := []int{fi9, 42, 694}
                    v18 := GenS2{F0: 114, F1: 1}
                    var v19 string = v6
                    _ = fi16
                    _ = v17
                    _ = v18
                    _ = v19
                }
                for fi20 := 0; fi20 < 12; fi20++ {
                    v3 = (v3 | -150)
                    v21 := GenS2{F0: 234, F1: 684}
                    _ = fi20
                    _ = v21
                }
                _ = v15
            } else {
                v3 = (v3 >> ((450) & 63))
                fmt.Println(v4)
                v5 -= (^(v5))
            }
            var v22 uint64 = 738
            _ = v22
        }
        for fi23 := 0; fi23 < 15; fi23++ {
            fmt.Println(v5, v5, v4)
            v5 += v5
            _ = fi23
        }
        _ = fi9
    }
    var v24 uint64 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
