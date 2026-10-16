package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 uint32 = 958
    var v2 uint = 1
    var v3 string = ("foo bar" + "foo bar")
    var v4 uint64 = 791
    v4 = 4294967295
    fmt.Println(v1)
    v1 = 169
    v4 -= (v4 / ((18446744073709551615) | 1))
    if (!false) {
        var v5 uint = 18446744073709551615
        v6 := []int{-9223372036854775808, -716}
        _ = v5
        _ = v6
    } else {
        if ((v3 + "a") <= v3) {
            for fi7 := 0; fi7 < 9; fi7++ {
                if (false || false) {
                    v8 := []int{-824, fi7, 42, (fi7 << ((65535) & 63))}
                    v2 = 0
                    if (-68 > 1) {
                        v4 = (v4 / ((0) | 1))
                        if (!true) {
                            if (false || true) {
                                _ = len(v8)
                                if len(v8) > 0 {
                                    v8[0] = fi7
                                }
                            } else {
                                for fi9 := 0; fi9 < 7; fi9++ {
                                    v2 -= 2
                                    fmt.Println(fi9)
                                    _ = fi9
                                }
                            }
                        }
                        fmt.Println(v1)
                    } else {
                        var v10 string = v3
                        _ = v10
                    }
                    _ = v8
                }
                _ = fi7
            }
            var v11 int64 = -1
            _ = v11
        }
        for fi12 := 0; fi12 < 11; fi12++ {
            if (144 > 1) {
                v13 := []int{(fi12 - fi12), fi12, (fi12 + (fi12 % ((fi12) | 1)))}
                _ = v13
            }
            _ = fi12
        }
        fmt.Println(v2, v1)
    }
    v2 = 1
    v14 := GenS1{F0: v3}
    if (false && ((v4 <= v4) && true)) {
        var v15 float32 = 1e10
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
