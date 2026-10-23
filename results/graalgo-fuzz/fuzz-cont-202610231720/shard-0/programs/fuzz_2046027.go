package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }

func main() {
    var v1 int8 = 28
    var v2 string = ""
    var v3 uint32 = 2147483647
    var v4 bool = (!((!(!(!true))) && (((1 != 42) || true) && true)))
    v5 := GenS1{F0: v4, F1: 2}
    var v6 float64 = -2.25
    var v7 uint = 4294967295
    fmt.Println(v1, v4)
    if v4 {
        var v8 int32 = 2147483647
        for fi9 := 0; fi9 < 17; fi9++ {
            var v10 uint = v7
            for fi11 := 0; fi11 < 3; fi11++ {
                v10 += (^(v7))
                _ = fi11
            }
            fmt.Println(v2, v4, v6)
            _ = fi9
            _ = v10
        }
        fmt.Println(v7, v7, v7)
        _ = v8
    } else {
        v12 := []int{0, -1, 35, 0}
        _ = v12
    }
    v13 := []int{}
    for fi14 := 0; fi14 < 9; fi14++ {
        if ((fi14 % (((fi14 - fi14)) | 1)) != 32767) {
            fmt.Println(v6, v2)
        }
        if len(v13) > 0 {
            v13[0] = (fi14 % ((998) | 1))
        }
        _ = len(v13)
        _ = fi14
    }
    if (32767 >= 1) {
        for fi15 := 0; fi15 < 14; fi15++ {
            if len(v13) > 0 {
                v13[0] = fi15
            }
            switch v2 {
            case "\n\t":
                var v16 uint32 = v3
                _ = v16
            case "hello":
                if v4 {
                    v13 = append(v13, (fi15 * fi15))
                    var v17 int8 = (v1 * v1)
                    if ((v4 || v4) && (!v4)) {
                        _ = len(v13)
                    } else {
                        for fi18 := 0; fi18 < 14; fi18++ {
                            v19 := GenS1{F0: true, F1: 32767}
                            if len(v13) >= 2 {
                                v13 = v13[1:len(v13)]
                            }
                            _ = fi18
                            _ = v19
                        }
                    }
                    _ = v17
                }
            case "a":
                v2 = v2
            }
            _ = fi15
        }
        v3 = 817
    }
    var v20 bool = (float32(((1.5 + v6) / -2.25)) < ((0.0 / float32(v6)) - ((0.0 - float32(((v6 + 1e10) - 1.5))) + 3.14159)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
