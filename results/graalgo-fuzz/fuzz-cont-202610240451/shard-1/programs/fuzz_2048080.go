package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint64 }
type GenS2 struct { F0 uint16; F1 string; F2 uint }
func genFn0(p0 int64) int32 {
    v5 := GenS1{F0: -1, F1: 0, F2: 42}
    var v6 bool = false
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 int64 = (p0 - (p0 >> ((1) & 63)))
        p0 -= (p0 + (v8 / (((v8 >> ((42) & 63))) | 1)))
        _ = fi7
        _ = v8
    }
    v6 = true
    _ = p0
    _ = v5
    _ = v6
    return -536
}

func main() {
    var v1 uint32 = 0
    v2 := []string{}
    var v3 string = "hello"
    var v4 uint = 42
    v2 = append(v2, "\n\t")
    v1 = 2147483647
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v9 := GenS2{F0: 255, F1: v3, F2: v4}
    var v10 string = (((("\n\t" + (v3 + ((((("\n\t" + "\n\t") + (v3 + v3)) + ((("hello" + (v3 + v3)) + (v3 + ("a" + "hello"))) + "hello")) + v3) + v3))) + (v3 + "a")) + v3) + "hello")
    var v11 int = 537
    switch (v11 % ((0) | 1)) {
    case 9223372036854775807:
        if (42 > 255) {
            for fi12 := 0; fi12 < 20; fi12++ {
                v3 = v3
                _ = fi12
            }
            v2 = append(v2, v10)
        }
    case 2147483647:
        for fi13 := 0; fi13 < 11; fi13++ {
            var v14 uint32 = v1
            _ = fi13
            _ = v14
        }
    case 1:
        for fi15 := 0; fi15 < 18; fi15++ {
            v10 = v10
            v11 += int(v4)
            _ = len(v2)
            _ = fi15
        }
    default:
        switch v3 {
        case "\n\t":
            if len(v2) > 0 {
                v2[0] = v3
            }
        default:
            if (!(v4 < v4)) {
                if (!(!(((0.0 / 3.14159) / 1.5) > (3.14159 * 1e10)))) {
                    switch "\n\t" {
                    case "foo bar":
                        switch v3 {
                        case "a":
                            v2 = append(v2, v3)
                        default:
                            v1 -= 1
                        }
                    case "\n\t":
                        var v16 uint16 = 798
                        _ = v16
                    default:
                        fmt.Println(v3)
                    }
                } else {
                    for fi17 := 0; fi17 < 14; fi17++ {
                        if len(v2) > 0 {
                            v2[0] = v3
                        }
                        v4 = (v4 % ((v4) | 1))
                        _ = fi17
                    }
                    v2 = append(v2, "foo bar")
                }
            }
        }
    }
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
