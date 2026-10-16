package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := []string{"", "\n\t", "foo bar"}
    v2 := []string{"foo bar", ("foo bar" + ("" + "\n\t")), "hello", "hello"}
    var v3 uint8 = 42
    var v4 uint32 = 231
    var v5 float64 = (0.0 + ((1e10 - (0.0 - 0.0)) + 0.0))
    fmt.Println(v5, v5)
    switch -517 {
    case 0:
        v5 += v5
    default:
        if len(v2) > 0 {
            v2[0] = "foo bar"
        }
    }
    switch "foo bar" {
    case "":
        v6 := GenS1{F0: 9223372036854775807}
        _ = v6
    case "\n\t":
        if len(v1) > 0 {
            v1[0] = "\n\t"
        }
    case "hello":
        for fi7 := 0; fi7 < 18; fi7++ {
            for fi8 := 0; fi8 < 17; fi8++ {
                v5 = v5
                v4 -= 0
                var v9 uint32 = v4
                switch ("\n\t" + "foo bar") {
                case "\n\t":
                    _ = len(v2)
                case "hello":
                    var v10 uint64 = 42
                    _ = v10
                case "":
                    v4 = 42
                default:
                    fmt.Println(fi7, v5)
                }
                _ = fi8
                _ = v9
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi7
        }
    }
    var v11 uint64 = uint64(v4)
    for fi12 := 0; fi12 < 5; fi12++ {
        v5 -= 1.5
        v2 = append(v2, ("a" + "hello"))
        _ = fi12
    }
    v4 -= 65535
    for fi13 := 0; fi13 < 20; fi13++ {
        var v14 int64 = -198
        var v15 uint64 = 42
        if (2147483647 == (^(fi13))) {
            var v16 int8 = 1
            var v17 int8 = 1
            for fi18 := 0; fi18 < 1; fi18++ {
                var v19 float64 = ((float64(1e10) - float64(-2.25)) / float64(3.14159))
                _ = fi18
                _ = v19
            }
            _ = v16
            _ = v17
        } else {
            for fi20 := 0; fi20 < 14; fi20++ {
                var v21 float32 = 3.14159
                switch ("" + (("a" + ("a" + "hello")) + "")) {
                case "\n\t":
                    var v22 uint = 65535
                    _ = v22
                default:
                    v5 = (0.0 * v5)
                }
                var v23 bool = (v14 > 183)
                _ = fi20
                _ = v21
                _ = v23
            }
            var v24 uint64 = (v11 * v11)
            _ = v24
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    v25 := GenS1{F0: 0}
    fmt.Println(v3, v5, v4)
    v26 := GenS1{F0: -9223372036854775808}
    v3 = uint8(v11)
    v11 = (v11 % ((v11) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
