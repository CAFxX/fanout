package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
func genFn0(p0 int32, p1 bool, p2 uint16) bool {
    if (!(2147483648 >= 1)) {
        p1 = (!(0 >= 2147483648))
    }
    v5 := []int{-1, -1, 32768}
    p0 = p0
    if len(v5) > 0 {
        v5[0] = -9223372036854775808
    }
    var v6 uint = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (false && true)
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint = 42
    v3 := []uint32{}
    var v4 string = ""
    v2 = (v2 + 0)
    v2 = (v2 / ((v2) | 1))
    delete(v1, 2147483647)
    if (!(false || (false && (!((65535 == 42) || false))))) {
        v7 := GenS1{F0: (v2 | v2), F1: -479}
        _ = v7
    }
    var v8 uint = v2
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 uint = v8
        switch (fi9 ^ 2147483647) {
        case 1:
            v10 -= (^(v10))
        case 0:
            var v11 bool = false
            _ = v11
        }
        var v12 bool = false
        _ = fi9
        _ = v10
        _ = v12
    }
    if (false && (false || true)) {
        _ = len(v3)
    }
    v8 += 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
