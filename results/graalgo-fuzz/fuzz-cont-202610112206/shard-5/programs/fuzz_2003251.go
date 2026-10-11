package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
func genFn0(p0 uint8) int16 {
    v4 := []int{4294967295, 4294967295, -502, -2147483648}
    p0 += 1
    fmt.Println(p0)
    _ = p0
    _ = v4
    return -549
}

func main() {
    v1 := []int{}
    var v2 float64 = 0.0
    var v3 uint64 = 241
    var v5 int = 466
    v3 -= (^(v3))
    v6 := GenS1{F0: "foo bar", F1: 778}
    v5 -= 9223372036854775807
    if ((false && false) || true) {
        var v7 uint = 633
        v7 += (v7 - (v7 - v7))
        if len(v1) > 0 {
            v1[0] = v5
        }
        _ = v7
    }
    var v8 bool = ((100 <= 0) || false)
    v8 = false
    if true {
        var v9 int16 = 42
        _ = v9
    } else {
        var v10 float64 = 1e10
        for fi11 := 0; fi11 < 13; fi11++ {
            if true {
                for fi12 := 0; fi12 < 19; fi12++ {
                    v13 := GenS1{F0: ("" + "a"), F1: 0}
                    _ = fi12
                    _ = v13
                }
                var v14 uint32 = 2
                _ = v14
            }
            if (386 == 255) {
                if len(v1) > 0 {
                    v1[0] = v5
                }
            } else {
                var v15 int32 = 42
                _ = v15
            }
            if (!false) {
                if false {
                    v2 = (3.14159 - 1.5)
                    var v16 uint8 = 255
                    var v17 uint8 = v16
                    _ = v16
                    _ = v17
                } else {
                    if true {
                        switch (fi11 >> ((42) & 63)) {
                        case -1:
                            v3 = (v3 << ((65535) & 63))
                        case 1:
                            v2 = (3.14159 * ((v10 + ((0.0 * v2) - v2)) * -2.25))
                        case 4294967295:
                            for fi18 := 0; fi18 < 19; fi18++ {
                                switch "foo bar" {
                                case "":
                                    v8 = v8
                                case "foo bar":
                                    v10 -= -2.25
                                    fmt.Println(fi11, v5, v2)
                                }
                                fmt.Println(v8)
                                _ = fi18
                            }
                        }
                        var v19 float64 = ((v2 - v10) * 1e10)
                        _ = len(v1)
                        _ = v19
                    } else {
                        v1 = append(v1, v5)
                    }
                }
                v5 -= 0
                v1 = append(v1, v5)
            }
            _ = fi11
        }
        _ = v10
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v8
}
