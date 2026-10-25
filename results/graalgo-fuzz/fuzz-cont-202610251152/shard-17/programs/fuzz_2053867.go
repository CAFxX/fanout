package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint8; F1 string; F2 string }
func genFn0() int {
    var v6 int8 = -1
    v7 := []int{0}
    v8 := GenS1{F0: 18446744073709551615, F1: 0}
    v9 := GenS1{F0: 1, F1: 39}
    var v10 uint64 = 2
    _ = v10
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 309
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = 0.0
    var v3 uint16 = 0
    v4 := make(map[string]string)
    var v5 uint8 = 127
    v5 -= v5
    switch "" {
    case "":
        var v11 bool = (!(27 == 100))
        _ = v11
    case "\n\t":
        v5 += 128
    case "foo bar":
        v12 := GenS1{F0: 4294967295, F1: (v5 % ((v5) | 1))}
        _ = v12
    }
    var v13 uint = 357
    v4["\n\t"] = "hello"
    v14 := GenS1{F0: 18446744073709551615, F1: 128}
    for fi15 := 0; fi15 < 1; fi15++ {
        if false {
            v3 = v3
            if false {
                v16 := []int{(fi15 - (fi15 & fi15)), (fi15 ^ fi15), (fi15 % ((fi15) | 1)), (^(fi15))}
                if true {
                    v17 := GenS1{F0: 18446744073709551615, F1: (v5 + v5)}
                    _ = v17
                }
                var v18 int64 = -360
                _ = v16
                _ = v18
            }
        }
        v13 = v13
        _ = fi15
    }
    var v19 int = 42
    v3 -= v3
    v20 := GenS2{F0: 2, F1: "\n\t", F2: "a"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
