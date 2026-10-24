package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int16 }

func main() {
    v1 := []int{-9223372036854775808, 2147483647, 66}
    var v2 uint32 = 1
    v3 := make(map[uint32]bool)
    fmt.Println(v2)
    v2 += 42
    _ = len(v1)
    for fi4 := 0; fi4 < 13; fi4++ {
        fmt.Println(v2)
        v2 = (v2 * v2)
        _ = fi4
    }
    if (!(-56 < 0)) {
        var v5 uint32 = (v2 << ((65535) & 63))
        for fi6 := 0; fi6 < 3; fi6++ {
            var v7 int32 = -1
            delete(v3, v2)
            _ = fi6
            _ = v7
        }
        _ = v5
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        v2 += v2
        _ = fi8
    }
    var v9 uint16 = 351
    var v10 float32 = 1e10
    fmt.Println(v10)
    var v11 float64 = (0.1 / 1.5)
    for fi12 := 0; fi12 < 9; fi12++ {
        fmt.Println(v11, v10, fi12)
        if false {
            v2 = (v2 | 65535)
        } else {
            if false {
                v13 := GenS1{F0: 0, F1: (127 >= 0), F2: 1}
                if (false && ((!false) || true)) {
                    if (4294967295 < 594) {
                        v14 := GenS1{F0: -83, F1: true, F2: 1}
                        if (374 != 0) {
                            for fi15 := 0; fi15 < 9; fi15++ {
                                delete(v3, v2)
                                _ = fi15
                            }
                            var v16 bool = (false && true)
                            _ = v16
                        } else {
                            var v17 string = ("\n\t" + "")
                            _ = v17
                        }
                        _ = v14
                    }
                }
                _ = v13
            }
            var v18 string = "foo bar"
            var v19 uint = 4294967295
            _ = v18
            _ = v19
        }
        if false {
            _ = len(v1)
        }
        _ = fi12
    }
    var v20 uint16 = (^(v9))
    v10 += v10
    fmt.Println(v9, v11)
    var v21 uint = 18446744073709551615
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v9
}
