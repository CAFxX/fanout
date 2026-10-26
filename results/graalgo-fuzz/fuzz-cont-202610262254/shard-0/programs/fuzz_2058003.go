package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
func genFn0(p0 uint64, p1 int8) int64 {
    v4 := []int{}
    var v5 uint64 = (p0 << ((0) & 63))
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return -9223372036854775808
}
func genFn1(p0 string, p1 uint) int64 {
    if (1 > -1) {
        v6 := []int{607}
        v7 := GenS1{F0: ((p1 % (((^(p1))) | 1)) >= (^(p1))), F1: 18446744073709551615}
        _ = v6
        _ = v7
    }
    var v8 uint16 = 957
    _ = p0
    _ = p1
    _ = v8
    return genFn0(18446744073709551615, -1)
}

func main() {
    var v1 int16 = -1
    var v2 uint64 = 18446744073709551615
    var v3 uint8 = 200
    for fi9 := 0; fi9 < 20; fi9++ {
        v1 = (v1 << ((842) & 63))
        _ = fi9
    }
    if (0.0 != 1e10) {
        v2 -= 18446744073709551615
        var v10 int8 = int8(v3)
        for fi11 := 0; fi11 < 9; fi11++ {
            v12 := GenS1{F0: false, F1: v2}
            v10 = 127
            v13 := GenS1{F0: (("foo bar" + "hello") <= ("a" + ((("foo bar" + "hello") + "foo bar") + "a"))), F1: (v2 / ((v2) | 1))}
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = v10
    }
    v14 := GenS1{F0: (0.0 >= (0.0 - 0.1)), F1: 42}
    v15 := GenS1{F0: false, F1: 0}
    v3 -= (v3 >> ((572) & 63))
    for fi16 := 0; fi16 < 4; fi16++ {
        v2 = v2
        _ = fi16
    }
    fmt.Println(v1)
    v17 := []int{int(v1), 128, 127}
    var v18 int32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
