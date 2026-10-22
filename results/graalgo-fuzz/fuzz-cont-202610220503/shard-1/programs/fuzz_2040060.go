package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 uint64 }
type GenS2 struct { F0 int32 }
func genFn0() uint8 {
    var v4 string = "foo bar"
    switch -729 {
    case 128:
        var v5 int64 = 42
        _ = v5
    case -1:
        v6 := []int{}
        _ = v6
    }
    v4 = ((v4 + v4) + v4)
    var v7 float32 = 0.0
    for fi8 := 0; fi8 < 2; fi8++ {
        v9 := GenS1{F0: 789, F1: (fi8 - 127), F2: 42}
        _ = fi8
        _ = v9
    }
    _ = v4
    _ = v7
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 42
    var v3 bool = false
    fmt.Println(v3, v2)
    if v3 {
        var v10 int32 = -2147483648
        v10 -= 0
        _ = v10
    }
    _ = len(v1)
    if v3 {
        for fi11 := 0; fi11 < 20; fi11++ {
            v2 += 18446744073709551615
            _ = fi11
        }
        var v12 float32 = 1e10
        _ = v12
    }
    v3 = false
    var v13 uint8 = 2
    var v14 int = 9223372036854775807
    for fi15 := 0; fi15 < 2; fi15++ {
        v16 := []int{(fi15 / ((-390) | 1)), (^(v14)), -1}
        if false {
            for fi17 := 0; fi17 < 20; fi17++ {
                var v18 int8 = 42
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
                switch fi15 {
                case -9223372036854775808:
                    if v3 {
                        v2 -= 2
                        delete(v1, 86)
                        var v19 int8 = v18
                        _ = v19
                    }
                }
                _ = fi17
                _ = v18
            }
            v14 = fi15
        } else {
            if (v2 <= v2) {
                v20 := GenS2{F0: -2147483648}
                v16 = append(v16, 394)
                _ = v20
            }
            if v3 {
                if true {
                    fmt.Println(fi15)
                    var v21 int64 = 1
                    _ = v21
                }
                v22 := GenS1{F0: v14, F1: fi15, F2: v2}
                _ = v22
            } else {
                switch "foo bar" {
                case "foo bar":
                    switch (^(v14)) {
                    case -9223372036854775808:
                        if len(v16) >= 2 {
                            v16 = v16[1:len(v16)]
                        }
                    }
                case "\n\t":
                    v2 = 1
                default:
                    v14 -= -345
                }
                var v23 int8 = -128
                switch v14 {
                case 42:
                    var v24 int64 = int64(v2)
                    _ = v24
                case -1:
                    v25 := GenS1{F0: v14, F1: v14, F2: 856}
                    _ = v25
                case 128:
                    fmt.Println(v14, fi15)
                }
                _ = v23
            }
            v16 = append(v16, fi15)
        }
        _ = fi15
        _ = v16
    }
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
}
