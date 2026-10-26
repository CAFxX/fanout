package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32 }
func genFn0() bool {
    switch -9223372036854775808 {
    case 2147483647:
        if (false && (true && true)) {
            var v6 uint16 = 414
            _ = v6
        }
    case 9223372036854775807:
        var v7 int8 = 127
        _ = v7
    default:
        for fi8 := 0; fi8 < 6; fi8++ {
            v9 := []int{fi8, (-(fi8)), fi8, (fi8 & fi8)}
            for fi10 := 0; fi10 < 9; fi10++ {
                if (-128 <= -128) {
                    v11 := GenS1{F0: true}
                    var v12 int16 = int16(fi10)
                    fmt.Println(fi8)
                    _ = v11
                    _ = v12
                } else {
                    switch "a" {
                    case "\n\t":
                        if true {
                            var v13 int64 = 42
                            if true {
                                var v14 string = "foo bar"
                                _ = v14
                            } else {
                                var v15 uint32 = 2147483647
                                v13 -= (v13 * (v13 + (v13 ^ (v13 % (((^(v13))) | 1)))))
                                _ = v15
                            }
                            _ = v13
                        } else {
                            fmt.Println(fi8, fi8)
                            _ = len(v9)
                            fmt.Println(fi10, fi10)
                        }
                    case "foo bar":
                        fmt.Println(fi8)
                    }
                }
                v16 := GenS1{F0: true}
                var v17 int8 = -16
                if true {
                    var v18 uint16 = 42
                    v17 += (v17 % ((v17) | 1))
                    _ = v18
                }
                _ = fi10
                _ = v16
                _ = v17
            }
            _ = fi8
            _ = v9
        }
    }
    v19 := []int{516, 0}
    _ = v19
    return (127 == 255)
}

func main() {
    var v1 bool = true
    var v2 int64 = 886
    v3 := make(map[int]bool)
    var v4 string = "\n\t"
    var v5 bool = (42 != 42)
    fmt.Println(v1, v5, v5)
    var v20 uint8 = 2
    v21 := []int{}
    for fi22 := 0; fi22 < 8; fi22++ {
        fmt.Println(v2)
        v5 = (!(2 <= 255))
        _ = fi22
    }
    var v23 float32 = (float32((1e10 / (1e10 * 0.1))) * 0.1)
    switch 1 {
    case 32768:
        v24 := GenS2{F0: 415}
        _ = v24
    }
    var v25 uint64 = 0
    v25 = v25
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
