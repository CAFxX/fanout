package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint }

func main() {
    var v1 int = 9223372036854775807
    v2 := make(map[int]bool)
    v3 := []int{v1, (v1 * -960), (v1 | 4294967295)}
    var v4 int64 = 142
    fmt.Println(v4, v4)
    var v5 int = v1
    v5 = (v1 >> ((42) & 63))
    var v6 int32 = -1
    var v7 float32 = 0.0
    v7 = (v7 + float32((float64(v7) + 1.5)))
    if len(v3) > 0 {
        v3[0] = v1
    }
    if false {
        v8 := GenS1{F0: 42, F1: (v1 * v5), F2: 536}
        _ = v8
    }
    switch ("a" + "foo bar") {
    case "a":
        v9 := v2[(-(v1))]
        _ = v9
    case "\n\t":
        v6 = int32(v4)
    default:
        switch "foo bar" {
        case "\n\t":
            v10 := v2[(v5 << ((113) & 63))]
            _ = v10
        }
    }
    if ((v1 + (v5 + v5)) == -2147483648) {
        v5 = (v5 & -288)
        for fi11 := 0; fi11 < 11; fi11++ {
            fmt.Println(v4, v7, v5)
            switch -637 {
            case -1:
                switch fi11 {
                case -169:
                    v4 -= (^(v4))
                default:
                    switch (v1 ^ (^(v5))) {
                    case -1:
                        var v12 uint8 = 2
                        _ = v12
                    case 4294967295:
                        for fi13 := 0; fi13 < 3; fi13++ {
                            v5 -= fi13
                            _ = fi13
                        }
                    case 285:
                        v3 = append(v3, v5)
                    }
                }
            case 1:
                for fi14 := 0; fi14 < 7; fi14++ {
                    var v15 uint8 = 128
                    _ = fi14
                    _ = v15
                }
            default:
                var v16 float64 = (1e10 / 1.0)
                _ = v16
            }
            if ("foo bar" <= ("hello" + "hello")) {
                v7 -= v7
            } else {
                _ = len(v3)
                if ("a" == "a") {
                    fmt.Println(v6, v5, fi11)
                    var v17 int = v5
                    _ = v17
                }
                _ = len(v2)
            }
            _ = fi11
        }
    }
    var v18 float32 = v7
    for fi19 := 0; fi19 < 4; fi19++ {
        if (v4 == 1) {
            for fi20 := 0; fi20 < 10; fi20++ {
                if true {
                    fmt.Println(v6)
                } else {
                    v21 := GenS1{F0: 2, F1: (v1 >> ((18446744073709551615) & 63)), F2: 65535}
                    var v22 uint = 18446744073709551615
                    if false {
                        var v23 float64 = -2.25
                        var v24 int8 = 0
                        var v25 int8 = v24
                        _ = v23
                        _ = v24
                        _ = v25
                    }
                    _ = v21
                    _ = v22
                }
                _ = fi20
            }
        } else {
            fmt.Println(v6)
        }
        _ = fi19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
