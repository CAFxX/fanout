package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0() int {
    v7 := GenS1{F0: -9223372036854775808}
    switch "foo bar" {
    case "hello":
        for fi8 := 0; fi8 < 5; fi8++ {
            var v9 int16 = 74
            var v10 bool = false
            _ = fi8
            _ = v10
            _ = v9
        }
    case "\n\t":
        var v11 uint = 1
        _ = v11
    }
    _ = v7
    return 32768
}

func main() {
    v1 := []uint32{42, 1, 42, 65535}
    v2 := make(map[int]uint32)
    var v3 int16 = 0
    v4 := make(map[int]int)
    v5 := []string{}
    v6 := make(map[string]int)
    var v12 float64 = 1.5
    if false {
        v13 := GenS1{F0: 63}
        _ = v13
    }
    v5 = append(v5, "")
    var v14 uint64 = 0
    var v15 int64 = 0
    fmt.Println(v14)
    for fi16 := 0; fi16 < 8; fi16++ {
        if ((2 == 1) && (18446744073709551615 < 2)) {
            if false {
                if true {
                    for fi17 := 0; fi17 < 16; fi17++ {
                        fmt.Println(fi17)
                        _ = fi17
                    }
                    var v18 uint16 = 42
                    _ = v18
                }
                v3 = (v3 / ((0) | 1))
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if true {
                    v19 := v4[(fi16 * int(v14))]
                    _ = v19
                }
            }
        } else {
            fmt.Println(fi16)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v20 := GenS1{F0: v15}
            _ = v20
        }
        v14 += v14
        _ = fi16
    }
    v21 := v2[-2147483648]
    fmt.Println(v14)
    _ = len(v1)
    v22 := GenS1{F0: v15}
    var v23 float64 = v12
    v23 -= 1e10
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
