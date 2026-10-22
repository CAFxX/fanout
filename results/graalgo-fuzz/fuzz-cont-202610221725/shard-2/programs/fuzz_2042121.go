package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 uint64 }

func main() {
    var v1 uint32 = 0
    v2 := make(map[string]int)
    v3 := []int{1, 4294967295, 778}
    var v4 int = -175
    switch v4 {
    case 127:
        if true {
            switch ("foo bar" + ("foo bar" + "\n\t")) {
            case "\n\t":
                for fi5 := 0; fi5 < 9; fi5++ {
                    for fi6 := 0; fi6 < 15; fi6++ {
                        v7 := GenS1{F0: 221}
                        if true {
                            v2["\n\t"] = fi6
                        } else {
                            var v8 uint64 = 815
                            v4 -= v4
                            var v9 bool = ((v8 > (v8 >> ((0) & 63))) && (42 >= -9223372036854775808))
                            _ = v8
                            _ = v9
                        }
                        _ = fi6
                        _ = v7
                    }
                    if (((1e10 + (0.0 + float32((0.0 / 1.0)))) * 1.5) > 0.1) {
                        v10 := GenS1{F0: (v4 & v4)}
                        _ = v10
                    }
                    v1 = (^(v1))
                    _ = fi5
                }
            default:
                v11 := v2[""]
                _ = v11
            }
            var v12 uint = 328
            _ = v12
        }
    case 2147483647:
        var v13 uint = 65535
        _ = v13
    }
    fmt.Println(v1)
    if len(v3) > 0 {
        v3[0] = (^(v4))
    }
    var v14 string = "foo bar"
    fmt.Println(v1, v14, v1)
    if (!false) {
        _ = len(v3)
    }
    if true {
        for fi15 := 0; fi15 < 13; fi15++ {
            var v16 uint8 = 42
            v16 += 42
            _ = fi15
            _ = v16
        }
        var v17 string = v14
        _ = v17
    }
    v2["hello"] = (v4 / ((-2147483648) | 1))
    for fi18 := 0; fi18 < 7; fi18++ {
        fmt.Println(v4, v4, v1)
        _ = len(v2)
        _ = fi18
    }
    v19 := GenS2{F0: 127, F1: 42}
    v1 -= v1
    var v20 uint64 = 664
    var v21 int16 = -207
    var v22 int32 = -2147483648
    fmt.Println(v21, v14)
    if ((v22 * v22) > 1) {
        fmt.Println(v14)
        if false {
            if false {
                v21 = -907
            }
            v21 += (^(v21))
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
