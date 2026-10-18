package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 uint16; F2 int64 }
func genFn0(p0 uint32, p1 int64) string {
    fmt.Println(p0, p1)
    var v4 int64 = (p1 % ((0) | 1))
    v5 := []int{128}
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 uint16 = 65535
        _ = fi6
        _ = v7
    }
    var v8 float32 = 0.1
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v8
    return ""
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 650
    v3 := []int{}
    v9 := GenS1{F0: "hello"}
    v2 = v2
    fmt.Println(v2)
    if (1 > 100) {
        var v10 bool = true
        v10 = true
        _ = v10
    } else {
        var v11 uint = 223
        v12 := GenS1{F0: "foo bar"}
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        if ((42 == 682) && (false || true)) {
            var v14 uint32 = 2147483647
            _ = v14
        }
        v2 -= v2
        _ = fi13
    }
    var v15 string = "hello"
    if true {
        v2 -= v2
        var v16 int32 = 0
        _ = v16
    }
    var v17 bool = true
    var v18 int8 = -1
    for fi19 := 0; fi19 < 3; fi19++ {
        v2 = (^(v2))
        v2 += v2
        v17 = true
        _ = fi19
    }
    fmt.Println(v15, v2, v18)
    if (v17 && true) {
        for fi20 := 0; fi20 < 4; fi20++ {
            v21 := GenS2{F0: 808, F1: 2, F2: 42}
            _ = fi20
            _ = v21
        }
        for fi22 := 0; fi22 < 17; fi22++ {
            var v23 uint32 = 707
            v23 = 65535
            if v17 {
                v24 := GenS1{F0: (v15 + v15)}
                var v25 uint32 = v23
                v1[2147483647] = ""
                _ = v24
                _ = v25
            } else {
                v23 += 42
                _ = len(v1)
                switch v15 {
                case "":
                    delete(v1, 42)
                }
            }
            _ = fi22
            _ = v23
        }
    } else {
        v18 -= (^(v18))
        _ = len(v1)
        v17 = v17
    }
    v18 += 0
    var v26 int8 = (-(v18))
    v18 = (-(v26))
    v17 = (-380 > 1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v9
}
