package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{"a"}
    var v3 int32 = 42
    var v4 int8 = 1
    v5 := GenS1{F0: 42, F1: 2}
    var v6 bool = true
    fmt.Println(v4, v4, v3)
    var v7 int64 = -1
    var v8 bool = true
    var v9 uint16 = 65535
    v2 = append(v2, "hello")
    if v8 {
        v7 = (v7 & (v7 * -645))
        v10 := GenS1{F0: 18446744073709551615, F1: 1}
        _ = v10
    } else {
        _ = len(v2)
        v4 -= (^(v4))
    }
    if v8 {
        var v11 uint = 4294967295
        if (1.5 < 1e10) {
            if (("a" + "hello") == "\n\t") {
                if v8 {
                    var v12 uint = (v11 % (((v11 * (v11 + 4294967295))) | 1))
                    _ = v12
                }
                if v6 {
                    v8 = false
                    var v13 uint16 = 65535
                    var v14 float32 = float32(-2.25)
                    _ = v13
                    _ = v14
                }
            } else {
                fmt.Println(v9, v8, v8)
            }
        } else {
            var v15 uint16 = 398
            for fi16 := 0; fi16 < 10; fi16++ {
                for fi17 := 0; fi17 < 16; fi17++ {
                    var v18 int = (^(fi17))
                    _ = fi17
                    _ = v18
                }
                var v19 uint32 = 4294967295
                if v6 {
                    v7 = (v7 / ((42) | 1))
                    for fi20 := 0; fi20 < 15; fi20++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v21 string = "foo bar"
                        _ = fi20
                        _ = v21
                    }
                    if (float64(0.0) < (0.0 * -2.25)) {
                        fmt.Println(v11)
                    }
                }
                _ = fi16
                _ = v19
            }
            _ = v15
        }
        _ = v11
    }
    var v22 string = "\n\t"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
