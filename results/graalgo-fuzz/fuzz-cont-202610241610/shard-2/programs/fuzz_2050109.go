package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0() int {
    var v5 int16 = 32767
    fmt.Println(v5)
    for fi6 := 0; fi6 < 9; fi6++ {
        if true {
            v5 += (v5 / (((v5 / (((v5 >> ((0) & 63))) | 1))) | 1))
            var v7 int64 = 1
            _ = v7
        } else {
            v8 := []int{(fi6 >> ((1) & 63)), fi6, (fi6 << ((511) & 63))}
            _ = len(v8)
            _ = v8
        }
        v9 := []int{}
        _ = fi6
        _ = v9
    }
    if (!(1 == 513)) {
        v5 -= (v5 / (((-(v5))) | 1))
        v10 := []int{}
        _ = v10
    }
    _ = v5
    return -365
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 bool = false
    var v3 int64 = -225
    var v4 int16 = 543
    v3 -= v3
    v11 := []int{-96}
    if (!(65535 <= 18446744073709551615)) {
        switch 9223372036854775807 {
        case 0:
            switch 4294967295 {
            case 127:
                v4 += (v4 % (((v4 << ((4294967295) & 63))) | 1))
            }
        default:
            for fi12 := 0; fi12 < 5; fi12++ {
                for fi13 := 0; fi13 < 11; fi13++ {
                    v14 := GenS1{F0: v2}
                    _ = fi13
                    _ = v14
                }
                var v15 uint16 = 65535
                v3 += 9223372036854775807
                _ = fi12
                _ = v15
            }
        }
        for fi16 := 0; fi16 < 7; fi16++ {
            v4 += v4
            v3 -= (v3 >> ((255) & 63))
            _ = fi16
        }
        var v17 uint16 = 246
        _ = v17
    }
    v18 := GenS1{F0: v2}
    fmt.Println(v2, v3, v3)
    var v19 string = "foo bar"
    v11 = append(v11, 55)
    v2 = true
    if v2 {
        v3 = (v3 - v3)
        v4 = (v4 ^ (v4 / ((v4) | 1)))
    } else {
        _ = len(v11)
    }
    v2 = v2
    for fi20 := 0; fi20 < 9; fi20++ {
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        var v21 float64 = ((3.14159 - float64(1e10)) / 1.0)
        v22 := GenS1{F0: v2}
        _ = fi20
        _ = v21
        _ = v22
    }
    v19 = v19
    var v23 float64 = (1e10 - -2.25)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
