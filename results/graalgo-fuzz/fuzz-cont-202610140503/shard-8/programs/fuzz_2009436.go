package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int64 }
func genFn0() bool {
    switch "hello" {
    case "foo bar":
        v7 := []int{-508, 609}
        _ = v7
    case "\n\t":
        var v8 uint8 = 1
        _ = v8
    }
    if false {
        v9 := []int{42, -2147483648}
        v10 := GenS1{F0: 126, F1: 0, F2: -915}
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 17; fi11++ {
        v12 := []int{fi11, (fi11 + -508)}
        _ = fi11
        _ = v12
    }
    return false
}
func genFn1(p0 int, p1 bool) bool {
    v13 := GenS1{F0: 216, F1: -1, F2: 0}
    p0 = (-(p0))
    var v14 uint64 = 42
    p1 = p1
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return (1 < 128)
}

func main() {
    var v1 bool = ((!((210 >= 53) && true)) || true)
    var v2 uint16 = 65535
    var v3 uint = 499
    var v4 uint8 = 42
    var v5 bool = false
    var v6 uint64 = 818
    var v15 uint = v3
    for fi16 := 0; fi16 < 9; fi16++ {
        v17 := GenS1{F0: -536, F1: -2147483648, F2: -9223372036854775808}
        v18 := []int{}
        _ = fi16
        _ = v17
        _ = v18
    }
    if true {
        v15 -= 65535
        var v19 uint = 777
        _ = v19
    } else {
        for fi20 := 0; fi20 < 18; fi20++ {
            if ((0.1 - 0.0) == 3.14159) {
                v15 -= v3
            } else {
                v21 := []int{(fi20 >> ((v3) & 63)), (fi20 >> ((uint(v2)) & 63))}
                var v22 uint32 = 2147483648
                _ = v21
                _ = v22
            }
            _ = fi20
        }
        v23 := []int{9223372036854775807}
        _ = v23
    }
    if v5 {
        for fi24 := 0; fi24 < 14; fi24++ {
            for fi25 := 0; fi25 < 2; fi25++ {
                v6 -= v6
                v26 := []int{(^(fi24)), 0, (fi25 + (^(fi25))), (fi24 | 1)}
                v15 -= (v3 | 1)
                _ = fi25
                _ = v26
            }
            var v27 uint64 = (v6 | v6)
            _ = fi24
            _ = v27
        }
        v1 = v5
    } else {
        var v28 uint = (v15 << ((2) & 63))
        var v29 bool = (v15 < (v15 % ((v15) | 1)))
        _ = v28
        _ = v29
    }
    v30 := []int{949, int(v2)}
    var v31 int32 = -2147483648
    v15 = v15
    fmt.Println(v2, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
