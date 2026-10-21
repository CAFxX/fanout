package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 string, p1 string) uint16 {
    fmt.Println(p1)
    var v4 uint32 = 329
    _ = p0
    _ = p1
    _ = v4
    return 42
}
func genFn1() int8 {
    var v5 uint = 1
    v5 = (v5 ^ (v5 * v5))
    for fi6 := 0; fi6 < 12; fi6++ {
        if (0 == 430) {
            fmt.Println(fi6, fi6)
        }
        v7 := []int{}
        _ = fi6
        _ = v7
    }
    _ = v5
    return 127
}

func main() {
    var v1 uint = 2
    var v2 string = ("a" + "")
    var v3 int8 = 3
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 int = (fi8 - (fi8 << (((^(v1))) & 63)))
        fmt.Println(fi8)
        if (1000 > -32768) {
            var v10 int64 = -408
            _ = v10
        }
        _ = fi8
        _ = v9
    }
    v3 -= (^(v3))
    if (-825 != 1) {
        v2 = "foo bar"
    }
    var v11 int32 = -2147483648
    var v12 uint8 = 91
    v1 = (v1 / (((v1 % ((v1) | 1))) | 1))
    for fi13 := 0; fi13 < 12; fi13++ {
        v14 := []int{42, fi13, -44, (fi13 + 575)}
        if ((v12 > 135) && true) {
            v15 := GenS1{F0: (v11 - (v11 - v11))}
            fmt.Println(v12, v1, v1)
            _ = v15
        } else {
            for fi16 := 0; fi16 < 6; fi16++ {
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
                var v17 float32 = -2.25
                _ = fi16
                _ = v17
            }
        }
        for fi18 := 0; fi18 < 17; fi18++ {
            var v19 string = ((v2 + v2) + "\n\t")
            _ = fi18
            _ = v19
        }
        _ = fi13
        _ = v14
    }
    if (true && false) {
        for fi20 := 0; fi20 < 18; fi20++ {
            fmt.Println(v1, fi20, v2)
            _ = fi20
        }
        var v21 float32 = float32(-2.25)
        v11 = v11
        _ = v21
    } else {
        if true {
            v22 := GenS1{F0: (v11 - int32(v12))}
            v2 = ("foo bar" + (v2 + (((((v2 + (v2 + v2)) + v2) + "a") + (v2 + v2)) + "")))
            _ = v22
        }
        var v23 int = -32
        _ = v23
    }
    v2 = (((v2 + v2) + v2) + v2)
    v2 = (v2 + v2)
    var v24 int64 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v24
    _ = v3
}
