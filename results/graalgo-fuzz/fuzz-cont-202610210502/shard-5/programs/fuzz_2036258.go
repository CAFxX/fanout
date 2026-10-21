package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 string, p1 uint8, p2 int) uint16 {
    var v6 int8 = 42
    p0 = "foo bar"
    for fi7 := 0; fi7 < 9; fi7++ {
        for fi8 := 0; fi8 < 9; fi8++ {
            p2 += p2
            if (42 > 0) {
                var v9 uint16 = 65535
                _ = v9
            }
            p0 = (p0 + p0)
            _ = fi8
        }
        var v10 bool = ((true && true) && ((-618 < 42) && false))
        v11 := GenS1{F0: v10}
        _ = fi7
        _ = v10
        _ = v11
    }
    var v12 int32 = -1
    if false {
        var v13 uint16 = 229
        var v14 uint16 = 240
        _ = v13
        _ = v14
    } else {
        v15 := GenS1{F0: true}
        fmt.Println(p1, p1, p2)
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v6
    return 32767
}

func main() {
    var v1 int16 = 42
    var v2 int8 = 0
    v3 := []int{909, -2147483648, 9223372036854775807}
    var v4 int8 = (v2 % ((v2) | 1))
    var v5 int32 = -1
    if false {
        _ = len(v3)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    fmt.Println(v4)
    for fi16 := 0; fi16 < 3; fi16++ {
        v4 = v4
        v4 = (^(v2))
        _ = fi16
    }
    v5 += v5
    fmt.Println(v2, v5, v5)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v17 := GenS1{F0: false}
    v5 = (v5 * v5)
    var v18 int16 = (v1 & (-(v1)))
    fmt.Println(v4, v1)
    _ = len(v3)
    if false {
        for fi19 := 0; fi19 < 5; fi19++ {
            var v20 string = ("" + "hello")
            switch fi19 {
            case 719:
                v5 += 284
            case -166:
                var v21 uint32 = 249
                _ = v21
            }
            _ = fi19
            _ = v20
        }
        v4 -= -128
    } else {
        if true {
            v22 := GenS1{F0: (1 != v2)}
            _ = v22
        }
        v2 = v4
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
