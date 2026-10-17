package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int64 }

func main() {
    var v1 bool = (!(false || (true && false)))
    v2 := []int{}
    var v3 int64 = -1
    var v4 float64 = 1.5
    var v5 uint8 = 47
    for fi6 := 0; fi6 < 5; fi6++ {
        v2 = append(v2, fi6)
        _ = fi6
    }
    v4 += (((3.14159 - 1e10) - 3.14159) + v4)
    if len(v2) > 0 {
        v2[0] = 9223372036854775807
    }
    v5 = v5
    if (v3 != (v3 >> ((42) & 63))) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if v1 {
            _ = len(v2)
            v7 := GenS1{F0: "", F1: 2, F2: 9223372036854775807}
            _ = v7
        } else {
            if v1 {
                if (84 >= 1) {
                    for fi8 := 0; fi8 < 11; fi8++ {
                        v3 = 492
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v9 int64 = (v3 & -256)
                        _ = fi8
                        _ = v9
                    }
                    var v10 uint = 0
                    _ = v10
                } else {
                    var v11 uint32 = 2147483647
                    _ = v11
                }
            } else {
                fmt.Println(v4)
                var v12 int16 = 42
                _ = v12
            }
            v13 := GenS1{F0: "foo bar", F1: 403, F2: v3}
            _ = v13
        }
    }
    fmt.Println(v4)
    v5 -= v5
    v14 := GenS1{F0: (("a" + "") + ""), F1: 422, F2: v3}
    _ = len(v2)
    if (!((42 <= 42) || (33 != 536))) {
        v2 = append(v2, 0)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
