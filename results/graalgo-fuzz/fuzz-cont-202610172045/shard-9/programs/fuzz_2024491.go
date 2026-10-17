package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 bool) int32 {
    var v7 int64 = -33
    v8 := GenS1{F0: p0}
    v7 += (v7 & v7)
    _ = p0
    _ = v7
    _ = v8
    return 395
}

func main() {
    var v1 bool = false
    v2 := make(map[uint32]bool)
    v3 := make(map[int]string)
    v4 := make(map[string]int)
    var v5 int = 634
    var v6 uint32 = 65535
    for fi9 := 0; fi9 < 15; fi9++ {
        v10 := []int{(v5 ^ v5), 128, v5}
        fmt.Println(v1)
        _ = fi9
        _ = v10
    }
    switch "\n\t" {
    case "":
        if (!false) {
            v11 := GenS1{F0: (-1 > 22)}
            _ = v11
        } else {
            var v12 int8 = -1
            v4["a"] = (v5 & (^(v5)))
            for fi13 := 0; fi13 < 11; fi13++ {
                if true {
                    if (255 >= 199) {
                        if false {
                            v5 += (v5 % ((fi13) | 1))
                        }
                        v12 = 42
                        if false {
                            var v14 bool = false
                            var v15 bool = v1
                            _ = v14
                            _ = v15
                        }
                    }
                    fmt.Println(v6, v1)
                    var v16 uint32 = (^(v6))
                    _ = v16
                } else {
                    switch "foo bar" {
                    case "hello":
                        fmt.Println(v5, fi13)
                    case "":
                        var v17 int64 = -9223372036854775808
                        _ = v17
                    }
                    v1 = (v1 || v1)
                }
                _ = fi13
            }
            _ = v12
        }
    case "foo bar":
        var v18 string = "foo bar"
        _ = v18
    }
    v19 := GenS1{F0: (v1 && ((1.5 * 0.1) >= 3.14159))}
    v5 = v5
    switch "a" {
    case "foo bar":
        for fi20 := 0; fi20 < 12; fi20++ {
            var v21 uint32 = v6
            var v22 float64 = (3.14159 - float64((float32(-2.25) / 0.1)))
            if true {
                v23 := GenS1{F0: false}
                for fi24 := 0; fi24 < 10; fi24++ {
                    v25 := GenS1{F0: (false && (-1 >= -1))}
                    var v26 int8 = 0
                    v27 := v3[32768]
                    _ = fi24
                    _ = v25
                    _ = v26
                    _ = v27
                }
                _ = v23
            }
            _ = fi20
            _ = v21
            _ = v22
        }
    }
    var v28 bool = (!v1)
    var v29 int16 = -1
    v29 = (v29 / (((v29 & (v29 % (((v29 >> ((42) & 63))) | 1)))) | 1))
    var v30 uint64 = 4294967295
    v6 = (v6 % (((v6 >> ((4294967295) & 63))) | 1))
    var v31 int64 = -9223372036854775808
    if false {
        v31 = v31
    }
    var v32 int = 871
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
