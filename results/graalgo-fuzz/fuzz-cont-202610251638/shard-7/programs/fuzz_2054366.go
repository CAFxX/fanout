package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }

func main() {
    var v1 int = -2147483648
    v2 := []int{}
    var v3 int16 = -847
    var v4 uint = 391
    var v5 int16 = (v3 ^ 1000)
    var v6 int64 = 1
    fmt.Println(v6)
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 int = int(v4)
        var v9 uint32 = 4294967295
        for fi10 := 0; fi10 < 10; fi10++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi10
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    var v11 string = ("hello" + "hello")
    fmt.Println(v1)
    for fi12 := 0; fi12 < 4; fi12++ {
        switch 603 {
        case -9223372036854775808:
            if true {
                v6 += -9223372036854775808
                switch (v1 - fi12) {
                case -373:
                    var v13 int64 = (v6 % (((v6 + (v6 ^ (^(v6))))) | 1))
                    _ = v13
                case 0:
                    var v14 int64 = v6
                    _ = v14
                case -652:
                    var v15 uint16 = 815
                    _ = v15
                default:
                    v16 := GenS1{F0: 899, F1: fi12}
                    _ = v16
                }
                var v17 uint64 = 42
                _ = v17
            }
        default:
            fmt.Println(v1, v1, v4)
        }
        _ = fi12
    }
    v18 := GenS1{F0: 0, F1: (^(v1))}
    if false {
        if ((v4 | (v4 | (v4 * (v4 - v4)))) <= 4294967295) {
            var v19 uint16 = 42
            if true {
                for fi20 := 0; fi20 < 4; fi20++ {
                    fmt.Println(v4, fi20)
                    v6 -= -1
                    var v21 bool = (521 == 145)
                    _ = fi20
                    _ = v21
                }
                var v22 int = v1
                _ = v22
            }
            _ = v19
        }
        v3 += (v3 | (v5 & 1))
        var v23 int64 = v6
        _ = v23
    }
    v24 := GenS1{F0: 18446744073709551615, F1: 9223372036854775807}
    v2 = append(v2, v1)
    v6 = v6
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
