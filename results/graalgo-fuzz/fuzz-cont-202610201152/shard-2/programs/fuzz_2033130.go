package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 int32 = -1
    v2 := make(map[string]uint32)
    var v3 uint = 4294967295
    v1 = -1
    var v4 int = 32768
    for fi5 := 0; fi5 < 14; fi5++ {
        for fi6 := 0; fi6 < 13; fi6++ {
            v3 += (v3 * v3)
            if (!true) {
                v4 = (fi5 / ((fi5) | 1))
            }
            _ = fi6
        }
        v7 := v2[("foo bar" + ("a" + ("foo bar" + "\n\t")))]
        _ = fi5
        _ = v7
    }
    v4 -= 9223372036854775807
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 int8 = -1
        for fi10 := 0; fi10 < 16; fi10++ {
            var v11 uint = v3
            _ = fi10
            _ = v11
        }
        if ((true || false) || (false || (!(false || ((0.0 != (-2.25 + 1.5)) || (0.0 >= 0.0)))))) {
            v9 = -27
            v12 := GenS1{F0: 0, F1: -649}
            v9 -= v9
            _ = v12
        } else {
            var v13 int32 = 153
            v14 := GenS1{F0: fi8, F1: 1}
            _ = v13
            _ = v14
        }
        _ = fi8
        _ = v9
    }
    v4 -= (v4 - 2147483647)
    fmt.Println(v1, v1, v3)
    if false {
        var v15 string = "a"
        if false {
            delete(v2, v15)
            var v16 int8 = -128
            v17 := []int{(v4 >> ((v3) & 63))}
            _ = v16
            _ = v17
        } else {
            var v18 int8 = -1
            var v19 uint8 = 190
            _ = v18
            _ = v19
        }
        _ = v15
    } else {
        v20 := []int{}
        v21 := GenS2{F0: 715}
        var v22 int8 = int8(v4)
        _ = v20
        _ = v21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
