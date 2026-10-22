package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint64 }
type GenS2 struct { F0 int64; F1 uint32 }

func main() {
    v1 := []string{"", ((("" + ("foo bar" + ("a" + ("hello" + "\n\t")))) + "foo bar") + "\n\t"), "foo bar"}
    var v2 uint64 = 0
    var v3 int64 = 1
    var v4 uint32 = 4294967295
    var v5 uint = 820
    v6 := make(map[string]int)
    fmt.Println(v3, v2, v3)
    if (false || (319 <= 863)) {
        _ = len(v1)
        for fi7 := 0; fi7 < 4; fi7++ {
            if false {
                _ = len(v1)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            } else {
                v4 = v4
                for fi8 := 0; fi8 < 3; fi8++ {
                    v2 = (v2 + (v2 + v2))
                    _ = fi8
                }
            }
            fmt.Println(v2, v4, v5)
            _ = len(v1)
            _ = fi7
        }
        _ = len(v1)
    }
    var v9 int = -2147483648
    v10 := GenS1{F0: 127, F1: "", F2: (v2 >> ((v5) & 63))}
    v9 -= (v9 * -194)
    v6[("foo bar" + "\n\t")] = int(v5)
    var v11 int = (-(v9))
    switch v11 {
    case -9223372036854775808:
        var v12 int16 = 996
        _ = v12
    case -2147483648:
        v2 -= v2
    default:
        v4 = 4294967295
    }
    v5 = (v5 / (((v5 + 488)) | 1))
    if len(v1) > 0 {
        v1[0] = ("\n\t" + "a")
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 += (v2 | v2)
    for fi13 := 0; fi13 < 15; fi13++ {
        switch v9 {
        case 9223372036854775807:
            v3 -= (v3 / (((v3 / ((-9223372036854775808) | 1))) | 1))
        }
        v14 := GenS2{F0: 9223372036854775807, F1: (v4 % ((497) | 1))}
        _ = fi13
        _ = v14
    }
    var v15 bool = false
    var v16 int16 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
