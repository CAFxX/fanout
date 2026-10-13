package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint32 }
func genFn0(p0 uint, p1 int32, p2 uint32) string {
    var v4 int8 = 74
    if (-1 > 0) {
        if (-32768 < 1000) {
            for fi5 := 0; fi5 < 7; fi5++ {
                v6 := GenS1{F0: true, F1: p2, F2: p2}
                _ = fi5
                _ = v6
            }
        }
    } else {
        for fi7 := 0; fi7 < 3; fi7++ {
            var v8 float64 = float64(1e10)
            v9 := []int{fi7, (fi7 >> ((847) & 63)), 127}
            var v10 float64 = -2.25
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
    }
    v11 := GenS1{F0: (!((int32(p2) >= -2147483648) && true)), F1: (^(p2)), F2: p2}
    var v12 int32 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v4
    return "hello"
}

func main() {
    var v1 uint8 = 105
    v2 := make(map[uint32]bool)
    var v3 uint32 = 4294967295
    v3 = (^(v3))
    v1 += v1
    v3 = 2
    switch ("hello" + "hello") {
    case "foo bar":
        for fi13 := 0; fi13 < 19; fi13++ {
            v3 -= v3
            _ = fi13
        }
    default:
        fmt.Println(v1, v3)
    }
    v3 = (^(v3))
    var v14 bool = false
    v3 += 166
    if (((-195 != -9223372036854775808) || v14) && v14) {
        for fi15 := 0; fi15 < 4; fi15++ {
            v1 += 1
            _ = fi15
        }
        v14 = v14
        fmt.Println(v1)
    }
    if false {
        fmt.Println(v14, v1, v14)
        v3 = v3
        v3 = 869
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        v1 = v1
        switch (fi16 % (((fi16 % ((fi16) | 1))) | 1)) {
        case -1:
            var v17 string = ""
            _ = v17
        case -936:
            v3 = (v3 % ((302) | 1))
        default:
            v18 := []int{568}
            _ = v18
        }
        _ = fi16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
}
