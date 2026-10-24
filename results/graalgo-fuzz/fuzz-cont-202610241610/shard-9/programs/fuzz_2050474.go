package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }

func main() {
    var v1 uint8 = 2
    v2 := make(map[string]int)
    var v3 uint16 = 42
    var v4 bool = (int(v1) == -916)
    v3 -= (v3 & v3)
    v5 := []int{}
    var v6 int8 = -1
    fmt.Println(v1, v3, v3)
    var v7 uint = 2
    var v8 uint8 = (v1 * (^(v1)))
    var v9 uint = 0
    switch "foo bar" {
    case "":
        fmt.Println(v8, v1, v9)
    default:
        fmt.Println(v9, v8)
    }
    if v4 {
        var v10 float32 = 1e10
        var v11 uint64 = 722
        delete(v2, "\n\t")
        _ = v10
        _ = v11
    } else {
        v9 -= (v9 ^ 255)
        fmt.Println(v3, v6, v6)
    }
    if (132 <= 42) {
        v9 += v9
        v3 = (v3 / (((v3 * (v3 - uint16(v9)))) | 1))
        for fi12 := 0; fi12 < 6; fi12++ {
            v13 := GenS1{F0: 18446744073709551615, F1: 1}
            v8 -= 1
            _ = fi12
            _ = v13
        }
    } else {
        var v14 bool = v4
        if (304 > 42) {
            v15 := GenS1{F0: 952, F1: 895}
            var v16 string = ((("foo bar" + "a") + "foo bar") + "")
            _ = v15
            _ = v16
        } else {
            if (77 >= (v1 % ((v1) | 1))) {
                for fi17 := 0; fi17 < 14; fi17++ {
                    fmt.Println(v9, v1, v14)
                    var v18 int32 = 42
                    _ = fi17
                    _ = v18
                }
                v19 := v2["\n\t"]
                _ = v19
            } else {
                if false {
                    var v20 uint32 = 2
                    _ = v20
                } else {
                    for fi21 := 0; fi21 < 16; fi21++ {
                        var v22 int = -872
                        v8 -= (v8 * 1)
                        _ = fi21
                        _ = v22
                    }
                    fmt.Println(v8)
                    switch "" {
                    case "":
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                }
                v3 += v3
            }
            var v23 bool = v4
            _ = v23
        }
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
