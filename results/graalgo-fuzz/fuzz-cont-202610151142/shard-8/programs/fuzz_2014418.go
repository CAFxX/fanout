package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 int; F1 int }

func main() {
    var v1 bool = (9223372036854775807 < -228)
    v2 := make(map[uint32]int)
    var v3 uint32 = 2147483648
    var v4 uint = 2
    v5 := make(map[int]bool)
    v6 := []uint32{4294967295}
    fmt.Println(v4, v3)
    for fi7 := 0; fi7 < 1; fi7++ {
        for fi8 := 0; fi8 < 13; fi8++ {
            v2[v3] = 32767
            var v9 int = (fi8 % ((fi8) | 1))
            var v10 uint32 = (^(v3))
            _ = fi8
            _ = v10
            _ = v9
        }
        v11 := GenS1{F0: v1, F1: 127}
        if (989 >= 42) {
            if v1 {
                for fi12 := 0; fi12 < 20; fi12++ {
                    v13 := v5[fi7]
                    _ = fi12
                    _ = v13
                }
                var v14 int32 = 0
                _ = v14
            } else {
                fmt.Println(v1, v4)
                if len(v6) > 0 {
                    v6[0] = v3
                }
            }
            fmt.Println(v4)
            v1 = false
        } else {
            for fi15 := 0; fi15 < 8; fi15++ {
                var v16 int32 = 42
                var v17 uint32 = (v3 % (((^(v3))) | 1))
                v18 := GenS1{F0: ((v16 | (v16 - (-(v16)))) >= (v16 + (-(v16)))), F1: (^(fi7))}
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
        }
        _ = fi7
        _ = v11
    }
    v3 = (^(v3))
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v4 += (v4 + v4)
    if v1 {
        v3 -= v3
        switch 634 {
        case 127:
            for fi19 := 0; fi19 < 15; fi19++ {
                fmt.Println(v3, v4, v1)
                var v20 int = (fi19 << ((42) & 63))
                delete(v5, 236)
                _ = fi19
                _ = v20
            }
        case 32767:
            v3 -= (v3 << (((v4 & (^(v4)))) & 63))
        case 587:
            v6 = append(v6, (v3 ^ (v3 & uint32(v4))))
        }
    }
    fmt.Println(v4)
    var v21 int16 = -435
    v1 = ((("foo bar" + ("hello" + "hello")) + "") < ((("a" + "\n\t") + "hello") + ""))
    var v22 uint64 = 0
    v23 := GenS2{F0: 2147483647, F1: 2147483647}
    for fi24 := 0; fi24 < 20; fi24++ {
        v4 = v4
        var v25 int16 = v21
        v26 := GenS1{F0: (0 > -1), F1: 9223372036854775807}
        _ = fi24
        _ = v25
        _ = v26
    }
    var v27 int32 = -1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
