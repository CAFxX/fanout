package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 uint64 }

func main() {
    var v1 int32 = 42
    var v2 float64 = 1.5
    v3 := []int{32767}
    var v4 int32 = 291
    var v5 uint16 = 1
    v4 = v1
    var v6 uint = 213
    v7 := GenS1{F0: 42, F1: 165, F2: 1}
    var v8 uint16 = (^(v5))
    for fi9 := 0; fi9 < 12; fi9++ {
        for fi10 := 0; fi10 < 18; fi10++ {
            switch "hello" {
            case "a":
                v11 := GenS1{F0: 0, F1: 9223372036854775807, F2: 76}
                _ = v11
            }
            _ = fi10
        }
        _ = fi9
    }
    v12 := GenS1{F0: 762, F1: 42, F2: 42}
    for fi13 := 0; fi13 < 2; fi13++ {
        v4 -= (v4 % ((425) | 1))
        _ = fi13
    }
    var v14 int8 = 42
    v14 = (-(v14))
    if ("a" <= "foo bar") {
        if ((false && true) && true) {
            v14 = (v14 * (-(v14)))
            v5 += 378
        } else {
            if ((v14 - v14) > v14) {
                var v15 uint8 = 128
                v2 += (float64(0.1) - 3.14159)
                _ = v15
            } else {
                switch -9223372036854775808 {
                case -2147483648:
                    v4 = v4
                }
                v8 += v8
                _ = len(v3)
            }
            for fi16 := 0; fi16 < 19; fi16++ {
                v17 := GenS1{F0: -1, F1: 1, F2: 4294967295}
                _ = fi16
                _ = v17
            }
        }
        fmt.Println(v1, v4, v14)
    }
    if (("hello" + "foo bar") != "a") {
        var v18 uint = 4294967295
        v19 := GenS1{F0: int64(v8), F1: 42, F2: 1}
        for fi20 := 0; fi20 < 10; fi20++ {
            var v21 uint8 = 149
            switch (fi20 * -418) {
            case -2147483648:
                v4 = (-(v4))
            case -415:
                for fi22 := 0; fi22 < 3; fi22++ {
                    if (v18 > (v18 * v6)) {
                        fmt.Println(fi22)
                        var v23 uint16 = v8
                        _ = v23
                    }
                    var v24 float64 = (3.14159 * (float64(1e10) * ((float64(0.0) + v2) + v2)))
                    _ = fi22
                    _ = v24
                }
            }
            var v25 float64 = -2.25
            _ = fi20
            _ = v21
            _ = v25
        }
        _ = v18
        _ = v19
    }
    v6 = (v6 | v6)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
