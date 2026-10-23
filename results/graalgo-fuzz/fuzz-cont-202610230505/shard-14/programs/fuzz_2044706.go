package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int32 }

func main() {
    var v1 bool = (("" + (("foo bar" + (((((("\n\t" + ("foo bar" + "a")) + "foo bar") + ("\n\t" + "hello")) + "foo bar") + "") + "hello")) + "foo bar")) <= "\n\t")
    v2 := make(map[int]uint32)
    var v3 uint64 = 0
    var v4 uint64 = (v3 / ((v3) | 1))
    v5 := []uint32{uint32(v4), 732}
    var v6 bool = true
    v7 := v2[32768]
    var v8 int8 = 42
    if v6 {
        for fi9 := 0; fi9 < 3; fi9++ {
            if true {
                delete(v2, (-(fi9)))
                v10 := GenS1{F0: 259, F1: 428, F2: -2147483648}
                for fi11 := 0; fi11 < 14; fi11++ {
                    v5 = append(v5, (^(v7)))
                    var v12 bool = false
                    v7 = 1
                    _ = fi11
                    _ = v12
                }
                _ = v10
            } else {
                var v13 uint32 = (v7 / (((v7 | (^(v7)))) | 1))
                v14 := GenS1{F0: 1, F1: 0, F2: -270}
                _ = v13
                _ = v14
            }
            fmt.Println(v6, v7, v4)
            v3 -= (v4 / ((1) | 1))
            _ = fi9
        }
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        v1 = false
        _ = fi15
    }
    delete(v2, 42)
    var v16 int32 = 42
    delete(v2, 134)
    var v17 string = "hello"
    v2[int(v8)] = v7
    v18 := v2[-621]
    if true {
        v5 = append(v5, (^(v7)))
    }
    v4 = (v4 * (v3 | (v3 | 2)))
    if len(v5) > 0 {
        v5[0] = 2147483647
    }
    var v19 int64 = 442
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
