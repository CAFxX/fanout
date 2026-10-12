package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8 }

func main() {
    v1 := make(map[int]bool)
    var v2 int8 = -48
    var v3 int = 4294967295
    var v4 uint32 = 2147483647
    var v5 uint8 = 0
    v6 := []int{v3, (v3 & (v3 | int(v4))), v3, (v3 << ((830) & 63))}
    v7 := GenS1{F0: 786}
    for fi8 := 0; fi8 < 6; fi8++ {
        fmt.Println(v3, v5, v5)
        switch v3 {
        case 1:
            fmt.Println(v3)
        case 127:
            var v9 int32 = 145
            _ = v9
        case 4294967295:
            _ = len(v6)
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi8
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        v2 = v2
        var v11 uint8 = 42
        _ = fi10
        _ = v11
    }
    if false {
        v12 := GenS2{F0: (v2 ^ (v2 % (((v2 & (v2 << ((42) & 63)))) | 1)))}
        _ = v12
    } else {
        v2 += 98
        if (1e10 < 1e10) {
            switch "a" {
            case "":
                if ((("hello" <= ("a" + (("a" + "a") + "\n\t"))) && true) || ((v4 + v4) >= v4)) {
                    for fi13 := 0; fi13 < 3; fi13++ {
                        v14 := GenS1{F0: v4}
                        _ = fi13
                        _ = v14
                    }
                } else {
                    v15 := GenS2{F0: v2}
                    _ = v15
                }
            case "\n\t":
                var v16 uint = 461
                _ = v16
            default:
                fmt.Println(v2, v2)
            }
            var v17 bool = (255 < 32767)
            _ = v17
        }
    }
    var v18 bool = true
    v6 = append(v6, 4294967295)
    fmt.Println(v4, v3, v18)
    v19 := GenS1{F0: (^(v4))}
    if v18 {
        v5 = 1
        v18 = (-68 != (v2 | (v2 + 1)))
    }
    var v20 uint = 636
    v21 := GenS1{F0: v4}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
