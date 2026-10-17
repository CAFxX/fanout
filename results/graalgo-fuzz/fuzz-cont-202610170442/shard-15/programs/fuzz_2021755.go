package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 uint }
func genFn0(p0 uint64, p1 string, p2 bool) uint8 {
    var v7 string = ("foo bar" + ("foo bar" + (p1 + p1)))
    v8 := GenS1{F0: 598}
    v9 := GenS1{F0: 18446744073709551615}
    if p2 {
        var v10 int8 = -111
        _ = v10
    }
    fmt.Println(p0, p0, v7)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 255
}

func main() {
    var v1 int = 0
    var v2 bool = false
    var v3 int = (v1 % ((v1) | 1))
    v4 := make(map[uint32]int)
    var v5 string = "foo bar"
    v6 := make(map[string]string)
    var v11 float64 = ((1e10 / 1.0) * 3.14159)
    _ = len(v4)
    v12 := GenS1{F0: 18446744073709551615}
    v11 = v11
    var v13 uint32 = 886
    for fi14 := 0; fi14 < 9; fi14++ {
        if (648 >= -436) {
            var v15 uint64 = 0
            v5 = "foo bar"
            _ = v15
        } else {
            for fi16 := 0; fi16 < 19; fi16++ {
                v17 := []int{(v1 >> ((477) & 63)), 2147483647}
                var v18 bool = ((!(0 != 934)) && (!v2))
                _ = fi16
                _ = v17
                _ = v18
            }
        }
        v13 += v13
        for fi19 := 0; fi19 < 2; fi19++ {
            if v2 {
                fmt.Println(v5, fi14)
                var v20 float64 = 0.1
                fmt.Println(v13)
                _ = v20
            }
            var v21 float64 = v11
            v22 := GenS2{F0: 100, F1: 6}
            if ((940 == (v3 - (fi14 & fi19))) || v2) {
                v1 = fi19
                v23 := v6[v5]
                v24 := GenS1{F0: 18446744073709551615}
                _ = v23
                _ = v24
            }
            _ = fi19
            _ = v21
            _ = v22
        }
        _ = fi14
    }
    v11 = 0.0
    if true {
        if v2 {
            var v25 uint16 = 971
            _ = v25
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
