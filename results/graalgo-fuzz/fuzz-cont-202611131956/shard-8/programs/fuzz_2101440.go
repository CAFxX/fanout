package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := []string{"foo bar"}
    v2 := make(map[int]string)
    var v3 uint = 77
    var v4 int32 = 887
    var v5 int64 = -1
    var v6 bool = true
    switch 488 {
    case 2147483647:
        v4 += v4
    }
    v3 = (v3 & (v3 | (v3 | (^(v3)))))
    switch -9223372036854775808 {
    case 128:
        for fi7 := 0; fi7 < 12; fi7++ {
            v8 := GenS1{F0: ("hello" + ""), F1: (v5 & (-(v5))), F2: 0}
            _ = fi7
            _ = v8
        }
    case 42:
        var v9 int = 0
        _ = v9
    case -2147483648:
        v6 = v6
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 bool = (true || true)
            switch (-(fi11)) {
            case 32768:
                v13 := v2[310]
                _ = v13
            case -1:
                var v14 float64 = 3.14159
                _ = v14
            }
            if len(v1) > 0 {
                v1[0] = "\n\t"
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v15 := GenS1{F0: "", F1: (v5 << ((185) & 63)), F2: 255}
            _ = fi11
            _ = v12
            _ = v15
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v16 int8 = -1
        _ = fi10
        _ = v16
    }
    var v17 int32 = 0
    var v18 uint = (v3 | (v3 ^ (v3 - v3)))
    v19 := GenS2{F0: 509}
    if ((((4294967295 < 944) || (v3 > 42)) && false) || ((v5 - v5) == (v5 % ((v5) | 1)))) {
        if len(v1) > 0 {
            v1[0] = ("\n\t" + "")
        }
    }
    switch 2147483647 {
    case -9223372036854775808:
        if v6 {
            switch -743 {
            case 0:
                v4 = v4
            case 386:
                switch "foo bar" {
                case "foo bar":
                    if ((4294967295 < 2) || ((!(v6 && v6)) && v6)) {
                        for fi20 := 0; fi20 < 17; fi20++ {
                            var v21 uint8 = 101
                            _ = fi20
                            _ = v21
                        }
                        v2[-2147483648] = "hello"
                    }
                }
            default:
                fmt.Println(v5)
            }
            var v22 uint32 = 42
            fmt.Println(v6, v4, v22)
            _ = v22
        }
    default:
        switch 9223372036854775807 {
        case 628:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        default:
            var v23 uint8 = 100
            _ = v23
        }
    }
    var v24 uint32 = 65535
    v3 -= 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
