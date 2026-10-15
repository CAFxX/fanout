package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 int8; F1 int64; F2 int8 }
func genFn0(p0 string, p1 int16, p2 int16) string {
    v4 := GenS1{F0: 2, F1: "foo bar"}
    v5 := []int{42, 42, 5, 922}
    p1 = (p2 + p1)
    p2 -= (p1 << ((42) & 63))
    _ = len(v5)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return p0
}

func main() {
    v1 := []string{}
    v2 := []string{"hello"}
    v3 := make(map[uint32]bool)
    if (true || ((0 >= 2) && (0 <= 1))) {
        if (!false) {
            var v6 int16 = -32768
            v3[735] = true
            if false {
                v6 -= (^(v6))
                for fi7 := 0; fi7 < 4; fi7++ {
                    v8 := GenS2{F0: 98, F1: 602, F2: -44}
                    for fi9 := 0; fi9 < 3; fi9++ {
                        fmt.Println(fi7, fi7, fi9)
                        _ = fi9
                    }
                    if (127 < 40) {
                        for fi10 := 0; fi10 < 11; fi10++ {
                            fmt.Println(fi7, fi7, fi10)
                            _ = fi10
                        }
                        for fi11 := 0; fi11 < 16; fi11++ {
                            v1 = append(v1, ("a" + ("foo bar" + "")))
                            _ = fi11
                        }
                        v3[0] = true
                    }
                    _ = fi7
                    _ = v8
                }
                for fi12 := 0; fi12 < 7; fi12++ {
                    v6 = v6
                    v3[503] = false
                    _ = fi12
                }
            } else {
                v13 := v3[948]
                _ = v13
            }
            _ = v6
        } else {
            var v14 int = 128
            v14 += (v14 & (-(v14)))
            _ = v14
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    _ = len(v3)
    var v15 uint8 = 1
    fmt.Println(v15)
    _ = len(v2)
    v15 += (v15 & (^(v15)))
    v16 := v3[0]
    var v17 uint8 = 2
    v17 += v17
    var v18 uint32 = 2147483647
    v18 += (v18 >> ((2) & 63))
    v18 -= (^(v18))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
