package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 int; F2 int64 }

func main() {
    var v1 uint32 = 2147483648
    var v2 bool = (!(!true))
    v3 := make(map[int]bool)
    var v4 string = ""
    var v5 uint = 2
    fmt.Println(v1, v1, v4)
    v6 := []int{127}
    v2 = false
    v6 = append(v6, 522)
    v5 -= (v5 ^ 1)
    var v7 int8 = -47
    for fi8 := 0; fi8 < 17; fi8++ {
        if v2 {
            if true {
                v7 -= int8(v5)
                if true {
                    if (-1 >= 406) {
                        var v9 uint = (v5 | (^(v5)))
                        v10 := GenS1{F0: 42}
                        _ = v10
                        _ = v9
                    }
                    fmt.Println(v5)
                }
                if false {
                    fmt.Println(v7, v1, v5)
                    v4 = v4
                }
            } else {
                if len(v6) > 0 {
                    v6[0] = (fi8 << ((v5) & 63))
                }
                fmt.Println(v4)
            }
            v1 += 2
        } else {
            var v11 string = "foo bar"
            v12 := GenS1{F0: 65535}
            if len(v6) > 0 {
                v6[0] = 955
            }
            _ = v11
            _ = v12
        }
        fmt.Println(v5)
        _ = fi8
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        var v14 float64 = -2.25
        v15 := GenS2{F0: 42, F1: (-(fi13)), F2: 760}
        if (61 <= -2147483648) {
            for fi16 := 0; fi16 < 11; fi16++ {
                var v17 uint16 = 906
                _ = fi16
                _ = v17
            }
            v7 -= 38
        } else {
            for fi18 := 0; fi18 < 6; fi18++ {
                v4 = v4
                _ = fi18
            }
            v19 := v3[(fi13 | (fi13 * fi13))]
            _ = v19
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
