package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 bool) uint64 {
    var v4 string = (("a" + ("hello" + "hello")) + "\n\t")
    p0 = (743 > 2147483648)
    v5 := []int{-9223372036854775808, -743}
    _ = p0
    _ = v4
    _ = v5
    return 18446744073709551615
}
func genFn1(p0 uint64, p1 bool, p2 int8) uint64 {
    p2 = int8(p0)
    v6 := GenS1{F0: (p2 & 33)}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p0
}

func main() {
    var v1 int = -2147483648
    v2 := []string{"", ("foo bar" + ("a" + ("a" + ("a" + ("\n\t" + "\n\t")))))}
    var v3 uint8 = 128
    v3 = (v3 / (((v3 - (v3 << ((uint(v3)) & 63)))) | 1))
    v1 += (^(v1))
    for fi7 := 0; fi7 < 18; fi7++ {
        if len(v2) > 0 {
            v2[0] = ("" + "a")
        }
        var v8 int32 = -198
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 42}
    fmt.Println(v1)
    var v10 uint = 0
    fmt.Println(v1)
    fmt.Println(v1)
    if (!true) {
        for fi11 := 0; fi11 < 15; fi11++ {
            _ = len(v2)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi11
        }
        for fi12 := 0; fi12 < 18; fi12++ {
            if ((float64((1.5 / 1.5)) + 0.1) >= ((0.0 / 1.5) * ((-2.25 + (0.0 / 1.0)) - (0.1 / 3.14159)))) {
                for fi13 := 0; fi13 < 2; fi13++ {
                    _ = len(v2)
                    _ = fi13
                }
            } else {
                v3 = v3
            }
            _ = fi12
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v9
}
