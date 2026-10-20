package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 int64, p1 uint64) uint32 {
    v6 := []int{-2147483648}
    if ((p1 ^ (^(p1))) <= p1) {
        v7 := GenS1{F0: 65535}
        if false {
            var v8 uint8 = 2
            _ = v8
        }
        _ = len(v6)
        _ = v7
    } else {
        switch 57 {
        case 574:
            v6 = append(v6, -946)
        case -84:
            var v9 uint16 = 805
            _ = v9
        default:
            p0 = p0
        }
    }
    _ = p0
    _ = p1
    _ = v6
    return 1
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{"hello", "foo bar"}
    var v3 uint32 = 2147483648
    v4 := []string{}
    var v5 uint8 = 0
    var v10 int32 = 1
    if true {
        var v11 uint8 = v5
        _ = v11
    } else {
        var v12 int16 = 1000
        if true {
            _ = len(v2)
            v3 = (^(v3))
            v5 = (v5 / ((v5) | 1))
        } else {
            if ((v10 ^ (v10 & (-(v10)))) <= (v10 & (^(v10)))) {
                var v13 string = ("" + ("a" + "\n\t"))
                _ = v13
            }
            var v14 uint = 430
            fmt.Println(v10, v5)
            _ = v14
        }
        var v15 uint8 = v5
        _ = v12
        _ = v15
    }
    fmt.Println(v3, v10)
    for fi16 := 0; fi16 < 3; fi16++ {
        _ = len(v2)
        _ = fi16
    }
    var v17 uint32 = genFn0(int64(v3), 104)
    v3 -= v17
    var v18 uint64 = 4294967295
    v17 = (^(v17))
    if false {
        v10 = v10
        for fi19 := 0; fi19 < 13; fi19++ {
            fmt.Println(v10, fi19)
            fmt.Println(v10, v3)
            _ = fi19
        }
        delete(v1, 32768)
    }
    var v20 bool = ((1.5 / -2.25) != 1.5)
    v17 = v3
    v1[-9223372036854775808] = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
