package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 int64 = 0
    var v2 int = -864
    var v3 float32 = 0.0
    if (v1 <= v1) {
        var v4 bool = false
        _ = v4
    }
    v5 := GenS1{F0: 1}
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := []int{(fi6 / ((v2) | 1)), 127, (v2 >> ((0) & 63)), (fi6 - (v2 * 32767))}
        _ = fi6
        _ = v7
    }
    if (761 != 2) {
        if (true && false) {
            for fi8 := 0; fi8 < 11; fi8++ {
                var v9 string = "a"
                _ = fi8
                _ = v9
            }
            v10 := []int{789}
            _ = v10
        } else {
            if ((v2 ^ v2) > v2) {
                var v11 int32 = 1
                _ = v11
            }
            v3 = (1e10 * float32(1e10))
        }
    } else {
        v12 := []int{(v2 & 780)}
        if true {
            v1 -= int64(v2)
            if (49 >= 421) {
                v3 += v3
                if len(v12) > 0 {
                    v12[0] = (^(v2))
                }
            } else {
                for fi13 := 0; fi13 < 8; fi13++ {
                    var v14 int64 = (v1 % (((^(v1))) | 1))
                    var v15 uint = 839
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            }
        }
        var v16 bool = (((!(32767 != (v2 & v2))) || true) || false)
        _ = v12
        _ = v16
    }
    v17 := []int{v2, (^(v2)), v2, v2}
    fmt.Println(v2)
    for fi18 := 0; fi18 < 9; fi18++ {
        for fi19 := 0; fi19 < 7; fi19++ {
            var v20 int16 = 1
            _ = fi19
            _ = v20
        }
        v1 = v1
        _ = fi18
    }
    for fi21 := 0; fi21 < 1; fi21++ {
        for fi22 := 0; fi22 < 12; fi22++ {
            var v23 float64 = 3.14159
            _ = fi22
            _ = v23
        }
        if len(v17) >= 2 {
            v17 = v17[1:len(v17)]
        }
        _ = fi21
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v5
}
