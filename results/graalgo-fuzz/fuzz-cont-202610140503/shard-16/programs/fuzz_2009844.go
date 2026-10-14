package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 string; F1 int64; F2 int64 }
func genFn0(p0 bool, p1 int) uint {
    var v7 uint16 = 42
    v8 := []int{408, 32767, -1}
    var v9 int64 = -631
    var v10 uint32 = 551
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 18446744073709551615
}

func main() {
    var v1 int16 = -1
    var v2 string = ""
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]uint32)
    v5 := make(map[int]int)
    var v6 uint16 = 0
    v11 := GenS1{F0: "a", F1: 113}
    v6 = (v6 + v6)
    v12 := v3[0]
    if false {
        var v13 string = v2
        for fi14 := 0; fi14 < 18; fi14++ {
            if true {
                var v15 int16 = (v1 >> ((18446744073709551615) & 63))
                v16 := GenS2{F0: "", F1: -1, F2: 1}
                _ = v15
                _ = v16
            } else {
                v17 := v5[fi14]
                _ = v17
            }
            v2 = (v13 + v13)
            _ = fi14
        }
        _ = v13
    }
    switch v2 {
    case "a":
        var v18 int = -893
        _ = v18
    }
    v19 := []int{}
    v20 := v3[(v12 - v12)]
    v21 := v5[-9223372036854775808]
    switch 9223372036854775807 {
    case 566:
        for fi22 := 0; fi22 < 19; fi22++ {
            if (false || ((true && (true && true)) && false)) {
                v2 = v2
            } else {
                v12 = 2147483647
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
            }
            switch (v21 - (v21 ^ (v21 + v21))) {
            case -1:
                v23 := GenS1{F0: v2, F1: 42}
                _ = v23
            case 922:
                for fi24 := 0; fi24 < 8; fi24++ {
                    delete(v5, v21)
                    v20 += (v20 * 0)
                    _ = fi24
                }
            }
            var v25 float64 = 3.14159
            _ = fi22
            _ = v25
        }
    case -443:
        for fi26 := 0; fi26 < 3; fi26++ {
            v21 -= (v21 ^ v21)
            var v27 uint = 0
            _ = fi26
            _ = v27
        }
    case 0:
        var v28 string = ((("a" + (v2 + v2)) + v2) + "")
        _ = v28
    }
    var v29 uint16 = (v6 << ((42) & 63))
    v29 += (^(v29))
    v29 = (v6 & (v6 * v6))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
